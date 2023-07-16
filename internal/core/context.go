package core

import "conglindeng/integration/internal/model"

type Context struct {
	TraceId   string
	NodeCode  string
	Nodes     []model.Node
	AllPram   map[string]Param // what is key ？
	Param     Param
	Result    Result
	Variables map[string]interface{}
}
