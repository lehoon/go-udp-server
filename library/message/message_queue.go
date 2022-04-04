package message

import (
	"time"
)

//响应内容
type Request struct {
	Opcode      uint16       //命令字
	SequenceNo  uint16       //流水号
	Length      int          //消息长度
	Buffer      []byte       //下发的消息内容
	Created     string       //生成时间
}

//响应内容
type Response struct {
	Opcode      uint16       //命令字
	SequenceNo  uint16       //流水号
	Length      int          //消息长度
	Buffer      []byte       //下发的消息内容
	Created     string       //生成时间
}

type Queue struct {
	msgCh   chan interface{}   //消息通道
}

var messageSizeMax  = 1024

var (
	inputQueue   *Queue        //gateway过来的event
	outputQueue  *Queue        //回复gateway的事件
	requestQueue *Queue        //agent主动发起的事件队列
)

func init()  {
	//初始化2个消息队列 一个request、一个response
	inputQueue = &Queue {
		msgCh: make(chan interface{}, messageSizeMax),
	}

	outputQueue = &Queue {
		msgCh: make(chan interface{}, messageSizeMax),
	}

	requestQueue = &Queue{
		msgCh: make(chan interface{}, messageSizeMax),
	}
}

func InputMessageQueue() *Queue {
	return inputQueue
}

func OutPutMessageQueue() *Queue {
	return outputQueue
}

func RequestMessageQueue() *Queue {
	return requestQueue
}

//消息队列尾追加消息
func (self *Queue) Push(msg interface{}) {
	self.msgCh <- msg
}

//从消息队列头取消息
func (self *Queue) Poll() interface{} {
	select {
	case msg := <- self.msgCh:
		return msg
	case <- time.After(time.Second * 3):
	default:
	}
	return nil
}