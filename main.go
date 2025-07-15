package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	force := flag.Bool("f", false, "falseforce deletion, ignore nonexistent files, no prompt")
	interactive := flag.Bool("i", false, "prompt before every removal")
	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "rm: missing operand")
		os.Exit(1)
		flag.Usage()
		return
	}
}

func confirm(file string) bool {
	fmt.Printf("rm: remove '%s'? [y/N]: ", file)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	return answer == "y" || answer == "yes"
}
