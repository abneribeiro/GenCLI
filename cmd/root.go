/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "Gen",
	Short: "Generate your projects from the terminal",
	Long: `terminalProjects is a CLI application that allows you to generate your projects from the terminal.
	
	You can create a new project with a simple command and the project will be created with the structure you want.

	You can also create a new project with a template that you have already created.
	`,
	
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


