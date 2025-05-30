//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file contains the service VtTablet exposes for queries.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.3
// source: queryservice.proto

package queryservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
	binlogdata "vitess.io/vitess/go/vt/proto/binlogdata"
	query "vitess.io/vitess/go/vt/proto/query"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_queryservice_proto protoreflect.FileDescriptor

const file_queryservice_proto_rawDesc = "" +
	"\n" +
	"\x12queryservice.proto\x12\fqueryservice\x1a\vquery.proto\x1a\x10binlogdata.proto2\x95\x12\n" +
	"\x05Query\x12:\n" +
	"\aExecute\x12\x15.query.ExecuteRequest\x1a\x16.query.ExecuteResponse\"\x00\x12N\n" +
	"\rStreamExecute\x12\x1b.query.StreamExecuteRequest\x1a\x1c.query.StreamExecuteResponse\"\x000\x01\x124\n" +
	"\x05Begin\x12\x13.query.BeginRequest\x1a\x14.query.BeginResponse\"\x00\x127\n" +
	"\x06Commit\x12\x14.query.CommitRequest\x1a\x15.query.CommitResponse\"\x00\x12=\n" +
	"\bRollback\x12\x16.query.RollbackRequest\x1a\x17.query.RollbackResponse\"\x00\x12:\n" +
	"\aPrepare\x12\x15.query.PrepareRequest\x1a\x16.query.PrepareResponse\"\x00\x12O\n" +
	"\x0eCommitPrepared\x12\x1c.query.CommitPreparedRequest\x1a\x1d.query.CommitPreparedResponse\"\x00\x12U\n" +
	"\x10RollbackPrepared\x12\x1e.query.RollbackPreparedRequest\x1a\x1f.query.RollbackPreparedResponse\"\x00\x12X\n" +
	"\x11CreateTransaction\x12\x1f.query.CreateTransactionRequest\x1a .query.CreateTransactionResponse\"\x00\x12F\n" +
	"\vStartCommit\x12\x19.query.StartCommitRequest\x1a\x1a.query.StartCommitResponse\"\x00\x12F\n" +
	"\vSetRollback\x12\x19.query.SetRollbackRequest\x1a\x1a.query.SetRollbackResponse\"\x00\x12^\n" +
	"\x13ConcludeTransaction\x12!.query.ConcludeTransactionRequest\x1a\".query.ConcludeTransactionResponse\"\x00\x12R\n" +
	"\x0fReadTransaction\x12\x1d.query.ReadTransactionRequest\x1a\x1e.query.ReadTransactionResponse\"\x00\x12g\n" +
	"\x16UnresolvedTransactions\x12$.query.UnresolvedTransactionsRequest\x1a%.query.UnresolvedTransactionsResponse\"\x00\x12I\n" +
	"\fBeginExecute\x12\x1a.query.BeginExecuteRequest\x1a\x1b.query.BeginExecuteResponse\"\x00\x12]\n" +
	"\x12BeginStreamExecute\x12 .query.BeginStreamExecuteRequest\x1a!.query.BeginStreamExecuteResponse\"\x000\x01\x12N\n" +
	"\rMessageStream\x12\x1b.query.MessageStreamRequest\x1a\x1c.query.MessageStreamResponse\"\x000\x01\x12C\n" +
	"\n" +
	"MessageAck\x12\x18.query.MessageAckRequest\x1a\x19.query.MessageAckResponse\"\x00\x12O\n" +
	"\x0eReserveExecute\x12\x1c.query.ReserveExecuteRequest\x1a\x1d.query.ReserveExecuteResponse\"\x00\x12^\n" +
	"\x13ReserveBeginExecute\x12!.query.ReserveBeginExecuteRequest\x1a\".query.ReserveBeginExecuteResponse\"\x00\x12c\n" +
	"\x14ReserveStreamExecute\x12\".query.ReserveStreamExecuteRequest\x1a#.query.ReserveStreamExecuteResponse\"\x000\x01\x12r\n" +
	"\x19ReserveBeginStreamExecute\x12'.query.ReserveBeginStreamExecuteRequest\x1a(.query.ReserveBeginStreamExecuteResponse\"\x000\x01\x12:\n" +
	"\aRelease\x12\x15.query.ReleaseRequest\x1a\x16.query.ReleaseResponse\"\x00\x12K\n" +
	"\fStreamHealth\x12\x1a.query.StreamHealthRequest\x1a\x1b.query.StreamHealthResponse\"\x000\x01\x12F\n" +
	"\aVStream\x12\x1a.binlogdata.VStreamRequest\x1a\x1b.binlogdata.VStreamResponse\"\x000\x01\x12R\n" +
	"\vVStreamRows\x12\x1e.binlogdata.VStreamRowsRequest\x1a\x1f.binlogdata.VStreamRowsResponse\"\x000\x01\x12X\n" +
	"\rVStreamTables\x12 .binlogdata.VStreamTablesRequest\x1a!.binlogdata.VStreamTablesResponse\"\x000\x01\x12[\n" +
	"\x0eVStreamResults\x12!.binlogdata.VStreamResultsRequest\x1a\".binlogdata.VStreamResultsResponse\"\x000\x01\x12B\n" +
	"\tGetSchema\x12\x17.query.GetSchemaRequest\x1a\x18.query.GetSchemaResponse\"\x000\x01B+Z)vitess.io/vitess/go/vt/proto/queryserviceb\x06proto3"

