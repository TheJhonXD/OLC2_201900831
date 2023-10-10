grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "Server/interfaces"
    import "Server/environment"
    import "Server/expressions"
    import "Server/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

instruction returns [interfaces.Instruction inst]
: printstmt  { $inst = $printstmt.prnt}
| callfuncinstruction { $inst = $callfuncinstruction.callfuncinstr}
| ifstmt { $inst = $ifstmt.ifinstr }
| varstmt { $inst = $varstmt.var}
| varasgmt { $inst = $varasgmt.asgmt}
| conststmt { $inst = $conststmt.const}
| switchstmt { $inst = $switchstmt.switchinstr}
| whilestmt { $inst = $whilestmt.whileinstr}
| forstmt { $inst = $forstmt.forinstr}
| guardstmt { $inst = $guardstmt.guardinstr}
| transferstmt { $inst = $transferstmt.trns}
| vectorstmt { $inst = $vectorstmt.vectorinstr}
| methodvec { $inst = $methodvec.methodinstr}
| vecaccess { $inst = $vecaccess.vecacc}
| funcstmt { $inst = $funcstmt.funcinstr}
;

//* Instrucción print
printstmt returns [interfaces.Instruction prnt]
@init{
    var listExpr []interface{}
}
: PRINT PARIZQ lista+=printexprlist+ PARDER { 
    for _, e := range localctx.(*PrintstmtContext).GetLista() {
        listExpr = append(listExpr, e.GetPrntexpr())
    }
    $prnt = instructions.NewPrint($PRINT.line, $PRINT.pos, listExpr)
}
;

printexprlist returns [interfaces.Expression prntexpr]
: expr COMA { $prntexpr = $expr.e }
| expr { $prntexpr = $expr.e }
;

//* Sentencia if else
ifstmt returns [interfaces.Instruction ifinstr]
: IF expr LLAVEIZQ block LLAVEDER { $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil, nil) }
| IF expr LLAVEIZQ firstBlk=block LLAVEDER elif+=elseifstmt+ elsestmt{ 
    var ifInterfaces []interface{}
    // fmt.Println($elif)
    for _, e := range localctx.(*IfstmtContext).GetElif() {
        ifInterfaces = append(ifInterfaces, e.GetElifinstr())
        fmt.Println(e.GetElifinstr())
    }
    $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $firstBlk.blk, ifInterfaces, $elsestmt.elseinstr)
}
| IF expr LLAVEIZQ firstBlk=block LLAVEDER elif+=elseifstmt+ {
    var ifInterfaces []interface{}
    // fmt.Println($elif)
    for _, e := range localctx.(*IfstmtContext).GetElif() {
        ifInterfaces = append(ifInterfaces, e.GetElifinstr())
        fmt.Println(e.GetElifinstr())
    }
    $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $firstBlk.blk, ifInterfaces, nil)
}
| IF expr LLAVEIZQ firstBlk=block LLAVEDER elsestmt { $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $firstBlk.blk, nil, $elsestmt.elseinstr) }
;

//* Instrucción else if
elseifstmt returns [interfaces.Instruction elifinstr]
: ELSE IF expr LLAVEIZQ block LLAVEDER { 
    $elifinstr = instructions.NewIf($ELSE.line, $ELSE.pos, $expr.e, nil, $block.blk, nil)
}
;

//* Instrucción else
elsestmt returns [[]interface{} elseinstr]
: ELSE LLAVEIZQ block LLAVEDER { $elseinstr = $block.blk }
;

//* Declaracion de variables
varstmt returns [interfaces.Instruction var]
: VAR ID COLON tipo IG expr {$var = instructions.NewStmt($VAR.line, $VAR.pos, $ID.text, $tipo.rtipo, $expr.e, false)}
| VAR ID IG expr            {$var = instructions.NewStmt($VAR.line, $VAR.pos, $ID.text, environment.NULL, $expr.e, false)}
| VAR ID COLON tipo Q_MARK  {$var = instructions.NewStmt($VAR.line, $VAR.pos, $ID.text, $tipo.rtipo, expressions.NewPrimitive($Q_MARK.line, $Q_MARK.pos, environment.WILDCARD, environment.WILDCARD), false)}
;

//* Tipo  de variables
tipo returns [environment.TipoExpresion rtipo]
:INT {$rtipo = 0}
|FLOAT {$rtipo = 1}
|STR {$rtipo = 2}
|BOOL {$rtipo = 3}
|CHARACTER {$rtipo = 5}
|CORCHIZQ INT CORCHDER {$rtipo = 7}
;

