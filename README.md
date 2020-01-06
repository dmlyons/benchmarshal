# Benchmarshal

A little test package to benchmark a few serialization methods for arbitrary data (no schema provided)

```
19:59 $ go test -bench=. 
goos: linux
goarch: amd64
pkg: github.com/dmlyons/benchmarshal
BenchmarkAll/JSON_1/marshal-8 	  936465	      1258 ns/op
BenchmarkAll/JSON_1/unmarshal-8         	  587817	      2103 ns/op
BenchmarkAll/Gob_1/encode-8             	  216973	      5991 ns/op
BenchmarkAll/Gob_1/decode-8             	   55053	     22527 ns/op
BenchmarkAll/Msgpack_1/marshal-8        	 1723753	       632 ns/op
BenchmarkAll/Msgpack_1/unmarshal-8      	 1313677	       939 ns/op
BenchmarkAll/JSON_10/marshal-8          	   99745	     11798 ns/op
BenchmarkAll/JSON_10/unmarshal-8        	   59775	     20855 ns/op
BenchmarkAll/Gob_10/encode-8            	   88365	     13617 ns/op
BenchmarkAll/Gob_10/decode-8            	   39987	     31302 ns/op
BenchmarkAll/Msgpack_10/marshal-8       	  311244	      4193 ns/op
BenchmarkAll/Msgpack_10/unmarshal-8     	  173952	      6616 ns/op
BenchmarkAll/JSON_100/marshal-8         	   10000	    118162 ns/op
BenchmarkAll/JSON_100/unmarshal-8       	    6302	    210355 ns/op
BenchmarkAll/Gob_100/encode-8           	   14257	     87443 ns/op
BenchmarkAll/Gob_100/decode-8           	   10000	    109395 ns/op
BenchmarkAll/Msgpack_100/marshal-8      	   27106	     42498 ns/op
BenchmarkAll/Msgpack_100/unmarshal-8    	   17080	     67384 ns/op
BenchmarkAll/JSON_1000/marshal-8        	    1026	   1213590 ns/op
BenchmarkAll/JSON_1000/unmarshal-8      	     556	   2134011 ns/op
BenchmarkAll/Gob_1000/encode-8          	    1524	    804020 ns/op
BenchmarkAll/Gob_1000/decode-8          	    1458	    864205 ns/op
BenchmarkAll/Msgpack_1000/marshal-8     	    3022	    437851 ns/op
BenchmarkAll/Msgpack_1000/unmarshal-8   	    1932	    693652 ns/op
BenchmarkAll/JSON_10000/marshal-8       	      92	  12584321 ns/op
BenchmarkAll/JSON_10000/unmarshal-8     	      46	  21767520 ns/op
BenchmarkAll/Gob_10000/encode-8         	     132	   9636253 ns/op
BenchmarkAll/Gob_10000/decode-8         	     134	   8902010 ns/op
BenchmarkAll/Msgpack_10000/marshal-8    	     286	   4292337 ns/op
BenchmarkAll/Msgpack_10000/unmarshal-8  	     170	   6655325 ns/op
BenchmarkAll/JSON_100000/marshal-8      	       8	 138721057 ns/op
BenchmarkAll/JSON_100000/unmarshal-8    	       5	 227712145 ns/op
BenchmarkAll/Gob_100000/encode-8        	      13	  94916227 ns/op
BenchmarkAll/Gob_100000/decode-8        	      12	  91523425 ns/op
BenchmarkAll/Msgpack_100000/marshal-8   	      24	  44834214 ns/op
BenchmarkAll/Msgpack_100000/unmarshal-8 	      16	  72257349 ns/op
PASS
ok  	github.com/dmlyons/benchmarshal	57.789s
```
