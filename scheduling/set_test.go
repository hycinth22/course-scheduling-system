package scheduling

import (
	"testing"
)

func Test_createHashSet(t *testing.T) {
	if got := createHashSet(1); got == nil {
		t.Errorf("createHashSet() = %v", got)
	}
}

/*
func Test_hashSet_Del(t *testing.T) {
	type fields struct {
		m map[interface{}]interface{}
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"del nil", fields{m: map[interface{}]interface{}{nil: 0}}, args{key: nil}},
		{"del int", fields{m: map[interface{}]interface{}{0: 0}}, args{key: 0}},
		{"del int", fields{m: map[interface{}]interface{}{111: 0}}, args{key: 111}},
		{"del string", fields{m: map[interface{}]interface{}{"0": 0}}, args{key: "0"}},
		{"del string", fields{m: map[interface{}]interface{}{"ttt": 0}}, args{key: "tttt"}},
		{"del struct", fields{m: map[interface{}]interface{}{
			pair{timespanInDay{321, 3}, 5}: 0,
		}}, args{
			key: pair{timespanInDay{321, 3}, 5},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &traversalHashSet{
				m: tt.fields.m,
			}
			s.Del(tt.args.key)
			_, exist := s.m[tt.args.key]
			if exist {
				t.Fail()
			}
		})
	}
}

func Test_hashSet_Has(t *testing.T) {
	type fields struct {
		m map[interface{}]interface{}
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"has nil", fields{m: map[interface{}]interface{}{nil: 0}}, args{key: nil}, true},
		{"has int", fields{m: map[interface{}]interface{}{0: 0}}, args{key: 0}, true},
		{"has int", fields{m: map[interface{}]interface{}{111: 0}}, args{key: 111}, true},
		{"has string", fields{m: map[interface{}]interface{}{"0": 0}}, args{key: "0"}, true},
		{"has string", fields{m: map[interface{}]interface{}{"ttt": 0}}, args{key: "ttt"}, true},
		{"has struct", fields{m: map[interface{}]interface{}{
			pair{timespanInDay{1, 3}, 5}: 0,
		}}, args{
			key: pair{timespanInDay{1, 3}, 5},
		}, true},

		{"not has nil", fields{m: map[interface{}]interface{}{11: 0}}, args{key: nil}, false},
		{"not has int", fields{m: map[interface{}]interface{}{0: 0}}, args{key: 111}, false},
		{"not has int", fields{m: map[interface{}]interface{}{444: 0}}, args{key: 111}, false},
		{"not has string", fields{m: map[interface{}]interface{}{"22": 0}}, args{key: "0"}, false},
		{"not has string", fields{m: map[interface{}]interface{}{"33": 0}}, args{key: "tttt"}, false},
		{"not has struct", fields{m: map[interface{}]interface{}{
			pair{timespanInDay{321, 3}, 5}: 0,
		}}, args{
			key: pair{timespanInDay{1, 5}, 5},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &traversalHashSet{
				m: tt.fields.m,
			}
			if got := s.Has(tt.args.key); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashSet_Insert(t *testing.T) {
	type fields struct {
		m map[interface{}]interface{}
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"insert nil", fields{m: make(map[interface{}]interface{})}, args{key: nil}},
		{"insert int", fields{m: make(map[interface{}]interface{})}, args{key: 0}},
		{"insert int", fields{m: make(map[interface{}]interface{})}, args{key: 111}},
		{"insert string", fields{m: make(map[interface{}]interface{})}, args{key: "0"}},
		{"insert string", fields{m: make(map[interface{}]interface{})}, args{key: "tttt"}},
		{"insert struct", fields{m: make(map[interface{}]interface{})}, args{
			key: pair{timespanInDay{1, 3}, 5},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &traversalHashSet{
				m: tt.fields.m,
			}
			s.Insert(tt.args.key)
			_, exist := s.m[tt.args.key]
			if !exist {
				t.Fail()
			}
		})
	}
}

*/
