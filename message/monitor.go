package message

import (
	"errors"
	"github.com/Islooper/warning/analyse"
)

//type Reader interface {
//	Read(rc chan string)
//}
//
//// 暂时提供监控服务 时间关系 如果后续需要支持多种数据来源方式 将这个文件改写策略
//type Service struct {
//	// 简单提供些配置
//	Name  string
//	Level string
//	data  string
//}
//
//func NewService(name string, level string) *Service {
//	return &Service{Name: name, Level: level}
//}
//
//func (s *Service) Read(rc chan string) {
//	fmt.Println("读取中...")
//	rc <- s.data
//}
//
//

type SendMessageEr interface {
	SendMessage(data string) error
}

type Message struct {
	Name  string
	Level string
	*analyse.WarningProcess
}

func NewMessage(name string, level string, warningProcess *analyse.WarningProcess) *Message {
	return &Message{Name: name, Level: level, WarningProcess: warningProcess}
}

func (m *Message) SendMessage(data string, des string, err error) error {
	if data == "" {
		return errors.New("param is empty")
	}

	if m.WarningProcess == nil {
		return errors.New("param is empty")
	}

	msg := "监控中：\r\n" + m.Name + ":\r\n" + m.Level + ":\r\n" + data + "\r\n==========" + des + "\r\n" + err.Error()
	m.R <- msg
	return nil
}
