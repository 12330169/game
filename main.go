package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func getLuckyNumbers() ([33]int, [16]int) {
	var redBall int
	var numbers [6]int
	var lucky int
	var redCount [33]int
	var blueCount [16]int

	//红球算法
	var redBalls [33]int
	for i := 0; i < 33; i++ {
		redBalls[i] = i + 1
	}
	for i := 0; i < 33; i++ {
		redBall = rand.Intn(33)
		redBalls[i], redBalls[redBall] = redBalls[redBall], redBalls[i]
	}
	for i := 0; i < 6; i++ {
		numbers[i] = redBalls[i]
		redCount[numbers[i]-1]++
	}

	//蓝球算法
	lucky = rand.Intn(16) + 1
	blueCount[lucky-1]++
	fmt.Printf("本期中奖号码为：%d,%d\n", numbers, lucky)
	return redCount, blueCount
}
func getLuckyNumbers2() ([33]int, [16]int) {
	var redBall int
	var numbers [6]int
	var lucky int
	var redCount [33]int
	var blueCount [16]int

	//红球算法
	for i := 0; i < 6; i++ {
		for {
			var isExist = false
			redBall = rand.Intn(33) + 1
			for j := 0; j < 6; j++ {
				if redBall == numbers[j] {
					isExist = true
					break
				}
			}
			if !isExist {
				numbers[i] = redBall
				redCount[redBall-1]++
				break
			}
		}

	}
	//蓝球算法
	lucky = rand.Intn(16) + 1
	blueCount[lucky-1]++
	fmt.Printf("本期中奖号码为：%d,%d\n", numbers, lucky)
	return redCount, blueCount
}
func selfChoose() {
	var i int
	size := 1
	var redCount [33]int
	var blueCount [16]int
	var redCount1 [33]int
	var blueCount1 [16]int
	fmt.Println("请输入您选择的红球号码：")

	for i = 0; i < size; i++ {
		redCount1, blueCount1 = getLuckyNumbers()
		for j := 0; j < 33; j++ {
			redCount[j] = redCount[j] + redCount1[j]
		}
		for j := 0; j < 16; j++ {
			blueCount[j] = blueCount[j] + blueCount1[j]
		}
	}
	fmt.Printf("统计%d次开奖结果各个数出现的次数及概率：\n", size)
	fmt.Println("蓝球:")
	for i = 0; i < 33; i++ {
		var red = float32(redCount[i])
		var size = float32(size)
		fmt.Printf("[%d]:%d,%.01f%%\n", i+1, redCount[i], red/size*100)
	}
	fmt.Println("红球")
	for i = 0; i < 16; i++ {
		var blue = float32(blueCount[i])
		var size = float32(size)
		fmt.Printf("[%d]:%d,%.01f%%\n", i+1, blueCount[i], blue/size*100)
	}
}
func machineChoose() {
	var i int
	var redCount [33]int
	var blueCount [16]int
	var redCount1 [33]int
	var blueCount1 [16]int
	size := 1
	for i = 0; i < size; i++ {
		redCount1, blueCount1 = getLuckyNumbers2()
		for j := 0; j < 33; j++ {
			redCount[j] = redCount[j] + redCount1[j]
		}
		for j := 0; j < 16; j++ {
			blueCount[j] = blueCount[j] + blueCount1[j]
		}
	}
	fmt.Printf("统计%d次开奖结果各个数出现的次数及概率：\n", size)
	fmt.Println("蓝球:")
	for i = 0; i < 33; i++ {
		var red = float32(redCount[i])
		var size = float32(size)
		fmt.Printf("[%d]:%d,%.01f%%\n", i+1, redCount[i], red/size*100)
	}
	fmt.Println("红球")
	for i = 0; i < 16; i++ {
		var blue = float32(blueCount[i])
		var size = float32(size)
		fmt.Printf("[%d]:%d,%.01f%%\n", i+1, blueCount[i], blue/size*100)
	}

}
func conculatPrize(red, blue int) {
	switch blue {
	case 0:

	case 1:
	default:
		fmt.Println("非法数据")
	}
}
func conculatProbability(red [33]int, blue [16]int) {
	var newRed [6]int
	var newBlue int
	oldRed := red
	oldBlue := blue
	for i := 0; i < 33; i++ {
		for j := 32; j > i; j-- {
			if red[i] < red[j] {
				red[i], red[j] = red[j], red[i]
			}
		}
	}
	for i := 0; i < 16; i++ {
		for j := 15; j > i; j-- {
			if blue[i] < blue[j] {
				blue[i], blue[j] = blue[j], blue[i]
			}
		}
	}
	fmt.Println(red)
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			for k := 0; k < 33; k++ {
				if red[i] == oldRed[k] {
					newRed[i] = k
				}
			}

		}
	}
	for i := 0; i < 16; i++ {
		if oldBlue[i] == blue[0] {
			newBlue = i
		}
	}
	fmt.Println("概率最高的红球和蓝球分别是:")
	fmt.Println(newRed)
	fmt.Println(newBlue)
}

func main() {
	for {
		fmt.Println("投注方式：\n1.自选 2.机选 3.退出系统")
		var choseType int
		fmt.Println("请选择:")
		fmt.Scanln(&choseType)
		switch choseType {
		case 1:
			selfChoose()
		case 2:
			machineChoose()
		default:
			fmt.Println("无效输入")
		}
	}

	// test()
	// getLuckyNumbers()
}
