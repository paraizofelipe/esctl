package cmd

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/elastic/go-elasticsearch/v8/typedapi/security/putuser"
	"github.com/paraizofelipe/esctl/internal/client"
	"github.com/urfave/cli/v2"
)

type SecurityUser struct {
	putuser.Request
}

func ChangeSecurityCommand() *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Modify Elasticsearch security settings, such as user configurations",
		Subcommands: []*cli.Command{
			ChangeSecurityUserCommand(),
		},
	}
}

func ApplySecurityUsers(ctx *cli.Context, users []SecurityUser) error {
	es := ctx.Context.Value("esClient").(client.ElasticClient)
	for _, user := range users {
		body := esutil.NewJSONReader(user)
		request := &esapi.SecurityPutUserRequest{
			Username: *user.Username,
			Body:     body,
		}
		jsonBytes, err := es.ExecRequest(ctx.Context, request)
		fmt.Println(string(jsonBytes))
		return err
	}
	return nil
}

func ChangeSecurityUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "Update settings for an Elasticsearch user",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "body",
				Usage:   "JSON-formatted body defining the new user settings",
				Aliases: []string{"b"},
			},
		},
		Action: func(ctx *cli.Context) error {
			var body *strings.Reader

			es := ctx.Context.Value("esClient").(client.ElasticClient)
			body = strings.NewReader(ctx.String("body"))
			request := &esapi.SecurityPutUserRequest{
				Username: ctx.Args().First(),
				Body:     body,
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			fmt.Println(string(jsonBytes))
			return err
		},
	}
}

func DescribeSecurityUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "Retrieve and display the security settings of specified Elasticsearch users",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "name",
				Usage:   "Specify the name(s) of the Elasticsearch user(s) to describe",
				Aliases: []string{"n"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			users := ctx.StringSlice("name")
			request := &esapi.SecurityGetUserRequest{
				Username: users,
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			fmt.Println(string(jsonBytes))
			return err
		},
	}
}

func DeleteSecurityUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "Delete a specified user from Elasticsearch security settings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Usage:   "Specify the name of the Elasticsearch user to be deleted",
				Aliases: []string{"n"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(client.ElasticClient)
			user := ctx.String("name")
			request := &esapi.SecurityDeleteUserRequest{
				Username: user,
			}
			jsonBytes, err := es.ExecRequest(ctx.Context, request)
			fmt.Println(string(jsonBytes))
			return err
		},
	}
}

func DeleteSecurityCommand() *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Remove security settings, including user configurations, in Elasticsearch",
		Subcommands: []*cli.Command{
			DeleteSecurityUserCommand(),
		},
	}
}

func DescribeSecurityCommand() *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "View detailed security settings in Elasticsearch, including user configurations",
		Subcommands: []*cli.Command{
			DescribeSecurityUserCommand(),
		},
	}
}
