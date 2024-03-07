package cmd

import (
	"log"
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "proj",
	Short: "generate your projects from the terminal",
	Long: `GenCli is a CLI application that allows you to generate your projects from the terminal.

	For example:
	terminalProjects projects -n myProject
	`,
	Run: genProjects,
}

func init() {

	rootCmd.AddCommand(projectsCmd)

	projectsCmd.Flags().StringP("name", "n","","Specify the project name")

	projectsCmd.MarkFlagRequired("name")

}


func genProjects(cmd *cobra.Command, args []string) {
	projectName, _ := cmd.Flags().GetString("name")

	// Validate input
	if projectName == "" {
		log.Fatal("Project name is required")
	}

	// Extract file extension
	fileExt := filepath.Ext(projectName)
	if fileExt == "" {
		log.Fatal("Invalid file extension")
	}

	// Remove leading dot from file extension
	fileExt = strings.TrimPrefix(fileExt, ".")

	file, err := os.Create(projectName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("File created successfully")
	fmt.Println("File Extension:", fileExt)
}



