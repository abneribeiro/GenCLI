package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// ProjectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "proj",
	Short: "generate your projects from the terminal",
	Long: `terminalProjects is a CLI application that allows you to generate your projects from the terminal.
	For example:
	terminalProjects projects -n myProject -t .go
	`,
	Run: genProjects,
}

func init() {

	rootCmd.AddCommand(projectsCmd)

	projectsCmd.Flags().StringP("name", "n","","Specify the programming language")
	projectsCmd.Flags().StringP("type", "t","","Specify the programming language")

	projectsCmd.MarkFlagRequired("name")
	projectsCmd.MarkFlagRequired("type")

}

func genProjects(cmd *cobra.Command, args []string) {
	projectName, _ := cmd.Flags().GetString("name")
	projectType, _ := cmd.Flags().GetString("type")

	project := projectName + projectType

	file, err := os.Create(project)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.Println("File created successfully")

}