//* Asignacion de valor a una variable
varasgmt returns [interfaces.Instruction asgmt]
: ID IG expr {$asgmt = instructions.NewAsgmt($ID.line, $ID.pos, $ID.text, $expr.e)}
| ID ADD IG expr {$asgmt = instructions.NewOpAsgmt($ID.line, $ID.pos, $ID.text, $expr.e, "+")}
| ID SUB IG expr {$asgmt = instructions.NewOpAsgmt($ID.line, $ID.pos, $ID.text, $expr.e, "-")}
;

//* Declaracion de constantes
conststmt returns [interfaces.Instruction const]
: LET ID IG expr {$const = instructions.NewStmt($LET.line, $LET.pos, $ID.text, environment.NULL, $expr.e, true)}
| LET ID COLON tipo IG expr {$const = instructions.NewStmt($LET.line, $LET.pos, $ID.text, $tipo.rtipo, $expr.e, true)}
;

//* Sentencia Switch-Case
switchstmt returns [interfaces.Instruction switchinstr]
@init{
    var switchInterfaces = []interface{}{}
    var interfacelist []ICasestmtContext
}
: SWITCH expr LLAVEIZQ casesvar+=casestmt+ defaultstmt? LLAVEDER { 
    interfacelist = localctx.(*SwitchstmtContext).GetCasesvar()
    for _, e := range interfacelist {
        switchInterfaces = append(switchInterfaces, e.GetCaseinstr())
    }
    $switchinstr = instructions.NewSwitch($SWITCH.line, $SWITCH.pos, $expr.e, switchInterfaces, nil)
}
;

//* Instrucción case
casestmt returns [interfaces.Instruction caseinstr]
: CASE expr COLON block BREAK? { 
    $caseinstr = instructions.NewSwitch($CASE.line, $CASE.pos, $expr.e, $block.blk, nil) 
}
;

//* Instrucción default
defaultstmt returns [[]interface{} definstr]
: DEFAULT COLON block BREAK? { $definstr = $block.blk }
;

//* Sentencia While
whilestmt returns [interfaces.Instruction whileinstr]
: WHILE expr LLAVEIZQ block LLAVEDER { $whileinstr = instructions.NewWhile($WHILE.line, $WHILE.pos, $expr.e, $block.blk) }
;

//* Sentencia For
forstmt returns [interfaces.Instruction forinstr]
@init{
    var cadena interfaces.Expression
    var str string
}
: FOR ID IN STRING LLAVEIZQ block LLAVEDER { 
    str = $STRING.text
    cadena = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1], environment.STRING)
    $forinstr = instructions.NewForIn($FOR.line, $FOR.pos, $ID.text, cadena, nil, nil, "", $block.blk) 
}
| FOR ID IN left=expr RANGEPTS right=expr LLAVEIZQ block LLAVEDER { 
    $forinstr = instructions.NewForIn($FOR.line, $FOR.pos, $ID.text, nil, $left.e, $right.e, "", $block.blk)
}
| FOR first=ID IN second=ID LLAVEIZQ block LLAVEDER {
    $forinstr = instructions.NewForIn($FOR.line, $FOR.pos, $first.text, nil, nil, nil, $second.text, $block.blk)
}
;

//* Instrucción Guard
guardstmt returns [interfaces.Instruction guardinstr]
: GUARD expr ELSE LLAVEIZQ block LLAVEDER { $guardinstr = instructions.NewGuard($GUARD.line, $GUARD.pos, $expr.e, $block.blk) }
;


//* Sentencias de transferencia
transferstmt returns [interfaces.Instruction trns]
: BREAK { $trns = instructions.NewBreak($BREAK.line, $BREAK.pos) }
| CONTINUE { $trns = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos) }
| RETURN { $trns = instructions.NewReturn($RETURN.line, $RETURN.pos, nil) }
| RETURN expr { $trns = instructions.NewReturn($RETURN.line, $RETURN.pos, $expr.e) }
;

//* Declaración de vectores
vectorstmt returns [interfaces.Instruction vectorinstr]
: VAR ID COLON CORCHIZQ tipo CORCHDER definestmt { 
    $vectorinstr = instructions.NewVectorStmt($VAR.line, $VAR.pos, $ID.text, $tipo.rtipo, $definestmt.defineinstr)
}
;

//* Definicion de valor de un vector
definestmt returns [[]interface{} defineinstr]
@init{
    var defVecInterfaces []interface{}
}
: IG CORCHIZQ lista += listexpr+ CORCHDER{ 
    for _, e := range localctx.(*DefinestmtContext).GetLista() {
        // fmt.Println(fmt.Sprintf("%T", e.GetListe()))
        defVecInterfaces = append(defVecInterfaces, e.GetListe())
    }
    $defineinstr = defVecInterfaces
}
| IG CORCHIZQ CORCHDER { $defineinstr = nil }
;

