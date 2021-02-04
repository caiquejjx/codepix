/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/caiquejjx/codepix/codepix-go/application/grpc"
	"github.com/caiquejjx/codepix/codepix-go/infraestructure/db"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var port int
var database *gorm.DB

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC Server",

	Run: func(cmd *cobra.Command, args []string) {
		database = db.ConnectDB(os.Getenv("env"))

		grpc.StartGrpcServer(database, port)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&port, "port", "p", 50051, "grpc Server port")

}
