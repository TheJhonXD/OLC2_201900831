package server

import (
	"Server/environment"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

type CodeEntry struct {
	Code string `json:"code"`
}

type AllInOne struct {
	Message string                `json:"message"`
	Error   []environment.Error_  `json:"error"`
	Symbol  []environment.SymbolT `json:"symbol"`
}

type Data struct {
	Grammar    string `json:"grammar"`
	Lexgrammar string `json:"lexgrammar"`
	Input      string `json:"input"`
	Start      string `json:"start"`
}

type Result struct {
	Rules   []string `json:"rules"`
	Input   string   `json:"input"`
	Symbols []string `json:"symbols"`
	Svgtree string   `json:"svgtree"`
}

type Tree struct {
	Warnings     []string `json:"warnings"`
	ParserErrors []string `json:"parser_grammar_errors"`
	LexerErrors  []string `json:"lexer_grammar_errors"`
	Result       Result   `json:"result"`
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello world!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func InputOutput(w http.ResponseWriter, r *http.Request) {
	request := CodeEntry{}
	response := AllInOne{}
	// var aux []environment.Error_
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	response.Message, response.Error, response.Symbol = analyzer(request.Code)
	// fmt.Println(aux)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	filescanner := bufio.NewScanner(file)
	// filescanner.Split(bufio.ScanLines)
	var text string
	for filescanner.Scan() {
		text += filescanner.Text() + "\n"
	}
	return text
}

func TreeParse(w http.ResponseWriter, r *http.Request) {
	url := "http://lab.antlr.org/parse/"
	/************ CODIGO DE ENTRADA ************/
	request := CodeEntry{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	/* *************************************** */
	grammar := readFile("other/grammar.txt")
	lexgrammar := readFile("other/lexgrammar.txt")
	res := &Tree{}
	/* data := Data{
		Grammar:    "parser grammar ExprParser;\noptions { tokenVocab=ExprLexer; }\n\nprogram : stat EOF | def EOF;\nstat: ID '=' expr ';' | expr ';';\ndef : ID '(' ID (',' ID)* ')' '{' stat* '}' ;\nexpr: ID | INT | func | 'not' expr | expr 'and' expr | expr 'or' expr;\nfunc : ID '(' expr (',' expr)* ')' ;",
		Lexgrammar: "lexer grammar ExprLexer;\nAND : 'and' ;\nOR : 'or' ;\nNOT : 'not' ;\nEQ : '=' ;\nCOMMA : ',' ;\nSEMI : ';' ;\nLPAREN : '(' ;\nRPAREN : ')' ;\nLCURLY : '{' ;\nRCURLY : '}' ;\nINT : [0-9]+ ;\nID: [a-zA-Z_][a-zA-Z_0-9]* ;\nWS: [ \\t\\n\\r\\f]+ -> skip ;",
		Input:      "f(x,y) {\n  a = 3;\n  x and y;\n}",
		Start:      "program",
	} */
	data := Data{
		Grammar:    grammar,
		Lexgrammar: lexgrammar,
		Input:      request.Code,
		Start:      "s",
	}
	content, err := json.Marshal(data)
	resp, err := http.Post(url, "Content-Type: application/json", bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("Error al hacer la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
