package cmd

import (
	"log"
	"net"
	"strconv"

	"github.com/amine7536/echo-server/conf"
	"github.com/amine7536/echo-server/server"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "echo-server",
	Run: run,
}

var version string
var progName string

// NewRootCmd will setup and return the root command
func NewRootCmd(v string, p string) *cobra.Command {
	// Set Version and ProgramName
	version = v
	progName = p

	rootCmd.PersistentFlags().StringP("config", "c", "", "Config file to use")
	rootCmd.PersistentFlags().StringP("listen", "l", "localhost", "Listen interface")
	rootCmd.Flags().IntP("port", "p", 3000, "Port to bind")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	logger, err := conf.ConfigureLogging(&config.LogConfig)
	if err != nil {
		log.Fatal("Failed to configure logging: " + err.Error())
	}

	logger.Infof("Starting with config: %+v", config)

	// Service string
	service := string(config.Host) + ":" + strconv.FormatInt(config.Port, 10)

	echoFunc := func(req server.EchoRequest) {
		// Read request data
		readError := req.Read()
		if readError != nil {
			logger.Errorf("Error reading from socket, %s", readError.Error())
		}

		// Log request body if log level is debug
		logger.Debugf("req.Body=%s", req.Body)

		// Reply with the same data
		writeError := req.Write()
		if writeError != nil {
			logger.Errorf("Error reading from socket, %s", readError.Error())
		}
	}

	conn, err := net.Listen("tcp", service)
	if err != nil {
		logger.Fatalf("Error %s", err.Error())
	}

	s := server.New(conn, logger)
	s.ListenAndEcho(echoFunc)

}
