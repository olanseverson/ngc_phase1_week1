package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

// ++++++++++++++++++ SOAL 5 +++++++++++++++++
// func downloadFile(url string, destination string) {
// 	outFile, err := os.Create(destination)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer outFile.Close()

// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		log.Printf("server returned non-200 status code : %v\n", response.StatusCode)
// 		return
// 	}

// 	_, err = io.Copy(outFile, response.Body)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	link := []string{
// 		"https://github.com/olanseverson/shashin/blob/master/asset/images/jpeg/1.jpg",
// 		"https://github.com/olanseverson/shashin/blob/master/asset/images/jpeg/2.jpg",
// 		"https://github.com/olanseverson/shashin/blob/master/asset/images/jpeg/3.jpg",
// 	}

// 	for i, v := range link {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			path := fmt.Sprintf("./download/%d.jpg", i)
// 			fmt.Println(v, path)
// 			downloadFile(v, path)
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println("all the links downloaded succesfuly")
// }

// ++++++++++++++++++ SOAL 4 +++++++++++++++++
// func scheduleTask(task func(float64), wg *sync.WaitGroup, param float64) {
// 	defer wg.Done()
// 	fmt.Printf("this is function %p\n", task)
// 	task(param)
// }

// func diagonal(l float64) {
// 	fmt.Printf("Diagonal is : %.2f\n", l*math.Sqrt(2))
// }

// func area(l float64) {
// 	fmt.Printf("Area is is : %.2f\n", l*l)
// }

// func volume(l float64) {
// 	fmt.Printf("Volume is : %.2f\n", l*l*l)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	fmt.Println("Enter length of a square:")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)
// 	x, ok := strconv.Atoi(input)
// 	if ok != nil {
// 		return
// 	}

// 	wg.Add(3)
// 	go scheduleTask(diagonal, &wg, float64(x))
// 	go scheduleTask(area, &wg, float64(x))
// 	go scheduleTask(volume, &wg, float64(x))

// 	time.Sleep(3 * time.Second) // represents another function or processes
// 	wg.Wait()
// 	fmt.Println("programs end.")
// }

// ++++++++++++++++++ SOAL 3 +++++++++++++++++

// func countScrabble(in string, ch chan int) {
// 	var count int
// 	for _, v := range in {
// 		char := string(v)
// 		switch char { // switch case value is dummy here, but you got the idea
// 		case "A":
// 			count += 1
// 		case "B":
// 			count += 2
// 		case "C":
// 			count += 3
// 		case "D":
// 			count += 4
// 		default:
// 			count += 5
// 		}
// 	}
// 	// fmt.Println("count1")
// 	ch <- count
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("---")
// }

// func main() {
// 	var wg sync.WaitGroup

// 	ch := make(chan int, 10)

// 	for {
// 		fmt.Println("input string:")
// 		reader := bufio.NewReader(os.Stdin)
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		input = strings.ToLower(input)
// 		if input == "done" {
// 			break
// 		}

// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			countScrabble(input, ch)
// 		}()

// 		out := <-ch
// 		fmt.Println("THIS IS SCORE FOR ", input, " ->", out)
// 	}

// 	wg.Wait()
// 	fmt.Println("end of program")
// }

// ++++++++++++++++++ SOAL 2 +++++++++++++++++

func acronimify(input string) {
	inputParsed := regexp.MustCompile(`[^a-zA-Z ]+`).ReplaceAllString(input, " ") // remove the nonchar
	// fmt.Println(inputParsed)
	slice := strings.Split(inputParsed, " ")
	// fmt.Println(len(slice))
	var out []string
	for i := 0; i < len(slice); i++ {
		out = append(out, string(slice[i][0]))
	}
	fmt.Println(input, " - ", strings.Join(out, ""))
}

func main() {
	var wg sync.WaitGroup
	for {
		fmt.Print("input the text: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println(input)
		if input == "done" {
			break
		}

		wg.Add(1)
		go func(in string) {
			defer wg.Done()
			acronimify(in)
		}(input)
	}

	wg.Wait()
	fmt.Println("All the process completed")

}

// ++++++++++++++++++ SOAL 1 +++++++++++++++++
// func fibonacci(x int) int {
// 	if x < 3 {
// 		return x - 1
// 	} else {
// 		return fibonacci(x-2) + fibonacci(x-1)
// 	}
// }

// func main for soal1
// func main() {
// 	var wg sync.WaitGroup

// 	// start go routine
// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Printf("Iteration: %v -> %v\n", i, fibonacci(i))
// 		}()
// 	}

// 	// wait for all goroutines have finished
// 	wg.Wait()
// 	fmt.Println("all goroutines have finished")
// }
