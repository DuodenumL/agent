// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/projecteru2/agent/types"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// GetIdentifier provides a mock function with given fields: ctx
func (_m *Store) GetIdentifier(ctx context.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNode provides a mock function with given fields: ctx, nodename
func (_m *Store) GetNode(ctx context.Context, nodename string) (*types.Node, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Node); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeStatus provides a mock function with given fields: ctx, nodename
func (_m *Store) GetNodeStatus(ctx context.Context, nodename string) (*types.NodeStatus, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *types.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.NodeStatus); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NodeStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPodNodes provides a mock function with given fields: ctx, all, podname, labels
func (_m *Store) ListPodNodes(ctx context.Context, all bool, podname string, labels map[string]string) ([]*types.Node, error) {
	ret := _m.Called(ctx, all, podname, labels)

	var r0 []*types.Node
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, map[string]string) []*types.Node); ok {
		r0 = rf(ctx, all, podname, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, bool, string, map[string]string) error); ok {
		r1 = rf(ctx, all, podname, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeStatusStream provides a mock function with given fields: ctx
func (_m *Store) NodeStatusStream(ctx context.Context) (<-chan *types.NodeStatus, <-chan error) {
	ret := _m.Called(ctx)

	var r0 <-chan *types.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *types.NodeStatus); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *types.NodeStatus)
		}
	}

	var r1 <-chan error
	if rf, ok := ret.Get(1).(func(context.Context) <-chan error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// SetNodeStatus provides a mock function with given fields: ctx, ttl
func (_m *Store) SetNodeStatus(ctx context.Context, ttl int64) error {
	ret := _m.Called(ctx, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWorkloadStatus provides a mock function with given fields: ctx, status, ttl
func (_m *Store) SetWorkloadStatus(ctx context.Context, status *types.WorkloadStatus, ttl int64) error {
	ret := _m.Called(ctx, status, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.WorkloadStatus, int64) error); ok {
		r0 = rf(ctx, status, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
