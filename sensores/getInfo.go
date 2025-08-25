package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
	Os sensores receberão os dados do ambiente (um arquivo, nesse caso) em bytes
	e apenas passarão os processadores ainda em bytes
*/

/*
	Essa função que será chamada de forma concorrente varias vezes(para vários arquivos, representando
	o ambiente perceptível àquele sensor) passará os dados por meio de um buffer(canal) para a próxima
	linha de execução que é os processadores
*/

type Info struct {
	Name  []byte
	Min   []byte
	Max   []byte
	Sum   []byte
	Times []byte
}

func main() {
	n_sensores := flag.Int("n_sensores", 1, "Número de sensores")
	ambientes := flag.String("ambientes", "", "Lista de ambientes")
	n_buffer := flag.Int("n_buffer", 0, "Tamanho do buffer")
	flag.Parse()

	listaDeAmbientes := strings.Split(*ambientes, " ")
	if(len(ListaDeAmbientes) > *n_sensores){
		fmt.Println("Você passou mais ambientes que sensores\n Por favor, passa um numero menor ou igual ao numero de sensores.")
		return
	}

	buffer := make(chan Info, *n_buffer)

	for _, ambiente := range listaDeAmbientes {
		go getInfo(ambiente, buffer)
	}
}

func getInfo(path string, buffer chan<-){

}
