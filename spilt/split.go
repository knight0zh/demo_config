package split

import (
	"fmt"
	"sync"
)

var (
	dbm *DatabaseMap
	one sync.Once
)

type (
	Type1 map[string]interface{}
	Type2 map[string]interface{}
	Type3 map[string]interface{}
	Type4 map[string]interface{}
	Type5 map[string]interface{}

	DatabaseMap struct {
		Lock     sync.Mutex
		SplitMap SplitMap
	}
	SplitMap struct {
		T1 Type1
		T2 Type2
		T3 Type3
		T4 Type4
		T5 Type5
	}
)

func (t Type1) Split(tb, s string) string {
	return fmt.Sprintf("%s_0%s", tb, s[len(s)-1:])
}

func (t Type2) Split(tb, s string) string {
	a := s[len(s)-1:]
	b := []byte(a)
	return fmt.Sprintf("%s_0%d", tb, b[len(b)-1]%10)
}

func (t Type3) Split(tb, s string) string {
	a := s[len(s)-4 : len(s)-3]
	b := []byte(a)
	return fmt.Sprintf("%s_0%d", tb, b[len(b)-1]%10)
}

func (t Type4) Split(tb, s string) string {
	a := s[:1]
	b := []byte(a)
	return fmt.Sprintf("%s_0%d", tb, b[len(b)-1]%10)
}

func (t Type5) Split(tb, s string) string {
	a := s[len(s)-3 : len(s)-2]
	b := []byte(a)
	return fmt.Sprintf("%s_0%d", tb, b[len(b)-1]%10)
}

func GetDbMap() *DatabaseMap {
	one.Do(func() {
		var m SplitMap
		m = SplitMap{
			T1: GetType1(),
			T2: GetType2(),
			T3: GetType3(),
			T4: GetType4(),
			T5: GetType5(),
		}

		dbm = &DatabaseMap{
			SplitMap: m,
		}
	})
	return dbm
}

func (d *DatabaseMap) Split(db, tb, split string) string {
	if split == "" {
		return tb
	}
	d.Lock.Lock()
	defer d.Lock.Unlock()

	if v, ok := d.SplitMap.T1[db]; ok {
		if _, ok := v.(map[string]interface{})[tb]; ok {
			return db + "." + d.SplitMap.T1.Split(tb, split)
		}
	}
	if v, ok := d.SplitMap.T2[db]; ok {
		if _, ok := v.(map[string]interface{})[tb]; ok {
			return db + "." + d.SplitMap.T2.Split(tb, split)
		}
	}
	if v, ok := d.SplitMap.T3[db]; ok {
		if _, ok := v.(map[string]interface{})[tb]; ok {
			return db + "." + d.SplitMap.T3.Split(tb, split)
		}
	}
	if v, ok := d.SplitMap.T4[db]; ok {
		if _, ok := v.(map[string]interface{})[tb]; ok {
			return db + "." + d.SplitMap.T4.Split(tb, split)
		}
	}
	if v, ok := d.SplitMap.T5[db]; ok {
		if _, ok := v.(map[string]interface{})[tb]; ok {
			return db + "." + d.SplitMap.T5.Split(tb, split)
		}
	}

	return db + "." + tb
}
