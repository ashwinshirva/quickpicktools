package main

import (
	"github.com/ashwinshirva/quickpicktools/grpc-translator/services"
	_ "github.com/rookie-ninja/rk-grpc/boot"
)

func main() {
	services.RegisterServices()

	/* time.Sleep(5)

	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background()) */
}
