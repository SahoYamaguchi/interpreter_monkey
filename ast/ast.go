package ast

import ( "github.com/SahoYamaguchi/interpreter_monkey/token"
  "bytes"
)

type Node interface{
  TokenLiteral() string
  String() string
}

type Statement interface{
  Node
  statementNode()
}

type Expression interface{
  Node
  expressionNode()
}

type Program struct{
  Statements []Statement
}

func (p *Program) TokenLiteral() string{
  if len(p.Statements) > 0{
    return p.Statements[0].TokenLiteral()
  }else{
    return ""
  }
}

func (p *Program) String() string{
  var out bytes.Buffer

  for _,s := range p.Statements{
    out.WriteString(s.String())
  }

  return out.String()
}

type LetStatement struct{
  Token token.Token //token.LET
  Name *Identifier
  Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct{
  Token token.Token //token.IDENT
  Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct{
  Token token.Token //'return'
  ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct{
  Token token.Token //式の最初のトークン
  Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string{ return es.Token.Literal }

type IntegerLiteral struct{
  Token token.Token
  Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {return il.Token.Literal}
func (il *IntegerLiteral) String() string {return il.Token.Literal}

func(ls *LetStatement) String() string{
  var out bytes.Buffer

  out.WriteString(ls.TokenLiteral() + " ")
  out.WriteString(ls.Name.String())
  out.WriteString(" = ")

  if ls.Value != nil{
    out.WriteString(ls.Value.String())
  }

  out.WriteString(";")

  return out.String()
}

func(rs *ReturnStatement) String() string{
  var out bytes.Buffer

  out.WriteString(rs.TokenLiteral() + " ")

  if rs.ReturnValue != nil{
    out.WriteString(rs.ReturnValue.String())
  }

  out.WriteString(";")

  return out.String()
}

func (es *ExpressionStatement) String() string{
  if es.Expression != nil{
    return es.Expression.String()
  }
  return ""
}

func (i *Identifier) String() string { return i.Value }

type PlusExpression struct{
  Token token.Token //式の最初のトークン
  LeftExp Expression
  RightExp Expression
}

func (pe *PlusExpression) statementNode() {}
func (pe *PlusExpression) TokenLiteral() string{ return pe.Token.Literal }

func (pe *PlusExpression) String() string{
  return fmt.Sprintf("%V + %V", LeftExp, RightExp)
}
