package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"tg-bot/processor"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./config/.env")
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err := processor.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	defer trace.Stop()

	p.StartTG()

}
