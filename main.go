package main

import (
	"fmt"
	"net/http"

	"github.com/murilosolino/app-web/config/routes"
)

func main() {
	fmt.Println("Serviço Ativo na porta :8080")
	routes.CarregaRotas()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}
