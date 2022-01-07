package antlrimp

import (
	parser "engine/engine/antlr"
	"engine/engine/pojo"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/golang-collections/collections/stack"
	"strings"
)

func NewEngineListenerImpl() *EngineListenerImpl {
	return &EngineListenerImpl{
		mulExpression: pojo.MulExpression{},
		Stack:         stack.New(),
	}
}

type EngineListenerImpl struct {
	parser.BaseengineListener
	mulExpression pojo.MulExpression
	Stack         *stack.Stack
}

// VisitTerminal is called when a terminal node is visited.
func (s *EngineListenerImpl) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *EngineListenerImpl) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *EngineListenerImpl) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *EngineListenerImpl) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *EngineListenerImpl) EnterStart(ctx *parser.StartContext) {}

// ExitStart is called when production start is exited.
func (s *EngineListenerImpl) ExitStart(ctx *parser.StartContext) {}

// EnterRule is called when production rule is entered.
func (s *EngineListenerImpl) EnterRule(ctx *parser.RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *EngineListenerImpl) ExitRule(ctx *parser.RuleContext) {

	mulExpression := pojo.MulExpression{}

	rightExp := s.Stack.Pop()
	switch rightExp.(type) {
	case pojo.SingleExpression:
		rightExp := rightExp.(pojo.SingleExpression)
		mulExpression.RightExpression = &rightExp
	case pojo.MulExpression:
		rightExp := rightExp.(pojo.MulExpression)
		mulExpression.RightExpression = &rightExp
	}

	oper := s.Stack.Pop().(string)
	mulExpression.Operator = pojo.Operator{AndOrCmp: oper}

	leftExp := s.Stack.Pop()
	switch leftExp.(type) {
	case pojo.SingleExpression:
		leftExp := leftExp.(pojo.SingleExpression)
		mulExpression.LeftExpression = &leftExp
	case pojo.MulExpression:
		leftExp := leftExp.(pojo.MulExpression)
		mulExpression.LeftExpression = &leftExp
	}

	s.Stack.Push(mulExpression)
}

// EnterMulExpression is called when production mulExpression is entered.
func (s *EngineListenerImpl) EnterMulExpression(ctx *parser.MulExpressionContext) {}

// ExitMulExpression is called when production mulExpression is exited.
func (s *EngineListenerImpl) ExitMulExpression(ctx *parser.MulExpressionContext) {
	rightExp := s.Stack.Pop().(pojo.SingleExpression)
	oper := s.Stack.Pop().(string)
	leftExp := s.Stack.Pop().(pojo.SingleExpression)
	mulExpression := pojo.MulExpression{
		LeftExpression:  &leftExp,
		RightExpression: &rightExp,
		Operator:        pojo.Operator{AndOrCmp: oper},
	}
	s.Stack.Push(mulExpression)
}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *EngineListenerImpl) EnterSingleExpression(ctx *parser.SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *EngineListenerImpl) ExitSingleExpression(ctx *parser.SingleExpressionContext) {
	varCmp := s.Stack.Pop()
	cmpOpe := s.Stack.Pop().(string)
	varDot := s.Stack.Pop()
	singleExpression := pojo.SingleExpression{
		VarDot:   varDot,
		VarCmp:   varCmp,
		Operator: pojo.Operator{CompareOpe: cmpOpe},
	}
	s.Stack.Push(singleExpression)
}

// EnterVariable is called when production variable is entered.
func (s *EngineListenerImpl) EnterVariable(ctx *parser.VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *EngineListenerImpl) ExitVariable(ctx *parser.VariableContext) {
	s.Stack.Push(strings.Trim(ctx.GetText(), "\""))
}

// EnterCompareEqual is called when production compareEqual is entered.
func (s *EngineListenerImpl) EnterCompareEqual(ctx *parser.CompareEqualContext) {

}

// ExitCompareEqual is called when production compareEqual is exited.
func (s *EngineListenerImpl) ExitCompareEqual(ctx *parser.CompareEqualContext) {
	s.Stack.Push(ctx.GetText())
}

// EnterCompareIn is called when production compareIn is entered.
func (s *EngineListenerImpl) EnterCompareIn(ctx *parser.CompareInContext) {}

// ExitCompareIn is called when production compareIn is exited.
func (s *EngineListenerImpl) ExitCompareIn(ctx *parser.CompareInContext) {
	s.Stack.Push(ctx.GetText())
}

// EnterLogicOper is called when production logicOper is entered.
func (s *EngineListenerImpl) EnterLogicOper(ctx *parser.LogicOperContext) {}

// ExitLogicOper is called when production logicOper is exited.
func (s *EngineListenerImpl) ExitLogicOper(ctx *parser.LogicOperContext) {
	s.Stack.Push(ctx.GetText())
}
