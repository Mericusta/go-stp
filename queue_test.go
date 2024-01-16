package stp

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_PoolDequeue(t *testing.T) {
	var (
		wg            = &sync.WaitGroup{}
		producerCount = 1 << 16
		produceCount  = 1
		consumerCount = 1
		p             = NewPoolDequeue[int](producerCount * produceCount)
		producer      = func(wg *sync.WaitGroup, p PoolDequeue[int], v any) {
			if produceCount > 1 {
				for i := 0; i < produceCount; i++ {
					p.PushHead(i)
				}
			} else {
				p.PushHead(v)
			}
			wg.Done()
			// fmt.Printf("producer %v done\n", v)
		}
		consumer = func(wg *sync.WaitGroup, p PoolDequeue[int]) {
			i := 0
			for i < producerCount {
				v, ok := p.PopTail()
				if !ok {
					time.Sleep(time.Millisecond)
					fmt.Printf("consumer continue at %v\n", i)
					continue
				}
				fmt.Printf("consumer receive value %v, i = %v\n", v, i)
				if v == nil {
					fmt.Printf("consumer receive nil value at %v\n", i)
					continue
				}
				i++
				if i == producerCount*2 {
					panic("overload")
				}
			}
			fmt.Printf("consumer done\n")
			wg.Done()
		}
	)

	wg.Add(producerCount)

	// producer
	for i := 0; i < producerCount; i++ {
		go producer(wg, p, i)
	}

	wg.Wait()
	wg.Add(consumerCount)

	// consumer
	for i := 0; i < consumerCount; i++ {
		go consumer(wg, p)
	}

	wg.Wait()

	fmt.Printf("done\n")
}
