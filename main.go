import "github.com/seu-usuario/net-monitor-log/internal"

func main() {
	internal.InitLogger()
	internal.Log.Info("Monitoramento iniciado")
}
