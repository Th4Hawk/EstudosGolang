# Variaveis em go

As variáveis são uma forma de armazenar dados temporariamente na memória, em golang, podemos fazer isso de duas formas. Na primeira forma, usando a palavra-chave "var", após isso informamos o nome da variável (que não pode ser o nome de uma palavra reservada) e por fim informamos seu tipo. Na segunda forma, informamos o nome da variável, seguido de ":=" e informamos seu valor. Na segunda forma, o go realiza de forma automática a tipagem.

```go
package main

import "fmt"

func main() {
    var nome string = "joao"
    idade := 18
    fmt.Println(nome, idade)
}
```