package yiigo_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/iiinsomnia/yiigo"
)

var dataInts = []int{1, 4, 7, 9, 0, 3, 5, 2, 7, 9, 8, 1, 6}
var uniqueInts = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var dataInt64s = []int64{1, 6, 4, 7, 9, 0, 3, 5, 2, 7, 9, 8, 4}
var uniqueInt64s = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var dataFloat64s = []float64{1.2, 4.1, 7.3, 9.5, 0.1, 3.2, 5.4, 2.3, 7.3, 9.5, 8.7, 1.2, 6.0}
var uniqueFloat64s = []float64{0.1, 1.2, 2.3, 3.2, 4.1, 5.4, 6.0, 7.3, 8.7, 9.5}

var dataStrings = []string{"hello", "world", "golang", "wechat", "alipay", "hello", "wechat"}
var uniqueStrings = []string{"alipay", "golang", "hello", "wechat", "world"}

func TestSortInt64s(t *testing.T) {
	data := dataInt64s
	a := yiigo.Int64Slice(data)
	yiigo.SortInt64s(a)

	if !sort.IsSorted(a) {
		t.Errorf("sorted %v", dataInt64s)
		t.Errorf("   got %v", data)
	}
}

func TestSearchInt64s(t *testing.T) {
	a := uniqueInt64s
	i := yiigo.SearchInt64s(a, 4)

	if i != 4 {
		t.Errorf("expected index 4; got %d", i)
	}
}

func TestUniqueInt(t *testing.T) {
	a := dataInts
	r := yiigo.UniqueInt(a)

	if !reflect.DeepEqual(r, uniqueInts) {
		t.Error("test UniqueInt failed")
	}
}

func TestUniqueInt64(t *testing.T) {
	a := dataInt64s
	r := yiigo.UniqueInt64(a)

	if !reflect.DeepEqual(r, uniqueInt64s) {
		t.Error("test UniqueInt64 failed")
	}
}

func TestUniqueFloat64(t *testing.T) {
	a := dataFloat64s
	r := yiigo.UniqueFloat64(a)

	if !reflect.DeepEqual(r, uniqueFloat64s) {
		t.Error("test UniqueFloat64 failed")
	}
}

func TestUniqueString(t *testing.T) {
	a := dataStrings
	r := yiigo.UniqueString(a)

	if !reflect.DeepEqual(r, uniqueStrings) {
		t.Error("test UniqueString failed")
	}
}

func TestInSliceInt(t *testing.T) {
	a := dataInts
	r := yiigo.InSliceInt(9, a)

	if !r {
		t.Error("test InSliceInt failed")
	}
}

func TestInSliceInt64(t *testing.T) {
	a := dataInt64s
	r := yiigo.InSliceInt64(9, a)

	if !r {
		t.Error("test InSliceInt64 failed")
	}
}

func TestInSliceFloat64(t *testing.T) {
	a := dataFloat64s
	r := yiigo.InSliceFloat64(9.5, a)

	if !r {
		t.Error("test InSliceFloat64 failed")
	}
}

func TestInSliceString(t *testing.T) {
	a := dataStrings
	r := yiigo.InSliceString("golang", a)

	if !r {
		t.Error("test InSliceString failed")
	}
}
