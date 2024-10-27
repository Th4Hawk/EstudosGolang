package main

import (
	"fmt"
	"os"
)

// Atividade: modificar o programa echo para exibir o os.Args[0]
// Codigo do echo: gopl.io/ch1/echo1
// Minha resolução: modificar o valor atribuido na variavel i para 0
func questao1() {
	var s, sep string

    for i := 0; i < len(os.Args); i++ {

		sep = " "

        s += sep + os.Args[i]
    }

    fmt.Println(s)
}

// Atividade: Modificar o programa echo para exibir o indice e o valor dos argumentos
// Codigo do echo: gopl.io/ch1/echo1
func questao2() {
	var s string

    for i := 1; i < len(os.Args); i++ {

        s += os.Args[i]

		fmt.Printf("Indice: %d Argumento: %s\n", i, s)
    }

}

// Medir a tempo de execução dos programas echo1, echo2 e echo3
// Codigo do echo1: gopl.io/ch1/echo1
// Codigo do echo2: gopl.io/ch1/echo2
// Codigo do echo3: gopl.io/ch1/echo3
func questao3() {
	
}

func main() {
	fmt.Println("Resolução da atividade 1.1")
	questao1()
	fmt.Println("Resolução da atividade 1.2")
	questao2()
}