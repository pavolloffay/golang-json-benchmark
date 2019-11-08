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
GOMAXPROCS=1 go test -bench=. -test.benchtime=1s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkUnmarshalFastjson/default.json         	   48613	     20792 ns/op	   29440 B/op	     158 allocs/op
BenchmarkUnmarshalFastjson/default-unicode.json 	   84498	     13744 ns/op	   16160 B/op	     115 allocs/op
BenchmarkUnmarshalFastjson/default-tagmap.json  	  117885	      9750 ns/op	   14416 B/op	      54 allocs/op
BenchmarkMarshalGojay/default.json              	  177536	      6020 ns/op	    4240 B/op	       9 allocs/op
BenchmarkMarshalGojay/default-unicode.json      	  343594	      3491 ns/op	    1808 B/op	       8 allocs/op
BenchmarkMarshalGojay/default-tagmap.json       	  243284	      4237 ns/op	    2607 B/op	       8 allocs/op
BenchmarkUnmarshalGojay/default.json            	   50919	     21961 ns/op	    9959 B/op	      75 allocs/op
BenchmarkUnmarshalGojay/default-unicode.json    	   76257	     14708 ns/op	    6858 B/op	      49 allocs/op
BenchmarkUnmarshalGojay/default-tagmap.json     	   91306	     14085 ns/op	    5486 B/op	      61 allocs/op
BenchmarkUnmarshalStdlib/default.json           	   35632	     30156 ns/op	    1520 B/op	      73 allocs/op
BenchmarkUnmarshalStdlib/default-unicode.json   	   52899	     22596 ns/op	    1136 B/op	      53 allocs/op
BenchmarkUnmarshalStdlib/default-tagmap.json    	   53526	     22119 ns/op	    2192 B/op	      95 allocs/op
BenchmarkMarshalStdlib/default.json             	  201351	      5434 ns/op	    1536 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-unicode.json     	  260739	      4068 ns/op	    1024 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-tagmap.json      	  129321	      8683 ns/op	    2752 B/op	      37 allocs/op
BenchmarkUnmarshalJsoninter/default.json        	  144920	      7895 ns/op	    1184 B/op	      76 allocs/op
BenchmarkUnmarshalJsoninter/default-unicode.json         	  193281	      5618 ns/op	     816 B/op	      56 allocs/op
BenchmarkUnmarshalJsoninter/default-tagmap.json          	  181873	      6748 ns/op	    1456 B/op	      85 allocs/op
BenchmarkMarshalJsoniter/default.json                    	  169333	      6726 ns/op	    1800 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-unicode.json            	  216128	      4982 ns/op	    1288 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-tagmap.json             	   86020	     13352 ns/op	    7929 B/op	      50 allocs/op
```

```
GOMAXPROCS=1 go test -bench=. -test.benchtime=3s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkUnmarshalFastjson/default.json         	  168582	     22137 ns/op	   29440 B/op	     158 allocs/op
BenchmarkUnmarshalFastjson/default-unicode.json 	  239068	     14376 ns/op	   16160 B/op	     115 allocs/op
BenchmarkUnmarshalFastjson/default-tagmap.json  	  335364	     10715 ns/op	   14416 B/op	      54 allocs/op
BenchmarkMarshalGojay/default.json              	  513777	      6692 ns/op	    4240 B/op	       9 allocs/op
BenchmarkMarshalGojay/default-unicode.json      	  741099	      4054 ns/op	    1808 B/op	       8 allocs/op
BenchmarkMarshalGojay/default-tagmap.json       	  619767	      5136 ns/op	    2606 B/op	       8 allocs/op
BenchmarkUnmarshalGojay/default.json            	  118202	     30924 ns/op	   10170 B/op	      75 allocs/op
BenchmarkUnmarshalGojay/default-unicode.json    	  171378	     19002 ns/op	    6971 B/op	      49 allocs/op
BenchmarkUnmarshalGojay/default-tagmap.json     	  268003	     12953 ns/op	    5375 B/op	      61 allocs/op
BenchmarkUnmarshalStdlib/default.json           	  108180	     31704 ns/op	    1520 B/op	      73 allocs/op
BenchmarkUnmarshalStdlib/default-unicode.json   	  154046	     22705 ns/op	    1137 B/op	      53 allocs/op
BenchmarkUnmarshalStdlib/default-tagmap.json    	  157510	     22280 ns/op	    2192 B/op	      95 allocs/op
BenchmarkMarshalStdlib/default.json             	  583240	      5437 ns/op	    1536 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-unicode.json     	  803704	      4018 ns/op	    1024 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-tagmap.json      	  374754	      8523 ns/op	    2752 B/op	      37 allocs/op
BenchmarkUnmarshalJsoninter/default.json        	  430677	      7908 ns/op	    1184 B/op	      76 allocs/op
BenchmarkUnmarshalJsoninter/default-unicode.json         	  564072	      5638 ns/op	     816 B/op	      56 allocs/op
BenchmarkUnmarshalJsoninter/default-tagmap.json          	  504270	      6624 ns/op	    1456 B/op	      85 allocs/op
BenchmarkMarshalJsoniter/default.json                    	  511934	      6608 ns/op	    1800 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-unicode.json            	  651517	      4922 ns/op	    1288 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-tagmap.json             	  246168	     13414 ns/op	    7930 B/op	      50 allocs/op
```

```
GOMAXPROCS=4 go test -bench=. -test.benchtime=1s -benchmem -cpuprofile profile_cpu.out ./pkg/jsontest
goos: linux
goarch: amd64
pkg: json-benchmark/pkg/jsontest
BenchmarkUnmarshalFastjson/default.json-4         	   57165	     20604 ns/op	   29440 B/op	     158 allocs/op
BenchmarkUnmarshalFastjson/default-unicode.json-4 	   85374	     13539 ns/op	   16160 B/op	     115 allocs/op
BenchmarkUnmarshalFastjson/default-tagmap.json-4  	  123814	      9267 ns/op	   14416 B/op	      54 allocs/op
BenchmarkMarshalGojay/default.json-4              	  256896	      4498 ns/op	    4240 B/op	       9 allocs/op
BenchmarkMarshalGojay/default-unicode.json-4      	  364513	      3073 ns/op	    1808 B/op	       8 allocs/op
BenchmarkMarshalGojay/default-tagmap.json-4       	  317397	      3646 ns/op	    2607 B/op	       8 allocs/op
BenchmarkUnmarshalGojay/default.json-4            	   65073	     17697 ns/op	    9885 B/op	      75 allocs/op
BenchmarkUnmarshalGojay/default-unicode.json-4    	   98305	     11553 ns/op	    6777 B/op	      49 allocs/op
BenchmarkUnmarshalGojay/default-tagmap.json-4     	  100416	     11720 ns/op	    5424 B/op	      61 allocs/op
BenchmarkUnmarshalStdlib/default.json-4           	   41959	     28671 ns/op	    1520 B/op	      73 allocs/op
BenchmarkUnmarshalStdlib/default-unicode.json-4   	   56792	     20310 ns/op	    1136 B/op	      53 allocs/op
BenchmarkUnmarshalStdlib/default-tagmap.json-4    	   54284	     20627 ns/op	    2192 B/op	      95 allocs/op
BenchmarkMarshalStdlib/default.json-4             	  195253	      5766 ns/op	    1536 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-unicode.json-4     	  266449	      4248 ns/op	    1024 B/op	       1 allocs/op
BenchmarkMarshalStdlib/default-tagmap.json-4      	  139760	      8397 ns/op	    2752 B/op	      37 allocs/op
BenchmarkUnmarshalJsoninter/default.json-4        	  143780	      8010 ns/op	    1184 B/op	      76 allocs/op
BenchmarkUnmarshalJsoninter/default-unicode.json-4         	  194613	      5730 ns/op	     816 B/op	      56 allocs/op
BenchmarkUnmarshalJsoninter/default-tagmap.json-4          	  173272	      6613 ns/op	    1456 B/op	      85 allocs/op
BenchmarkMarshalJsoniter/default.json-4                    	  178272	      6423 ns/op	    1800 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-unicode.json-4            	  233517	      4799 ns/op	    1288 B/op	       6 allocs/op
BenchmarkMarshalJsoniter/default-tagmap.json-4             	  115671	     10225 ns/op	    7933 B/op	      50 allocs/op
```


## Conclusions

* stdlib marshalling performance drops when unmarshalling tag map (`span.tag`/`span.process.tag`) `map[string]interface{}`

[ci-img]: https://github.com/pavolloffay/golang-json-benchmark/workflows/Benchmark/badge.svg
[ci]: https://github.com/pavolloffay/golang-json-benchmark/actions
