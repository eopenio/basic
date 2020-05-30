package wxclient

import (
	"github.com/eopenio/basic/config"
	"github.com/eopenio/workwx"
	log "github.com/micro/go-micro/v2/logger"
	"sync"
)

var (
	app       *workwx.WorkwxApp
	m         sync.RWMutex
	inited    bool
)

type WxRecipient workwx.Recipient

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
		log.Info("开始初始化Workwx")

		initWorkwx(workwxConfig)

		log.Info("初始化Workwx完成，开始检测连接...")
		app.SpawnAccessTokenRefresher()

		testUser := workwx.Recipient{
			UserIDs: []string{workwxConfig.GetAdminUser()},
		}
		testDesc := "企业微信正在进行初始化测试, 该信息请忽略。"
		err := app.SendTextMessage(&testUser, testDesc, false)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Info("初始化Workwx完成，检测连接完成。")
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

// UserIDs 成员ID列表（消息接收者），最多支持1000个
func SetRecipientToUser(uids []string) WxRecipient {
	return WxRecipient{UserIDs: uids}
}

// PartyIDs 部门ID列表，最多支持100个。
func SetRecipientToDept(depts []string) WxRecipient {
	return WxRecipient{PartyIDs: depts}
}

// TagIDs 标签ID列表，最多支持100个
func SetRecipientToTag(tags []string) WxRecipient {
	return WxRecipient{TagIDs: tags}
}

// ChatID 应用关联群聊ID，仅用于【发送消息到群聊会话】
func SetRecipientToChat(chatid string) WxRecipient {
	return WxRecipient{ChatID: chatid}
}
