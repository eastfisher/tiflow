package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pingcap/log"
	"github.com/pingcap/tiflow/cdc/contextutil"
	"github.com/pingcap/tiflow/cdc/model"
	"github.com/pingcap/tiflow/pkg/config"
)

type httpPluginSink struct {
	id model.ChangeFeedID

	httpClient http.Client
	httpURL    string
}

func NewHTTPPluginSink(ctx context.Context, sinkURI *url.URL,
	replicaConfig *config.ReplicaConfig, opts map[string]string, errCh chan error,
) (*httpPluginSink, error) {
	changeFeedID := contextutil.ChangefeedIDFromCtx(ctx)
	log.Info("init http plugin sink")
	return &httpPluginSink{
		id:         changeFeedID,
		httpClient: http.Client{Timeout: time.Duration(1) * time.Second},
		httpURL:    opts["httpurl"], // sinkURI.String(),
	}, nil
}

func (hp *httpPluginSink) callHTTPRequest(ctx context.Context, data []byte) ([]byte, error) {
	postData := bytes.NewBuffer(data)
	req, err := http.NewRequestWithContext(ctx, "POST", hp.httpURL, postData)
	if err != nil {
		log.Error("http plugin sink NewRequestWithContext" + err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", `application/json`)

	resp, err := hp.httpClient.Do(req)
	if err != nil {
		log.Error("http plugin sink httpClient.Do" + err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("http plugin sink io read" + err.Error())

		return nil, err
	}

	return body, nil
	// return nil, nil
}

func (hp *httpPluginSink) AddTable(tableID model.TableID) error {
	ctx := context.Background()

	req := AddTableReq{TableID: tableID}
	return hp.doExecHTTPCall(ctx, "sink_add_table", req)
}

func (hp *httpPluginSink) EmitRowChangedEvents(ctx context.Context, rows ...*model.RowChangedEvent) error {
	return hp.doExecHTTPCall(ctx, "sink_emit_row_changed_events", rows)
}

func (hp *httpPluginSink) EmitDDLEvent(ctx context.Context, ddl *model.DDLEvent) error {
	return hp.doExecHTTPCall(ctx, "sink_emit_ddl_event", ddl)
}

func (hp *httpPluginSink) FlushRowChangedEvents(ctx context.Context, tableID model.TableID, resolved model.ResolvedTs) (model.ResolvedTs, error) {
	// Do nothing now, implement this function when we need it.
	return resolved, nil
}

func (hp *httpPluginSink) EmitCheckpointTs(ctx context.Context, ts uint64, tables []model.TableName) error {
	// Do nothing now, implement this function when we need it.
	return nil
}

func (hp *httpPluginSink) RemoveTable(ctx context.Context, tableID model.TableID) error {
	req := RemoveTableReq{TableID: tableID}
	return hp.doExecHTTPCall(ctx, "sink_remove_table", req)
}

func (hp *httpPluginSink) Close(ctx context.Context) error {
	return nil
}

func (hp *httpPluginSink) doExecHTTPCall(ctx context.Context, operation string, req interface{}) error {
	reqData := &PluginRequest{
		Operation: operation,
		Data:      req,
	}
	reqBytes, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	// // use http client
	hp.callHTTPRequest(ctx, reqBytes)
	// if err != nil {
	// 	return err
	// }

	// var resp CommonExecResp
	// if err := json.Unmarshal(respBytes, &resp); err != nil {
	// 	return err
	// }
	// if resp.Code != HTTPRespCodeSuccess {
	// 	return resp
	// }
	return nil
}
