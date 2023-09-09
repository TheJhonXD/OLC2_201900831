// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'",
		"'break'", "'continue'", "'return'", "'struct'", "'self'", "'mutating'",
		"'func'", "'append'", "'removeLast'", "'remove'", "'isEmpty'", "'count'",
		"'inout'", "'_'", "", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'",
		"'='", "'>='", "'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "'?'", "'->'", "','", "'.'",
		"':'", "'&'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL",
		"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE",
		"FOR", "IN", "RANGEPTS", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT",
		"SELF", "MUTATING", "FUNC", "APPEND", "REMOVELAST", "REMOVE", "EMPTY",
		"COUNT", "INOUT", "CAME", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG",
		"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD",
		"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"CORCHIZQ", "CORCHDER", "Q_MARK", "ARROW", "COMA", "PUNTO", "COLON",
		"AMP", "PUNTOCOMA", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "printexprlist", "ifstmt",
		"elseifstmt", "elsestmt", "varstmt", "tipo", "varasgmt", "conststmt",
		"switchstmt", "casestmt", "defaultstmt", "whilestmt", "forstmt", "guardstmt",
		"transferstmt", "vectorstmt", "definestmt", "listexpr", "methodvec",
		"methodvecrtrn", "vecaccess", "funcstmt", "listparam", "callfunc", "callfuncinstruction",
		"listparamcall", "funcembed", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 728, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 70, 8, 1, 11, 1, 12, 1, 71, 1,
		1, 1, 1, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 84,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 96, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 102, 8, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 108, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 114, 8, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 120, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 126, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 132, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		138, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 144, 8, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 150, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 156, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 162, 8, 2, 1, 2, 1, 2, 3, 2, 166, 8, 2, 1, 3, 1, 3, 1,
		3, 4, 3, 171, 8, 3, 11, 3, 12, 3, 172, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 185, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 200, 8, 5, 11,
		5, 12, 5, 201, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4,
		5, 213, 8, 5, 11, 5, 12, 5, 214, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 227, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 264, 8, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 276, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 295, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		3, 11, 311, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 4, 12, 317, 8, 12, 11, 12,
		12, 12, 318, 1, 12, 3, 12, 322, 8, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 332, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 340, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 380, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 400, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 4, 20, 414, 8, 20, 11, 20, 12, 20, 415,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 425, 8, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 434, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 458, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 474, 8, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 495, 8, 24,
		1, 25, 1, 25, 1, 25, 1, 25, 4, 25, 501, 8, 25, 11, 25, 12, 25, 502, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 4, 25, 517, 8, 25, 11, 25, 12, 25, 518, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 3, 25, 547, 8, 25, 1, 26, 3, 26, 550, 8, 26, 1, 26, 1, 26, 1,
		26, 3, 26, 555, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 562, 8,
		26, 1, 26, 1, 26, 1, 26, 3, 26, 567, 8, 26, 1, 26, 1, 26, 1, 26, 3, 26,
		572, 8, 26, 1, 27, 1, 27, 1, 27, 4, 27, 577, 8, 27, 11, 27, 12, 27, 578,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 588, 8, 27, 1,
		28, 1, 28, 1, 28, 4, 28, 593, 8, 28, 11, 28, 12, 28, 594, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 604, 8, 28, 1, 29, 1, 29, 3,
		29, 608, 8, 29, 1, 29, 3, 29, 611, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 619, 8, 29, 1, 29, 3, 29, 622, 8, 29, 1, 29, 1, 29, 1,
		29, 3, 29, 627, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 3, 30, 647, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 686,
		8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 723, 8, 31, 10, 31, 12, 31, 726,
		9, 31, 1, 31, 0, 1, 62, 32, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 0, 6, 2, 0, 35, 35, 39, 39, 1, 0, 50, 52, 1, 0, 53, 54, 2, 0, 46,
		46, 48, 48, 2, 0, 47, 47, 49, 49, 1, 0, 40, 41, 798, 0, 64, 1, 0, 0, 0,
		2, 69, 1, 0, 0, 0, 4, 165, 1, 0, 0, 0, 6, 167, 1, 0, 0, 0, 8, 184, 1, 0,
		0, 0, 10, 226, 1, 0, 0, 0, 12, 228, 1, 0, 0, 0, 14, 236, 1, 0, 0, 0, 16,
		263, 1, 0, 0, 0, 18, 275, 1, 0, 0, 0, 20, 294, 1, 0, 0, 0, 22, 310, 1,
		0, 0, 0, 24, 312, 1, 0, 0, 0, 26, 326, 1, 0, 0, 0, 28, 335, 1, 0, 0, 0,
		30, 343, 1, 0, 0, 0, 32, 379, 1, 0, 0, 0, 34, 381, 1, 0, 0, 0, 36, 399,
		1, 0, 0, 0, 38, 401, 1, 0, 0, 0, 40, 424, 1, 0, 0, 0, 42, 433, 1, 0, 0,
		0, 44, 457, 1, 0, 0, 0, 46, 473, 1, 0, 0, 0, 48, 494, 1, 0, 0, 0, 50, 546,
		1, 0, 0, 0, 52, 571, 1, 0, 0, 0, 54, 587, 1, 0, 0, 0, 56, 603, 1, 0, 0,
		0, 58, 626, 1, 0, 0, 0, 60, 646, 1, 0, 0, 0, 62, 685, 1, 0, 0, 0, 64, 65,
		3, 2, 1, 0, 65, 66, 5, 0, 0, 1, 66, 67, 6, 0, -1, 0, 67, 1, 1, 0, 0, 0,
		68, 70, 3, 4, 2, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 6, 1, -1, 0, 74,
		3, 1, 0, 0, 0, 75, 77, 3, 6, 3, 0, 76, 78, 5, 67, 0, 0, 77, 76, 1, 0, 0,
		0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 6, 2, -1, 0, 80, 166,
		1, 0, 0, 0, 81, 83, 3, 56, 28, 0, 82, 84, 5, 67, 0, 0, 83, 82, 1, 0, 0,
		0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 6, 2, -1, 0, 86, 166,
		1, 0, 0, 0, 87, 89, 3, 10, 5, 0, 88, 90, 5, 67, 0, 0, 89, 88, 1, 0, 0,
		0, 89, 90, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 6, 2, -1, 0, 92, 166,
		1, 0, 0, 0, 93, 95, 3, 16, 8, 0, 94, 96, 5, 67, 0, 0, 95, 94, 1, 0, 0,
		0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 6, 2, -1, 0, 98, 166,
		1, 0, 0, 0, 99, 101, 3, 20, 10, 0, 100, 102, 5, 67, 0, 0, 101, 100, 1,
		0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 6, 2, -1,
		0, 104, 166, 1, 0, 0, 0, 105, 107, 3, 22, 11, 0, 106, 108, 5, 67, 0, 0,
		107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		110, 6, 2, -1, 0, 110, 166, 1, 0, 0, 0, 111, 113, 3, 24, 12, 0, 112, 114,
		5, 67, 0, 0, 113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 116, 6, 2, -1, 0, 116, 166, 1, 0, 0, 0, 117, 119, 3, 30, 15,
		0, 118, 120, 5, 67, 0, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 122, 6, 2, -1, 0, 122, 166, 1, 0, 0, 0, 123, 125,
		3, 32, 16, 0, 124, 126, 5, 67, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 6, 2, -1, 0, 128, 166, 1, 0, 0,
		0, 129, 131, 3, 34, 17, 0, 130, 132, 5, 67, 0, 0, 131, 130, 1, 0, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134, 6, 2, -1, 0, 134,
		166, 1, 0, 0, 0, 135, 137, 3, 36, 18, 0, 136, 138, 5, 67, 0, 0, 137, 136,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 6, 2,
		-1, 0, 140, 166, 1, 0, 0, 0, 141, 143, 3, 38, 19, 0, 142, 144, 5, 67, 0,
		0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		146, 6, 2, -1, 0, 146, 166, 1, 0, 0, 0, 147, 149, 3, 44, 22, 0, 148, 150,
		5, 67, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 152, 6, 2, -1, 0, 152, 166, 1, 0, 0, 0, 153, 155, 3, 48, 24,
		0, 154, 156, 5, 67, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156,
		157, 1, 0, 0, 0, 157, 158, 6, 2, -1, 0, 158, 166, 1, 0, 0, 0, 159, 161,
		3, 50, 25, 0, 160, 162, 5, 67, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1,
		0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 6, 2, -1, 0, 164, 166, 1, 0, 0,
		0, 165, 75, 1, 0, 0, 0, 165, 81, 1, 0, 0, 0, 165, 87, 1, 0, 0, 0, 165,
		93, 1, 0, 0, 0, 165, 99, 1, 0, 0, 0, 165, 105, 1, 0, 0, 0, 165, 111, 1,
		0, 0, 0, 165, 117, 1, 0, 0, 0, 165, 123, 1, 0, 0, 0, 165, 129, 1, 0, 0,
		0, 165, 135, 1, 0, 0, 0, 165, 141, 1, 0, 0, 0, 165, 147, 1, 0, 0, 0, 165,
		153, 1, 0, 0, 0, 165, 159, 1, 0, 0, 0, 166, 5, 1, 0, 0, 0, 167, 168, 5,
		11, 0, 0, 168, 170, 5, 55, 0, 0, 169, 171, 3, 8, 4, 0, 170, 169, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 175, 5, 56, 0, 0, 175, 176, 6, 3, -1, 0, 176,
		7, 1, 0, 0, 0, 177, 178, 3, 62, 31, 0, 178, 179, 5, 63, 0, 0, 179, 180,
		6, 4, -1, 0, 180, 185, 1, 0, 0, 0, 181, 182, 3, 62, 31, 0, 182, 183, 6,
		4, -1, 0, 183, 185, 1, 0, 0, 0, 184, 177, 1, 0, 0, 0, 184, 181, 1, 0, 0,
		0, 185, 9, 1, 0, 0, 0, 186, 187, 5, 12, 0, 0, 187, 188, 3, 62, 31, 0, 188,
		189, 5, 57, 0, 0, 189, 190, 3, 2, 1, 0, 190, 191, 5, 58, 0, 0, 191, 192,
		6, 5, -1, 0, 192, 227, 1, 0, 0, 0, 193, 194, 5, 12, 0, 0, 194, 195, 3,
		62, 31, 0, 195, 196, 5, 57, 0, 0, 196, 197, 3, 2, 1, 0, 197, 199, 5, 58,
		0, 0, 198, 200, 3, 12, 6, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		204, 3, 14, 7, 0, 204, 205, 6, 5, -1, 0, 205, 227, 1, 0, 0, 0, 206, 207,
		5, 12, 0, 0, 207, 208, 3, 62, 31, 0, 208, 209, 5, 57, 0, 0, 209, 210, 3,
		2, 1, 0, 210, 212, 5, 58, 0, 0, 211, 213, 3, 12, 6, 0, 212, 211, 1, 0,
		0, 0, 213, 214, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0,
		215, 216, 1, 0, 0, 0, 216, 217, 6, 5, -1, 0, 217, 227, 1, 0, 0, 0, 218,
		219, 5, 12, 0, 0, 219, 220, 3, 62, 31, 0, 220, 221, 5, 57, 0, 0, 221, 222,
		3, 2, 1, 0, 222, 223, 5, 58, 0, 0, 223, 224, 3, 14, 7, 0, 224, 225, 6,
		5, -1, 0, 225, 227, 1, 0, 0, 0, 226, 186, 1, 0, 0, 0, 226, 193, 1, 0, 0,
		0, 226, 206, 1, 0, 0, 0, 226, 218, 1, 0, 0, 0, 227, 11, 1, 0, 0, 0, 228,
		229, 5, 13, 0, 0, 229, 230, 5, 12, 0, 0, 230, 231, 3, 62, 31, 0, 231, 232,
		5, 57, 0, 0, 232, 233, 3, 2, 1, 0, 233, 234, 5, 58, 0, 0, 234, 235, 6,
		6, -1, 0, 235, 13, 1, 0, 0, 0, 236, 237, 5, 13, 0, 0, 237, 238, 5, 57,
		0, 0, 238, 239, 3, 2, 1, 0, 239, 240, 5, 58, 0, 0, 240, 241, 6, 7, -1,
		0, 241, 15, 1, 0, 0, 0, 242, 243, 5, 9, 0, 0, 243, 244, 5, 39, 0, 0, 244,
		245, 5, 65, 0, 0, 245, 246, 3, 18, 9, 0, 246, 247, 5, 45, 0, 0, 247, 248,
		3, 62, 31, 0, 248, 249, 6, 8, -1, 0, 249, 264, 1, 0, 0, 0, 250, 251, 5,
		9, 0, 0, 251, 252, 5, 39, 0, 0, 252, 253, 5, 45, 0, 0, 253, 254, 3, 62,
		31, 0, 254, 255, 6, 8, -1, 0, 255, 264, 1, 0, 0, 0, 256, 257, 5, 9, 0,
		0, 257, 258, 5, 39, 0, 0, 258, 259, 5, 65, 0, 0, 259, 260, 3, 18, 9, 0,
		260, 261, 5, 61, 0, 0, 261, 262, 6, 8, -1, 0, 262, 264, 1, 0, 0, 0, 263,
		242, 1, 0, 0, 0, 263, 250, 1, 0, 0, 0, 263, 256, 1, 0, 0, 0, 264, 17, 1,
		0, 0, 0, 265, 266, 5, 1, 0, 0, 266, 276, 6, 9, -1, 0, 267, 268, 5, 2, 0,
		0, 268, 276, 6, 9, -1, 0, 269, 270, 5, 4, 0, 0, 270, 276, 6, 9, -1, 0,
		271, 272, 5, 3, 0, 0, 272, 276, 6, 9, -1, 0, 273, 274, 5, 5, 0, 0, 274,
		276, 6, 9, -1, 0, 275, 265, 1, 0, 0, 0, 275, 267, 1, 0, 0, 0, 275, 269,
		1, 0, 0, 0, 275, 271, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 19, 1, 0,
		0, 0, 277, 278, 5, 39, 0, 0, 278, 279, 5, 45, 0, 0, 279, 280, 3, 62, 31,
		0, 280, 281, 6, 10, -1, 0, 281, 295, 1, 0, 0, 0, 282, 283, 5, 39, 0, 0,
		283, 284, 5, 53, 0, 0, 284, 285, 5, 45, 0, 0, 285, 286, 3, 62, 31, 0, 286,
		287, 6, 10, -1, 0, 287, 295, 1, 0, 0, 0, 288, 289, 5, 39, 0, 0, 289, 290,
		5, 54, 0, 0, 290, 291, 5, 45, 0, 0, 291, 292, 3, 62, 31, 0, 292, 293, 6,
		10, -1, 0, 293, 295, 1, 0, 0, 0, 294, 277, 1, 0, 0, 0, 294, 282, 1, 0,
		0, 0, 294, 288, 1, 0, 0, 0, 295, 21, 1, 0, 0, 0, 296, 297, 5, 10, 0, 0,
		297, 298, 5, 39, 0, 0, 298, 299, 5, 45, 0, 0, 299, 300, 3, 62, 31, 0, 300,
		301, 6, 11, -1, 0, 301, 311, 1, 0, 0, 0, 302, 303, 5, 10, 0, 0, 303, 304,
		5, 39, 0, 0, 304, 305, 5, 65, 0, 0, 305, 306, 3, 18, 9, 0, 306, 307, 5,
		45, 0, 0, 307, 308, 3, 62, 31, 0, 308, 309, 6, 11, -1, 0, 309, 311, 1,
		0, 0, 0, 310, 296, 1, 0, 0, 0, 310, 302, 1, 0, 0, 0, 311, 23, 1, 0, 0,
		0, 312, 313, 5, 14, 0, 0, 313, 314, 3, 62, 31, 0, 314, 316, 5, 57, 0, 0,
		315, 317, 3, 26, 13, 0, 316, 315, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318,
		316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 321, 1, 0, 0, 0, 320, 322,
		3, 28, 14, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 323, 1,
		0, 0, 0, 323, 324, 5, 58, 0, 0, 324, 325, 6, 12, -1, 0, 325, 25, 1, 0,
		0, 0, 326, 327, 5, 15, 0, 0, 327, 328, 3, 62, 31, 0, 328, 329, 5, 65, 0,
		0, 329, 331, 3, 2, 1, 0, 330, 332, 5, 22, 0, 0, 331, 330, 1, 0, 0, 0, 331,
		332, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 6, 13, -1, 0, 334, 27,
		1, 0, 0, 0, 335, 336, 5, 16, 0, 0, 336, 337, 5, 65, 0, 0, 337, 339, 3,
		2, 1, 0, 338, 340, 5, 22, 0, 0, 339, 338, 1, 0, 0, 0, 339, 340, 1, 0, 0,
		0, 340, 341, 1, 0, 0, 0, 341, 342, 6, 14, -1, 0, 342, 29, 1, 0, 0, 0, 343,
		344, 5, 17, 0, 0, 344, 345, 3, 62, 31, 0, 345, 346, 5, 57, 0, 0, 346, 347,
		3, 2, 1, 0, 347, 348, 5, 58, 0, 0, 348, 349, 6, 15, -1, 0, 349, 31, 1,
		0, 0, 0, 350, 351, 5, 18, 0, 0, 351, 352, 5, 39, 0, 0, 352, 353, 5, 19,
		0, 0, 353, 354, 5, 38, 0, 0, 354, 355, 5, 57, 0, 0, 355, 356, 3, 2, 1,
		0, 356, 357, 5, 58, 0, 0, 357, 358, 6, 16, -1, 0, 358, 380, 1, 0, 0, 0,
		359, 360, 5, 18, 0, 0, 360, 361, 5, 39, 0, 0, 361, 362, 5, 19, 0, 0, 362,
		363, 3, 62, 31, 0, 363, 364, 5, 20, 0, 0, 364, 365, 3, 62, 31, 0, 365,
		366, 5, 57, 0, 0, 366, 367, 3, 2, 1, 0, 367, 368, 5, 58, 0, 0, 368, 369,
		6, 16, -1, 0, 369, 380, 1, 0, 0, 0, 370, 371, 5, 18, 0, 0, 371, 372, 5,
		39, 0, 0, 372, 373, 5, 19, 0, 0, 373, 374, 5, 39, 0, 0, 374, 375, 5, 57,
		0, 0, 375, 376, 3, 2, 1, 0, 376, 377, 5, 58, 0, 0, 377, 378, 6, 16, -1,
		0, 378, 380, 1, 0, 0, 0, 379, 350, 1, 0, 0, 0, 379, 359, 1, 0, 0, 0, 379,
		370, 1, 0, 0, 0, 380, 33, 1, 0, 0, 0, 381, 382, 5, 21, 0, 0, 382, 383,
		3, 62, 31, 0, 383, 384, 5, 13, 0, 0, 384, 385, 5, 57, 0, 0, 385, 386, 3,
		2, 1, 0, 386, 387, 5, 58, 0, 0, 387, 388, 6, 17, -1, 0, 388, 35, 1, 0,
		0, 0, 389, 390, 5, 22, 0, 0, 390, 400, 6, 18, -1, 0, 391, 392, 5, 23, 0,
		0, 392, 400, 6, 18, -1, 0, 393, 394, 5, 24, 0, 0, 394, 400, 6, 18, -1,
		0, 395, 396, 5, 24, 0, 0, 396, 397, 3, 62, 31, 0, 397, 398, 6, 18, -1,
		0, 398, 400, 1, 0, 0, 0, 399, 389, 1, 0, 0, 0, 399, 391, 1, 0, 0, 0, 399,
		393, 1, 0, 0, 0, 399, 395, 1, 0, 0, 0, 400, 37, 1, 0, 0, 0, 401, 402, 5,
		9, 0, 0, 402, 403, 5, 39, 0, 0, 403, 404, 5, 65, 0, 0, 404, 405, 5, 59,
		0, 0, 405, 406, 3, 18, 9, 0, 406, 407, 5, 60, 0, 0, 407, 408, 3, 40, 20,
		0, 408, 409, 6, 19, -1, 0, 409, 39, 1, 0, 0, 0, 410, 411, 5, 45, 0, 0,
		411, 413, 5, 59, 0, 0, 412, 414, 3, 42, 21, 0, 413, 412, 1, 0, 0, 0, 414,
		415, 1, 0, 0, 0, 415, 413, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 417,
		1, 0, 0, 0, 417, 418, 5, 60, 0, 0, 418, 419, 6, 20, -1, 0, 419, 425, 1,
		0, 0, 0, 420, 421, 5, 45, 0, 0, 421, 422, 5, 59, 0, 0, 422, 423, 5, 60,
		0, 0, 423, 425, 6, 20, -1, 0, 424, 410, 1, 0, 0, 0, 424, 420, 1, 0, 0,
		0, 425, 41, 1, 0, 0, 0, 426, 427, 3, 62, 31, 0, 427, 428, 5, 63, 0, 0,
		428, 429, 6, 21, -1, 0, 429, 434, 1, 0, 0, 0, 430, 431, 3, 62, 31, 0, 431,
		432, 6, 21, -1, 0, 432, 434, 1, 0, 0, 0, 433, 426, 1, 0, 0, 0, 433, 430,
		1, 0, 0, 0, 434, 43, 1, 0, 0, 0, 435, 436, 5, 39, 0, 0, 436, 437, 5, 64,
		0, 0, 437, 438, 5, 29, 0, 0, 438, 439, 5, 55, 0, 0, 439, 440, 3, 62, 31,
		0, 440, 441, 5, 56, 0, 0, 441, 442, 6, 22, -1, 0, 442, 458, 1, 0, 0, 0,
		443, 444, 5, 39, 0, 0, 444, 445, 5, 64, 0, 0, 445, 446, 5, 30, 0, 0, 446,
		447, 5, 55, 0, 0, 447, 448, 5, 56, 0, 0, 448, 458, 6, 22, -1, 0, 449, 450,
		5, 39, 0, 0, 450, 451, 5, 64, 0, 0, 451, 452, 5, 31, 0, 0, 452, 453, 5,
		55, 0, 0, 453, 454, 3, 62, 31, 0, 454, 455, 5, 56, 0, 0, 455, 456, 6, 22,
		-1, 0, 456, 458, 1, 0, 0, 0, 457, 435, 1, 0, 0, 0, 457, 443, 1, 0, 0, 0,
		457, 449, 1, 0, 0, 0, 458, 45, 1, 0, 0, 0, 459, 460, 5, 39, 0, 0, 460,
		461, 5, 64, 0, 0, 461, 462, 5, 32, 0, 0, 462, 474, 6, 23, -1, 0, 463, 464,
		5, 39, 0, 0, 464, 465, 5, 64, 0, 0, 465, 466, 5, 33, 0, 0, 466, 474, 6,
		23, -1, 0, 467, 468, 5, 39, 0, 0, 468, 469, 5, 59, 0, 0, 469, 470, 3, 62,
		31, 0, 470, 471, 5, 60, 0, 0, 471, 472, 6, 23, -1, 0, 472, 474, 1, 0, 0,
		0, 473, 459, 1, 0, 0, 0, 473, 463, 1, 0, 0, 0, 473, 467, 1, 0, 0, 0, 474,
		47, 1, 0, 0, 0, 475, 476, 5, 39, 0, 0, 476, 477, 5, 59, 0, 0, 477, 478,
		3, 62, 31, 0, 478, 479, 5, 60, 0, 0, 479, 480, 5, 45, 0, 0, 480, 481, 5,
		39, 0, 0, 481, 482, 5, 59, 0, 0, 482, 483, 3, 62, 31, 0, 483, 484, 5, 60,
		0, 0, 484, 485, 6, 24, -1, 0, 485, 495, 1, 0, 0, 0, 486, 487, 5, 39, 0,
		0, 487, 488, 5, 59, 0, 0, 488, 489, 3, 62, 31, 0, 489, 490, 5, 60, 0, 0,
		490, 491, 5, 45, 0, 0, 491, 492, 3, 62, 31, 0, 492, 493, 6, 24, -1, 0,
		493, 495, 1, 0, 0, 0, 494, 475, 1, 0, 0, 0, 494, 486, 1, 0, 0, 0, 495,
		49, 1, 0, 0, 0, 496, 497, 5, 28, 0, 0, 497, 498, 5, 39, 0, 0, 498, 500,
		5, 55, 0, 0, 499, 501, 3, 52, 26, 0, 500, 499, 1, 0, 0, 0, 501, 502, 1,
		0, 0, 0, 502, 500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 504, 1, 0, 0,
		0, 504, 505, 5, 56, 0, 0, 505, 506, 5, 62, 0, 0, 506, 507, 3, 18, 9, 0,
		507, 508, 5, 57, 0, 0, 508, 509, 3, 2, 1, 0, 509, 510, 5, 58, 0, 0, 510,
		511, 6, 25, -1, 0, 511, 547, 1, 0, 0, 0, 512, 513, 5, 28, 0, 0, 513, 514,
		5, 39, 0, 0, 514, 516, 5, 55, 0, 0, 515, 517, 3, 52, 26, 0, 516, 515, 1,
		0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 516, 1, 0, 0, 0, 518, 519, 1, 0, 0,
		0, 519, 520, 1, 0, 0, 0, 520, 521, 5, 56, 0, 0, 521, 522, 5, 57, 0, 0,
		522, 523, 3, 2, 1, 0, 523, 524, 5, 58, 0, 0, 524, 525, 6, 25, -1, 0, 525,
		547, 1, 0, 0, 0, 526, 527, 5, 28, 0, 0, 527, 528, 5, 39, 0, 0, 528, 529,
		5, 55, 0, 0, 529, 530, 5, 56, 0, 0, 530, 531, 5, 62, 0, 0, 531, 532, 3,
		18, 9, 0, 532, 533, 5, 57, 0, 0, 533, 534, 3, 2, 1, 0, 534, 535, 5, 58,
		0, 0, 535, 536, 6, 25, -1, 0, 536, 547, 1, 0, 0, 0, 537, 538, 5, 28, 0,
		0, 538, 539, 5, 39, 0, 0, 539, 540, 5, 55, 0, 0, 540, 541, 5, 56, 0, 0,
		541, 542, 5, 57, 0, 0, 542, 543, 3, 2, 1, 0, 543, 544, 5, 58, 0, 0, 544,
		545, 6, 25, -1, 0, 545, 547, 1, 0, 0, 0, 546, 496, 1, 0, 0, 0, 546, 512,
		1, 0, 0, 0, 546, 526, 1, 0, 0, 0, 546, 537, 1, 0, 0, 0, 547, 51, 1, 0,
		0, 0, 548, 550, 7, 0, 0, 0, 549, 548, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0,
		550, 551, 1, 0, 0, 0, 551, 552, 5, 39, 0, 0, 552, 554, 5, 65, 0, 0, 553,
		555, 5, 34, 0, 0, 554, 553, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555, 556,
		1, 0, 0, 0, 556, 557, 3, 18, 9, 0, 557, 558, 5, 63, 0, 0, 558, 559, 6,
		26, -1, 0, 559, 572, 1, 0, 0, 0, 560, 562, 7, 0, 0, 0, 561, 560, 1, 0,
		0, 0, 561, 562, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 564, 5, 39, 0, 0,
		564, 566, 5, 65, 0, 0, 565, 567, 5, 34, 0, 0, 566, 565, 1, 0, 0, 0, 566,
		567, 1, 0, 0, 0, 567, 568, 1, 0, 0, 0, 568, 569, 3, 18, 9, 0, 569, 570,
		6, 26, -1, 0, 570, 572, 1, 0, 0, 0, 571, 549, 1, 0, 0, 0, 571, 561, 1,
		0, 0, 0, 572, 53, 1, 0, 0, 0, 573, 574, 5, 39, 0, 0, 574, 576, 5, 55, 0,
		0, 575, 577, 3, 58, 29, 0, 576, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0,
		578, 576, 1, 0, 0, 0, 578, 579, 1, 0, 0, 0, 579, 580, 1, 0, 0, 0, 580,
		581, 5, 56, 0, 0, 581, 582, 6, 27, -1, 0, 582, 588, 1, 0, 0, 0, 583, 584,
		5, 39, 0, 0, 584, 585, 5, 55, 0, 0, 585, 586, 5, 56, 0, 0, 586, 588, 6,
		27, -1, 0, 587, 573, 1, 0, 0, 0, 587, 583, 1, 0, 0, 0, 588, 55, 1, 0, 0,
		0, 589, 590, 5, 39, 0, 0, 590, 592, 5, 55, 0, 0, 591, 593, 3, 58, 29, 0,
		592, 591, 1, 0, 0, 0, 593, 594, 1, 0, 0, 0, 594, 592, 1, 0, 0, 0, 594,
		595, 1, 0, 0, 0, 595, 596, 1, 0, 0, 0, 596, 597, 5, 56, 0, 0, 597, 598,
		6, 28, -1, 0, 598, 604, 1, 0, 0, 0, 599, 600, 5, 39, 0, 0, 600, 601, 5,
		55, 0, 0, 601, 602, 5, 56, 0, 0, 602, 604, 6, 28, -1, 0, 603, 589, 1, 0,
		0, 0, 603, 599, 1, 0, 0, 0, 604, 57, 1, 0, 0, 0, 605, 606, 5, 39, 0, 0,
		606, 608, 5, 65, 0, 0, 607, 605, 1, 0, 0, 0, 607, 608, 1, 0, 0, 0, 608,
		610, 1, 0, 0, 0, 609, 611, 5, 66, 0, 0, 610, 609, 1, 0, 0, 0, 610, 611,
		1, 0, 0, 0, 611, 612, 1, 0, 0, 0, 612, 613, 3, 62, 31, 0, 613, 614, 5,
		63, 0, 0, 614, 615, 6, 29, -1, 0, 615, 627, 1, 0, 0, 0, 616, 617, 5, 39,
		0, 0, 617, 619, 5, 65, 0, 0, 618, 616, 1, 0, 0, 0, 618, 619, 1, 0, 0, 0,
		619, 621, 1, 0, 0, 0, 620, 622, 5, 66, 0, 0, 621, 620, 1, 0, 0, 0, 621,
		622, 1, 0, 0, 0, 622, 623, 1, 0, 0, 0, 623, 624, 3, 62, 31, 0, 624, 625,
		6, 29, -1, 0, 625, 627, 1, 0, 0, 0, 626, 607, 1, 0, 0, 0, 626, 618, 1,
		0, 0, 0, 627, 59, 1, 0, 0, 0, 628, 629, 5, 4, 0, 0, 629, 630, 5, 55, 0,
		0, 630, 631, 3, 62, 31, 0, 631, 632, 5, 56, 0, 0, 632, 633, 6, 30, -1,
		0, 633, 647, 1, 0, 0, 0, 634, 635, 5, 1, 0, 0, 635, 636, 5, 55, 0, 0, 636,
		637, 3, 62, 31, 0, 637, 638, 5, 56, 0, 0, 638, 639, 6, 30, -1, 0, 639,
		647, 1, 0, 0, 0, 640, 641, 5, 2, 0, 0, 641, 642, 5, 55, 0, 0, 642, 643,
		3, 62, 31, 0, 643, 644, 5, 56, 0, 0, 644, 645, 6, 30, -1, 0, 645, 647,
		1, 0, 0, 0, 646, 628, 1, 0, 0, 0, 646, 634, 1, 0, 0, 0, 646, 640, 1, 0,
		0, 0, 647, 61, 1, 0, 0, 0, 648, 649, 6, 31, -1, 0, 649, 650, 5, 54, 0,
		0, 650, 651, 3, 62, 31, 20, 651, 652, 6, 31, -1, 0, 652, 686, 1, 0, 0,
		0, 653, 654, 5, 42, 0, 0, 654, 655, 3, 62, 31, 19, 655, 656, 6, 31, -1,
		0, 656, 686, 1, 0, 0, 0, 657, 658, 5, 55, 0, 0, 658, 659, 3, 62, 31, 0,
		659, 660, 5, 56, 0, 0, 660, 661, 6, 31, -1, 0, 661, 686, 1, 0, 0, 0, 662,
		663, 5, 36, 0, 0, 663, 686, 6, 31, -1, 0, 664, 665, 5, 37, 0, 0, 665, 686,
		6, 31, -1, 0, 666, 667, 5, 38, 0, 0, 667, 686, 6, 31, -1, 0, 668, 669,
		5, 6, 0, 0, 669, 686, 6, 31, -1, 0, 670, 671, 3, 54, 27, 0, 671, 672, 6,
		31, -1, 0, 672, 686, 1, 0, 0, 0, 673, 674, 3, 60, 30, 0, 674, 675, 6, 31,
		-1, 0, 675, 686, 1, 0, 0, 0, 676, 677, 5, 7, 0, 0, 677, 686, 6, 31, -1,
		0, 678, 679, 5, 39, 0, 0, 679, 686, 6, 31, -1, 0, 680, 681, 3, 46, 23,
		0, 681, 682, 6, 31, -1, 0, 682, 686, 1, 0, 0, 0, 683, 684, 5, 8, 0, 0,
		684, 686, 6, 31, -1, 0, 685, 648, 1, 0, 0, 0, 685, 653, 1, 0, 0, 0, 685,
		657, 1, 0, 0, 0, 685, 662, 1, 0, 0, 0, 685, 664, 1, 0, 0, 0, 685, 666,
		1, 0, 0, 0, 685, 668, 1, 0, 0, 0, 685, 670, 1, 0, 0, 0, 685, 673, 1, 0,
		0, 0, 685, 676, 1, 0, 0, 0, 685, 678, 1, 0, 0, 0, 685, 680, 1, 0, 0, 0,
		685, 683, 1, 0, 0, 0, 686, 724, 1, 0, 0, 0, 687, 688, 10, 18, 0, 0, 688,
		689, 7, 1, 0, 0, 689, 690, 3, 62, 31, 19, 690, 691, 6, 31, -1, 0, 691,
		723, 1, 0, 0, 0, 692, 693, 10, 17, 0, 0, 693, 694, 7, 2, 0, 0, 694, 695,
		3, 62, 31, 18, 695, 696, 6, 31, -1, 0, 696, 723, 1, 0, 0, 0, 697, 698,
		10, 16, 0, 0, 698, 699, 7, 3, 0, 0, 699, 700, 3, 62, 31, 17, 700, 701,
		6, 31, -1, 0, 701, 723, 1, 0, 0, 0, 702, 703, 10, 15, 0, 0, 703, 704, 7,
		4, 0, 0, 704, 705, 3, 62, 31, 16, 705, 706, 6, 31, -1, 0, 706, 723, 1,
		0, 0, 0, 707, 708, 10, 14, 0, 0, 708, 709, 7, 5, 0, 0, 709, 710, 3, 62,
		31, 15, 710, 711, 6, 31, -1, 0, 711, 723, 1, 0, 0, 0, 712, 713, 10, 13,
		0, 0, 713, 714, 5, 44, 0, 0, 714, 715, 3, 62, 31, 14, 715, 716, 6, 31,
		-1, 0, 716, 723, 1, 0, 0, 0, 717, 718, 10, 12, 0, 0, 718, 719, 5, 43, 0,
		0, 719, 720, 3, 62, 31, 13, 720, 721, 6, 31, -1, 0, 721, 723, 1, 0, 0,
		0, 722, 687, 1, 0, 0, 0, 722, 692, 1, 0, 0, 0, 722, 697, 1, 0, 0, 0, 722,
		702, 1, 0, 0, 0, 722, 707, 1, 0, 0, 0, 722, 712, 1, 0, 0, 0, 722, 717,
		1, 0, 0, 0, 723, 726, 1, 0, 0, 0, 724, 722, 1, 0, 0, 0, 724, 725, 1, 0,
		0, 0, 725, 63, 1, 0, 0, 0, 726, 724, 1, 0, 0, 0, 59, 71, 77, 83, 89, 95,
		101, 107, 113, 119, 125, 131, 137, 143, 149, 155, 161, 165, 172, 184, 201,
		214, 226, 263, 275, 294, 310, 318, 321, 331, 339, 379, 399, 415, 424, 433,
		457, 473, 494, 502, 518, 546, 549, 554, 561, 566, 571, 578, 587, 594, 603,
		607, 610, 618, 621, 626, 646, 685, 722, 724,
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
	SwiftGrammarParserRANGEPTS     = 20
	SwiftGrammarParserGUARD        = 21
	SwiftGrammarParserBREAK        = 22
	SwiftGrammarParserCONTINUE     = 23
	SwiftGrammarParserRETURN       = 24
	SwiftGrammarParserSTRUCT       = 25
	SwiftGrammarParserSELF         = 26
	SwiftGrammarParserMUTATING     = 27
	SwiftGrammarParserFUNC         = 28
	SwiftGrammarParserAPPEND       = 29
	SwiftGrammarParserREMOVELAST   = 30
	SwiftGrammarParserREMOVE       = 31
	SwiftGrammarParserEMPTY        = 32
	SwiftGrammarParserCOUNT        = 33
	SwiftGrammarParserINOUT        = 34
	SwiftGrammarParserCAME         = 35
	SwiftGrammarParserNUMBER       = 36
	SwiftGrammarParserCHAR         = 37
	SwiftGrammarParserSTRING       = 38
	SwiftGrammarParserID           = 39
	SwiftGrammarParserDIF          = 40
	SwiftGrammarParserIG_IG        = 41
	SwiftGrammarParserNOT          = 42
	SwiftGrammarParserOR           = 43
	SwiftGrammarParserAND          = 44
	SwiftGrammarParserIG           = 45
	SwiftGrammarParserMAY_IG       = 46
	SwiftGrammarParserMEN_IG       = 47
	SwiftGrammarParserMAYOR        = 48
	SwiftGrammarParserMENOR        = 49
	SwiftGrammarParserMOD          = 50
	SwiftGrammarParserMUL          = 51
	SwiftGrammarParserDIV          = 52
	SwiftGrammarParserADD          = 53
	SwiftGrammarParserSUB          = 54
	SwiftGrammarParserPARIZQ       = 55
	SwiftGrammarParserPARDER       = 56
	SwiftGrammarParserLLAVEIZQ     = 57
	SwiftGrammarParserLLAVEDER     = 58
	SwiftGrammarParserCORCHIZQ     = 59
	SwiftGrammarParserCORCHDER     = 60
	SwiftGrammarParserQ_MARK       = 61
	SwiftGrammarParserARROW        = 62
	SwiftGrammarParserCOMA         = 63
	SwiftGrammarParserPUNTO        = 64
	SwiftGrammarParserCOLON        = 65
	SwiftGrammarParserAMP          = 66
	SwiftGrammarParserPUNTOCOMA    = 67
	SwiftGrammarParserWHITESPACE   = 68
	SwiftGrammarParserCOMMENT      = 69
	SwiftGrammarParserLINE_COMMENT = 70
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                   = 0
	SwiftGrammarParserRULE_block               = 1
	SwiftGrammarParserRULE_instruction         = 2
	SwiftGrammarParserRULE_printstmt           = 3
	SwiftGrammarParserRULE_printexprlist       = 4
	SwiftGrammarParserRULE_ifstmt              = 5
	SwiftGrammarParserRULE_elseifstmt          = 6
	SwiftGrammarParserRULE_elsestmt            = 7
	SwiftGrammarParserRULE_varstmt             = 8
	SwiftGrammarParserRULE_tipo                = 9
	SwiftGrammarParserRULE_varasgmt            = 10
	SwiftGrammarParserRULE_conststmt           = 11
	SwiftGrammarParserRULE_switchstmt          = 12
	SwiftGrammarParserRULE_casestmt            = 13
	SwiftGrammarParserRULE_defaultstmt         = 14
	SwiftGrammarParserRULE_whilestmt           = 15
	SwiftGrammarParserRULE_forstmt             = 16
	SwiftGrammarParserRULE_guardstmt           = 17
	SwiftGrammarParserRULE_transferstmt        = 18
	SwiftGrammarParserRULE_vectorstmt          = 19
	SwiftGrammarParserRULE_definestmt          = 20
	SwiftGrammarParserRULE_listexpr            = 21
	SwiftGrammarParserRULE_methodvec           = 22
	SwiftGrammarParserRULE_methodvecrtrn       = 23
	SwiftGrammarParserRULE_vecaccess           = 24
	SwiftGrammarParserRULE_funcstmt            = 25
	SwiftGrammarParserRULE_listparam           = 26
	SwiftGrammarParserRULE_callfunc            = 27
	SwiftGrammarParserRULE_callfuncinstruction = 28
	SwiftGrammarParserRULE_listparamcall       = 29
	SwiftGrammarParserRULE_funcembed           = 30
	SwiftGrammarParserRULE_expr                = 31
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
		p.SetState(64)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(65)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(68)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
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

	// Get_callfuncinstruction returns the _callfuncinstruction rule contexts.
	Get_callfuncinstruction() ICallfuncinstructionContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_varstmt returns the _varstmt rule contexts.
	Get_varstmt() IVarstmtContext

	// Get_varasgmt returns the _varasgmt rule contexts.
	Get_varasgmt() IVarasgmtContext

	// Get_conststmt returns the _conststmt rule contexts.
	Get_conststmt() IConststmtContext

	// Get_switchstmt returns the _switchstmt rule contexts.
	Get_switchstmt() ISwitchstmtContext

	// Get_whilestmt returns the _whilestmt rule contexts.
	Get_whilestmt() IWhilestmtContext

	// Get_forstmt returns the _forstmt rule contexts.
	Get_forstmt() IForstmtContext

	// Get_guardstmt returns the _guardstmt rule contexts.
	Get_guardstmt() IGuardstmtContext

	// Get_transferstmt returns the _transferstmt rule contexts.
	Get_transferstmt() ITransferstmtContext

	// Get_vectorstmt returns the _vectorstmt rule contexts.
	Get_vectorstmt() IVectorstmtContext

	// Get_methodvec returns the _methodvec rule contexts.
	Get_methodvec() IMethodvecContext

	// Get_vecaccess returns the _vecaccess rule contexts.
	Get_vecaccess() IVecaccessContext

	// Get_funcstmt returns the _funcstmt rule contexts.
	Get_funcstmt() IFuncstmtContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_callfuncinstruction sets the _callfuncinstruction rule contexts.
	Set_callfuncinstruction(ICallfuncinstructionContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_varstmt sets the _varstmt rule contexts.
	Set_varstmt(IVarstmtContext)

	// Set_varasgmt sets the _varasgmt rule contexts.
	Set_varasgmt(IVarasgmtContext)

	// Set_conststmt sets the _conststmt rule contexts.
	Set_conststmt(IConststmtContext)

	// Set_switchstmt sets the _switchstmt rule contexts.
	Set_switchstmt(ISwitchstmtContext)

	// Set_whilestmt sets the _whilestmt rule contexts.
	Set_whilestmt(IWhilestmtContext)

	// Set_forstmt sets the _forstmt rule contexts.
	Set_forstmt(IForstmtContext)

	// Set_guardstmt sets the _guardstmt rule contexts.
	Set_guardstmt(IGuardstmtContext)

	// Set_transferstmt sets the _transferstmt rule contexts.
	Set_transferstmt(ITransferstmtContext)

	// Set_vectorstmt sets the _vectorstmt rule contexts.
	Set_vectorstmt(IVectorstmtContext)

	// Set_methodvec sets the _methodvec rule contexts.
	Set_methodvec(IMethodvecContext)

	// Set_vecaccess sets the _vecaccess rule contexts.
	Set_vecaccess(IVecaccessContext)

	// Set_funcstmt sets the _funcstmt rule contexts.
	Set_funcstmt(IFuncstmtContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	PUNTOCOMA() antlr.TerminalNode
	Callfuncinstruction() ICallfuncinstructionContext
	Ifstmt() IIfstmtContext
	Varstmt() IVarstmtContext
	Varasgmt() IVarasgmtContext
	Conststmt() IConststmtContext
	Switchstmt() ISwitchstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Transferstmt() ITransferstmtContext
	Vectorstmt() IVectorstmtContext
	Methodvec() IMethodvecContext
	Vecaccess() IVecaccessContext
	Funcstmt() IFuncstmtContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser               antlr.Parser
	inst                 interfaces.Instruction
	_printstmt           IPrintstmtContext
	_callfuncinstruction ICallfuncinstructionContext
	_ifstmt              IIfstmtContext
	_varstmt             IVarstmtContext
	_varasgmt            IVarasgmtContext
	_conststmt           IConststmtContext
	_switchstmt          ISwitchstmtContext
	_whilestmt           IWhilestmtContext
	_forstmt             IForstmtContext
	_guardstmt           IGuardstmtContext
	_transferstmt        ITransferstmtContext
	_vectorstmt          IVectorstmtContext
	_methodvec           IMethodvecContext
	_vecaccess           IVecaccessContext
	_funcstmt            IFuncstmtContext
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

func (s *InstructionContext) Get_callfuncinstruction() ICallfuncinstructionContext {
	return s._callfuncinstruction
}

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_varstmt() IVarstmtContext { return s._varstmt }

func (s *InstructionContext) Get_varasgmt() IVarasgmtContext { return s._varasgmt }

func (s *InstructionContext) Get_conststmt() IConststmtContext { return s._conststmt }

func (s *InstructionContext) Get_switchstmt() ISwitchstmtContext { return s._switchstmt }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_guardstmt() IGuardstmtContext { return s._guardstmt }

func (s *InstructionContext) Get_transferstmt() ITransferstmtContext { return s._transferstmt }

func (s *InstructionContext) Get_vectorstmt() IVectorstmtContext { return s._vectorstmt }

func (s *InstructionContext) Get_methodvec() IMethodvecContext { return s._methodvec }

func (s *InstructionContext) Get_vecaccess() IVecaccessContext { return s._vecaccess }

func (s *InstructionContext) Get_funcstmt() IFuncstmtContext { return s._funcstmt }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_callfuncinstruction(v ICallfuncinstructionContext) {
	s._callfuncinstruction = v
}

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_varstmt(v IVarstmtContext) { s._varstmt = v }

func (s *InstructionContext) Set_varasgmt(v IVarasgmtContext) { s._varasgmt = v }

func (s *InstructionContext) Set_conststmt(v IConststmtContext) { s._conststmt = v }

func (s *InstructionContext) Set_switchstmt(v ISwitchstmtContext) { s._switchstmt = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_guardstmt(v IGuardstmtContext) { s._guardstmt = v }

func (s *InstructionContext) Set_transferstmt(v ITransferstmtContext) { s._transferstmt = v }

func (s *InstructionContext) Set_vectorstmt(v IVectorstmtContext) { s._vectorstmt = v }

func (s *InstructionContext) Set_methodvec(v IMethodvecContext) { s._methodvec = v }

func (s *InstructionContext) Set_vecaccess(v IVecaccessContext) { s._vecaccess = v }

func (s *InstructionContext) Set_funcstmt(v IFuncstmtContext) { s._funcstmt = v }

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

func (s *InstructionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTOCOMA, 0)
}

func (s *InstructionContext) Callfuncinstruction() ICallfuncinstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallfuncinstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallfuncinstructionContext)
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

func (s *InstructionContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *InstructionContext) Transferstmt() ITransferstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransferstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransferstmtContext)
}

