// Code generated from C:/Users/foyjo/GolandProjects/engine/engine/antlr\engine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // engine

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by engineParser.
type engineVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by engineParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by engineParser#rule.
	VisitRule(ctx *RuleContext) interface{}

	// Visit a parse tree produced by engineParser#mulExpression.
	VisitMulExpression(ctx *MulExpressionContext) interface{}

	// Visit a parse tree produced by engineParser#singleExpression.
	VisitSingleExpression(ctx *SingleExpressionContext) interface{}

	// Visit a parse tree produced by engineParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by engineParser#compareEqual.
	VisitCompareEqual(ctx *CompareEqualContext) interface{}

	// Visit a parse tree produced by engineParser#compareIn.
	VisitCompareIn(ctx *CompareInContext) interface{}

	// Visit a parse tree produced by engineParser#logicOper.
	VisitLogicOper(ctx *LogicOperContext) interface{}

}