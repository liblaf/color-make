package logging

import "github.com/pterm/pterm"

func Fatalln(err ...any) {
	pterm.Fatal.Println(err...)
}
