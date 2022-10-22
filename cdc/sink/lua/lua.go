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

	addTableFn    tableFunc
	removeTableFn tableFunc
	rowChangedFn  rowChangedFunc
	ddlFn         ddlFunc
	cancel        context.CancelFunc
}

type tableFunc func(ctx context.Context, tableID int64) error
type rowChangedFunc func(context.Context, ...*model.RowChangedEvent) error
type ddlFunc func(context.Context, *model.DDLEvent) error

func NewLuaSink(ctx context.Context, sinkURI *url.URL,
	replicaConfig *config.ReplicaConfig, opts map[string]string, errCh chan error,
) (*luaSink, error) {
	changefeedID := contextutil.ChangefeedIDFromCtx(ctx)

	L := lua.NewState()
	defer L.Close() //you need to serialize and deserialize rows data, send it to lua and process.
	// for c++, u need to define 
	
	the header and give some guidelines for users to write code.
	ctx, cancel := context.WithCancel(context.Background())
	L.SetContext(ctx)

	L.DoString(`
    function coro()
          local i = 0
          while true do
            coroutine.yield(i)
                i = i+1
          end
          return i
    end
`)
	co, cocancel := L.NewThread()
	defer cocancel()
	fn := L.GetGlobal("coro").(*LFunction)

	_, err, values := L.Resume(co, fn) // err is nil

	cancel() // cancel the parent context

	_, err, values = L.Resume(co, fn) // err is NOT nil : child context was canceled

	s := &luaSink{id: changefeedID, cancel: cancel}

	return s, nil
}

func (l *luaSink) AddTable(tableID model.TableID) error {
	return l.addTableFn(context.Background(), tableID)
}

func (l *luaSink) EmitRowChangedEvents(ctx context.Context, rows ...*model.RowChangedEvent) error {
	return l.rowChangedFn(ctx, rows...)
}

func (l *luaSink) EmitDDLEvent(ctx context.Context, ddl *model.DDLEvent) error {
	return l.ddlFn(ctx, ddl)
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
	return l.removeTableFn(ctx, tableID)
}

func (l *luaSink) Close(ctx context.Context) error {
	l.cancel()
	return nil
}
