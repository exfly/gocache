# gocache

go multi backend cache

# usage

go get -u github.com/exfly/gocache

# Bankmark

```
goos: darwin
goarch: amd64
pkg: github.com/ExFly/gocache/memory
Benchmark_Set-4                 200000000              386 ns/op
Benchmark_SetParallel-4         300000000              324 ns/op
Benchmark_Get-4                 500000000              152 ns/op
Benchmark_GetParallel-4         1000000000             106 ns/op
```

# cache link

- https://github.com/patrickmn/go-cache
- https://github.com/go-macaron/cache
- https://github.com/karlseguin/ccache
- https://github.com/muesli/cache2go
