package butil

// Ptr 返回输入值的指针。
func Ptr[E interface{}](message E) *E {
	return &message
}
