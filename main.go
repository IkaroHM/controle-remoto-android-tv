package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/IkaroHM/controle-remoto-android-tv/Tv"
	"github.com/IkaroHM/controle-remoto-android-tv/routes"
	"github.com/joho/godotenv"
)

//go:embed public
var arquivosWeb embed.FS

func main() {

    if err := godotenv.Load(); err != nil {
        log.Println("AVISO: .env não encontrado, usando variáveis do ambiente")
    }

    ipDaTv := os.Getenv("IP_DA_TV")
    if ipDaTv == "" {
        log.Fatal("IP_DA_TV não definido")
    }
    Tv.SetIP(ipDaTv)

    if err := Tv.Conectar(); err != nil {
        log.Println("AVISO: não conectado à TV:", err)
        log.Println("Acesse /certificados/criar/ para criar os certificados")
    }

    pastaPublic, err := fs.Sub(arquivosWeb, "public")
    if err != nil {
        log.Fatal(err)
    }

    router := routes.CriarRouter(pastaPublic)
    router.StartRotas()

    fmt.Println("Server rodando na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}