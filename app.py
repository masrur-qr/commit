# # Replace these values with your GitHub username, repository name, and personal access token
# github_username = 'masrur-qr'
# repo_name = 'commit'
# access_token = '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'

import os
import requests
import schedule
import time

# GitHub credentials
username = "masrur-qr"
repo_name = "commit"
token = "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"

# GitHub API URL
api_url = f"https://api.github.com/repos/{username}/{repo_name}/git/commits"

# Directory containing the files to commit
repo_directory = "./"

def create_commit_and_push():
    # Change to the repository directory
    os.chdir(repo_directory)
    print("hello")

    # Read the content of a sample file (you can customize this part)
    with open("sample_file.txt", "a") as file:
        file.write("\nCommit made at: " + time.strftime("%Y-%m-%d %H:%M:%S"))

    # Stage and commit changes
    os.system("git add .")
    os.system(f'git commit -m "Auto commit at: {time.strftime("%Y-%m-%d %H:%M:%S")}"')

    # Push to the repository
    os.system("git push origin main")

# def job():
#     print("Creating commit and pushing to GitHub...")
create_commit_and_push()
print("hello")


# # Schedule the job to run every hour
# schedule.every().minute.do(job)

# # Run the scheduler
# while True:
#     schedule.run_pending()
#     time.sleep(1)
