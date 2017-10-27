// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package main

import (
	"fmt"
	"os"

	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/flag"

	"github.com/joshdk/metrics/cmd/client/query"
	"github.com/joshdk/metrics/cmd/client/write"
	"github.com/joshdk/metrics/cmd/version"
)

var (
	addressParam = flag.StringFlag{
		Name:  "address",
		Usage: "Address of metrics-server",
	}
)

func Cmd() *cli.App {

	app := cli.NewApp()

	app.Name = "metrics-client"
	app.Description = "Interact with a metrics server"
	app.Version = version.Version()

	app.Flags = []flag.Flag{
		addressParam,
	}

	app.Subcommands = []cli.Command{
		query.Command(),
		write.Command(),
	}

	app.ErrorHandler = func(ctx cli.Context, err error) int {
		fmt.Printf("%s: %s\n", app.Name, err.Error())
		return 1
	}

	return app
}

func main() {
	os.Exit(Cmd().Run(os.Args))
}
