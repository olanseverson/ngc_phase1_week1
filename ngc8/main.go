// ====================== SOAL 5 ========================

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func countWords(s string) int {
	// open the file
	file, err := os.Open("./" + s)
	if err != nil {
		fmt.Println("Error", err)
		return -1
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	var cntWords int
	for scanner.Scan() {
		line := scanner.Text()
		cntWords += len(line)
	}

	// check for error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("errorrr:", err)
		return -1
	}
	return cntWords
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg sync.WaitGroup

	wg.Add(len(files))

	countCH := make(chan map[string]int, 3)
	for _, v := range files {
		go func(wg *sync.WaitGroup, v string, c chan map[string]int) {
			defer wg.Done()
			data := map[string]int{v: countWords(v)}
			c <- data
		}(&wg, v, countCH)
	}

	wg.Wait()
	close(countCH)

	for val := range countCH {
		fmt.Println(val)
	}

}

// ====================== SOAL 4 ========================
// package main

// import (
// 	"fmt"
// 	"math/rand/v2"
// 	"sync"
// 	"time"
// )

// func calculateSumofSquares(num []int, numGoRoutine int) int {
// 	var wg sync.WaitGroup

// 	wg.Add(numGoRoutine)

// 	// divide the slice of num equally
// 	var dividedSlice [][]int
// 	for i := 1; i <= numGoRoutine; i++ {
// 		dividedSlice = append(dividedSlice, make([]int, 0))
// 	}
// 	for i := 0; i < len(num); i++ {
// 		dividedSlice[i%numGoRoutine] = append(dividedSlice[i%numGoRoutine], num[i])
// 	}
// 	// fmt.Println(dividedSlice)

// 	// create channel for receiving the calculated data
// 	dividedChan := make(chan int, numGoRoutine)

// 	// run 'n' number of go routine according to 'numGoRoutine' t4 execute divided slice of num
// 	for i := 1; i <= numGoRoutine; i++ {
// 		go func(wg *sync.WaitGroup, data []int, ch chan int) {
// 			defer wg.Done()
// 			var total int
// 			for _, v := range data {
// 				total += v * v
// 			}
// 			dividedChan <- total

// 		}(&wg, dividedSlice[i-1], dividedChan)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(dividedChan)
// 	}()

// 	var combinetotal int
// 	for msg := range dividedChan {
// 		combinetotal += msg
// 	}

// 	return combinetotal
// }

// func main() {
// 	// rand.Seed(time.Now().UnixNano())
// 	var numbers []int
// 	for i := 0; i < 5000000; i++ {
// 		numbers = append(numbers, rand.IntN(100))
// 	}
// 	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}
// 	x := time.Now()
// 	a := calculateSumofSquares(numbers, 2)
// 	y := time.Since(x)
// 	fmt.Println(y.Microseconds())

// 	x = time.Now()
// 	b := calculateSumofSquares(numbers, 50)
// 	y = time.Since(x)
// 	fmt.Println(y.Microseconds())
// 	fmt.Println(a, b)
// }

// ====================== SOAL 3 ========================

// package main

// import (
// 	"fmt"
// 	"math"
// )

// const (
// 	RECTANGLE = "rectangle"
// 	CIRCLE    = "circle"
// 	TRIANGLE  = "triangle"
// )

// type Shape struct {
// 	ShapeType string
// 	Length    int
// 	Area      float64
// }

// func main() {
// 	// var wg sync.WaitGroup

// 	input := []Shape{
// 		{ShapeType: RECTANGLE, Length: 5},
// 		{ShapeType: CIRCLE, Length: 3},
// 		{ShapeType: TRIANGLE, Length: 5},
// 		{ShapeType: RECTANGLE, Length: 15},
// 		{ShapeType: CIRCLE, Length: 5},
// 	}

// 	rect := make(chan Shape)
// 	cir := make(chan Shape)
// 	tri := make(chan Shape)

// 	// wg.Add(1)
// 	go func() {
// 		defer close(rect)
// 		defer close(cir)
// 		defer close(tri)
// 		for _, v := range input {
// 			switch v.ShapeType {
// 			case RECTANGLE:
// 				rect <- v
// 			case CIRCLE:
// 				cir <- v
// 			case TRIANGLE:
// 				tri <- v
// 			}
// 		}
// 	}()

// 	for {
// 		select {
// 		case s, ok := <-rect:
// 			if ok {
// 				s.Area = float64(s.Length) * float64(s.Length)
// 				fmt.Println("Rect: ", s)
// 			} else {
// 				rect = nil
// 			}
// 		case s, ok := <-cir:
// 			if ok {
// 				s.Area = math.Pi * float64(s.Length) * float64(s.Length)
// 				fmt.Println("Circle: ", s)
// 			} else {
// 				cir = nil
// 			}
// 		case s, ok := <-tri:
// 			if ok {
// 				s.Area = 0.5 * float64(s.Length) * float64(s.Length)
// 				fmt.Println("Triangle: ", s)
// 			} else {
// 				tri = nil
// 			}
// 		}

// 		if rect == nil && cir == nil && tri == nil {
// 			break
// 		}
// 	}
// }

// ====================== SOAL 2 ========================

// package main

// import (
// 	"fmt"
// )

// func SumSquare(num int, result chan int) {
// 	var total int
// 	for i := 0; i < num; i++ {
// 		total += i * i
// 	}
// 	// fmt.Println(total)
// 	result <- total
// }

// func SquareSum(num int, result chan int) {
// 	var total int
// 	for i := 0; i < num; i++ {
// 		total += i
// 	}
// 	// fmt.Println(total)
// 	result <- total * total
// }

// func main() {

// 	sumsquare := make(chan int)
// 	squaresum := make(chan int)

// 	go SumSquare(100, sumsquare)
// 	go SquareSum(100, squaresum)

// 	fmt.Println("SumSquare", <-sumsquare)
// 	fmt.Println("SquareSum", <-squaresum)

// 	close(sumsquare)
// 	close(squaresum)

// }

// ====================== SOAL 1 ========================
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func updateChan(wg *sync.WaitGroup, num chan<- int) {
// 	defer wg.Done()
// 	for i := 0; i < 100; i++ {
// 		fmt.Println("-x-")
// 		num <- i
// 	}

// }

// func fizzbuzz(x int) {
// 	if x%3 == 0 && x%5 == 0 {
// 		fmt.Printf("%vFizzBuzz\n", x)
// 	} else if x%3 == 0 {
// 		fmt.Printf("%vFizz\n", x)
// 	} else if x%5 == 0 {
// 		fmt.Printf("%vBuzz\n", x)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	num := make(chan int)

// 	wg.Add(1)
// 	go updateChan(&wg, num)

// 	go func() {
// 		wg.Wait()
//      close(num)
// 	}()

// 	var total int
// 	var countEven int
// 	for msg := range num {
// 		go fizzbuzz(msg)
// 		total += msg
// 		if msg%2 == 0 {
// 			countEven++
// 		}
// 	}

// 	fmt.Println("total: ", total)
// 	fmt.Println("bilangan genap:", countEven)
// 	fmt.Println("bilangan ganjil:", 100-countEven)
// 	fmt.Println("all task completed")
// }
