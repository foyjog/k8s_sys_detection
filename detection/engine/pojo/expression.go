package pojo

import v1 "k8s.io/apiserver/pkg/apis/audit/v1"

type Expression interface {
	Evaluate(event *v1.Event) bool
}
