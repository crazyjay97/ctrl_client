package core

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func WiSunLoader(path string, eraseAll bool, blMode bool, process chan int, port string) {
	var args = []string{
		path,
		"-p",
		port,
	}
	if eraseAll {
		args = append(args, "-e")
	}
	if blMode {
		args = append(args, "-s")
	}
	cmd := exec.Command("wisun-loader", args...)
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(stdout)
	for {
		line, _, err := reader.ReadLine()
		output := string(line)
		log.Println(output)
		substr := "progress value = "
		idx := strings.LastIndex(output, substr)
		if strings.Index(output, "ERROR") != -1 {
			cmd.Process.Kill()
			process <- -1
			return
		}
		if idx != -1 {
			v, err := strconv.Atoi(output[idx+len(substr):])
			if err == nil {
				process <- v
			}
		}
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
