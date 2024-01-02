package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	// "github.com/google/go-github/github"
	// "golang.org/x/oauth2"
)

// GitHub credentials
const (
	username     = "masrur-qr"
	repoName     = "commit"
	accessToken  = "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"
	repoDirectory = "./"
)

// GitHub API URL
const apiURL = "https://api.github.com/repos/"

func createCommitAndPush() {
	// Change to the repository directory
	os.Chdir(repoDirectory)
	fmt.Println("hello")

	// Read the content of a sample file (you can customize this part)
	file, err := os.OpenFile("sample_file.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nCommit made at: " + time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Stage and commit changes
	cmd := exec.Command("git", "add", ".")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error staging changes:", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", fmt.Sprintf("Auto commit at: %s", time.Now().Format("2006-01-02 15:04:05")))
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Push to the repository
	cmd = exec.Command("git", "push", "origin", "main")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error pushing changes:", err)
		return
	}
}

func job() {
	fmt.Println("Creating commit and pushing to GitHub...")
	createCommitAndPush()
	fmt.Println("hello")
}

func main() {
	// Schedule the job to run every hour
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			job()
		}
	}
}