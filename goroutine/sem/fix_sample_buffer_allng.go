package main

import (
	"fmt"
	"sync"
)

func main() {
	if err := send(); err != nil {
		fmt.Printf("error occured. %v\n", err)
	}
	fmt.Println("finished")
}

const concurrency = 2

func send() error {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	numbers := []int{1, 2, 3, 4, 5, 6}

	sem := make(chan struct{}, concurrency)
	errs := make(chan error, 3)
	var wg sync.WaitGroup
	for _, num := range numbers {
		select {
		case sem <- struct{}{}:
		}
		// sem <- struct{}{}
		// defer func() { <-sem }()

		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			defer func() { <-sem }()
			if err := fnA(n); err != nil {
				errs <- fmt.Errorf("failed to A, %v", err)
				fmt.Printf("--> fnA len(errs) %d\n", len(errs))
				return
			}
			fmt.Println("--> succeeded to do A. num:", n)

			if err := fnB(n); err != nil {
				errs <- fmt.Errorf("failed to B, %v", err)
				fmt.Printf("--> fnB len(errs) %d\n", len(errs))
				return
			}
			fmt.Println("--> succeeded to do B. num:", n)
		}(num)
	}

	go func() {
		wg.Wait()
		close(sem)
		close(errs)
	}()
	if err, ok := <-errs; ok {
		return err
	}
	return nil
}

func fnA(n int) error {
	fmt.Printf("do fnA. \n")
	fmt.Printf("--> failed to do fnA.\n")
	return fmt.Errorf("error A. num: %d", n)
}
func fnB(n int) error {
	fmt.Printf("do fnB. \n")
	fmt.Printf("--> failed to do fnB.\n")
	return fmt.Errorf("error B. num: %d", n)
}
