grammar ParserSea;


progAbilities: (decl | expr | assign | funcDecl);

prog: progAbilities+ EOF
    ;

decl: typesKeyword ID '=' (types_of_tokens | expr);

block: '{' progAbilities* '}';

assign: ID '=' (types_of_tokens | expr);

flags : (PRIVATE | PUBLIC);

types_of_tokens: (NUM | STRING | FLOAT);
typesKeyword: INT_TYPE | STRING_TYPE | FLOAT_TYPE;
param: typesKeyword ID;

/* function decleration */
paramList: '(' (param (',' param)*)? ')';
returnList: ('(' typesKeyword (',' typesKeyword)* ')')? ;
funcDecl:
        flags*?
        DEFINE
        ID
        paramList
        returnList
        block
        ;




/* ANTLR resoleves ambiguities in favor of the alternative given first*/
expr: ID '(' (expr (',' expr)*)? ')'
    | expr PLUS expr
    | expr MUL expr
    | ID
    | types_of_tokens
    ;

/* key words */
RETURN: 'return';
DEFINE: 'def';
INT_TYPE : 'int';
STRING_TYPE : 'string';
FLOAT_TYPE : 'float';
PUBLIC : 'pub';
PRIVATE : 'pri';
PLUS : '+';
MUL : '*';

/* tokens */


ID : [a-z][a-zA-Z0-9_]*; // identifire
FLOAT : '-'? [1-9]+ '.'  [1-9]+;
NUM : '0' | '-'?[1-9][0-9]*;
COMMENT : '#' ~[\r\n]* -> skip;
STRING : '"' ~[\r\n"]* '"'
       | '\'' ~[\r\n']* '\''
       ;

WS : [ \t\r\n]+ -> skip;
