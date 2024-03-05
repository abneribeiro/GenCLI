/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Gen",
	Short: "Generate your projects from the terminal",
	Long: `terminalProjects is a CLI application that allows you to generate your projects from the terminal.
	
	You can create a new project with a simple command and the project will be created with the structure you want.

	You can also create a new project with a template that you have already created.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.terminalProjects.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


