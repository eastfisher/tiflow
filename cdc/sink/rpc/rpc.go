package rpc

import (
	"context"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/url"

	"github.com/labstack/gommon/log"
	"github.com/pingcap/tiflow/cdc/contextutil"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/config"
)

type rpcPluginSink struct {
	id model.ChangeFeedID

	rpcClient *rpc.Client
	httpURL   string
}

func NewRPCPluginSink(ctx context.Context, sinkURI *url.URL,
	replicaConfig *config.ReplicaConfig, opts map[string]string, errCh chan error,
) (*rpcPluginSink, error) {
	changeFeedID := contextutil.ChangefeedIDFromCtx(ctx)

	rpcClient, err := jsonrpc.Dial("tcp", opts["httpurl"])
	if err != nil {
		log.Fatalf("dial sink URI failed, err:%v", err)
		return nil, err
	}
	return &rpcPluginSink{
		id:        changeFeedID,
		rpcClient: rpcClient,
	}, nil
}

func (hp *rpcPluginSink) callRPC(ctx context.Context, pr *PluginRequest) (*RespStatus, error) {
	reply := &RespStatus{}
	err := hp.rpcClient.Call("SinkSyncService.SinkSync", &pr, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (hp *rpcPluginSink) AddTable(tableID model.TableID) error {
	ctx := context.Background()

	req := AddTableReq{TableID: tableID}
	return hp.doExecRPCCall(ctx, "sink_add_table", req)
}

func (hp *rpcPluginSink) EmitRowChangedEvents(ctx context.Context, rows ...*model.RowChangedEvent) error {
	return hp.doExecRPCCall(ctx, "sink_emit_row_changed_events", rows)
}

func (hp *rpcPluginSink) EmitDDLEvent(ctx context.Context, ddl *model.DDLEvent) error {
	return hp.doExecRPCCall(ctx, "sink_emit_ddl_event", ddl)
}

func (hp *rpcPluginSink) FlushRowChangedEvents(ctx context.Context, tableID model.TableID, resolved model.ResolvedTs) (model.ResolvedTs, error) {
	// Do nothing now, implement this function when we need it.
	return resolved, nil
}

func (hp *rpcPluginSink) EmitCheckpointTs(ctx context.Context, ts uint64, tables []model.TableName) error {
	// Do nothing now, implement this function when we need it.
	return nil
}

func (hp *rpcPluginSink) RemoveTable(ctx context.Context, tableID model.TableID) error {
	req := RemoveTableReq{TableID: tableID}
	return hp.doExecRPCCall(ctx, "sink_remove_table", req)
}

func (hp *rpcPluginSink) Close(ctx context.Context) error {
	return nil
}

func (hp *rpcPluginSink) doExecRPCCall(ctx context.Context, operation string, req interface{}) error {
	reqData := &PluginRequest{
		Operation: operation,
		Data:      req,
	}

	// use json rpc client
	resp, err := hp.callRPC(ctx, reqData)
	if err != nil {
		return err
	}

	if resp.Code != RPCRespCodeSuccess {
		return resp
	}
	return nil
}
