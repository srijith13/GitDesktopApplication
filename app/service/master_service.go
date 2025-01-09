package service

import (
	"fmt"
	"git-visualizer/app/dto"
	"os"
	"os/exec"
	"strings"
)

// CloneRepo clones a Git repository to a specified directory using the `git` CLI.
func GitCloneRepo(req *dto.Request) error {
	cmd := exec.Command("git", "clone", req.RepoURL, req.Dir)
	cmd.Env = append(os.Environ(), fmt.Sprintf("GIT_ASKPASS=echo"), fmt.Sprintf("GIT_USERNAME=%s", req.Username), fmt.Sprintf("GIT_PASSWORD=%s", req.Password))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone repo: %w", err)
	}
	return nil
}

func GetRepoBranches(req *dto.Request) ([]string, error) {
	cmd := exec.Command("git", "-C", req.Dir, "branch", "-a") // "  origin/HEAD -> origin/main",
	// cmd := exec.Command("", "-c", fmt.Sprintf("git -C %s branch -r | awk -F'/' '{print $2}'", req.Dir)) //ideal for mac and linus
	// cmd := exec.Command("powershell", "-Command", fmt.Sprintf("git -C %s branch -r | ForEach-Object { ($_ -split '/')[1] }", req.Dir))// for windowa powershell but need to make fixes

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	// Output is a list of branches, so we need to process it
	branches := []string{}
	lines := strings.Split(string(output), "\n")
	for i := 0; i < len(lines)-1; i++ {
		if !strings.Contains(lines[i], "origin/HEAD") {
			branchName := lines[i]
			branches = append(branches, branchName)
		}
	}

	// // to remove origin/ for remote branches
	// for i := 0; i < len(lines)-1; i++ {
	// 	line := strings.TrimSpace(lines[i])
	// 	if line != "" {
	// 		parts := strings.Split(line, "/")
	// 		if len(parts) > 1 {
	// 			branches = append(branches, parts[1])
	// 		}
	// 	}
	// }

	// remove duplicates in branch name from remote origins
	// make separate  array for remote only branches which are not present in local

	// to get only local branches  use command git -C <DIR> branch and without current branch use git -C <DIR>  branch --format="%(refname:short)"
	// to get only remote branches git -C <DIR> branch -r and to remove origin prefix use git -C <DIR>  branch -r | awk -F'/' '{print $2}'  ideal for mac and linux
	// to get all local and remote branch use  git -C <DIR> branch -a

	return branches, nil

}

func GetBranchStatus(req *dto.Request) ([]string, error) {
	cmd := exec.Command("git", "-C", req.Dir, "status", "-b")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed get branch status: %w", err)
	}

	// Summary of Common Flags
	// -s or --short: Short output.
	// -b or --branch: Show the current branch.
	// -u <mode> or --untracked-files <mode>: Control visibility of untracked files.
	// --ignored: Include ignored files in the output.
	// -v or --verbose: Show more detailed output.
	// --porcelain: Output in a machine-readable format.
	// --show-stash: Show any stashes.
	// --ahead-behind: Display ahead/behind info.
	// --no-ahead-behind: Hide ahead/behind info.
	// -h or --help: Show help.

	status := strings.Split(string(output), "\n")
	if len(status) > 0 && status[len(status)-1] == "" {
		status = status[:len(status)-1]
	}

	return status, nil

}

func GetLogs(req *dto.Request) ([]string, error) {
	cmd := exec.Command("git", "-C", req.Dir, "log", "--pretty=format:%h %an %ad %s", "--date=iso")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed get branch status: %w", err)
	}

	// Summary of Common git log Flags:
	// --oneline: Compact commit summary.
	// -n <number> or --max-count=<number>: Limit the number of commits.
	// --author=<author>: Filter by author.
	// --since=<date>  or --after=<date>: Show commits after a date.
	// 	//  // You can use various date formats like: yyyy-mm-dd, yesterday, 2 weeks ago,  3 days ago
	// 	--until=<date> or --before=<date>: Show commits before a date.
	// --grep=<pattern>: Filter by commit message.
	// --stat: Show file change statistics.
	// --patch or -p: Show diffs (patches) for commits.
	// --pretty=<format>: Customize output format.   --pretty=format:"%h %an %ad %s"
	// 	//  // This shows:
	// 	//  //	// %h: Abbreviated commit hash
	// 	//  //	// %an: Author name
	// 	//  //	// %ad: Author date (in default format)
	// 	//  //	// %s: Commit message
	// --graph: Visualize commit history with a graph.   --oneline --graph --all
	// --all: Include all branches in the history.
	// --merges/--no-merges: Filter merge commits.
	// --reverse: Show commits from oldest to newest.
	// --date=<format>: Customize date format.
	// 	//  // Customizes the date format. Common options:
	// 	//  // default (default Git date format)
	// 	//  // relative (e.g., "2 weeks ago")
	// 	//  // iso (e.g., "2025-01-08 10:00:00")
	// 	//  // rfc (e.g., "Wed, 08 Jan 2025 10:00:00 -0500")
	// 	//  // short (e.g., "2025-01-08")
	// --skip=<number>: Skip the first N commits.
	// --max-parents=<number> :Limits the number of parent commits to a specific number. This is useful for showing only commits with a particular number of parents (e.g., merge commits).
	// -show-pull-statistics : Shows statistics about pull requests (if using GitHub, for example).
	logs := strings.Split(string(output), "\n")

	if len(logs) > 0 && logs[len(logs)-1] == "" {
		logs = logs[:len(logs)-1]
	}

	return logs, nil
}

