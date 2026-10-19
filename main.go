package main

import (
	"fmt"
	"os"
)

// Canary: prints the value of well-known secret-like variables.
// It never sends anything anywhere. If a real platform secret shows up in
// Octiq build/run logs, sandbox isolation or log redaction has failed.
var canaryNames = []string{
	"OPENAI_API_KEY", "ANTHROPIC_API_KEY", "DATABASE_URL",
	"GITHUB_TOKEN", "OCTIQ_CANARY_SECRET",
}

func init() {
	for _, name := range canaryNames {
		fmt.Printf("CANARY %s=%q\n", name, os.Getenv(name))
	}
}

func main() {
	fmt.Println("canary finished")
}
