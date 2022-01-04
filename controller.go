package main

import (
	"context"
	"fmt"
	"selfmade-webframework/framework"
	"time"
)

func FooControllerHandler(c *framework.Context) error {

	//繼承request的Context，創建一個設置超時的Context
	//這裏設置超時時間為一秒
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	//所有事情結束後調用 cancel，告知 durationCtx 之後的子 Context 們都结束
	defer cancel()

	//建立一個buffer為1的channel,負責通知結束
	finish := make(chan struct{}, 1)

	//這個channel則是負責通知 panic 異常
	panicChan := make(chan interface{}, 1)


	go func() {

		//這裏做異常處理
		defer func ()  {
			if p := recover(); p != nil{
				panicChan <- p
			}
		}()

		//這裏做具體業務
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		//新的 goroutine 結束的時候通過一個 finish channel 來告知父 goroutine
		finish <- struct{}{}
	}()

	
	select {
		// 监听 panic
		// case  p := <-panicChan:
		case  <-panicChan:
			c.Json(500, "panic")
		// 监听结束事件
	  	case <-finish:
			fmt.Println("finish")
		// 监听超时事件
	  	case <-durationCtx.Done():
			c.Json(500, "time out")
	  }
	  return nil

  }

