package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

// ref: https://encore.dev/blog/advanced-go-concurrency
// 上のドキュメントに触発されて書いたテストコード

func TestErrSem(t *testing.T) {
	tests := map[string]struct {
		cities []string
		want   bool
	}{
		"1": {
			cities: []string{"a100", "b100", "c100", "d100"},
			want:   true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			infos, err := errSem(ctx, tt.cities)
			if err != nil {
				t.Fatal(err)
			}
			for _, v := range infos {
				fmt.Println(*v)
			}

		})
	}
}

// Info is city's information.
type Info struct {
	City         string
	TempC, TempF int    // temperature in Celsius and Farenheit
	Conditions   string // "sunny", "snowing", etc
}

var group singleflight.Group

func errSem(ctx context.Context, cities []string) ([]*Info, error) {
	start := time.Now()
	defer func() {
		log.Printf("Total time: %fs", time.Since(start).Seconds())
	}()
	eg, ctx := errgroup.WithContext(ctx)
	var mu sync.Mutex
	res := make([]*Info, len(cities)) // res[i] corresponds to cities[i]
	sem := semaphore.NewWeighted(1)   // 100 chars processed concurrently
	for i, city := range cities {
		i, city := i, city // create locals for closure below
		// cost := int64(len(city))
		cost := int64(1)
		log.Println("city cost", city, cost)
		if err := sem.Acquire(ctx, cost); err != nil {
			fmt.Println("failed to acquire:", err)
			return nil, err
		}
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			info, err := City(ctx, city)
			mu.Lock()
			log.Println("city info", city, info)
			res[i] = info
			mu.Unlock()
			sem.Release(cost)
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

func City(ctx context.Context, city string) (*Info, error) {
	time.Sleep(time.Second)
	return &Info{
		City:       city,
		TempC:      10,
		TempF:      32,
		Conditions: "sunny",
	}, nil
	// results, err, _ := group.Do(city, func() (interface{}, error) {
	// 	info, err := fetchWeatherFromDB(city) // slow operation
	// 	return info, err
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("weather.City %s: %w", city, err)
	// }
	// return results.(*Info), nil
}

// func fetchWeatherFromDB(city string) (*Info, error) {
// 	time.Sleep
// 	return Info, nil
// }
