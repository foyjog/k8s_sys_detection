grammar engine;

start: rule;

rule: (singleExpression|mulExpression) logicOper (singleExpression|mulExpression)
| singleExpression
| mulExpression;

mulExpression:'('singleExpression logicOper singleExpression')';

singleExpression: singleExpression logicOper singleExpression
| variable compareEqual variable
| variable compareIn variable;

variable: VAR_DOT | INTEGER| STRING;
compareEqual: EQUAL|NOT_EQUAL;
compareIn: compareEqual | IN;
logicOper: (AND|OR);

EQUAL: '==';
NOT_EQUAL: '!=';
AND: 'and';
OR: 'or';
INTEGER: '-'?[0-9]+;
IN: 'in';
STRING: '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
VAR_DOT: VAR | VAR DOT VAR_DOT;
VAR: [a-zA-Z0-9]+;
DOT: '.';
SL_COMMENT: '//' .*? '\n' -> skip ;
WS  :   [ \t\n\r]+ -> skip ;