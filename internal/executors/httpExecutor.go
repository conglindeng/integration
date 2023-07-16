package executors

import "conglindeng/integration/internal/core"

type HttpExecutor struct {
	BaseExecutor
	
}

func (h *HttpExecutor) Act(core.Param) core.Result {
	return core.Sucess_Result(core.Data{Data: 1})
}
