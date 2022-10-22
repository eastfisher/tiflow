package http

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/pingcap/tiflow/cdc/model"
	"github.com/stretchr/testify/require"
)

func TestInitHTTPPluginSink(t *testing.T) {
	ctx := context.Background()
	sinkURI, err := url.Parse("http://localhost:5005/sink_sync")
	require.NoError(t, err)
	sink, err := NewHTTPPluginSink(ctx, sinkURI, nil, nil, nil)
	require.NoError(t, err)
	err = sink.AddTable(time.Now().Unix())
	require.NoError(t, err)
}

func TestHTTPPluginRowEvent(t *testing.T) {
	ctx := context.Background()
	sinkURI, err := url.Parse("http://localhost:5005/sink_sync")
	require.NoError(t, err)
	sink, err := NewHTTPPluginSink(ctx, sinkURI, nil, nil, nil)
	require.NoError(t, err)
	err = sink.AddTable(time.Now().Unix())
	require.NoError(t, err)

	events := getTestRowEvents()
	err = sink.EmitRowChangedEvents(ctx, events...)
	require.NoError(t, err)
}

func getTestRowEvents() []*model.RowChangedEvent {
	e1 := &model.RowChangedEvent{
		StartTs:  1,
		CommitTs: 1,
		RowID:    1,
		Table: &model.TableName{
			Schema:      "db1",
			Table:       "tbl1",
			TableID:     1,
			IsPartition: false,
		},
		ColInfos:         nil,
		TableInfoVersion: 0,
		ReplicaID:        0,
		Columns: []*model.Column{
			{Value: "v11"},
			{Value: 111},
			{Value: false},
			{Value: "v12"},
			{Value: 222},
			{Value: true},
		},
		PreColumns:          nil,
		IndexColumns:        nil,
		ApproximateDataSize: 0,
		SplitTxn:            false,
		ReplicatingTs:       0,
	}
	e2 := &model.RowChangedEvent{
		StartTs:  1,
		CommitTs: 1,
		RowID:    1,
		Table: &model.TableName{
			Schema:      "db1",
			Table:       "tbl1",
			TableID:     1,
			IsPartition: false,
		},
		ColInfos:         nil,
		TableInfoVersion: 0,
		ReplicaID:        0,
		Columns: []*model.Column{
			{Value: "v21"},
			{Value: 111},
			{Value: false},
			{Value: "v22"},
			{Value: 222},
			{Value: true},
		},
		PreColumns:          nil,
		IndexColumns:        nil,
		ApproximateDataSize: 0,
		SplitTxn:            false,
		ReplicatingTs:       0,
	}
	return []*model.RowChangedEvent{e1, e2}
}
