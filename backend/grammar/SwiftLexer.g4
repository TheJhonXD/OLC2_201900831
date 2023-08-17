lexer grammar SwiftLexer;

// --------------- Tokens
// types
INT: 'Int';
FLOAT: 'Float';
BOOL: 'Bool';
STR: 'String';
CHARACTER: 'Character';

// reserved words
TRU: 'true';
FAL: 'false';
NIL: 'nil';
VAR: 'var';
LET: 'let';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
WHILE: 'while';
FOR: 'for';
IN: 'in';
GUARD: 'guard';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
STRUCT: 'struct';
SELF: 'self';
MUTATING: 'mutating';
FUNC: 'func';

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
CHAR: '(\\(["\'\\bfnrt]|u[0-9A-Fa-f]{4})|[^\\\'])';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

// symbols

DIF:      '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:          '=';
MAY_IG:     '>=';
MEN_IG:     '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
Q_MARK:         '?';
ARROW:         '->';
COMA:           ',';
PUNTO:          '.';
COLON:      ':';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;