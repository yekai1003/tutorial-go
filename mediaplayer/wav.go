/*
   author:Yekai
   company:Pdj
   filename:wav.go
*/
package main

import (
	"fmt"
	"time"
)

//wav 播放器
type WavPlayer struct {
	stat     int //状态
	progress int //进度
}

//播放音乐，显示进度
func (p *WavPlayer) Play(source string) {
	fmt.Println("Playing wav music", source)
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 这一定是一个假的播放器，只有进度，没有声音
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
