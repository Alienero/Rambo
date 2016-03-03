package prefixtree

import "testing"

func TestPrefixTree(t *testing.T) {
	pt := NewPrefixTree()
	key := []interface{}{"1", "2"}
	pt.Add(key, "3")
	t.Log(pt.Get(key).(string))
	pt.Update(key, "666")
	t.Log(pt.Get(key).(string))
	if rest := pt.Get(key).(string); rest != "666" {
		t.Error("update fail")
	}
	pt.Del(key)
	if pt.Get(key) != nil || !pt.IsEmpty() {
		t.Error("del error")
	}
}