func (s *InstructionContext) Vectorstmt() IVectorstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorstmtContext)
}

func (s *InstructionContext) Methodvec() IMethodvecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodvecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodvecContext)
}

func (s *InstructionContext) Vecaccess() IVecaccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVecaccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVecaccessContext)
}

func (s *InstructionContext) Funcstmt() IFuncstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncstmtContext)
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
	var _la int

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(76)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)

			var _x = p.Callfuncinstruction()

			localctx.(*InstructionContext)._callfuncinstruction = _x
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(82)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_callfuncinstruction().GetCallfuncinstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(88)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)

			var _x = p.Varstmt()

			localctx.(*InstructionContext)._varstmt = _x
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(94)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_varstmt().GetVar_()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)

			var _x = p.Varasgmt()

			localctx.(*InstructionContext)._varasgmt = _x
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(100)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_varasgmt().GetAsgmt()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)

			var _x = p.Conststmt()

			localctx.(*InstructionContext)._conststmt = _x
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(106)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_conststmt().GetConst_()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(112)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSwitchinstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(117)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(118)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhileinstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(123)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(124)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetForinstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(129)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(130)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGuardinstr()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(135)

			var _x = p.Transferstmt()

			localctx.(*InstructionContext)._transferstmt = _x
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(136)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transferstmt().GetTrns()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(141)

			var _x = p.Vectorstmt()

			localctx.(*InstructionContext)._vectorstmt = _x
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(142)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vectorstmt().GetVectorinstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(147)

			var _x = p.Methodvec()

			localctx.(*InstructionContext)._methodvec = _x
		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(148)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_methodvec().GetMethodinstr()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(153)

			var _x = p.Vecaccess()

			localctx.(*InstructionContext)._vecaccess = _x
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(154)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vecaccess().GetVecacc()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(159)

			var _x = p.Funcstmt()

			localctx.(*InstructionContext)._funcstmt = _x
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(160)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_funcstmt().GetFuncinstr()

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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_printexprlist returns the _printexprlist rule contexts.
	Get_printexprlist() IPrintexprlistContext

	// Set_printexprlist sets the _printexprlist rule contexts.
	Set_printexprlist(IPrintexprlistContext)

	// GetLista returns the lista rule context list.
	GetLista() []IPrintexprlistContext

	// SetLista sets the lista rule context list.
	SetLista([]IPrintexprlistContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	AllPrintexprlist() []IPrintexprlistContext
	Printexprlist(i int) IPrintexprlistContext

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	prnt           interfaces.Instruction
	_PRINT         antlr.Token
	_printexprlist IPrintexprlistContext
	lista          []IPrintexprlistContext
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

func (s *PrintstmtContext) Get_printexprlist() IPrintexprlistContext { return s._printexprlist }

func (s *PrintstmtContext) Set_printexprlist(v IPrintexprlistContext) { s._printexprlist = v }

func (s *PrintstmtContext) GetLista() []IPrintexprlistContext { return s.lista }

func (s *PrintstmtContext) SetLista(v []IPrintexprlistContext) { s.lista = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) AllPrintexprlist() []IPrintexprlistContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintexprlistContext); ok {
			len++
		}
	}

	tst := make([]IPrintexprlistContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintexprlistContext); ok {
			tst[i] = t.(IPrintexprlistContext)
			i++
		}
	}

	return tst
}

