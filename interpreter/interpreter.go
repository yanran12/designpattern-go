package interpreter

import "fmt"

type Context struct {
	VariableMap map[*VariableExpression]int
}

func (ctx *Context) Lookup(Variable *VariableExpression) int {
	return ctx.VariableMap[Variable]
}

type Expression interface {
	Interpret(context *Context) int
}

// 终结符
type VariableExpression struct {
}

func (ve *VariableExpression) Interpret(context *Context) int {
	return context.Lookup(ve)
}

// 加
type AddExpression struct {
	A Expression
	B Expression
}

func (ae *AddExpression) Interpret(context *Context) int {
	return ae.A.Interpret(context) + ae.B.Interpret(context)
}

// 减
type SubtractExpression struct {
	A Expression
	B Expression
}

func (se *SubtractExpression) Interpret(context *Context) int {
	return se.A.Interpret(context) - se.B.Interpret(context)
}

func Run() {

	x := &VariableExpression{}
	y := &VariableExpression{}

	ctx := &Context{VariableMap: map[*VariableExpression]int{x: 3, y: 7}}

	result := &SubtractExpression{
		A: x,
		B: &AddExpression{
			A: y,
			B: x,
		},
	}

	fmt.Println(result.Interpret(ctx))

}
