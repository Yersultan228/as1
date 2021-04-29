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

//declaration of shared channel for producer and consumer
var tweetsCh chan *Tweet

func producer(stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			//only receiver should close channel
			close(tweetsCh)
			return
		}
		//sending tweets to channel
		tweetsCh <- tweet
	}
}

func consumer(tweetsChannel <-chan *Tweet) {
	for {
		//receiving tweets from channel
		t, err := <-tweetsChannel

		//break loop used for listening channel if no signals/data send
		if !err {
			break
		}

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

	//creation of channel
	tweetsCh = make(chan *Tweet)

	// Producer - runs on separate goroutine
	go producer(stream)

	// Consumer
	consumer(tweetsCh)

	fmt.Printf("Process took %s\n", time.Since(start))
}
