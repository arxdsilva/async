package async

import (
	"context"
	"log"
	"runtime"
	"sync"
)

type WorkFunc func(msg interface{}) error

// Async opens the number of goroutines equals to the available CPUs
// then ensures that all data is sent to them and consumed properly in the provided WorkFunc
// if any error rises it will be sent to the error channel
func Async(ctx context.Context, fn WorkFunc, data <-chan interface{}, errChan chan<- error) {
	cpus := runtime.NumCPU()
	var wg sync.WaitGroup
	for i := 0; i < cpus; i++ {
		wg.Add(1)
		go work(ctx, &wg, fn, data, errChan)
	}
	log.Printf("async: job started with %v processors\n", cpus)
	wg.Wait()
	log.Println("async: jobs stopped")
}

func work(ctx context.Context, wg *sync.WaitGroup, fn WorkFunc, data <-chan interface{}, errChan chan<- error) {
	defer wg.Done()
	select {
	case msg := <-data:
		err := fn(msg)
		if err != nil {
			log.Printf("async: error %v \n", err.Error())
			errChan <- err
		}
	case <-ctx.Done():
		log.Println("async: context ended")
		return
	}
}

// New starts a async func executor and returns the necessary channels
// options are just placeholders for future features
// channels are always buffered, at least for now
//
// possible future options:
// 1. buffered or unbuffered
// 2. logs
func New(ctx context.Context, fn WorkFunc, options ...Option) (dataChan chan interface{}, errChan chan error) {
	cfg := &config{
		chanSize: 10,
	}
	for _, option := range options {
		option(cfg)
	}
	dataChan = make(chan interface{}, cfg.chanSize)
	errChan = make(chan error, cfg.chanSize)
	go Async(ctx, fn, dataChan, errChan)
	return
}
