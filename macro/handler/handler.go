// Package handler is the highest level module of the macro package which makes use the rest of the macro package,
// it is mainly used, internally, by the router package.
package handler

import (
	"github.com/radiantrfid/iris/context"
	"github.com/radiantrfid/iris/macro"
)

// CanMakeHandler reports whether a macro template needs a special macro's evaluator handler to be validated
// before procceed to the next handler(s).
// If the template does not contain any dynamic attributes and a special handler is NOT required
// then it returns false.
func CanMakeHandler(tmpl macro.Template) (needsMacroHandler bool) {
	if len(tmpl.Params) == 0 {
		return
	}

	// check if we have params like: {name:string} or {name} or {anything:path} without else keyword or any functions used inside these params.
	// 1. if we don't have, then we don't need to add a handler before the main route's handler (as I said, no performance if macro is not really used)
	// 2. if we don't have any named params then we don't need a handler too.
	for _, p := range tmpl.Params {
		if p.CanEval() {
			// if at least one needs it, then create the handler.
			needsMacroHandler = true
			break
		}
	}

	return
}

// MakeHandler creates and returns a handler from a macro template, the handler evaluates each of the parameters if necessary at all.
// If the template does not contain any dynamic attributes and a special handler is NOT required
// then it returns a nil handler.
func MakeHandler(tmpl macro.Template) context.Handler {
	filter := MakeFilter(tmpl)

	return func(ctx context.Context) {
		if !filter(ctx) {
			ctx.StopExecution()
			return
		}

		// if all passed, just continue.
		ctx.Next()
	}
}

// MakeFilter returns a Filter which reports whether a specific macro template
// and its parameters pass the serve-time validation.
func MakeFilter(tmpl macro.Template) context.Filter {
	if !CanMakeHandler(tmpl) {
		return nil
	}

	return func(ctx context.Context) bool {
		for _, p := range tmpl.Params {
			if !p.CanEval() {
				continue // allow.
			}

			// 07-29-2019
			// changed to retrieve by param index in order to support
			// different parameter names for routes with
			// different param types (and probably different param names i.e {name:string}, {id:uint64})
			// in the exact same path pattern.
			//
			// Same parameter names are not allowed, different param types in the same path
			// should have different name e.g. {name} {id:uint64};
			// something like {name} and {name:uint64}
			// is bad API design and we do NOT allow it by-design.
			entry, found := ctx.Params().Store.GetEntryAt(p.Index)
			if !found {
				// should never happen.
				return false
			}

			if !p.Eval(entry.String(), &ctx.Params().Store) {
				ctx.StatusCode(p.ErrCode)
				return false
			}
		}

		return true
	}
}
