package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/kleriston/net-monitor-log-kleriston/internal"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia o monitoramento de conectividade",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Monitoramento iniciado...")
		internal.TestPing()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
