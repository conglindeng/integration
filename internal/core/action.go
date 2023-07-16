package core


//communicate by channel
type Action interface {
	Act(param Param) Result
}
