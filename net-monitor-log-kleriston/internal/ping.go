package internal

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func TestPing() {
	pinger, err := ping.NewPinger("google.com")
	if err != nil {
		fmt.Println("Erro ao criar pinger:", err)
		return
	}
	pinger.Count = 3
	pinger.Timeout = 5 * time.Second

	fmt.Println("Enviando ping para google.com...")
	err = pinger.Run()
	if err != nil {
		fmt.Println("Erro ao executar ping:", err)
		return
	}

	stats := pinger.Statistics()
	fmt.Printf("Perda de pacotes: %.2f%%\n", stats.PacketLoss)
}
