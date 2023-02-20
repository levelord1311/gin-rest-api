package consumer

import (
	"gin-rest-api/internal/model"
	"sync"
	"time"
)

type EventRepo interface {
	Lock(n uint64) ([]model.UserEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.UserEvent) error
	Remove(eventIDs []uint64) error
}

type consumer struct {
	n      uint64
	events chan<- model.UserEvent

	repo EventRepo

	batchSize uint64
	timeout   time.Duration

	done chan bool
	wg   *sync.WaitGroup
}

func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo EventRepo,
	events chan<- model.UserEvent,
) *consumer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &consumer{
		n:         n,
		events:    events,
		repo:      repo,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		done:      done,
		wg:        wg,
	}
}

func (c *consumer) Start() {
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-c.done:
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	close(c.done)
	c.wg.Wait()
}
