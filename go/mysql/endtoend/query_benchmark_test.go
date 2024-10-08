/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package endtoend

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"vitess.io/vitess/go/mysql"
	vttestpb "vitess.io/vitess/go/vt/proto/vttest"
	"vitess.io/vitess/go/vt/vttest"
)

// This file contains various long-running tests for mysql.

// BenchmarkWithRealDatabase runs a real MySQL database, and runs all kinds
// of benchmarks on it. To minimize overhead, we only run one database, and
// run all the benchmarks on it.
func BenchmarkWithRealDatabase(b *testing.B) {
	// Launch MySQL.
	// We need a Keyspace in the topology, so the DbName is set.
	// We need a Shard too, so the database 'vttest' is created.
	cfg := vttest.Config{
		Topology: &vttestpb.VTTestTopology{
			Keyspaces: []*vttestpb.Keyspace{
				{
					Name: "vttest",
					Shards: []*vttestpb.Shard{
						{
							Name:           "0",
							DbNameOverride: "vttest",
						},
					},
				},
			},
		},
		OnlyMySQL: true,
	}
	if err := cfg.InitSchemas("vttest", "create table a(id int, name varchar(128), primary key(id))", nil); err != nil {
		b.Fatalf("InitSchemas failed: %v\n", err)
	}
	defer os.RemoveAll(cfg.SchemaDir)
	cluster := vttest.LocalCluster{
		Config: cfg,
	}
	if err := cluster.Setup(); err != nil {
		b.Fatalf("could not launch mysql: %v\n", err)
	}
	defer cluster.TearDown()
	params := cluster.MySQLConnParams()

	b.Run("Inserts", func(b *testing.B) {
		benchmarkInserts(b, &params)
	})
	b.Run("ParallelReads", func(b *testing.B) {
		benchmarkParallelReads(b, &params, 10)
	})
}

func benchmarkInserts(b *testing.B, params *mysql.ConnParams) {
	// Connect.
	ctx := context.Background()
	conn, err := mysql.Connect(ctx, params)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	// Delete what we may already have in the database.
	if _, err := conn.ExecuteFetch("delete from a", 0, false); err != nil {
		b.Fatalf("delete failed: %v", err)
	}

	// Now reset timer.
	b.ResetTimer()

	// Do the insert.
	for i := range b.N {
		_, err := conn.ExecuteFetch(fmt.Sprintf("insert into a(id, name) values(%v, 'nice name %v')", i, i), 0, false)
		if err != nil {
			b.Fatalf("ExecuteFetch(%v) failed: %v", i, err)
		}
	}
}

func benchmarkParallelReads(b *testing.B, params *mysql.ConnParams, parallelCount int) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	for i := range parallelCount {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			conn, err := mysql.Connect(ctx, params)
			if err != nil {
				b.Error(err)
			}

			for j := range b.N {
				if _, err := conn.ExecuteFetch("select * from a", 20000, true); err != nil {
					b.Errorf("ExecuteFetch(%v, %v) failed: %v", i, j, err)
				}
			}
			conn.Close()
		}(i)
	}
	wg.Wait()
}

