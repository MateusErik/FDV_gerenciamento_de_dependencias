package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netmonitor",
	Short: "Monitorador de conectividade com a internet",
	Long:  "Uma ferramenta CLI para monitorar conectividade de rede e registrar falhas em log.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
