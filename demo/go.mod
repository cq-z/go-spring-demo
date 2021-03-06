module github.com/cq-z/go-spring-demo/demo

go 1.16

require (
	github.com/cq-z/go-spring-demo/demo/admin v0.0.0
	github.com/go-spring/spring-boot v1.0.6-0.20201107041159-ef7b31e3a8fa
	github.com/go-spring/spring-web v1.0.6-0.20201109120017-baa11539c3ba
	github.com/go-spring/starter-gin v1.0.6-0.20201108112112-9d1b4a0067c3
	github.com/go-spring/starter-grpc v1.0.5
	google.golang.org/grpc v1.36.1
)

replace github.com/cq-z/go-spring-demo/demo/admin => ./admin
