package udflibc

import (
	"context"
	"net/url"

	"plugin"

	"github.com/pingcap/tiflow/cdc/contextutil"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/config"
)

type udflibcSink struct {
	id model.ChangeFeedID

	addTableFn    tableFunc
	removeTableFn tableFunc
	rowChangedFn  rowChangedFunc
	ddlFn         ddlFunc
}

type tableFunc func(ctx context.Context, tableID int64) error
type rowChangedFunc func(context.Context, ...*model.RowChangedEvent) error
type ddlFunc func(context.Context, *model.DDLEvent) error

func NewUdflibcSink(ctx context.Context, sinkURI *url.URL,
	replicaConfig *config.ReplicaConfig, opts map[string]string, errCh chan error,
) (*udflibcSink, error) {
	changefeedID := contextutil.ChangefeedIDFromCtx(ctx)

	p := opts["udflibcpath"]
	plugin, err := plugin.Open(p)
	if err != nil {
		return nil, err
	}

	s := &udflibcSink{
		id: changefeedID,
	}

	if addTableFn, err := plugin.Lookup("AddTable"); err == nil {
		s.addTableFn = addTableFn.(func(ctx context.Context, tableID int64) error)
	}
	if removeTableFn, err := plugin.Lookup("RemoveTable"); err == nil {
		s.removeTableFn = removeTableFn.(func(ctx context.Context, tableID int64) error)
	}
	if rowChangedFn, err := plugin.Lookup("RowChanged"); err == nil {
		s.rowChangedFn = rowChangedFn.(func(context.Context, ...*model.RowChangedEvent) error)
	}
	if ddlFn, err := plugin.Lookup("DDL"); err == nil {
		s.ddlFn = ddlFn.(func(context.Context, *model.DDLEvent) error)
	}
	return s, nil
}

func (l *udflibcSink) AddTable(tableID model.TableID) error {
	return l.addTableFn(context.Background(), tableID)
}

func (l *udflibcSink) EmitRowChangedEvents(ctx context.Context, rows ...*model.RowChangedEvent) error {
	return l.rowChangedFn(ctx, rows...)
}

func (l *udflibcSink) EmitDDLEvent(ctx context.Context, ddl *model.DDLEvent) error {
	return l.ddlFn(ctx, ddl)
}

func (l *udflibcSink) FlushRowChangedEvents(ctx context.Context, tableID model.TableID, resolved model.ResolvedTs) (model.ResolvedTs, error) {
	// Do nothing now, implement this function when we need it.
	return resolved, nil
}

func (l *udflibcSink) EmitCheckpointTs(ctx context.Context, ts uint64, tables []model.TableName) error {
	// Do nothing now, implement this function when we need it.
	return nil
}

func (l *udflibcSink) RemoveTable(ctx context.Context, tableID model.TableID) error {
	return l.removeTableFn(ctx, tableID)
}

func (l *udflibcSink) Close(ctx context.Context) error {
	return nil
}
