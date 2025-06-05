package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ping/ping"
)

func TestConnection(logPath string) {
	pinger, err := ping.NewPinger("8.8.8.8") // Google DNS
	if err != nil {
		log.Fatalf("Erro ao criar pinger: %v", err)
	}

	pinger.Count = 3
	pinger.Timeout = 5 * time.Second

	err = pinger.Run()
	if err != nil {
		writeLog(logPath, false)
		fmt.Println("Conexão: ❌ (Erro ao executar ping)")
	} else {
		stats := pinger.Statistics()
		if stats.PacketsRecv == 0 {
			writeLog(logPath, false)
			fmt.Println("Conexão: ❌ (Sem pacotes recebidos)")
		} else {
			writeLog(logPath, true)
			fmt.Printf("Conexão: ✅ Latência média: %v\n", stats.AvgRtt)
		}
	}
}

func writeLog(path string, online bool) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Erro ao abrir o arquivo de log: %v", err)
		return
	}
	defer f.Close()

	status := "OFFLINE"
	if online {
		status = "ONLINE"
	}

	entry := fmt.Sprintf("[%s] Conexão: %s\n", time.Now().Format("2006-01-02 15:04:05"), status)
	if _, err := f.WriteString(entry); err != nil {
		log.Printf("Erro ao escrever no arquivo de log: %v", err)
	}
}
