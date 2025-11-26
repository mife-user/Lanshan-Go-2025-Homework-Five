package cyjj

import (
	"Lanshan-homework/five/role"
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type cgmax func() //cg函数类型
// Cyjj结构体
type Cyjj struct {
	Name string
	Age  int
	Good role.Good
	cg   cgmax
	Cg   bool
}

// 初始化cyjj
func (cyjj *Cyjj) New() {
	cyjj.Name = "cyjj"
	cyjj.Age = 19
	cyjj.Good = 0
	cyjj.Cg = false
}

// 好感度判断
func (cyjj *Cyjj) Tellcg() {
	if cyjj.Good > 90 {
		fmt.Println("初始化cg中...>-<")
		file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/cyjj/cyjj_cg.txt")
		if err != nil {
			fmt.Println("打开cg文件失败>~<")
			return
		}
		var input bufio.Reader = *bufio.NewReader(file)
		cyjj.cg = func() {
			defer file.Close()
			var cgput []string
			for {
				cgputtemp, err := input.ReadString('\n')
				cgput = append(cgput, cgputtemp)
				if err != nil {
					if err == io.EOF {
						fmt.Println("读取cg文件完成>^<")
						cyjj.Cg = true
						break
					}
					fmt.Println("读取cg文件失败>~<")
					break
				}
			}
			fmt.Println("cg内容如下: ")
			for _, line := range cgput {
				fmt.Print(line)
			} //编译器自带ai填充，挺好用的>,<(甚至还会模仿我的语言风格，可能也觉得jj们可爱吧，开玩笑的, 什么你问我为啥代码和kqjj一模一样？因为我直接复制粘贴改名字了呀~>-<	)
		}
		cyjj.cg()
	} else {
		fmt.Println("你的好感度为：", cyjj.Good)
		fmt.Println("好感度不足90，无法触发cg>-<")
	}
}

// cyjj好感度增加函数
func (cyjj *Cyjj) Talk(talkinput string) {
	rand.Seed(time.Now().UnixNano())
	talkchoose := rand.Intn(3)
	switch talkchoose {
	case 0:
		fmt.Println("cyjj：鸭梨玛氏勒！")
		cyjj.Good += 10
	case 1:
		fmt.Println("cyjj:喵~")
		cyjj.Good += 15
	case 2:
		fmt.Println("cyjj：嗯，我不太明白你的意思。")
		cyjj.Good += 5
	}
}

// cyjj故事函数
func (cyjj *Cyjj) Story() {
	file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/cyjj/cyjj_story.txt")
	if err != nil {
		fmt.Println("打开故事文件失败>~<")
		return
	}
	defer file.Close()
	var intput bufio.Reader = *bufio.NewReader(file)
	var storyput []string
	for {
		storyputtemp, err := intput.ReadString('\n')
		storyput = append(storyput, storyputtemp)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取故事文件完成>^<")
				break
			}
			fmt.Println("读取故事文件失败>~<")
			break
		}
	}
	fmt.Println("cyjj的故事内容如下: ")
	for _, line := range storyput {
		fmt.Print(line)
	}
}
func (cyjj *Cyjj) Getgood() role.Good {
	return cyjj.Good
}
func (cyjj *Cyjj) Getcg() bool {
	return cyjj.Cg
}
