package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

//rand.Intn返还[0,n)的随机数
//Seed值固定，那么保证随机数一定固定
func RandNWithSeed() {
	rand.Seed(100)
	answers := []string{
		"多得你嘅帮助，好多谢你。",
		"多谢，多谢你嘅好意。",
		"呢次你帮左我好大嘅忙。",
		"呢个周末得闲吗？",
		"你听日得闲吗？",
		"对唔住，今日下午无时间。",
		"我想请你去参观科技馆。 ",
		"多谢你嘅鼓励。 ",
		"多谢你嘅邀请。  ",
		"唔使多谢，我自愿嘅。 ",
		"唔使客气。",
		"多谢你嘅祝福。 ",
		"多谢你嘅指导。",
		"对于你嘅建议，我会认真考虑。",
		"对于你嘅关心，我好感激。",
		"呢次你帮左我好大嘅忙。",
	}

	fmt.Println("粤语答复:", answers[rand.Intn(len(answers))])
	// 输出:	//粤语答复: 多谢你嘅鼓励。

	//10的数值排列
	fmt.Println("数值排列:", rand.Perm(10))
	//数值排列: [7 9 8 6 3 0 1 2 4 5],由于seed的值固定了。

	//current time as seed value
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//It is always different when you run the code.
	ns := r.Perm(10)
	fmt.Println("数值排列:", ns)
	//loop print the slice's elements
	for _, n := range ns {
		fmt.Printf("%d\t", n)
	}
}

func cryptoRand() {
	//生成质数
	for n := 2; n < 10; n++ {
		p, err := crand.Prime(crand.Reader, n)
		if err != nil {
			fmt.Errorf("Can't generate %d-bit prime: %v", n, err)
		}
		if p.BitLen() != n {
			fmt.Errorf("%v is not %d-bit", p, n)
		}
		if !p.ProbablyPrime(32) {
			fmt.Errorf("%v is not prime", p)
		}
		fmt.Printf("crypto/rand prime %d:\n", p)
	}

	//1>Int [0, 100000) random number
	b := new(big.Int).SetInt64(int64(100000))
	if i, err := crand.Int(crand.Reader, b); err != nil {
		fmt.Errorf("Can't generate random value: %v, %v", i, err)
	} else {
		fmt.Printf("rand [0, 100000) random value : %d\n", i)
	}

	//2> big int output
	b1 := new(big.Int).SetInt64(int64(math.MaxInt64))
	if i1, err := crand.Int(crand.Reader, b1); err != nil {
		fmt.Errorf("Can't generate random value: %v, %v", i1, err)
	} else {
		fmt.Println("rand [0, math.MaxInt64) random value : ", i1)
	}
}

func main() {
	RandNWithSeed()
	cryptoRand()
}
