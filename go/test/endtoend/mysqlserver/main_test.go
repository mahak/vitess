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

package mysqlserver

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"testing"

	"vitess.io/vitess/go/mysql"
	"vitess.io/vitess/go/test/endtoend/cluster"
	"vitess.io/vitess/go/vt/utils"
)

var (
	clusterInstance       *cluster.LocalProcessCluster
	vtParams              mysql.ConnParams
	hostname              = "localhost"
	keyspaceName          = "test_keyspace"
	tableACLConfig        = "/table_acl_config.json"
	mysqlAuthServerStatic = "/mysql_auth_server_static.json"
	cell                  = "zone1"
	sqlSchema             = `create table vt_insert_test (
		id bigint auto_increment,
		msg varchar(64),
		keyspace_id bigint(20) unsigned NOT NULL,
		data longblob,
		primary key (id)
		) Engine=InnoDB;
	create table vt_partition_test (
		c1 int NOT NULL,
		logdata BLOB NOT NULL,
		created DATETIME NOT NULL,
		PRIMARY KEY(c1, created)
		)
		PARTITION BY HASH( TO_DAYS(created) )
		PARTITIONS 10;
`
	createProcSQL = `
CREATE PROCEDURE testing()
BEGIN
	delete from vt_insert_test;
	delete from vt_partition_test;
END;
`
)

func TestMain(m *testing.M) {
	flag.Parse()

	// setting grpc max size
	if os.Getenv("grpc-max-message-size") == "" {
		os.Setenv("grpc-max-message-size", strconv.FormatInt(16*1024*1024, 10))
	}

	exitcode, err := func() (int, error) {
		clusterInstance = cluster.NewCluster(cell, hostname)
		defer clusterInstance.Teardown()

		// Start topo server
		if err := clusterInstance.StartTopo(); err != nil {
			return 1, err
		}

		// create acl config
		ACLConfig := `{
			"table_groups": [
				{
					"table_names_or_prefixes": ["vt_insert_test", "vt_partition_test", "dual"],
					"readers": ["vtgate client 1"],
					"writers": ["vtgate client 1"],
					"admins": ["vtgate client 1"]
				}
			]
		}`
		if err := createConfig(tableACLConfig, ACLConfig); err != nil {
			return 1, err
		}

		// create auth server config
		SQLConfig := `{
			"testuser1": {
				"Password": "testpassword1",
				"UserData": "vtgate client 1"
			},
			"testuser2": {
				"Password": "testpassword2",
				"UserData": "vtgate client 2"
			}
		}`
		if err := createConfig(mysqlAuthServerStatic, SQLConfig); err != nil {
			return 1, err
		}

		clusterInstance.VtGateExtraArgs = []string{
			"--vschema_ddl_authorized_users=%",
			utils.GetFlagVariantForTests("--mysql-server-query-timeout"), "1s",
			utils.GetFlagVariantForTests("--mysql-auth-server-impl"), "static",
			"--mysql_auth_server_static_file", clusterInstance.TmpDirectory + mysqlAuthServerStatic,
			utils.GetFlagVariantForTests("--mysql-server-version"), "8.0.16-7",
			utils.GetFlagVariantForTests("--warn-sharded-only") + "=true",
		}

		clusterInstance.VtTabletExtraArgs = []string{
			"--table-acl-config", clusterInstance.TmpDirectory + tableACLConfig,
			"--queryserver-config-strict-table-acl",
		}

		// Start keyspace
		keyspace := &cluster.Keyspace{
			Name:      keyspaceName,
			SchemaSQL: sqlSchema,
		}
		if err := clusterInstance.StartUnshardedKeyspace(*keyspace, 1, false); err != nil {
			return 1, err
		}

		// Start vtgate
		if err := clusterInstance.StartVtgate(); err != nil {
			return 1, err
		}

		vtParams = mysql.ConnParams{
			Host:  clusterInstance.Hostname,
			Port:  clusterInstance.VtgateMySQLPort,
			Uname: "testuser1",
			Pass:  "testpassword1",
		}

		primaryTabletProcess := clusterInstance.Keyspaces[0].Shards[0].PrimaryTablet().VttabletProcess
		if _, err := primaryTabletProcess.QueryTablet(createProcSQL, keyspaceName, true); err != nil {
			return 1, err
		}

		return m.Run(), nil
	}()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		os.Exit(exitcode)
	}

}
