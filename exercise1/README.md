# Producer-Consumer Szenario

The producer reads in tweets from a mockstream and a consumer is processing the data to find out whether someone has tweeted about golang or not. The task is to modify the code inside `main.go` so that producer and consumer can run concurrently to increase the throughput of this program.

## Expected results:

Before: 
```
davecheney 	tweets about golang
beertocode 	does not tweet about golang
ironzeb 	tweets about golang
beertocode 	tweets about golang
vampirewalk666 	tweets about golang
Process took 2.499851159s
```

After:
```
davecheney 	tweets about golang
beertocode 	does not tweet about golang
ironzeb 	tweets about golang
beertocode 	tweets about golang
vampirewalk666 	tweets about golang
Process took 1.443451149s
```