var file_queryservice_proto_goTypes = []any{
	(*query.ExecuteRequest)(nil),                    // 0: query.ExecuteRequest
	(*query.StreamExecuteRequest)(nil),              // 1: query.StreamExecuteRequest
	(*query.BeginRequest)(nil),                      // 2: query.BeginRequest
	(*query.CommitRequest)(nil),                     // 3: query.CommitRequest
	(*query.RollbackRequest)(nil),                   // 4: query.RollbackRequest
	(*query.PrepareRequest)(nil),                    // 5: query.PrepareRequest
	(*query.CommitPreparedRequest)(nil),             // 6: query.CommitPreparedRequest
	(*query.RollbackPreparedRequest)(nil),           // 7: query.RollbackPreparedRequest
	(*query.CreateTransactionRequest)(nil),          // 8: query.CreateTransactionRequest
	(*query.StartCommitRequest)(nil),                // 9: query.StartCommitRequest
	(*query.SetRollbackRequest)(nil),                // 10: query.SetRollbackRequest
	(*query.ConcludeTransactionRequest)(nil),        // 11: query.ConcludeTransactionRequest
	(*query.ReadTransactionRequest)(nil),            // 12: query.ReadTransactionRequest
	(*query.UnresolvedTransactionsRequest)(nil),     // 13: query.UnresolvedTransactionsRequest
	(*query.BeginExecuteRequest)(nil),               // 14: query.BeginExecuteRequest
	(*query.BeginStreamExecuteRequest)(nil),         // 15: query.BeginStreamExecuteRequest
	(*query.MessageStreamRequest)(nil),              // 16: query.MessageStreamRequest
	(*query.MessageAckRequest)(nil),                 // 17: query.MessageAckRequest
	(*query.ReserveExecuteRequest)(nil),             // 18: query.ReserveExecuteRequest
	(*query.ReserveBeginExecuteRequest)(nil),        // 19: query.ReserveBeginExecuteRequest
	(*query.ReserveStreamExecuteRequest)(nil),       // 20: query.ReserveStreamExecuteRequest
	(*query.ReserveBeginStreamExecuteRequest)(nil),  // 21: query.ReserveBeginStreamExecuteRequest
	(*query.ReleaseRequest)(nil),                    // 22: query.ReleaseRequest
	(*query.StreamHealthRequest)(nil),               // 23: query.StreamHealthRequest
	(*binlogdata.VStreamRequest)(nil),               // 24: binlogdata.VStreamRequest
	(*binlogdata.VStreamRowsRequest)(nil),           // 25: binlogdata.VStreamRowsRequest
	(*binlogdata.VStreamTablesRequest)(nil),         // 26: binlogdata.VStreamTablesRequest
	(*binlogdata.VStreamResultsRequest)(nil),        // 27: binlogdata.VStreamResultsRequest
	(*query.GetSchemaRequest)(nil),                  // 28: query.GetSchemaRequest
	(*query.ExecuteResponse)(nil),                   // 29: query.ExecuteResponse
	(*query.StreamExecuteResponse)(nil),             // 30: query.StreamExecuteResponse
	(*query.BeginResponse)(nil),                     // 31: query.BeginResponse
	(*query.CommitResponse)(nil),                    // 32: query.CommitResponse
	(*query.RollbackResponse)(nil),                  // 33: query.RollbackResponse
	(*query.PrepareResponse)(nil),                   // 34: query.PrepareResponse
	(*query.CommitPreparedResponse)(nil),            // 35: query.CommitPreparedResponse
	(*query.RollbackPreparedResponse)(nil),          // 36: query.RollbackPreparedResponse
	(*query.CreateTransactionResponse)(nil),         // 37: query.CreateTransactionResponse
	(*query.StartCommitResponse)(nil),               // 38: query.StartCommitResponse
	(*query.SetRollbackResponse)(nil),               // 39: query.SetRollbackResponse
	(*query.ConcludeTransactionResponse)(nil),       // 40: query.ConcludeTransactionResponse
	(*query.ReadTransactionResponse)(nil),           // 41: query.ReadTransactionResponse
	(*query.UnresolvedTransactionsResponse)(nil),    // 42: query.UnresolvedTransactionsResponse
	(*query.BeginExecuteResponse)(nil),              // 43: query.BeginExecuteResponse
	(*query.BeginStreamExecuteResponse)(nil),        // 44: query.BeginStreamExecuteResponse
	(*query.MessageStreamResponse)(nil),             // 45: query.MessageStreamResponse
	(*query.MessageAckResponse)(nil),                // 46: query.MessageAckResponse
	(*query.ReserveExecuteResponse)(nil),            // 47: query.ReserveExecuteResponse
	(*query.ReserveBeginExecuteResponse)(nil),       // 48: query.ReserveBeginExecuteResponse
	(*query.ReserveStreamExecuteResponse)(nil),      // 49: query.ReserveStreamExecuteResponse
	(*query.ReserveBeginStreamExecuteResponse)(nil), // 50: query.ReserveBeginStreamExecuteResponse
	(*query.ReleaseResponse)(nil),                   // 51: query.ReleaseResponse
	(*query.StreamHealthResponse)(nil),              // 52: query.StreamHealthResponse
	(*binlogdata.VStreamResponse)(nil),              // 53: binlogdata.VStreamResponse
	(*binlogdata.VStreamRowsResponse)(nil),          // 54: binlogdata.VStreamRowsResponse
	(*binlogdata.VStreamTablesResponse)(nil),        // 55: binlogdata.VStreamTablesResponse
	(*binlogdata.VStreamResultsResponse)(nil),       // 56: binlogdata.VStreamResultsResponse
	(*query.GetSchemaResponse)(nil),                 // 57: query.GetSchemaResponse
}
var file_queryservice_proto_depIdxs = []int32{
	0,  // 0: queryservice.Query.Execute:input_type -> query.ExecuteRequest
	1,  // 1: queryservice.Query.StreamExecute:input_type -> query.StreamExecuteRequest
	2,  // 2: queryservice.Query.Begin:input_type -> query.BeginRequest
	3,  // 3: queryservice.Query.Commit:input_type -> query.CommitRequest
	4,  // 4: queryservice.Query.Rollback:input_type -> query.RollbackRequest
	5,  // 5: queryservice.Query.Prepare:input_type -> query.PrepareRequest
	6,  // 6: queryservice.Query.CommitPrepared:input_type -> query.CommitPreparedRequest
	7,  // 7: queryservice.Query.RollbackPrepared:input_type -> query.RollbackPreparedRequest
	8,  // 8: queryservice.Query.CreateTransaction:input_type -> query.CreateTransactionRequest
	9,  // 9: queryservice.Query.StartCommit:input_type -> query.StartCommitRequest
	10, // 10: queryservice.Query.SetRollback:input_type -> query.SetRollbackRequest
	11, // 11: queryservice.Query.ConcludeTransaction:input_type -> query.ConcludeTransactionRequest
	12, // 12: queryservice.Query.ReadTransaction:input_type -> query.ReadTransactionRequest
	13, // 13: queryservice.Query.UnresolvedTransactions:input_type -> query.UnresolvedTransactionsRequest
	14, // 14: queryservice.Query.BeginExecute:input_type -> query.BeginExecuteRequest
	15, // 15: queryservice.Query.BeginStreamExecute:input_type -> query.BeginStreamExecuteRequest
	16, // 16: queryservice.Query.MessageStream:input_type -> query.MessageStreamRequest
	17, // 17: queryservice.Query.MessageAck:input_type -> query.MessageAckRequest
	18, // 18: queryservice.Query.ReserveExecute:input_type -> query.ReserveExecuteRequest
	19, // 19: queryservice.Query.ReserveBeginExecute:input_type -> query.ReserveBeginExecuteRequest
	20, // 20: queryservice.Query.ReserveStreamExecute:input_type -> query.ReserveStreamExecuteRequest
	21, // 21: queryservice.Query.ReserveBeginStreamExecute:input_type -> query.ReserveBeginStreamExecuteRequest
	22, // 22: queryservice.Query.Release:input_type -> query.ReleaseRequest
	23, // 23: queryservice.Query.StreamHealth:input_type -> query.StreamHealthRequest
	24, // 24: queryservice.Query.VStream:input_type -> binlogdata.VStreamRequest
	25, // 25: queryservice.Query.VStreamRows:input_type -> binlogdata.VStreamRowsRequest
	26, // 26: queryservice.Query.VStreamTables:input_type -> binlogdata.VStreamTablesRequest
	27, // 27: queryservice.Query.VStreamResults:input_type -> binlogdata.VStreamResultsRequest
	28, // 28: queryservice.Query.GetSchema:input_type -> query.GetSchemaRequest
	29, // 29: queryservice.Query.Execute:output_type -> query.ExecuteResponse
	30, // 30: queryservice.Query.StreamExecute:output_type -> query.StreamExecuteResponse
	31, // 31: queryservice.Query.Begin:output_type -> query.BeginResponse
	32, // 32: queryservice.Query.Commit:output_type -> query.CommitResponse
	33, // 33: queryservice.Query.Rollback:output_type -> query.RollbackResponse
	34, // 34: queryservice.Query.Prepare:output_type -> query.PrepareResponse
	35, // 35: queryservice.Query.CommitPrepared:output_type -> query.CommitPreparedResponse
	36, // 36: queryservice.Query.RollbackPrepared:output_type -> query.RollbackPreparedResponse
	37, // 37: queryservice.Query.CreateTransaction:output_type -> query.CreateTransactionResponse
	38, // 38: queryservice.Query.StartCommit:output_type -> query.StartCommitResponse
	39, // 39: queryservice.Query.SetRollback:output_type -> query.SetRollbackResponse
	40, // 40: queryservice.Query.ConcludeTransaction:output_type -> query.ConcludeTransactionResponse
	41, // 41: queryservice.Query.ReadTransaction:output_type -> query.ReadTransactionResponse
	42, // 42: queryservice.Query.UnresolvedTransactions:output_type -> query.UnresolvedTransactionsResponse
	43, // 43: queryservice.Query.BeginExecute:output_type -> query.BeginExecuteResponse
	44, // 44: queryservice.Query.BeginStreamExecute:output_type -> query.BeginStreamExecuteResponse
	45, // 45: queryservice.Query.MessageStream:output_type -> query.MessageStreamResponse
	46, // 46: queryservice.Query.MessageAck:output_type -> query.MessageAckResponse
	47, // 47: queryservice.Query.ReserveExecute:output_type -> query.ReserveExecuteResponse
	48, // 48: queryservice.Query.ReserveBeginExecute:output_type -> query.ReserveBeginExecuteResponse
	49, // 49: queryservice.Query.ReserveStreamExecute:output_type -> query.ReserveStreamExecuteResponse
	50, // 50: queryservice.Query.ReserveBeginStreamExecute:output_type -> query.ReserveBeginStreamExecuteResponse
	51, // 51: queryservice.Query.Release:output_type -> query.ReleaseResponse
	52, // 52: queryservice.Query.StreamHealth:output_type -> query.StreamHealthResponse
	53, // 53: queryservice.Query.VStream:output_type -> binlogdata.VStreamResponse
	54, // 54: queryservice.Query.VStreamRows:output_type -> binlogdata.VStreamRowsResponse
	55, // 55: queryservice.Query.VStreamTables:output_type -> binlogdata.VStreamTablesResponse
	56, // 56: queryservice.Query.VStreamResults:output_type -> binlogdata.VStreamResultsResponse
	57, // 57: queryservice.Query.GetSchema:output_type -> query.GetSchemaResponse
	29, // [29:58] is the sub-list for method output_type
	0,  // [0:29] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_queryservice_proto_init() }
func file_queryservice_proto_init() {
	if File_queryservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_queryservice_proto_rawDesc), len(file_queryservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queryservice_proto_goTypes,
		DependencyIndexes: file_queryservice_proto_depIdxs,
	}.Build()
	File_queryservice_proto = out.File
	file_queryservice_proto_goTypes = nil
	file_queryservice_proto_depIdxs = nil
}
