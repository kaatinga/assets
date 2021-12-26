[![GitHub release](https://img.shields.io/github/release/kaatinga/assets.svg)](https://github.com/kaatinga/assets/releases)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/assets/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/assets/branch/master/graph/badge.svg)](https://codecov.io/gh/kaatinga/assets)
[![lint workflow](https://github.com/kaatinga/assets/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/assets/actions?query=workflow%3Alinter)
[![help wanted](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/assets/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# The assets package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32(), Bytes2Uint32(), String2Byte(), etc

The functions are faster alternative to `strconv.Atoi()`.

```
cpu: Intel(R) Core(TM) i7-4870HQ CPU @ 2.50GHz
BenchmarkString2Byte
BenchmarkString2Byte-8     	359347491	         3.097 ns/op	       0 B/op	       0 allocs/op
BenchmarkString2Uint32
BenchmarkString2Uint32-8   	92657037	        11.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconv_Atoi
BenchmarkStrvconv_Atoi-8   	51235477	        22.03 ns/op	       0 B/op	       0 allocs/op
PASS
```

## Uint162String, Byte2String, etc

The functions are faster alternative to `strconv.Itoa()`.

```
cpu: AMD Ryzen 5 3400G with Radeon Vega Graphics    
BenchmarkUint162String
BenchmarkUint162String-8   	68361266	        17.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkStrvconvItoa
BenchmarkStrvconvItoa-8    	30279402	        39.60 ns/op	       5 B/op	       1 allocs/op
```

Warning. For the sake of maximum processing speed, all the converters have limited support of the leading zeros.
The length of the string cannot exceed the maximum length of the string of every certain type. For example, `012` is
supported fot the `byte` type, whereas `0255` is not as it consist of 4 characters what is impossible for `byte` type.
In case you process a date number, like month that is formatted with zero, like `09`, you still can use the fast
converters of the assets package, but if you still need to process formatted numbers like `0000000000000255` as byte, use
`DeleteLeadingZeros` helper to remove leading zeros in the string first:

```go
numberINeed, err = assets.String2Byte(DeleteLeadingZeros("0000000000000255"))
if err != nil {
	...
}
```