func (s *PrintstmtContext) Printexprlist(i int) IPrintexprlistContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintexprlistContext); ok {
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

	return t.(IPrintexprlistContext)
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

	var listExpr []interface{}

	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54048624367108566) != 0) {
		{
			p.SetState(169)

			var _x = p.Printexprlist()

			localctx.(*PrintstmtContext)._printexprlist = _x
		}
		localctx.(*PrintstmtContext).lista = append(localctx.(*PrintstmtContext).lista, localctx.(*PrintstmtContext)._printexprlist)

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	for _, e := range localctx.(*PrintstmtContext).GetLista() {
		listExpr = append(listExpr, e.GetPrntexpr())
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
	}()), listExpr)

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

// IPrintexprlistContext is an interface to support dynamic dispatch.
type IPrintexprlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetPrntexpr returns the prntexpr attribute.
	GetPrntexpr() interfaces.Expression

	// SetPrntexpr sets the prntexpr attribute.
	SetPrntexpr(interfaces.Expression)

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode

	// IsPrintexprlistContext differentiates from other interfaces.
	IsPrintexprlistContext()
}

type PrintexprlistContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	prntexpr interfaces.Expression
	_expr    IExprContext
}

func NewEmptyPrintexprlistContext() *PrintexprlistContext {
	var p = new(PrintexprlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printexprlist
	return p
}

func InitEmptyPrintexprlistContext(p *PrintexprlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printexprlist
}

func (*PrintexprlistContext) IsPrintexprlistContext() {}

func NewPrintexprlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintexprlistContext {
	var p = new(PrintexprlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printexprlist

	return p
}

func (s *PrintexprlistContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintexprlistContext) Get_expr() IExprContext { return s._expr }

func (s *PrintexprlistContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintexprlistContext) GetPrntexpr() interfaces.Expression { return s.prntexpr }

func (s *PrintexprlistContext) SetPrntexpr(v interfaces.Expression) { s.prntexpr = v }

func (s *PrintexprlistContext) Expr() IExprContext {
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

func (s *PrintexprlistContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *PrintexprlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintexprlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintexprlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintexprlist(s)
	}
}

func (s *PrintexprlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintexprlist(s)
	}
}

func (p *SwiftGrammarParser) Printexprlist() (localctx IPrintexprlistContext) {
	localctx = NewPrintexprlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_printexprlist)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)

			var _x = p.expr(0)

			localctx.(*PrintexprlistContext)._expr = _x
		}
		{
			p.SetState(178)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintexprlistContext).prntexpr = localctx.(*PrintexprlistContext).Get_expr().GetE()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)

			var _x = p.expr(0)

			localctx.(*PrintexprlistContext)._expr = _x
		}
		localctx.(*PrintexprlistContext).prntexpr = localctx.(*PrintexprlistContext).Get_expr().GetE()

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

	// GetFirstBlk returns the firstBlk rule contexts.
	GetFirstBlk() IBlockContext

	// Get_elseifstmt returns the _elseifstmt rule contexts.
	Get_elseifstmt() IElseifstmtContext

	// Get_elsestmt returns the _elsestmt rule contexts.
	Get_elsestmt() IElsestmtContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetFirstBlk sets the firstBlk rule contexts.
	SetFirstBlk(IBlockContext)

	// Set_elseifstmt sets the _elseifstmt rule contexts.
	Set_elseifstmt(IElseifstmtContext)

	// Set_elsestmt sets the _elsestmt rule contexts.
	Set_elsestmt(IElsestmtContext)

	// GetElif returns the elif rule context list.
	GetElif() []IElseifstmtContext

	// SetElif sets the elif rule context list.
	SetElif([]IElseifstmtContext)

	// GetIfinstr returns the ifinstr attribute.
	GetIfinstr() interfaces.Instruction

	// SetIfinstr sets the ifinstr attribute.
	SetIfinstr(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	Elsestmt() IElsestmtContext
	AllElseifstmt() []IElseifstmtContext
	Elseifstmt(i int) IElseifstmtContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	ifinstr     interfaces.Instruction
	_IF         antlr.Token
	_expr       IExprContext
	_block      IBlockContext
	firstBlk    IBlockContext
	_elseifstmt IElseifstmtContext
	elif        []IElseifstmtContext
	_elsestmt   IElsestmtContext
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

func (s *IfstmtContext) GetFirstBlk() IBlockContext { return s.firstBlk }

func (s *IfstmtContext) Get_elseifstmt() IElseifstmtContext { return s._elseifstmt }

func (s *IfstmtContext) Get_elsestmt() IElsestmtContext { return s._elsestmt }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfstmtContext) SetFirstBlk(v IBlockContext) { s.firstBlk = v }

