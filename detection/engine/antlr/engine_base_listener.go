// Code generated from C:/Users/foyjo/GolandProjects/engine/engine/antlr\engine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // engine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseengineListener is a complete listener for a parse tree produced by engineParser.
type BaseengineListener struct{}

var _ engineListener = &BaseengineListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseengineListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseengineListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseengineListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseengineListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseengineListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseengineListener) ExitStart(ctx *StartContext) {}

// EnterRule is called when production rule is entered.
func (s *BaseengineListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *BaseengineListener) ExitRule(ctx *RuleContext) {}

// EnterMulExpression is called when production mulExpression is entered.
func (s *BaseengineListener) EnterMulExpression(ctx *MulExpressionContext) {}

// ExitMulExpression is called when production mulExpression is exited.
func (s *BaseengineListener) ExitMulExpression(ctx *MulExpressionContext) {}

// EnterSingleExpression is called when production singleExpression is entered.
func (s *BaseengineListener) EnterSingleExpression(ctx *SingleExpressionContext) {}

// ExitSingleExpression is called when production singleExpression is exited.
func (s *BaseengineListener) ExitSingleExpression(ctx *SingleExpressionContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseengineListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseengineListener) ExitVariable(ctx *VariableContext) {}

// EnterCompareEqual is called when production compareEqual is entered.
func (s *BaseengineListener) EnterCompareEqual(ctx *CompareEqualContext) {}

// ExitCompareEqual is called when production compareEqual is exited.
func (s *BaseengineListener) ExitCompareEqual(ctx *CompareEqualContext) {}

// EnterCompareIn is called when production compareIn is entered.
func (s *BaseengineListener) EnterCompareIn(ctx *CompareInContext) {}

// ExitCompareIn is called when production compareIn is exited.
func (s *BaseengineListener) ExitCompareIn(ctx *CompareInContext) {}

// EnterLogicOper is called when production logicOper is entered.
func (s *BaseengineListener) EnterLogicOper(ctx *LogicOperContext) {}

// ExitLogicOper is called when production logicOper is exited.
func (s *BaseengineListener) ExitLogicOper(ctx *LogicOperContext) {}
