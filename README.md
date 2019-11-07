# golang-json-benchmark
Json benchmark for various golang JSON libraries


## Run
```bash
go test -bench=Stdlib -benchmem ./pkg/jsontest
```

Add `-run=xxx` to exclude tests from the execution.
