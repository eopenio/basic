package wxclient

import (
	"github.com/eopenio/basic/config"
	"github.com/eopenio/workwx"
	log "github.com/micro/go-micro/v2/logger"
	"sync"
)

var (
	app    *workwx.WorkwxApp
	m      sync.RWMutex
	inited bool
)

// Init 初始化Redis
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("已经初始化过Workwx...")
		return
	}

	workwxConfig := config.GetWorkwxConfig()

	// 打开才加载
	if workwxConfig != nil && workwxConfig.GetEnabled() {
		log.Info("初始化Workwx...")

		initWorkwx(workwxConfig)

		log.Info("初始化Workwx，检测连接...")
		app.SpawnAccessTokenRefresher()

		testUser := workwx.Recipient{
			UserIDs: []string{workwxConfig.GetAdminUser()},
		}
		testDesc := "企业微信正在进行初始化测试, 该信息请忽略。"
		err := app.SendTextMessage(&testUser, testDesc, false)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	inited = true
}

// GetWorkwx 获取 wxclient
func GetWorkwx() *workwx.WorkwxApp {
	return app
}

func initWorkwx(workwxConfig config.WorkwxConfig) {
	client := workwx.New(workwxConfig.GetCorpId())
	app = client.WithApp(workwxConfig.GetCorpSecret(), workwxConfig.GetAgentId())
}
