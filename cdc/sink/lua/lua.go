package lua

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cjoudrey/gluahttp"
	"github.com/pingcap/log"
	"github.com/pingcap/tiflow/cdc/contextutil"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/config"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
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
	changeFeedID := contextutil.ChangefeedIDFromCtx(ctx)

	s := &luaSink{id: changeFeedID}
	s.addTableFn = opts["luaAddTable"]       // file path
	s.removeTableFn = opts["luaRemoveTable"] // remove table file path
	s.rowChangedFn = opts["luaRowChanged"]
	s.ddlFn = opts["luaDDL"]

	return s, nil
}

func (l *luaSink) AddTable(tableID model.TableID) error {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)

	err := L.DoFile(l.addTableFn)
	if err != nil {
		log.Error("DoFile failed", zap.Error(err))
		return err
	}
	err = l.doExecLuaCall(L, "addTable", tableID)
	if err != nil {
		log.Error("doExecLuaCall failed", zap.Error(err))
		return err
	}
	return nil
}

func (l *luaSink) EmitRowChangedEvents(ctx context.Context, rows ...*model.RowChangedEvent) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.rowChangedFn) // FIXME: rows is not passed to function!
	if err != nil {
		log.Error("DoFile failed, err", zap.Error(err))
		return err
	}
	data, err := json.Marshal(rows)
	if err != nil {
		log.Error("marshal failed", zap.Error(err))
		return err
	}
	err = l.doExecLuaCall2(L, "rowChanged", string(data))
	if err != nil {
		log.Error("doExecLuaCall2 failed", zap.Error(err))
		return err
	}
	return err
}

func (l *luaSink) EmitDDLEvent(ctx context.Context, ddl *model.DDLEvent) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.ddlFn)
	if err != nil {
		log.Error("DoFile failed", zap.Error(err))
		return err
	}
	data, err := json.Marshal(ddl)
	if err != nil {
		log.Error("marshal failed", zap.Error(err))
		return err
	}
	err = l.doExecLuaCall2(L, "ddl", string(data))
	if err != nil {
		log.Error("doExecLuaCall2 failed", zap.Error(err))
		return err
	}
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

func (l *luaSink) RemoveTable(tableID model.TableID) error {
	L := lua.NewState()
	defer L.Close()

	err := L.DoFile(l.removeTableFn)
	if err != nil {
		log.Error("DoFile failed", zap.Error(err))
		return err
	}
	err = l.doExecLuaCall(L, "removeTable", tableID)
	if err != nil {
		log.Error("doExecLuaCall failed", zap.Error(err))
		return err
	}
	return err
}

func (l *luaSink) Close(ctx context.Context) error {
	return nil
}

func (l *luaSink) doExecLuaCall(L *lua.LState, op string, data any) error {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(op),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(data.(int64))); err != nil {
		log.Error("callByParam failed", zap.Error(err))
		return err
	}
	ret, _ := L.Get(-1).(lua.LString)
	log.Info("doExecLuaCall success", zap.Any("ret", ret))
	return nil
}

func (l *luaSink) doExecLuaCall2(L *lua.LState, op string, data string) error {
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal(op),
		NRet:    1,
		Protect: true,
	}, lua.LString(data)); err != nil {
		log.Error("callByParam failed", zap.Error(err))
		return err
	}
	ret, _ := L.Get(-1).(lua.LString)
	log.Info("doExecLuaCall2 success", zap.Any("ret", ret))
	return nil
}
