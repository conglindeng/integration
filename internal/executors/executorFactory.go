package executors

import "conglindeng/integration/internal/core"

type ExecutorFacotry struct {
	executors map[int]func() core.Action
}

func CreateExecutorFactory() *ExecutorFacotry {
	f := &ExecutorFacotry{executors: map[int]func() core.Action{}}
	f.registerExecutor(0, func() core.Action { return &HttpExecutor{} })

	return f
}

func (f *ExecutorFacotry) registerExecutor(subType int, creator func() core.Action) {
	f.executors[subType] = creator
}

func (f *ExecutorFacotry) createExecutor(subType int) core.Action {
	return f.executors[subType]()
}
