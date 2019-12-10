[![Build Status][ci-img]][ci]

# Golang JSON benchmark
Golang JSON benchmark for Jaeger Elasticsearch span model.

## Run benchmark
```bash
make bench
```

## Understanding benchmark results
Name, number of runs of the loop per `-test.benchtime`, average [ns] of a run, average allocated B of memory per run, average number of allocations per run
```
BenchmmarkXX    44541     27849 ns/op    29440 B/op     158 allocs/op
```

## Test results
Tests results can be found on Github CI/Actions page or I provide here results from my local runs
on Lenovo P50 i7-6820HQ CPU @ 2.70GHz.

```
GOMAXPROCS=1 go test -bench=Marshal -test.benchtime=3s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkMarshalGojay/default.json         	  516730	      6654 ns/op	    4240 B/op	       9 allocs/op
BenchmarkMarshalGojay/default-unicode.json 	  904315	      3617 ns/op	    1808 B/op	       8 allocs/op
BenchmarkMarshalGojay/default-tagmap.json  	  802406	      4899 ns/op	    2607 B/op	       8 allocs/op
BenchmarkMarshalStdlib/default.json        	  558175	      5728 ns/op	    1536 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-unicode.json         	  768222	      4183 ns/op	    1024 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-tagmap.json          	  362587	      8962 ns/op	    2752 B/op	      37 allocs/op
BenchmarkMarshalJsoniter/default.json               	  499992	      6932 ns/op	    1800 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-unicode.json       	  612794	      5121 ns/op	    1288 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-tagmap.json        	  245335	     13910 ns/op	    7930 B/op	      50 allocs/op
BenchmarkMarshalJettison/default.json      	  399955	      8176 ns/op	    1744 B/op	      14 allocs/op
BenchmarkMarshalJettison/default-unicode.json         	  598621	      5165 ns/op	    1152 B/op	       9 allocs/op
BenchmarkMarshalJettison/default-tagmap.json          	  323354	     10128 ns/op	    2719 B/op	      27 allocs/op

GOMAXPROCS=1 go test -bench=Unmarshal -test.benchtime=3s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkUnmarshalFastjson/default.json         	  170601	     21980 ns/op	   29440 B/op	     158 allocs/op
BenchmarkUnmarshalFastjson/default-unicode.json 	  235021	     14834 ns/op	   16160 B/op	     115 allocs/op
BenchmarkUnmarshalFastjson/default-tagmap.json  	  312003	     10821 ns/op	   14416 B/op	      54 allocs/op
BenchmarkUnmarshalGojay/default.json            	  148288	     22983 ns/op	   10159 B/op	      75 allocs/op
BenchmarkUnmarshalGojay/default-unicode.json    	  237422	     14631 ns/op	    6818 B/op	      49 allocs/op
BenchmarkUnmarshalGojay/default-tagmap.json     	  266962	     12868 ns/op	    5377 B/op	      61 allocs/op
BenchmarkUnmarshalStdlib/default.json           	  103725	     32835 ns/op	    1520 B/op	      73 allocs/op
BenchmarkUnmarshalStdlib/default-unicode.json   	  159901	     23717 ns/op	    1136 B/op	      53 allocs/op
BenchmarkUnmarshalStdlib/default-tagmap.json    	  151378	     22602 ns/op	    2192 B/op	      95 allocs/op
BenchmarkUnmarshalJsoninter/default.json        	  417264	      8225 ns/op	    1184 B/op	      76 allocs/op
BenchmarkUnmarshalJsoninter/default-unicode.json         	  518607	      5856 ns/op	     816 B/op	      56 allocs/op
BenchmarkUnmarshalJsoninter/default-tagmap.json          	  480778	      6971 ns/op	    1456 B/op	      85 allocs/op
```

