package join

import (
	"reflect"
	"sort"
)

func JoinSort(data [][2]interface{}, key string) {
	sort.Sort(sortableJoin{key, data})
}

type sortableJoin struct {
	sort string
	data [][2]interface{}
}

func (d sortableJoin) Len() int { return len(d.data) }

func (d sortableJoin) Less(i, j int) bool {
	iv := d.value(i, 0, d.sort)
	jv := d.value(j, 0, d.sort)
	if iv == nil && jv == nil {
		iv = d.value(i, 1, d.sort)
		jv = d.value(j, 1, d.sort)
	}
	return d.less(iv, jv)
}

func (d sortableJoin) less(iv, jv interface{}) bool {
	if iv != nil && jv != nil {
		switch v := iv.(type) {
		case int:
			return v < jv.(int)
		default:
			return false
		}
	} else if iv == nil {
		return false
	} else if jv == nil {
		return true
	} else {
		return false
	}
}

func (d sortableJoin) value(index, which int, name string) interface{} {
	o := d.data[index][which]
	if o == nil {
		return nil
	}
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Struct {
		return v.FieldByName(name).Interface()
	} else {
		return v.MapIndex(reflect.ValueOf(name)).Interface()
	}
}

func (d sortableJoin) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}