func (s *IfstmtContext) Set_elseifstmt(v IElseifstmtContext) { s._elseifstmt = v }

func (s *IfstmtContext) Set_elsestmt(v IElsestmtContext) { s._elsestmt = v }

func (s *IfstmtContext) GetElif() []IElseifstmtContext { return s.elif }

func (s *IfstmtContext) SetElif(v []IElseifstmtContext) { s.elif = v }

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

func (s *IfstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *IfstmtContext) Block() IBlockContext {
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

func (s *IfstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *IfstmtContext) Elsestmt() IElsestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsestmtContext)
}

func (s *IfstmtContext) AllElseifstmt() []IElseifstmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifstmtContext); ok {
			len++
		}
	}

	tst := make([]IElseifstmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifstmtContext); ok {
			tst[i] = t.(IElseifstmtContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Elseifstmt(i int) IElseifstmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifstmtContext); ok {
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

	return t.(IElseifstmtContext)
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
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_ifstmt)
	var _la int

	var _alt int

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(188)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(190)
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
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), nil, nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(195)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)

			var _x = p.Block()

			localctx.(*IfstmtContext).firstBlk = _x
		}
		{
			p.SetState(197)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(198)

					var _x = p.Elseifstmt()

					localctx.(*IfstmtContext)._elseifstmt = _x
				}
				localctx.(*IfstmtContext).elif = append(localctx.(*IfstmtContext).elif, localctx.(*IfstmtContext)._elseifstmt)

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(203)

			var _x = p.Elsestmt()

			localctx.(*IfstmtContext)._elsestmt = _x
		}

		var ifInterfaces []interface{}
		// fmt.Println(localctx.(*IfstmtContext).GetElif())
		for _, e := range localctx.(*IfstmtContext).GetElif() {
			ifInterfaces = append(ifInterfaces, e.GetElifinstr())
			fmt.Println(e.GetElifinstr())
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
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetFirstBlk().GetBlk(), ifInterfaces, localctx.(*IfstmtContext).Get_elsestmt().GetElseinstr())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(208)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)

			var _x = p.Block()

			localctx.(*IfstmtContext).firstBlk = _x
		}
		{
			p.SetState(210)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SwiftGrammarParserELSE {
			{
				p.SetState(211)

				var _x = p.Elseifstmt()

				localctx.(*IfstmtContext)._elseifstmt = _x
			}
			localctx.(*IfstmtContext).elif = append(localctx.(*IfstmtContext).elif, localctx.(*IfstmtContext)._elseifstmt)

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		var ifInterfaces []interface{}
		// fmt.Println(localctx.(*IfstmtContext).GetElif())
		for _, e := range localctx.(*IfstmtContext).GetElif() {
			ifInterfaces = append(ifInterfaces, e.GetElifinstr())
			fmt.Println(e.GetElifinstr())
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
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetFirstBlk().GetBlk(), ifInterfaces, nil)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(218)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(220)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)

			var _x = p.Block()

			localctx.(*IfstmtContext).firstBlk = _x
		}
		{
			p.SetState(222)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)

			var _x = p.Elsestmt()

			localctx.(*IfstmtContext)._elsestmt = _x
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
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetFirstBlk().GetBlk(), nil, localctx.(*IfstmtContext).Get_elsestmt().GetElseinstr())

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

// IElseifstmtContext is an interface to support dynamic dispatch.
type IElseifstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ELSE returns the _ELSE token.
	Get_ELSE() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetElifinstr returns the elifinstr attribute.
	GetElifinstr() interfaces.Instruction

	// SetElifinstr sets the elifinstr attribute.
	SetElifinstr(interfaces.Instruction)

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsElseifstmtContext differentiates from other interfaces.
	IsElseifstmtContext()
}

type ElseifstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	elifinstr interfaces.Instruction
	_ELSE     antlr.Token
	_expr     IExprContext
	_block    IBlockContext
}

func NewEmptyElseifstmtContext() *ElseifstmtContext {
	var p = new(ElseifstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseifstmt
	return p
}

func InitEmptyElseifstmtContext(p *ElseifstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseifstmt
}

func (*ElseifstmtContext) IsElseifstmtContext() {}

func NewElseifstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifstmtContext {
	var p = new(ElseifstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elseifstmt

	return p
}

func (s *ElseifstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifstmtContext) Get_ELSE() antlr.Token { return s._ELSE }

func (s *ElseifstmtContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *ElseifstmtContext) Get_expr() IExprContext { return s._expr }

func (s *ElseifstmtContext) Get_block() IBlockContext { return s._block }

func (s *ElseifstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ElseifstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ElseifstmtContext) GetElifinstr() interfaces.Instruction { return s.elifinstr }

func (s *ElseifstmtContext) SetElifinstr(v interfaces.Instruction) { s.elifinstr = v }

func (s *ElseifstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *ElseifstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *ElseifstmtContext) Expr() IExprContext {
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

func (s *ElseifstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElseifstmtContext) Block() IBlockContext {
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

func (s *ElseifstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElseifstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElseifstmt(s)
	}
}

func (s *ElseifstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElseifstmt(s)
	}
}

func (p *SwiftGrammarParser) Elseifstmt() (localctx IElseifstmtContext) {
	localctx = NewElseifstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_elseifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)

		var _m = p.Match(SwiftGrammarParserELSE)

		localctx.(*ElseifstmtContext)._ELSE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)

		var _x = p.expr(0)

		localctx.(*ElseifstmtContext)._expr = _x
	}
	{
		p.SetState(231)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)

		var _x = p.Block()

		localctx.(*ElseifstmtContext)._block = _x
	}
	{
		p.SetState(233)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*ElseifstmtContext).elifinstr = instructions.NewIf((func() int {
		if localctx.(*ElseifstmtContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*ElseifstmtContext).Get_ELSE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ElseifstmtContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*ElseifstmtContext).Get_ELSE().GetColumn()
		}
	}()), localctx.(*ElseifstmtContext).Get_expr().GetE(), nil, localctx.(*ElseifstmtContext).Get_block().GetBlk(), nil)

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

// IElsestmtContext is an interface to support dynamic dispatch.
type IElsestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetElseinstr returns the elseinstr attribute.
	GetElseinstr() []interface{}

	// SetElseinstr sets the elseinstr attribute.
	SetElseinstr([]interface{})

	// Getter signatures
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsElsestmtContext differentiates from other interfaces.
	IsElsestmtContext()
}

type ElsestmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	elseinstr []interface{}
	_block    IBlockContext
}

func NewEmptyElsestmtContext() *ElsestmtContext {
	var p = new(ElsestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elsestmt
	return p
}

func InitEmptyElsestmtContext(p *ElsestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elsestmt
}

func (*ElsestmtContext) IsElsestmtContext() {}

func NewElsestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsestmtContext {
	var p = new(ElsestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elsestmt

	return p
}

func (s *ElsestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsestmtContext) Get_block() IBlockContext { return s._block }

func (s *ElsestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ElsestmtContext) GetElseinstr() []interface{} { return s.elseinstr }

func (s *ElsestmtContext) SetElseinstr(v []interface{}) { s.elseinstr = v }

func (s *ElsestmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *ElsestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElsestmtContext) Block() IBlockContext {
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

func (s *ElsestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElsestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElsestmt(s)
	}
}

func (s *ElsestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElsestmt(s)
	}
}

func (p *SwiftGrammarParser) Elsestmt() (localctx IElsestmtContext) {
	localctx = NewElsestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_elsestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)

		var _x = p.Block()

		localctx.(*ElsestmtContext)._block = _x
	}
	{
		p.SetState(239)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ElsestmtContext).elseinstr = localctx.(*ElsestmtContext).Get_block().GetBlk()

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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_varstmt)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)

			var _x = p.Tipo()

			localctx.(*VarstmtContext)._tipo = _x
		}
		{
			p.SetState(246)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)

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
			p.SetState(250)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

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
			p.SetState(256)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)

			var _x = p.Tipo()

			localctx.(*VarstmtContext)._tipo = _x
		}
		{
			p.SetState(260)

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
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_tipo)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
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
			p.SetState(267)
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
			p.SetState(269)
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
			p.SetState(271)
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
			p.SetState(273)
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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_varasgmt)
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

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
			p.SetState(282)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(SwiftGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)

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
			p.SetState(288)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)

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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_conststmt)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*ConststmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConststmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)

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
			p.SetState(302)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*ConststmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConststmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)

			var _x = p.Tipo()

			localctx.(*ConststmtContext)._tipo = _x
		}
		{
			p.SetState(306)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)

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

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_casestmt returns the _casestmt rule contexts.
	Get_casestmt() ICasestmtContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_casestmt sets the _casestmt rule contexts.
	Set_casestmt(ICasestmtContext)

	// GetCasesvar returns the casesvar rule context list.
	GetCasesvar() []ICasestmtContext

	// SetCasesvar sets the casesvar rule context list.
	SetCasesvar([]ICasestmtContext)

	// GetSwitchinstr returns the switchinstr attribute.
	GetSwitchinstr() interfaces.Instruction

	// SetSwitchinstr sets the switchinstr attribute.
	SetSwitchinstr(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Defaultstmt() IDefaultstmtContext
	AllCasestmt() []ICasestmtContext
	Casestmt(i int) ICasestmtContext

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	switchinstr interfaces.Instruction
	_SWITCH     antlr.Token
	_expr       IExprContext
	_casestmt   ICasestmtContext
	casesvar    []ICasestmtContext
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *SwitchstmtContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *SwitchstmtContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchstmtContext) Get_casestmt() ICasestmtContext { return s._casestmt }

func (s *SwitchstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchstmtContext) Set_casestmt(v ICasestmtContext) { s._casestmt = v }

func (s *SwitchstmtContext) GetCasesvar() []ICasestmtContext { return s.casesvar }

func (s *SwitchstmtContext) SetCasesvar(v []ICasestmtContext) { s.casesvar = v }

func (s *SwitchstmtContext) GetSwitchinstr() interfaces.Instruction { return s.switchinstr }

func (s *SwitchstmtContext) SetSwitchinstr(v interfaces.Instruction) { s.switchinstr = v }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
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

func (s *SwitchstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchstmtContext) Defaultstmt() IDefaultstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultstmtContext)
}

func (s *SwitchstmtContext) AllCasestmt() []ICasestmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICasestmtContext); ok {
			len++
		}
	}

	tst := make([]ICasestmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICasestmtContext); ok {
			tst[i] = t.(ICasestmtContext)
			i++
		}
	}

	return tst
}

func (s *SwitchstmtContext) Casestmt(i int) ICasestmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasestmtContext); ok {
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

	return t.(ICasestmtContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (p *SwiftGrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_switchstmt)

	var switchInterfaces = []interface{}{}
	var interfacelist []ICasestmtContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)

		var _m = p.Match(SwiftGrammarParserSWITCH)

		localctx.(*SwitchstmtContext)._SWITCH = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)

		var _x = p.expr(0)

		localctx.(*SwitchstmtContext)._expr = _x
	}
	{
		p.SetState(314)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCASE {
		{
			p.SetState(315)

			var _x = p.Casestmt()

			localctx.(*SwitchstmtContext)._casestmt = _x
		}
		localctx.(*SwitchstmtContext).casesvar = append(localctx.(*SwitchstmtContext).casesvar, localctx.(*SwitchstmtContext)._casestmt)

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDEFAULT {
		{
			p.SetState(320)
			p.Defaultstmt()
		}

	}
	{
		p.SetState(323)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	interfacelist = localctx.(*SwitchstmtContext).GetCasesvar()
	for _, e := range interfacelist {
		switchInterfaces = append(switchInterfaces, e.GetCaseinstr())
	}
	localctx.(*SwitchstmtContext).switchinstr = instructions.NewSwitch((func() int {
		if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
			return 0
		} else {
			return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
		}
	}()), (func() int {
		if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
			return 0
		} else {
			return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
		}
	}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), switchInterfaces, nil)

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

// ICasestmtContext is an interface to support dynamic dispatch.
type ICasestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCaseinstr returns the caseinstr attribute.
	GetCaseinstr() interfaces.Instruction

	// SetCaseinstr sets the caseinstr attribute.
	SetCaseinstr(interfaces.Instruction)

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	COLON() antlr.TerminalNode
	Block() IBlockContext
	BREAK() antlr.TerminalNode

	// IsCasestmtContext differentiates from other interfaces.
	IsCasestmtContext()
}

type CasestmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	caseinstr interfaces.Instruction
	_CASE     antlr.Token
	_expr     IExprContext
	_block    IBlockContext
}

