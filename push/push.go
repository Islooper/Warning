package push

import (
	"github.com/CatchZeng/dingtalk"
	_ "github.com/CatchZeng/dingtalk"
)

type Writer interface {
	Write(wc chan string)
}

// PUSH 实例
// 暂时提供钉钉的输出方式 时间关系 后续有新的输出方式如日志 将这个文件改写成策略 增加log实例 实现Writer接口
type DingTalk struct {
	AccessToken string
	Secret      string
	PushLevel   string
	client      *dingtalk.Client
}

func NewDingTalk(accessToken string, secret string, pushLevel string) *DingTalk {
	client := createClient(accessToken, secret)
	return &DingTalk{AccessToken: accessToken, Secret: secret, PushLevel: pushLevel, client: client}
}

func (d *DingTalk) Write(wc chan string) {
	for x := range wc {
		// 发送到钉钉
		msg := dingtalk.NewTextMessage().SetContent(x).SetAt([]string{"177010xxx60"}, false)
		_, _ = d.client.Send(msg)
	}
}

func createClient(accessToken string, secret string) *dingtalk.Client {
	return dingtalk.NewClient(accessToken, secret)
}
