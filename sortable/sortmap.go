package sortable

type SortedMap struct {
	Map      []KeyValuePair
	Property int
	Mode     int
}

type KeyValuePair struct {
	Key   interface{}
	Value interface{}
}

const (
	SortByKey = iota + 1
	SortByValue
)

const (
	ASC = iota + 1
	DESC
)

func SortableMap(m map[interface{}]interface{}, property int, mode int) SortedMap {
	var sm SortedMap
	sm.Map = make([]KeyValuePair, 0, len(m))
	for k, v := range m {
		sm.Map = append(sm.Map, KeyValuePair{k, v})
	}
	sm.Property = property
	sm.Mode = mode
	return sm
}

func (sm SortedMap) Len() int {
	return len(sm.Map)
}

func (sm SortedMap) Less(i, j int) bool {
	var flag bool
	switch sm.Property {
	case SortByKey:
		flag = compare(sm.Map[i].Key, sm.Map[j].Key)
	case SortByValue:
		flag = compare(sm.Map[i].Value, sm.Map[j].Value)
	default:
		panic("sortBy must be set!")
	}
	switch sm.Mode {
	case ASC:
		return flag
	case DESC:
		return !flag
	default:
		panic("mode must be set!")
	}
}

func (sm SortedMap) Swap(i, j int) {
	sm.Map[i], sm.Map[j] = sm.Map[j], sm.Map[i]
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
