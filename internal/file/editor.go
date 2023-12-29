package file

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Editor interface {
	ExecEditor(filePath string) (content string, err error)
	Open(filePath string) (err error)
}

type TextEditor struct{}

func NewTextEditor() Editor {
	return &TextEditor{}
}

func (e *TextEditor) selectArgs(editorName string, fileName string) (args []string) {
	if editorName == "vim" || editorName == "nvim" {
		args = []string{"-c", "set filetype=json", fileName}
	}
	return
}

func (e *TextEditor) createFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func (e *TextEditor) validateFilePath(filePath string) (string, error) {
	var err error

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		filePath = fmt.Sprintf("/tmp/%d", rand.Int())
		if err = e.createFile(filePath); err != nil {
			return filePath, err
		}
	}
	return filePath, err
}

func (e *TextEditor) Open(filePath string) (err error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	args := e.selectArgs(editor, filePath)
	cmd := exec.Command(editor, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return
	}

	return
}

func (e *TextEditor) ExecEditor(filePath string) (content string, err error) {
	var (
		beforeFileInfo os.FileInfo
		afterFileInfo  os.FileInfo
		beforeModified time.Time
	)

	filePath, err = e.validateFilePath(filePath)
	if err != nil {
		return
	}

	beforeFileInfo, err = os.Stat(filePath)
	if err != nil {
		return
	}
	beforeModified = beforeFileInfo.ModTime()

	if err = e.Open(filePath); err != nil {
		return
	}

	if afterFileInfo, err = os.Stat(filePath); err != nil {
		return
	}
	afterModified := afterFileInfo.ModTime()

	if !afterModified.After(beforeModified) {
		return
	}

	bContent, err := os.ReadFile(filePath)
	content = string(bContent)
	if err != nil {
		return
	}

	return
}
