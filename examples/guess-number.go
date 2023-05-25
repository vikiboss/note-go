package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	count := 0

	fmt.Println("✨ 猜数字游戏开始啦！")

	for {
		var guess int
		fmt.Print("⏰ 请输入你猜测的数字（1-100）：")
		fmt.Scan(&guess)

		count++

		if guess < target {
			fmt.Println("🟢 猜测的数字太小了！请继续猜测。")
		} else if guess > target {
			fmt.Println("🔴 猜测的数字太大了！请继续猜测。")
		} else {
			fmt.Printf("🎉 猜对啦！答案是 %d，你一共猜了 %d 次", target, count)
			break
		}
	}
}
