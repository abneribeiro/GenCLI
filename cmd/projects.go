package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var projectsCmd = &cobra.Command{
	Use:   "proj",
	Short: "generate your projects from the terminal",
	Long: `GenCli is a CLI application that allows you to generate your projects from the terminal.

	For example:
	terminalProjects projects -n myProject.go
	`,
	Run: genProjects,
}

func init() {

	rootCmd.AddCommand(projectsCmd)

	projectsCmd.Flags().StringP("name", "n", "", "Specify the project name")

	projectsCmd.MarkFlagRequired("name")

}

func genProjects(cmd *cobra.Command, args []string) {
	projectName, _ := cmd.Flags().GetString("name")

	// Validate input
	if projectName == "" {
		log.Fatal("Project name is required")
	}

	err := os.Mkdir(projectName, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Create Go file inside the project directory
	filePath := filepath.Join(projectName, "main.go")
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = os.Chdir(projectName)
	if err != nil {
		log.Fatal(err)
	}

	gitCmd := exec.Command("git", "init")

	err = gitCmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("File created successfully")
}
