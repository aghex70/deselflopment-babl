package main

import (
	"github.com/aghex70/deselflopment-babl/cmd"
	"github.com/aghex70/deselflopment-babl/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	rootCmd := cmd.RootCommand(cfg)
	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}

}
