// https://austburn.me/blog/a-better-fan-in-fan-out-example.html
package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

// Item is
type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func prepareItems(ctx context.Context) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		Item{0, "Shirt", 1 * time.Second},
		Item{1, "Legos", 1 * time.Second},
		Item{2, "TV", 5 * time.Second},
		Item{3, "Bananas", 2 * time.Second},
		Item{4, "Hat", 1 * time.Second},
		Item{5, "Phone", 2 * time.Second},
		Item{6, "Plates", 3 * time.Second},
		Item{7, "Computer", 5 * time.Second},
		Item{8, "Pint Glass", 3 * time.Second},
		Item{9, "Watch", 2 * time.Second},
	}
	go func() {
		for _, item := range itemsToShip {
			select {
			case <-ctx.Done():
				return
			case items <- item:
			}
		}
		close(items)
	}()
	return items
}

func packItems(ctx context.Context, items <-chan Item, workerID int) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-ctx.Done():
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
				fmt.Printf("Worker #%d: Shipping package no. %d,took %ds to pack\n", workerID, item.ID, item.PackingEffort/time.Second)
			}
		}
		close(packages)
	}()
	return packages
}

func merge(ctx context.Context, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	outgoingPackages := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-ctx.Done():
				return
			case outgoingPackages <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(outgoingPackages)
	}()
	return outgoingPackages
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	start := time.Now()

	log.Printf("cpu: %d\n", runtime.NumCPU())
	log.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	items := prepareItems(ctx)
	sem := 8
	workers := make([]<-chan int, sem)
	for i := 0; i < sem; i++ {
		workers[i] = packItems(ctx, items, i)
	}

	numPackages := 0
	for range merge(ctx, workers...) {
		numPackages++
		//fmt.Printf("Shipping package no. %d\n", p)
	}

	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
