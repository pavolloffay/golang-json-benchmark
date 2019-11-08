# Golang JSON benchmark
Golang JSON benchmark for Jaeger Elasticsearch span model.

## Run benchmark
```bash
make bench
```

## Understanding benchmark results
Name, average [ns] of a run, average allocated B of memory per run, average number of allocations per run
```
BenchmmarkXX    44541     27849 ns/op    29440 B/op     158 allocs/op
```

## Conclusions

* stdlib marshalling performance drops when unmarshalling `map[string]interface{}`
