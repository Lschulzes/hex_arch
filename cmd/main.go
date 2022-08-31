package main

import (
	"github.com/lschulzes/hex_arch/internal/adapters/app/api"
	"github.com/lschulzes/hex_arch/internal/adapters/core/arithmetic"
	"github.com/lschulzes/hex_arch/internal/adapters/framework/left/grpc"
	"github.com/lschulzes/hex_arch/internal/adapters/framework/right/database"
	"github.com/lschulzes/hex_arch/internal/ports"
	"os"
)

func main() {
	var core ports.ArithmeticPort
	var dbAdapter ports.DBPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	driver, dataSrcName := os.Getenv("DATABASE_DRIVER"), os.Getenv("DATA_SRC_NAME")
	dbAdapter = database.NewAdapter(driver, dataSrcName)
	defer dbAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbAdapter)
	gRPCAdapter = grpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
