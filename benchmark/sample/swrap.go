package swrap

type SWrap []byte

// create swrap from []byte and return
func New(a []byte) SWrap {
	return SWrap(a)
}

func (sw *SWrap) Add(a byte) {
	*sw = append(*sw, a)
}

func (sw *SWrap) Len() int {
	return len(*sw)
}
