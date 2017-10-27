// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package impl

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Blank import to load the postgres driver
	"github.com/pkg/errors"

	"github.com/joshdk/metrics/protocol"
)

const (
	createTableStmt = `
		CREATE TABLE IF NOT EXISTS metrics (
			id    VARCHAR(32) PRIMARY KEY NOT NULL,
			count INTEGER                 NOT NULL,
			time  INTEGER                 NOT NULL
		)
	`

	insertStmt = `
		INSERT INTO metrics (id, count, time)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET count = metrics.count + $2
	`

	selectStmt = `
		SELECT count, time
		FROM metrics
		WHERE $1 <= time AND time <= $2
		LIMIT $3
	`
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(host string, port int, user string, password string, dbname string) (*Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open connection")
	}

	log.Printf("Connected to database at %s:%d/%s with %s\n", host, port, dbname, user)

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping connection")
	}

	if err := create(db); err != nil {
		return nil, errors.Wrap(err, "failed to create table")
	}

	return &Postgres{
		db: db,
	}, nil
}

func (s *Postgres) WriteMetric(ctx context.Context, in *protocol.MetricWriteRequest) (*protocol.MetricWriteReply, error) {

	if err := insert(s.db, in.Id, in.Count, time.Now().Unix()); err != nil {
		return nil, err
	}

	return &protocol.MetricWriteReply{
		Error: nil,
	}, nil
}

func (s *Postgres) QueryMetrics(ctx context.Context, in *protocol.QueryMetricsRequest) (*protocol.QueryMetricsReply, error) {

	results, err := query(s.db, in.Start, in.End, in.Count)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query")
	}

	return &protocol.QueryMetricsReply{
		Results: results,
		Error:   nil,
	}, nil
}

func create(db *sql.DB) error {

	_, err := db.Exec(createTableStmt)
	if err != nil {
		return err
	}

	log.Printf("Created table %q\n", "metrics")

	return nil
}

func insert(db *sql.DB, id string, count uint32, time int64) error {

	_, err := db.Exec(insertStmt, id, count, time)
	if err != nil {
		return errors.Wrap(err, "failed to insert value")
	}

	log.Printf("Inserted row %d @ %d %q\n", count, time, id)

	return nil
}

func query(db *sql.DB, start uint32, end uint32, count uint32) ([]*protocol.QueryMetricsReply_Result, error) {

	rows, err := db.Query(selectStmt, start, end, count)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err.Error())
		}
	}()

	points := []*protocol.QueryMetricsReply_Result{}

	for rows.Next() {

		var point protocol.QueryMetricsReply_Result

		if err := rows.Scan(&point.Count, &point.Time); err != nil {
			return nil, err
		}

		log.Printf("Retrieved row %d @ %d\n", point.Count, point.Time)

		points = append(points, &point)
	}

	// Check errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return points, nil
}
