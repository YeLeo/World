package util

type Property int

const (
	SortByKey Property = iota + 1
	SortByValue
)

type Mode int

const (
	ASC Mode = iota + 1
	DESC
)

type SortableMap struct {
	KeyValuePairs []KeyValuePair
	SortProperty  Property
	SortMode      Mode
}

type KeyValuePair struct {
	Key   interface{}
	Value interface{}
}

func CreateSM(m map[interface{}]interface{}, property Property, mode Mode) SortableMap {
	var sm SortableMap
	sm.KeyValuePairs = make([]KeyValuePair, 0, len(m))
	for k, v := range m {
		sm.KeyValuePairs = append(sm.KeyValuePairs, KeyValuePair{k, v})
	}
	sm.SortProperty = property
	sm.SortMode = mode
	return sm
}

func (sm SortableMap) Len() int {
	return len(sm.KeyValuePairs)
}

func (sm SortableMap) Less(i, j int) bool {
	var flag bool
	switch sm.SortProperty {
	case SortByKey:
		flag = compare(sm.KeyValuePairs[i].Key, sm.KeyValuePairs[j].Key)
	case SortByValue:
		flag = compare(sm.KeyValuePairs[i].Value, sm.KeyValuePairs[j].Value)
	default:
		panic("sortBy must be set!")
	}
	switch sm.SortMode {
	case ASC:
		return flag
	case DESC:
		return !flag
	default:
		panic("mode must be set!")
	}
}

func (sm SortableMap) Swap(i, j int) {
	sm.KeyValuePairs[i], sm.KeyValuePairs[j] = sm.KeyValuePairs[j], sm.KeyValuePairs[i]
}

func compare(x, y interface{}) bool {
	switch x.(type) {
	case string:
		return x.(string) < y.(string)
	case int:
		return x.(int) < y.(int)
	case float32:
		return x.(float32) < y.(float32)
	case float64:
		return x.(float64) < y.(float64)
	default:
		panic("the type can't compare!")
	}
}
