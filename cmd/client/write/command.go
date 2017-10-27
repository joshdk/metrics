// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package write

import (
	"fmt"

	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/joshdk/metrics/protocol"
	//"errors"
)

var (
	idParam = flag.StringFlag{
		Name:  "id",
		Usage: "Unique metric id",
	}
	countParam = flag.IntFlag{
		Name:  "count",
		Usage: "Value of metric point",
	}
)

func Command() cli.Command {
	cmd := cli.Command{}

	cmd.Name = "write"
	cmd.Description = "Write a metric value"

	cmd.Flags = []flag.Flag{
		idParam,
		countParam,
	}

	cmd.Action = func(ctx cli.Context) error {
		addr := ctx.String("address")
		if addr == "" {
			return fmt.Errorf("Missing flag %s", "address")
		}

		id := ctx.String(idParam.Name)
		if id == "" {
			return fmt.Errorf("Missing flag %s", idParam.Name)
		}

		count := ctx.Int(countParam.Name)
		if count == 0 {
			return fmt.Errorf("Missing flag %s", countParam.Name)
		}

		return Write(addr, id, count)
	}

	return cmd
}

func Write(address string, id string, count int) error {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err.Error())
		}
	}()

	client := pb.NewMetricsClient(conn)

	req := pb.MetricWriteRequest{
		Id:    id,
		Count: uint32(count),
	}

	fmt.Printf("writing metric %v\n", req)
	_, err = client.WriteMetric(context.Background(), &req)
	if err != nil {
		panic(err.Error())
	}

	return nil
}
