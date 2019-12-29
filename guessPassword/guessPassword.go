package main

import (
	"fmt"
	"os"
	"sort"
)


// 筛选符合条件的数组
func seekOut(ca [][]int) []int {

	result := []int{}

	LOOP:
	for ii:=0 ; ii < len(ca) ; ii++ {

		// 初始化检测值
		ck_1 := 1
		ck_2 := 0
		ck_3 := 0
		ck_4 := 0
		ck_5 := 0

		lenVa := len(ca[ii])
		va := ca[ii]

		// 检测一 条件四
		for i := 0 ; i < lenVa ; i++ {
			if va[i] == 1 || va[i] == 4 || va[i] == 7 {
				ck_1 = 0
				goto LOOP
			}
		}

		// 检测二 条件一
		if va[0] == 2 || va[2] == 6 {
			ck_2 = 1
		}

		// 检测三 条件二
		for i := 0 ; i < lenVa ; i++ {
			if (va[i] == 2 && i != 0) || (va[i] == 5 && i!=1) || (va[i] == 8 && i!=2) {
			 	ck_3 = 1
			}
		}

		// 检测四 条件三
		correctNumber := 0
		for i := 0 ; i < lenVa ; i++ {
			if va[i] == 6 || va[i] == 9 || va[i] == 2 {
				correctNumber += 1
			}
		}

		// 2个号码正确
		if correctNumber == 2 {
			if (va[0] != 6 && va[1] != 9 && va[2] != 2) || (va[1] != 9 && va[2] != 2 && va[0] != 6) || (va[0] != 6 && va[2] != 2 && va[1] != 9) {
				ck_4 = 1
			}
		}

		// 检测五 条件五
		if va[0] == 9 || va[1] == 9 {
			ck_5 = 1
		}

		// 筛选输出符合5个条件的数组
		result := ck_1 * ck_2 * ck_3 * ck_4 * ck_5
		if result == 1 {
			return va
		}

	}

	return result
}


// 排列组合
func sequence(ma []int) [][]int {

	// 前两位号码
	var clt [][]int
	var st1 []int
	var st2 []int
	st1 = append(st1,ma[0])
	st1 = append(st1,ma[1])
	clt = append(clt,st1)
	st2 = append(st2,ma[1])
	st2 = append(st2,ma[0])
	clt = append(clt,st2)

	var tmp [][]int
	for i := 2 ; i < len(ma) ; i++ {
		for j := 0 ; j < len(clt) ; j++ {
			// 组成新序列
			var cb []int
			cb = append(cb,ma[i])
			cb = append(cb,clt[j]...)
			for k := 0 ; k < len(cb) ; k++ {
				var cm []int
				cm = append(cm,cb[1:]...)
				cm = append(cm,cb[0])
				cb=cm
				tmp = append(tmp,cb)
			}
		}
		clt=tmp
	}

	return clt
}


func main() {

	// 5个已知条件
	conditions := [][]int{
		{2,4,6}, /* 条件一: 1个号码正确位置正确 */
		{2,5,8}, /* 条件二: 1个号码正确位置不正确 */
		{6,9,2}, /* 条件三: 2个号码正确位置都不正确 */
		{1,7,4}, /* 条件二: 条件四: 没有一个号码正确 */
		{4,1,9}, /* 条件五: 1个号码正确位置不正确*/
	}

	// 可能的开锁号码
	ca := []int{}
	cdt := []int{}
	m := make(map[int] bool)

	// 合并去重
	for _, v := range conditions {
		for _,e := range v {
			if _, ok := m[e]; !ok {
				cdt = append(cdt, e)
				m[e] = true
			}
		}

	}

	// 排除1，4，7
	for _, v := range cdt {
		if (v != 1) && (v != 4) && (v != 7)  {
			ca = append(ca,v)
		}
	}

	// 升序排序
	sort.Ints(ca[:])
	lenCa := len(ca)
	// 二维切片包含各种可能的组合-n取3
	var mtva [][]int
	for i := 0 ; i < lenCa ; i++ {
		for j := i+1 ; j < lenCa ; j++ {
			for k := j+1 ; k < lenCa ; k++ {
				tmp := []int{}
				tmp = append(tmp,ca[i])
				tmp = append(tmp,ca[j])
				tmp = append(tmp,ca[k])
				mtva = append(mtva,tmp)
			}
		}
	}


	for i := 0 ; i < len(mtva) ; i++ {
		// 排列组合
		va := sequence(mtva[i])
		answer := seekOut(va)
		if len(answer) > 0 {
			fmt.Print("开锁密码是:")
			fmt.Print(answer)
			os.Exit(0)
		}

	}


}
