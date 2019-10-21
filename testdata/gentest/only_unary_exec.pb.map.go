// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata/gentest/only_unary_exec.proto

package gentest

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"

	//protoc-gen-map packages
	bytes "bytes"
	context "context"
	sql "database/sql"
	mapper "github.com/jackskj/protoc-gen-map/mapper"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	log "log"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Code generated by protoc-gen-map. DO NOT EDIT.
// To Use:
// 1. Instantiate MapperServers with sql.DB connection
// 2. Register MapperServer as the gRPC service server
// 3. Begin serving

type OnlyExecServiceMapServer struct {
	DB           *sql.DB
	InsertMapper *mapper.Mapper

	mapperGenMux sync.Mutex
}

func (m *OnlyExecServiceMapServer) Insert(ctx context.Context, r *OnlyExec) (*OnlyExec, error) {
	sqlBuffer := &bytes.Buffer{}
	if err := sqlTemplate.ExecuteTemplate(sqlBuffer, "Insert", r); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rawSql := sqlBuffer.String()

	_, err := m.DB.Exec(rawSql)
	if err != nil {
		log.Printf("error executing query.\n OnlyExec request: %s \n,query: %s \n error: %s", r, rawSql, err)
		return nil, status.Error(codes.InvalidArgument, "request generated malformed query")
	}
	resp := OnlyExec{}
	return &resp, nil

}
