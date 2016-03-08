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

func main() {
	start := time.Now()

	var tweets []*Tweet

	stream := GetMockStream()

	// Producer
	for {
		tweet, err := stream.Next()
		if err == EOF {
			break
		}

		if len(tweets) == 0 {
			tweets = []*Tweet{tweet}
		} else {
			tweets = append(tweets, tweet)
		}
	}

	// Consumer
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}

	fmt.Printf("Process took %s\n", time.Since(start))
}
