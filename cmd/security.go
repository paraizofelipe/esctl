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

func ChangeSecurityCommand() *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Change security settings",
		Subcommands: []*cli.Command{
			ChangeSecurityUserCommand(),
		},
	}
}

func ApplySecurityUsers(ctx *cli.Context, users []SecurityUser) error {
	es := ctx.Context.Value("esClient").(*client.Elastic)
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

func ChangeSecurityUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "Change user settings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "body",
				Aliases: []string{"b"},
			},
		},
		Action: func(ctx *cli.Context) error {
			var body *strings.Reader

			es := ctx.Context.Value("esClient").(*client.Elastic)
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

func DescribeSecurityUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "Describe user settings",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "name",
				Aliases: []string{"n"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			users := ctx.StringSlice("name")
			request := &esapi.SecurityGetUserRequest{
				Username: users,
				Pretty:   true,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DeleteSecurityUserCommand() *cli.Command {
	return &cli.Command{
		Name:  "user",
		Usage: "Delete user settings",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
			},
		},
		Action: func(ctx *cli.Context) error {
			es := ctx.Context.Value("esClient").(*client.Elastic)
			user := ctx.String("name")
			request := &esapi.SecurityDeleteUserRequest{
				Username: user,
				Pretty:   true,
			}
			return es.ExecRequest(ctx.Context, request)
		},
	}
}

func DeleteSecurityCommand() *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Delete security settings",
		Subcommands: []*cli.Command{
			DeleteSecurityUserCommand(),
		},
	}
}

func DescribeSecurityCommand() *cli.Command {
	return &cli.Command{
		Name:  "security",
		Usage: "Describe security settings",
		Subcommands: []*cli.Command{
			DescribeSecurityUserCommand(),
		},
	}
}
