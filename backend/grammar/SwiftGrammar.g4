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
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
;

ifstmt returns [interfaces.Instruction ifinstr]
: IF  expr LLAVEIZQ block LLAVEDER { $ifinstr = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk) }
| IF expr LLAVEIZQ block LLAVEDER ELSE LLAVEIZQ block LLAVEDER
| IF expr LLAVEIZQ block LLAVEDER ELSE ifstmt
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