// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package query

import (
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/joshdk/metrics/protocol"
)

var (
	startTimeParam = flag.StringFlag{
		Name:  "start",
		Usage: "Beginning of time interval",
	}
	endTimeParam = flag.StringFlag{
		Name:  "end",
		Usage: "End of time interval",
	}
	countParam = flag.IntFlag{
		Name:  "count",
		Usage: "Desired number of points",
	}
)

func Command() cli.Command {
	cmd := cli.Command{}

	cmd.Name = "query"
	cmd.Description = "Query for metrics withing a time interval"

	cmd.Flags = []flag.Flag{
		startTimeParam,
		endTimeParam,
		countParam,
	}

	cmd.Action = func(ctx cli.Context) error {
		addr := ctx.String("address")
		if addr == "" {
			return fmt.Errorf("Missing flag %s", "address")
		}

		start := ctx.String(startTimeParam.Name)
		if start == "" {
			return fmt.Errorf("Missing flag %s", startTimeParam.Name)
		}

		end := ctx.String(endTimeParam.Name)
		if end == "" {
			return fmt.Errorf("Missing flag %s", endTimeParam.Name)
		}

		count := ctx.Int(countParam.Name)
		if count == 0 {
			return fmt.Errorf("Missing flag %s", countParam.Name)
		}

		startTime, err := dateparse.ParseLocal(start)
		if err != nil {
			return fmt.Errorf("Could not parse start time value %q", start)

		}

		endTime, err := dateparse.ParseLocal(end)
		if err != nil {
			return fmt.Errorf("Could not parse start time value %q", end)

		}

		return Query(addr, startTime.Unix(), endTime.Unix(), count)
	}

	return cmd
}

func Query(address string, start int64, end int64, count int) error {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err.Error())
		}
	}()

	client := protocol.NewMetricsClient(conn)

	req := protocol.QueryMetricsRequest{
		Start: uint32(start),
		End:   uint32(end),
		Count: uint32(count),
	}

	reply, err := client.QueryMetrics(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d results\n", len(reply.Results))
	for _, result := range reply.Results {
		fmt.Printf(">>> %d @ %d\n", result.Count, result.Time)
	}

	return nil
}
