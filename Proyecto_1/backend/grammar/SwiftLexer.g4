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
RANGEPTS: '...';
GUARD: 'guard';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
STRUCT: 'struct';
SELF: 'self';
MUTATING: 'mutating';
FUNC: 'func';
APPEND: 'append';
REMOVELAST: 'removeLast';
REMOVE: 'remove';
EMPTY: 'isEmpty';
COUNT: 'count';
INOUT: 'inout';
CAME: '_';

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
CHAR: '"'(.)'"';
STRING: '"'~["]*'"';
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
MOD:            '%';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
CORCHIZQ:       '[';
CORCHDER:       ']';
Q_MARK:         '?';
ARROW:         '->';
COMA:           ',';
PUNTO:          '.';
COLON:      ':';
AMP:            '&';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;