func BenchmarkSetVarsWithQueryHints(b *testing.B) {
	ctx := context.Background()
	conn, err := mysql.Connect(ctx, &connParams)
	if err != nil {
		b.Fatal(err)
	}

	_, err = conn.ExecuteFetch("create table t(id int primary key, name varchar(100))", 1, false)
	require.NoError(b, err)

	defer func() {
		_, err = conn.ExecuteFetch("drop table t", 1, false)
		require.NoError(b, err)
	}()

	for _, sleepDuration := range []time.Duration{0, 1 * time.Millisecond} {
		b.Run(fmt.Sprintf("Sleep %d ms", sleepDuration/time.Millisecond), func(b *testing.B) {
			for i := range b.N {
				_, err := conn.ExecuteFetch(fmt.Sprintf("insert /*+ SET_VAR(sql_mode = ' ') SET_VAR(sql_safe_updates = 0) */ into t(id) values (%d)", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, err = conn.ExecuteFetch(fmt.Sprintf("select /*+ SET_VAR(sql_mode = ' ') SET_VAR(sql_safe_updates = 0) */ * from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, err = conn.ExecuteFetch(fmt.Sprintf("update /*+ SET_VAR(sql_mode = ' ') SET_VAR(sql_safe_updates = 0) */ t set name = 'foo' where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, err = conn.ExecuteFetch(fmt.Sprintf("delete /*+ SET_VAR(sql_mode = ' ') SET_VAR(sql_safe_updates = 0) */ from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				time.Sleep(sleepDuration)
			}
		})
	}
}

func BenchmarkSetVarsMultipleSets(b *testing.B) {
	ctx := context.Background()
	conn, err := mysql.Connect(ctx, &connParams)
	if err != nil {
		b.Fatal(err)
	}

	_, err = conn.ExecuteFetch("create table t(id int primary key, name varchar(100))", 1, false)
	require.NoError(b, err)

	defer func() {
		_, err = conn.ExecuteFetch("drop table t", 1, false)
		require.NoError(b, err)
	}()

	setFunc := func() {
		_, err = conn.ExecuteFetch("set sql_mode = '', sql_safe_updates = 0;", 1, false)
		if err != nil {
			b.Fatal(err)
		}
	}

	for _, sleepDuration := range []time.Duration{0, 1 * time.Millisecond} {
		b.Run(fmt.Sprintf("Sleep %d ms", sleepDuration/time.Millisecond), func(b *testing.B) {
			for i := range b.N {
				setFunc()

				_, err = conn.ExecuteFetch(fmt.Sprintf("insert into t(id) values (%d)", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				setFunc()

				_, err = conn.ExecuteFetch(fmt.Sprintf("select * from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				setFunc()

				_, err = conn.ExecuteFetch(fmt.Sprintf("update t set name = 'foo' where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				setFunc()

				_, err = conn.ExecuteFetch(fmt.Sprintf("delete from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				time.Sleep(sleepDuration)
			}
		})
	}
}

func BenchmarkSetVarsMultipleSetsInSameStmt(b *testing.B) {
	ctx := context.Background()
	conn, err := mysql.Connect(ctx, &connParams)
	if err != nil {
		b.Fatal(err)
	}

	_, err = conn.ExecuteFetch("create table t(id int primary key, name varchar(100))", 1, false)
	require.NoError(b, err)

	defer func() {
		_, err = conn.ExecuteFetch("drop table t", 1, false)
		require.NoError(b, err)
	}()

	for _, sleepDuration := range []time.Duration{0, 1 * time.Millisecond} {
		b.Run(fmt.Sprintf("Sleep %d ms", sleepDuration/time.Millisecond), func(b *testing.B) {
			for i := range b.N {
				_, _, err := conn.ExecuteFetchMulti(fmt.Sprintf("set sql_mode = '', sql_safe_updates = 0 ; insert into t(id) values (%d)", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				_, _, _, err = conn.ReadQueryResult(1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, _, err = conn.ExecuteFetchMulti(fmt.Sprintf("set sql_mode = '', sql_safe_updates = 0 ; select * from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				_, _, _, err = conn.ReadQueryResult(1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, _, err = conn.ExecuteFetchMulti(fmt.Sprintf("set sql_mode = '', sql_safe_updates = 0 ; update t set name = 'foo' where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				_, _, _, err = conn.ReadQueryResult(1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, _, err = conn.ExecuteFetchMulti(fmt.Sprintf("set sql_mode = '', sql_safe_updates = 0 ; delete from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				_, _, _, err = conn.ReadQueryResult(1, false)
				if err != nil {
					b.Fatal(err)
				}
				time.Sleep(sleepDuration)
			}
		})
	}
}

func BenchmarkSetVarsSingleSet(b *testing.B) {
	ctx := context.Background()
	conn, err := mysql.Connect(ctx, &connParams)
	if err != nil {
		b.Fatal(err)
	}

	_, err = conn.ExecuteFetch("set sql_mode = '', sql_safe_updates = 0", 1, false)
	require.NoError(b, err)

	_, err = conn.ExecuteFetch("create table t(id int primary key, name varchar(100))", 1, false)
	require.NoError(b, err)

	defer func() {
		_, err = conn.ExecuteFetch("drop table t", 1, false)
		require.NoError(b, err)
	}()

	for _, sleepDuration := range []time.Duration{0, 1 * time.Millisecond} {
		b.Run(fmt.Sprintf("Sleep %d ms", sleepDuration/time.Millisecond), func(b *testing.B) {
			for i := range b.N {
				_, err = conn.ExecuteFetch(fmt.Sprintf("insert into t(id) values (%d)", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, err = conn.ExecuteFetch(fmt.Sprintf("select * from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, err = conn.ExecuteFetch(fmt.Sprintf("update t set name = 'foo' where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}

				_, err = conn.ExecuteFetch(fmt.Sprintf("delete from t where id = %d", i), 1, false)
				if err != nil {
					b.Fatal(err)
				}
				time.Sleep(sleepDuration)
			}
		})
	}

}
