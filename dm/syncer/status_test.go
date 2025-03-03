// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package syncer

import (
	"sync"

	"github.com/go-mysql-org/go-mysql/mysql"
	. "github.com/pingcap/check"
	"github.com/pingcap/tidb/parser/model"
	"github.com/pingcap/tidb/util/filter"
	"go.uber.org/zap"

	"github.com/pingcap/tiflow/dm/dm/config"
	"github.com/pingcap/tiflow/dm/dm/pb"
	"github.com/pingcap/tiflow/dm/pkg/binlog"
	tcontext "github.com/pingcap/tiflow/dm/pkg/context"
	"github.com/pingcap/tiflow/dm/pkg/log"
	"github.com/pingcap/tiflow/dm/syncer/shardddl"
)

var _ = Suite(&statusSuite{})

type statusSuite struct{}

func (t *statusSuite) TestStatusRace(c *C) {
	s := &Syncer{}

	l := log.With(zap.String("unit test", "TestStatueRace"))
	s.tctx = tcontext.Background().WithLogger(l)
	s.cfg = &config.SubTaskConfig{}
	s.checkpoint = &mockCheckpoint{}
	s.pessimist = shardddl.NewPessimist(&l, nil, "", "")
	s.optimist = shardddl.NewOptimist(&l, nil, "", "")

	sourceStatus := &binlog.SourceStatus{
		Location: binlog.Location{
			Position: mysql.Position{
				Name: "mysql-bin.000123",
				Pos:  223,
			},
		},
		Binlogs: binlog.FileSizes(nil),
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ret := s.Status(sourceStatus)
			status := ret.(*pb.SyncStatus)
			c.Assert(status.MasterBinlog, Equals, "(mysql-bin.000123, 223)")
			c.Assert(status.SyncerBinlog, Equals, "(mysql-bin.000123, 123)")
		}()
	}
	wg.Wait()
}

type mockCheckpoint struct {
	CheckPoint
}

func (*mockCheckpoint) FlushedGlobalPoint() binlog.Location {
	return binlog.Location{
		Position: mysql.Position{
			Name: "mysql-bin.000123",
			Pos:  123,
		},
	}
}

func (*mockCheckpoint) SaveTablePoint(_ *filter.Table, _ binlog.Location, _ *model.TableInfo) {}
