package lua

import (
	"context"
	"net/url"

	"github.com/pingcap/tiflow/cdc/contextutil"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/config"
	lua "github.com/yuin/gopher-lua"
)

type luaSink struct {
	id model.ChangeFeedID

	addTableFn    string
	removeTableFn string
	rowChangedFn  string
	ddlFn         string
}

func NewLuaSink(ctx context.Context, sinkURI *url.URL,
	replicaConfig *config.ReplicaConfig, opts map[string]string, errCh chan error,
) (*luaSink, error) {
	changefeedID := contextutil.ChangefeedIDFromCtx(ctx)

	s := &luaSink{id: changefeedID}
	s.addTableFn = opts["luaAddTable"]       // file path
	s.removeTableFn = opts["luaRemoveTable"] // remove table file path
	s.rowChangedFn = opts["luaRowChanged"]
	s.ddlFn = opts["luaDDL"]

	return s, nil
}

func (l *luaSink) AddTable(tableID model.TableID) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.addTableFn)

	return err
}

func (l *luaSink) EmitRowChangedEvents(ctx context.Context, rows ...*model.RowChangedEvent) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.rowChangedFn) // FIXME: rows is not passed to function!

	return err
}

func (l *luaSink) EmitDDLEvent(ctx context.Context, ddl *model.DDLEvent) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.ddlFn)

	return err
}

func (l *luaSink) FlushRowChangedEvents(ctx context.Context, tableID model.TableID, resolved model.ResolvedTs) (model.ResolvedTs, error) {
	// Do nothing now, implement this function when we need it.
	return resolved, nil
}

func (l *luaSink) EmitCheckpointTs(ctx context.Context, ts uint64, tables []model.TableName) error {
	// Do nothing now, implement this function when we need it.
	return nil
}

func (l *luaSink) RemoveTable(ctx context.Context, tableID model.TableID) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.removeTableFn)

	return err
}

func (l *luaSink) Close(ctx context.Context) error {
	return nil
}
