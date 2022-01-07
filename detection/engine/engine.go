package engine

import (
	parser "engine/engine/antlr"
	"engine/engine/antlrimp"
	"engine/engine/pojo"
	"engine/engine/yaml"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	v1 "k8s.io/apiserver/pkg/apis/audit/v1"
	"log"
)

var K8sRules = pojo.Rules{}

func InitRules() {
	log.Println("Engine init：Loading rules.")
	K8sRules.MulExpression = make([]pojo.MulExpression, 0)
	rules := yaml.ReadK8sRulesFromYml("engine/yaml/files/k8s_rules.yml")
	log.Printf("Engine init：Loading rules done. loading %d rules\n", len(rules))
	for _, rule := range rules["rules"] {
		//fmt.Println("Adding new rule:", rule.Condition)
		in := antlr.NewInputStream(rule.Condition)
		lexer := parser.NewengineLexer(in)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		listener := antlrimp.NewEngineListenerImpl()

		engineParser := parser.NewengineParser(stream)
		engineParser.BuildParseTrees = true

		antlr.ParseTreeWalkerDefault.Walk(listener, engineParser.Start())
		mulExpression := listener.Stack.Pop().(pojo.MulExpression)
		mulExpression.RuleName = rule.Name
		mulExpression.RuleDescription = rule.Description
		K8sRules.MulExpression = append(K8sRules.MulExpression, mulExpression)
	}

}

func Fire(event *v1.Event) {
	for _, mulExpression := range K8sRules.MulExpression {
		if mulExpression.Evaluate(event) == true {
			log.Println("Alert: found", mulExpression.RuleDescription)
		}
	}
}
