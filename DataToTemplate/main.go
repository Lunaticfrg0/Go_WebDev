package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type ObjectToPass struct {
	WhereToStream string
	SongsList     []string
}

func init() {
	tpl = template.Must(template.ParseFiles("tmplt.gohtml"))
}
func main() {
	fmt.Println("stream Radical Optimism by Dua Lipa")
	oTP := ObjectToPass{
		"any platform",
		[]string{"End if an era", "Houdini", "Ilusion", "Falling Forever", "Maria", "These Walls", "Training season"},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tmplt.gohtml", oTP)
	if err != nil {
		log.Println(err)
	}
}
