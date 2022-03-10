package main

import "fmt"

func main() {
	loop1()
	loop2()
	loop3()
	loop4()
	loop5()
}

// for initialization; condition; post {}
func loop1() {
	fmt.Println("loop1 start.")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(list); i++ {
		fmt.Print(list[i])
	}
	fmt.Println("\nloop1 end.")
}

// for range {}
func loop2() {
	fmt.Println("loop2 start.")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := range list {
		fmt.Print(list[i])
	}
	fmt.Println("\nloop2 end.")
}

// for range {}
func loop3() {
	fmt.Println("loop3 start.")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range list {
		fmt.Print(v)
	}
	fmt.Println("\nloop3 end.")
}

// for condition {}
func loop4() {
	fmt.Println("loop4 start.")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for len(list) > 0 {
		fmt.Print(list[0]) // 先頭を取り出し
		list = list[1:]    // 2番目移行を切り出し
	}
	fmt.Println("\nloop4 end.")
}

// for {} 無限ループ
func loop5() {
	fmt.Println("loop5 start.")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for {
		fmt.Print(list[0]) // 先頭を取り出し
		list = list[1:]    // 2番目移行を切り出し
		if len(list) == 0 {
			break
		}
	}
	fmt.Println("\nloop5 end.")
}
