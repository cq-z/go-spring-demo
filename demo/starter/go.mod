module github.com/cq-z/go-spring-demo/demo/starter

go 1.16

require (
	github.com/cq-z/go-spring-demo/demo/admin v0.0.0
	github.com/go-spring/spring-boot v1.0.6-0.20201107041159-ef7b31e3a8fa
	github.com/spf13/cobra v1.1.3
)

replace github.com/cq-z/go-spring-demo/demo/admin => ../admin
