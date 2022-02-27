package hello

//go:noinline
func HelloNoInline() interface{} {
	return struct{}{}
}

func Hello() interface{} {
	return struct{}{}
}

// ref: https://qiita.com/tutuz/items/caa5d85544c398a2da9a
