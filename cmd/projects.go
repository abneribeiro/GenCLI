package cmd

import "github.com/spf13/cobra"

// ProjectsCmd represents the projects command
var ProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "generate your projects from the terminal",
	Long: `terminalProjects is a CLI application that allows you to generate your projects from the terminal.`,
	Run: genProjects,
}

func init() {
	
	rootCmd.AddCommand(ProjectsCmd)

}

func genProjects(cmd *cobra.Command, args []string) {

}