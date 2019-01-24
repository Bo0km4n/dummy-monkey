package evaluator

import (
	"fmt"

	"github.com/Bo0km4n/dummy-monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: _builtinLen,
	},
	"puts": &object.Builtin{
		Fn: _builtinPuts,
	},
	"first": &object.Builtin{
		Fn: _buildinFirst,
	},
}

func _builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func _builtinPuts(args ...object.Object) object.Object {
	in := []interface{}{}

	for _, v := range args {
		switch v := v.(type) {
		case *object.String:
			in = append(in, v.Value)
		case *object.Integer:
			in = append(in, v.Value)
		case *object.Boolean:
			in = append(in, v.Value)
		case *object.Error:
			in = append(in, v.Message)
		default:
			return newError("argument to `puts` not supported, got %s", args[0].Type())
		}
	}
	fmt.Println(in...)
	return &object.Integer{Value: int64(len(args))}
}

func _buildinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got=%s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}
	return NULL
}
