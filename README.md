# compare-go-cases

In Go, there are uncertain cases when one is not sure which code to choose. 

Let's compare performances of those cases.


## 1. slice vs map for 'contains'

Since map is O(1), it should be better to use map instead of iterating slice to check contains.. right? But there can be an overhead constructing map, so is it really a good idea to use map?
Let's compare.

BenchmarkContainsCheck/construct_map_:_size_10-8                10290870               116.0 ns/op
BenchmarkContainsCheck/map-8                                    84140656                12.54 ns/op
BenchmarkContainsCheck/slice-8                                  155193526                8.894 ns/op
BenchmarkContainsCheck/generic_slice_contains-8                 124357770                9.225 ns/op
BenchmarkContainsCheck/construct_map_:_size_10#01-8             10450783               115.2 ns/op
BenchmarkContainsCheck/map#01-8                                 95395822                12.31 ns/op
BenchmarkContainsCheck/slice#01-8                               156002023                7.690 ns/op
BenchmarkContainsCheck/generic_slice_contains#01-8              125440598                9.233 ns/op
BenchmarkContainsCheck/construct_map_:_size_100-8                 903336              1326 ns/op
BenchmarkContainsCheck/map#02-8                                 60912675                19.48 ns/op
BenchmarkContainsCheck/slice#02-8                               22311034                53.87 ns/op
BenchmarkContainsCheck/generic_slice_contains#02-8              21528444                55.51 ns/op
BenchmarkContainsCheck/construct_map_:_size_100#01-8              902517              1318 ns/op
BenchmarkContainsCheck/map#03-8                                 60944898                19.60 ns/op
BenchmarkContainsCheck/slice#03-8                               22099752                54.24 ns/op
BenchmarkContainsCheck/generic_slice_contains#03-8              21531390                55.36 ns/op
BenchmarkContainsCheck/construct_map_:_size_1000-8                113653             10595 ns/op
BenchmarkContainsCheck/map#04-8                                 97368351                12.31 ns/op
BenchmarkContainsCheck/slice#04-8                                2262998               530.1 ns/op
BenchmarkContainsCheck/generic_slice_contains#04-8               2260411               530.8 ns/op
BenchmarkContainsCheck/construct_map_:_size_1000#01-8             113384             10549 ns/op
BenchmarkContainsCheck/map#05-8                                 95551221                12.30 ns/op
BenchmarkContainsCheck/slice#05-8                                2264047               530.3 ns/op
BenchmarkContainsCheck/generic_slice_contains#05-8               2260662               530.2 ns/op

-> You'll get the advantage of map only if you're checking more than 20 times in case of 1000 elements size.