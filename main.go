package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// GitHub credentials
const (
	username      = "masrur-qr"
	repoName      = "commit"
	accessToken   = "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"
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


	// git config --global user.email "you@example.com"
	cmd = exec.Command("git", "config", "--global", "user.email",`"you@example.com"`)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	//  git config --global user.name "Your Name"
	cmd = exec.Command("git", "config", "--global", "user.name",`"masrur"`)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", fmt.Sprintf("Auto commit at: %s", time.Now().Format("2006-01-02 15:04:05")))
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Push to the repository
	cmd = exec.Command("git", "push", "origin", "golang")
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
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			job()
		}
	}
}
