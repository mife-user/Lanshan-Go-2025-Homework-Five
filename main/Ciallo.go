package main

import (
	"Lanshan-homework/five/cyjj"
	"Lanshan-homework/five/kqjj"
	"Lanshan-homework/five/role"
	"Lanshan-homework/five/rtjj"
	"fmt"
	"os"
	"sync"
	"time"
)

var cy cyjj.Cyjj
var kq kqjj.Kqjj
var rt rtjj.Rtjj
var wg sync.WaitGroup
var my_everyone role.Now_everyone
var mu sync.Mutex
var cyChan chan Role = make(chan Role, 1)
var kqChan chan Role = make(chan Role, 1)
var rtChan chan Role = make(chan Role, 1)
var cygoodChan chan role.Good = make(chan role.Good, 1)
var kqgoodChan chan role.Good = make(chan role.Good, 1)
var rtgoodChan chan role.Good = make(chan role.Good, 1)
var cycgChan chan bool = make(chan bool, 1)
var kqcgChan chan bool = make(chan bool, 1)
var rtcgChan chan bool = make(chan bool, 1)

type Role struct {
	Action string
	Text   string
	Name   string
}

func fetch(f role.Personer, c chan Role) {
	defer wg.Done()
	for Rolet := range c {
		switch Rolet.Name {
		case "cyjj":
			switch Rolet.Action {
			case "talk":
				f.Talk(Rolet.Text)
				cygoodChan <- f.Getgood()
			case "tellcg":
				f.Tellcg()
				cycgChan <- f.Getcg()
			case "story":
				f.Story()
			}
		case "kqjj":
			switch Rolet.Action {
			case "talk":
				f.Talk(Rolet.Text)
				kqgoodChan <- f.Getgood()
			case "tellcg":
				f.Tellcg()
				kqcgChan <- f.Getcg()
			case "story":
				f.Story()
			}
		case "rtjj":
			switch Rolet.Action {
			case "talk":
				f.Talk(Rolet.Text)
				rtgoodChan <- f.Getgood()
			case "tellcg":
				f.Tellcg()
				rtcgChan <- f.Getcg()
			case "story":
				f.Story()
			}
		}
	}
}

func searchperson() {
	mu.Lock()
	fmt.Println("当前角色状态：")
	for name, good := range my_everyone.One_good {
		fmt.Printf("角色: %s, 好感度: %d\n", name, good)
	}
	fmt.Println("是否解锁cg：")
	for name, cg := range my_everyone.One_cg {
		fmt.Printf("角色: %s, 状态: %t\n", name, cg)
	}
	mu.Unlock()
}

func new() {
	mu.Lock()
	defer mu.Unlock()

	// 初始化映射
	my_everyone.One_good = make(map[string]role.Good)
	my_everyone.One_cg = make(map[string]bool)

	file, err := os.Open("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/user.txt")
	if err != nil {
		fmt.Println("未找到存档文件，创建默认数据...")
		// 设置默认值
		my_everyone.One_good["cyjj"] = 0
		my_everyone.One_good["kqjj"] = 0
		my_everyone.One_good["rtjj"] = 0
		my_everyone.One_cg["cyjj"] = false
		my_everyone.One_cg["kqjj"] = false
		my_everyone.One_cg["rtjj"] = false
		return
	}
	defer file.Close()

	fmt.Println("读取存档中...")
	var name string
	var good role.Good
	var cg bool
	for {
		_, err := fmt.Fscanf(file, "%s %d %t\n", &name, &good, &cg)
		if err != nil {
			break
		}
		my_everyone.One_good[name] = good
		my_everyone.One_cg[name] = cg
	}
}

func write() {
	file, err := os.Create("D:/vscode/VsCodeWork/Lanshan-Go-2025-Homework/five/user.txt")
	if err != nil {
		fmt.Println("写入失败...")
		return
	}
	defer file.Close()

	for name, good := range my_everyone.One_good {
		cg := my_everyone.One_cg[name]
		fmt.Fprintf(file, "%s %d %t\n", name, good, cg)
	}
	fmt.Println("存档保存成功!")
}

