// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Server/interfaces"
import "Server/environment"
import "Server/expressions"
import "Server/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'",
		"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'while'", "'for'", "'in'", "'guard'", "'break'",
		"'continue'", "'return'", "'struct'", "'self'", "'mutating'", "'func'",
		"", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='",
		"'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
		"'{'", "'}'", "'?'", "'->'", "','", "'.'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL",
		"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE",
		"FOR", "IN", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", "SELF",
		"MUTATING", "FUNC", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD",
		"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"Q_MARK", "ARROW", "COMA", "PUNTO", "COLON", "WHITESPACE", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "ifstmt", "varstmt", "tipo",
		"varasgmt", "conststmt", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 216, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 4, 1, 26, 8, 1, 11, 1, 12, 1, 27, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 80, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		103, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 115, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 134, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 3, 8, 150, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 174, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 211, 8, 9, 10, 9, 12, 9, 214, 9, 9, 1, 9,
		0, 1, 18, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 5, 1, 0, 42, 44, 1,
		0, 45, 46, 2, 0, 38, 38, 40, 40, 2, 0, 39, 39, 41, 41, 1, 0, 32, 33, 235,
		0, 20, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6, 48, 1, 0, 0,
		0, 8, 79, 1, 0, 0, 0, 10, 102, 1, 0, 0, 0, 12, 114, 1, 0, 0, 0, 14, 133,
		1, 0, 0, 0, 16, 149, 1, 0, 0, 0, 18, 173, 1, 0, 0, 0, 20, 21, 3, 2, 1,
		0, 21, 22, 5, 0, 0, 1, 22, 23, 6, 0, -1, 0, 23, 1, 1, 0, 0, 0, 24, 26,
		3, 4, 2, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0,
		27, 28, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30, 6, 1, -1, 0, 30, 3, 1,
		0, 0, 0, 31, 32, 3, 6, 3, 0, 32, 33, 6, 2, -1, 0, 33, 47, 1, 0, 0, 0, 34,
		35, 3, 8, 4, 0, 35, 36, 6, 2, -1, 0, 36, 47, 1, 0, 0, 0, 37, 38, 3, 10,
		5, 0, 38, 39, 6, 2, -1, 0, 39, 47, 1, 0, 0, 0, 40, 41, 3, 14, 7, 0, 41,
		42, 6, 2, -1, 0, 42, 47, 1, 0, 0, 0, 43, 44, 3, 16, 8, 0, 44, 45, 6, 2,
		-1, 0, 45, 47, 1, 0, 0, 0, 46, 31, 1, 0, 0, 0, 46, 34, 1, 0, 0, 0, 46,
		37, 1, 0, 0, 0, 46, 40, 1, 0, 0, 0, 46, 43, 1, 0, 0, 0, 47, 5, 1, 0, 0,
		0, 48, 49, 5, 11, 0, 0, 49, 50, 5, 47, 0, 0, 50, 51, 3, 18, 9, 0, 51, 52,
		5, 48, 0, 0, 52, 53, 6, 3, -1, 0, 53, 7, 1, 0, 0, 0, 54, 55, 5, 12, 0,
		0, 55, 56, 3, 18, 9, 0, 56, 57, 5, 49, 0, 0, 57, 58, 3, 2, 1, 0, 58, 59,
		5, 50, 0, 0, 59, 60, 6, 4, -1, 0, 60, 80, 1, 0, 0, 0, 61, 62, 5, 12, 0,
		0, 62, 63, 3, 18, 9, 0, 63, 64, 5, 49, 0, 0, 64, 65, 3, 2, 1, 0, 65, 66,
		5, 50, 0, 0, 66, 67, 5, 13, 0, 0, 67, 68, 5, 49, 0, 0, 68, 69, 3, 2, 1,
		0, 69, 70, 5, 50, 0, 0, 70, 80, 1, 0, 0, 0, 71, 72, 5, 12, 0, 0, 72, 73,
		3, 18, 9, 0, 73, 74, 5, 49, 0, 0, 74, 75, 3, 2, 1, 0, 75, 76, 5, 50, 0,
		0, 76, 77, 5, 13, 0, 0, 77, 78, 3, 8, 4, 0, 78, 80, 1, 0, 0, 0, 79, 54,
		1, 0, 0, 0, 79, 61, 1, 0, 0, 0, 79, 71, 1, 0, 0, 0, 80, 9, 1, 0, 0, 0,
		81, 82, 5, 9, 0, 0, 82, 83, 5, 31, 0, 0, 83, 84, 5, 55, 0, 0, 84, 85, 3,
		12, 6, 0, 85, 86, 5, 37, 0, 0, 86, 87, 3, 18, 9, 0, 87, 88, 6, 5, -1, 0,
		88, 103, 1, 0, 0, 0, 89, 90, 5, 9, 0, 0, 90, 91, 5, 31, 0, 0, 91, 92, 5,
		37, 0, 0, 92, 93, 3, 18, 9, 0, 93, 94, 6, 5, -1, 0, 94, 103, 1, 0, 0, 0,
		95, 96, 5, 9, 0, 0, 96, 97, 5, 31, 0, 0, 97, 98, 5, 55, 0, 0, 98, 99, 3,
		12, 6, 0, 99, 100, 5, 51, 0, 0, 100, 101, 6, 5, -1, 0, 101, 103, 1, 0,
		0, 0, 102, 81, 1, 0, 0, 0, 102, 89, 1, 0, 0, 0, 102, 95, 1, 0, 0, 0, 103,
		11, 1, 0, 0, 0, 104, 105, 5, 1, 0, 0, 105, 115, 6, 6, -1, 0, 106, 107,
		5, 2, 0, 0, 107, 115, 6, 6, -1, 0, 108, 109, 5, 4, 0, 0, 109, 115, 6, 6,
		-1, 0, 110, 111, 5, 3, 0, 0, 111, 115, 6, 6, -1, 0, 112, 113, 5, 5, 0,
		0, 113, 115, 6, 6, -1, 0, 114, 104, 1, 0, 0, 0, 114, 106, 1, 0, 0, 0, 114,
		108, 1, 0, 0, 0, 114, 110, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 13, 1,
		0, 0, 0, 116, 117, 5, 31, 0, 0, 117, 118, 5, 37, 0, 0, 118, 119, 3, 18,
		9, 0, 119, 120, 6, 7, -1, 0, 120, 134, 1, 0, 0, 0, 121, 122, 5, 31, 0,
		0, 122, 123, 5, 45, 0, 0, 123, 124, 5, 37, 0, 0, 124, 125, 3, 18, 9, 0,
		125, 126, 6, 7, -1, 0, 126, 134, 1, 0, 0, 0, 127, 128, 5, 31, 0, 0, 128,
		129, 5, 46, 0, 0, 129, 130, 5, 37, 0, 0, 130, 131, 3, 18, 9, 0, 131, 132,
		6, 7, -1, 0, 132, 134, 1, 0, 0, 0, 133, 116, 1, 0, 0, 0, 133, 121, 1, 0,
		0, 0, 133, 127, 1, 0, 0, 0, 134, 15, 1, 0, 0, 0, 135, 136, 5, 10, 0, 0,
		136, 137, 5, 31, 0, 0, 137, 138, 5, 37, 0, 0, 138, 139, 3, 18, 9, 0, 139,
		140, 6, 8, -1, 0, 140, 150, 1, 0, 0, 0, 141, 142, 5, 10, 0, 0, 142, 143,
		5, 31, 0, 0, 143, 144, 5, 55, 0, 0, 144, 145, 3, 12, 6, 0, 145, 146, 5,
		37, 0, 0, 146, 147, 3, 18, 9, 0, 147, 148, 6, 8, -1, 0, 148, 150, 1, 0,
		0, 0, 149, 135, 1, 0, 0, 0, 149, 141, 1, 0, 0, 0, 150, 17, 1, 0, 0, 0,
		151, 152, 6, 9, -1, 0, 152, 153, 5, 46, 0, 0, 153, 154, 3, 18, 9, 15, 154,
		155, 6, 9, -1, 0, 155, 174, 1, 0, 0, 0, 156, 157, 5, 47, 0, 0, 157, 158,
		3, 18, 9, 0, 158, 159, 5, 48, 0, 0, 159, 160, 6, 9, -1, 0, 160, 174, 1,
		0, 0, 0, 161, 162, 5, 28, 0, 0, 162, 174, 6, 9, -1, 0, 163, 164, 5, 29,
		0, 0, 164, 174, 6, 9, -1, 0, 165, 166, 5, 30, 0, 0, 166, 174, 6, 9, -1,
		0, 167, 168, 5, 6, 0, 0, 168, 174, 6, 9, -1, 0, 169, 170, 5, 7, 0, 0, 170,
		174, 6, 9, -1, 0, 171, 172, 5, 31, 0, 0, 172, 174, 6, 9, -1, 0, 173, 151,
		1, 0, 0, 0, 173, 156, 1, 0, 0, 0, 173, 161, 1, 0, 0, 0, 173, 163, 1, 0,
		0, 0, 173, 165, 1, 0, 0, 0, 173, 167, 1, 0, 0, 0, 173, 169, 1, 0, 0, 0,
		173, 171, 1, 0, 0, 0, 174, 212, 1, 0, 0, 0, 175, 176, 10, 14, 0, 0, 176,
		177, 7, 0, 0, 0, 177, 178, 3, 18, 9, 15, 178, 179, 6, 9, -1, 0, 179, 211,
		1, 0, 0, 0, 180, 181, 10, 13, 0, 0, 181, 182, 7, 1, 0, 0, 182, 183, 3,
		18, 9, 14, 183, 184, 6, 9, -1, 0, 184, 211, 1, 0, 0, 0, 185, 186, 10, 12,
		0, 0, 186, 187, 7, 2, 0, 0, 187, 188, 3, 18, 9, 13, 188, 189, 6, 9, -1,
		0, 189, 211, 1, 0, 0, 0, 190, 191, 10, 11, 0, 0, 191, 192, 7, 3, 0, 0,
		192, 193, 3, 18, 9, 12, 193, 194, 6, 9, -1, 0, 194, 211, 1, 0, 0, 0, 195,
		196, 10, 10, 0, 0, 196, 197, 7, 4, 0, 0, 197, 198, 3, 18, 9, 11, 198, 199,
		6, 9, -1, 0, 199, 211, 1, 0, 0, 0, 200, 201, 10, 9, 0, 0, 201, 202, 5,
		36, 0, 0, 202, 203, 3, 18, 9, 10, 203, 204, 6, 9, -1, 0, 204, 211, 1, 0,
		0, 0, 205, 206, 10, 8, 0, 0, 206, 207, 5, 35, 0, 0, 207, 208, 3, 18, 9,
		9, 208, 209, 6, 9, -1, 0, 209, 211, 1, 0, 0, 0, 210, 175, 1, 0, 0, 0, 210,
		180, 1, 0, 0, 0, 210, 185, 1, 0, 0, 0, 210, 190, 1, 0, 0, 0, 210, 195,
		1, 0, 0, 0, 210, 200, 1, 0, 0, 0, 210, 205, 1, 0, 0, 0, 211, 214, 1, 0,
		0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 19, 1, 0, 0, 0,
		214, 212, 1, 0, 0, 0, 10, 27, 46, 79, 102, 114, 133, 149, 173, 210, 212,
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

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserINT          = 1
	SwiftGrammarParserFLOAT        = 2
	SwiftGrammarParserBOOL         = 3
	SwiftGrammarParserSTR          = 4
	SwiftGrammarParserCHARACTER    = 5
	SwiftGrammarParserTRU          = 6
	SwiftGrammarParserFAL          = 7
	SwiftGrammarParserNIL          = 8
	SwiftGrammarParserVAR          = 9
	SwiftGrammarParserLET          = 10
	SwiftGrammarParserPRINT        = 11
	SwiftGrammarParserIF           = 12
	SwiftGrammarParserELSE         = 13
	SwiftGrammarParserSWITCH       = 14
	SwiftGrammarParserCASE         = 15
	SwiftGrammarParserDEFAULT      = 16
	SwiftGrammarParserWHILE        = 17
	SwiftGrammarParserFOR          = 18
	SwiftGrammarParserIN           = 19
	SwiftGrammarParserGUARD        = 20
	SwiftGrammarParserBREAK        = 21
	SwiftGrammarParserCONTINUE     = 22
	SwiftGrammarParserRETURN       = 23
	SwiftGrammarParserSTRUCT       = 24
	SwiftGrammarParserSELF         = 25
	SwiftGrammarParserMUTATING     = 26
	SwiftGrammarParserFUNC         = 27
	SwiftGrammarParserNUMBER       = 28
	SwiftGrammarParserCHAR         = 29
	SwiftGrammarParserSTRING       = 30
	SwiftGrammarParserID           = 31
	SwiftGrammarParserDIF          = 32
	SwiftGrammarParserIG_IG        = 33
	SwiftGrammarParserNOT          = 34
	SwiftGrammarParserOR           = 35
	SwiftGrammarParserAND          = 36
	SwiftGrammarParserIG           = 37
	SwiftGrammarParserMAY_IG       = 38
	SwiftGrammarParserMEN_IG       = 39
	SwiftGrammarParserMAYOR        = 40
	SwiftGrammarParserMENOR        = 41
	SwiftGrammarParserMOD          = 42
	SwiftGrammarParserMUL          = 43
	SwiftGrammarParserDIV          = 44
	SwiftGrammarParserADD          = 45
	SwiftGrammarParserSUB          = 46
	SwiftGrammarParserPARIZQ       = 47
	SwiftGrammarParserPARDER       = 48
	SwiftGrammarParserLLAVEIZQ     = 49
	SwiftGrammarParserLLAVEDER     = 50
	SwiftGrammarParserQ_MARK       = 51
	SwiftGrammarParserARROW        = 52
	SwiftGrammarParserCOMA         = 53
	SwiftGrammarParserPUNTO        = 54
	SwiftGrammarParserCOLON        = 55
	SwiftGrammarParserWHITESPACE   = 56
	SwiftGrammarParserCOMMENT      = 57
	SwiftGrammarParserLINE_COMMENT = 58
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s           = 0
	SwiftGrammarParserRULE_block       = 1
	SwiftGrammarParserRULE_instruction = 2
	SwiftGrammarParserRULE_printstmt   = 3
	SwiftGrammarParserRULE_ifstmt      = 4
	SwiftGrammarParserRULE_varstmt     = 5
	SwiftGrammarParserRULE_tipo        = 6
	SwiftGrammarParserRULE_varasgmt    = 7
	SwiftGrammarParserRULE_conststmt   = 8
	SwiftGrammarParserRULE_expr        = 9
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(21)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2147491328) != 0) {
		{
			p.SetState(24)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_varstmt returns the _varstmt rule contexts.
	Get_varstmt() IVarstmtContext

	// Get_varasgmt returns the _varasgmt rule contexts.
	Get_varasgmt() IVarasgmtContext

	// Get_conststmt returns the _conststmt rule contexts.
	Get_conststmt() IConststmtContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_varstmt sets the _varstmt rule contexts.
	Set_varstmt(IVarstmtContext)

	// Set_varasgmt sets the _varasgmt rule contexts.
	Set_varasgmt(IVarasgmtContext)

	// Set_conststmt sets the _conststmt rule contexts.
	Set_conststmt(IConststmtContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	Ifstmt() IIfstmtContext
	Varstmt() IVarstmtContext
	Varasgmt() IVarasgmtContext
	Conststmt() IConststmtContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_printstmt IPrintstmtContext
	_ifstmt    IIfstmtContext
	_varstmt   IVarstmtContext
	_varasgmt  IVarasgmtContext
	_conststmt IConststmtContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_varstmt() IVarstmtContext { return s._varstmt }

func (s *InstructionContext) Get_varasgmt() IVarasgmtContext { return s._varasgmt }

func (s *InstructionContext) Get_conststmt() IConststmtContext { return s._conststmt }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_varstmt(v IVarstmtContext) { s._varstmt = v }

func (s *InstructionContext) Set_varasgmt(v IVarasgmtContext) { s._varasgmt = v }

func (s *InstructionContext) Set_conststmt(v IConststmtContext) { s._conststmt = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Varstmt() IVarstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarstmtContext)
}

func (s *InstructionContext) Varasgmt() IVarasgmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarasgmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarasgmtContext)
}

