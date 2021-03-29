/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "github.com/cq-z/go-spring-demo/demo/admin/pkg/helloworld"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
	_ "github.com/go-spring/starter-gin"
	_ "github.com/go-spring/starter-grpc/client"
)

const (
	defaultName = "world"
)

func init() {
	SpringBoot.RegisterBean(new(GreeterClientController)).Init(func(c *GreeterClientController) {
		SpringBoot.GetMapping("/", c.index)
	})
}

type GreeterClientController struct {
	GreeterClient pb.GreeterClient `autowire:""`
}

func (c *GreeterClientController) index(webCtx SpringWeb.WebContext) {
	r, err := c.GreeterClient.SayHello(webCtx.Request().Context(), &pb.HelloRequest{Name: defaultName})
	SpringWeb.ERROR.Panic(err).When(err != nil)
	webCtx.String("Greeting: " + r.GetMessage())
}

func init() {
	SpringBoot.RegisterGRpcClient(pb.NewGreeterClient, "greeter-client")
}

func main() {
// Set up a connection to the server.
conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithBlock())
if err != nil {
	log.Fatalf("did not connect: %v", err)
}
defer conn.Close()
c := pb.NewGreeterClient(conn)

// Contact the server and print out its response.
name := defaultName
if len(os.Args) > 1 {
	name = os.Args[1]
}
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
if err != nil {
	log.Fatalf("could not greet: %v", err)
}
log.Printf("Greeting: %s", r.GetMessage())



// 	r, err := c.GreeterClient.SayHello(webCtx.Request().Context(), &pb.HelloRequest{Name: defaultName})
// 	if err != nil {
//         log.Fatalf("could not greet: %v", err)
// }
// fmt.Printf("Greeting: %s", r.GetMessage())
	// SpringBoot.SetProperty("grpc.endpoint.greeter-client.address", "127.0.0.1:50051")
	// SpringBoot.SetProperty("spring.application.name", "GreeterClient")
	// SpringBoot.RunApplication()
}
