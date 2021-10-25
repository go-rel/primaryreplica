package primaryreplica

import (
	"context"
	"sync/atomic"

	"github.com/go-rel/rel"
)

type PrimaryReplica struct {
	primary     rel.Adapter
	replicas    []rel.Adapter
	replicasPtr int64
	replicasLen int64
}

func (pr *PrimaryReplica) Close() error {
	for i := range pr.replicas {
		if err := pr.replicas[i].Close(); err != nil {
			return err
		}
	}

	return pr.primary.Close()
}

func (pr *PrimaryReplica) Instrumentation(instrumenter rel.Instrumenter) {
	for i := range pr.replicas {
		pr.replicas[i].Instrumentation(instrumenter)
	}

	pr.primary.Instrumentation(instrumenter)
}

func (pr *PrimaryReplica) Ping(ctx context.Context) error {
	for i := range pr.replicas {
		if err := pr.replicas[i].Ping(ctx); err != nil {
			return err
		}
	}

	return pr.primary.Ping(ctx)
}

func (pr *PrimaryReplica) Aggregate(ctx context.Context, query rel.Query, mode string, field string) (int, error) {
	return pr.readAdapter(query).Aggregate(ctx, query, mode, field)
}

func (pr *PrimaryReplica) Query(ctx context.Context, query rel.Query) (rel.Cursor, error) {
	if query.LockQuery != "" {
		return pr.writeAdapter().Query(ctx, query)
	}

	return pr.readAdapter(query).Query(ctx, query)
}

func (pr *PrimaryReplica) Exec(ctx context.Context, stmt string, args []interface{}) (int64, int64, error) {
	return pr.writeAdapter().Exec(ctx, stmt, args)
}

func (pr *PrimaryReplica) Insert(ctx context.Context, query rel.Query, primaryField string, mutates map[string]rel.Mutate) (interface{}, error) {
	return pr.writeAdapter().Insert(ctx, query, primaryField, mutates)
}

func (pr *PrimaryReplica) InsertAll(ctx context.Context, query rel.Query, primaryField string, fields []string, bulkMutates []map[string]rel.Mutate) ([]interface{}, error) {
	return pr.writeAdapter().InsertAll(ctx, query, primaryField, fields, bulkMutates)
}

func (pr *PrimaryReplica) Update(ctx context.Context, query rel.Query, primaryField string, mutates map[string]rel.Mutate) (int, error) {
	return pr.writeAdapter().Update(ctx, query, primaryField, mutates)
}

func (pr *PrimaryReplica) Delete(ctx context.Context, query rel.Query) (int, error) {
	return pr.writeAdapter().Delete(ctx, query)
}

func (pr *PrimaryReplica) Apply(ctx context.Context, migration rel.Migration) error {
	return pr.writeAdapter().Apply(ctx, migration)
}

func (pr *PrimaryReplica) Begin(ctx context.Context) (rel.Adapter, error) {
	return pr.writeAdapter().Begin(ctx)
}

func (pr *PrimaryReplica) Commit(ctx context.Context) error {
	// this line shouldn't be accessible because transaction doesn't use this adapter
	return pr.writeAdapter().Commit(ctx)
}

func (pr *PrimaryReplica) Rollback(ctx context.Context) error {
	// this line shouldn't be accessible because transaction doesn't use this adapter
	return pr.writeAdapter().Rollback(ctx)
}

func (pr *PrimaryReplica) readAdapter(query rel.Query) rel.Adapter {
	if query.UsePrimaryDb {
		return pr.primary
	}

	return pr.replicas[atomic.AddInt64(&pr.replicasPtr, 1)%pr.replicasLen]
}

func (pr *PrimaryReplica) writeAdapter() rel.Adapter {
	return pr.primary
}

func New(primary rel.Adapter, replicas ...rel.Adapter) rel.Adapter {
	if len(replicas) == 0 {
		panic("rel: at least 1 replica is required")
	}

	return &PrimaryReplica{
		primary:     primary,
		replicas:    replicas,
		replicasLen: int64(len(replicas)),
	}
}
