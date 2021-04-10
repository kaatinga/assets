# The assets package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32(), Bytes2Uint32(), String2Byte(), etc
The functions are faster alternative to strconv.Atoi().

```
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkString2Byte
BenchmarkString2Byte-8                        	447266503	       2.612 ns/op	       0 B/op	       0 allocs/op
BenchmarkString2Uint32
BenchmarkString2Uint32-8                      	126690867	       9.335 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8                      	 64822033	       17.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytes2Uint32
BenchmarkBytes2Uint32-8                       	131627535	       9.239 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_AtoiPlusBytesConversion
BenchmarkStrvconv_AtoiPlusBytesConversion-8   	 11436552	       103.1 ns/op	      54 B/op	       3 allocs/op
PASS
```