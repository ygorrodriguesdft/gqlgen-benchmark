FROM golang
COPY . /go/src/gqlgen-benchmark
WORKDIR /go/src/gqlgen-benchmark
RUN go mod download
CMD go test -benchmem -run=^$ -bench BenchmarkGraphqlGO