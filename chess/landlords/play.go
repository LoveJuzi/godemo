package landlords

import (
	"fmt"
	"godemo/chess/utils/task"
	"time"
)

// PLAYERSNUM 玩家的个数
const PLAYERSNUM = 3

// PLTYPE 玩家类型
type PLTYPE int

// Stauts 当前斗地主的状态定义
type Stauts struct {
	Plcrds [PLAYERSNUM][]string
	Plts   [PLAYERSNUM]PLTYPE
	Win    [PLAYERSNUM]bool
	ID     int
	Cards  []string
}

func nextid(id int) int {
	return (id + 1) % PLAYERSNUM
}

// NewGame 初始化一个新的游戏
func NewGame() Stauts {
	return Stauts{
		Plcrds: [PLAYERSNUM][]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		},
		Plts:  [PLAYERSNUM]PLTYPE{0, 0, 1},
		Cards: []string{},
		Win:   [PLAYERSNUM]bool{false, false, false},
	}
}

// Play 斗地主主程序
func Play(st Stauts) Stauts {
	ch1s := make(chan Stauts, 1)
	ch1r := make(chan Stauts, 1)

	ch2s := make(chan Stauts, 1)
	ch2r := make(chan Stauts, 1)

	ch3s := make(chan Stauts, 1)
	ch3r := make(chan Stauts, 1)

	ch := make(chan int, 1)
	task.SyncTask(taskScanf(ch), func() { close(ch) })
	defer func() { close(ch) }()

	T1 := task.SyncTask(taskNextCards(ch1r, ch1s, 0, ch))
	T2 := task.SyncTask(taskNextCards(ch2r, ch2s, 1, ch))
	T3 := task.SyncTask(taskNextCards(ch3r, ch3s, 2, ch))

	for {
		// 给玩家发送当前牌的信息
		ch1r <- st
		ch2r <- st
		ch3r <- st

		// 接收玩家出牌
		var id int
		select {
		case st = <-ch1s:
			id = 0
		case st = <-ch2s:
			id = 1
		case st = <-ch3s:
			id = 2
		}

		// 判定玩家输赢
		if len(st.Plcrds[id]) == 0 {
			plt := st.Plts[id]
			for i := 0; i < PLAYERSNUM; i++ {
				if st.Plts[id] == plt {
					st.Win[id] = true
				}
			}
			break
		}

		// 更新下一个出牌的玩家
		st.ID = nextid(st.ID)
	}

	// 流程控制
	close(ch1r)
	close(ch2r)
	close(ch3r)
	<-T1
	<-T2
	<-T3

	return st
}

func taskNextCards(chr <-chan Stauts, chs chan<- Stauts, id int, ch <-chan int) func() {
	return func() {
		for st := range chr {
			st = NextCards(id, st, ch)
			if id == st.ID {
				chs <- st
			}
		}
	}
}

func taskScanf(ch chan<- int) func() {
	return func() {
		Scanf(ch)
	}
}

// Scanf 从标准输入读入数据
func Scanf(ch chan<- int) {
	var idx int
	for {
		fmt.Scanf("%d", &idx)
		ch <- idx
	}
}

// NextCards 选手出牌
func NextCards(id int, st Stauts, ch <-chan int) Stauts {
	// 输出当前选手的手牌
	fmt.Printf("当前的选手号：%d, 手牌：%v\n", id, st.Plcrds[id])
	// ch2 chan<- string
	// ch2 <- 输出信息

	if id != st.ID {
		return st
	}

	// 决定出什么牌
	midx := len(st.Plcrds[id])
	ch1 := make(chan int, 1)       // 接收异步输入值
	done := make(chan struct{}, 1) // 超时通知信号
	// 异步有效输入
	task.SyncTask(func() {
		var idx int
		for {
			fmt.Printf("选手号：%d，输入范围0-%d\n", id, midx-1)
			// ch2 chan<- string
			// ch2 <- 输入提示信息
			select {
			case idx = <-ch:
				if idx >= -1 && idx < midx {
					// 还要判定输入是否有效
					fmt.Println(idx)
					ch1 <- idx
					goto DONE
				}
			case <-done:
				goto DONE
			}
		}
	DONE:
	})
	// 超时输入
	var idx int
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("超时")
		done <- struct{}{} // 停止异步输入
		idx = -1
	case idx = <-ch1:
	}
	close(ch1)
	close(done)

	if idx >= 0 && idx < midx {
		// 保存出的牌
		st.Cards = []string{}
		st.Cards = append(st.Cards, st.Plcrds[id][idx])

		// 更新选手的手牌
		j := 0
		for i := 0; i < midx; i++ {
			if i == idx {
				continue
			}
			st.Plcrds[id][j] = st.Plcrds[id][i]
			j++
		}
		st.Plcrds[id] = st.Plcrds[id][:midx-1]
	}

	return st
}
