package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//--------------指针---------------------------
	//fmt.Println("----func21--start---")
	//num := 10
	//func21(&num)
	//fmt.Println(num)
	//fmt.Println("----func21--end---")

	//fmt.Println("----func22--start---")
	//num := []int{1, 2, 3}
	//func22(&num)
	//fmt.Println(num)
	//fmt.Println("----func22--end---")

	//--------Goroutine-----------
	//fmt.Println("----func23--start---")
	//func23()
	//fmt.Println("----func23--end---")

	//fmt.Println("----func24--start---")
	//func24()
	//fmt.Println("----func24--end---")

	//--------------面向对象--------------
	//fmt.Println("----func25--start---")
	//func25()
	//fmt.Println("----func25--end---")

	//fmt.Println("----func26--start---")
	//func26()
	//fmt.Println("----func26--end---")

	//---------------Channel--------------------
	//fmt.Println("----func27--start---")
	//func27()
	//fmt.Println("----func27--end---")

	//fmt.Println("----func28--start---")
	//func28()
	//fmt.Println("----func28--end---")

	//---------------锁机制---------------------

	//fmt.Println("----func29--start---")
	//func29()
	//fmt.Println("----func29--end---")

	fmt.Println("----func30--start---")
	func30()
	fmt.Println("----func30--end---")
}

/*
*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func func21(num *int) {
	*num = *num + 10
}

/*
*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func func22(num *[]int) {
	temp := *num
	for i := range temp {
		temp[i] *= 2
	}
}

/*
*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func goWork1(index int) {
	for i := 1; i <= 10; i++ {
		if i%2 == index {
			fmt.Println("work", index, i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
func func23() {
	go goWork1(0)
	go goWork1(1)
	time.Sleep(2 * time.Second)
}

/*
*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func func24() {
	var wg sync.WaitGroup
	resultChan := make(chan string)
	wg.Add(2)                 //等待两个 done
	for j := 0; j <= 1; j++ { //for 循环，开两个线程
		go func() {
			defer wg.Done()
			start := time.Now()
			for i := 1; i <= 100; i++ {
				if i%2 == j {
					//fmt.Println("work", j, i)
					time.Sleep(10 * time.Millisecond)
				}
			}
			since := time.Since(start)
			time.Sleep(2 * time.Second)
			resultChan <- "work" + strconv.Itoa(j) + ":" + since.String()
		}()
	}
	go func() {
		wg.Wait()         //开线程等待done都提交
		close(resultChan) // 关闭通道
	}()
	for s := range resultChan { //循环展示
		fmt.Println(s)
	}
}

/*
*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
}
type Circle struct {
}

func (a Rectangle) Area() {
	fmt.Println("我是Rectangle的Area方法")
}
func (a Rectangle) Perimeter() {
	fmt.Println("我是Rectangle的Perimeter方法")
}
func (a Circle) Area() {
	fmt.Println("我是Circle的Area方法")
}
func (a Circle) Perimeter() {
	fmt.Println("我是Circle的Perimeter方法")
}

func func25() {
	var re Rectangle
	re.Area()
	re.Perimeter()
	var ci Circle
	ci.Area()
	ci.Perimeter()
}

/*
*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     Person
	EmployeeID string
}

func (a Employee) PrintInfo() {
	fmt.Println(a.Person.Name, a.Person.Age, a.EmployeeID)
}
func func26() {
	em := Employee{Person: Person{Name: "张三", Age: 18}, EmployeeID: "id123"}
	em.PrintInfo()
}

/*
*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func func27() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
	time.Sleep(2 * time.Second)
}

/*
*题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func func28() {
	ch := make(chan int, 10)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
	time.Sleep(10 * time.Second)
}

/*
*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func func29() {
	var mu sync.Mutex
	num := 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				num++
				mu.Unlock()
			}
		}()
	}
	time.Sleep(10 * time.Second)
	fmt.Println(num)
}

/*
*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func func30() {
	var num int64 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(num)
}
