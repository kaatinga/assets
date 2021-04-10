# The assets package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32() and Bytes2Uint32()
The functions are faster alternative to strconv.Atoi(). They are optimized for uint32 integer type.
String2Uint32() and Bytes2Uint32() are faster than strconv.Atoi().

```
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkString2Uint32
BenchmarkString2Uint32-8                        120055279           9.790 ns/op         0 B/op          0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8                        60553940            18.78 ns/op         0 B/op          0 allocs/op
BenchmarkBytes2Uint32
BenchmarkBytes2Uint32-8                         140027144           8.520 ns/op         0 B/op          0 allocs/op
BenchmarkStrvconv_AtoiPlusBytesConversion
BenchmarkStrvconv_AtoiPlusBytesConversion-8   	11169204            104.1 ns/op         54 B/op         3 allocs/op
PASS
```