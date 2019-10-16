package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func printStringChars(str string, wg *sync.WaitGroup) {
	chars := strings.Split(str, "")
	for i := range chars {
		fmt.Print(chars[i])
	}
	wg.Done()
}

func main()  {
	go func(str string) {
		chars := strings.Split(str, "")
		for i := range chars {
			fmt.Print(chars[i])
		}
	}(lower)
	go func(str string) {
		chars := strings.Split(str, "")
		for i := range chars {
			fmt.Print(chars[i])
		}
	}(upper)
	time.Sleep(time.Second)
	fmt.Println("\nDone with standalone routines\n")


	var wg sync.WaitGroup
	wg.Add(1)
	printStringChars(lower, &wg)
	wg.Add(1)
	printStringChars(upper, &wg)
	wg.Wait()
	fmt.Println("\nDone with WaitGroup\n")


	channel := make(chan string, 56)
	wg.Add(1)
	lowChars := strings.Split(lower, "")
	upperChars := strings.Split(upper, "")

	printStrings(lowChars, upperChars, channel, &wg)
	wg.Wait()
	close(channel)
	for item := range channel {
		fmt.Print(item)
	}
	fmt.Println("\nDone with channels")
}

func printStrings(lowChars []string, upperChars []string, channel chan<- string, group *sync.WaitGroup) {
	defer group.Done()
	for i := range lowChars {
		channel <- lowChars[i]
		channel <- upperChars[i]
	}
}