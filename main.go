package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GitHub credentials
const (
	username    = "masrur-qr"
	repoName    = "commit"
	accessToken = "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"
)

func createCommitAndPush(client *github.Client) {
	// Change to the repository directory
	os.Chdir(repoName)

	// Read the content of a sample file (you can customize this part)
	file, err := os.OpenFile("sample_file.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	_, err = file.WriteString("\nCommit made at: " + time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	// Stage and commit changes
	cmd := exec.Command("git", "add", ".")
	err = cmd.Run()
	if err != nil {
		log.Fatal("Error staging changes:", err)
	}

	cmd = exec.Command("git", "commit", "-m", fmt.Sprintf("Auto commit at: %s", time.Now().Format("2006-01-02 15:04:05")))
	err = cmd.Run()
	if err != nil {
		log.Fatal("Error committing changes:", err)
	}

	
	// Now create a commit using the GitHub API
	commitMessage := fmt.Sprintf("Auto commit at: %s", time.Now().Format("2006-01-02 15:04:05"))
	opts := &github.RepositoryContentFileOptions{
		Content: []byte(commitMessage),
		Message: github.String(commitMessage),
	}
	
	_, _, err = client.Repositories.CreateFile(context.Background(), username, repoName, "sample_file.txt", opts)
	if err != nil {
		log.Fatal("Error creating commit via GitHub API:", err)
	}
	// Push to the repository
	cmd = exec.Command("git", "push", "origin", "golang")
	err = cmd.Run()
	if err != nil {
		log.Fatal("Error pushing changes:", err)
	}
}

func job(client *github.Client) {
	fmt.Println("Creating commit and pushing to GitHub...")
	createCommitAndPush(client)
	fmt.Println("hello")
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Schedule the job to run every hour
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			job(client)
		}
	}
}
