# Limit your crawler

Given is a crawler (modified from the Go tour), which would pull pages excessively. The task is to modify the `main.go` file, so that

 - the crawler cannot crawl more than 2 pages per second
 - without losing it's concurrency