func (s *InstructionContext) Conststmt() IConststmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConststmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConststmtContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case SwiftGrammarParserIF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinstr()

	case SwiftGrammarParserVAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)

			var _x = p.Varstmt()

			localctx.(*InstructionContext)._varstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_varstmt().GetVar_()

	case SwiftGrammarParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)

			var _x = p.Varasgmt()

			localctx.(*InstructionContext)._varasgmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_varasgmt().GetAsgmt()

	case SwiftGrammarParserLET:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(43)

			var _x = p.Conststmt()

			localctx.(*InstructionContext)._conststmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_conststmt().GetConst_()

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	prnt   interfaces.Instruction
	_PRINT antlr.Token
	_expr  IExprContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_expr() IExprContext { return s._expr }

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(51)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
		}
	}()), localctx.(*PrintstmtContext).Get_expr().GetE())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetIfinstr returns the ifinstr attribute.
	GetIfinstr() interfaces.Instruction

	// SetIfinstr sets the ifinstr attribute.
	SetIfinstr(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Ifstmt() IIfstmtContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	ifinstr interfaces.Instruction
	_IF     antlr.Token
	_expr   IExprContext
	_block  IBlockContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_IF() antlr.Token { return s._IF }

func (s *IfstmtContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *IfstmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfstmtContext) Get_block() IBlockContext { return s._block }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfstmtContext) GetIfinstr() interfaces.Instruction { return s.ifinstr }

