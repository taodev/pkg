package geodb

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"regexp"
	"strings"
)

type Helper struct {
	File         string
	Line, Column int
}

func WrapError(err error, filename string, line, column int) error {
	// 打印 err 的类型
	switch codeErr := err.(type) {
	case scanner.Error:
		codeErr.Pos.Filename = filename
		codeErr.Pos.Line += line
		codeErr.Pos.Column += column
		err = codeErr
	case scanner.ErrorList:
		for _, e := range codeErr {
			e.Pos.Filename = filename
			e.Pos.Line += line
			e.Pos.Column += column
		}
		err = codeErr
	default:
		err = fmt.Errorf("%s:%d:%d: %w", filename, line, column, err)
	}
	return err
}

func LoadRule(code string) (*Rule, error) {
	expr, err := parser.ParseExpr(code)
	if err != nil {
		return nil, err
	}
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return nil, fmt.Errorf("must like action(...), code: %s", code)
	}
	action, ok := call.Fun.(*ast.Ident)
	if !ok {
		// 动作必须是标识符
		return nil, fmt.Errorf("action must be ident")
	}
	if len(call.Args) == 0 {
		// 必须包含一个或多个条件参数
		return nil, fmt.Errorf("must have at least one condition")
	}
	cond, err := buildExpr(call.Args[0])
	if err != nil {
		return nil, err
	}
	for _, arg := range call.Args[1:] {
		right, err := buildExpr(arg)
		if err != nil {
			return nil, err
		}
		cond = &OrExpr{Left: cond, Right: right}
	}

	return &Rule{Action: action.Name, Cond: cond}, nil
}

func buildExpr(node ast.Expr) (Expr, error) {
	switch v := node.(type) {
	case *ast.CallExpr:
		fn, ok := v.Fun.(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf("不支持的函数表达式")
		}
		switch fn.Name {
		case DomainKeyGeoSite:
			return buildTagExpr(v.Args, func(tags []string) (Expr, error) {
				var expr GeoSiteExpr
				for _, tag := range tags {
					matcher, err := Site("geosite.dat", tag)
					if err != nil {
						return nil, fmt.Errorf("geosite:%s - %w", tag, err)
					}
					expr.Matcher = append(expr.Matcher, matcher.(*DomainMatcher))
				}
				return &expr, nil
			})
		case DomainKeyFull:
			return buildStringExpr(v.Args, func(val []string) Expr {
				return &FullExpr{Values: val}
			})
		case DomainKeyKeyword:
			return buildStringExpr(v.Args, func(val []string) Expr {
				return &KeywordExpr{Values: val}
			})
		case DomainKeySuffix:
			return buildStringExpr(v.Args, func(val []string) Expr {
				return &SuffixExpr{Values: val}
			})
		case DomainKeyRegex:
			return buildStringExpr(v.Args, func(val []string) Expr {
				var regexes []*regexp.Regexp
				for _, v := range val {
					regexes = append(regexes, regexp.MustCompile(v))
				}
				return &RegexExpr{Regexes: regexes}
			})
		default:
			// 不支持的函数表达式
			return nil, fmt.Errorf("unsupported function expression: %s", fn.Name)
		}
	case *ast.UnaryExpr:
		if v.Op == token.NOT {
			expr, err := buildExpr(v.X)
			if err != nil {
				return nil, err
			}
			return &NotExpr{Expr: expr}, nil
		}
	case *ast.BinaryExpr:
		left, err := buildExpr(v.X)
		if err != nil {
			return nil, err
		}
		right, err := buildExpr(v.Y)
		if err != nil {
			return nil, err
		}
		switch v.Op {
		case token.LAND:
			return &AndExpr{Left: left, Right: right}, nil
		case token.LOR:
			return &OrExpr{Left: left, Right: right}, nil
		default:
			// 不支持的逻辑操作符
			return nil, fmt.Errorf("unsupported logical operator: %s", v.Op.String())
		}
	case *ast.ParenExpr:
		return buildExpr(v.X)
	}
	// 无法识别表达式类型
	return nil, fmt.Errorf("unsupported expression type: %T", node)
}

func buildTagExpr(args []ast.Expr, f func([]string) (Expr, error)) (Expr, error) {
	var tags []string
	for _, arg := range args {
		s, err := extractString(arg)
		if err != nil {
			return nil, err
		}
		tags = append(tags, strings.ToLower(s))
	}
	return f(tags)
}

func buildStringExpr(args []ast.Expr, f func([]string) Expr) (Expr, error) {
	if len(args) < 1 {
		// 至少有一个参数
		return nil, fmt.Errorf("must have at least one string argument")
	}
	var tags []string
	for _, arg := range args {
		s, err := extractString(arg)
		if err != nil {
			return nil, err
		}
		tags = append(tags, s)
	}
	return f(tags), nil
}

func extractString(expr ast.Expr) (string, error) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		// 必须是字符串字面量
		return "", fmt.Errorf("must be string literal")
	}
	return strings.ToLower(strings.Trim(lit.Value, `"`)), nil
}
