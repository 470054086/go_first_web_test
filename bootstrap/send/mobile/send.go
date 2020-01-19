package mobile

import (
	"first_web/bootstrap"
	"first_web/bootstrap/ini"
	"first_web/bootstrap/log"
	"first_web/bootstrap/send"
	"fmt"
	"math/rand"
	"strconv"
)

const chanLength = 10

type MobileSend struct {
	sendChan    []chan *send.SendMessage
	sendChanNum int //使用多少个chan来获取数据
	sendErrorChan chan *send.SendMessage
}

var G_Send MobileSend

func init() {
	//启动服务
	bootstrap.Func.AddProviders(func() {
		num := ini.Cfg.GetSelect("Message").GetKeyDefault("sendChanNum", "4")
		atoi, err := strconv.Atoi(num)
		if err != nil {
			panic(fmt.Sprintf("启动消息队列发生错误%v", err))
		}
		mobileSend := MobileSend{
			sendChanNum: atoi,
			sendErrorChan:make(chan *send.SendMessage),
		}
		for i := 0; i <= mobileSend.sendChanNum; i++ {
			mobileSend.sendChan = append(mobileSend.sendChan, make(chan
			*send.SendMessage, chanLength))
		}
		mobileSend.run()
		mobileSend.failRun()
		G_Send = mobileSend
	})
}

/**
随时添加一个channel
*/
func (s *MobileSend) Send(m *send.SendMessage) {
	//随机获取一个chan
	randInt := rand.Intn(s.sendChanNum)
	s.sendChan[randInt] <- m
}

/**
启动监听程序
*/
func (s *MobileSend) run() {
	for i := 0; i <= s.sendChanNum; i++ {
		go func(i int) {
			for {
				message := <-s.sendChan[i]
				fmt.Printf("当前使用的通道id为%d", i)
				//通过不同的参数 调用不同的协议
				log.Logger.Info(message)
				//如果发生了错误的话 将channel推送给发送错误的channel
				// todo 模拟错误的发生
				s.sendErrorChan <- message
			}
		}(i)
	}
}

func (s *MobileSend) failRun() {
	go func() {
		 for {
			 message := <-s.sendErrorChan
			 log.Logger.Info("我是错误处理")
			 log.Logger.Info(message)
		 }

	}()
}
