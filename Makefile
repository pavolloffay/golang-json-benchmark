
.PHONY:test
test:
	go test ./pkg/jsontest

.PHONY:bench
bench:
	go test -bench=. -benchmem -run=XX ./pkg/jsontest
