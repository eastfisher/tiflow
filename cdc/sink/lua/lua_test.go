package lua

import (
	"context"
	"testing"

	"github.com/pingcap/tiflow/cdc/model"
	"github.com/stretchr/testify/require"
)

func TestLuaPluginSink(t *testing.T) {
	ctx := context.Background()
	opts := map[string]string{
		"luaAddTable":    "scripts/addTable.lua",
		"luaRemoveTable": "scripts/removeTable.lua",
	}
	sink, err := NewLuaSink(ctx, nil, nil, opts, nil)
	require.NoError(t, err)
	err = sink.AddTable(100)
	require.NoError(t, err)

	err = sink.RemoveTable(100)
	require.NoError(t, err)
}

func TestLuaRowEvent(t *testing.T) {
	ctx := context.Background()
	opts := map[string]string{
		"luaAddTable":    "scripts/addTable.lua",
		"luaRemoveTable": "scripts/removeTable.lua",
		"luaRowChanged":  "scripts/rowChanged.lua",
		"luaDDL":         "scripts/ddl.lua",
	}
	sink, err := NewLuaSink(ctx, nil, nil, opts, nil)
	require.NoError(t, err)
	err = sink.AddTable(100)
	require.NoError(t, err)

	events := getTestRowEvents()
	err = sink.EmitRowChangedEvents(ctx, events...)
	require.NoError(t, err)

	ddl := &model.DDLEvent{
		StartTs:  1,
		CommitTs: 1,
		Query:    "select * from 1",
		Type:     0,
		Done:     false,
	}
	err = sink.EmitDDLEvent(ctx, ddl)
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