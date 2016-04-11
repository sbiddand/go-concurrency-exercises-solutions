# Limit your crawler

Given is a crawler (modified from the Go tour), which would pull pages excessively. The task is to modify the `main.go` file, so that

 - the crawler cannot crawl more than 2 pages per second
 - without losing it's concurrency

# Test your solution

Use `go test` to verify if your solution is correct.

Correct solution:
```
PASS
ok      github.com/mindworker/go-concurrency-exercises/0-limit-crawler  13.009s
```

Incorrect solution:
```
--- FAIL: TestMain (7.80s)
        main_test.go:18: There exists a two crawls who were executed less than 1 sec apart.
	        main_test.go:19: Solution is incorrect.
		FAIL
		exit status 1
		FAIL    github.com/mindworker/go-concurrency-exercises/0-limit-crawler  7.808s
```