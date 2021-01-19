package main

import (
	"fmt"
	"strconv"
)

const (
	//簇内普通节点个数
	CH_NODE_NUM = 4

	//初始权值
	INIT_WEIGHT = 1.0

	//设置的阈值
	THERSHOLD = 0.6

)

//设置轮次的值
var round = [][CH_NODE_NUM]int{{1,0,1,0}, {0,0,0,1}, {1,0,0,1}, {1,1,1,0}}

func main() {
	Ch, Ch_weight := SetChValue(CH_NODE_NUM)
	fmt.Printf("簇头节点的值为:\n 	%v \n簇内节点的初始权值为:\n	%v\n", Ch, Ch_weight)

	w := CalculationWeight(Ch, Ch_weight, round)
	fmt.Printf("权值为:%v\n", w)

}

//设置簇头节点的value为1
func SetChValue(num int) (Ch []int, Ch_weight []float64) {
	var ch []int
	var ch_weight []float64
	for i:=0; i<num; i++ {
		j := 1
		ch = append(ch, j)
		ch_weight = append(ch_weight, INIT_WEIGHT)
	}
	return ch, ch_weight
}

//计算权值
func CalculationWeight(ch []int, ch_weight []float64, round [][CH_NODE_NUM]int) []float64{
	//var Weight []float64
	var weight float64
	for i:=1; i<len(round); i++ {
		for j:=0; j<len(ch_weight); j++ {
			if (round[i-1][j] & ch[j] != 1) {
				weight = float64(ch_weight[j]) - 0.2
			} else {
				weight = ch_weight[j]
			}
			ch_weight[j] = Tools(weight)
		}
		fmt.Printf("第%d轮计算的权值为:%v\n", i, ch_weight)

		if !IfMailciousNode(ch_weight) {
			break
		} else {
			continue
		}
	}
	return ch_weight
}

//判断恶意节点
func IfMailciousNode(ch_weight []float64) bool{
	for k, v := range ch_weight {
		if v < THERSHOLD {
			//v = Tools(v)
			fmt.Printf("簇内第%d个节点的权值为%.2f, 低于0.6,所以为恶意节点!\n", k+1, v)
			return false
		}
	}
	return true
}

func Tools(v float64) float64 {
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", v), 64)
	return v
}
