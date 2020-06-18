/*
	用于更新生成Token的秘钥
	每日凌晨0点自动更新，会
	保留昨日的秘钥并更新今天
	的秘钥
	author：Kyda
	date:06.18.2020
*/

package main

import (
	"math/rand"
	"time"
)

const keyLength = 30 // 秘钥长度
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 暂时放这里，后面上etcd再改
var (
	yesterdayKey string
	todayKey     string
)

// createTokenKey 用于生成密钥
func createTokenKey(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// run
func run() {
	todayKey := createTokenKey(keyLength)
	for {
		now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
		t := time.NewTimer(next.Sub(now))                                                    //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		yesterdayKey = todayKey
		todayKey = createTokenKey(keyLength)
		// fmt.Println(yesterdayKey, todayKey)
	}
}

func main() {
	run()
}

//func Run()  {
//	waitTime = time.Second*10
//	for {
//		time.Sleep(waitTime)
//		yesterdayKey = todayKey
//		todayKey = createTokenKey(keyLength)
//		fmt.Println(yesterdayKey, todayKey)
//		fmt.Println()
//	}
//}
