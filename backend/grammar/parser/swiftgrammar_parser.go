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
		"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'",
		"'break'", "'continue'", "'return'", "'struct'", "'self'", "'mutating'",
		"'func'", "'append'", "'removeLast'", "'remove'", "'isEmpty'", "'count'",
		"", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='",
		"'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'?'", "'->'", "','", "'.'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL",
		"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE",
		"FOR", "IN", "RANGEPTS", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT",
		"SELF", "MUTATING", "FUNC", "APPEND", "REMOVELAST", "REMOVE", "EMPTY",
		"COUNT", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR",
		"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", "MUL", "DIV",
		"ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ",
		"CORCHDER", "Q_MARK", "ARROW", "COMA", "PUNTO", "COLON", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "ifstmt", "elseifstmt", "elsestmt",
		"varstmt", "tipo", "varasgmt", "conststmt", "switchstmt", "casestmt",
		"defaultstmt", "whilestmt", "forstmt", "guardstmt", "transferstmt",
		"vectorstmt", "definestmt", "listexpr", "methodvec", "methodvecrtrn",
		"vecaccess", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 66, 452, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 4, 1, 56, 8, 1, 11, 1, 12, 1, 57, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 101, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 122,
		8, 4, 11, 4, 12, 4, 123, 1, 4, 1, 4, 3, 4, 128, 8, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 141, 8, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 177, 8, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 189,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 208, 8, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 224, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 230, 8, 11, 11,
		11, 12, 11, 231, 1, 11, 3, 11, 235, 8, 11, 1, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 245, 8, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 3, 13, 253, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 284, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 304, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 4, 19, 318, 8, 19,
		11, 19, 12, 19, 319, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 332, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 356, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 372, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 410, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 447, 8, 24, 10, 24, 12,
		24, 450, 9, 24, 1, 24, 0, 1, 48, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 5, 1, 0,
		48, 50, 1, 0, 51, 52, 2, 0, 44, 44, 46, 46, 2, 0, 45, 45, 47, 47, 1, 0,
		38, 39, 481, 0, 50, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4, 100, 1, 0, 0, 0,
		6, 102, 1, 0, 0, 0, 8, 140, 1, 0, 0, 0, 10, 142, 1, 0, 0, 0, 12, 150, 1,
		0, 0, 0, 14, 176, 1, 0, 0, 0, 16, 188, 1, 0, 0, 0, 18, 207, 1, 0, 0, 0,
		20, 223, 1, 0, 0, 0, 22, 225, 1, 0, 0, 0, 24, 239, 1, 0, 0, 0, 26, 248,
		1, 0, 0, 0, 28, 256, 1, 0, 0, 0, 30, 283, 1, 0, 0, 0, 32, 285, 1, 0, 0,
		0, 34, 303, 1, 0, 0, 0, 36, 305, 1, 0, 0, 0, 38, 314, 1, 0, 0, 0, 40, 331,
		1, 0, 0, 0, 42, 355, 1, 0, 0, 0, 44, 371, 1, 0, 0, 0, 46, 373, 1, 0, 0,
		0, 48, 409, 1, 0, 0, 0, 50, 51, 3, 2, 1, 0, 51, 52, 5, 0, 0, 1, 52, 53,
		6, 0, -1, 0, 53, 1, 1, 0, 0, 0, 54, 56, 3, 4, 2, 0, 55, 54, 1, 0, 0, 0,
		56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 60, 6, 1, -1, 0, 60, 3, 1, 0, 0, 0, 61, 62, 3, 6, 3, 0, 62,
		63, 6, 2, -1, 0, 63, 101, 1, 0, 0, 0, 64, 65, 3, 8, 4, 0, 65, 66, 6, 2,
		-1, 0, 66, 101, 1, 0, 0, 0, 67, 68, 3, 14, 7, 0, 68, 69, 6, 2, -1, 0, 69,
		101, 1, 0, 0, 0, 70, 71, 3, 18, 9, 0, 71, 72, 6, 2, -1, 0, 72, 101, 1,
		0, 0, 0, 73, 74, 3, 20, 10, 0, 74, 75, 6, 2, -1, 0, 75, 101, 1, 0, 0, 0,
		76, 77, 3, 22, 11, 0, 77, 78, 6, 2, -1, 0, 78, 101, 1, 0, 0, 0, 79, 80,
		3, 28, 14, 0, 80, 81, 6, 2, -1, 0, 81, 101, 1, 0, 0, 0, 82, 83, 3, 30,
		15, 0, 83, 84, 6, 2, -1, 0, 84, 101, 1, 0, 0, 0, 85, 86, 3, 32, 16, 0,
		86, 87, 6, 2, -1, 0, 87, 101, 1, 0, 0, 0, 88, 89, 3, 34, 17, 0, 89, 90,
		6, 2, -1, 0, 90, 101, 1, 0, 0, 0, 91, 92, 3, 36, 18, 0, 92, 93, 6, 2, -1,
		0, 93, 101, 1, 0, 0, 0, 94, 95, 3, 42, 21, 0, 95, 96, 6, 2, -1, 0, 96,
		101, 1, 0, 0, 0, 97, 98, 3, 46, 23, 0, 98, 99, 6, 2, -1, 0, 99, 101, 1,
		0, 0, 0, 100, 61, 1, 0, 0, 0, 100, 64, 1, 0, 0, 0, 100, 67, 1, 0, 0, 0,
		100, 70, 1, 0, 0, 0, 100, 73, 1, 0, 0, 0, 100, 76, 1, 0, 0, 0, 100, 79,
		1, 0, 0, 0, 100, 82, 1, 0, 0, 0, 100, 85, 1, 0, 0, 0, 100, 88, 1, 0, 0,
		0, 100, 91, 1, 0, 0, 0, 100, 94, 1, 0, 0, 0, 100, 97, 1, 0, 0, 0, 101,
		5, 1, 0, 0, 0, 102, 103, 5, 11, 0, 0, 103, 104, 5, 53, 0, 0, 104, 105,
		3, 48, 24, 0, 105, 106, 5, 54, 0, 0, 106, 107, 6, 3, -1, 0, 107, 7, 1,
		0, 0, 0, 108, 109, 5, 12, 0, 0, 109, 110, 3, 48, 24, 0, 110, 111, 5, 55,
		0, 0, 111, 112, 3, 2, 1, 0, 112, 113, 5, 56, 0, 0, 113, 114, 6, 4, -1,
		0, 114, 141, 1, 0, 0, 0, 115, 116, 5, 12, 0, 0, 116, 117, 3, 48, 24, 0,
		117, 118, 5, 55, 0, 0, 118, 119, 3, 2, 1, 0, 119, 121, 5, 56, 0, 0, 120,
		122, 3, 10, 5, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 121,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 5, 13,
		0, 0, 126, 128, 3, 12, 6, 0, 127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0,
		128, 129, 1, 0, 0, 0, 129, 130, 6, 4, -1, 0, 130, 141, 1, 0, 0, 0, 131,
		132, 5, 12, 0, 0, 132, 133, 3, 48, 24, 0, 133, 134, 5, 55, 0, 0, 134, 135,
		3, 2, 1, 0, 135, 136, 5, 56, 0, 0, 136, 137, 5, 13, 0, 0, 137, 138, 3,
		12, 6, 0, 138, 139, 6, 4, -1, 0, 139, 141, 1, 0, 0, 0, 140, 108, 1, 0,
		0, 0, 140, 115, 1, 0, 0, 0, 140, 131, 1, 0, 0, 0, 141, 9, 1, 0, 0, 0, 142,
		143, 5, 13, 0, 0, 143, 144, 5, 12, 0, 0, 144, 145, 3, 48, 24, 0, 145, 146,
		5, 55, 0, 0, 146, 147, 3, 2, 1, 0, 147, 148, 5, 56, 0, 0, 148, 149, 6,
		5, -1, 0, 149, 11, 1, 0, 0, 0, 150, 151, 5, 55, 0, 0, 151, 152, 3, 2, 1,
		0, 152, 153, 5, 56, 0, 0, 153, 154, 6, 6, -1, 0, 154, 13, 1, 0, 0, 0, 155,
		156, 5, 9, 0, 0, 156, 157, 5, 37, 0, 0, 157, 158, 5, 63, 0, 0, 158, 159,
		3, 16, 8, 0, 159, 160, 5, 43, 0, 0, 160, 161, 3, 48, 24, 0, 161, 162, 6,
		7, -1, 0, 162, 177, 1, 0, 0, 0, 163, 164, 5, 9, 0, 0, 164, 165, 5, 37,
		0, 0, 165, 166, 5, 43, 0, 0, 166, 167, 3, 48, 24, 0, 167, 168, 6, 7, -1,
		0, 168, 177, 1, 0, 0, 0, 169, 170, 5, 9, 0, 0, 170, 171, 5, 37, 0, 0, 171,
		172, 5, 63, 0, 0, 172, 173, 3, 16, 8, 0, 173, 174, 5, 59, 0, 0, 174, 175,
		6, 7, -1, 0, 175, 177, 1, 0, 0, 0, 176, 155, 1, 0, 0, 0, 176, 163, 1, 0,
		0, 0, 176, 169, 1, 0, 0, 0, 177, 15, 1, 0, 0, 0, 178, 179, 5, 1, 0, 0,
		179, 189, 6, 8, -1, 0, 180, 181, 5, 2, 0, 0, 181, 189, 6, 8, -1, 0, 182,
		183, 5, 4, 0, 0, 183, 189, 6, 8, -1, 0, 184, 185, 5, 3, 0, 0, 185, 189,
		6, 8, -1, 0, 186, 187, 5, 5, 0, 0, 187, 189, 6, 8, -1, 0, 188, 178, 1,
		0, 0, 0, 188, 180, 1, 0, 0, 0, 188, 182, 1, 0, 0, 0, 188, 184, 1, 0, 0,
		0, 188, 186, 1, 0, 0, 0, 189, 17, 1, 0, 0, 0, 190, 191, 5, 37, 0, 0, 191,
		192, 5, 43, 0, 0, 192, 193, 3, 48, 24, 0, 193, 194, 6, 9, -1, 0, 194, 208,
		1, 0, 0, 0, 195, 196, 5, 37, 0, 0, 196, 197, 5, 51, 0, 0, 197, 198, 5,
		43, 0, 0, 198, 199, 3, 48, 24, 0, 199, 200, 6, 9, -1, 0, 200, 208, 1, 0,
		0, 0, 201, 202, 5, 37, 0, 0, 202, 203, 5, 52, 0, 0, 203, 204, 5, 43, 0,
		0, 204, 205, 3, 48, 24, 0, 205, 206, 6, 9, -1, 0, 206, 208, 1, 0, 0, 0,
		207, 190, 1, 0, 0, 0, 207, 195, 1, 0, 0, 0, 207, 201, 1, 0, 0, 0, 208,
		19, 1, 0, 0, 0, 209, 210, 5, 10, 0, 0, 210, 211, 5, 37, 0, 0, 211, 212,
		5, 43, 0, 0, 212, 213, 3, 48, 24, 0, 213, 214, 6, 10, -1, 0, 214, 224,
		1, 0, 0, 0, 215, 216, 5, 10, 0, 0, 216, 217, 5, 37, 0, 0, 217, 218, 5,
		63, 0, 0, 218, 219, 3, 16, 8, 0, 219, 220, 5, 43, 0, 0, 220, 221, 3, 48,
		24, 0, 221, 222, 6, 10, -1, 0, 222, 224, 1, 0, 0, 0, 223, 209, 1, 0, 0,
		0, 223, 215, 1, 0, 0, 0, 224, 21, 1, 0, 0, 0, 225, 226, 5, 14, 0, 0, 226,
		227, 3, 48, 24, 0, 227, 229, 5, 55, 0, 0, 228, 230, 3, 24, 12, 0, 229,
		228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232,
		1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233, 235, 3, 26, 13, 0, 234, 233, 1,
		0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 5, 56, 0,
		0, 237, 238, 6, 11, -1, 0, 238, 23, 1, 0, 0, 0, 239, 240, 5, 15, 0, 0,
		240, 241, 3, 48, 24, 0, 241, 242, 5, 63, 0, 0, 242, 244, 3, 2, 1, 0, 243,
		245, 5, 22, 0, 0, 244, 243, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246,
		1, 0, 0, 0, 246, 247, 6, 12, -1, 0, 247, 25, 1, 0, 0, 0, 248, 249, 5, 16,
		0, 0, 249, 250, 5, 63, 0, 0, 250, 252, 3, 2, 1, 0, 251, 253, 5, 22, 0,
		0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		255, 6, 13, -1, 0, 255, 27, 1, 0, 0, 0, 256, 257, 5, 17, 0, 0, 257, 258,
		3, 48, 24, 0, 258, 259, 5, 55, 0, 0, 259, 260, 3, 2, 1, 0, 260, 261, 5,
		56, 0, 0, 261, 262, 6, 14, -1, 0, 262, 29, 1, 0, 0, 0, 263, 264, 5, 18,
		0, 0, 264, 265, 5, 37, 0, 0, 265, 266, 5, 19, 0, 0, 266, 267, 5, 36, 0,
		0, 267, 268, 5, 55, 0, 0, 268, 269, 3, 2, 1, 0, 269, 270, 5, 56, 0, 0,
		270, 271, 6, 15, -1, 0, 271, 284, 1, 0, 0, 0, 272, 273, 5, 18, 0, 0, 273,
		274, 5, 37, 0, 0, 274, 275, 5, 19, 0, 0, 275, 276, 3, 48, 24, 0, 276, 277,
		5, 20, 0, 0, 277, 278, 3, 48, 24, 0, 278, 279, 5, 55, 0, 0, 279, 280, 3,
		2, 1, 0, 280, 281, 5, 56, 0, 0, 281, 282, 6, 15, -1, 0, 282, 284, 1, 0,
		0, 0, 283, 263, 1, 0, 0, 0, 283, 272, 1, 0, 0, 0, 284, 31, 1, 0, 0, 0,
		285, 286, 5, 21, 0, 0, 286, 287, 3, 48, 24, 0, 287, 288, 5, 13, 0, 0, 288,
		289, 5, 55, 0, 0, 289, 290, 3, 2, 1, 0, 290, 291, 5, 56, 0, 0, 291, 292,
		6, 16, -1, 0, 292, 33, 1, 0, 0, 0, 293, 294, 5, 22, 0, 0, 294, 304, 6,
		17, -1, 0, 295, 296, 5, 23, 0, 0, 296, 304, 6, 17, -1, 0, 297, 298, 5,
		24, 0, 0, 298, 304, 6, 17, -1, 0, 299, 300, 5, 24, 0, 0, 300, 301, 3, 48,
		24, 0, 301, 302, 6, 17, -1, 0, 302, 304, 1, 0, 0, 0, 303, 293, 1, 0, 0,
		0, 303, 295, 1, 0, 0, 0, 303, 297, 1, 0, 0, 0, 303, 299, 1, 0, 0, 0, 304,
		35, 1, 0, 0, 0, 305, 306, 5, 9, 0, 0, 306, 307, 5, 37, 0, 0, 307, 308,
		5, 63, 0, 0, 308, 309, 5, 57, 0, 0, 309, 310, 3, 16, 8, 0, 310, 311, 5,
		58, 0, 0, 311, 312, 3, 38, 19, 0, 312, 313, 6, 18, -1, 0, 313, 37, 1, 0,
		0, 0, 314, 315, 5, 43, 0, 0, 315, 317, 5, 57, 0, 0, 316, 318, 3, 40, 20,
		0, 317, 316, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319,
		320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 322, 5, 58, 0, 0, 322, 323,
		6, 19, -1, 0, 323, 39, 1, 0, 0, 0, 324, 325, 3, 48, 24, 0, 325, 326, 5,
		61, 0, 0, 326, 327, 6, 20, -1, 0, 327, 332, 1, 0, 0, 0, 328, 329, 3, 48,
		24, 0, 329, 330, 6, 20, -1, 0, 330, 332, 1, 0, 0, 0, 331, 324, 1, 0, 0,
		0, 331, 328, 1, 0, 0, 0, 332, 41, 1, 0, 0, 0, 333, 334, 5, 37, 0, 0, 334,
		335, 5, 62, 0, 0, 335, 336, 5, 29, 0, 0, 336, 337, 5, 53, 0, 0, 337, 338,
		3, 48, 24, 0, 338, 339, 5, 54, 0, 0, 339, 340, 6, 21, -1, 0, 340, 356,
		1, 0, 0, 0, 341, 342, 5, 37, 0, 0, 342, 343, 5, 62, 0, 0, 343, 344, 5,
		30, 0, 0, 344, 345, 5, 53, 0, 0, 345, 346, 5, 54, 0, 0, 346, 356, 6, 21,
		-1, 0, 347, 348, 5, 37, 0, 0, 348, 349, 5, 62, 0, 0, 349, 350, 5, 31, 0,
		0, 350, 351, 5, 53, 0, 0, 351, 352, 3, 48, 24, 0, 352, 353, 5, 54, 0, 0,
		353, 354, 6, 21, -1, 0, 354, 356, 1, 0, 0, 0, 355, 333, 1, 0, 0, 0, 355,
		341, 1, 0, 0, 0, 355, 347, 1, 0, 0, 0, 356, 43, 1, 0, 0, 0, 357, 358, 5,
		37, 0, 0, 358, 359, 5, 62, 0, 0, 359, 360, 5, 32, 0, 0, 360, 372, 6, 22,
		-1, 0, 361, 362, 5, 37, 0, 0, 362, 363, 5, 62, 0, 0, 363, 364, 5, 33, 0,
		0, 364, 372, 6, 22, -1, 0, 365, 366, 5, 37, 0, 0, 366, 367, 5, 57, 0, 0,
		367, 368, 3, 48, 24, 0, 368, 369, 5, 58, 0, 0, 369, 370, 6, 22, -1, 0,
		370, 372, 1, 0, 0, 0, 371, 357, 1, 0, 0, 0, 371, 361, 1, 0, 0, 0, 371,
		365, 1, 0, 0, 0, 372, 45, 1, 0, 0, 0, 373, 374, 5, 37, 0, 0, 374, 375,
		5, 57, 0, 0, 375, 376, 3, 48, 24, 0, 376, 377, 5, 58, 0, 0, 377, 378, 5,
		43, 0, 0, 378, 379, 5, 37, 0, 0, 379, 380, 5, 57, 0, 0, 380, 381, 3, 48,
		24, 0, 381, 382, 5, 58, 0, 0, 382, 383, 6, 23, -1, 0, 383, 47, 1, 0, 0,
		0, 384, 385, 6, 24, -1, 0, 385, 386, 5, 52, 0, 0, 386, 387, 3, 48, 24,
		16, 387, 388, 6, 24, -1, 0, 388, 410, 1, 0, 0, 0, 389, 390, 5, 53, 0, 0,
		390, 391, 3, 48, 24, 0, 391, 392, 5, 54, 0, 0, 392, 393, 6, 24, -1, 0,
		393, 410, 1, 0, 0, 0, 394, 395, 5, 34, 0, 0, 395, 410, 6, 24, -1, 0, 396,
		397, 5, 35, 0, 0, 397, 410, 6, 24, -1, 0, 398, 399, 5, 36, 0, 0, 399, 410,
		6, 24, -1, 0, 400, 401, 5, 6, 0, 0, 401, 410, 6, 24, -1, 0, 402, 403, 5,
		7, 0, 0, 403, 410, 6, 24, -1, 0, 404, 405, 5, 37, 0, 0, 405, 410, 6, 24,
		-1, 0, 406, 407, 3, 44, 22, 0, 407, 408, 6, 24, -1, 0, 408, 410, 1, 0,
		0, 0, 409, 384, 1, 0, 0, 0, 409, 389, 1, 0, 0, 0, 409, 394, 1, 0, 0, 0,
		409, 396, 1, 0, 0, 0, 409, 398, 1, 0, 0, 0, 409, 400, 1, 0, 0, 0, 409,
		402, 1, 0, 0, 0, 409, 404, 1, 0, 0, 0, 409, 406, 1, 0, 0, 0, 410, 448,
		1, 0, 0, 0, 411, 412, 10, 15, 0, 0, 412, 413, 7, 0, 0, 0, 413, 414, 3,
		48, 24, 16, 414, 415, 6, 24, -1, 0, 415, 447, 1, 0, 0, 0, 416, 417, 10,
		14, 0, 0, 417, 418, 7, 1, 0, 0, 418, 419, 3, 48, 24, 15, 419, 420, 6, 24,
		-1, 0, 420, 447, 1, 0, 0, 0, 421, 422, 10, 13, 0, 0, 422, 423, 7, 2, 0,
		0, 423, 424, 3, 48, 24, 14, 424, 425, 6, 24, -1, 0, 425, 447, 1, 0, 0,
		0, 426, 427, 10, 12, 0, 0, 427, 428, 7, 3, 0, 0, 428, 429, 3, 48, 24, 13,
		429, 430, 6, 24, -1, 0, 430, 447, 1, 0, 0, 0, 431, 432, 10, 11, 0, 0, 432,
		433, 7, 4, 0, 0, 433, 434, 3, 48, 24, 12, 434, 435, 6, 24, -1, 0, 435,
		447, 1, 0, 0, 0, 436, 437, 10, 10, 0, 0, 437, 438, 5, 42, 0, 0, 438, 439,
		3, 48, 24, 11, 439, 440, 6, 24, -1, 0, 440, 447, 1, 0, 0, 0, 441, 442,
		10, 9, 0, 0, 442, 443, 5, 41, 0, 0, 443, 444, 3, 48, 24, 10, 444, 445,
		6, 24, -1, 0, 445, 447, 1, 0, 0, 0, 446, 411, 1, 0, 0, 0, 446, 416, 1,
		0, 0, 0, 446, 421, 1, 0, 0, 0, 446, 426, 1, 0, 0, 0, 446, 431, 1, 0, 0,
		0, 446, 436, 1, 0, 0, 0, 446, 441, 1, 0, 0, 0, 447, 450, 1, 0, 0, 0, 448,
		446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 49, 1, 0, 0, 0, 450, 448, 1,
		0, 0, 0, 22, 57, 100, 123, 127, 140, 176, 188, 207, 223, 231, 234, 244,
		252, 283, 303, 319, 331, 355, 371, 409, 446, 448,
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
	SwiftGrammarParserNUMBER       = 34
	SwiftGrammarParserCHAR         = 35
	SwiftGrammarParserSTRING       = 36
	SwiftGrammarParserID           = 37
	SwiftGrammarParserDIF          = 38
	SwiftGrammarParserIG_IG        = 39
	SwiftGrammarParserNOT          = 40
	SwiftGrammarParserOR           = 41
	SwiftGrammarParserAND          = 42
	SwiftGrammarParserIG           = 43
	SwiftGrammarParserMAY_IG       = 44
	SwiftGrammarParserMEN_IG       = 45
	SwiftGrammarParserMAYOR        = 46
	SwiftGrammarParserMENOR        = 47
	SwiftGrammarParserMOD          = 48
	SwiftGrammarParserMUL          = 49
	SwiftGrammarParserDIV          = 50
	SwiftGrammarParserADD          = 51
	SwiftGrammarParserSUB          = 52
	SwiftGrammarParserPARIZQ       = 53
	SwiftGrammarParserPARDER       = 54
	SwiftGrammarParserLLAVEIZQ     = 55
	SwiftGrammarParserLLAVEDER     = 56
	SwiftGrammarParserCORCHIZQ     = 57
	SwiftGrammarParserCORCHDER     = 58
	SwiftGrammarParserQ_MARK       = 59
	SwiftGrammarParserARROW        = 60
	SwiftGrammarParserCOMA         = 61
	SwiftGrammarParserPUNTO        = 62
	SwiftGrammarParserCOLON        = 63
	SwiftGrammarParserWHITESPACE   = 64
	SwiftGrammarParserCOMMENT      = 65
	SwiftGrammarParserLINE_COMMENT = 66
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s             = 0
	SwiftGrammarParserRULE_block         = 1
	SwiftGrammarParserRULE_instruction   = 2
	SwiftGrammarParserRULE_printstmt     = 3
	SwiftGrammarParserRULE_ifstmt        = 4
	SwiftGrammarParserRULE_elseifstmt    = 5
	SwiftGrammarParserRULE_elsestmt      = 6
	SwiftGrammarParserRULE_varstmt       = 7
	SwiftGrammarParserRULE_tipo          = 8
	SwiftGrammarParserRULE_varasgmt      = 9
	SwiftGrammarParserRULE_conststmt     = 10
	SwiftGrammarParserRULE_switchstmt    = 11
	SwiftGrammarParserRULE_casestmt      = 12
	SwiftGrammarParserRULE_defaultstmt   = 13
	SwiftGrammarParserRULE_whilestmt     = 14
	SwiftGrammarParserRULE_forstmt       = 15
	SwiftGrammarParserRULE_guardstmt     = 16
	SwiftGrammarParserRULE_transferstmt  = 17
	SwiftGrammarParserRULE_vectorstmt    = 18
	SwiftGrammarParserRULE_definestmt    = 19
	SwiftGrammarParserRULE_listexpr      = 20
	SwiftGrammarParserRULE_methodvec     = 21
	SwiftGrammarParserRULE_methodvecrtrn = 22
	SwiftGrammarParserRULE_vecaccess     = 23
	SwiftGrammarParserRULE_expr          = 24
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
		p.SetState(50)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(51)
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(54)

				var _x = p.Instruction()

				localctx.(*BlockContext)._instruction = _x
			}
			localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(57)
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
	Switchstmt() ISwitchstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Transferstmt() ITransferstmtContext
	Vectorstmt() IVectorstmtContext
	Methodvec() IMethodvecContext
	Vecaccess() IVecaccessContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	inst          interfaces.Instruction
	_printstmt    IPrintstmtContext
	_ifstmt       IIfstmtContext
	_varstmt      IVarstmtContext
	_varasgmt     IVarasgmtContext
	_conststmt    IConststmtContext
	_switchstmt   ISwitchstmtContext
	_whilestmt    IWhilestmtContext
	_forstmt      IForstmtContext
	_guardstmt    IGuardstmtContext
	_transferstmt ITransferstmtContext
	_vectorstmt   IVectorstmtContext
	_methodvec    IMethodvecContext
	_vecaccess    IVecaccessContext
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

