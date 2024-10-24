package xerr

import (
	"github.com/pkg/errors"

	"github.com/smokecat/goweb-components/pkg/xcode"
)

type XErr interface {
	Code() xcode.Code
	Msg() string
	Err() error
	ErrType() ErrType
	Meta() map[string]any
	WithMeta(key string, value any) XErr
	IsTypePublic() bool
	Error() string
}

type ErrType int

const (
	ErrTypePublic ErrType = iota
	ErrTypePrivate
)

type xe struct {
	code    xcode.Code
	msg     string
	err     error
	errType ErrType
	meta    map[string]any
}

func New(code xcode.Code, msg string) XErr {
	return newXe(code, msg)
}

func Wrap(err error, code xcode.Code, msg string) XErr {
	if err == nil {
		return newXe(code, msg)
	}

	var joinErr *xe
	if e, ok := err.(*xe); ok {
		joinErr = e
		joinErr.code = code
		joinErr.msg = msg
		joinErr.err = errors.Wrap(e.err, msg)
	} else {
		joinErr = newXe(code, msg)
		joinErr.err = errors.Wrap(err, msg)
	}

	return joinErr
}

func (e *xe) Code() xcode.Code {
	return e.code
}

func (e *xe) Msg() string {
	return e.msg
}

func (e *xe) Err() error {
	if e.code.IsNil() || e.code.IsOK() {
		return nil
	}

	if e.err != nil {
		return e.err
	}

	return newError(e.code, e.msg)
}

func (e *xe) ErrType() ErrType {
	return e.errType
}

func (e *xe) Meta() map[string]any {
	return e.meta
}

func (e *xe) WithMeta(key string, value any) XErr {
	if e == nil {
		return nil
	}

	if e.meta == nil {
		e.meta = make(map[string]any)
	}

	e.meta[key] = value
	return e
}

func (e *xe) IsTypePublic() bool {
	return e.equalType(ErrTypePublic)
}

func (e *xe) Error() string {
	if e.msg != "" {
		return e.msg
	}

	if e.err != nil {
		return e.err.Error()
	}

	return e.code.Text()
}

func (e *xe) equalType(errType ErrType) bool {
	return e.errType == errType
}

func newXe(code xcode.Code, msg string) *xe {
	if msg == "" {
		msg = code.Text()
	}

	return &xe{
		code:    code,
		msg:     msg,
		err:     newError(code, msg),
		errType: ErrTypePublic,
		meta:    nil,
	}
}

func newError(code xcode.Code, msg string) error {
	if code.IsNil() || code.IsOK() {
		return nil
	}
	if msg == "" {
		return errors.New(code.Text())
	}
	return errors.New(msg)
}
