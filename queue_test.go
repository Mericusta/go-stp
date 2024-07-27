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
		producerCount = 1 << 3
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
			fmt.Printf("producer %v done\n", v)
			wg.Done()
		}
		consumer = func(wg *sync.WaitGroup, p PoolDequeue[int]) {
			i := 0
			for i < producerCount-1 {
				v, ok := p.PopTail()
				if !ok {
					time.Sleep(time.Millisecond)
					fmt.Printf("consumer continue, i = %v\n", i)
					continue
				}
				fmt.Printf("consumer receive value %v, i = %v\n", v, i)
				if v == nil {
					fmt.Printf("consumer receive nil value, i = %v\n", i)
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

	// producer
	wg.Add(producerCount - 1)
	// TODO: 这里必须全部塞满，底下 consumer 才会正确读取
	for i := 0; i < producerCount-1; i++ {
		go producer(wg, p, i)
	}
	wg.Wait()

	// consumer
	wg.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go consumer(wg, p)
	}
	wg.Wait()

	fmt.Printf("done\n")
}

func Test_PoolChain(t *testing.T) {
	var (
		wg            = &sync.WaitGroup{}
		producerCount = 1 << 3
		produceCount  = 1
		consumerCount = 1
		p             = NewPoolChain[int]()
		producer      = func(wg *sync.WaitGroup, p PoolDequeue[int], v any) {
			if produceCount > 1 {
				for i := 0; i < produceCount; i++ {
					p.PushHead(i)
				}
			} else {
				p.PushHead(v)
			}
			fmt.Printf("producer %v done\n", v)
			wg.Done()
		}
		consumer = func(wg *sync.WaitGroup, p PoolDequeue[int]) {
			i := 0
			for i < producerCount-1 {
				v, ok := p.PopTail()
				if !ok {
					time.Sleep(time.Millisecond)
					fmt.Printf("consumer continue, i = %v\n", i)
					continue
				}
				fmt.Printf("consumer receive value %v, i = %v\n", v, i)
				if v == nil {
					fmt.Printf("consumer receive nil value, i = %v\n", i)
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

	// producer
	wg.Add(producerCount - 1)
	for i := 0; i < producerCount-1; i++ {
		go producer(wg, p, i)
	}
	wg.Wait()

	// consumer
	wg.Add(consumerCount)
	for i := 0; i < consumerCount; i++ {
		go consumer(wg, p)
	}
	wg.Wait()

	fmt.Printf("done\n")
}