//* Lista de expresiones
listexpr returns [interfaces.Expression liste]
: expr COMA { $liste = $expr.e }
| expr { $liste = $expr.e }
;

//* Metodos de vectores
methodvec returns [interfaces.Instruction methodinstr]
: ID PUNTO APPEND PARIZQ expr PARDER { $methodinstr = instructions.NewVectorMethod($ID.line, $ID.pos, $ID.text, $expr.e, $APPEND.text) }
| ID PUNTO REMOVELAST PARIZQ PARDER { $methodinstr = instructions.NewVectorMethod($ID.line, $ID.pos, $ID.text, nil, $REMOVELAST.text) }
| ID PUNTO REMOVE PARIZQ expr PARDER { $methodinstr = instructions.NewVectorMethod($ID.line, $ID.pos, $ID.text, $expr.e, $REMOVE.text) }
;

//* Metodos de vectores que retornan un valor
methodvecrtrn returns [interfaces.Expression methodinstrtrn]
: ID PUNTO EMPTY { $methodinstrtrn = expressions.NewVector($ID.line, $ID.pos, $ID.text, nil, $EMPTY.text) }
| ID PUNTO COUNT { $methodinstrtrn = expressions.NewVector($ID.line, $ID.pos, $ID.text, nil, $COUNT.text) }
| ID CORCHIZQ expr CORCHDER { $methodinstrtrn = expressions.NewVector($ID.line, $ID.pos, $ID.text, $expr.e, "access") }
;

//* Acceso a vectores
vecaccess returns [interfaces.Instruction vecacc]
: firstId=ID CORCHIZQ first=expr CORCHDER IG secondId=ID CORCHIZQ second=expr CORCHDER { 
    $vecacc = instructions.NewVectorAsgmt($firstId.line, $firstId.pos, $firstId.text, $first.e, $second.e, $secondId.text) 
}
| ID CORCHIZQ first=expr CORCHDER IG second=expr { $vecacc = instructions.NewVectorAsgmt($ID.line, $ID.pos, $ID.text, $first.e, $second.e, "") }
;


//* Declaración de matrices
/* tipomatrix 
: CORCHIZQ tipomatrix CORCHIZQ
| tipo
;

defmatrix 
: listvalmatrix
simplematrix
;

listvalmatrix
: listvalmatrix2
;

listvalmatrix2
: listvalmatrix2 COMA listvalmatrix
| listvalmatrix
| listexpr
;

simplematrix
: tipomatrix PARIZQ simplematrix COMA NUMBER PARDER
| tipomatrix PARIZQ expr COMA simplematrix PARDER
; */

//* Declaración de Structs
structstmt
: STRUCT ID LLAVEIZQ list+=structattrlist+ LLAVEDER
;

structattrlist 
: (LET | VAR) ID (: tipo)? (IG expr)?
;


//* Funciones
funcstmt returns [interfaces.Instruction funcinstr]
@init{
    var args []interface{}
}
: FUNC ID PARIZQ lista+=listparam+ PARDER ARROW tipo LLAVEIZQ block LLAVEDER {
    for _, e := range localctx.(*FuncstmtContext).GetLista() {
        // fmt.Println(fmt.Sprintf("%T", e.GetListparaminstr()))
        args = append(args, e.GetListparaminstr())
    }
    $funcinstr = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, args, $tipo.rtipo, $block.blk)
}
| FUNC ID PARIZQ lista+=listparam+ PARDER LLAVEIZQ block LLAVEDER {
    for _, e := range localctx.(*FuncstmtContext).GetLista() {
        // fmt.Println(fmt.Sprintf("%T", e.GetListparaminstr()))
        args = append(args, e.GetListparaminstr())
    }
    $funcinstr = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, args, 4, $block.blk)
}
| FUNC ID PARIZQ PARDER ARROW tipo LLAVEIZQ block LLAVEDER { 
    $funcinstr = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, nil, $tipo.rtipo, $block.blk) 
}
| FUNC ID PARIZQ PARDER LLAVEIZQ block LLAVEDER { 
    $funcinstr = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, nil, 4, $block.blk) 
}
;

