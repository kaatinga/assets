# The assets package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32(), Bytes2Uint32(), String2Byte(), etc
The functions are faster alternative to strconv.Atoi().

```
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkString2Byte
BenchmarkString2Byte-8                        	775535930	  1.519 ns/op	   0 B/op	   0 allocs/op
BenchmarkString2Uint32
BenchmarkString2Uint32-8                      	177936165	  6.592 ns/op	   0 B/op	   0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8                      	64663167	  18.33 ns/op	   0 B/op	   0 allocs/op
BenchmarkBytes2Uint32
BenchmarkBytes2Uint32-8                       	202709667	  5.853 ns/op	   0 B/op	   0 allocs/op
BenchmarkStrvconv_AtoiPlusBytesConversion
BenchmarkStrvconv_AtoiPlusBytesConversion-8   	10781844	  104.4 ns/op	   54 B/op	   3 allocs/op
PASS
```