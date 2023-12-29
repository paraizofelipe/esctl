package file

import (
	"os"
	reflect "reflect"
	"testing"
)

func TestSelectArgs(t *testing.T) {
	editor := NewTextEditor().(*TextEditor)

	tests := []struct {
		editorName string
		want       []string
	}{
		{"vim", []string{"-c", "set filetype=json", "file.json"}},
		{"nvim", []string{"-c", "set filetype=json", "file.json"}},
		{"emacs", nil},
	}

	for _, tt := range tests {
		got := editor.selectArgs(tt.editorName, "file.json")
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("selectArgs(%q) = %v, want %v", tt.editorName, got, tt.want)
		}
	}
}

func TestCreateFile(t *testing.T) {
	editor := NewTextEditor().(*TextEditor)
	filePath := "/tmp/testfile"

	defer os.Remove(filePath)

	err := editor.createFile(filePath)
	if err != nil {
		t.Errorf("createFile() error = %v, wantErr nil", err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("createFile() file %s does not exist", filePath)
	}
}

func TestValidateFilePath(t *testing.T) {
	editor := NewTextEditor().(*TextEditor)
	filePath := "/testfile.json"

	currentPath, err := editor.validateFilePath(filePath)
	if err != nil {
		t.Errorf("validateFilePath() error = %v, wantErr nil", err)
	}

	defer os.Remove(currentPath)

	if _, err := os.Stat(currentPath); os.IsNotExist(err) {
		t.Errorf("validateFilePath() file %s does not exist", filePath)
	}
}
