package rpc

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/stretchr/testify/require"
	"net/rpc/jsonrpc"
	"net/url"
	"testing"
	"time"
)

func TestRPCClient(t *testing.T) {
	rpcClient, err := jsonrpc.Dial("tcp", "localhost:5006")
	require.NoError(t, err)
	args := PluginRequest{"2.0", "d87198f0-af92-49f8-9a7d-ab8bed5c4d17"}
	var reply string

	err = rpcClient.Call("SinkSyncService.SinkSync", args, &reply)
	require.NoError(t, err)
	log.Printf("Response: %d", reply)
}

func TestInitRPCPluginSink(t *testing.T) {
	sinkURI, err := url.Parse("localhost:5006")
	require.NoError(t, err)
	sink, err := NewRPCPluginSink(context.Background(), sinkURI, nil, nil, nil)
	require.NoError(t, err)
	err = sink.AddTable(time.Now().Unix())
	require.NoError(t, err)
}

func TestRPCPluginRowEvent(t *testing.T) {
	ctx := context.Background()
	sinkURI, err := url.Parse("localhost:5006")
	require.NoError(t, err)
	sink, err := NewRPCPluginSink(ctx, sinkURI, nil, nil, nil)
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
