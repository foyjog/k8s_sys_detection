// Code generated from C:/Users/foyjo/GolandProjects/engine/engine/antlr\engine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // engine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// engineListener is a complete listener for a parse tree produced by engineParser.
type engineListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterRule is called when entering the rule production.
	EnterRule(c *RuleContext)

	// EnterMulExpression is called when entering the mulExpression production.
	EnterMulExpression(c *MulExpressionContext)

	// EnterSingleExpression is called when entering the singleExpression production.
	EnterSingleExpression(c *SingleExpressionContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterCompareEqual is called when entering the compareEqual production.
	EnterCompareEqual(c *CompareEqualContext)

	// EnterCompareIn is called when entering the compareIn production.
	EnterCompareIn(c *CompareInContext)

	// EnterLogicOper is called when entering the logicOper production.
	EnterLogicOper(c *LogicOperContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitRule is called when exiting the rule production.
	ExitRule(c *RuleContext)

	// ExitMulExpression is called when exiting the mulExpression production.
	ExitMulExpression(c *MulExpressionContext)

	// ExitSingleExpression is called when exiting the singleExpression production.
	ExitSingleExpression(c *SingleExpressionContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitCompareEqual is called when exiting the compareEqual production.
	ExitCompareEqual(c *CompareEqualContext)

	// ExitCompareIn is called when exiting the compareIn production.
	ExitCompareIn(c *CompareInContext)

	// ExitLogicOper is called when exiting the logicOper production.
	ExitLogicOper(c *LogicOperContext)
}
