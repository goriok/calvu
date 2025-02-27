package git

import (
  "fmt"
  "os/exec"
  "time"
  "bytes"
  "strings"
)

func HeadDate() (*time.Time, error){
  err := validate()
  if err != nil {
    return nil, err
  }

	cmd := exec.Command("git", "log", "-1", "--format=%cd", "--date=iso")

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to get HEAD commit date: %v", err)
	}

	headTs := strings.TrimSpace(out.String())
	headTime, err := time.Parse("2006-01-02 15:04:05 -0700", headTs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse commit date: %v", err)
	}

  return &headTime, nil
}

func Bump(tag string) error {
  err := validate()
  if err != nil {
    return err
  }

	cmd := exec.Command("git", "tag", "-a", tag, "-m", "Release "+tag)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create tag '%s': %v", tag, err)
	}

	cmdPush := exec.Command("git", "push", "origin", tag)
	cmdPush.Stdout = &out
	cmdPush.Stderr = &out

	err = cmdPush.Run()
	if err != nil {
		return fmt.Errorf("failed to push tag '%s': %v", tag, err)
	}

	return nil
}

func Commit(filePath, message string) error {
  err := validate()
  if err != nil {
		return err
	}

  stageClean, err := isStageClean()
  if err != nil {
    return err
  }

  if !stageClean {
    return fmt.Errorf("commit action requires a clean git stage")
  }

	cmdAdd := exec.Command("git", "add", filePath)
	var out bytes.Buffer
	cmdAdd.Stdout = &out
	cmdAdd.Stderr = &out
	err = cmdAdd.Run()
	if err != nil {
		return fmt.Errorf("failed to stage file '%s': %v", filePath, err)
	}

	cmdCommit := exec.Command("git", "commit", "-m", message, "--", filePath)
	cmdCommit.Stdout = &out
	cmdCommit.Stderr = &out
	err = cmdCommit.Run()
	if err != nil {
		return fmt.Errorf("failed to commit file '%s': %v", filePath, err)
	}

	fmt.Printf("✅ File '%s' committed successfully with message: '%s'\n", filePath, message)
	return nil
}

func Push() error {
  err := validate()
  if err != nil {
		return err
	}

  stageClean, err := isStageClean()
  if err != nil {
    return err
  }

  if !stageClean {
    return fmt.Errorf("push action requires a clean git stage")
  }

	var out bytes.Buffer
  cmdPush := exec.Command("git", "push")
	cmdPush.Stdout = &out
	cmdPush.Stderr = &out
	err = cmdPush.Run()
	if err != nil {
		return fmt.Errorf("failed to push")
	}

	return nil
}

func validate() error {
  cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
  if err != nil {
    return err
  }

  if strings.TrimSpace(out.String()) != "true" {
    return fmt.Errorf("This cmd must run inside a git worktree")
  }

  return nil
}

func isStageClean() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("failed to check git status: %v", err)
	}

	return strings.TrimSpace(out.String()) == "", nil
}
