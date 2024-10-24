package xsql

import (
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func TestNewNilNullInt64(t *testing.T) {
	tests := []struct {
		name  string
		wantN sql.NullInt64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNilNullInt64(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNilNullInt64() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewNilNullString(t *testing.T) {
	tests := []struct {
		name  string
		wantN sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNilNullString(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNilNullString() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewNilNullTime(t *testing.T) {
	tests := []struct {
		name  string
		wantN sql.NullTime
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNilNullTime(); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNilNullTime() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewNull(t *testing.T) {
	type args[T any] struct {
		v T
	}
	type testCase[T any] struct {
		name  string
		args  T
		wantN T
	}
	tests := []testCase[any]{
		{"case0", 10, 10},
		{"case1", "abc", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNull(tt.args); !reflect.DeepEqual(gotN.V, tt.wantN) {
				t.Errorf("NewNull() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewNullInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name  string
		args  args
		wantN sql.NullInt64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNullInt64(tt.args.v); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNullInt64() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewNullString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name  string
		args  args
		wantN sql.NullString
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNullString(tt.args.v); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNullString() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestNewNullTime(t *testing.T) {
	type args struct {
		v time.Time
	}
	tests := []struct {
		name  string
		args  args
		wantN sql.NullTime
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := NewNullTime(tt.args.v); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("NewNullTime() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
