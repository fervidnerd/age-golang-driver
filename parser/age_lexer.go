// Code generated from /home/harry/projects/age/drivers/golang/parser/Age.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type AgeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AgeLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agelexerLexerInit() {
	staticData := &AgeLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'['", "','", "']'", "'{'", "'}'", "':'", "'::vertex'", "'::edge'",
		"'::path'", "'::numeric'", "", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "KW_VERTEX", "KW_EDGE", "KW_PATH", "KW_NUMERIC",
		"STRING", "BOOL", "NULL", "NUMBER", "FLOAT_EXPR", "NUMERIC", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "KW_VERTEX", "KW_EDGE",
		"KW_PATH", "KW_NUMERIC", "STRING", "BOOL", "NULL", "ESC", "UNICODE",
		"HEX", "SAFECODEPOINT", "NUMBER", "FLOAT_EXPR", "NUMERIC", "INT", "EXP",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 210, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 96, 8, 10, 10, 10, 12, 10, 99,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 112, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 3, 13, 122, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 3, 17, 135, 8, 17, 1, 17, 1, 17,
		1, 17, 4, 17, 140, 8, 17, 11, 17, 12, 17, 141, 3, 17, 144, 8, 17, 1, 17,
		3, 17, 147, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 3, 18, 169, 8, 18, 1, 19, 3, 19, 172, 8, 19, 1, 19, 1, 19,
		1, 19, 4, 19, 177, 8, 19, 11, 19, 12, 19, 178, 3, 19, 181, 8, 19, 1, 19,
		3, 19, 184, 8, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 5, 20, 191, 8, 20,
		10, 20, 12, 20, 194, 9, 20, 3, 20, 196, 8, 20, 1, 21, 1, 21, 3, 21, 200,
		8, 21, 1, 21, 1, 21, 1, 22, 4, 22, 205, 8, 22, 11, 22, 12, 22, 206, 1,
		22, 1, 22, 0, 0, 23, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 0, 29, 0, 31, 0, 33, 0, 35,
		14, 37, 15, 39, 16, 41, 0, 43, 0, 45, 17, 1, 0, 8, 8, 0, 34, 34, 47, 47,
		92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65,
		70, 97, 102, 3, 0, 0, 31, 34, 34, 92, 92, 1, 0, 48, 57, 1, 0, 49, 57, 2,
		0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 3, 0, 9, 10, 13, 13, 32, 32,
		221, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0,
		39, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 1, 47, 1, 0, 0, 0, 3, 49, 1, 0, 0, 0,
		5, 51, 1, 0, 0, 0, 7, 53, 1, 0, 0, 0, 9, 55, 1, 0, 0, 0, 11, 57, 1, 0,
		0, 0, 13, 59, 1, 0, 0, 0, 15, 68, 1, 0, 0, 0, 17, 75, 1, 0, 0, 0, 19, 82,
		1, 0, 0, 0, 21, 92, 1, 0, 0, 0, 23, 111, 1, 0, 0, 0, 25, 113, 1, 0, 0,
		0, 27, 118, 1, 0, 0, 0, 29, 123, 1, 0, 0, 0, 31, 129, 1, 0, 0, 0, 33, 131,
		1, 0, 0, 0, 35, 134, 1, 0, 0, 0, 37, 168, 1, 0, 0, 0, 39, 171, 1, 0, 0,
		0, 41, 195, 1, 0, 0, 0, 43, 197, 1, 0, 0, 0, 45, 204, 1, 0, 0, 0, 47, 48,
		5, 91, 0, 0, 48, 2, 1, 0, 0, 0, 49, 50, 5, 44, 0, 0, 50, 4, 1, 0, 0, 0,
		51, 52, 5, 93, 0, 0, 52, 6, 1, 0, 0, 0, 53, 54, 5, 123, 0, 0, 54, 8, 1,
		0, 0, 0, 55, 56, 5, 125, 0, 0, 56, 10, 1, 0, 0, 0, 57, 58, 5, 58, 0, 0,
		58, 12, 1, 0, 0, 0, 59, 60, 5, 58, 0, 0, 60, 61, 5, 58, 0, 0, 61, 62, 5,
		118, 0, 0, 62, 63, 5, 101, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5, 116,
		0, 0, 65, 66, 5, 101, 0, 0, 66, 67, 5, 120, 0, 0, 67, 14, 1, 0, 0, 0, 68,
		69, 5, 58, 0, 0, 69, 70, 5, 58, 0, 0, 70, 71, 5, 101, 0, 0, 71, 72, 5,
		100, 0, 0, 72, 73, 5, 103, 0, 0, 73, 74, 5, 101, 0, 0, 74, 16, 1, 0, 0,
		0, 75, 76, 5, 58, 0, 0, 76, 77, 5, 58, 0, 0, 77, 78, 5, 112, 0, 0, 78,
		79, 5, 97, 0, 0, 79, 80, 5, 116, 0, 0, 80, 81, 5, 104, 0, 0, 81, 18, 1,
		0, 0, 0, 82, 83, 5, 58, 0, 0, 83, 84, 5, 58, 0, 0, 84, 85, 5, 110, 0, 0,
		85, 86, 5, 117, 0, 0, 86, 87, 5, 109, 0, 0, 87, 88, 5, 101, 0, 0, 88, 89,
		5, 114, 0, 0, 89, 90, 5, 105, 0, 0, 90, 91, 5, 99, 0, 0, 91, 20, 1, 0,
		0, 0, 92, 97, 5, 34, 0, 0, 93, 96, 3, 27, 13, 0, 94, 96, 3, 33, 16, 0,
		95, 93, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 100, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100,
		101, 5, 34, 0, 0, 101, 22, 1, 0, 0, 0, 102, 103, 5, 116, 0, 0, 103, 104,
		5, 114, 0, 0, 104, 105, 5, 117, 0, 0, 105, 112, 5, 101, 0, 0, 106, 107,
		5, 102, 0, 0, 107, 108, 5, 97, 0, 0, 108, 109, 5, 108, 0, 0, 109, 110,
		5, 115, 0, 0, 110, 112, 5, 101, 0, 0, 111, 102, 1, 0, 0, 0, 111, 106, 1,
		0, 0, 0, 112, 24, 1, 0, 0, 0, 113, 114, 5, 110, 0, 0, 114, 115, 5, 117,
		0, 0, 115, 116, 5, 108, 0, 0, 116, 117, 5, 108, 0, 0, 117, 26, 1, 0, 0,
		0, 118, 121, 5, 92, 0, 0, 119, 122, 7, 0, 0, 0, 120, 122, 3, 29, 14, 0,
		121, 119, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 28, 1, 0, 0, 0, 123, 124,
		5, 117, 0, 0, 124, 125, 3, 31, 15, 0, 125, 126, 3, 31, 15, 0, 126, 127,
		3, 31, 15, 0, 127, 128, 3, 31, 15, 0, 128, 30, 1, 0, 0, 0, 129, 130, 7,
		1, 0, 0, 130, 32, 1, 0, 0, 0, 131, 132, 8, 2, 0, 0, 132, 34, 1, 0, 0, 0,
		133, 135, 5, 45, 0, 0, 134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 143, 3, 41, 20, 0, 137, 139, 5, 46, 0, 0, 138, 140,
		7, 3, 0, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0,
		0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143, 137, 1, 0, 0, 0,
		143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 147, 3, 43, 21, 0, 146,
		145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 36, 1, 0, 0, 0, 148, 149, 5,
		78, 0, 0, 149, 150, 5, 97, 0, 0, 150, 169, 5, 78, 0, 0, 151, 152, 5, 45,
		0, 0, 152, 153, 5, 73, 0, 0, 153, 154, 5, 110, 0, 0, 154, 155, 5, 102,
		0, 0, 155, 156, 5, 105, 0, 0, 156, 157, 5, 110, 0, 0, 157, 158, 5, 105,
		0, 0, 158, 159, 5, 116, 0, 0, 159, 169, 5, 121, 0, 0, 160, 161, 5, 73,
		0, 0, 161, 162, 5, 110, 0, 0, 162, 163, 5, 102, 0, 0, 163, 164, 5, 105,
		0, 0, 164, 165, 5, 110, 0, 0, 165, 166, 5, 105, 0, 0, 166, 167, 5, 116,
		0, 0, 167, 169, 5, 121, 0, 0, 168, 148, 1, 0, 0, 0, 168, 151, 1, 0, 0,
		0, 168, 160, 1, 0, 0, 0, 169, 38, 1, 0, 0, 0, 170, 172, 5, 45, 0, 0, 171,
		170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 180,
		3, 41, 20, 0, 174, 176, 5, 46, 0, 0, 175, 177, 7, 3, 0, 0, 176, 175, 1,
		0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0,
		0, 179, 181, 1, 0, 0, 0, 180, 174, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		183, 1, 0, 0, 0, 182, 184, 3, 43, 21, 0, 183, 182, 1, 0, 0, 0, 183, 184,
		1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 3, 19, 9, 0, 186, 40, 1, 0,
		0, 0, 187, 196, 5, 48, 0, 0, 188, 192, 7, 4, 0, 0, 189, 191, 7, 3, 0, 0,
		190, 189, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 187,
		1, 0, 0, 0, 195, 188, 1, 0, 0, 0, 196, 42, 1, 0, 0, 0, 197, 199, 7, 5,
		0, 0, 198, 200, 7, 6, 0, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 201, 1, 0, 0, 0, 201, 202, 3, 41, 20, 0, 202, 44, 1, 0, 0, 0, 203,
		205, 7, 7, 0, 0, 204, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 204,
		1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 6, 22,
		0, 0, 209, 46, 1, 0, 0, 0, 18, 0, 95, 97, 111, 121, 134, 141, 143, 146,
		168, 171, 178, 180, 183, 192, 195, 199, 206, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AgeLexerInit initializes any static state used to implement AgeLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAgeLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgeLexerInit() {
	staticData := &AgeLexerLexerStaticData
	staticData.once.Do(agelexerLexerInit)
}

// NewAgeLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAgeLexer(input antlr.CharStream) *AgeLexer {
	AgeLexerInit()
	l := new(AgeLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AgeLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Age.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AgeLexer tokens.
const (
	AgeLexerT__0       = 1
	AgeLexerT__1       = 2
	AgeLexerT__2       = 3
	AgeLexerT__3       = 4
	AgeLexerT__4       = 5
	AgeLexerT__5       = 6
	AgeLexerKW_VERTEX  = 7
	AgeLexerKW_EDGE    = 8
	AgeLexerKW_PATH    = 9
	AgeLexerKW_NUMERIC = 10
	AgeLexerSTRING     = 11
	AgeLexerBOOL       = 12
	AgeLexerNULL       = 13
	AgeLexerNUMBER     = 14
	AgeLexerFLOAT_EXPR = 15
	AgeLexerNUMERIC    = 16
	AgeLexerWS         = 17
)