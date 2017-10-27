[![License](https://img.shields.io/github/license/joshdk/metrics.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/metrics?status.svg)](https://godoc.org/github.com/joshdk/metrics)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/metrics)](https://goreportcard.com/report/github.com/joshdk/metrics)
[![CircleCI](https://circleci.com/gh/joshdk/metrics.svg?&style=shield)](https://circleci.com/gh/joshdk/metrics/tree/master)

# Metrics

📈 Sample metrics api server

## Usage

### Building the required Docker image

You can run the following command to produce a docker image of `metrics-server:latest`.

```
$ ./godelw dist && ./godelw docker build
Building metrics-server for linux-amd64 at /Users/josh/Code/src/github.com/joshdk/metrics/build/bin/0.0.0-4-g9fd97a2.dirty/linux-amd64/metrics-server
...
Building docker image for metrics-server and tagging it as metrics-server:latest
```

### Launching the Docker Compose cluster

You can start the cluster (Postgres database + metrics server) by running the following.

```
$ cd example

$ $ docker-compose up
Creating example_db_1 ...
Creating example_db_1 ... done
Creating example_metrics-server_1 ...
Creating example_metrics-server_1 ... done
Attaching to example_db_1, example_metrics-server_1
metrics-server_1  | 2017/10/27 11:11:56 Connected to database at db:5432/metrics with postgres
metrics-server_1  | 2017/10/27 11:11:56 Created table "metrics"
metrics-server_1  | 2017/10/27 11:11:56 Now listening on 0.0.0.0:9001
metrics-server_1  | 2017/10/27 11:11:56 Now serving
db_1              | The files belonging to this database system will be owned by user "postgres".
db_1              | This user must also own the server process.
db_1              |
db_1              | The database cluster will be initialized with locale "en_US.utf8".
...
db_1              | PostgreSQL init process complete; ready for start up.
```

### Verification

You can verify that the metrics service is listing on port `9001`, by running the following.

```
$ lsof -nP -i4TCP | grep LISTEN
vpnkit    54882 josh   23u  IPv4 0xad85ffeb367269d1      0t0  TCP *:9001 (LISTEN)
```

### Writing Metrics

This project comes pre-bundled with an example client that wraps the gRPC bindings.

You can write a metric like so. Where `VERSION` is the current value reported by `git describe --tags`, and `UUID` is a random identifier meant to collect multiple reports of the same metric point.

```
$ ./build/bin/<VERSION>/darwin-amd64/metrics-client --address localhost:9001 write --id <UUID> --count 83
writing metric {<UUID> 83}
```

### Querying Metrics

You can write a metric like so. Where `start` and `end` are Unix epoc timestamps specifying a search range, and `count` is the maximum number of points to report.

```
$ ./build/bin/0.0.0-4-g9fd97a2.dirty/darwin-amd64/metrics-client --address localhost:9001 query --start 0 --end 1519103734 --count 10
Found 3 results
>>> 56 @ 1509103430
>>> 76 @ 1509103771
>>> 83 @ 1509103774
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.
