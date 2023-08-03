package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/elct9620/servant"
	"github.com/spf13/cobra"
)

const ServePort = 8080

func main() {
	root := &cobra.Command{
		Use:   "servantd",
		Short: "Start servant daemon",
		RunE:  runDaemon,
	}

	ping := &cobra.Command{
		Use:   "healthz",
		Short: "Check servant daemon health",
		RunE:  runHealthz,
	}

	root.AddCommand(ping)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
	}
}

func runDaemon(cmd *cobra.Command, args []string) error {
	api := servant.NewApi()
	return http.ListenAndServe(fmt.Sprintf(":%d", ServePort), api)
}

func runHealthz(cmd *cobra.Command, args []string) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/livez", ServePort))
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}

	return nil
}
