package main

import (
	"fmt"
	"monkey-study/repl"
	"os"
	"os/user"
)

const MONKEY_LOGO = `    __  ___            __
   /  |/  /___  ____  / /_____  __  __
  / /|_/ / __ \/ __ \/ //_/ _ \/ / / /
 / /  / / /_/ / / / / ,< /  __/ /_/ /
/_/  /_/\____/_/ /_/_/|_|\___/\__, /
                             /____/
`

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(MONKEY_LOGO)
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Name)
	fmt.Printf("Feel free to type in commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
