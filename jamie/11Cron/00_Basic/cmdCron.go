package main

import (
	"github.com/go-co-op/gocron"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

func job() {
	cmd := exec.Command("notepad")
	err := cmd.Run()
	log.Println("5분마다")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minutes().Do(job)
	s.StartAsync()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
