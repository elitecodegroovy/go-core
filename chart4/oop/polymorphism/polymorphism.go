package main

import (
	"strconv"
	"fmt"
	"bytes"
)

type WeChatAlert interface {
	NotifyWeChat() bool
}

type ManagerSystem struct {
	Id 		int64
	Name 	string
	Issue 	string
}

func (m *ManagerSystem) NotifyWeChat() bool {
	var buff bytes.Buffer
	buff.WriteString("ManagerSystem Name:")
	buff.WriteString(m.Name)
	buff.WriteString(", Id: ")
	buff.WriteString(strconv.FormatInt(m.Id, 10))

	buff.WriteString(" issue:")
	buff.WriteString(m.Issue)
	//TODO ... send msg to wechat account
	fmt.Printf("ManagerSystem sent msg [%s] successfully!\n", buff.String())
	return true
}

type MonitorSystem struct {
	Id 		int64
	Name    string
	Issue	string
}

func (m *MonitorSystem) NotifyWeChat() bool {
	var buff bytes.Buffer
	buff.WriteString("MonitorSystem Name:")
	buff.WriteString(m.Name)
	buff.WriteString(", Id: ")
	buff.WriteString(strconv.FormatInt(m.Id, 10))

	buff.WriteString(" issue:")
	buff.WriteString(m.Issue)
	//TODO ... send msg to wechat account
	fmt.Printf("MonitorSystem sent msg [%s] successfully!\n", buff.String())
	return true
}

func Notify(alert WeChatAlert){
	alert.NotifyWeChat()
}

func main(){
	var system = ManagerSystem{100, "CRM System", " 数据连接数过大"}
	Notify(&system)

	var monitor = MonitorSystem{101, "应用程序监控", "请求超时" }
	Notify(&monitor)
}