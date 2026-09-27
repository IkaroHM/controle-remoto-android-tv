package Tv

import (
	"crypto/tls"
	"fmt"
	"os"
	"sync"

	"log"

	"github.com/drosocode/atvremote/pkg/common"
	"github.com/drosocode/atvremote/pkg/v2/command"
	"github.com/drosocode/atvremote/pkg/v2/pairing"
)

var (
    mu           sync.Mutex
    comandoTV    *command.Command
    pareamento   *pairing.Pairing
    ipDaTV       string
)

func SetIP(ip string) {
    ipDaTV = ip
}

func GetIP() string {
    return ipDaTV
}

func Conectar() error {
    mu.Lock()
    defer mu.Unlock()

    if comandoTV != nil {
        return nil
    }

    if _, err := os.Stat("certs/cert.pem"); os.IsNotExist(err) {
        return fmt.Errorf("certificados não encontrados")
    }

    certs, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/key.pem")
    if err != nil {
        return err
    }

    cmd := command.New(ipDaTV, 6466, &certs)
    if err := cmd.Connect(); err != nil {
        return fmt.Errorf("Erro ao criar conexao: %v", err)
    }

    comandoTV = &cmd
    return nil
}

func conectarSemLock() error {
    if comandoTV != nil {
        return fmt.Errorf("A conexao já existe")
    }

    certs, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/key.pem")
    if err != nil {
        return err
    }

    cmd := command.New(ipDaTV, 6466, &certs)
    if err := cmd.Connect(); err != nil {
        return err
    }

    comandoTV = &cmd
    return nil
}

func GetComando() *command.Command {
    mu.Lock()
    defer mu.Unlock()
    return comandoTV
}

func SetComando(cmd *command.Command) {
    mu.Lock()
    defer mu.Unlock()
    comandoTV = cmd
}

func IniciarPareamento() (*pairing.Pairing, error) {
    mu.Lock()
    defer mu.Unlock()

    certs, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/key.pem")
    if err != nil {
        return nil, err
    }

    p := pairing.New(ipDaTV, 6467, &certs)
    if err := p.Connect(); err != nil {
        return nil, err
    }

    pareamento = &p
    return &p, nil
}

func GetPareamento() *pairing.Pairing {
    mu.Lock()
    defer mu.Unlock()
    return pareamento
}

func FinalizarPareamento(codigo string) error {
    mu.Lock()
    defer mu.Unlock()

    if pareamento == nil {
        return fmt.Errorf("nenhum pareamento em andamento")
    }

    if len(codigo) != 6{
        return fmt.Errorf("código precisa ter 6 dígitos, recebi %d", len(codigo))
    }

    if err := pareamento.Secret(codigo); err != nil {
        return err
    }

    pareamento = nil

    comandoTV = nil
    return conectarSemLock()

}

func EnviarTecla(codigo common.RemoteKeyCode) error {
    mu.Lock()
    defer mu.Unlock()
    
    var err error
    for tentativa := 0; tentativa <= 3; tentativa++{

        if comandoTV == nil {
            if err = conectarSemLock(); err != nil {
                continue
            }
        }

        err = comandoTV.SendKey(codigo)
        if err == nil {
            return nil
        }

        log.Printf("Tentativa %d falhou: %v", tentativa+1, err)

        comandoTV = nil
        
    }

    return fmt.Errorf("Falhou após 3 tentativas: %w", err)
}