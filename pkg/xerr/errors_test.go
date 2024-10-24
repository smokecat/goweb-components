package xerr

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/smokecat/goweb-components/pkg/xcode"
)

func TestJoin(t *testing.T) {
	type args struct {
		err  error
		code xcode.Code
		msg  string
	}
	tests := []struct {
		name string
		args args
		want *xe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.err, tt.args.code, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Run("New with empty msg", func(t *testing.T) {
		e0 := New(xcode.CodeOK, "")
		if fmt.Sprintf("%+v", e0) != xcode.CodeOK.Text() {
			t.Errorf("New() = %v, want %v", e0, xcode.CodeOK.Text())
		}

		if e0.Error() != xcode.CodeOK.Text() || e0.ErrType() != ErrTypePublic || e0.Meta() != nil {
			t.Errorf("New() = %v, want %v", e0, &xe{
				code:    xcode.CodeOK,
				msg:     xcode.CodeOK.Text(),
				err:     fmt.Errorf("%s", xcode.CodeOK.Text()),
				errType: ErrTypePublic,
				meta:    nil,
			})
		}
	})

	t.Run("New with non-empty msg", func(t *testing.T) {
		const e1Msg = "Custom e1 msg"
		e1 := New(xcode.CodeOK, e1Msg)
		if fmt.Sprintf("%+v", e1) != e1Msg {
			t.Errorf("New() = %v, want %v", e1, e1Msg)
		}
	})
}

func Test_xe_Error(t *testing.T) {
	type fields struct {
		code    xcode.Code
		msg     string
		err     error
		errType ErrType
		meta    map[string]any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &xe{
				code:    tt.fields.code,
				msg:     tt.fields.msg,
				err:     tt.fields.err,
				errType: tt.fields.errType,
				meta:    tt.fields.meta,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xe_WithMeta(t *testing.T) {
	type fields struct {
		code    xcode.Code
		msg     string
		err     error
		errType ErrType
		meta    map[string]any
	}
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *xe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &xe{
				code:    tt.fields.code,
				msg:     tt.fields.msg,
				err:     tt.fields.err,
				errType: tt.fields.errType,
				meta:    tt.fields.meta,
			}
			if got := e.WithMeta(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
