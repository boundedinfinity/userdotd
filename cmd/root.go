package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var gContext context.Context
var cancelFunc context.CancelFunc
var osSignals chan os.Signal
var gLogger *log.Logger

var RootCmd = &cobra.Command{
	Use:   "userdotd",
	Short: "userdotd shell .d management utility",
	Long:  "userdotd utility",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		gLogger.Fatalf(err.Error())
	}
}

func handleError(err error) {
	log.Fatal(err.Error())
}

func init() {
	gLogger = log.New(os.Stdout, "", 0)
	osSignals = make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	gContext, cancelFunc = context.WithCancel(context.Background())

	go func() {
		s := <-osSignals
		gLogger.Printf("Received signal: %v", s)
		cancelFunc()
	}()

	cobra.OnInitialize(InitConfigEnvironment)

	ConfigureDebug(RootCmd.PersistentFlags())
	ConfigureFormat(RootCmd.PersistentFlags())
}
