package errors

const (
	ErrUnknownCode         = "GTH-000"
	ErrBadRequestCode      = "GTH-400"
	ErrNotFoundCode        = "GTH-404"
	ErrClientClosedRequest = "GTH-499"
	ErrInternalCode        = "GTH-500"
)

type Error struct {
	Code     string
	Msg      string
	Metadata map[string]string
}

func (e *Error) Error() string {
	return e.Msg
}

func New(err error, code string, metadata ...map[string]string) *Error {
	if e, ok := err.(*Error); ok {
		if len(metadata) > 0 {
			setMetaData(e, metadata[0])
		}
		return e
	}
	n := &Error{Code: code, Msg: err.Error()}
	if len(metadata) > 0 {
		n.Metadata = metadata[0]
	}
	return n
}

func setMetaData(err *Error, metadata map[string]string) {
	if err.Metadata == nil {
		err.Metadata = metadata
		return
	}
	for k, v := range metadata {
		err.Metadata[k] = v
	}
}

func GetErrorCode(err error) string {
	if err == nil {
		return ""
	}
	e, ok := err.(*Error)
	if !ok {
		return ErrUnknownCode
	}
	if e == nil {
		return ""
	}
	return e.Code
}