func GitCheckoutBranches(req *dto.Request) error { // make switch case to handle various checkout flags

	// cmd := exec.Command("git", "-C", req.Dir, "checkout", req.Branch)

	// output, err := cmd.CombinedOutput() // Try in future when the errors are not needed individually to be debunged

	args := []string{"-C", req.Dir, "checkout", req.Branch}
	switch req.Flag {
	case "new":
		args = append(args, "-b")
	case "force":
		args = append(args, "-f")
	}
	// args = append(args, req.Flags) // Add all additional flags

	// Create the command
	cmd := exec.Command("git", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("failed checkout to branch %s : %w", req.Branch, err)
	}
	// Switching to an existing branch:
	// 	//git checkout <branch-name>
	// Creating a new branch and switching to it:
	// 	//git checkout -b <new-branch-name>
	// Checking out a commit (detached HEAD state):
	// 	//git checkout <commit-hash>
	// Restoring a file to the latest commit:
	// 	//git checkout -- <file-path>
	// Checkout and create a new branch that tracks the remote:
	// 	//git checkout --track origin/<remote-branch-name>
	// Force checkout of a branch, discarding local changes:
	// 	//git checkout -f <branch-name>

	return nil
}

func GitStash(req *dto.Request) ([]string, error) {

	args := []string{"-C", req.Dir, "stash"}
	switch req.Flag {
	case "push":
		args = append(args, "push")
	case "pop":
		args = append(args, "pop")
	case "apply":
		args = append(args, "apply")
	case "list":
		args = append(args, "list")
	case "drop":
		args = append(args, "drop")
	case "clear":
		args = append(args, "clear")
	case "diff":
		args = append(args, "diff")
	}

	// cmd := exec.Command("git", "-C", req.Dir, "stash", "",)
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to stash changes: %w", err)
	}
	stash := strings.Split(string(output), "\n")

	if len(stash) > 0 && stash[len(stash)-1] == "" {
		stash = stash[:len(stash)-1]
	} else {
		stash = nil
	}

	return stash, nil
}

func GitAddCommitFiles(req *dto.Request) error {
	args := []string{"-C", req.Dir, "add"}
	if len(req.Files) > 0 {
		fmt.Println("this not")
		args = append(args, req.Files...)
	} else {
		fmt.Println("this")
		args = append(args, ".")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("failed to delete last commit: %w", err)
	}

	return nil
}

func GitCommitChanges(req *dto.Request) error {
	// add flags later with different conditions
	// args := []string{"-C", req.Dir, "add"}
	// if len(req.Files) > 0 {
	// 	fmt.Println("this not")
	// 	args = append(args, req.Files...)
	// } else {
	// 	fmt.Println("this")
	// 	args = append(args, ".")
	// }

	// cmd := exec.Command("git", args...)

	cmd := exec.Command("git", "-C", req.Dir, "commit", "-m", req.Message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to delete last commit: %w", err)
	}

	return nil
}

func GitDeleteLastCommit(req *dto.Request) error {
	cmd := exec.Command("git", "-C", req.Dir, "reset", "--hard", "HEAD~1")

	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to delete last commit: %w", err)
	}

	return nil
}

func GitPushChanges(req *dto.Request) error {
	cmd := exec.Command("git", "-C", req.Dir, "push", req.RemoteBranch, req.Branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to delete last commit: %w", err)
	}

	return nil
}

func GitPullChanges(req *dto.Request) error {
	cmd := exec.Command("git", "-C", req.Dir, "pull", req.RemoteBranch, req.Branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to delete last commit: %w", err)
	}

	return nil
}