//* Lista de parametros
listparam returns [interfaces.Instruction listparaminstr]
: extr=(ID | CAME )? inter=ID COLON INOUT? tipo COMA { 
    var flag bool
    var aux string
    if $INOUT != nil {
        flag = true
    }else{
        flag = false
    }
    if $extr != nil {
        aux = $extr.text
    }
    $listparaminstr = instructions.NewParam($extr.line, $extr.pos, aux, $inter.text, flag, $tipo.rtipo) 
}
| extr=(ID | CAME )? inter=ID COLON INOUT? tipo { 
    var flag bool
    var aux string
    if $INOUT != nil {
        flag = true
    }else{
        flag = false
    }
    if $extr != nil {
        aux = $extr.text
    }
    $listparaminstr = instructions.NewParam($extr.line, $extr.pos, aux, $inter.text, flag, $tipo.rtipo) 
}
;

//TODO: Hacer lo mismo que en funcstmt

callfunc returns [interfaces.Expression callfuncexpr]
@init{
    var args []interface{}
}
: ID PARIZQ lista+=listparamcall+ PARDER {
    for _, e := range localctx.(*CallfuncContext).GetLista() {
        // fmt.Println(fmt.Sprintf("%T", e.GetListparamcallinstr()))
        args = append(args, e.GetListparamcallinstr())
    }
    $callfuncexpr = expressions.NewCallFunc($ID.line, $ID.pos, $ID.text, args)
}
| ID PARIZQ PARDER { $callfuncexpr = expressions.NewCallFunc($ID.line, $ID.pos, $ID.text, nil) }
;

callfuncinstruction returns [interfaces.Instruction callfuncinstr]
@init{
    var args []interface{}
}
: ID PARIZQ lista+=listparamcall+ PARDER {
    for _, e := range localctx.(*CallfuncinstructionContext).GetLista() {
        // fmt.Println(fmt.Sprintf("%T", e.GetListparamcallinstr()))
        args = append(args, e.GetListparamcallinstr())
    }
    $callfuncinstr = instructions.NewCallFunc($ID.line, $ID.pos, $ID.text, args)
}
| ID PARIZQ PARDER { $callfuncinstr = instructions.NewCallFunc($ID.line, $ID.pos, $ID.text, nil) }
;

listparamcall returns [interfaces.Expression listparamcallinstr]
: (ID COLON)? AMP? expr COMA { 
    var flag bool
    if $AMP != nil {
        flag = true
    }
    $listparamcallinstr = expressions.NewCallParams($expr.start.GetLine(), $expr.start.GetColumn(), $ID.text, flag, $expr.e) 
}
| (ID COLON)? AMP? expr { 
    var flag bool
    if $AMP != nil {
        flag = true
    }
    $listparamcallinstr = expressions.NewCallParams($expr.start.GetLine(), $expr.start.GetColumn(), $ID.text, flag, $expr.e)
 }
;

funcembed returns [interfaces.Expression funcembedexpr]
: STR PARIZQ expr PARDER { $funcembedexpr = expressions.NewString($STR.line, $STR.pos, $expr.e) }
| INT PARIZQ expr PARDER { $funcembedexpr = expressions.NewInteger($INT.line, $INT.pos, $expr.e) }
| FLOAT PARIZQ expr PARDER { $funcembedexpr = expressions.NewFloat($FLOAT.line, $FLOAT.pos, $expr.e) }
;

//* Gramatica para Expresiones
expr returns [interfaces.Expression e]
//Agregar para menos unario
: op=SUB right=expr { $e = expressions.NewOperation($right.start.GetLine(), $right.start.GetColumn(), $right.e, "unario", expressions.NewPrimitive($op.line, $op.pos, -1, environment.INTEGER)) }
| op=NOT right=expr { $e = expressions.NewOperation($right.start.GetLine(), $right.start.GetColumn(), $right.e, $op.text, expressions.NewPrimitive($op.line, $op.pos, false, environment.BOOLEAN)) }
| left=expr op=(MUL|DIV|MOD) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $AND.text, $right.e) }
| left=expr OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $OR.text, $right.e) }
| PARIZQ expr PARDER { $e = $expr.e }
| NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| CHAR { 
        str := $CHAR.text
        $e = expressions.NewPrimitive($CHAR.line, $CHAR.pos, str[1:len(str)-1], environment.CHAR) 
    }
| STRING
    {
        str := $STRING.text
        $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }                        
| TRU { $e = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| callfunc { $e = $callfunc.callfuncexpr }
| funcembed { $e = $funcembed.funcembedexpr }
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
| ID { $e = expressions.NewVar($ID.line, $ID.pos, $ID.text) }
| methodvecrtrn { $e = $methodvecrtrn.methodinstrtrn }
| NIL { $e = expressions.NewPrimitive($NIL.line, $NIL.pos, "nil", environment.NULL) }
;