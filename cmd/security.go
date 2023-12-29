package cmd

import (
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

func ChangeSecurityCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Modify Elasticsearch security settings, such as user configurations",
		Subcommands: []*cli.Command{
			ChangeSecurityUserCommand(es),
		},
	}
}

func ApplySecurityUsers(ctx *cli.Context, es client.ElasticClient, users []SecurityUser) error {
	for _, user := range users {
		body := esutil.NewJSONReader(user)
		request := &esapi.SecurityPutUserRequest{
			Username: *user.Username,
			Body:     body,
			Pretty:   true,
		}
		return es.ExecRequest(ctx.Context, request)
	}
	return nil
}

func ChangeSecurityUserCommand(es client.ElasticClient) *cli.Command {
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

			body = strings.NewReader(ctx.String("body"))
			request := &esapi.SecurityPutUserRequest{
				Username: ctx.Args().First(),
				Body:     body,
				Pretty:   true,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DescribeSecurityUserCommand(es client.ElasticClient) *cli.Command {
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
			users := ctx.StringSlice("name")
			request := &esapi.SecurityGetUserRequest{
				Username: users,
				Pretty:   true,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DeleteSecurityUserCommand(es client.ElasticClient) *cli.Command {
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
			user := ctx.String("name")
			request := &esapi.SecurityDeleteUserRequest{
				Username: user,
				Pretty:   true,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DeleteSecurityCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Remove security settings, including user configurations, in Elasticsearch",
		Subcommands: []*cli.Command{
			DeleteSecurityUserCommand(es),
		},
	}
}

func DescribeSecurityCommand(es client.ElasticClient) *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "View detailed security settings in Elasticsearch, including user configurations",
		Subcommands: []*cli.Command{
			DescribeSecurityUserCommand(es),
		},
	}
}
