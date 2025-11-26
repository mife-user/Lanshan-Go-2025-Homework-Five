package rtjj

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
// Rtjj结构体
type Rtjj struct {
	Name string
	Age  int
	Good role.Good
	cg   cgmax
	Cg   bool
}

// 初始化cyjj
func (rtjj *Rtjj) New() {
	rtjj.Name = "rtjj"
	rtjj.Age = 19
	rtjj.Good = 0
	rtjj.Cg = false
}

// 好感度判断
func (rtjj *Rtjj) Tellcg() {
	if rtjj.Good > 90 {
		fmt.Println("初始化cg中...>-<")
		file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/rtjj/rtjj_cg.txt")
		if err != nil {
			fmt.Println("打开cg文件失败>~<")
			return
		}
		var input bufio.Reader = *bufio.NewReader(file)
		rtjj.cg = func() {
			defer file.Close()
			var cgput []string
			for {
				cgputtemp, err := input.ReadString('\n')
				cgput = append(cgput, cgputtemp)
				if err != nil {
					if err == io.EOF {
						fmt.Println("读取cg文件完成>^<")
						rtjj.Cg = true
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
		rtjj.cg()
	} else {
		fmt.Println("你的好感度为：", rtjj.Good)

		fmt.Println("好感度不足90，无法触发cg>-<")
	}
}

// rtjj好感度增加函数
func (rtjj *Rtjj) Talk(talkinput string) {
	rand.Seed(time.Now().UnixNano())
	talkchoose := rand.Intn(3)
	switch talkchoose {
	case 0:
		fmt.Println("rtjj：Ciallo～(∠・ω< )⌒☆")
		rtjj.Good += 10
	case 1:
		fmt.Println(`rtjj:你怎么能直接 commit 到我的 main 分支啊？！GitHub 上不是这样！
					你应该先 fork 我的仓库，然后从 develop 分支 checkout 一个新的 feature
					 分支，比如叫feature/confession。然后你把你的心意写成代码，并为它写好
					 单元测试和集成测试，确保代码覆盖率达到95%以上。接着你要跑一下 Linter
					 ，通过所有的代码风格检查。然后你再 commit，commit message 要遵循 Con
					 ventional Commits 规范。之后你把这个分支 push 到你自己的远程仓库，然
					 后给我提一个 Pull Request。在 PR 描述里，你要详细说明你的功能改动和
					 实现思路，并且 @ 我和至少两个其他的评审。我们会 review 你的代码，可
					 能会留下一些评论，你需要解决所有的 thread。等 CI/CD 流水线全部通过
					 ，并且拿到至少两个 LGTM 之后，我才会考虑把你的分支 squash and merg
					 e 到 develop 里，等待下一个版本发布。你怎么直接上来就想 force push
					  到 main？！GitHub 上根本不是这样！我拒绝合并！`)
		rtjj.Good += 15
	case 2:
		fmt.Println("rtjj：嗯，我不太明白你的意思。")
		rtjj.Good += 5
	}
}

// rtjj故事函数
func (rtjj *Rtjj) Story() {
	file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/rtjj/rtjj_story.txt")
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
	fmt.Println("rtjj的故事内容如下: ")
	for _, line := range storyput {
		fmt.Print(line)
	}
}
func (rtjj *Rtjj) Getgood() role.Good {
	return rtjj.Good
}
func (rtjj *Rtjj) Getcg() bool {
	return rtjj.Cg
}
