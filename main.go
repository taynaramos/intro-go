package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibirIntroducao()

	for {
		exibirMenu()
		comando := lerComando()

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa...")
		// } else {
		// 	fmt.Println("Comando inválido")
		// }

		switch comando {
		case 1:
			iniciarMonitoranmento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) // o 0 indica que o usuario saiu do programa e ocorreu como o esperado
		default:
			fmt.Println("Comando inválido")
			os.Exit(-1) // o -1 indica que algo não saiu como esperado
		}
	}
}

func exibirIntroducao() {
	nome := "Tayná"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func lerComando() int {
	var comandoLido int
	// pegar o valor digitado pelo usuário e armazenar na variável comando (endereço de memória da variável)
	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoranmento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://random-status-code.herokuapp.com/",
	// 	"https://www.alura.com.br", "https://www.caelum.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for _, site := range sites {
			testaSite(site)
		}
		// dar um delay de 5 segundos
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}
	fmt.Println("")
}

func testaSite(site string) *http.Response {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resp.StatusCode)
		registraLog(site, false)
	}
	return resp
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		// removendo espaços
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		// io.EOF equivale a end of file, pois o arquivo acabou
		if err == io.EOF {
			break
		}
	}
	// fmt.Println(string(arquivo))
	fmt.Println(sites)

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/06/2006 15:04:05") + " " + site + " - online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))

}
