// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

package openapi

import (
	"cuelang.org/go/internal/core/adt"
	"cuelang.org/go/internal/pkg"
)

func init() {
	pkg.Register("openapi", p)
}

var _ = adt.TopKind // in case the adt package isn't used

var p = &pkg.Package{
	Native: []*pkg.Builtin{{
		Name: "Gen",
		Params: []pkg.Param{
			{Kind: adt.TopKind},
		},
		Result: adt.StringKind,
		Func: func(c *pkg.CallCtxt) {
			s := c.Schema(0)
			if c.Do() {
				c.Ret, c.Err = Gen(s)
			}
		},
	}},
}
