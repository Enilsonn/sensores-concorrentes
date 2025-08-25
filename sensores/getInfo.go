package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
	O arquivo passará os dados separados por ponto-e-virgula e por linha
	Brazil;TempMin;TempMax
*/

/*
	Os sensores receberão os dados do ambiente (um arquivo, nesse caso) em bytes
	e apenas passarão os processadores ainda em bytes
*/

/*
	Essa função que será chamada de forma concorrente varias vezes(para vários arquivos, representando
	o ambiente perceptível àquele sensor) passará os dados por meio de um buffer(canal) para a próxima
	linha de execução que é os processadores
*/

/*
	Como não estamos salvando os dados em memoria ainda, para passá-los todos de uma vez, não usaremos mutex
*/

/*
	n_buffer aqui funciona como um semáforo, se o numero de threads(gorotines) em execução por igual a ele,
	as novas threads param de executar até que uma conceda recurso(tempo de CPU), ou seja as novas
	threads são barradas de executar
*/

type Info struct {
	Name []byte
	Min  []byte
	Max  []byte
}

func main() {
	/*	flags
		--n_sensores=n
		--ambientes=path1 path2 ... pathn
		--n_buffer=n
	*/
	n_sensores := flag.Int("n_sensores", 1, "Número de sensores")
	ambientes := flag.String("ambientes", "", "Lista de ambientes")
	n_buffer := flag.Int("n_buffer", 0, "Tamanho do buffer")
	flag.Parse()

	listaDeAmbientes := strings.Split(*ambientes, " ")
	if len(listaDeAmbientes) > *n_sensores {
		fmt.Println("você passou mais ambientes que sensores\n Por favor, passa um numero menor ou igual ao numero de sensores.")
		return
	}

	buffer := make(chan Info, *n_buffer)

	for _, ambiente := range listaDeAmbientes {
		go getInfo(ambiente, buffer)
	}
}

func getInfo(path string, buffer chan<- Info) {
	// abrindo o arquivo
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("erro ao acessar o arquivo")
		return
	}
	defer file.Close()

	// scaner para ler o arquivo
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linha := scanner.Text()
		dado := strings.Split(linha, ";")
		if len(dado) < 3 || len(dado) > 3 {
			continue // linha mal formatada
		}
		dadoEstruturado := Info{
			Name: []byte(dado[0]),
			Min:  []byte(dado[1]),
			Max:  []byte(dado[2]),
		}

		buffer <- dadoEstruturado
	}
}
