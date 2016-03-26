# Race condition in caching szenario

Given is some code to cache key-value pairs from a mock database into
the main memory (to reduce access time). The code is buggy and
contains a race conditition. Change the code to make this thread safe.

Use the following command to test for race conditions:
```
go test -race
```

*Note*: that golang's map are not entirely thread safe. Multiple readers
are fine, but multiple writers are not.
