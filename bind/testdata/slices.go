package slices

type ByteArray []byte

type S struct {
	B []byte
	I []int32
	b []byte
	i []int32
}

func (s *S) SingleSetBytes(b1 []byte) {
	s.b = b1
}

func (s *S) SingleGetBytes() []byte {
	return s.b
}

func (s *S) DoubleSetBytes(b1 []byte, b2 []byte) {
	s.b = b1
	s.B = b2
}

func (s *S) SingleSetInts(i1 []int32) {
	s.i = i1
}

func (s *S) SingleGetInts() []int32 {
	return s.i
}

func (s *S) DoubleSetInts(i1 []int32, i2 []int32) {
	s.i = i1
	s.I = i2
}
