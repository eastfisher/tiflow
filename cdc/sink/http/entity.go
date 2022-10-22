package http

import (
	"fmt"
)

// HTTPRespCodeSuccess http plugin response code success
const HTTPRespCodeSuccess = 200

// RespStatus response status
type RespStatus struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

// AddTableReq add table request
type AddTableReq struct {
	TableID int64 `json:"table_id"`
}

// RemoveTableReq remove table request
type RemoveTableReq struct {
	TableID int64 `json:"table_id"`
}

type CommonExecResp struct {
	RespStatus
}

type PluginRequest struct {
	Operation string      `json:"operation"`
	Data      interface{} `json:"data"`
}

func (r RespStatus) Error() string {
	return fmt.Sprintf("code: %d, msg: %s, cause: %s", r.Code, r.Message, r.Cause)
}
