package scheduling

type hashSet struct {
	m map[interface{}]interface{}
}

func createHashSet(cap int) *hashSet {
	return &hashSet{
		m: make(map[interface{}]interface{}, cap),
	}
}

func (s *hashSet) Has(key interface{}) bool {
	_, exist := s.m[key]
	return exist
}

func (s *hashSet) Insert(key interface{}) {
	s.m[key] = 0
}

func (s *hashSet) Del(key interface{}) {
	delete(s.m, key)
}

func (s *hashSet) Free() {
	s.m = nil
}

type pair struct {
	first, second interface{}
}
