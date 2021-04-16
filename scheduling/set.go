package scheduling

import "sync"

type HashSet struct {
	m map[interface{}]interface{}
}

func createHashSet(cap int) *HashSet {
	return &HashSet{
		m: make(map[interface{}]interface{}, cap),
	}
}

func (s *HashSet) Has(key interface{}) bool {
	_, exist := s.m[key]
	return exist
}

func (s *HashSet) Insert(key interface{}) {
	s.m[key] = 0
}

func (s *HashSet) Del(key interface{}) {
	delete(s.m, key)
}

func (s *HashSet) Free() {
	s.m = nil
}

type sliceSet struct {
	m []interface{}

	pool *sync.Pool
}

func createSliceSet(pool *sync.Pool) *sliceSet {
	return &sliceSet{
		m:    pool.Get().([]interface{})[:0],
		pool: pool,
	}
}

func (s *sliceSet) Has(key interface{}) bool {
	for i := range s.m {
		if s.m[i] == key {
			return true
		}
	}
	return false
}

func (s *sliceSet) Insert(key interface{}) {
	s.m = append(s.m, key)
}

func (s *sliceSet) Del(key interface{}) {
	for i := range s.m {
		if s.m[i] == key {
			s.m = append(s.m[:i], s.m[i+1:]...)
			break
		}
	}
	return
}

func (s *sliceSet) Free() {
	s.pool.Put(s.m)
	s.m = nil
}
