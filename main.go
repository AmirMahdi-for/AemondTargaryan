package main

import (
	"github.com/amirmahdi-for/AemondTargaryen/cmd"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const description = "AemondTargaryan"
	root := &cobra.Command{Short: "description"}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Server{}.Command(trap),
	)

	err := root.Execute()
	if err != nil {
		log.Fatalf("failed to execute root command\n%v", err)
	}

}
