package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

var (
	names = []string{"one", "two", "three"}
	outputs = []byte{}
)

func main() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())
	for _, name := range names {
		wg.Add(1)
		go func(n string) {
			out, err := exec.Command("php", "./test.php", n).Output()
			if err != nil {
				wg.Done()
				fmt.Printf("failed to execute: %s %s %s", "php", "./test.php", n)
				panic(err)
			}
			outputs = append(outputs, out...)
			wg.Done()
		}(name)
	}
	wg.Wait()

	for _, output := range outputs {
		fmt.Printf("%c", output)
	}
}
