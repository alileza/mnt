package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	t := time.Now()
	cmd.Run()
	processTime := time.Since(t)

	f, err := os.OpenFile("mnt.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}
	hostname, _ := os.Hostname()

	fmt.Fprintf(f, "%s;%s;%s;%s!\n", t.Format(time.RFC3339), hostname, processTime.String(), strings.Join(os.Args[1:], " "))
	f.Close()
}
