package main

import (
	"fmt"
	"time"
)

func after() {
	//time.After() 超时一次后，就放弃的定时器
	tchan := time.After(time.Second * 3)
	fmt.Print(time.Now().String(), "tchan type:%T", tchan)
	fmt.Println(time.Now().String(), "mark 1")
	//channel 取出数据之后，发现超时时间为3秒
	fmt.Println(time.Now().String(), "tchan = ", <-tchan)
	fmt.Println(time.Now().String(), "mark 2")
}

//每隔2秒从channel获取一次数据
func tick() {
	//time.Tick()每隔一段时间就启用的定时器
	c := time.Tick(2 * time.Second)
	for next := range c {
		fmt.Printf("%v \n", next)
	}
}

//time.NewTicker()控制条件的定时器
func ticker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		//5秒后触发stop关闭ticker
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!!")
			return
		//获取此定时器channl 取出数据
		case t := <-ticker.C:
			fmt.Println("current time: ", t)
		}
	}
}

func main() {
	ticker()
	//after()
	//tick()
}