4 CPUs
```
GOMAXPROCS=4 go test -bench=Marshal -test.benchtime=3s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkMarshalGojay/default.json-4         	  661621	      4624 ns/op	    4240 B/op	       9 allocs/op
BenchmarkMarshalGojay/default-unicode.json-4 	 1000000	      3157 ns/op	    1808 B/op	       8 allocs/op
BenchmarkMarshalGojay/default-tagmap.json-4  	  868977	      3669 ns/op	    2608 B/op	       8 allocs/op
BenchmarkMarshalStdlib/default.json-4        	  565506	      5802 ns/op	    1536 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-unicode.json-4         	  800090	      4386 ns/op	    1024 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-tagmap.json-4          	  415911	      8413 ns/op	    2752 B/op	      37 allocs/op
BenchmarkMarshalJsoniter/default.json-4               	  529898	      6596 ns/op	    1800 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-unicode.json-4       	  615393	      4976 ns/op	    1288 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-tagmap.json-4        	  345463	     10480 ns/op	    7933 B/op	      50 allocs/op
BenchmarkMarshalJettison/default.json-4      	  413160	      8200 ns/op	    1744 B/op	      14 allocs/op
BenchmarkMarshalJettison/default-unicode.json-4         	  640046	      5256 ns/op	    1152 B/op	       9 allocs/op
BenchmarkMarshalJettison/default-tagmap.json-4          	  351909	      9658 ns/op	    2721 B/op	      27 allocs/op
GOMAXPROCS=4 go test -bench=Unmarshal -test.benchtime=3s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkUnmarshalFastjson/default.json-4         	  161304	     20712 ns/op	   29440 B/op	     158 allocs/op
BenchmarkUnmarshalFastjson/default-unicode.json-4 	  251964	     13733 ns/op	   16160 B/op	     115 allocs/op
BenchmarkUnmarshalFastjson/default-tagmap.json-4  	  332452	      9770 ns/op	   14416 B/op	      54 allocs/op
BenchmarkUnmarshalGojay/default.json-4            	  194344	     18322 ns/op	    9983 B/op	      75 allocs/op
BenchmarkUnmarshalGojay/default-unicode.json-4    	  285207	     11772 ns/op	    6932 B/op	      49 allocs/op
BenchmarkUnmarshalGojay/default-tagmap.json-4     	  287793	     12263 ns/op	    5470 B/op	      61 allocs/op
BenchmarkUnmarshalStdlib/default.json-4           	  123186	     29579 ns/op	    1520 B/op	      73 allocs/op
BenchmarkUnmarshalStdlib/default-unicode.json-4   	  158065	     21486 ns/op	    1136 B/op	      53 allocs/op
BenchmarkUnmarshalStdlib/default-tagmap.json-4    	  146869	     21530 ns/op	    2192 B/op	      95 allocs/op
BenchmarkUnmarshalJsoninter/default.json-4        	  379986	      8116 ns/op	    1184 B/op	      76 allocs/op
BenchmarkUnmarshalJsoninter/default-unicode.json-4         	  571060	      5947 ns/op	     816 B/op	      56 allocs/op
BenchmarkUnmarshalJsoninter/default-tagmap.json-4          	  523004	      6824 ns/op	    1456 B/op	      85 allocs/op
```

## Conclusions

* Fastjson does not provide a way to unmarshal `interface{}` in tag map, therefore it should not be considered in results
* stdlib marshalling performance drops (especially allocations increase) when unmarshalling tag map (`span.tag`/`span.process.tag`) `map[string]interface{}`

### Unmarshalling
* jsoniter is the fastest for unmarshalling, lowest CPU and memory allocations. 1/3 CPU and about the same memory compared to stdlib
* gojay is 1/2 CPU but the memory usage is 2-3 times higher than stdlib

### Marshaling
* gojay performs the best for marshalling (`span.tag`/`span.process.tag`) `map[string]interface{}`. It does not increase allocations and CPU also looks stable
* jsoninter is slover than stdlib CPU and memory wise
* gojay might be a bit faster than stdlib but it consumes more memory.

[ci-img]: https://github.com/pavolloffay/golang-json-benchmark/workflows/Benchmark/badge.svg
[ci]: https://github.com/pavolloffay/golang-json-benchmark/actions
