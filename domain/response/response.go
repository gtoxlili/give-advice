package response

// R is a struct that contains the result of a request.
type R[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func Ok[T any](data T) *R[T] {
	return &R[T]{Code: 200, Data: data}
}

func Fail(code int, message string) R[struct{}] {
	return R[struct{}]{Code: code, Message: message}
}

func FailWith(err error) R[struct{}] {
	return R[struct{}]{Code: 500, Message: err.Error()}
}
