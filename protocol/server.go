// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package protocol

import (
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve(address string, impl MetricsServer) error {

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	log.Printf("Now listening on %s\n", address)

	server := grpc.NewServer()
	RegisterMetricsServer(server, impl)
	reflection.Register(server)

	log.Println("Now serving")

	if err := server.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to server")
	}

	return nil
}
