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
: printstmt { $inst = $printstmt.prnt}
| ifstmt { $inst = $ifstmt.ifinstr }
| varstmt { $inst = $varstmt.var}
| varasgmt { $inst = $varasgmt.asgmt}
| conststmt { $inst = $conststmt.const}
| switchstmt { $inst = $switchstmt.switchinstr}
| whilestmt { $inst = $whilestmt.whileinstr}
| forstmt { $inst = $forstmt.forinstr}
| guardstmt { $inst = $guardstmt.guardinstr}
| transferstmt { $inst = $transferstmt.trns}
;

//* Instrucción print
printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
;

//* Sentencia if else
ifstmt returns [interfaces.Instruction ifinstr]
: IF expr LLAVEIZQ block LLAVEDER { $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil, nil) }
| IF expr LLAVEIZQ firstBlk=block LLAVEDER elif+=elseifstmt+ ELSE elsestmt?{ 
    var ifInterfaces []interface{}
    // fmt.Println($elif)
    for _, e := range localctx.(*IfstmtContext).GetElif() {
        ifInterfaces = append(ifInterfaces, e.GetElifinstr())
        fmt.Println(e.GetElifinstr())
    }
    $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $firstBlk.blk, ifInterfaces, $elsestmt.elseinstr)
 }
| IF expr LLAVEIZQ firstBlk=block LLAVEDER ELSE elsestmt { $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $firstBlk.blk, nil, $elsestmt.elseinstr) }
;

//* Instrucción else if
elseifstmt returns [interfaces.Instruction elifinstr]
: ELSE IF expr LLAVEIZQ block LLAVEDER { 
    $elifinstr = instructions.NewIf($ELSE.line, $ELSE.pos, $expr.e, nil, $block.blk, nil)
}
;

//* Instrucción else
elsestmt returns [[]interface{} elseinstr]
: LLAVEIZQ block LLAVEDER { $elseinstr = $block.blk }
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
    $forinstr = instructions.NewForIn($FOR.line, $FOR.pos, $ID.text, cadena, nil, nil, $block.blk) 
}
| FOR ID IN left=expr RANGEPTS right=expr LLAVEIZQ block LLAVEDER { 
    $forinstr = instructions.NewForIn($FOR.line, $FOR.pos, $ID.text, nil, $left.e, $right.e, $block.blk)
}
;

//* Instrucción Guard
guardstmt returns [interfaces.Instruction guardinstr]
: GUARD expr ELSE LLAVEIZQ block LLAVEDER { $guardinstr = instructions.NewGuard($GUARD.line, $GUARD.pos, $expr.e, $block.blk) }
;

transferstmt returns [interfaces.Instruction trns]
: BREAK { $trns = instructions.NewBreak($BREAK.line, $BREAK.pos) }
| CONTINUE { $trns = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos) }
| RETURN { $trns = instructions.NewReturn($RETURN.line, $RETURN.pos, nil) }
| RETURN expr { $trns = instructions.NewReturn($RETURN.line, $RETURN.pos, $expr.e) }
;

/* rangestmt 
: expr PUNTO expr */

//* Gramatica para Expresiones
expr returns [interfaces.Expression e]
//Agregar para menos unario
: op=SUB right=expr { $e = expressions.NewOperation($right.start.GetLine(), $right.start.GetColumn(), $right.e, "unario", expressions.NewPrimitive($op.line, $op.pos, -1, environment.INTEGER)) }
| left=expr op=(MUL|DIV|MOD) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
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
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
| ID { $e = expressions.NewVar($ID.line, $ID.pos, $ID.text) }
;

/* primitive [interfaces.Expression p]  
: NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $p = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $p = expressions.NewPrimitive($STRING.line, $STRING.pos,str[1:len(str)-1],environment.STR)
    }                        
| TRU { $p = expressions.NewPrimitive($TRU.line, $TRU.pos,true,environment.BOOLEAN) }
| FAL { $p = expressions.NewPrimitive($FAL.line, $FAL.pos,false,environment.BOOLEAN) }
; */