package main

import "flag"

func main() {
	force := flag.Bool("f", false ,"falseforce deletion, ignore nonexistent files, no prompt")
	interactive := flag.Bool("i", false, "prompt before every removal")
	flag.Parse()
}