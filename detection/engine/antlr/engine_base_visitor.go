// Code generated from C:/Users/foyjo/GolandProjects/engine/engine/antlr\engine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // engine

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseengineVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseengineVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitRule(ctx *RuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitMulExpression(ctx *MulExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitSingleExpression(ctx *SingleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitCompareEqual(ctx *CompareEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitCompareIn(ctx *CompareInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseengineVisitor) VisitLogicOper(ctx *LogicOperContext) interface{} {
	return v.VisitChildren(ctx)
}
