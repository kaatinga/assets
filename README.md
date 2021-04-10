# The assets package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32()
String2Uint32() is a faster alternative to strconv.Atoi(), and optimized for uint32 integer type.
String2Uint32() function is faster than strconv.Atoi().

```
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkString2Uint32
BenchmarkString2Uint32-8   	100000000	        10.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8   	62046132	        18.94 ns/op	       0 B/op	       0 allocs/op
PASS
```