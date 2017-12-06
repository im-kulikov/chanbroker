package main

import (
	"fmt"
	"github.com/myself659/ChanBroker"
	"strconv"
	"time"
)

type event struct {
	id   int
	info string
}

func main() {
	b := ChanBroker.NewChanBroker(time.Second)

	sub1 := b.RegSubscriber(0)
	sub2 := b.RegSubscriber(2)
	sub3 := b.RegSubscriber(4)
	go func(sub1 ChanBroker.Subscriber, b *ChanBroker.ChanBroker) {
		i := 0
		for c := range sub1 {
			switch t := c.(type) {
			case string:
				fmt.Println(sub1, "string:", t)
				i++
				if i == 1 {
					b.UnRegSubscriber(sub1)
				}
			case int:
				fmt.Println(sub1, "int:", t)

			default:

			}
		}

	}(sub1, b)

	go func(sub2 ChanBroker.Subscriber, b *ChanBroker.ChanBroker) {
		for c := range sub2 {
			switch t := c.(type) {
			case string:
				fmt.Println(sub2, "string:", t)
			case int:
				fmt.Println(sub2, "int:", t)
			case event:
				fmt.Println(sub2, "event:", t)
			default:

			}
		}

	}(sub2, b)

	go func(sub3 ChanBroker.Subscriber, b *ChanBroker.ChanBroker) {
		c := <-sub3
		switch t := c.(type) {
		case string:
			fmt.Println(sub3, "string:", t)
		case int:
			fmt.Println(sub3, "int:", t)
		case event:
			fmt.Println(sub3, "event:", t)
		default:

		}

		//b.UnRegSubscriber(sub3)
	}(sub3, b)

	ticker := time.NewTicker(time.Second)
	go func() {
		var prefix string = "pub_"
		i := 0
		for range ticker.C {
			i++
			if i%3 == 0 {
				var temp string
				temp = prefix + strconv.Itoa(i)
				b.PubContent(temp)
			} else if i%3 == 1 {
				b.PubContent(i)
			} else {
				ev := event{i, "event"}
				b.PubContent(ev)
			}

		}
	}()

	<-time.After(7 * time.Second)
}
