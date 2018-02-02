package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
)

type Location struct {
	Ip 		string		`json:"ip"`
	Country string		`json:"country"`
	Region  string		`json:"region"`
	City 	string		`json:"city"`
	Isp 	string 		`json:"isp"`
}

type ResponseResult struct {
	Code 	int 				`json:"code"`
	Data 	map[string]string 	`json:"data"`
}

type IpInfo struct {
	Location
	err 	error				`json:"err"`
}


type CustomerInfo struct {
	Level 		int			`json:"level"`
	LevelName	string		`json:"levelName"`
	Username	string		`json:"username"`
	Ips         []IpInfo	`json:"ipInfo"`
}
func (c *CustomerInfo) GetIpLocation(ip string, ipR *IpInfo) {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		ipR.err = err
		return
	}
	var result ResponseResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err == nil {
		fmt.Println("rsult:: ", result)
		ipR.City = result.Data["city"]
		ipR.Region = result.Data["region"]
		ipR.Country = result.Data["country"]
		ipR.Ip = ip
		ipR.Isp = result.Data["isp"]
		fmt.Println("rsult:: ", ipR)
	}else {
		ipR.err = err
	}
}

func (c *CustomerInfo)CallTask(ip string, cs chan IpInfo) {
	var ipR IpInfo
	c.GetIpLocation(ip, &ipR)
	cs <- ipR
}

func (c *CustomerInfo)QueryCustomerInfo(username string, finished chan bool){
	//TODO ....query handler
	time.Sleep(100 * time.Millisecond)
	c.Username = username
	if username == "李强" {
		c.Level = 1
		c.LevelName = "钻石会员"
	}else {
		c.Level = 2
		c.LevelName = "普通会员"
	}
	fmt.Println("查询会员登记结束...")
	finished <- true
	close(finished)
}

func QueryCustomerInfo(ips []string, username string) (CustomerInfo, error) {
	cIp := make(chan IpInfo, len(ips))
	finished := make(chan bool)
	var userInfo CustomerInfo
	//查询ip信息的任务
	for _, ip := range ips {
		go userInfo.CallTask(ip, cIp)
	}
	//查询用户等级任务
	go userInfo.QueryCustomerInfo(username, finished)
	//等待任务执行完毕
	//开启的goroutine个数
	count := len(ips) + 1
	for {
		if count == 0 {
			break
		}
		select {
		case <- finished:
			count--
			fmt.Println("查询用户等级任务完成")
		case ipInfo := <- cIp:
			count--
			userInfo.Ips = append(userInfo.Ips, ipInfo)
			fmt.Println("ip信息:", ipInfo)
		}
	}
	return userInfo, nil
}

func GetCustomerInfo(){
	ips := []string{"45.32.251.98", "221.4.38.21"}
	customerInfo , err := QueryCustomerInfo(ips, "李强")
	if err == nil {
		data , _ := json.Marshal(customerInfo)
		fmt.Printf("info %s", data)
	}else {
		fmt.Errorf(" error %s", err.Error())
	}
}

func main(){
	GetCustomerInfo()
}