package execute

import "conglindeng/integration/internal/core"

func Execute(context core.Context) core.Result {

	//todo:dcl  channel?  log param&result?  find a component for execution
	// p := context.Param

	doExecute(context)

	return context.Result
}

func doExecute(context core.Context) {

}
