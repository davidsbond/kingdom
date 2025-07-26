package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	"github.com/davidsbond/kingdom/internal/server"
)

var (
	//go:embed usage.txt
	usage string
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	var (
		host    string
		port    int
		keyPath string
		debug   bool
	)

	cmd := &cobra.Command{
		Use:               "kingdom",
		Short:             "Host a game of Kingdom",
		Long:              usage,
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		RunE: func(cmd *cobra.Command, args []string) error {
			level := log.InfoLevel
			if debug {
				level = log.DebugLevel
			}

			logger := log.NewWithOptions(os.Stderr, log.Options{
				Level:           level,
				ReportTimestamp: true,
				ReportCaller:    debug,
				Formatter:       log.TextFormatter,
			})

			return server.Run(cmd.Context(), server.Config{
				Host:    host,
				Port:    port,
				KeyPath: keyPath,
				Logger:  logger,
			})
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&host, "host", "H", "0.0.0.0", "host for serving SSH connections")
	flags.IntVarP(&port, "port", "p", 22, "port for serving SSH connections")
	flags.StringVarP(&keyPath, "key-path", "k", ".ssh/id_ed25519", "path to SSH keys")
	flags.BoolVarP(&debug, "debug", "d", false, "enable debug logging")

	if err := fang.Execute(ctx, cmd); err != nil {
		os.Exit(1)
	}
}
