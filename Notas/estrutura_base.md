# Estrutura de um programa go

Começamos a escrever um programa go informando o nome do pacote, é importante ressaltar a obrigatoriedade da existência de um pacote "main" em qualquer programa go. Com o nome do pacote informado, seguimos para as importações e, após isso, para as funções (em go, tudo é uma função).

```go
// Informa que este pacote é o principal
package main
// Biblioteca de I/O
import "fmt"
// Função principal do programa
func main() {
    fmt.Println("Olá mundo.")
}
```