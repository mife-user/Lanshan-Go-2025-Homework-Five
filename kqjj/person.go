package kqjj

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
type Good int     //好感度>~<
// Kqjj结构体
type Kqjj struct {
	Name string
	Age  int
	Good role.Good
	cg   cgmax
	Cg   bool
}

// 初始化kqjj
func (kqjj *Kqjj) New() {
	kqjj.Name = "kqjj"
	kqjj.Age = 18
	kqjj.Good = 0
	kqjj.Cg = false
}

// 好感度判断
func (kqjj *Kqjj) Tellcg() {
	if kqjj.Good > 90 {
		fmt.Println("初始化cg中...>-<")
		file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/kqjj/kqjj_cg.txt")
		if err != nil {
			fmt.Println("打开cg文件失败>~<")
			return
		}
		var input bufio.Reader = *bufio.NewReader(file)
		kqjj.cg = func() {
			defer file.Close()
			var cgput []string
			for {
				cgputtemp, err := input.ReadString('\n')
				cgput = append(cgput, cgputtemp)
				if err != nil {
					if err == io.EOF {
						fmt.Println("读取cg文件完成>^<")
						kqjj.Cg = true
						break
					}
					fmt.Println("读取cg文件失败>~<")
					break
				}
			}
			fmt.Println("cg内容如下: ")
			for _, line := range cgput {
				fmt.Print(line)
			} //编译器自带ai填充，挺好用的>,<(甚至还会模仿我的语言风格，可能也觉得jj们可爱吧，开玩笑的, 什么你问我为啥代码和cyjj一模一样？因为我直接复制粘贴改名字了呀~>-<	)
		}
		kqjj.cg()
	} else {
		fmt.Println("你的好感度为：", kqjj.Good)
		fmt.Println("好感度不足90，无法触发cg>-<")
	}
}

// kqjj好感度增加函数
func (kqjj *Kqjj) Talk(talkinput string) {
	rand.Seed(time.Now().UnixNano())
	talkchoose := rand.Intn(3)
	switch talkchoose {
	case 0:
		fmt.Println("kqjj：你说的真有趣！")
		kqjj.Good += 10
	case 1:
		fmt.Println("kqjj：哈哈，你真幽默！")
		kqjj.Good += 15
	case 2:
		fmt.Println("kqjj：嗯，我不太明白你的意思。")
		kqjj.Good += 5
	}
}

// kqjj故事函数
func (kqjj *Kqjj) Story() {
	file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/kqjj/kqjj_story.txt")
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
	fmt.Println("kqjj的故事内容如下: ")
	for _, line := range storyput {
		fmt.Print(line)
	}
}
func (kqjj *Kqjj) Getgood() role.Good {
	return kqjj.Good
}
func (kqjj *Kqjj) Getcg() bool {
	return kqjj.Cg
}