func NewEmptyCasestmtContext() *CasestmtContext {
	var p = new(CasestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_casestmt
	return p
}

func InitEmptyCasestmtContext(p *CasestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_casestmt
}

func (*CasestmtContext) IsCasestmtContext() {}

func NewCasestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasestmtContext {
	var p = new(CasestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_casestmt

	return p
}

func (s *CasestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CasestmtContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CasestmtContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CasestmtContext) Get_expr() IExprContext { return s._expr }

func (s *CasestmtContext) Get_block() IBlockContext { return s._block }

func (s *CasestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CasestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *CasestmtContext) GetCaseinstr() interfaces.Instruction { return s.caseinstr }

func (s *CasestmtContext) SetCaseinstr(v interfaces.Instruction) { s.caseinstr = v }

func (s *CasestmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *CasestmtContext) Expr() IExprContext {
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

func (s *CasestmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *CasestmtContext) Block() IBlockContext {
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

func (s *CasestmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *CasestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCasestmt(s)
	}
}

func (s *CasestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCasestmt(s)
	}
}

func (p *SwiftGrammarParser) Casestmt() (localctx ICasestmtContext) {
	localctx = NewCasestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_casestmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*CasestmtContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)

		var _x = p.expr(0)

		localctx.(*CasestmtContext)._expr = _x
	}
	{
		p.SetState(328)
		p.Match(SwiftGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)

		var _x = p.Block()

		localctx.(*CasestmtContext)._block = _x
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserBREAK {
		{
			p.SetState(330)
			p.Match(SwiftGrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

	localctx.(*CasestmtContext).caseinstr = instructions.NewSwitch((func() int {
		if localctx.(*CasestmtContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasestmtContext).Get_CASE().GetLine()
		}
	}()), (func() int {
		if localctx.(*CasestmtContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasestmtContext).Get_CASE().GetColumn()
		}
	}()), localctx.(*CasestmtContext).Get_expr().GetE(), localctx.(*CasestmtContext).Get_block().GetBlk(), nil)

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

// IDefaultstmtContext is an interface to support dynamic dispatch.
type IDefaultstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetDefinstr returns the definstr attribute.
	GetDefinstr() []interface{}

	// SetDefinstr sets the definstr attribute.
	SetDefinstr([]interface{})

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	COLON() antlr.TerminalNode
	Block() IBlockContext
	BREAK() antlr.TerminalNode

	// IsDefaultstmtContext differentiates from other interfaces.
	IsDefaultstmtContext()
}

type DefaultstmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	definstr []interface{}
	_block   IBlockContext
}

func NewEmptyDefaultstmtContext() *DefaultstmtContext {
	var p = new(DefaultstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultstmt
	return p
}

func InitEmptyDefaultstmtContext(p *DefaultstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultstmt
}

func (*DefaultstmtContext) IsDefaultstmtContext() {}

func NewDefaultstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultstmtContext {
	var p = new(DefaultstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defaultstmt

	return p
}

func (s *DefaultstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultstmtContext) Get_block() IBlockContext { return s._block }

func (s *DefaultstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *DefaultstmtContext) GetDefinstr() []interface{} { return s.definstr }

func (s *DefaultstmtContext) SetDefinstr(v []interface{}) { s.definstr = v }

func (s *DefaultstmtContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *DefaultstmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *DefaultstmtContext) Block() IBlockContext {
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

func (s *DefaultstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *DefaultstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefaultstmt(s)
	}
}

func (s *DefaultstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefaultstmt(s)
	}
}

func (p *SwiftGrammarParser) Defaultstmt() (localctx IDefaultstmtContext) {
	localctx = NewDefaultstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_defaultstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(SwiftGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(SwiftGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)

		var _x = p.Block()

		localctx.(*DefaultstmtContext)._block = _x
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserBREAK {
		{
			p.SetState(338)
			p.Match(SwiftGrammarParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	localctx.(*DefaultstmtContext).definstr = localctx.(*DefaultstmtContext).Get_block().GetBlk()

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

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetWhileinstr returns the whileinstr attribute.
	GetWhileinstr() interfaces.Instruction

	// SetWhileinstr sets the whileinstr attribute.
	SetWhileinstr(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	whileinstr interfaces.Instruction
	_WHILE     antlr.Token
	_expr      IExprContext
	_block     IBlockContext
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilestmtContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilestmtContext) Get_expr() IExprContext { return s._expr }

func (s *WhilestmtContext) Get_block() IBlockContext { return s._block }

func (s *WhilestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *WhilestmtContext) GetWhileinstr() interfaces.Instruction { return s.whileinstr }

func (s *WhilestmtContext) SetWhileinstr(v interfaces.Instruction) { s.whileinstr = v }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
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

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
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

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(345)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(347)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilestmtContext).whileinstr = instructions.NewWhile((func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilestmtContext).Get_expr().GetE(), localctx.(*WhilestmtContext).Get_block().GetBlk())

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

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// GetFirst returns the first token.
	GetFirst() antlr.Token

	// GetSecond returns the second token.
	GetSecond() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// SetFirst sets the first token.
	SetFirst(antlr.Token)

	// SetSecond sets the second token.
	SetSecond(antlr.Token)

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetForinstr returns the forinstr attribute.
	GetForinstr() interfaces.Instruction

	// SetForinstr sets the forinstr attribute.
	SetForinstr(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	IN() antlr.TerminalNode
	STRING() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	RANGEPTS() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	forinstr interfaces.Instruction
	_FOR     antlr.Token
	_ID      antlr.Token
	_STRING  antlr.Token
	_block   IBlockContext
	left     IExprContext
	right    IExprContext
	first    antlr.Token
	second   antlr.Token
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ForstmtContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ForstmtContext) GetFirst() antlr.Token { return s.first }

func (s *ForstmtContext) GetSecond() antlr.Token { return s.second }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ForstmtContext) SetFirst(v antlr.Token) { s.first = v }

func (s *ForstmtContext) SetSecond(v antlr.Token) { s.second = v }

func (s *ForstmtContext) Get_block() IBlockContext { return s._block }

func (s *ForstmtContext) GetLeft() IExprContext { return s.left }

func (s *ForstmtContext) GetRight() IExprContext { return s.right }

func (s *ForstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ForstmtContext) SetLeft(v IExprContext) { s.left = v }

func (s *ForstmtContext) SetRight(v IExprContext) { s.right = v }

func (s *ForstmtContext) GetForinstr() interfaces.Instruction { return s.forinstr }

func (s *ForstmtContext) SetForinstr(v interfaces.Instruction) { s.forinstr = v }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ForstmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
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

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) RANGEPTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRANGEPTS, 0)
}

func (s *ForstmtContext) AllExpr() []IExprContext {
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

func (s *ForstmtContext) Expr(i int) IExprContext {
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

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_forstmt)

	var cadena interfaces.Expression
	var str string

	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ForstmtContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(356)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str = (func() string {
			if localctx.(*ForstmtContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_STRING().GetText()
			}
		}())
		cadena = expressions.NewPrimitive((func() int {
			if localctx.(*ForstmtContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)
		localctx.(*ForstmtContext).forinstr = instructions.NewForIn((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), cadena, nil, nil, "", localctx.(*ForstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).left = _x
		}
		{
			p.SetState(363)
			p.Match(SwiftGrammarParserRANGEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).right = _x
		}
		{
			p.SetState(365)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(367)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ForstmtContext).forinstr = instructions.NewForIn((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*ForstmtContext).GetLeft().GetE(), localctx.(*ForstmtContext).GetRight().GetE(), "", localctx.(*ForstmtContext).Get_block().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(370)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext).first = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext).second = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(376)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ForstmtContext).forinstr = instructions.NewForIn((func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForstmtContext).Get_FOR() == nil {
				return 0
			} else {
				return localctx.(*ForstmtContext).Get_FOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForstmtContext).GetFirst() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).GetFirst().GetText()
			}
		}()), nil, nil, nil, (func() string {
			if localctx.(*ForstmtContext).GetSecond() == nil {
				return ""
			} else {
				return localctx.(*ForstmtContext).GetSecond().GetText()
			}
		}()), localctx.(*ForstmtContext).Get_block().GetBlk())

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

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetGuardinstr returns the guardinstr attribute.
	GetGuardinstr() interfaces.Instruction

	// SetGuardinstr sets the guardinstr attribute.
	SetGuardinstr(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	guardinstr interfaces.Instruction
	_GUARD     antlr.Token
	_expr      IExprContext
	_block     IBlockContext
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardstmtContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardstmtContext) Get_expr() IExprContext { return s._expr }

func (s *GuardstmtContext) Get_block() IBlockContext { return s._block }

func (s *GuardstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *GuardstmtContext) GetGuardinstr() interfaces.Instruction { return s.guardinstr }

func (s *GuardstmtContext) SetGuardinstr(v interfaces.Instruction) { s.guardinstr = v }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
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

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
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

func (s *GuardstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (p *SwiftGrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(381)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(383)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(386)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*GuardstmtContext).guardinstr = instructions.NewGuard((func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardstmtContext).Get_expr().GetE(), localctx.(*GuardstmtContext).Get_block().GetBlk())

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

// ITransferstmtContext is an interface to support dynamic dispatch.
type ITransferstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetTrns returns the trns attribute.
	GetTrns() interfaces.Instruction

	// SetTrns sets the trns attribute.
	SetTrns(interfaces.Instruction)

	// Getter signatures
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsTransferstmtContext differentiates from other interfaces.
	IsTransferstmtContext()
}

type TransferstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	trns      interfaces.Instruction
	_BREAK    antlr.Token
	_CONTINUE antlr.Token
	_RETURN   antlr.Token
	_expr     IExprContext
}

func NewEmptyTransferstmtContext() *TransferstmtContext {
	var p = new(TransferstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_transferstmt
	return p
}

func InitEmptyTransferstmtContext(p *TransferstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_transferstmt
}

func (*TransferstmtContext) IsTransferstmtContext() {}

func NewTransferstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferstmtContext {
	var p = new(TransferstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_transferstmt

	return p
}

func (s *TransferstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferstmtContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *TransferstmtContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *TransferstmtContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *TransferstmtContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *TransferstmtContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *TransferstmtContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *TransferstmtContext) Get_expr() IExprContext { return s._expr }

func (s *TransferstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *TransferstmtContext) GetTrns() interfaces.Instruction { return s.trns }

func (s *TransferstmtContext) SetTrns(v interfaces.Instruction) { s.trns = v }

func (s *TransferstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *TransferstmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *TransferstmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *TransferstmtContext) Expr() IExprContext {
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

func (s *TransferstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransferstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTransferstmt(s)
	}
}

func (s *TransferstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTransferstmt(s)
	}
}

func (p *SwiftGrammarParser) Transferstmt() (localctx ITransferstmtContext) {
	localctx = NewTransferstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_transferstmt)
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)

			var _m = p.Match(SwiftGrammarParserBREAK)

			localctx.(*TransferstmtContext)._BREAK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TransferstmtContext).trns = instructions.NewBreak((func() int {
			if localctx.(*TransferstmtContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_BREAK().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_BREAK().GetColumn()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(391)

			var _m = p.Match(SwiftGrammarParserCONTINUE)

			localctx.(*TransferstmtContext)._CONTINUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TransferstmtContext).trns = instructions.NewContinue((func() int {
			if localctx.(*TransferstmtContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_CONTINUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_CONTINUE().GetColumn()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(393)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TransferstmtContext).trns = instructions.NewReturn((func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetColumn()
			}
		}()), nil)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(395)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)

			var _x = p.expr(0)

			localctx.(*TransferstmtContext)._expr = _x
		}
		localctx.(*TransferstmtContext).trns = instructions.NewReturn((func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*TransferstmtContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*TransferstmtContext).Get_RETURN().GetColumn()
			}
		}()), localctx.(*TransferstmtContext).Get_expr().GetE())

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

// IVectorstmtContext is an interface to support dynamic dispatch.
type IVectorstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_definestmt returns the _definestmt rule contexts.
	Get_definestmt() IDefinestmtContext

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_definestmt sets the _definestmt rule contexts.
	Set_definestmt(IDefinestmtContext)

	// GetVectorinstr returns the vectorinstr attribute.
	GetVectorinstr() interfaces.Instruction

	// SetVectorinstr sets the vectorinstr attribute.
	SetVectorinstr(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	Tipo() ITipoContext
	CORCHDER() antlr.TerminalNode
	Definestmt() IDefinestmtContext

	// IsVectorstmtContext differentiates from other interfaces.
	IsVectorstmtContext()
}

type VectorstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	vectorinstr interfaces.Instruction
	_VAR        antlr.Token
	_ID         antlr.Token
	_tipo       ITipoContext
	_definestmt IDefinestmtContext
}

func NewEmptyVectorstmtContext() *VectorstmtContext {
	var p = new(VectorstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorstmt
	return p
}

func InitEmptyVectorstmtContext(p *VectorstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectorstmt
}

func (*VectorstmtContext) IsVectorstmtContext() {}

func NewVectorstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorstmtContext {
	var p = new(VectorstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectorstmt

	return p
}

func (s *VectorstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorstmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *VectorstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *VectorstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *VectorstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VectorstmtContext) Get_tipo() ITipoContext { return s._tipo }

func (s *VectorstmtContext) Get_definestmt() IDefinestmtContext { return s._definestmt }

func (s *VectorstmtContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *VectorstmtContext) Set_definestmt(v IDefinestmtContext) { s._definestmt = v }

func (s *VectorstmtContext) GetVectorinstr() interfaces.Instruction { return s.vectorinstr }

func (s *VectorstmtContext) SetVectorinstr(v interfaces.Instruction) { s.vectorinstr = v }

func (s *VectorstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *VectorstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VectorstmtContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *VectorstmtContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *VectorstmtContext) Tipo() ITipoContext {
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

func (s *VectorstmtContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *VectorstmtContext) Definestmt() IDefinestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinestmtContext)
}

func (s *VectorstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectorstmt(s)
	}
}

func (s *VectorstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectorstmt(s)
	}
}

func (p *SwiftGrammarParser) Vectorstmt() (localctx IVectorstmtContext) {
	localctx = NewVectorstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_vectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)

		var _m = p.Match(SwiftGrammarParserVAR)

		localctx.(*VectorstmtContext)._VAR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VectorstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Match(SwiftGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)

		var _x = p.Tipo()

		localctx.(*VectorstmtContext)._tipo = _x
	}
	{
		p.SetState(406)
		p.Match(SwiftGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)

		var _x = p.Definestmt()

		localctx.(*VectorstmtContext)._definestmt = _x
	}

	localctx.(*VectorstmtContext).vectorinstr = instructions.NewVectorStmt((func() int {
		if localctx.(*VectorstmtContext).Get_VAR() == nil {
			return 0
		} else {
			return localctx.(*VectorstmtContext).Get_VAR().GetLine()
		}
	}()), (func() int {
		if localctx.(*VectorstmtContext).Get_VAR() == nil {
			return 0
		} else {
			return localctx.(*VectorstmtContext).Get_VAR().GetColumn()
		}
	}()), (func() string {
		if localctx.(*VectorstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*VectorstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*VectorstmtContext).Get_tipo().GetRtipo(), localctx.(*VectorstmtContext).Get_definestmt().GetDefineinstr())

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

// IDefinestmtContext is an interface to support dynamic dispatch.
type IDefinestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listexpr returns the _listexpr rule contexts.
	Get_listexpr() IListexprContext

	// Set_listexpr sets the _listexpr rule contexts.
	Set_listexpr(IListexprContext)

	// GetLista returns the lista rule context list.
	GetLista() []IListexprContext

	// SetLista sets the lista rule context list.
	SetLista([]IListexprContext)

	// GetDefineinstr returns the defineinstr attribute.
	GetDefineinstr() []interface{}

	// SetDefineinstr sets the defineinstr attribute.
	SetDefineinstr([]interface{})

	// Getter signatures
	IG() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	CORCHDER() antlr.TerminalNode
	AllListexpr() []IListexprContext
	Listexpr(i int) IListexprContext

	// IsDefinestmtContext differentiates from other interfaces.
	IsDefinestmtContext()
}

type DefinestmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	defineinstr []interface{}
	_listexpr   IListexprContext
	lista       []IListexprContext
}

func NewEmptyDefinestmtContext() *DefinestmtContext {
	var p = new(DefinestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_definestmt
	return p
}

func InitEmptyDefinestmtContext(p *DefinestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_definestmt
}

func (*DefinestmtContext) IsDefinestmtContext() {}

func NewDefinestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinestmtContext {
	var p = new(DefinestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_definestmt

	return p
}

func (s *DefinestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinestmtContext) Get_listexpr() IListexprContext { return s._listexpr }

func (s *DefinestmtContext) Set_listexpr(v IListexprContext) { s._listexpr = v }

func (s *DefinestmtContext) GetLista() []IListexprContext { return s.lista }

func (s *DefinestmtContext) SetLista(v []IListexprContext) { s.lista = v }

func (s *DefinestmtContext) GetDefineinstr() []interface{} { return s.defineinstr }

func (s *DefinestmtContext) SetDefineinstr(v []interface{}) { s.defineinstr = v }

func (s *DefinestmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DefinestmtContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *DefinestmtContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *DefinestmtContext) AllListexpr() []IListexprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListexprContext); ok {
			len++
		}
	}

	tst := make([]IListexprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListexprContext); ok {
			tst[i] = t.(IListexprContext)
			i++
		}
	}

	return tst
}

func (s *DefinestmtContext) Listexpr(i int) IListexprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListexprContext); ok {
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

	return t.(IListexprContext)
}

func (s *DefinestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefinestmt(s)
	}
}

func (s *DefinestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefinestmt(s)
	}
}

func (p *SwiftGrammarParser) Definestmt() (localctx IDefinestmtContext) {
	localctx = NewDefinestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_definestmt)

	var defVecInterfaces []interface{}

	var _la int

	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54048624367108566) != 0) {
			{
				p.SetState(412)

				var _x = p.Listexpr()

				localctx.(*DefinestmtContext)._listexpr = _x
			}
			localctx.(*DefinestmtContext).lista = append(localctx.(*DefinestmtContext).lista, localctx.(*DefinestmtContext)._listexpr)

			p.SetState(415)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(417)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		for _, e := range localctx.(*DefinestmtContext).GetLista() {
			// fmt.Println(fmt.Sprintf("%T", e.GetListe()))
			defVecInterfaces = append(defVecInterfaces, e.GetListe())
		}
		localctx.(*DefinestmtContext).defineinstr = defVecInterfaces

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DefinestmtContext).defineinstr = nil

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