func (s *IfstmtContext) SetIfinstr(v interfaces.Instruction) { s.ifinstr = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEIZQ)
}

func (s *IfstmtContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEDER)
}

func (s *IfstmtContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, i)
}

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *IfstmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_ifstmt)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(56)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(58)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinstr = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(SwiftGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)
			p.expr(0)
		}
		{
			p.SetState(63)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Block()
		}
		{
			p.SetState(65)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Block()
		}
		{
			p.SetState(69)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Match(SwiftGrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.expr(0)
		}
		{
			p.SetState(73)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Block()
		}
		{
			p.SetState(75)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Ifstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarstmtContext is an interface to support dynamic dispatch.
type IVarstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_Q_MARK returns the _Q_MARK token.
	Get_Q_MARK() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_Q_MARK sets the _Q_MARK token.
	Set_Q_MARK(antlr.Token)

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetVar_ returns the var_ attribute.
	GetVar_() interfaces.Instruction

	// SetVar_ sets the var_ attribute.
	SetVar_(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Tipo() ITipoContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	Q_MARK() antlr.TerminalNode

	// IsVarstmtContext differentiates from other interfaces.
	IsVarstmtContext()
}

type VarstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	var_    interfaces.Instruction
	_VAR    antlr.Token
	_ID     antlr.Token
	_tipo   ITipoContext
	_expr   IExprContext
	_Q_MARK antlr.Token
}

func NewEmptyVarstmtContext() *VarstmtContext {
	var p = new(VarstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_varstmt
	return p
}

func InitEmptyVarstmtContext(p *VarstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_varstmt
}

func (*VarstmtContext) IsVarstmtContext() {}

func NewVarstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarstmtContext {
	var p = new(VarstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_varstmt

	return p
}

func (s *VarstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarstmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *VarstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *VarstmtContext) Get_Q_MARK() antlr.Token { return s._Q_MARK }

func (s *VarstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *VarstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VarstmtContext) Set_Q_MARK(v antlr.Token) { s._Q_MARK = v }

func (s *VarstmtContext) Get_tipo() ITipoContext { return s._tipo }

func (s *VarstmtContext) Get_expr() IExprContext { return s._expr }

func (s *VarstmtContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *VarstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VarstmtContext) GetVar_() interfaces.Instruction { return s.var_ }

func (s *VarstmtContext) SetVar_(v interfaces.Instruction) { s.var_ = v }

func (s *VarstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *VarstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VarstmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *VarstmtContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VarstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VarstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarstmtContext) Q_MARK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserQ_MARK, 0)
}

func (s *VarstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVarstmt(s)
	}
}

