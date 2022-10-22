package wasm

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalCommonExecResp(t *testing.T) {
	s := []byte(`{"code":200,"message":"success","cause":null}`)
	var resp CommonExecResp
	err := WasmUnmarshal(s, &resp)
	assert.NoError(t, err)
	assert.Equal(t, int32(RespCodeSuccess), resp.Code)
}

func TestBase64(t *testing.T) {
	src := []byte("hello")
	ret := base64.URLEncoding.EncodeToString(src)
	t.Logf("ret: %v", ret)
	ret2, err := base64.URLEncoding.DecodeString(ret)
	if err != nil {
		t.FailNow()
	}
	t.Logf("ret2: %s", ret2)
}
