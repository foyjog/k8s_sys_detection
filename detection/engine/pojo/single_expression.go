package pojo

import (
	v1 "k8s.io/apiserver/pkg/apis/audit/v1"
	"strings"
)

type SingleExpression struct {
	VarDot   interface{}
	VarCmp   interface{}
	Operator Operator
}

func (e *SingleExpression) GetVarDotValue(event *v1.Event) interface{} {
	switch e.VarDot {
	case "event.ObjectRef.Resource":
		return event.ObjectRef.Resource
	case "event.ObjectRef.Name":
		return event.ObjectRef.Name
	case "event.Verb":
		return event.Verb
	case "event.UserAgent0":
		return strings.Split(event.UserAgent, "/")[0]
	case "event.ObjectRef.Namespace":
		return event.ObjectRef.Namespace
	default:
		return nil

	}
}

func (e *SingleExpression) Evaluate(event *v1.Event) (ret bool) {
	switch e.Operator.CompareOpe {
	case "==":
		varDotValue := e.GetVarDotValue(event)
		result := varDotValue == e.VarCmp
		return result
	case "!=":
		varDotValue := e.GetVarDotValue(event)
		result := varDotValue != e.VarCmp
		return result
	case "and":
		leftExpression := e.VarDot.(SingleExpression)
		rightExpression := e.VarCmp.(SingleExpression)
		return leftExpression.Evaluate(event) && rightExpression.Evaluate(event)
	case "or":
		leftExpression := e.VarDot.(SingleExpression)
		rightExpression := e.VarCmp.(SingleExpression)
		return leftExpression.Evaluate(event) || rightExpression.Evaluate(event)
	}
	return false
}
