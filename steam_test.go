package steams

import (
	"reflect"
	"testing"
)

func TestPosta(t *testing.T) {
	list := List[int]{1, 2, 3, 4, 5}
	res := list.Skip(1).Filter(filtro)
    res2 := CollectSteamToSteam2(res, func(v int) int { return v }, func(v int) int { return v + 1 })
	t.Logf("Posta %v\n", reflect.TypeOf(res.Collect()))
	t.Logf("Posta %v\n", res.Collect())
	t.Logf("Result %v\n", res2)

	m := Map[int, string]{
		1: "uno",
		2: "dos",
		3: "tres",
	}
	r := m.Filter(func(k int, v string) bool { return k%2 == 0 }).Collect()
	t.Logf("Map %v\n", r)
}

func filtro(v int) bool {
	return v%2 == 0
}