func (s *InstructionContext) Get_switchstmt() ISwitchstmtContext { return s._switchstmt }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_guardstmt() IGuardstmtContext { return s._guardstmt }

func (s *InstructionContext) Get_transferstmt() ITransferstmtContext { return s._transferstmt }

func (s *InstructionContext) Get_vectorstmt() IVectorstmtContext { return s._vectorstmt }

func (s *InstructionContext) Get_methodvec() IMethodvecContext { return s._methodvec }

func (s *InstructionContext) Get_vecaccess() IVecaccessContext { return s._vecaccess }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)

			var _x = p.Varstmt()

			localctx.(*InstructionContext)._varstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_varstmt().GetVar_()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)

			var _x = p.Varasgmt()

			localctx.(*InstructionContext)._varasgmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_varasgmt().GetAsgmt()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)

			var _x = p.Conststmt()

			localctx.(*InstructionContext)._conststmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_conststmt().GetConst_()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(76)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSwitchinstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(79)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhileinstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(82)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetForinstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(85)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGuardinstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(88)

			var _x = p.Transferstmt()

			localctx.(*InstructionContext)._transferstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_transferstmt().GetTrns()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(91)

			var _x = p.Vectorstmt()

			localctx.(*InstructionContext)._vectorstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vectorstmt().GetVectorinstr()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(94)

			var _x = p.Methodvec()

			localctx.(*InstructionContext)._methodvec = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_methodvec().GetMethodinstr()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(97)

			var _x = p.Vecaccess()

			localctx.(*InstructionContext)._vecaccess = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vecaccess().GetVecacc()

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
		p.SetState(102)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)

		var _x = p.expr(0)

		localctx.(*PrintstmtContext)._expr = _x
	}
	{
		p.SetState(105)
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
	ELSE() antlr.TerminalNode
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

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
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
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_ifstmt)
	var _la int

	var _alt int

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(110)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(112)
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
			p.SetState(115)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(117)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.Block()

			localctx.(*IfstmtContext).firstBlk = _x
		}
		{
			p.SetState(119)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(120)

					var _x = p.Elseifstmt()

					localctx.(*IfstmtContext)._elseifstmt = _x
				}
				localctx.(*IfstmtContext).elif = append(localctx.(*IfstmtContext).elif, localctx.(*IfstmtContext)._elseifstmt)

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserLLAVEIZQ {
			{
				p.SetState(126)

				var _x = p.Elsestmt()

				localctx.(*IfstmtContext)._elsestmt = _x
			}

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
			p.SetState(131)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(133)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.Block()

			localctx.(*IfstmtContext).firstBlk = _x
		}
		{
			p.SetState(135)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)

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
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_elseifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)

		var _m = p.Match(SwiftGrammarParserELSE)

		localctx.(*ElseifstmtContext)._ELSE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(SwiftGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)

		var _x = p.expr(0)

		localctx.(*ElseifstmtContext)._expr = _x
	}
	{
		p.SetState(145)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)

		var _x = p.Block()

		localctx.(*ElseifstmtContext)._block = _x
	}
	{
		p.SetState(147)
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
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_elsestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)

		var _x = p.Block()

		localctx.(*ElsestmtContext)._block = _x
	}
	{
		p.SetState(152)
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
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_varstmt)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)

			var _x = p.Tipo()

			localctx.(*VarstmtContext)._tipo = _x
		}
		{
			p.SetState(159)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

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
			p.SetState(163)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)

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
			p.SetState(169)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*VarstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)

			var _x = p.Tipo()

			localctx.(*VarstmtContext)._tipo = _x
		}
		{
			p.SetState(173)

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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_tipo)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
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
			p.SetState(180)
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
			p.SetState(182)
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
			p.SetState(184)
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
			p.SetState(186)
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
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_varasgmt)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)

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
			p.SetState(195)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Match(SwiftGrammarParserADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

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
			p.SetState(201)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VarasgmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(SwiftGrammarParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)

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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_conststmt)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*ConststmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConststmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)

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
			p.SetState(215)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*ConststmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConststmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(SwiftGrammarParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)

			var _x = p.Tipo()

			localctx.(*ConststmtContext)._tipo = _x
		}
		{
			p.SetState(219)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)

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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_switchstmt)

	var switchInterfaces = []interface{}{}
	var interfacelist []ICasestmtContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)

		var _m = p.Match(SwiftGrammarParserSWITCH)

		localctx.(*SwitchstmtContext)._SWITCH = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)

		var _x = p.expr(0)

		localctx.(*SwitchstmtContext)._expr = _x
	}
	{
		p.SetState(227)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCASE {
		{
			p.SetState(228)

			var _x = p.Casestmt()

			localctx.(*SwitchstmtContext)._casestmt = _x
		}
		localctx.(*SwitchstmtContext).casesvar = append(localctx.(*SwitchstmtContext).casesvar, localctx.(*SwitchstmtContext)._casestmt)

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserDEFAULT {
		{
			p.SetState(233)
			p.Defaultstmt()
		}

	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_casestmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*CasestmtContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)

		var _x = p.expr(0)

		localctx.(*CasestmtContext)._expr = _x
	}
	{
		p.SetState(241)
		p.Match(SwiftGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)

		var _x = p.Block()

		localctx.(*CasestmtContext)._block = _x
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserBREAK {
		{
			p.SetState(243)
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
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_defaultstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(SwiftGrammarParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(SwiftGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)

		var _x = p.Block()

		localctx.(*DefaultstmtContext)._block = _x
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SwiftGrammarParserBREAK {
		{
			p.SetState(251)
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
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(258)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(260)
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

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

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
	ID() antlr.TerminalNode
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

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) Set_STRING(v antlr.Token) { s._STRING = v }

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

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
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
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_forstmt)

	var cadena interfaces.Expression
	var str string

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ForstmtContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(269)
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
		}()), cadena, nil, nil, localctx.(*ForstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)

			var _m = p.Match(SwiftGrammarParserFOR)

			localctx.(*ForstmtContext)._FOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(SwiftGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).left = _x
		}
		{
			p.SetState(276)
			p.Match(SwiftGrammarParserRANGEPTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _x = p.expr(0)

			localctx.(*ForstmtContext).right = _x
		}
		{
			p.SetState(278)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

			var _x = p.Block()

			localctx.(*ForstmtContext)._block = _x
		}
		{
			p.SetState(280)
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
		}()), nil, localctx.(*ForstmtContext).GetLeft().GetE(), localctx.(*ForstmtContext).GetRight().GetE(), localctx.(*ForstmtContext).Get_block().GetBlk())

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
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(287)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(290)
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
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_transferstmt)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)

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
			p.SetState(295)

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
			p.SetState(297)

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
			p.SetState(299)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*TransferstmtContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)

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
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_vectorstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)

		var _m = p.Match(SwiftGrammarParserVAR)

		localctx.(*VectorstmtContext)._VAR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VectorstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(SwiftGrammarParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)

		var _x = p.Tipo()

		localctx.(*VectorstmtContext)._tipo = _x
	}
	{
		p.SetState(310)
		p.Match(SwiftGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)

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
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_definestmt)

	var defVecInterfaces []interface{}

	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13511056580149440) != 0) {
		{
			p.SetState(316)

			var _x = p.Listexpr()

			localctx.(*DefinestmtContext)._listexpr = _x
		}
		localctx.(*DefinestmtContext).lista = append(localctx.(*DefinestmtContext).lista, localctx.(*DefinestmtContext)._listexpr)

		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(321)
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
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_listexpr)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(324)

			var _x = p.expr(0)

			localctx.(*ListexprContext)._expr = _x
		}
		{
			p.SetState(325)
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
			p.SetState(328)

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
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_methodvec)
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)

			var _m = p.Match(SwiftGrammarParserAPPEND)

			localctx.(*MethodvecContext)._APPEND = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)

			var _x = p.expr(0)

			localctx.(*MethodvecContext)._expr = _x
		}
		{
			p.SetState(338)
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
			p.SetState(341)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)

			var _m = p.Match(SwiftGrammarParserREMOVELAST)

			localctx.(*MethodvecContext)._REMOVELAST = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
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
			p.SetState(347)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)

			var _m = p.Match(SwiftGrammarParserREMOVE)

			localctx.(*MethodvecContext)._REMOVE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)

			var _x = p.expr(0)

			localctx.(*MethodvecContext)._expr = _x
		}
		{
			p.SetState(352)
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
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_methodvecrtrn)
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecrtrnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)

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
			p.SetState(361)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecrtrnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)

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
			p.SetState(365)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*MethodvecrtrnContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)
			p.Match(SwiftGrammarParserCORCHIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)

			var _x = p.expr(0)

			localctx.(*MethodvecrtrnContext)._expr = _x
		}
		{
			p.SetState(368)
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

	// SetFirstId sets the firstId token.
	SetFirstId(antlr.Token)

	// SetSecondId sets the secondId token.
	SetSecondId(antlr.Token)

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

func (s *VecaccessContext) SetFirstId(v antlr.Token) { s.firstId = v }

func (s *VecaccessContext) SetSecondId(v antlr.Token) { s.secondId = v }

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
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_vecaccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VecaccessContext).firstId = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(375)

		var _x = p.expr(0)

		localctx.(*VecaccessContext).first = _x
	}
	{
		p.SetState(376)
		p.Match(SwiftGrammarParserCORCHDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(377)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VecaccessContext).secondId = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.Match(SwiftGrammarParserCORCHIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(380)

		var _x = p.expr(0)

		localctx.(*VecaccessContext).second = _x
	}
	{
		p.SetState(381)
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

	// Get_methodvecrtrn returns the _methodvecrtrn rule contexts.
	Get_methodvecrtrn() IMethodvecrtrnContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

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
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	CHAR() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	ID() antlr.TerminalNode
	Methodvecrtrn() IMethodvecrtrnContext
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
	_FAL           antlr.Token
	_ID            antlr.Token
	_methodvecrtrn IMethodvecrtrnContext
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

func (s *ExprContext) Get_methodvecrtrn() IMethodvecrtrnContext { return s._methodvecrtrn }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(385)

			var _m = p.Match(SwiftGrammarParserSUB)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)

			var _x = p.expr(16)

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
			p.SetState(389)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(391)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 3:
		{
			p.SetState(394)

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

	case 4:
		{
			p.SetState(396)

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

	case 5:
		{
			p.SetState(398)

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

	case 6:
		{
			p.SetState(400)

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

	case 7:
		{
			p.SetState(402)

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

	case 8:
		{
			p.SetState(404)

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

	case 9:
		{
			p.SetState(406)

			var _x = p.Methodvecrtrn()

			localctx.(*ExprContext)._methodvecrtrn = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_methodvecrtrn().GetMethodinstrtrn()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(446)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(412)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1970324836974592) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(413)

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

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(417)

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
					p.SetState(418)

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

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(422)

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
					p.SetState(423)

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

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(426)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(427)

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
					p.SetState(428)

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

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(432)

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
					p.SetState(433)

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

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(437)
					p.Match(SwiftGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(438)

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

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(442)
					p.Match(SwiftGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(443)

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

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(450)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	case 24:
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
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
