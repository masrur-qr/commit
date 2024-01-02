package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// GitHub credentials
const (
	username     = "masrur-qr"
	repoName     = "commit"
	accessToken  = "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"
	repoDirectory = "./"
)

// GitHub API URL
var  apiURL = fmt.Sprintf("https://%v@github.com/%v/%v.git",accessToken,username,repoName)
// const apiURL = "https://%vgithub.com/masrur-qr/commit.git"

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
	fmt.Println("git", "push", apiURL , "golang")
	cmd = exec.Command("git", "push", apiURL , "golang")
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
