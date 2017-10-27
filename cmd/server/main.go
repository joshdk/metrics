// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/flag"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/joshdk/metrics/cmd/version"
	"github.com/joshdk/metrics/protocol"
	"github.com/joshdk/metrics/protocol/impl"
)

var (
	configParam = flag.StringParam{
		Name:  "config",
		Usage: "Server config file",
	}
)

type PostgresConfig struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Config struct {
	Address  string         `yaml:"address"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func Cmd() *cli.App {

	app := cli.NewApp()

	app.Name = "metrics-server"
	app.Description = "Start a metrics server"
	app.Version = version.Version()

	app.Flags = []flag.Flag{
		configParam,
	}

	app.ErrorHandler = func(ctx cli.Context, err error) int {
		log.Fatalf("%s: %s\n", app.Name, err.Error())
		return 1
	}

	app.Action = func(ctx cli.Context) error {

		configPath := ctx.String(configParam.Name)
		if configPath == "" {
			configPath = "server.yml"
		}

		contents, err := ioutil.ReadFile(configPath)
		if err != nil {
			return errors.Wrap(err, "could not read config file")
		}

		config := Config{}

		if err := yaml.Unmarshal(contents, &config); err != nil {
			return errors.Wrap(err, "could not parse config file")
		}

		svr, err := impl.NewPostgres(
			config.Postgres.Hostname,
			config.Postgres.Port,
			config.Postgres.Username,
			config.Postgres.Password,
			config.Postgres.Dbname,
		)
		if err != nil {
			return errors.Wrap(err, "could not connect to Postgres database")
		}

		return protocol.Serve(config.Address, svr)
	}

	return app
}

func main() {
	os.Exit(Cmd().Run(os.Args))
}
