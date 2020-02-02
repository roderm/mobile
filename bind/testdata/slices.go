package slices

type ByteArray []byte

type S struct {
	B []byte
	b []byte
}

type (s *S) Set(b1 []byte, b2 []byte) {
	s.b = b1
	s.B = b2
}

type (s *S) Get() ([]byte, []byte){
	return s.b, s.B
}