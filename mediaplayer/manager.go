/*
   author:Yekai
   company:Pdj
   filename:manager.go
*/
package main

import "errors"

type MusicEntry struct {
	Id     string //编号
	Name   string //歌名
	Artist string //作者
	Source string //位置
	Type   string //类型
}

type MusicManager struct {
	musics []MusicEntry
}

//入口函数
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//返回曲库当前数量
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//通过编号获取歌曲
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

//查询歌曲信息
func (m *MusicManager) Find(name string) (me *MusicEntry, index int) {
	if len(m.musics) == 0 {
		return nil, -1
	}
	for k, m := range m.musics {
		if m.Name == name {
			return &m, k
		}
	}
	return nil, -1
}

//像歌曲库添加一首歌曲
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//按编号删除歌曲
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	// 从数组切片中删除元素
	if index < len(m.musics)-1 { // 中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { // 删除仅有的一个元素
		m.musics = make([]MusicEntry, 0)
	} else { // 删除的是最后一个元素
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) (me *MusicEntry, index int) {
	e, idx := m.Find(name)
	if e == nil {
		return nil, idx
	}
	m.Remove(idx)
	return e, idx
}
