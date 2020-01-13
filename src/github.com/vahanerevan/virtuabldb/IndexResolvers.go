package virtuabldb

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

type IndexResolver struct {
	Less      func(item, than interface{}) bool
	Transform func(item interface{}) interface{}
	OnNull    func() interface{}
}
type dataTypeResolver struct {
	Resolvers map[DataType]IndexResolver
}

func (dte *dataTypeResolver) Less(dataType DataType, item, than interface{}) bool {
	resolver, ok := dte.Resolvers[dataType]
	if false == ok {
		panic(fmt.Sprintf("Resolver for data type %v not found", dataType))
	}
	return resolver.Less(item, than)
}

func (dte *dataTypeResolver) Transform(dataType DataType, item interface{}) interface{} {
	resolver, ok := dte.Resolvers[dataType]
	if false == ok {
		panic(fmt.Sprintf("Resolver for data type %v not found", dataType))
	}
	if nil == resolver.Transform {
		return item
	}
	return resolver.Transform(item)
}

func (dte *dataTypeResolver) OnNull(dataType DataType) interface{} {
	resolver, ok := dte.Resolvers[dataType]
	if false == ok {
		panic(fmt.Sprintf("Resolver for data type %v not found", dataType))
	}
	return resolver.OnNull()
}

var IndexResolvers dataTypeResolver

func init() {
	IndexResolvers = dataTypeResolver{Resolvers: make(map[DataType]IndexResolver)}

	DefineIndexResolver(TYPE_INT, IndexResolver{
		Less: func(item, than interface{}) bool {
			return item.(int) < than.(int)
		},
		Transform: func(item interface{}) interface{} {
			result, err := cast.ToIntE(item)
			if nil != err {
				panic(err.Error())
			}
			return result
		},
	})

	DefineIndexResolver(TYPE_STRING, IndexResolver{
		Less: func(item, than interface{}) bool {
			return strings.Compare(item.(string), than.(string)) == -1
		},
	})

	DefineIndexResolver(TYPE_BOOL, IndexResolver{
		Less: func(item, than interface{}) bool {
			return (item.(int8)) < (than.(int8))
		},
		Transform: func(item interface{}) interface{} {
			return b2i(item.(bool))
		},
	})

	DefineIndexResolver(TYPE_FLOAT32, IndexResolver{
		Less: func(item, than interface{}) bool {
			return item.(float32) < than.(float32)
		},
		Transform: func(item interface{}) interface{} {
			return cast.ToFloat32(item)
		},
	})

	DefineIndexResolver(TYPE_FLOAT64, IndexResolver{
		Less: func(item, than interface{}) bool {
			return item.(float64) < than.(float64)
		},
		Transform: func(item interface{}) interface{} {
			return cast.ToFloat64(item)
		},
	})

}

func DefineIndexResolver(indexType DataType, resolver IndexResolver) {
	IndexResolvers.Resolvers[indexType] = resolver
}
