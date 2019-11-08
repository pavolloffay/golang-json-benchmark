BENCHTIME?=1s
GOMAXPROCS?=1

.PHONY:test
test:
	go test ./pkg/jsontest

.PHONY:bench
bench:
	GOMAXPROCS=${GOMAXPROCS} go test -bench=. -test.benchtime=${BENCHTIME} -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest

.PHONY:profile
profile:
	go tool pprof -http=:8080 profile_cpu.out

.PHONY:pprof
pprof:
	echo "Run this to profile Jaeger services"
	go tool pprof  -http=:8080 http://localhost:14269/debug/pprof/profile\?seconds\=15
