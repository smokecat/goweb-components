package xcode

import (
	"net/http"
	"strings"
	"unicode"
)

var (
	CodeNil             = New("NIL", 0)
	CodeOK              = NewHttp(http.StatusOK)
	CodeInvalidParams   = New("INVALID_PARAMS", http.StatusBadRequest)
	CodeInvalidArgument = New("INVALID_ARGUMENT", http.StatusBadRequest)
	CodePrivate         = New("PRIVATE_ERROR", http.StatusInternalServerError)
	CodeUnexpectedError = New("UNEXPECTED_ERROR", http.StatusInternalServerError)

	CodeBadRequest          = NewHttp(http.StatusBadRequest)
	CodeUnauthorized        = NewHttp(http.StatusUnauthorized)
	CodeForbidden           = NewHttp(http.StatusForbidden)
	CodeNotFound            = NewHttp(http.StatusNotFound)
	CodeMethodNotAllowed    = NewHttp(http.StatusMethodNotAllowed)
	CodeTooManyRequests     = NewHttp(http.StatusTooManyRequests)
	CodeInternalServerError = NewHttp(http.StatusInternalServerError)
	CodeNotImplemented      = NewHttp(http.StatusNotImplemented)
	CodeBadGateway          = NewHttp(http.StatusBadGateway)
	CodeServiceUnavailable  = NewHttp(http.StatusServiceUnavailable)
	CodeGatewayTimeout      = NewHttp(http.StatusGatewayTimeout)
)

type Code struct {
	text       string
	httpStatus int
}

func New(text string, httpStatus int) Code {
	return Code{text: text, httpStatus: httpStatus}
}

func WithCode(code Code, httpStatus int) Code {
	return Code{text: code.text, httpStatus: httpStatus}
}

func NewHttp(httpStatus int) Code {
	str := http.StatusText(httpStatus)
	if str == "" {
		str = http.StatusText(http.StatusInternalServerError)
	}
	return Code{text: codeText(str), httpStatus: httpStatus}
}

func (c Code) Text() string {
	return c.text
}

func (c Code) HttpStatus() int {
	return c.httpStatus
}

func (c Code) IsNil() bool {
	return c == CodeNil || (c.text == "" && c.httpStatus == 0)
}

func (c Code) IsOK() bool {
	return c == CodeOK
}

func codeText(str string) string {
	var builder strings.Builder
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune('_')
		}
	}
	return builder.String()
}
