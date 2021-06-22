package main

import (
	"fmt"
)

//数字>>右移位数,相当于除以2的几次方
func main() {
	//fmt.Println(88 >> 1)
	//fmt.Println(88<<1 | 2)
	//fmt.Println(4 | 2)

	/*goto 跳转到指定行*/
	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		goto label
	//	}
	//}
	//fmt.Printf("111")
	//label:
	//fmt.Printf("222")
	/*闭包: 匿名函数*/
	//s := getS()
	//fmt.Println(&s)

	/*channels*/
	//messages := make(chan string,3)
	////go func() { messages <- "ping" }()
	//messages <- "1"
	//messages <- "2"
	////msg := <-messages
	////fmt.Printf(msg)
	//close(messages)
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)
	//
	///*select*/
	//c1 := make(chan string)
	//c2 := make(chan string)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	c1 <- "one"
	//}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c2 <- "two"
	//}()
	//for i := 0; i < 2; i++ {
	//	select {
	//	case msg1 := <-c1:
	//		fmt.Println("received", msg1)
	//	case msg2 := <-c2:
	//		fmt.Println("received", msg2)
	//	}
	//}

	/*定时器*/
	//ticker := time.NewTicker(500 * time.Millisecond)
	//done := make(chan bool)
	//go func() {
	//	for {
	//		select {
	//		case <-done:
	//			return
	//		case t := <-ticker.C:
	//			fmt.Println("Tick at", t)
	//		}
	//	}
	//}()
	//time.Sleep(1600 * time.Millisecond)
	//ticker.Stop()
	//done <- false
	//fmt.Println("Ticker stopped")

	if true {
		defer fmt.Println("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")

}

func ReadWrite() bool {
	//failureX, _ := os.Open("file")
	// 做一些工作
	//if failureX {
	//	failureX.Close()
	//	return false
	//}
	//
	//if failureY {
	//	failureX.Close()
	//	return false
	//}
	//
	//failureX.Close()
	//return true
	return true
}

func printf() {
	fmt.Printf("%3.2f", 3.14)
}

func getS() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func pointer(int2 *int) {

}
