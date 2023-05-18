package planner

import "pi-coursework-server/processors"

type Plan struct {
	Plan []processors.IProcessor
}

func NewPlan() *Plan {
	return &Plan{
		Plan: make([]processors.IProcessor, 0),
	}
}
