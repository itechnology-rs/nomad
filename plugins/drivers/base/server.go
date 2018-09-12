package base

import (
	"golang.org/x/net/context"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/nomad/plugins/drivers/base/proto"
)

type driverPluginServer struct {
	broker *plugin.GRPCBroker
	impl   DriverPlugin
}

func (b *driverPluginServer) TaskConfigSchema(ctx context.Context, req *proto.TaskConfigSchemaRequest) (*proto.TaskConfigSchemaResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) Capabilities(ctx context.Context, req *proto.CapabilitiesRequest) (*proto.CapabilitiesResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) Fingerprint(req *proto.FingerprintRequest, srv proto.Driver_FingerprintServer) error {
	return nil
}

func (b *driverPluginServer) RecoverTask(ctx context.Context, req *proto.RecoverTaskRequest) (*proto.RecoverTaskResponse, error) {
	return nil, b.impl.RecoverTask(taskHandleFromProto(req.Handle))
}

func (b *driverPluginServer) StartTask(ctx context.Context, req *proto.StartTaskRequest) (*proto.StartTaskResponse, error) {
	handle, err := b.impl.StartTask(taskConfigFromProto(req.Task))
	if err != nil {
		return nil, err
	}

	return &proto.StartTaskResponse{
		Handle: taskHandleToProto(handle),
	}, nil
}

func (b *driverPluginServer) WaitTask(ctx context.Context, req *proto.WaitTaskRequest) (*proto.WaitTaskResponse, error) {
	ch := b.impl.WaitTask(req.TaskId)
	result := <-ch
	var errStr string
	if result.Err != nil {
		errStr = result.Err.Error()
	}
	return &proto.WaitTaskResponse{
		Err: errStr,
		Result: &proto.ExitResult{
			ExitCode:  int32(result.ExitCode),
			Signal:    int32(result.Signal),
			OomKilled: result.OOMKilled,
		},
	}, nil
}

func (b *driverPluginServer) StopTask(ctx context.Context, req *proto.StopTaskRequest) (*proto.StopTaskResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) DestroyTask(ctx context.Context, req *proto.DestroyTaskRequest) (*proto.DestroyTaskResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) InspectTask(ctx context.Context, req *proto.InspectTaskRequest) (*proto.InspectTaskResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) TaskStats(ctx context.Context, req *proto.TaskStatsRequest) (*proto.TaskStatsResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) ExecTask(ctx context.Context, req *proto.ExecTaskRequest) (*proto.ExecTaskResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) SignalTask(ctx context.Context, req *proto.SignalTaskRequest) (*proto.SignalTaskResponse, error) {
	return nil, nil
}

func (b *driverPluginServer) TaskEvents(req *proto.TaskEventsRequest, srv proto.Driver_TaskEventsServer) error {
	return nil
}