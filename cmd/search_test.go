package cmd

import (
	"context"
	"os"
	"testing"

	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/paraizofelipe/esctl/internal/file"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"go.uber.org/mock/gomock"
)

func helperWriteFile(t *testing.T, filePath string, content string) {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		assert.NoError(t, err)
	}
}

func helperCreateFile(t *testing.T, filename string) {
	testFile, err := os.Create(filename)
	if err != nil {
		assert.NoError(t, err)
	}
	testFile.Close()
}

func setup(t *testing.T, fileName string) func() {
	t.Log("setup")
	helperCreateFile(t, fileName)
	return func() {
		t.Log("teardown")
		os.Remove(fileName)
	}
}

func TestSearchCommandWithQueryFlag(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockElastic := client.NewMockElasticClient(ctrl)
	mockElastic.EXPECT().ExecRequest(gomock.Any(), gomock.Any()).Return(nil, nil)

	ctx := context.WithValue(context.Background(), "esClient", mockElastic)

	app := &cli.App{
		Commands: []*cli.Command{
			SearchCommand(file.NewMockEditor(ctrl)),
		},
	}

	args := []string{"app", "search", "-q", "query", "index"}
	err := app.RunContext(ctx, args)
	assert.NoError(t, err)
}

func TestSearchCommandWithEditorFlag(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockElastic := client.NewMockElasticClient(ctrl)
	mockElastic.EXPECT().ExecRequest(gomock.Any(), gomock.Any()).Return(nil, nil)
	mockEditor := file.NewMockEditor(ctrl)
	mockEditor.EXPECT().ExecEditor(gomock.Any()).Return(`{ "query": { "match_all": {} } }`, nil)

	ctx := context.WithValue(context.Background(), "esClient", mockElastic)

	app := &cli.App{
		Commands: []*cli.Command{
			SearchCommand(mockEditor),
		},
	}

	args := []string{"app", "search", "-e", "index"}
	err := app.RunContext(ctx, args)
	assert.NoError(t, err)
}

func TestSearchCommandWithFileFlag(t *testing.T) {
	var filename string = "file.json"

	teardown := setup(t, filename)
	defer teardown()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	helperWriteFile(t, filename, `{ "query": { "match_all": {} } }`)

	mockElastic := client.NewMockElasticClient(ctrl)
	mockElastic.EXPECT().ExecRequest(gomock.Any(), gomock.Any()).Return(nil, nil)
	mockEditor := file.NewMockEditor(ctrl)

	ctx := context.WithValue(context.Background(), "esClient", mockElastic)

	app := &cli.App{
		Commands: []*cli.Command{
			SearchCommand(mockEditor),
		},
	}

	args := []string{"app", "search", "-f", "file.json", "index"}
	err := app.RunContext(ctx, args)
	assert.NoError(t, err)
}
