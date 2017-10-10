package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	//"regexp"
	//"text/tabwriter"
)

func main() {
	//w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	var varg string
	for i := 1; i < len(os.Args); i++ {
		varg += " " + os.Args[i]
	}
	cmd := exec.Command("tbsql", varg)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "select * from dual;\n"+
			"select * from tab;\n"+
			"select * from dict where rownum < 10;\n"+
			"exit\n")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	//re := regexp.MustCompile("[.\t]+")
	fmt.Printf("%s\n", out)
	//w.Flush()
}