func menu() int {
	for {
		fmt.Println("选择：0-结束 1-继续")
		var starEnd int
		fmt.Scanln(&starEnd)
		if starEnd == 0 {
			return 0
		}

		fmt.Println("选择角色: 1-cy 2-kq 3-rt")
		var choice int
		fmt.Scanln(&choice)

		fmt.Println("选择操作: 0-talk 1-tellcg 2-search 3-story")
		var cz int
		fmt.Scanln(&cz)

		switch cz {
		case 0:
			fmt.Println("你和jj们愉快的聊天中...>^<")
			fmt.Println("请输入你想对jj说的话：")
			var word string
			fmt.Scanln(&word)

			switch choice {
			case 1:
				tool := Role{Action: "talk", Text: word, Name: "cyjj"}
				cyChan <- tool
				return 1
			case 2:
				tool := Role{Action: "talk", Text: word, Name: "kqjj"}
				kqChan <- tool
				return 1
			case 3:
				tool := Role{Action: "talk", Text: word, Name: "rtjj"}
				rtChan <- tool
				return 1
			default:
				fmt.Println("角色选择错误，请重新选择")
				continue
			}
		case 1:
			switch choice {
			case 1:
				tool := Role{Action: "tellcg", Name: "cyjj"}
				cyChan <- tool
				return 1
			case 2:
				tool := Role{Action: "tellcg", Name: "kqjj"}
				kqChan <- tool
				return 1
			case 3:
				tool := Role{Action: "tellcg", Name: "rtjj"}
				rtChan <- tool
				return 1
			default:
				fmt.Println("角色选择错误，请重新选择")
				continue
			}
		case 2:
			searchperson()
			return 1
		case 3:
			switch choice {
			case 1:
				tool := Role{Action: "story", Name: "cyjj"}
				cyChan <- tool
				return 1

			case 2:
				tool := Role{Action: "story", Name: "kqjj"}
				cyChan <- tool
				return 1

			case 3:
				tool := Role{Action: "story", Name: "rtjj"}
				cyChan <- tool
				return 1
			}
		default:
			fmt.Println("输入错误，重新输入...")
			continue

		}
	}
}

func getAll() {
	mu.Lock()
	defer mu.Unlock()

	select {
	case wsgcyg := <-cygoodChan:
		my_everyone.One_good["cyjj"] = wsgcyg
		fmt.Printf("cyjj好感度更新为: %d\n", wsgcyg)
	default:
	}

	select {
	case wsgcyc := <-cycgChan:
		my_everyone.One_cg["cyjj"] = wsgcyc
		if wsgcyc {
			fmt.Println("cyjj CG已解锁!")
		}
	default:
	}

	select {
	case wsgkqg := <-kqgoodChan:
		my_everyone.One_good["kqjj"] = wsgkqg
		fmt.Printf("kqjj好感度更新为: %d\n", wsgkqg)
	default:
	}

	select {
	case wsgkqc := <-kqcgChan:
		my_everyone.One_cg["kqjj"] = wsgkqc
		if wsgkqc {
			fmt.Println("kqjj CG已解锁!")
		}
	default:
	}

	select {
	case wsgrtg := <-rtgoodChan:
		my_everyone.One_good["rtjj"] = wsgrtg
		fmt.Printf("rtjj好感度更新为: %d\n", wsgrtg)
	default:
	}

	select {
	case wsgrtc := <-rtcgChan:
		my_everyone.One_cg["rtjj"] = wsgrtc
		if wsgrtc {
			fmt.Println("rtjj CG已解锁!")
		}
	default:
	}
}

func main() {
	fmt.Println("实际上用不上并发，我强塞的hhh，就当练习吧，我觉得非常有趣，虽然是用屎堆出来的>.<")
	cy.New()
	kq.New()
	rt.New()
	new()
	wg.Add(3)
	go fetch(&cy, cyChan)
	go fetch(&kq, kqChan)
	go fetch(&rt, rtChan)
	for {
		if menu() == 1 {
			fmt.Println("等待数据传输中...")
			time.Sleep(time.Second * 1)
			getAll()
			write()
		} else {
			close(cyChan)
			close(kqChan)
			close(rtChan)
			break
		}
	}
	wg.Wait()
}
