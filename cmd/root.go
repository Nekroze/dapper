package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var backend string

var rootCmd = &cobra.Command{
	Use:   "dapper",
	Short: "Dockerized application manager",
	Long:  `Dockerized application manager.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&backend, "backend", "docker", "container backend to use")

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
