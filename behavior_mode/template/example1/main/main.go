package main

import "go-patterns/behavior_mode/template"

func main1() {
	// 做西红柿
	xihongshi := &template.XiHongShi{}
	template.DoCook(xihongshi)

	// 做炒鸡蛋
	chaojidan := &template.ChaoJiDan{}
	template.DoCook(chaojidan)
}

func main() {
	cricketGame := new(template.Cricket)
	template.Play(cricketGame)

	footballGame := new(template.Football)
	template.Play(footballGame)
}
