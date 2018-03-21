package main

import (
	"superBrain/cmd"
	"os"
	"superBrain/engine"
)

func main(){
	go func() {
		engine.Run()
	}()

	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}







