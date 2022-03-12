package primaryreplica

import (
	"context"
	"errors"
	"testing"

	"github.com/go-rel/rel"
	"github.com/stretchr/testify/assert"
)

func TestAdapter_Close(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{})
	assert.Nil(t, adapter.Close())
}

func TestAdapter_Close_replcaError(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("replica error")})
	assert.Equal(t, errors.New("replica error"), adapter.Close())
}

func TestAdapter_Instrumentation(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{})
	assert.NotPanics(t, func() {
		adapter.Instrumentation(rel.DefaultLogger)
	})
}

func TestAdapter_Ping(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{})
	assert.Nil(t, adapter.Ping(context.TODO()))
}

func TestAdapter_Ping_replcaError(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("replica error")})
	assert.Equal(t, errors.New("replica error"), adapter.Ping(context.TODO()))
}

func TestAdapter_Aggregate(t *testing.T) {
	adapter := New(&nopAdapter{retError: errors.New("should not use primary")}, &nopAdapter{})
	_, err := adapter.Aggregate(context.TODO(), rel.From("users"), "COUNT", "id")
	assert.Nil(t, err)
}

func TestAdapter_Query(t *testing.T) {
	adapter := New(&nopAdapter{retError: errors.New("should not use primary")}, &nopAdapter{})
	_, err := adapter.Query(context.TODO(), rel.From("users"))
	assert.Nil(t, err)
}

func TestAdapter_Query_UsePrimary(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.Query(context.TODO(), rel.From("users").UsePrimary())
	assert.Nil(t, err)
}

func TestAdapter_Query_LockForUpdate(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.Query(context.TODO(), rel.From("users").Lock("FOR UPDATE"))
	assert.Nil(t, err)
}

func TestAdapter_Exec(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, _, err := adapter.Exec(context.TODO(), "PING", nil)
	assert.Nil(t, err)
}

func TestAdapter_Insert(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.Insert(context.TODO(), rel.From("users"), "id", nil, rel.OnConflict{})
	assert.Nil(t, err)
}

func TestAdapter_InsertAll(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.InsertAll(context.TODO(), rel.From("users"), "id", nil, nil, rel.OnConflict{})
	assert.Nil(t, err)
}

func TestAdapter_Update(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.Update(context.TODO(), rel.From("users"), "id", nil)
	assert.Nil(t, err)
}

func TestAdapter_Delete(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.Delete(context.TODO(), rel.From("users"))
	assert.Nil(t, err)
}

func TestAdapter_Apply(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	err := adapter.Apply(context.TODO(), nil)
	assert.Nil(t, err)
}

func TestAdapter_Begin(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	_, err := adapter.Begin(context.TODO())
	assert.Nil(t, err)
}

func TestAdapter_Commit(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	err := adapter.Commit(context.TODO())
	assert.Nil(t, err)
}

func TestAdapter_Rollback(t *testing.T) {
	adapter := New(&nopAdapter{}, &nopAdapter{retError: errors.New("should not use replica")})
	err := adapter.Rollback(context.TODO())
	assert.Nil(t, err)
}

func TestNew_noReplicas(t *testing.T) {
	assert.Panics(t, func() {
		New(&nopAdapter{})
	})
}
