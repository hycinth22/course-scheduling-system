package scheduling

import "sync"

//type HashSet struct {
//	m map[interface{}]interface{}
//}
//
//func createHashSet(cap int) *HashSet {
//	return &HashSet{
//		m: make(map[interface{}]interface{}, cap),
//	}
//}
//
//func (s *HashSet) Has(key interface{}) bool {
//	_, exist := s.m[key]
//	return exist
//}
//
//func (s *HashSet) Insert(key interface{}) {
//	s.m[key] = 0
//}
//
//func (s *HashSet) Del(key interface{}) {
//	delete(s.m, key)
//}
//
//func (s *HashSet) Free() {
//	s.m = nil
//}

type IEqual interface {
	Equal(IEqual) bool
}

type sliceSet struct {
	m []IEqual

	pool *sync.Pool
}

func createSliceSet(pool *sync.Pool) *sliceSet {
	return &sliceSet{
		m:    pool.Get().([]IEqual)[:0],
		pool: pool,
	}
}

func (s *sliceSet) Has(key IEqual) bool {
	for i := range s.m {
		if s.m[i].Equal(key) {
			return true
		}
	}
	return false
}

func (s *sliceSet) Get(key IEqual) (IEqual, bool) {
	for i := range s.m {
		if s.m[i].Equal(key) {
			return s.m[i], true
		}
	}
	return nil, false
}

func (s *sliceSet) Insert(key IEqual) {
	s.m = append(s.m, key)
}

func (s *sliceSet) Del(key IEqual) {
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
