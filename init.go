package alarm

import (
	"github.com/Islooper/warning/analyse"
	"github.com/Islooper/warning/push"
)

func InitWarning(accessToken, secret string) *analyse.WarningProcess {
	// 初始化
	//reader := InitReader()
	writer := InitWriter(accessToken, secret)
	warningAnalyse := InitWarningAnalyse(writer)

	//go reader.Read(warningAnalyse.R)
	go warningAnalyse.Analyse()
	go writer.Write(warningAnalyse.W)

	return warningAnalyse
}

//// 以下两个具体实例暂写成固定 后续有追求加内容改为由策略返回的抽象上下文
//func InitReader() *message.Service {
//	return message.NewService("serviceName" , "warning")
//}

func InitWriter(accessToken, secret string) *push.DingTalk {
	return push.NewDingTalk(accessToken, secret, "WARNING")
}

/////////// 分析器
func InitWarningAnalyse(writer push.Writer) *analyse.WarningProcess {
	return analyse.NewWarningProcess(make(chan string), make(chan string), writer)
}
