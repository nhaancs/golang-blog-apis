package pblocal

import (
	"context"
	"log"
	"nhaancs/common"
	"nhaancs/pubsub"
	"sync"
)

// A pb run locally (in-mem)
// It has a queue (buffer channel) at it's core and many group of subscribers.
// Because we want to send a message with a specific topic for many subscribers in a group can handle.
type localPubSub struct {
	messageQueue chan *pubsub.Message
	topicsMap    map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		topicsMap:    make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, msg *pubsub.Message) error {
	msg.SetTopic(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- msg
		log.Println("New event published:", msg.String(), "with data", msg.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)
	ps.locker.Lock()
	ps.topicsMap[topic] = append(ps.topicsMap[topic], c)
	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")
		if chans, ok := ps.topicsMap[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.topicsMap[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}

}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")
	go func() {
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess)
			if subs, ok := ps.topicsMap[mess.Topic()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
