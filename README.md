[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kaatinga/assets/blob/main/LICENSE)
[![codecov](https://codecov.io/gh/kaatinga/assets/branch/main/graph/badge.svg)](https://codecov.io/gh/kaatinga/assets)
[![lint workflow](https://github.com/kaatinga/assets/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/kaatinga/assets/actions?query=workflow%3Alinter)
[![Generic badge](https://img.shields.io/badge/Help%20wanted-True-yellow.svg)](https://github.com/kaatinga/assets/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22)

# The assets package

The package contain several useful functions intended for conversion between string,
[]byte and different integer types.

## String2Uint32(), Bytes2Uint32(), String2Byte(), etc
The functions are faster alternative to strconv.Atoi().

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