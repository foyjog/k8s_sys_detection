// Code generated from C:/Users/foyjo/GolandProjects/engine/engine/antlr\engine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // engine

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 70, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 23, 10, 3, 3, 3, 3, 
	3, 3, 3, 5, 3, 28, 10, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 4, 3, 4, 3, 4, 
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 
	5, 5, 49, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 55, 10, 5, 12, 5, 14, 5, 
	58, 11, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 66, 10, 8, 3, 9, 3, 
	9, 3, 9, 2, 3, 8, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 5, 4, 2, 9, 9, 11, 
	12, 3, 2, 5, 6, 3, 2, 7, 8, 2, 68, 2, 18, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 
	6, 33, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 61, 3, 2, 
	2, 2, 14, 65, 3, 2, 2, 2, 16, 67, 3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19, 3, 
	3, 2, 2, 2, 20, 23, 5, 8, 5, 2, 21, 23, 5, 6, 4, 2, 22, 20, 3, 2, 2, 2, 
	22, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 27, 5, 16, 9, 2, 25, 28, 5, 
	8, 5, 2, 26, 28, 5, 6, 4, 2, 27, 25, 3, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 
	32, 3, 2, 2, 2, 29, 32, 5, 8, 5, 2, 30, 32, 5, 6, 4, 2, 31, 22, 3, 2, 2, 
	2, 31, 29, 3, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 34, 7, 
	3, 2, 2, 34, 35, 5, 8, 5, 2, 35, 36, 5, 16, 9, 2, 36, 37, 5, 8, 5, 2, 37, 
	38, 7, 4, 2, 2, 38, 7, 3, 2, 2, 2, 39, 40, 8, 5, 1, 2, 40, 41, 5, 10, 6, 
	2, 41, 42, 5, 12, 7, 2, 42, 43, 5, 10, 6, 2, 43, 49, 3, 2, 2, 2, 44, 45, 
	5, 10, 6, 2, 45, 46, 5, 14, 8, 2, 46, 47, 5, 10, 6, 2, 47, 49, 3, 2, 2, 
	2, 48, 39, 3, 2, 2, 2, 48, 44, 3, 2, 2, 2, 49, 56, 3, 2, 2, 2, 50, 51, 
	12, 5, 2, 2, 51, 52, 5, 16, 9, 2, 52, 53, 5, 8, 5, 6, 53, 55, 3, 2, 2, 
	2, 54, 50, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 
	3, 2, 2, 2, 57, 9, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 9, 2, 2, 2, 
	60, 11, 3, 2, 2, 2, 61, 62, 9, 3, 2, 2, 62, 13, 3, 2, 2, 2, 63, 66, 5, 
	12, 7, 2, 64, 66, 7, 10, 2, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 
	66, 15, 3, 2, 2, 2, 67, 68, 9, 4, 2, 2, 68, 17, 3, 2, 2, 2, 8, 22, 27, 
	31, 48, 56, 65,
}
var literalNames = []string{
	"", "'('", "')'", "'=='", "'!='", "'and'", "'or'", "", "'in'", "", "", 
	"", "'.'",
}
var symbolicNames = []string{
	"", "", "", "EQUAL", "NOT_EQUAL", "AND", "OR", "INTEGER", "IN", "STRING", 
	"VAR_DOT", "VAR", "DOT", "SL_COMMENT", "WS",
}

var ruleNames = []string{
	"start", "rule", "mulExpression", "singleExpression", "variable", "compareEqual", 
	"compareIn", "logicOper",
}
type engineParser struct {
	*antlr.BaseParser
}

// NewengineParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *engineParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewengineParser(input antlr.TokenStream) *engineParser {
	this := new(engineParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "engine.g4"

	return this
}


// engineParser tokens.
const (
	engineParserEOF = antlr.TokenEOF
	engineParserT__0 = 1
	engineParserT__1 = 2
	engineParserEQUAL = 3
	engineParserNOT_EQUAL = 4
	engineParserAND = 5
	engineParserOR = 6
	engineParserINTEGER = 7
	engineParserIN = 8
	engineParserSTRING = 9
	engineParserVAR_DOT = 10
	engineParserVAR = 11
	engineParserDOT = 12
	engineParserSL_COMMENT = 13
	engineParserWS = 14
)

