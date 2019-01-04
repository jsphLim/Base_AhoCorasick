package main


import (
	"./ac"
	"fmt"
)

func main() {
	content := "wm.Lan is the excellent professor of JNU"
	aho := ac.CreateAhoCorasick()
	aho.Add("wm")
	aho.Add("is")
	aho.Add("lan")
	aho.Add("JNU")
	aho.BuildAhoCorasick()
	results := aho.ScanAhoCorasick(content)
	fmt.Println("匹配成功的词: ")
	for _, result := range results {
		//fmt.Println(result)
		fmt.Println(string([]rune(content)[result.Begin : result.End+1]))
	}
}
