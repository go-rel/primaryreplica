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

func (na *nopAdapter) Instrumentation(_ rel.Instrumenter) {
}

func (na *nopAdapter) Ping(_ context.Context) error {
	return na.retError
}

func (na *nopAdapter) Aggregate(_ context.Context, _ rel.Query, _, _ string) (int, error) {
	return 0, na.retError
}

func (na *nopAdapter) Begin(_ context.Context) (rel.Adapter, error) {
	return na, na.retError
}

func (na *nopAdapter) Commit(_ context.Context) error {
	return na.retError
}

func (na *nopAdapter) Delete(_ context.Context, _ rel.Query) (int, error) {
	return 1, na.retError
}

func (na *nopAdapter) Insert(_ context.Context, _ rel.Query, _ string, _ map[string]rel.Mutate) (interface{}, error) {
	return 1, na.retError
}

func (na *nopAdapter) InsertAll(_ context.Context, _ rel.Query, _ string, _ []string, bulkMutates []map[string]rel.Mutate) ([]interface{}, error) {
	var (
		ids = make([]interface{}, len(bulkMutates))
	)

	for i := range bulkMutates {
		ids[i] = i + 1
	}

	return ids, na.retError
}

func (na *nopAdapter) Query(_ context.Context, _ rel.Query) (rel.Cursor, error) {
	return nil, nil
}

func (na *nopAdapter) Rollback(_ context.Context) error {
	return na.retError
}

func (na *nopAdapter) Update(_ context.Context, _ rel.Query, _ string, _ map[string]rel.Mutate) (int, error) {
	return 1, na.retError
}

func (na *nopAdapter) Apply(_ context.Context, _ rel.Migration) error {
	return na.retError
}

func (na *nopAdapter) Exec(_ context.Context, _ string, _ []interface{}) (int64, int64, error) {
	return 0, 0, na.retError
}
