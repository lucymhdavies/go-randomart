package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/lucymhdavies/go-randomart"
)

func main() {

	data := []byte(strings.Join(os.Args[1:], " "))

	sshhash := randomart.NewSSH()

	sshhash.Write(data)

	s := sshhash.String()
	fmt.Println(s)

}
