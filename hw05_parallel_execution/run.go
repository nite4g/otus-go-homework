package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var errCount uint64

func consumer(ctx context.Context, wg *sync.WaitGroup, ch chan Task, m int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case t, ok := <-ch:

			if atomic.LoadUint64(&errCount) >= uint64(m) || !ok {
				//fmt.Printf("Routine: %+v  - %v\n", id, atomic.LoadUint64(&errCount))
				return
			}

			out := t()
			if out != nil {
				//  увеличиваем счетчик ошибок только в том случае если функция возвращает
				// то-то отличное от nil
				atomic.AddUint64(&errCount, 1)
			}
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	wg := &sync.WaitGroup{}
	taskChannel := make(chan Task)
	errCount = 0 // explicit init, avoid test races

	// Я читер :)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < n; i++ {
		wg.Add(1)
		go consumer(ctx, wg, taskChannel, m)
	}

	// producer
	for _, t := range tasks {
		if atomic.LoadUint64(&errCount) >= uint64(m) {
			close(taskChannel)
			break
		} else {
			taskChannel <- t
		}
	}

	cancel()
	wg.Wait()

	if atomic.LoadUint64(&errCount) >= uint64(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
