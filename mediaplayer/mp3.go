/*
   author:Yekai
   company:Pdj
   filename:mp3.go
*/
package main

import (
	"fmt"
	"time"
)

//mp3 播放器
type MP3Player struct {
	stat     int //状态
	progress int //进度
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music", source)
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 这一定是一个假的播放器，只有进度，没有声音
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
