package geodb

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type Context struct {
	Domain string
	IP     string
}

type Expr interface {
	Eval(ctx *Context) bool
}

type GeoSiteExpr struct {
	Matcher []*DomainMatcher
}

func (e *GeoSiteExpr) Eval(ctx *Context) bool {
	for _, matcher := range e.Matcher {
		if matcher.Match(ctx.Domain) {
			return true
		}
	}
	return false
}

func (g GeoSiteExpr) MarshalYAML() (any, error) {
	var names []string
	for _, m := range g.Matcher {
		if m != nil {
			names = append(names, fmt.Sprintf("%s:%d", m.Code, len(m.Params)))
		}
	}
	return names, nil
}

type FullExpr struct {
	Values []string
}

func (e *FullExpr) Eval(ctx *Context) bool {
	for _, v := range e.Values {
		if strings.EqualFold(ctx.Domain, v) {
			return true
		}
	}
	return false
}

type KeywordExpr struct {
	Values []string
}

func (e *KeywordExpr) Eval(ctx *Context) bool {
	for _, v := range e.Values {
		if strings.Contains(ctx.Domain, v) {
			return true
		}
	}
	return false
}

type SuffixExpr struct {
	Values []string
}

func (e *SuffixExpr) Eval(ctx *Context) bool {
	for _, v := range e.Values {
		if ctx.Domain == v || strings.HasSuffix(ctx.Domain, "."+strings.TrimPrefix(v, ".")) {
			return true
		}
	}
	return false
}

type RegexExpr struct {
	Regexes []*regexp.Regexp
}

func (e *RegexExpr) Eval(ctx *Context) bool {
	for _, re := range e.Regexes {
		if re.MatchString(ctx.Domain) {
			return true
		}
	}
	return false
}

type NotExpr struct {
	Expr Expr
}

func (e *NotExpr) Eval(ctx *Context) bool {
	return !e.Expr.Eval(ctx)
}

type AndExpr struct {
	Left, Right Expr
}

func (e *AndExpr) Eval(ctx *Context) bool {
	return e.Left.Eval(ctx) && e.Right.Eval(ctx)
}

type OrExpr struct {
	Left, Right Expr
}

func (e *OrExpr) Eval(ctx *Context) bool {
	fmt.Printf("domain: %s, or: %v, %v\n", ctx.Domain, e.Left, e.Right)
	return e.Left.Eval(ctx) || e.Right.Eval(ctx)
}

type Rule struct {
	Action string
	Cond   Expr
}

func (r *Rule) Match(ctx *Context) (string, bool) {
	if r.Cond.Eval(ctx) {
		return r.Action, true
	}
	return "", false
}

func DebugExpr(expr Expr) {
	message, err := yaml.Marshal(expr)
	if err != nil {
		return
	}
	log.Printf("\nexpr: %T\n%s", expr, string(message))
}
