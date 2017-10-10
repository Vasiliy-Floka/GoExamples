package main

import (
	//"bufio"
	"fmt"
	"log"
	//"math/rand"
	"os"
	"os/exec"
	//"strings"
	//	"time"
)

func main() {
	var v_arg string
	for i := 1; i < len(os.Args); i++ {
		v_arg += " " + os.Args[i]
	}
	var cmdChain = []*exec.Cmd{
		exec.Command("tbsql", v_arg)}
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout
	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Fatalln(err)
		} else {
			defer cmd.Process.Kill()
		}
	}
	fmt.Println("select * from dual;");
	fmt.Println("exit");
	for _, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Fatalln(err)
		}
	}
	
}