// engineParser rules.
const (
	engineParserRULE_start = 0
	engineParserRULE_rule = 1
	engineParserRULE_mulExpression = 2
	engineParserRULE_singleExpression = 3
	engineParserRULE_variable = 4
	engineParserRULE_compareEqual = 5
	engineParserRULE_compareIn = 6
	engineParserRULE_logicOper = 7
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Rule() IRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, engineParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Rule()
	}



	return localctx
}


// IRuleContext is an interface to support dynamic dispatch.
type IRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContext differentiates from other interfaces.
	IsRuleContext()
}

type RuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContext() *RuleContext {
	var p = new(RuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_rule
	return p
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) LogicOper() ILogicOperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicOperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicOperContext)
}

func (s *RuleContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *RuleContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *RuleContext) AllMulExpression() []IMulExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMulExpressionContext)(nil)).Elem())
	var tst = make([]IMulExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMulExpressionContext)
		}
	}

	return tst
}

func (s *RuleContext) MulExpression(i int) IMulExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMulExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMulExpressionContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterRule(s)
	}
}

func (s *RuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitRule(s)
	}
}

func (s *RuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitRule(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) Rule() (localctx IRuleContext) {
	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, engineParserRULE_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(20)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case engineParserINTEGER, engineParserSTRING, engineParserVAR_DOT:
			{
				p.SetState(18)
				p.singleExpression(0)
			}


		case engineParserT__0:
			{
				p.SetState(19)
				p.MulExpression()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(22)
			p.LogicOper()
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case engineParserINTEGER, engineParserSTRING, engineParserVAR_DOT:
			{
				p.SetState(23)
				p.singleExpression(0)
			}


		case engineParserT__0:
			{
				p.SetState(24)
				p.MulExpression()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.singleExpression(0)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.MulExpression()
		}

	}


	return localctx
}


// IMulExpressionContext is an interface to support dynamic dispatch.
type IMulExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulExpressionContext differentiates from other interfaces.
	IsMulExpressionContext()
}

type MulExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulExpressionContext() *MulExpressionContext {
	var p = new(MulExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_mulExpression
	return p
}

func (*MulExpressionContext) IsMulExpressionContext() {}

func NewMulExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulExpressionContext {
	var p = new(MulExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_mulExpression

	return p
}

func (s *MulExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MulExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *MulExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MulExpressionContext) LogicOper() ILogicOperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicOperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicOperContext)
}

func (s *MulExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MulExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterMulExpression(s)
	}
}

func (s *MulExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitMulExpression(s)
	}
}

func (s *MulExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitMulExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) MulExpression() (localctx IMulExpressionContext) {
	localctx = NewMulExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, engineParserRULE_mulExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(engineParserT__0)
	}
	{
		p.SetState(32)
		p.singleExpression(0)
	}
	{
		p.SetState(33)
		p.LogicOper()
	}
	{
		p.SetState(34)
		p.singleExpression(0)
	}
	{
		p.SetState(35)
		p.Match(engineParserT__1)
	}



	return localctx
}


// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *SingleExpressionContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *SingleExpressionContext) CompareEqual() ICompareEqualContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareEqualContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareEqualContext)
}

func (s *SingleExpressionContext) CompareIn() ICompareInContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareInContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareInContext)
}

func (s *SingleExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *SingleExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *SingleExpressionContext) LogicOper() ILogicOperContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicOperContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicOperContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SingleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterSingleExpression(s)
	}
}

func (s *SingleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitSingleExpression(s)
	}
}

func (s *SingleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitSingleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *engineParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *engineParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, engineParserRULE_singleExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(38)
			p.Variable()
		}
		{
			p.SetState(39)
			p.CompareEqual()
		}
		{
			p.SetState(40)
			p.Variable()
		}


	case 2:
		{
			p.SetState(42)
			p.Variable()
		}
		{
			p.SetState(43)
			p.CompareIn()
		}
		{
			p.SetState(44)
			p.Variable()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSingleExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, engineParserRULE_singleExpression)
			p.SetState(48)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(49)
				p.LogicOper()
			}
			{
				p.SetState(50)
				p.singleExpression(4)
			}


		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}



	return localctx
}


// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VAR_DOT() antlr.TerminalNode {
	return s.GetToken(engineParserVAR_DOT, 0)
}

func (s *VariableContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(engineParserINTEGER, 0)
}

func (s *VariableContext) STRING() antlr.TerminalNode {
	return s.GetToken(engineParserSTRING, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, engineParserRULE_variable)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << engineParserINTEGER) | (1 << engineParserSTRING) | (1 << engineParserVAR_DOT))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ICompareEqualContext is an interface to support dynamic dispatch.
type ICompareEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareEqualContext differentiates from other interfaces.
	IsCompareEqualContext()
}

type CompareEqualContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareEqualContext() *CompareEqualContext {
	var p = new(CompareEqualContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_compareEqual
	return p
}

func (*CompareEqualContext) IsCompareEqualContext() {}

func NewCompareEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareEqualContext {
	var p = new(CompareEqualContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_compareEqual

	return p
}

func (s *CompareEqualContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareEqualContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(engineParserEQUAL, 0)
}

func (s *CompareEqualContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(engineParserNOT_EQUAL, 0)
}

func (s *CompareEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareEqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CompareEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterCompareEqual(s)
	}
}

func (s *CompareEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitCompareEqual(s)
	}
}

func (s *CompareEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitCompareEqual(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) CompareEqual() (localctx ICompareEqualContext) {
	localctx = NewCompareEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, engineParserRULE_compareEqual)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(_la == engineParserEQUAL || _la == engineParserNOT_EQUAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ICompareInContext is an interface to support dynamic dispatch.
type ICompareInContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareInContext differentiates from other interfaces.
	IsCompareInContext()
}

type CompareInContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareInContext() *CompareInContext {
	var p = new(CompareInContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_compareIn
	return p
}

func (*CompareInContext) IsCompareInContext() {}

func NewCompareInContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareInContext {
	var p = new(CompareInContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_compareIn

	return p
}

func (s *CompareInContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareInContext) CompareEqual() ICompareEqualContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareEqualContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareEqualContext)
}

func (s *CompareInContext) IN() antlr.TerminalNode {
	return s.GetToken(engineParserIN, 0)
}

func (s *CompareInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareInContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CompareInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterCompareIn(s)
	}
}

func (s *CompareInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitCompareIn(s)
	}
}

func (s *CompareInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitCompareIn(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) CompareIn() (localctx ICompareInContext) {
	localctx = NewCompareInContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, engineParserRULE_compareIn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case engineParserEQUAL, engineParserNOT_EQUAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.CompareEqual()
		}


	case engineParserIN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Match(engineParserIN)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ILogicOperContext is an interface to support dynamic dispatch.
type ILogicOperContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicOperContext differentiates from other interfaces.
	IsLogicOperContext()
}

type LogicOperContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicOperContext() *LogicOperContext {
	var p = new(LogicOperContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = engineParserRULE_logicOper
	return p
}

func (*LogicOperContext) IsLogicOperContext() {}

func NewLogicOperContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicOperContext {
	var p = new(LogicOperContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = engineParserRULE_logicOper

	return p
}

func (s *LogicOperContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicOperContext) AND() antlr.TerminalNode {
	return s.GetToken(engineParserAND, 0)
}

func (s *LogicOperContext) OR() antlr.TerminalNode {
	return s.GetToken(engineParserOR, 0)
}

func (s *LogicOperContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicOperContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicOperContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.EnterLogicOper(s)
	}
}

func (s *LogicOperContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(engineListener); ok {
		listenerT.ExitLogicOper(s)
	}
}

func (s *LogicOperContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case engineVisitor:
		return t.VisitLogicOper(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *engineParser) LogicOper() (localctx ILogicOperContext) {
	localctx = NewLogicOperContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, engineParserRULE_logicOper)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !(_la == engineParserAND || _la == engineParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


func (p *engineParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
			var t *SingleExpressionContext = nil
			if localctx != nil { t = localctx.(*SingleExpressionContext) }
			return p.SingleExpression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *engineParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

