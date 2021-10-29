# primaryreplica

[![GoDoc](https://godoc.org/github.com/go-rel/primaryreplica?status.svg)](https://pkg.go.dev/github.com/go-rel/primaryreplica)
[![Test](https://github.com/go-rel/primaryreplica/actions/workflows/test.yml/badge.svg)](https://github.com/go-rel/primaryreplica/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-rel/primaryreplica)](https://goreportcard.com/report/github.com/go-rel/primaryreplica)
[![codecov](https://codecov.io/gh/go-rel/primaryreplica/branch/main/graph/badge.svg?token=snplQW05GK)](https://codecov.io/gh/go-rel/primaryreplica)
[![Gitter chat](https://badges.gitter.im/go-rel/rel.png)](https://gitter.im/go-rel/rel)

Read Write separation for primary-replica topologies

## Example

```go
package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-rel/primaryreplica"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
)
func main() {
	// open mysql connections.
	// note: `clientFoundRows=true` is required for update and delete to works correctly.
	adapter := primaryreplica.New(
		mysql.MustOpen("root@(source:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
		mysql.MustOpen("root@(replica1:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
		mysql.MustOpen("root@(replica2:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
	)
	defer adapter.Close()

	// initialize REL's repo.
	repo := rel.New(adapter)
	repo.Ping(context.TODO())
}
```

## Load Balancing of Replicas

REL only implements a very primitive load balancing for multiple replicas.
For large scale application we recommend you to use external load balancing solution.
