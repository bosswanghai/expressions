package expressions

import (
	"strconv"

	"github.com/cyberfox/expressions/parser"
	"github.com/wxio/antlr4/runtime/Go/antlr"
)

type ExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

// var _ parser.StartContextVisitor = &ExprVisitor{}
// var _ parser.CodelineContextVisitor = &ExprVisitor{}
var _ parser.AddSubExprContextVisitor = &ExprVisitor{}
var _ parser.ParenExprContextVisitor = &ExprVisitor{}

// var _ parser.LiteralExprContextVisitor = &ExprVisitor{}
// var _ parser.UnaryExprContextVisitor = &ExprVisitor{}
// var _ parser.UnaryContextVisitor = &ExprVisitor{}
var _ parser.IntLiteralContextVisitor = &ExprVisitor{}

func (v *ExprVisitor) VisitIntLiteral(ctx parser.IIntLiteralContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	r, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		return nil
	}
	return r
}

func (v *ExprVisitor) VisitParenExpr(ctx parser.IParenExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = ctx.GetE().Visit(v)
	return
}

func (v *ExprVisitor) VisitAddSubExpr(ctx parser.IAddSubExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	op := ctx.GetOp().GetText()
	left := ctx.GetA().Visit(delegate)
	right := ctx.GetB().Visit(delegate)

	switch left.(type) {
	case int64:
		if op == "+" {
			result = left.(int64) + right.(int64)
		} else {
			result = left.(int64) - right.(int64)
		}
		// default:
		// 	return nil
	}
	return
}

func (v *ExprVisitor) VisitNegateExpr(ctx parser.INegateExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	val := ctx.GetVal().Visit(v)
	result = -val.(int64)
	return
}

func (v *ExprVisitor) VisitInvertExpr(ctx parser.IInvertExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	val := ctx.GetVal().Visit(v)
	result = ^val.(int64)
	return
}

func NewEvaluator() *ExprVisitor {
	visitor := new(ExprVisitor)
	return visitor
}
