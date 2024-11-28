package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type syncWorker struct {
	stop atomic.Bool
	lock sync.Mutex
	wg   sync.WaitGroup
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskChan := make(chan Task, len(tasks))
	syncWorker := syncWorker{}

	errorCount := 0

	worker := func(id int) {
		defer syncWorker.wg.Done()
		defer fmt.Printf("worker %d complete\n", id)

		for task := range taskChan {
			if syncWorker.stop.Load() {
				return
			}

			err := task()
			fmt.Printf("worker %d: complete task\n", id)

			if err != nil {
				syncWorker.lock.Lock()
				errorCount++
				if m > 0 && errorCount >= m {
					syncWorker.stop.Store(true)
					syncWorker.lock.Unlock()
					return
				}
				syncWorker.lock.Unlock()
			}
		}
	}

	// запускаем worker
	for i := 0; i < n; i++ {
		syncWorker.wg.Add(1)
		go worker(i)
	}

	// отдаем task в канал
	for _, task := range tasks {
		taskChan <- task
	}

	close(taskChan)

	syncWorker.wg.Wait()

	if syncWorker.stop.Load() {
		return ErrErrorsLimitExceeded
	}

	return nil
}
