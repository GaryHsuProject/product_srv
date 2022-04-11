package main

import (
	"shop/driver"
)

func main() {

	driver.Init()
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", driver.GlobalConfig.Port))
	// if err != nil {
	// 	panic(err)
	// }
	// grpcServer := grpc.NewServer()

}
