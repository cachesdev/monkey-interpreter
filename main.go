package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/cachesdev/monkey-interpreter/repl"
)

/* Codigo > Tokens > AST

Convertimos el codigo fuente en un set de tokens que luego se pasan a la funcion AST y crea
un arbol sintactico que mantiene el alcance contextual de cada token.

let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
*/

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Bienvenido %s a la REPL del lenguaje de programacion Monkey\n", user.Username)
	fmt.Printf("Puede proceder a escribir comandos\n")
	repl.Start(os.Stdin, os.Stdout)
}
