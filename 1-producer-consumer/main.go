//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, ch chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		ch <- tweet
	}
}

func consumer(ch <-chan *Tweet) {
	for t := range ch {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	prodChan := make(chan *Tweet)

	// Producer
	go func() {
		defer close(prodChan)
		producer(stream, prodChan)
	}()

	// Consumer
	consumer(prodChan)

	fmt.Printf("Process took %s\n", time.Since(start))
}
