/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gitinitCmd represents the gitinit command
var gitinitCmd = &cobra.Command{
	Use:   "gitinit",
	Short: "Initialize and push a Git repo with one command",
	Run: func(cmd *cobra.Command, args []string) {
		runGitInitFlow()
		fmt.Println("gitinit called")
	},
}

func init() {
	rootCmd.AddCommand(gitinitCmd)
}

// func runGitInitFlow() {
// 	var remoteURL, branch, message, gitignoreType string

// 	defaultUrl := getDefaultRemoteURL()

// 	// Prompts
// 	survey.AskOne(&survey.Input{
// 		Message: "Remote repository URL:",
// 		Default: defaultUrl,
// 	}, &remoteURL, survey.WithValidator(survey.Required))

// 	survey.AskOne(&survey.Input{
// 		Message: "Default branch name:",
// 		Default: "main",
// 	}, &branch)

// 	survey.AskOne(&survey.Input{
// 		Message: "Commit message:",
// 		Default: "Initial commit",
// 	}, &message)

// 	survey.AskOne(&survey.Select{
// 		Message: "Add .gitignore for:",
// 		Options: []string{"None", "Go", "Node.js", "Python"},
// 		Default: "None",
// 	}, &gitignoreType)

// 	// Optional: Add .gitignore file
// 	if gitignoreType != "None" {
// 		addGitignore(gitignoreType)
// 	}

// 	run("git", "init")
// 	run("git", "add", ".")
// 	run("git", "commit", "-m", message)
// 	run("git", "branch", "-M", branch)
// 	run("git", "remote", "add", "origin", remoteURL)
// 	run("git", "push", "-u", "origin", branch)

// 	fmt.Println("✅ Git repo initialized and pushed!")

// }

// func run(name string, args ...string) {
// 	cmd := exec.Command(name, args...)

// 	output, err := cmd.CombinedOutput()

// 	if err != nil {
// 		fmt.Printf("❌ Error running %s %v:\n%s\n", name, args, output)
// 	} else {
// 		fmt.Printf("✅ Ran: %s %v\n", name, args)
// 	}
// }

// func addGitignore(lang string) {
// 	content := ""
// 	switch lang {
// 	case "Go":
// 		content = "bin/\n*.exe\n*.log\n*.test\nvendor/\n"
// 	case "Node.js":
// 		content = "node_modules/\ndist/\n.env\n*.log\n"
// 	case "Python":
// 		content = "__pycache__/\n*.pyc\n.env\nvenv/\n"
// 	}
// 	if content != "" {
// 		_ = os.WriteFile(".gitignore", []byte(content), 0644)
// 		fmt.Println("📄 .gitignore added")
// 	}
// }

// func getCurrentFolderName() string {
// 	path, err := os.Getwd()

// 	if err != nil {
// 		return "my-repo"
// 	}

// 	return filepath.Base(path)
// }

// func getGitUsername() string {
// 	cmd := exec.Command("git", "config", "--get", "user.name")
// 	output, err := cmd.Output()

// 	if err != nil {
// 		return "your-username"
// 	}

// 	return strings.TrimSpace(string(output))
// }

// func getDefaultRemoteURL() string {
// 	username := getGitUsername()
// 	repo := getCurrentFolderName()
// 	return "https://github.com/" + username + "/" + repo + ".git"
// }
