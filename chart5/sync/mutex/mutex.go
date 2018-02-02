package main


import (
	"strconv"
	"sync"
	"fmt"
)



type Currency struct {
	amount float64
	code   string
	mu     sync.RWMutex
}

func (c *Currency) Add(i float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.amount += i/10
	fmt.Printf("\n 当前汇率：%s", c.Display())
}

func (c *Currency) Display() string {
	return strconv.FormatFloat(c.amount, 'f', 2, 64) + " " + c.code
}

func (c *Currency)handleCurrency(i float64, wg *sync.WaitGroup){
	c.Add(i)
	wg.Done()

}

func main(){
	var currency = &Currency{amount: 50.00, code: "CNY"}
	var wg sync.WaitGroup
	threadsNum := 10
	wg.Add(threadsNum)

	for i:=0 ; i < threadsNum; i++ {
		currency.handleCurrency(float64(i) , &wg)
	}
	wg.Wait()
	fmt.Println("\n程序结束...")
}