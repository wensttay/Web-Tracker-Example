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

const monitoriamentos = 3
const delay = 5

func main() {

	// registrarLog("saite-falseo", true)

	// exibeNomes()
	// nome, idade := devolveNomeEIdade()
	// outroNome, _ := devolveNomeEIdade()
	// fmt.Println("Nome:", nome+", Idade:", idade)
	// fmt.Println("OutroNome:", outroNome)

	exibirIntroducao()

	for {
		exibirMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Saindo ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando!")
			os.Exit(-1)
		}
	}

}

func devolveNomeEIdade() (string, int) {
	nome := "Wensttay"
	idade := 21
	return nome, idade
}

func exibirIntroducao() {
	// var nome = "Wensttay"
	// var idade = 21
	// var versao float32 = 1.1

	// fmt.Println("O tipo da váriavel nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da váriavel idade é", reflect.TypeOf(idade))
	// fmt.Println("O tipo da váriavel versão é", reflect.TypeOf(versao))

	// nome := "Wensttay"
	// idade := 21
	versao := 1.10

	// fmt.Println("Hello Sr.", nome, ",sua idade é:", idade)
	fmt.Println("###########################\n### Web Tracker Example ###\n###########################")
	fmt.Println()
	fmt.Println("Está programa está na versão:", versao)
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	// fmt.Scanf("%d", &comandoLido)
	fmt.Scan(&comandoLido)
	// fmt.Println("O endereço da variavel comandoLido é", &comandoLido)
	// fmt.Println("O camando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")

	sites := lerArquivoDeSites()

	for i := 0; i < monitoriamentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp != nil && resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com sucesso.")
		registrarLog(site, true)
	} else {
		fmt.Print("Site: ", site, " Está com problemas. ")
		registrarLog(site, false)
		if resp != nil {
			fmt.Print("Status code: ", resp.StatusCode)
		}
		fmt.Println()
	}
}

func exibirLogs() {
	fmt.Println("Exibindo Logs ...")
	printaLog()
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"} // SLICE
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Aparecida")
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}

func lerArquivoDeSites() []string {

	sites := []string{}
	arquivo, err := os.Open("sites.txt")
	leitor := bufio.NewReader(arquivo)

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler arquivo:", err)
	}

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registrarLog(site string, funcionando bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(funcionando) + "\n")

	arquivo.Close()
}

func printaLog() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