func (s *VarstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVarstmt(s)
	}
}

func (p *SwiftGrammarParser) Varstmt() (localctx IVarstmtContext) {
	localctx = NewVarstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_varstmt)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)

			var _x = p.Tipo()

			localctx.(*VarstmtContext)._tipo = _x
		}
		{
			p.SetState(85)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)

			var _x = p.expr(0)

			localctx.(*VarstmtContext)._expr = _x
		}
		localctx.(*VarstmtContext).var_ = instructions.NewStmt((func() int {
			if localctx.(*VarstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VarstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VarstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VarstmtContext).Get_tipo().GetRtipo(), localctx.(*VarstmtContext).Get_expr().GetE(), false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)

			var _x = p.expr(0)

			localctx.(*VarstmtContext)._expr = _x
		}
		localctx.(*VarstmtContext).var_ = instructions.NewStmt((func() int {
			if localctx.(*VarstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VarstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VarstmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*VarstmtContext).Get_expr().GetE(), false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)

			var _x = p.Tipo()

			localctx.(*VarstmtContext)._tipo = _x
		}
		{
			p.SetState(99)

			var _m = p.Match(SwiftGrammarParserQ_MARK)

			localctx.(*VarstmtContext)._Q_MARK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VarstmtContext).var_ = instructions.NewStmt((func() int {
			if localctx.(*VarstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VarstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VarstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VarstmtContext).Get_tipo().GetRtipo(), expressions.NewPrimitive((func() int {
			if localctx.(*VarstmtContext).Get_Q_MARK() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_Q_MARK().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarstmtContext).Get_Q_MARK() == nil {
				return 0
			} else {
				return localctx.(*VarstmtContext).Get_Q_MARK().GetColumn()
			}
		}()), environment.WILDCARD, environment.WILDCARD), false)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRtipo returns the rtipo attribute.
	GetRtipo() environment.TipoExpresion

	// SetRtipo sets the rtipo attribute.
	SetRtipo(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STR() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	rtipo  environment.TipoExpresion
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) GetRtipo() environment.TipoExpresion { return s.rtipo }

func (s *TipoContext) SetRtipo(v environment.TipoExpresion) { s.rtipo = v }

func (s *TipoContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TipoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TipoContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *TipoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TipoContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHARACTER, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *SwiftGrammarParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_tipo)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipoContext).rtipo = 0

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipoContext).rtipo = 1

	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(108)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipoContext).rtipo = 2

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipoContext).rtipo = 3

	case SwiftGrammarParserCHARACTER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Match(SwiftGrammarParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TipoContext).rtipo = 5

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarasgmtContext is an interface to support dynamic dispatch.
type IVarasgmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAsgmt returns the asgmt attribute.
	GetAsgmt() interfaces.Instruction

	// SetAsgmt sets the asgmt attribute.
	SetAsgmt(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsVarasgmtContext differentiates from other interfaces.
	IsVarasgmtContext()
}

type VarasgmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	asgmt  interfaces.Instruction
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyVarasgmtContext() *VarasgmtContext {
	var p = new(VarasgmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_varasgmt
	return p
}

func InitEmptyVarasgmtContext(p *VarasgmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_varasgmt
}

func (*VarasgmtContext) IsVarasgmtContext() {}

func NewVarasgmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarasgmtContext {
	var p = new(VarasgmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_varasgmt

	return p
}

func (s *VarasgmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VarasgmtContext) Get_ID() antlr.Token { return s._ID }

func (s *VarasgmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VarasgmtContext) Get_expr() IExprContext { return s._expr }

func (s *VarasgmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VarasgmtContext) GetAsgmt() interfaces.Instruction { return s.asgmt }

func (s *VarasgmtContext) SetAsgmt(v interfaces.Instruction) { s.asgmt = v }

func (s *VarasgmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VarasgmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VarasgmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarasgmtContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *VarasgmtContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *VarasgmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarasgmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarasgmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVarasgmt(s)
	}
}

func (s *VarasgmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVarasgmt(s)
	}
}

func (p *SwiftGrammarParser) Varasgmt() (localctx IVarasgmtContext) {
	localctx = NewVarasgmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_varasgmt)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.expr(0)

			localctx.(*VarasgmtContext)._expr = _x
		}
		localctx.(*VarasgmtContext).asgmt = instructions.NewAsgmt((func() int {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VarasgmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(SwiftGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)

			var _x = p.expr(0)

			localctx.(*VarasgmtContext)._expr = _x
		}
		localctx.(*VarasgmtContext).asgmt = instructions.NewOpAsgmt((func() int {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VarasgmtContext).Get_expr().GetE(), "+")

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)

			var _x = p.expr(0)

			localctx.(*VarasgmtContext)._expr = _x
		}
		localctx.(*VarasgmtContext).asgmt = instructions.NewOpAsgmt((func() int {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VarasgmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VarasgmtContext).Get_ID().GetText()
			}
		}()), localctx.(*VarasgmtContext).Get_expr().GetE(), "-")

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConststmtContext is an interface to support dynamic dispatch.
type IConststmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// GetConst_ returns the const_ attribute.
	GetConst_() interfaces.Instruction

	// SetConst_ sets the const_ attribute.
	SetConst_(interfaces.Instruction)

	// Getter signatures
	LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	Tipo() ITipoContext

	// IsConststmtContext differentiates from other interfaces.
	IsConststmtContext()
}

type ConststmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	const_ interfaces.Instruction
	_LET   antlr.Token
	_ID    antlr.Token
	_expr  IExprContext
	_tipo  ITipoContext
}

func NewEmptyConststmtContext() *ConststmtContext {
	var p = new(ConststmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_conststmt
	return p
}

func InitEmptyConststmtContext(p *ConststmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_conststmt
}

func (*ConststmtContext) IsConststmtContext() {}

func NewConststmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConststmtContext {
	var p = new(ConststmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_conststmt

	return p
}

func (s *ConststmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConststmtContext) Get_LET() antlr.Token { return s._LET }

func (s *ConststmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ConststmtContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *ConststmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ConststmtContext) Get_expr() IExprContext { return s._expr }

func (s *ConststmtContext) Get_tipo() ITipoContext { return s._tipo }

func (s *ConststmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ConststmtContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *ConststmtContext) GetConst_() interfaces.Instruction { return s.const_ }

func (s *ConststmtContext) SetConst_(v interfaces.Instruction) { s.const_ = v }

func (s *ConststmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *ConststmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ConststmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *ConststmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConststmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *ConststmtContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *ConststmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConststmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConststmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterConststmt(s)
	}
}

func (s *ConststmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitConststmt(s)
	}
}

func (p *SwiftGrammarParser) Conststmt() (localctx IConststmtContext) {
	localctx = NewConststmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_conststmt)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*ConststmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConststmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)

			var _x = p.expr(0)

			localctx.(*ConststmtContext)._expr = _x
		}
		localctx.(*ConststmtContext).const_ = instructions.NewStmt((func() int {
			if localctx.(*ConststmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConststmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConststmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConststmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ConststmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ConststmtContext).Get_ID().GetText()
			}
		}()), environment.NULL, localctx.(*ConststmtContext).Get_expr().GetE(), true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*ConststmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConststmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)

			var _x = p.Tipo()

			localctx.(*ConststmtContext)._tipo = _x
		}
		{
			p.SetState(145)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)

			var _x = p.expr(0)

			localctx.(*ConststmtContext)._expr = _x
		}
		localctx.(*ConststmtContext).const_ = instructions.NewStmt((func() int {
			if localctx.(*ConststmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConststmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConststmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*ConststmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ConststmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ConststmtContext).Get_ID().GetText()
			}
		}()), localctx.(*ConststmtContext).Get_tipo().GetRtipo(), localctx.(*ConststmtContext).Get_expr().GetE(), true)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	SUB() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	ID() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	e       interfaces.Expression
	left    IExprContext
	op      antlr.Token
	right   IExprContext
	_expr   IExprContext
	_NUMBER antlr.Token
	_CHAR   antlr.Token
	_STRING antlr.Token
	_TRU    antlr.Token
	_FAL    antlr.Token
	_ID     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCHAR, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMOD, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserSUB:
		{
			p.SetState(152)

			var _m = p.Match(SwiftGrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _x = p.expr(15)

			localctx.(*ExprContext).right = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
			if localctx.(*ExprContext).GetRight() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetRight().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetRight() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetRight().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetRight().GetE(), "unario", expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).GetOp().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).GetOp().GetColumn()
			}
		}()), -1, environment.INTEGER))

	case SwiftGrammarParserPARIZQ:
		{
			p.SetState(156)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(158)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case SwiftGrammarParserNUMBER:
		{
			p.SetState(161)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case SwiftGrammarParserCHAR:
		{
			p.SetState(163)

			var _m = p.Match(SwiftGrammarParserCHAR)

			localctx.(*ExprContext)._CHAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_CHAR().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CHAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CHAR().GetColumn()
			}
		}()), str[1:len(str)-1], environment.CHAR)

	case SwiftGrammarParserSTRING:
		{
			p.SetState(165)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case SwiftGrammarParserTRU:
		{
			p.SetState(167)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*ExprContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case SwiftGrammarParserFAL:
		{
			p.SetState(169)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*ExprContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case SwiftGrammarParserID:
		{
			p.SetState(171)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewVar((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(176)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(177)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(181)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(182)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(185)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(186)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(187)

					var _x = p.expr(13)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(191)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(192)

					var _x = p.expr(12)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(196)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(197)

					var _x = p.expr(11)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(201)
					p.Match(SwiftGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(202)

					var _x = p.expr(10)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(205)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(206)
					p.Match(SwiftGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(207)

					var _x = p.expr(9)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
