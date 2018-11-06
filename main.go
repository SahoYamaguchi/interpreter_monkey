// package main
//
// import(
//   "fmt"
//   "os"
//   "os/user"
//
//   "github.com/SahoYamaguchi/interpreter_monkey/repl"
// )
//
// func main() {
//   user,err := user.Current()
//   if err != nil{
//     panic(err)
//   }
//   fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
//   fmt.Printf("Feel free to type in commands\n")
//   repl.Start(os.Stdin, os.Stdout)
// }

package main

import (
    "fmt"

    "github.com/SahoYamaguchi/interpreter_monkey/lexer"
    "github.com/SahoYamaguchi/interpreter_monkey/parser"

)

func main() {
    input := "1 + 2 + 3;"
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()

    for _, s := range program.Statements {
        fmt.Printf("%+v\n", s)
    }
}
