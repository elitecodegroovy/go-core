package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"sync"
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

func (c *CustomerInfo)CallTask(ip string, wg *sync.WaitGroup) {
	defer wg.Done()
	var ipR IpInfo
	c.GetIpLocation(ip, &ipR)
	c.Ips = append(c.Ips, ipR)
}

func (c *CustomerInfo)QueryCustomerInfo(username string, wg *sync.WaitGroup){
	defer wg.Done()
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
}

func QueryCustomerInfo(ips []string, username string) (CustomerInfo, error) {
	var userInfo CustomerInfo
	var wg sync.WaitGroup
	wg.Add(len(ips) + 1)
	//查询ip信息的任务
	for _, ip := range ips {
		go userInfo.CallTask(ip, &wg)
	}
	//查询用户等级任务
	go userInfo.QueryCustomerInfo(username, &wg)
	wg.Wait()
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