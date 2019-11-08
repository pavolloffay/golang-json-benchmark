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
4 CPUs
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

* stdlib marshalling performance drops (especially allocations increase) when unmarshalling tag map (`span.tag`/`span.process.tag`) `map[string]interface{}`
* GoJay performs the best for marshalling (`span.tag`/`span.process.tag`) `map[string]interface{}`. It does not increase allocations
* Fastjson does not provide a way to unmarshal `interface{}` in tag map

[ci-img]: https://github.com/pavolloffay/golang-json-benchmark/workflows/Benchmark/badge.svg
[ci]: https://github.com/pavolloffay/golang-json-benchmark/actions