// IListexprContext is an interface to support dynamic dispatch.
type IListexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetListe returns the liste attribute.
	GetListe() interfaces.Expression

	// SetListe sets the liste attribute.
	SetListe(interfaces.Expression)

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode

	// IsListexprContext differentiates from other interfaces.
	IsListexprContext()
}

type ListexprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	liste  interfaces.Expression
	_expr  IExprContext
}

func NewEmptyListexprContext() *ListexprContext {
	var p = new(ListexprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listexpr
	return p
}

func InitEmptyListexprContext(p *ListexprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listexpr
}

func (*ListexprContext) IsListexprContext() {}

func NewListexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListexprContext {
	var p = new(ListexprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listexpr

	return p
}

func (s *ListexprContext) GetParser() antlr.Parser { return s.parser }

func (s *ListexprContext) Get_expr() IExprContext { return s._expr }

func (s *ListexprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListexprContext) GetListe() interfaces.Expression { return s.liste }

func (s *ListexprContext) SetListe(v interfaces.Expression) { s.liste = v }

func (s *ListexprContext) Expr() IExprContext {
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

func (s *ListexprContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListexprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListexpr(s)
	}
}

func (s *ListexprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListexpr(s)
	}
}

func (p *SwiftGrammarParser) Listexpr() (localctx IListexprContext) {
	localctx = NewListexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_listexpr)
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)

			var _x = p.expr(0)

			localctx.(*ListexprContext)._expr = _x
		}
		{
			p.SetState(427)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ListexprContext).liste = localctx.(*ListexprContext).Get_expr().GetE()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(430)

			var _x = p.expr(0)

			localctx.(*ListexprContext)._expr = _x
		}
		localctx.(*ListexprContext).liste = localctx.(*ListexprContext).Get_expr().GetE()

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

// IMethodvecContext is an interface to support dynamic dispatch.
type IMethodvecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_APPEND returns the _APPEND token.
	Get_APPEND() antlr.Token

	// Get_REMOVELAST returns the _REMOVELAST token.
	Get_REMOVELAST() antlr.Token

	// Get_REMOVE returns the _REMOVE token.
	Get_REMOVE() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_APPEND sets the _APPEND token.
	Set_APPEND(antlr.Token)

	// Set_REMOVELAST sets the _REMOVELAST token.
	Set_REMOVELAST(antlr.Token)

	// Set_REMOVE sets the _REMOVE token.
	Set_REMOVE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetMethodinstr returns the methodinstr attribute.
	GetMethodinstr() interfaces.Instruction

	// SetMethodinstr sets the methodinstr attribute.
	SetMethodinstr(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	REMOVE() antlr.TerminalNode

	// IsMethodvecContext differentiates from other interfaces.
	IsMethodvecContext()
}

type MethodvecContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	methodinstr interfaces.Instruction
	_ID         antlr.Token
	_APPEND     antlr.Token
	_expr       IExprContext
	_REMOVELAST antlr.Token
	_REMOVE     antlr.Token
}

func NewEmptyMethodvecContext() *MethodvecContext {
	var p = new(MethodvecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_methodvec
	return p
}

func InitEmptyMethodvecContext(p *MethodvecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_methodvec
}

func (*MethodvecContext) IsMethodvecContext() {}

func NewMethodvecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodvecContext {
	var p = new(MethodvecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_methodvec

	return p
}

func (s *MethodvecContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodvecContext) Get_ID() antlr.Token { return s._ID }

func (s *MethodvecContext) Get_APPEND() antlr.Token { return s._APPEND }

func (s *MethodvecContext) Get_REMOVELAST() antlr.Token { return s._REMOVELAST }

func (s *MethodvecContext) Get_REMOVE() antlr.Token { return s._REMOVE }

func (s *MethodvecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *MethodvecContext) Set_APPEND(v antlr.Token) { s._APPEND = v }

func (s *MethodvecContext) Set_REMOVELAST(v antlr.Token) { s._REMOVELAST = v }

func (s *MethodvecContext) Set_REMOVE(v antlr.Token) { s._REMOVE = v }

func (s *MethodvecContext) Get_expr() IExprContext { return s._expr }

func (s *MethodvecContext) Set_expr(v IExprContext) { s._expr = v }

func (s *MethodvecContext) GetMethodinstr() interfaces.Instruction { return s.methodinstr }

func (s *MethodvecContext) SetMethodinstr(v interfaces.Instruction) { s.methodinstr = v }

func (s *MethodvecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *MethodvecContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *MethodvecContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAPPEND, 0)
}

func (s *MethodvecContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *MethodvecContext) Expr() IExprContext {
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

func (s *MethodvecContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *MethodvecContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVELAST, 0)
}

func (s *MethodvecContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE, 0)
}

func (s *MethodvecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodvecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodvecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterMethodvec(s)
	}
}

func (s *MethodvecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitMethodvec(s)
	}
}

func (p *SwiftGrammarParser) Methodvec() (localctx IMethodvecContext) {
	localctx = NewMethodvecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_methodvec)
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)

			var _m = p.Match(SwiftGrammarParserAPPEND)

			localctx.(*MethodvecContext)._APPEND = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)

			var _x = p.expr(0)

			localctx.(*MethodvecContext)._expr = _x
		}
		{
			p.SetState(440)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MethodvecContext).methodinstr = instructions.NewVectorMethod((func() int {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetText()
			}
		}()), localctx.(*MethodvecContext).Get_expr().GetE(), (func() string {
			if localctx.(*MethodvecContext).Get_APPEND() == nil {
				return ""
			} else {
				return localctx.(*MethodvecContext).Get_APPEND().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)

			var _m = p.Match(SwiftGrammarParserREMOVELAST)

			localctx.(*MethodvecContext)._REMOVELAST = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MethodvecContext).methodinstr = instructions.NewVectorMethod((func() int {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetText()
			}
		}()), nil, (func() string {
			if localctx.(*MethodvecContext).Get_REMOVELAST() == nil {
				return ""
			} else {
				return localctx.(*MethodvecContext).Get_REMOVELAST().GetText()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(449)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)

			var _m = p.Match(SwiftGrammarParserREMOVE)

			localctx.(*MethodvecContext)._REMOVE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)

			var _x = p.expr(0)

			localctx.(*MethodvecContext)._expr = _x
		}
		{
			p.SetState(454)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MethodvecContext).methodinstr = instructions.NewVectorMethod((func() int {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MethodvecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MethodvecContext).Get_ID().GetText()
			}
		}()), localctx.(*MethodvecContext).Get_expr().GetE(), (func() string {
			if localctx.(*MethodvecContext).Get_REMOVE() == nil {
				return ""
			} else {
				return localctx.(*MethodvecContext).Get_REMOVE().GetText()
			}
		}()))

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

// IMethodvecrtrnContext is an interface to support dynamic dispatch.
type IMethodvecrtrnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_EMPTY returns the _EMPTY token.
	Get_EMPTY() antlr.Token

	// Get_COUNT returns the _COUNT token.
	Get_COUNT() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_EMPTY sets the _EMPTY token.
	Set_EMPTY(antlr.Token)

	// Set_COUNT sets the _COUNT token.
	Set_COUNT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetMethodinstrtrn returns the methodinstrtrn attribute.
	GetMethodinstrtrn() interfaces.Expression

	// SetMethodinstrtrn sets the methodinstrtrn attribute.
	SetMethodinstrtrn(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	EMPTY() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	CORCHIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORCHDER() antlr.TerminalNode

	// IsMethodvecrtrnContext differentiates from other interfaces.
	IsMethodvecrtrnContext()
}

type MethodvecrtrnContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	methodinstrtrn interfaces.Expression
	_ID            antlr.Token
	_EMPTY         antlr.Token
	_COUNT         antlr.Token
	_expr          IExprContext
}

func NewEmptyMethodvecrtrnContext() *MethodvecrtrnContext {
	var p = new(MethodvecrtrnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_methodvecrtrn
	return p
}

func InitEmptyMethodvecrtrnContext(p *MethodvecrtrnContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_methodvecrtrn
}

func (*MethodvecrtrnContext) IsMethodvecrtrnContext() {}

func NewMethodvecrtrnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodvecrtrnContext {
	var p = new(MethodvecrtrnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_methodvecrtrn

	return p
}

func (s *MethodvecrtrnContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodvecrtrnContext) Get_ID() antlr.Token { return s._ID }

func (s *MethodvecrtrnContext) Get_EMPTY() antlr.Token { return s._EMPTY }

func (s *MethodvecrtrnContext) Get_COUNT() antlr.Token { return s._COUNT }

func (s *MethodvecrtrnContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *MethodvecrtrnContext) Set_EMPTY(v antlr.Token) { s._EMPTY = v }

func (s *MethodvecrtrnContext) Set_COUNT(v antlr.Token) { s._COUNT = v }

func (s *MethodvecrtrnContext) Get_expr() IExprContext { return s._expr }

func (s *MethodvecrtrnContext) Set_expr(v IExprContext) { s._expr = v }

func (s *MethodvecrtrnContext) GetMethodinstrtrn() interfaces.Expression { return s.methodinstrtrn }

func (s *MethodvecrtrnContext) SetMethodinstrtrn(v interfaces.Expression) { s.methodinstrtrn = v }

func (s *MethodvecrtrnContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *MethodvecrtrnContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *MethodvecrtrnContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEMPTY, 0)
}

func (s *MethodvecrtrnContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *MethodvecrtrnContext) CORCHIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, 0)
}

func (s *MethodvecrtrnContext) Expr() IExprContext {
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

func (s *MethodvecrtrnContext) CORCHDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, 0)
}

func (s *MethodvecrtrnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodvecrtrnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodvecrtrnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterMethodvecrtrn(s)
	}
}

func (s *MethodvecrtrnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitMethodvecrtrn(s)
	}
}

func (p *SwiftGrammarParser) Methodvecrtrn() (localctx IMethodvecrtrnContext) {
	localctx = NewMethodvecrtrnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_methodvecrtrn)
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecrtrnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(460)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)

			var _m = p.Match(SwiftGrammarParserEMPTY)

			localctx.(*MethodvecrtrnContext)._EMPTY = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MethodvecrtrnContext).methodinstrtrn = expressions.NewVector((func() int {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetText()
			}
		}()), nil, (func() string {
			if localctx.(*MethodvecrtrnContext).Get_EMPTY() == nil {
				return ""
			} else {
				return localctx.(*MethodvecrtrnContext).Get_EMPTY().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(463)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecrtrnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(465)

			var _m = p.Match(SwiftGrammarParserCOUNT)

			localctx.(*MethodvecrtrnContext)._COUNT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MethodvecrtrnContext).methodinstrtrn = expressions.NewVector((func() int {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetText()
			}
		}()), nil, (func() string {
			if localctx.(*MethodvecrtrnContext).Get_COUNT() == nil {
				return ""
			} else {
				return localctx.(*MethodvecrtrnContext).Get_COUNT().GetText()
			}
		}()))

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(467)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecrtrnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)

			var _x = p.expr(0)

			localctx.(*MethodvecrtrnContext)._expr = _x
		}
		{
			p.SetState(470)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MethodvecrtrnContext).methodinstrtrn = expressions.NewVector((func() int {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*MethodvecrtrnContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MethodvecrtrnContext).Get_ID().GetText()
			}
		}()), localctx.(*MethodvecrtrnContext).Get_expr().GetE(), "access")

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

// IVecaccessContext is an interface to support dynamic dispatch.
type IVecaccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirstId returns the firstId token.
	GetFirstId() antlr.Token

	// GetSecondId returns the secondId token.
	GetSecondId() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// SetFirstId sets the firstId token.
	SetFirstId(antlr.Token)

	// SetSecondId sets the secondId token.
	SetSecondId(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetFirst returns the first rule contexts.
	GetFirst() IExprContext

	// GetSecond returns the second rule contexts.
	GetSecond() IExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IExprContext)

	// SetSecond sets the second rule contexts.
	SetSecond(IExprContext)

	// GetVecacc returns the vecacc attribute.
	GetVecacc() interfaces.Instruction

	// SetVecacc sets the vecacc attribute.
	SetVecacc(interfaces.Instruction)

	// Getter signatures
	AllCORCHIZQ() []antlr.TerminalNode
	CORCHIZQ(i int) antlr.TerminalNode
	AllCORCHDER() []antlr.TerminalNode
	CORCHDER(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsVecaccessContext differentiates from other interfaces.
	IsVecaccessContext()
}

type VecaccessContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	vecacc   interfaces.Instruction
	firstId  antlr.Token
	first    IExprContext
	secondId antlr.Token
	second   IExprContext
	_ID      antlr.Token
}

func NewEmptyVecaccessContext() *VecaccessContext {
	var p = new(VecaccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vecaccess
	return p
}

func InitEmptyVecaccessContext(p *VecaccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vecaccess
}

func (*VecaccessContext) IsVecaccessContext() {}

func NewVecaccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VecaccessContext {
	var p = new(VecaccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vecaccess

	return p
}

func (s *VecaccessContext) GetParser() antlr.Parser { return s.parser }

func (s *VecaccessContext) GetFirstId() antlr.Token { return s.firstId }

func (s *VecaccessContext) GetSecondId() antlr.Token { return s.secondId }

func (s *VecaccessContext) Get_ID() antlr.Token { return s._ID }

func (s *VecaccessContext) SetFirstId(v antlr.Token) { s.firstId = v }

func (s *VecaccessContext) SetSecondId(v antlr.Token) { s.secondId = v }

func (s *VecaccessContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VecaccessContext) GetFirst() IExprContext { return s.first }

func (s *VecaccessContext) GetSecond() IExprContext { return s.second }

func (s *VecaccessContext) SetFirst(v IExprContext) { s.first = v }

func (s *VecaccessContext) SetSecond(v IExprContext) { s.second = v }

func (s *VecaccessContext) GetVecacc() interfaces.Instruction { return s.vecacc }

func (s *VecaccessContext) SetVecacc(v interfaces.Instruction) { s.vecacc = v }

func (s *VecaccessContext) AllCORCHIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHIZQ)
}

func (s *VecaccessContext) CORCHIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHIZQ, i)
}

func (s *VecaccessContext) AllCORCHDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCORCHDER)
}

func (s *VecaccessContext) CORCHDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORCHDER, i)
}

func (s *VecaccessContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VecaccessContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *VecaccessContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *VecaccessContext) AllExpr() []IExprContext {
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

func (s *VecaccessContext) Expr(i int) IExprContext {
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

func (s *VecaccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecaccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VecaccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVecaccess(s)
	}
}

func (s *VecaccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVecaccess(s)
	}
}

func (p *SwiftGrammarParser) Vecaccess() (localctx IVecaccessContext) {
	localctx = NewVecaccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_vecaccess)
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecaccessContext).firstId = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)

			var _x = p.expr(0)

			localctx.(*VecaccessContext).first = _x
		}
		{
			p.SetState(478)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecaccessContext).secondId = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(481)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)

			var _x = p.expr(0)

			localctx.(*VecaccessContext).second = _x
		}
		{
			p.SetState(483)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*VecaccessContext).vecacc = instructions.NewVectorAsgmt((func() int {
			if localctx.(*VecaccessContext).GetFirstId() == nil {
				return 0
			} else {
				return localctx.(*VecaccessContext).GetFirstId().GetLine()
			}
		}()), (func() int {
			if localctx.(*VecaccessContext).GetFirstId() == nil {
				return 0
			} else {
				return localctx.(*VecaccessContext).GetFirstId().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VecaccessContext).GetFirstId() == nil {
				return ""
			} else {
				return localctx.(*VecaccessContext).GetFirstId().GetText()
			}
		}()), localctx.(*VecaccessContext).GetFirst().GetE(), localctx.(*VecaccessContext).GetSecond().GetE(), (func() string {
			if localctx.(*VecaccessContext).GetSecondId() == nil {
				return ""
			} else {
				return localctx.(*VecaccessContext).GetSecondId().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(486)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecaccessContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(487)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(488)

			var _x = p.expr(0)

			localctx.(*VecaccessContext).first = _x
		}
		{
			p.SetState(489)
			p.Match(SwiftGrammarParserCORCHDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)

			var _x = p.expr(0)

			localctx.(*VecaccessContext).second = _x
		}
		localctx.(*VecaccessContext).vecacc = instructions.NewVectorAsgmt((func() int {
			if localctx.(*VecaccessContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VecaccessContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*VecaccessContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*VecaccessContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VecaccessContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VecaccessContext).Get_ID().GetText()
			}
		}()), localctx.(*VecaccessContext).GetFirst().GetE(), localctx.(*VecaccessContext).GetSecond().GetE(), "")

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

// IFuncstmtContext is an interface to support dynamic dispatch.
type IFuncstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listparam returns the _listparam rule contexts.
	Get_listparam() IListparamContext

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_listparam sets the _listparam rule contexts.
	Set_listparam(IListparamContext)

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetLista returns the lista rule context list.
	GetLista() []IListparamContext

	// SetLista sets the lista rule context list.
	SetLista([]IListparamContext)

	// GetFuncinstr returns the funcinstr attribute.
	GetFuncinstr() interfaces.Instruction

	// SetFuncinstr sets the funcinstr attribute.
	SetFuncinstr(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	ARROW() antlr.TerminalNode
	Tipo() ITipoContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode
	AllListparam() []IListparamContext
	Listparam(i int) IListparamContext

	// IsFuncstmtContext differentiates from other interfaces.
	IsFuncstmtContext()
}

type FuncstmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	funcinstr  interfaces.Instruction
	_FUNC      antlr.Token
	_ID        antlr.Token
	_listparam IListparamContext
	lista      []IListparamContext
	_tipo      ITipoContext
	_block     IBlockContext
}

func NewEmptyFuncstmtContext() *FuncstmtContext {
	var p = new(FuncstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcstmt
	return p
}

func InitEmptyFuncstmtContext(p *FuncstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcstmt
}

func (*FuncstmtContext) IsFuncstmtContext() {}

func NewFuncstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncstmtContext {
	var p = new(FuncstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funcstmt

	return p
}

func (s *FuncstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncstmtContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *FuncstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncstmtContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *FuncstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncstmtContext) Get_listparam() IListparamContext { return s._listparam }

func (s *FuncstmtContext) Get_tipo() ITipoContext { return s._tipo }

func (s *FuncstmtContext) Get_block() IBlockContext { return s._block }

func (s *FuncstmtContext) Set_listparam(v IListparamContext) { s._listparam = v }

func (s *FuncstmtContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *FuncstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *FuncstmtContext) GetLista() []IListparamContext { return s.lista }

func (s *FuncstmtContext) SetLista(v []IListparamContext) { s.lista = v }

func (s *FuncstmtContext) GetFuncinstr() interfaces.Instruction { return s.funcinstr }

func (s *FuncstmtContext) SetFuncinstr(v interfaces.Instruction) { s.funcinstr = v }

func (s *FuncstmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FuncstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FuncstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncstmtContext) ARROW() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserARROW, 0)
}

func (s *FuncstmtContext) Tipo() ITipoContext {
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

func (s *FuncstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *FuncstmtContext) Block() IBlockContext {
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

func (s *FuncstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *FuncstmtContext) AllListparam() []IListparamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListparamContext); ok {
			len++
		}
	}

	tst := make([]IListparamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListparamContext); ok {
			tst[i] = t.(IListparamContext)
			i++
		}
	}

	return tst
}

