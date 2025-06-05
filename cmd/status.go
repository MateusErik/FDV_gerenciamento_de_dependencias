package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Mostra o status atual do monitoramento",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Status: monitoramento ativo (simulado)")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
