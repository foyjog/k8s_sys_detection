// Code generated from C:/Users/foyjo/GolandProjects/engine/engine/antlr\engine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 106, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 
	3, 8, 5, 8, 50, 10, 8, 3, 8, 6, 8, 53, 10, 8, 13, 8, 14, 8, 54, 3, 9, 3, 
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 66, 10, 10, 12, 
	10, 14, 10, 69, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 
	5, 11, 78, 10, 11, 3, 12, 6, 12, 81, 10, 12, 13, 12, 14, 12, 82, 3, 13, 
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 91, 10, 14, 12, 14, 14, 14, 94, 
	11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 6, 15, 101, 10, 15, 13, 15, 
	14, 15, 102, 3, 15, 3, 15, 3, 92, 2, 16, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 3, 
	2, 6, 3, 2, 50, 59, 4, 2, 36, 36, 94, 94, 5, 2, 50, 59, 67, 92, 99, 124, 
	5, 2, 11, 12, 15, 15, 34, 34, 2, 114, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 
	3, 2, 2, 2, 3, 31, 3, 2, 2, 2, 5, 33, 3, 2, 2, 2, 7, 35, 3, 2, 2, 2, 9, 
	38, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 45, 3, 2, 2, 2, 15, 49, 3, 2, 2, 
	2, 17, 56, 3, 2, 2, 2, 19, 59, 3, 2, 2, 2, 21, 77, 3, 2, 2, 2, 23, 80, 
	3, 2, 2, 2, 25, 84, 3, 2, 2, 2, 27, 86, 3, 2, 2, 2, 29, 100, 3, 2, 2, 2, 
	31, 32, 7, 42, 2, 2, 32, 4, 3, 2, 2, 2, 33, 34, 7, 43, 2, 2, 34, 6, 3, 
	2, 2, 2, 35, 36, 7, 63, 2, 2, 36, 37, 7, 63, 2, 2, 37, 8, 3, 2, 2, 2, 38, 
	39, 7, 35, 2, 2, 39, 40, 7, 63, 2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7, 99, 
	2, 2, 42, 43, 7, 112, 2, 2, 43, 44, 7, 102, 2, 2, 44, 12, 3, 2, 2, 2, 45, 
	46, 7, 113, 2, 2, 46, 47, 7, 116, 2, 2, 47, 14, 3, 2, 2, 2, 48, 50, 7, 
	47, 2, 2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 
	53, 9, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 
	2, 54, 55, 3, 2, 2, 2, 55, 16, 3, 2, 2, 2, 56, 57, 7, 107, 2, 2, 57, 58, 
	7, 112, 2, 2, 58, 18, 3, 2, 2, 2, 59, 67, 7, 36, 2, 2, 60, 61, 7, 94, 2, 
	2, 61, 66, 11, 2, 2, 2, 62, 63, 7, 36, 2, 2, 63, 66, 7, 36, 2, 2, 64, 66, 
	10, 3, 2, 2, 65, 60, 3, 2, 2, 2, 65, 62, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 
	66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 
	2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 71, 7, 36, 2, 2, 71, 20, 3, 2, 2, 2, 72, 
	78, 5, 23, 12, 2, 73, 74, 5, 23, 12, 2, 74, 75, 5, 25, 13, 2, 75, 76, 5, 
	21, 11, 2, 76, 78, 3, 2, 2, 2, 77, 72, 3, 2, 2, 2, 77, 73, 3, 2, 2, 2, 
	78, 22, 3, 2, 2, 2, 79, 81, 9, 4, 2, 2, 80, 79, 3, 2, 2, 2, 81, 82, 3, 
	2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 24, 3, 2, 2, 2, 84, 
	85, 7, 48, 2, 2, 85, 26, 3, 2, 2, 2, 86, 87, 7, 49, 2, 2, 87, 88, 7, 49, 
	2, 2, 88, 92, 3, 2, 2, 2, 89, 91, 11, 2, 2, 2, 90, 89, 3, 2, 2, 2, 91, 
	94, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 95, 3, 2, 2, 
	2, 94, 92, 3, 2, 2, 2, 95, 96, 7, 12, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 
	8, 14, 2, 2, 98, 28, 3, 2, 2, 2, 99, 101, 9, 5, 2, 2, 100, 99, 3, 2, 2, 
	2, 101, 102, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 
	104, 3, 2, 2, 2, 104, 105, 8, 15, 2, 2, 105, 30, 3, 2, 2, 2, 11, 2, 49, 
	54, 65, 67, 77, 82, 92, 102, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'=='", "'!='", "'and'", "'or'", "", "'in'", "", "", 
	"", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "EQUAL", "NOT_EQUAL", "AND", "OR", "INTEGER", "IN", "STRING", 
	"VAR_DOT", "VAR", "DOT", "SL_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "EQUAL", "NOT_EQUAL", "AND", "OR", "INTEGER", "IN", "STRING", 
	"VAR_DOT", "VAR", "DOT", "SL_COMMENT", "WS",
}
type engineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewengineLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *engineLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewengineLexer(input antlr.CharStream) *engineLexer {
	l := new(engineLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "engine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// engineLexer tokens.
const (
	engineLexerT__0 = 1
	engineLexerT__1 = 2
	engineLexerEQUAL = 3
	engineLexerNOT_EQUAL = 4
	engineLexerAND = 5
	engineLexerOR = 6
	engineLexerINTEGER = 7
	engineLexerIN = 8
	engineLexerSTRING = 9
	engineLexerVAR_DOT = 10
	engineLexerVAR = 11
	engineLexerDOT = 12
	engineLexerSL_COMMENT = 13
	engineLexerWS = 14
)

