# Performance Benchmark graphql-golang
Benchmark of two main libraries of graphql for golang, the glqgen and graphql-go.

## Benchmark functions
The functions used for building this benchmark are a query that returns "Hello world", a mutation that creates a Todo object and a query that list all Todo objects created.

After running a functions, it shows how many times that loop ran at witch speed, like:
```
BenchmarkRandInt-8   	68453040	   17.8 ns/op
```

## Benchmark results
After cloning this repo, to run the benchmarks you can run the following commands:

    // Run all the benchmarks
    go test -benchmem -run=^$ -bench Benchmark
    // Run only gqlgen benchmarks
    go test -benchmem -run=^$ -bench BenchmarkGqlgen
    // Run only graphqlgo benchmarks
    go test -benchmem -run=^$ -bench BenchmarkGraphqlGO

Possible results:
+ Gqlgen
```
BenchmarkGqlgenHello-8          1000000000               0.000074 ns/op        0 B/op          0 allocs/op
BenchmarkGqlgenCreateTodo-8        19322             64884 ns/op           27426 B/op        448 allocs/op
BenchmarkGqlgenListTodos-8         36224             37946 ns/op           19411 B/op        288 allocs/op
```
+ Graphql-go
```
BenchmarkGraphqlGOHello-8                  19789             61846 ns/op           31657 B/op        543 allocs/op
BenchmarkGraphqlGOCreateTodo-8              5478            205291 ns/op           75397 B/op       1373 allocs/op
BenchmarkGraphqlGOListTodos-8                 18          60148059 ns/op        34386488 B/op     542423 allocs/op
```

## Docker Benchmark
There are created images for docker to test the services with limited resources in docker swarm.

The [docker-compose.yml](https://github.com/ygorrodriguesdft/gqlgen-benchmark/blob/master/docker-compose.yml) file can be used as an example for deploying the test.

After the deploy I used the following docker stats in each virtual machine and docker service logs to obtain the results:

```
// docker stats --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.NetIO}}\t{{.BlockIO}}"
```
| NAME | CPU VARIATION % | MEM USAGE VARIATION | NET I/O AVERAGE | BLOCK I/O AVERAGE |
| ---- | ----- | ----------------- | ------- | --------- |
| graphql-go-service | 95% - 110% | 16.00MiB - 75.00MiB | 516B / 0B | 60.4MB / 471kB |
| gqlgen-service | 95% - 110% | 20.00MiB - 80.00MiB | 516B / 0B | 60.4MB / 471kB |

```
docker service logs service-id
BenchmarkGqlgenHello      	1000000000	         0.000087 ns/op	       0 B/op	       0 allocs/op
BenchmarkGqlgenCreateTodo 	   15496	     67402 ns/op	   27276 B/op	     443 allocs/op
BenchmarkGqlgenListTodos  	   31657	     39394 ns/op	   19314 B/op	     284 allocs/op
BenchmarkGraphqlGOHello      	   16490	     67632 ns/op	   31346 B/op     529 allocs/op
BenchmarkGraphqlGOCreateTodo 	    5788	    241571 ns/op	   74781 B/op    1301 allocs/op
BenchmarkGraphqlGOListTodos  	      10	 108447902 ns/op	36283162 B/op  572132 allocs/op
```