package main

import "fmt"

// 将切片通过参数传递给函数或者赋值操作，其实质是复制了slice结构体对象，两个slice结构体的字段值均相等。正常情况下，由于函数内slice结构体的array和函数外slice结构体的array指向的是同一底层数组，所以当对底层数组中的数据做修改时，两者均会受到影响。
// 但是存在这样的问题：如果指向底层数组的指针被覆盖或者修改（copy、重分配、append触发扩容），此时函数内部对数据的修改将不再影响到外部的切片，代表长度的len和容量cap也均不会被修改。
// 如果你只想修改切片中元素的值，而不会更改切片的容量与指向，则可以按值传递切片，否则你应该考虑按指针传递
func modifySlice(innerSlice []string) {
	fmt.Printf("%p %v   %p\n", &innerSlice, innerSlice, &innerSlice[0])
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
}

func main() {
	k1 := make([]int,1)
	k1[0] = 1
	fmt.Printf("%p %v   %p\n", &k1, k1, &k1[0])
	var  k2 []int
	k2 = k1
	fmt.Printf("%p %v   %p\n", &k2, k2, &k2[0])
	k2 = append(k2,2)
	k2[0] = 2
	fmt.Printf("%p %v   %p\n", &k2, k2, &k2[0])
	fmt.Println(k1)


	outerSlice := []string{"a", "a"}
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice[0])
	modifySlice(outerSlice)
	fmt.Printf("%p %v   %p\n", &outerSlice, outerSlice, &outerSlice[0])

	outerSlice1:= make([]string, 0, 3)
	outerSlice1 = append(outerSlice1, "a", "a")
	modifySlice1(outerSlice1)
	fmt.Println(outerSlice1)
}

func modifySlice1(innerSlice []string) {
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}