func (s *FuncstmtContext) Listparam(i int) IListparamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListparamContext); ok {
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

	return t.(IListparamContext)
}

func (s *FuncstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncstmt(s)
	}
}

func (s *FuncstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncstmt(s)
	}
}

func (p *SwiftGrammarParser) Funcstmt() (localctx IFuncstmtContext) {
	localctx = NewFuncstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_funcstmt)

	var args []interface{}

	var _la int

	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SwiftGrammarParserCAME || _la == SwiftGrammarParserID {
			{
				p.SetState(499)

				var _x = p.Listparam()

				localctx.(*FuncstmtContext)._listparam = _x
			}
			localctx.(*FuncstmtContext).lista = append(localctx.(*FuncstmtContext).lista, localctx.(*FuncstmtContext)._listparam)

			p.SetState(502)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(504)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Match(SwiftGrammarParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(506)

			var _x = p.Tipo()

			localctx.(*FuncstmtContext)._tipo = _x
		}
		{
			p.SetState(507)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(508)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(509)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		for _, e := range localctx.(*FuncstmtContext).GetLista() {
			// fmt.Println(fmt.Sprintf("%T", e.GetListparaminstr()))
			args = append(args, e.GetListparaminstr())
		}
		localctx.(*FuncstmtContext).funcinstr = instructions.NewFunction((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), args, localctx.(*FuncstmtContext).Get_tipo().GetRtipo(), localctx.(*FuncstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(512)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(514)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SwiftGrammarParserCAME || _la == SwiftGrammarParserID {
			{
				p.SetState(515)

				var _x = p.Listparam()

				localctx.(*FuncstmtContext)._listparam = _x
			}
			localctx.(*FuncstmtContext).lista = append(localctx.(*FuncstmtContext).lista, localctx.(*FuncstmtContext)._listparam)

			p.SetState(518)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(520)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(523)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		for _, e := range localctx.(*FuncstmtContext).GetLista() {
			// fmt.Println(fmt.Sprintf("%T", e.GetListparaminstr()))
			args = append(args, e.GetListparaminstr())
		}
		localctx.(*FuncstmtContext).funcinstr = instructions.NewFunction((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), args, 4, localctx.(*FuncstmtContext).Get_block().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(526)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(527)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(529)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Match(SwiftGrammarParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(531)

			var _x = p.Tipo()

			localctx.(*FuncstmtContext)._tipo = _x
		}
		{
			p.SetState(532)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(534)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*FuncstmtContext).funcinstr = instructions.NewFunction((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*FuncstmtContext).Get_tipo().GetRtipo(), localctx.(*FuncstmtContext).Get_block().GetBlk())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(537)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FuncstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FuncstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(541)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)

			var _x = p.Block()

			localctx.(*FuncstmtContext)._block = _x
		}
		{
			p.SetState(543)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*FuncstmtContext).funcinstr = instructions.NewFunction((func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FuncstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FuncstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncstmtContext).Get_ID().GetText()
			}
		}()), nil, 4, localctx.(*FuncstmtContext).Get_block().GetBlk())

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

// IListparamContext is an interface to support dynamic dispatch.
type IListparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExtr returns the extr token.
	GetExtr() antlr.Token

	// GetInter returns the inter token.
	GetInter() antlr.Token

	// Get_INOUT returns the _INOUT token.
	Get_INOUT() antlr.Token

	// SetExtr sets the extr token.
	SetExtr(antlr.Token)

	// SetInter sets the inter token.
	SetInter(antlr.Token)

	// Set_INOUT sets the _INOUT token.
	Set_INOUT(antlr.Token)

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// GetListparaminstr returns the listparaminstr attribute.
	GetListparaminstr() interfaces.Instruction

	// SetListparaminstr sets the listparaminstr attribute.
	SetListparaminstr(interfaces.Instruction)

	// Getter signatures
	COLON() antlr.TerminalNode
	Tipo() ITipoContext
	COMA() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	INOUT() antlr.TerminalNode
	CAME() antlr.TerminalNode

	// IsListparamContext differentiates from other interfaces.
	IsListparamContext()
}

type ListparamContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	listparaminstr interfaces.Instruction
	extr           antlr.Token
	inter          antlr.Token
	_INOUT         antlr.Token
	_tipo          ITipoContext
}

func NewEmptyListparamContext() *ListparamContext {
	var p = new(ListparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listparam
	return p
}

func InitEmptyListparamContext(p *ListparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listparam
}

func (*ListparamContext) IsListparamContext() {}

func NewListparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListparamContext {
	var p = new(ListparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listparam

	return p
}

func (s *ListparamContext) GetParser() antlr.Parser { return s.parser }

func (s *ListparamContext) GetExtr() antlr.Token { return s.extr }

func (s *ListparamContext) GetInter() antlr.Token { return s.inter }

func (s *ListparamContext) Get_INOUT() antlr.Token { return s._INOUT }

func (s *ListparamContext) SetExtr(v antlr.Token) { s.extr = v }

func (s *ListparamContext) SetInter(v antlr.Token) { s.inter = v }

func (s *ListparamContext) Set_INOUT(v antlr.Token) { s._INOUT = v }

func (s *ListparamContext) Get_tipo() ITipoContext { return s._tipo }

func (s *ListparamContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *ListparamContext) GetListparaminstr() interfaces.Instruction { return s.listparaminstr }

func (s *ListparamContext) SetListparaminstr(v interfaces.Instruction) { s.listparaminstr = v }

func (s *ListparamContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *ListparamContext) Tipo() ITipoContext {
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

func (s *ListparamContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListparamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ListparamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ListparamContext) INOUT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINOUT, 0)
}

func (s *ListparamContext) CAME() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCAME, 0)
}

func (s *ListparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListparam(s)
	}
}

func (s *ListparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListparam(s)
	}
}

func (p *SwiftGrammarParser) Listparam() (localctx IListparamContext) {
	localctx = NewListparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftGrammarParserRULE_listparam)
	var _la int

	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(549)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(548)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ListparamContext).extr = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SwiftGrammarParserCAME || _la == SwiftGrammarParserID) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ListparamContext).extr = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(551)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListparamContext).inter = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserINOUT {
			{
				p.SetState(553)

				var _m = p.Match(SwiftGrammarParserINOUT)

				localctx.(*ListparamContext)._INOUT = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(556)

			var _x = p.Tipo()

			localctx.(*ListparamContext)._tipo = _x
		}
		{
			p.SetState(557)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		var flag bool
		var aux string
		if localctx.(*ListparamContext).Get_INOUT() != nil {
			flag = true
		} else {
			flag = false
		}
		if localctx.(*ListparamContext).GetExtr() != nil {
			aux = (func() string {
				if localctx.(*ListparamContext).GetExtr() == nil {
					return ""
				} else {
					return localctx.(*ListparamContext).GetExtr().GetText()
				}
			}())
		}
		localctx.(*ListparamContext).listparaminstr = instructions.NewParam((func() int {
			if localctx.(*ListparamContext).GetExtr() == nil {
				return 0
			} else {
				return localctx.(*ListparamContext).GetExtr().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListparamContext).GetExtr() == nil {
				return 0
			} else {
				return localctx.(*ListparamContext).GetExtr().GetColumn()
			}
		}()), aux, (func() string {
			if localctx.(*ListparamContext).GetInter() == nil {
				return ""
			} else {
				return localctx.(*ListparamContext).GetInter().GetText()
			}
		}()), flag, localctx.(*ListparamContext).Get_tipo().GetRtipo())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(561)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(560)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ListparamContext).extr = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == SwiftGrammarParserCAME || _la == SwiftGrammarParserID) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ListparamContext).extr = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(563)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListparamContext).inter = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserINOUT {
			{
				p.SetState(565)

				var _m = p.Match(SwiftGrammarParserINOUT)

				localctx.(*ListparamContext)._INOUT = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(568)

			var _x = p.Tipo()

			localctx.(*ListparamContext)._tipo = _x
		}

		var flag bool
		var aux string
		if localctx.(*ListparamContext).Get_INOUT() != nil {
			flag = true
		} else {
			flag = false
		}
		if localctx.(*ListparamContext).GetExtr() != nil {
			aux = (func() string {
				if localctx.(*ListparamContext).GetExtr() == nil {
					return ""
				} else {
					return localctx.(*ListparamContext).GetExtr().GetText()
				}
			}())
		}
		localctx.(*ListparamContext).listparaminstr = instructions.NewParam((func() int {
			if localctx.(*ListparamContext).GetExtr() == nil {
				return 0
			} else {
				return localctx.(*ListparamContext).GetExtr().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListparamContext).GetExtr() == nil {
				return 0
			} else {
				return localctx.(*ListparamContext).GetExtr().GetColumn()
			}
		}()), aux, (func() string {
			if localctx.(*ListparamContext).GetInter() == nil {
				return ""
			} else {
				return localctx.(*ListparamContext).GetInter().GetText()
			}
		}()), flag, localctx.(*ListparamContext).Get_tipo().GetRtipo())

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

// ICallfuncContext is an interface to support dynamic dispatch.
type ICallfuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listparamcall returns the _listparamcall rule contexts.
	Get_listparamcall() IListparamcallContext

	// Set_listparamcall sets the _listparamcall rule contexts.
	Set_listparamcall(IListparamcallContext)

	// GetLista returns the lista rule context list.
	GetLista() []IListparamcallContext

	// SetLista sets the lista rule context list.
	SetLista([]IListparamcallContext)

	// GetCallfuncexpr returns the callfuncexpr attribute.
	GetCallfuncexpr() interfaces.Expression

	// SetCallfuncexpr sets the callfuncexpr attribute.
	SetCallfuncexpr(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	AllListparamcall() []IListparamcallContext
	Listparamcall(i int) IListparamcallContext

	// IsCallfuncContext differentiates from other interfaces.
	IsCallfuncContext()
}

type CallfuncContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	callfuncexpr   interfaces.Expression
	_ID            antlr.Token
	_listparamcall IListparamcallContext
	lista          []IListparamcallContext
}

func NewEmptyCallfuncContext() *CallfuncContext {
	var p = new(CallfuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callfunc
	return p
}

func InitEmptyCallfuncContext(p *CallfuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callfunc
}

func (*CallfuncContext) IsCallfuncContext() {}

func NewCallfuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallfuncContext {
	var p = new(CallfuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callfunc

	return p
}

func (s *CallfuncContext) GetParser() antlr.Parser { return s.parser }

func (s *CallfuncContext) Get_ID() antlr.Token { return s._ID }

func (s *CallfuncContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallfuncContext) Get_listparamcall() IListparamcallContext { return s._listparamcall }

func (s *CallfuncContext) Set_listparamcall(v IListparamcallContext) { s._listparamcall = v }

func (s *CallfuncContext) GetLista() []IListparamcallContext { return s.lista }

func (s *CallfuncContext) SetLista(v []IListparamcallContext) { s.lista = v }

func (s *CallfuncContext) GetCallfuncexpr() interfaces.Expression { return s.callfuncexpr }

func (s *CallfuncContext) SetCallfuncexpr(v interfaces.Expression) { s.callfuncexpr = v }

func (s *CallfuncContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallfuncContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *CallfuncContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *CallfuncContext) AllListparamcall() []IListparamcallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListparamcallContext); ok {
			len++
		}
	}

	tst := make([]IListparamcallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListparamcallContext); ok {
			tst[i] = t.(IListparamcallContext)
			i++
		}
	}

	return tst
}

func (s *CallfuncContext) Listparamcall(i int) IListparamcallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListparamcallContext); ok {
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

	return t.(IListparamcallContext)
}

func (s *CallfuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallfuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallfuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallfunc(s)
	}
}

func (s *CallfuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallfunc(s)
	}
}

func (p *SwiftGrammarParser) Callfunc() (localctx ICallfuncContext) {
	localctx = NewCallfuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftGrammarParserRULE_callfunc)

	var args []interface{}

	var _la int

	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(573)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*CallfuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(576)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54048624367108566) != 0) || _la == SwiftGrammarParserAMP {
			{
				p.SetState(575)

				var _x = p.Listparamcall()

				localctx.(*CallfuncContext)._listparamcall = _x
			}
			localctx.(*CallfuncContext).lista = append(localctx.(*CallfuncContext).lista, localctx.(*CallfuncContext)._listparamcall)

			p.SetState(578)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(580)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		for _, e := range localctx.(*CallfuncContext).GetLista() {
			// fmt.Println(fmt.Sprintf("%T", e.GetListparamcallinstr()))
			args = append(args, e.GetListparamcallinstr())
		}
		localctx.(*CallfuncContext).callfuncexpr = expressions.NewCallFunc((func() int {
			if localctx.(*CallfuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*CallfuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*CallfuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*CallfuncContext).Get_ID().GetText()
			}
		}()), args)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(583)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*CallfuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(584)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(585)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*CallfuncContext).callfuncexpr = expressions.NewCallFunc((func() int {
			if localctx.(*CallfuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*CallfuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*CallfuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*CallfuncContext).Get_ID().GetText()
			}
		}()), nil)

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

// ICallfuncinstructionContext is an interface to support dynamic dispatch.
type ICallfuncinstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listparamcall returns the _listparamcall rule contexts.
	Get_listparamcall() IListparamcallContext

	// Set_listparamcall sets the _listparamcall rule contexts.
	Set_listparamcall(IListparamcallContext)

	// GetLista returns the lista rule context list.
	GetLista() []IListparamcallContext

	// SetLista sets the lista rule context list.
	SetLista([]IListparamcallContext)

	// GetCallfuncinstr returns the callfuncinstr attribute.
	GetCallfuncinstr() interfaces.Instruction

	// SetCallfuncinstr sets the callfuncinstr attribute.
	SetCallfuncinstr(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	AllListparamcall() []IListparamcallContext
	Listparamcall(i int) IListparamcallContext

	// IsCallfuncinstructionContext differentiates from other interfaces.
	IsCallfuncinstructionContext()
}

type CallfuncinstructionContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	callfuncinstr  interfaces.Instruction
	_ID            antlr.Token
	_listparamcall IListparamcallContext
	lista          []IListparamcallContext
}

func NewEmptyCallfuncinstructionContext() *CallfuncinstructionContext {
	var p = new(CallfuncinstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callfuncinstruction
	return p
}

func InitEmptyCallfuncinstructionContext(p *CallfuncinstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callfuncinstruction
}

func (*CallfuncinstructionContext) IsCallfuncinstructionContext() {}

func NewCallfuncinstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallfuncinstructionContext {
	var p = new(CallfuncinstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callfuncinstruction

	return p
}

func (s *CallfuncinstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallfuncinstructionContext) Get_ID() antlr.Token { return s._ID }

func (s *CallfuncinstructionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallfuncinstructionContext) Get_listparamcall() IListparamcallContext {
	return s._listparamcall
}

func (s *CallfuncinstructionContext) Set_listparamcall(v IListparamcallContext) { s._listparamcall = v }

func (s *CallfuncinstructionContext) GetLista() []IListparamcallContext { return s.lista }

func (s *CallfuncinstructionContext) SetLista(v []IListparamcallContext) { s.lista = v }

func (s *CallfuncinstructionContext) GetCallfuncinstr() interfaces.Instruction {
	return s.callfuncinstr
}

func (s *CallfuncinstructionContext) SetCallfuncinstr(v interfaces.Instruction) { s.callfuncinstr = v }

func (s *CallfuncinstructionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallfuncinstructionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *CallfuncinstructionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *CallfuncinstructionContext) AllListparamcall() []IListparamcallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListparamcallContext); ok {
			len++
		}
	}

	tst := make([]IListparamcallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListparamcallContext); ok {
			tst[i] = t.(IListparamcallContext)
			i++
		}
	}

	return tst
}

