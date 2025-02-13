package node

import (
	"context"
	"errors"

	"github.com/projecteru2/agent/common"
	"github.com/projecteru2/agent/runtime"
	"github.com/projecteru2/agent/runtime/docker"
	runtimemocks "github.com/projecteru2/agent/runtime/mocks"
	"github.com/projecteru2/agent/runtime/yavirt"
	"github.com/projecteru2/agent/store"
	corestore "github.com/projecteru2/agent/store/core"
	storemocks "github.com/projecteru2/agent/store/mocks"
	"github.com/projecteru2/agent/types"
	"github.com/projecteru2/agent/utils"

	log "github.com/sirupsen/logrus"
)

// Manager manages node status
type Manager struct {
	config        *types.Config
	store         store.Store
	runtimeClient runtime.Runtime
}

// NewManager .
func NewManager(ctx context.Context, config *types.Config) (*Manager, error) {
	m := &Manager{config: config}
	switch config.Store {
	case common.GRPCStore:
		corestore.Init(ctx, config)
		m.store = corestore.Get()
		if m.store == nil {
			return nil, errors.New("failed to get store client")
		}
	case common.MocksStore:
		m.store = storemocks.FromTemplate()
	default:
		return nil, errors.New("unknown store type")
	}

	switch config.Runtime {
	case common.DockerRuntime:
		node, err := m.store.GetNode(ctx, config.HostName)
		if err != nil {
			log.Errorf("[NewManager] failed to get node %s, err: %s", config.HostName, err)
			return nil, err
		}

		nodeIP := utils.GetIP(node.Endpoint)
		if nodeIP == "" {
			nodeIP = common.LocalIP
		}
		docker.InitClient(config, nodeIP)
		m.runtimeClient = docker.GetClient()
		if m.runtimeClient == nil {
			return nil, errors.New("failed to get runtime client")
		}
	case common.YavirtRuntime:
		yavirt.InitClient(config)
		m.runtimeClient = yavirt.GetClient()
		if m.runtimeClient == nil {
			return nil, errors.New("failed to get runtime client")
		}
	case common.MocksRuntime:
		m.runtimeClient = runtimemocks.FromTemplate()
	default:
		return nil, errors.New("unknown runtime type")
	}

	return m, nil
}

// Run runs a node manager
func (m *Manager) Run(ctx context.Context) error {
	log.Info("[NodeManager] start node status heartbeat")
	go m.heartbeat(ctx)

	<-ctx.Done()
	log.Info("[NodeManager] exiting")
	return nil
}

// Exit .
func (m *Manager) Exit() error {
	log.Info("[NodeManager] exiting")
	log.Infof("[NodeManager] remove node status of %s", m.config.HostName)

	// ctx is now canceled. use a new context.
	var err error
	utils.WithTimeout(context.TODO(), m.config.GlobalConnectionTimeout, func(ctx context.Context) {
		// remove node status
		err = m.store.SetNodeStatus(ctx, -1)
	})
	if err != nil {
		log.Errorf("[NodeManager] failed to remove node status of %v, err: %s", m.config.HostName, err)
		return err
	}
	return nil
}
