package mt

const N = 312
const M = 156
const MATRIX_A = 0xB5026F5AA96619E9
const UPPER_MASK = 0xFFFFFFFF80000000
const LOWER_MASK = 0x7FFFFFFF

var MAG01 = [2]uint64{0, MATRIX_A}

type MT19937_64 struct {
	Array   [312]uint64 //state vector
	Index   uint64      // array index
	Counter uint64
}

func New(seed uint64, counter uint64) *MT19937_64 {
	ret := &MT19937_64{
		Index:   N + 1,
		Counter: counter,
	}
	ret.seed(seed)
	for i := uint64(0); i < counter; i++ {
		ret.Int63()
	}
	return ret
}

func (m *MT19937_64) seed(seed uint64) {
	m.Array[0] = seed
	for m.Index = 1; m.Index < N; m.Index++ {
		m.Array[m.Index] = (6364136223846793005*(m.Array[m.Index-1]^(m.Array[m.Index-1]>>62)) + m.Index)
	}
}

func (m *MT19937_64) Int63() uint64 {
	m.Counter++
	var i int
	var x uint64
	if m.Index >= N {
		if m.Index == N+1 {
			m.seed(5489)
		}

		for i = 0; i < N-M; i++ {
			x = (m.Array[i] & UPPER_MASK) | (m.Array[i+1] & LOWER_MASK)
			m.Array[i] = m.Array[i+(M)] ^ (x >> 1) ^ MAG01[int(x&uint64(1))]
		}
		for ; i < N-1; i++ {
			x = (m.Array[i] & UPPER_MASK) | (m.Array[i+1] & LOWER_MASK)
			m.Array[i] = m.Array[i+(M-N)] ^ (x >> 1) ^ MAG01[int(x&uint64(1))]
		}
		x = (m.Array[N-1] & UPPER_MASK) | (m.Array[0] & LOWER_MASK)
		m.Array[N-1] = m.Array[M-1] ^ (x >> 1) ^ MAG01[int(x&uint64(1))]
		m.Index = 0
	}
	x = m.Array[m.Index]
	m.Index++
	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)
	return x
}

func (m *MT19937_64) Next() int32 {
	return int32(m.Int63() % 10000)
}

func (m *MT19937_64) IntN(v uint64) int32 {
	return int32(m.Int63() % v)
}

func (m *MT19937_64) IntRange(min uint64, max uint64) uint64 {
	v := m.Int63() % (max - min)
	return min + v
}
