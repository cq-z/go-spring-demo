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

// Package main implements a server for Greeter service.
package cmd

import (
	"context"
	"fmt"
	pb "github.com/cq-z/go-spring-demo/demo/admin/pkg/helloworld"
	"github.com/go-spring/spring-boot"
	_ "github.com/go-spring/starter-grpc/server"
	"github.com/spf13/cobra"
	"log"
)

var(
  validArgs   = []string{"start"}

  grpcServer = &cobra.Command{
	Use:   "grpc [start]",
	Short: "staring grpc service",
	Long: `staring grpc service`,
	ValidArgs: validArgs,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Accepts 1 arg(s), received %d\n"+
				"Valid arg(s): %s\n", len(args), validArgs)
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
			case "start":
				configLocations := []string{
					"config/",
				}
				SpringBoot.SetProperty("grpc.server.port", 50051)
				SpringBoot.RunApplication(configLocations...)
			default:
				fmt.Errorf("Valid arg(s): %s\n", validArgs)
		}
	},
  }
)
func init() {
	SpringBoot.SetProperty("cmd.grpcServer", grpcServer)

	SpringBoot.RegisterGRpcServer(pb.RegisterGreeterServer,
		"helloworld.Greeter",
		new(GreeterServer),
	).ConditionOnOptionalPropertyValue("greeter-server.enable", true)
}

type GreeterServer struct {
	AppName string `value:"${spring.application.name:=test}"`
}

func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName() + " from " + s.AppName}, nil
}

