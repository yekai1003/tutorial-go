/*
   author:Yekai
   company:Pdj
   filename:play.go
*/
package main

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{0, 0}
	case "WAV":
		p = &WavPlayer{0, 0}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