func (s *CallfuncinstructionContext) Listparamcall(i int) IListparamcallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListparamcallContext); ok {
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

	return t.(IListparamcallContext)
}

func (s *CallfuncinstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallfuncinstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallfuncinstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallfuncinstruction(s)
	}
}

func (s *CallfuncinstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallfuncinstruction(s)
	}
}

func (p *SwiftGrammarParser) Callfuncinstruction() (localctx ICallfuncinstructionContext) {
	localctx = NewCallfuncinstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftGrammarParserRULE_callfuncinstruction)

	var args []interface{}

	var _la int

	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(589)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*CallfuncinstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(590)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(592)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54048624367108566) != 0) || _la == SwiftGrammarParserAMP {
			{
				p.SetState(591)

				var _x = p.Listparamcall()

				localctx.(*CallfuncinstructionContext)._listparamcall = _x
			}
			localctx.(*CallfuncinstructionContext).lista = append(localctx.(*CallfuncinstructionContext).lista, localctx.(*CallfuncinstructionContext)._listparamcall)

			p.SetState(594)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(596)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		for _, e := range localctx.(*CallfuncinstructionContext).GetLista() {
			// fmt.Println(fmt.Sprintf("%T", e.GetListparamcallinstr()))
			args = append(args, e.GetListparamcallinstr())
		}
		localctx.(*CallfuncinstructionContext).callfuncinstr = instructions.NewCallFunc((func() int {
			if localctx.(*CallfuncinstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncinstructionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*CallfuncinstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncinstructionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*CallfuncinstructionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*CallfuncinstructionContext).Get_ID().GetText()
			}
		}()), args)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(599)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*CallfuncinstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(600)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*CallfuncinstructionContext).callfuncinstr = instructions.NewCallFunc((func() int {
			if localctx.(*CallfuncinstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncinstructionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*CallfuncinstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*CallfuncinstructionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*CallfuncinstructionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*CallfuncinstructionContext).Get_ID().GetText()
			}
		}()), nil)

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

// IListparamcallContext is an interface to support dynamic dispatch.
type IListparamcallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_AMP returns the _AMP token.
	Get_AMP() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_AMP sets the _AMP token.
	Set_AMP(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetListparamcallinstr returns the listparamcallinstr attribute.
	GetListparamcallinstr() interfaces.Expression

	// SetListparamcallinstr sets the listparamcallinstr attribute.
	SetListparamcallinstr(interfaces.Expression)

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ID() antlr.TerminalNode
	COLON() antlr.TerminalNode
	AMP() antlr.TerminalNode

	// IsListparamcallContext differentiates from other interfaces.
	IsListparamcallContext()
}

type ListparamcallContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	listparamcallinstr interfaces.Expression
	_ID                antlr.Token
	_AMP               antlr.Token
	_expr              IExprContext
}

func NewEmptyListparamcallContext() *ListparamcallContext {
	var p = new(ListparamcallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listparamcall
	return p
}

func InitEmptyListparamcallContext(p *ListparamcallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listparamcall
}

func (*ListparamcallContext) IsListparamcallContext() {}

func NewListparamcallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListparamcallContext {
	var p = new(ListparamcallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listparamcall

	return p
}

func (s *ListparamcallContext) GetParser() antlr.Parser { return s.parser }

func (s *ListparamcallContext) Get_ID() antlr.Token { return s._ID }

func (s *ListparamcallContext) Get_AMP() antlr.Token { return s._AMP }

func (s *ListparamcallContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListparamcallContext) Set_AMP(v antlr.Token) { s._AMP = v }

func (s *ListparamcallContext) Get_expr() IExprContext { return s._expr }

func (s *ListparamcallContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListparamcallContext) GetListparamcallinstr() interfaces.Expression {
	return s.listparamcallinstr
}

func (s *ListparamcallContext) SetListparamcallinstr(v interfaces.Expression) {
	s.listparamcallinstr = v
}

func (s *ListparamcallContext) Expr() IExprContext {
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

func (s *ListparamcallContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListparamcallContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListparamcallContext) COLON() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOLON, 0)
}

func (s *ListparamcallContext) AMP() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAMP, 0)
}

func (s *ListparamcallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListparamcallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListparamcallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListparamcall(s)
	}
}

func (s *ListparamcallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListparamcall(s)
	}
}

func (p *SwiftGrammarParser) Listparamcall() (localctx IListparamcallContext) {
	localctx = NewListparamcallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_listparamcall)
	var _la int

	p.SetState(626)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(607)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(605)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListparamcallContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(606)
				p.Match(SwiftGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(610)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserAMP {
			{
				p.SetState(609)

				var _m = p.Match(SwiftGrammarParserAMP)

				localctx.(*ListparamcallContext)._AMP = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(612)

			var _x = p.expr(0)

			localctx.(*ListparamcallContext)._expr = _x
		}
		{
			p.SetState(613)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		var flag bool
		if localctx.(*ListparamcallContext).Get_AMP() != nil {
			flag = true
		}
		localctx.(*ListparamcallContext).listparamcallinstr = expressions.NewCallParams((func() antlr.Token {
			if localctx.(*ListparamcallContext).Get_expr() == nil {
				return nil
			} else {
				return localctx.(*ListparamcallContext).Get_expr().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ListparamcallContext).Get_expr() == nil {
				return nil
			} else {
				return localctx.(*ListparamcallContext).Get_expr().GetStart()
			}
		}()).GetColumn(), (func() string {
			if localctx.(*ListparamcallContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListparamcallContext).Get_ID().GetText()
			}
		}()), flag, localctx.(*ListparamcallContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(618)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(616)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListparamcallContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(617)
				p.Match(SwiftGrammarParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(621)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserAMP {
			{
				p.SetState(620)

				var _m = p.Match(SwiftGrammarParserAMP)

				localctx.(*ListparamcallContext)._AMP = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(623)

			var _x = p.expr(0)

			localctx.(*ListparamcallContext)._expr = _x
		}

		var flag bool
		if localctx.(*ListparamcallContext).Get_AMP() != nil {
			flag = true
		}
		localctx.(*ListparamcallContext).listparamcallinstr = expressions.NewCallParams((func() antlr.Token {
			if localctx.(*ListparamcallContext).Get_expr() == nil {
				return nil
			} else {
				return localctx.(*ListparamcallContext).Get_expr().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ListparamcallContext).Get_expr() == nil {
				return nil
			} else {
				return localctx.(*ListparamcallContext).Get_expr().GetStart()
			}
		}()).GetColumn(), (func() string {
			if localctx.(*ListparamcallContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListparamcallContext).Get_ID().GetText()
			}
		}()), flag, localctx.(*ListparamcallContext).Get_expr().GetE())

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

// IFuncembedContext is an interface to support dynamic dispatch.
type IFuncembedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STR returns the _STR token.
	Get_STR() antlr.Token

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Set_STR sets the _STR token.
	Set_STR(antlr.Token)

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetFuncembedexpr returns the funcembedexpr attribute.
	GetFuncembedexpr() interfaces.Expression

	// SetFuncembedexpr sets the funcembedexpr attribute.
	SetFuncembedexpr(interfaces.Expression)

	// Getter signatures
	STR() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsFuncembedContext differentiates from other interfaces.
	IsFuncembedContext()
}

type FuncembedContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	funcembedexpr interfaces.Expression
	_STR          antlr.Token
	_expr         IExprContext
	_INT          antlr.Token
	_FLOAT        antlr.Token
}

func NewEmptyFuncembedContext() *FuncembedContext {
	var p = new(FuncembedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcembed
	return p
}

func InitEmptyFuncembedContext(p *FuncembedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_funcembed
}

func (*FuncembedContext) IsFuncembedContext() {}

func NewFuncembedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncembedContext {
	var p = new(FuncembedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_funcembed

	return p
}

func (s *FuncembedContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncembedContext) Get_STR() antlr.Token { return s._STR }

func (s *FuncembedContext) Get_INT() antlr.Token { return s._INT }

func (s *FuncembedContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *FuncembedContext) Set_STR(v antlr.Token) { s._STR = v }

func (s *FuncembedContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *FuncembedContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *FuncembedContext) Get_expr() IExprContext { return s._expr }

func (s *FuncembedContext) Set_expr(v IExprContext) { s._expr = v }

func (s *FuncembedContext) GetFuncembedexpr() interfaces.Expression { return s.funcembedexpr }

func (s *FuncembedContext) SetFuncembedexpr(v interfaces.Expression) { s.funcembedexpr = v }

func (s *FuncembedContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *FuncembedContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FuncembedContext) Expr() IExprContext {
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

func (s *FuncembedContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FuncembedContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *FuncembedContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *FuncembedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncembedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncembedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFuncembed(s)
	}
}

func (s *FuncembedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFuncembed(s)
	}
}

func (p *SwiftGrammarParser) Funcembed() (localctx IFuncembedContext) {
	localctx = NewFuncembedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftGrammarParserRULE_funcembed)
	p.SetState(646)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(628)

			var _m = p.Match(SwiftGrammarParserSTR)

			localctx.(*FuncembedContext)._STR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(629)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)

			var _x = p.expr(0)

			localctx.(*FuncembedContext)._expr = _x
		}
		{
			p.SetState(631)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncembedContext).funcembedexpr = expressions.NewString((func() int {
			if localctx.(*FuncembedContext).Get_STR() == nil {
				return 0
			} else {
				return localctx.(*FuncembedContext).Get_STR().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncembedContext).Get_STR() == nil {
				return 0
			} else {
				return localctx.(*FuncembedContext).Get_STR().GetColumn()
			}
		}()), localctx.(*FuncembedContext).Get_expr().GetE())

	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)

			var _m = p.Match(SwiftGrammarParserINT)

			localctx.(*FuncembedContext)._INT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(635)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(636)

			var _x = p.expr(0)

			localctx.(*FuncembedContext)._expr = _x
		}
		{
			p.SetState(637)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncembedContext).funcembedexpr = expressions.NewInteger((func() int {
			if localctx.(*FuncembedContext).Get_INT() == nil {
				return 0
			} else {
				return localctx.(*FuncembedContext).Get_INT().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncembedContext).Get_INT() == nil {
				return 0
			} else {
				return localctx.(*FuncembedContext).Get_INT().GetColumn()
			}
		}()), localctx.(*FuncembedContext).Get_expr().GetE())

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(640)

			var _m = p.Match(SwiftGrammarParserFLOAT)

			localctx.(*FuncembedContext)._FLOAT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(641)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(642)

			var _x = p.expr(0)

			localctx.(*FuncembedContext)._expr = _x
		}
		{
			p.SetState(643)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FuncembedContext).funcembedexpr = expressions.NewFloat((func() int {
			if localctx.(*FuncembedContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*FuncembedContext).Get_FLOAT().GetLine()
			}
		}()), (func() int {
			if localctx.(*FuncembedContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*FuncembedContext).Get_FLOAT().GetColumn()
			}
		}()), localctx.(*FuncembedContext).Get_expr().GetE())

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

	// Get_NIL returns the _NIL token.
	Get_NIL() antlr.Token

	// Get_AND returns the _AND token.
	Get_AND() antlr.Token

	// Get_OR returns the _OR token.
	Get_OR() antlr.Token

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

	// Set_NIL sets the _NIL token.
	Set_NIL(antlr.Token)

	// Set_AND sets the _AND token.
	Set_AND(antlr.Token)

	// Set_OR sets the _OR token.
	Set_OR(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_callfunc returns the _callfunc rule contexts.
	Get_callfunc() ICallfuncContext

	// Get_funcembed returns the _funcembed rule contexts.
	Get_funcembed() IFuncembedContext

	// Get_methodvecrtrn returns the _methodvecrtrn rule contexts.
	Get_methodvecrtrn() IMethodvecrtrnContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_callfunc sets the _callfunc rule contexts.
	Set_callfunc(ICallfuncContext)

	// Set_funcembed sets the _funcembed rule contexts.
	Set_funcembed(IFuncembedContext)

	// Set_methodvecrtrn sets the _methodvecrtrn rule contexts.
	Set_methodvecrtrn(IMethodvecrtrnContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	SUB() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NOT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	Callfunc() ICallfuncContext
	Funcembed() IFuncembedContext
	FAL() antlr.TerminalNode
	ID() antlr.TerminalNode
	Methodvecrtrn() IMethodvecrtrnContext
	NIL() antlr.TerminalNode
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
	parser         antlr.Parser
	e              interfaces.Expression
	left           IExprContext
	op             antlr.Token
	right          IExprContext
	_expr          IExprContext
	_NUMBER        antlr.Token
	_CHAR          antlr.Token
	_STRING        antlr.Token
	_TRU           antlr.Token
	_callfunc      ICallfuncContext
	_funcembed     IFuncembedContext
	_FAL           antlr.Token
	_ID            antlr.Token
	_methodvecrtrn IMethodvecrtrnContext
	_NIL           antlr.Token
	_AND           antlr.Token
	_OR            antlr.Token
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

func (s *ExprContext) Get_NIL() antlr.Token { return s._NIL }

func (s *ExprContext) Get_AND() antlr.Token { return s._AND }

func (s *ExprContext) Get_OR() antlr.Token { return s._OR }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_NIL(v antlr.Token) { s._NIL = v }

func (s *ExprContext) Set_AND(v antlr.Token) { s._AND = v }

func (s *ExprContext) Set_OR(v antlr.Token) { s._OR = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_callfunc() ICallfuncContext { return s._callfunc }

func (s *ExprContext) Get_funcembed() IFuncembedContext { return s._funcembed }

func (s *ExprContext) Get_methodvecrtrn() IMethodvecrtrnContext { return s._methodvecrtrn }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_callfunc(v ICallfuncContext) { s._callfunc = v }

func (s *ExprContext) Set_funcembed(v IFuncembedContext) { s._funcembed = v }

func (s *ExprContext) Set_methodvecrtrn(v IMethodvecrtrnContext) { s._methodvecrtrn = v }

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

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
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

func (s *ExprContext) Callfunc() ICallfuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallfuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallfuncContext)
}

func (s *ExprContext) Funcembed() IFuncembedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncembedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncembedContext)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) Methodvecrtrn() IMethodvecrtrnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethodvecrtrnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethodvecrtrnContext)
}

func (s *ExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNIL, 0)
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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(685)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(649)

			var _m = p.Match(SwiftGrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(650)

			var _x = p.expr(20)

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

	case 2:
		{
			p.SetState(653)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(654)

			var _x = p.expr(19)

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
		}()).GetColumn(), localctx.(*ExprContext).GetRight().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()), expressions.NewPrimitive((func() int {
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
		}()), false, environment.BOOLEAN))

	case 3:
		{
			p.SetState(657)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(658)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(659)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 4:
		{
			p.SetState(662)

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

	case 5:
		{
			p.SetState(664)

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

	case 6:
		{
			p.SetState(666)

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

	case 7:
		{
			p.SetState(668)

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

	case 8:
		{
			p.SetState(670)

			var _x = p.Callfunc()

			localctx.(*ExprContext)._callfunc = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_callfunc().GetCallfuncexpr()

	case 9:
		{
			p.SetState(673)

			var _x = p.Funcembed()

			localctx.(*ExprContext)._funcembed = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_funcembed().GetFuncembedexpr()

	case 10:
		{
			p.SetState(676)

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

	case 11:
		{
			p.SetState(678)

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

	case 12:
		{
			p.SetState(680)

			var _x = p.Methodvecrtrn()

			localctx.(*ExprContext)._methodvecrtrn = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_methodvecrtrn().GetMethodinstrtrn()

	case 13:
		{
			p.SetState(683)

			var _m = p.Match(SwiftGrammarParserNIL)

			localctx.(*ExprContext)._NIL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetColumn()
			}
		}()), "nil", environment.NULL)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(724)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(722)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(687)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(688)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(689)

					var _x = p.expr(19)

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
				p.SetState(692)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(693)

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
					p.SetState(694)

					var _x = p.expr(18)

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
				p.SetState(697)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(698)

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
					p.SetState(699)

					var _x = p.expr(17)

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
				p.SetState(702)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(703)

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
					p.SetState(704)

					var _x = p.expr(16)

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
				p.SetState(707)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(708)

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
					p.SetState(709)

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

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(712)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(713)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext)._AND = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(714)

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
					if localctx.(*ExprContext).Get_AND() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).Get_AND().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(717)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(718)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext)._OR = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(719)

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
					if localctx.(*ExprContext).Get_OR() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).Get_OR().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(726)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
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
	case 31:
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
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 12)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
