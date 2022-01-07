package pojo

import v1 "k8s.io/apiserver/pkg/apis/audit/v1"

type MulExpression struct {
	LeftExpression  interface{ Expression }
	RightExpression interface{ Expression }
	Operator        Operator
	RuleName        string
	RuleDescription string
}

func (e *MulExpression) Evaluate(event *v1.Event) (ret bool) {
	switch e.Operator.AndOrCmp {
	case "and":
		return e.RightExpression.Evaluate(event) && e.LeftExpression.Evaluate(event)
	case "or":
		return e.RightExpression.Evaluate(event) || e.LeftExpression.Evaluate(event)
	default:
		return false
	}

}
