package sortable

type SortedMap []KeyValuePair

type KeyValuePair struct {
	Key   interface{}
	Value interface{}
}

func SortableMap(m map[interface{}]interface{}) SortedMap {
	sm := make(SortedMap, 0, len(m))
	for k, v := range m {
		sm = append(sm, KeyValuePair{k, v})
	}
	return sm
}
