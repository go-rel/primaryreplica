package primaryreplica

import (
	"context"

	"github.com/go-rel/rel"
)

type nopAdapter struct {
	retError error
}

func (na *nopAdapter) Close() error {
	return na.retError
}

func (na *nopAdapter) Instrumentation(instrumenter rel.Instrumenter) {
}

func (na *nopAdapter) Ping(ctx context.Context) error {
	return na.retError
}

func (na *nopAdapter) Aggregate(ctx context.Context, query rel.Query, mode string, field string) (int, error) {
	return 0, na.retError
}

func (na *nopAdapter) Begin(ctx context.Context) (rel.Adapter, error) {
	return na, na.retError
}

func (na *nopAdapter) Commit(ctx context.Context) error {
	return na.retError
}

func (na *nopAdapter) Delete(ctx context.Context, query rel.Query) (int, error) {
	return 1, na.retError
}

func (na *nopAdapter) Insert(ctx context.Context, query rel.Query, primaryField string, mutates map[string]rel.Mutate) (interface{}, error) {
	return 1, na.retError
}

func (na *nopAdapter) InsertAll(ctx context.Context, query rel.Query, primaryField string, fields []string, bulkMutates []map[string]rel.Mutate) ([]interface{}, error) {
	var (
		ids = make([]interface{}, len(bulkMutates))
	)

	for i := range bulkMutates {
		ids[i] = i + 1
	}

	return ids, na.retError
}

func (na *nopAdapter) Query(ctx context.Context, query rel.Query) (rel.Cursor, error) {
	return nil, nil
}

func (na *nopAdapter) Rollback(ctx context.Context) error {
	return na.retError
}

func (na *nopAdapter) Update(ctx context.Context, query rel.Query, primaryField string, mutates map[string]rel.Mutate) (int, error) {
	return 1, na.retError
}

func (na *nopAdapter) Apply(ctx context.Context, migration rel.Migration) error {
	return na.retError
}

func (na *nopAdapter) Exec(ctx context.Context, stmt string, args []interface{}) (int64, int64, error) {
	return 0, 0, na.retError
}
