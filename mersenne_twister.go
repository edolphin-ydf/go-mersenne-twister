package mt

import (
	"encoding/binary"
)

const (
	n         = 624
	m         = 397
	matrixA   = 0x9908b0df
	upperMask = 0x80000000
	lowerMask = 0x7fffffff
)

type MtRandom struct {
	mt  [n]uint32
	mti uint32
}

// create and initializes random sequence with a seed value
func New() *MtRandom {
	res := &MtRandom{
		mti: n + 1,
	}

	return res
}

func (r *MtRandom) Seed(s uint32) {
	r.mt[0] = s & 0xffffffff
	for r.mti = 1; r.mti < n; r.mti++ {
		r.mt[r.mti] = 1812433253*(r.mt[r.mti-1]^(r.mt[r.mti-1]>>30)) + r.mti
		r.mt[r.mti] &= 0xffffffff
	}
}

// InitByArray initializes random sequence with a slice
func (r *MtRandom) InitByArray(initKey []uint32) {
	var i, j, k uint32
	r.Seed(19650218)
	i = 1
	j = 0
	keyLength := uint32(len(initKey))
	k = keyLength
	if n > keyLength {
		k = n
	}
	for ; k != 0; k-- {
		r.mt[i] = (r.mt[i] ^ ((r.mt[i-1] ^ (r.mt[i-1] >> 30)) * 1664525)) + initKey[j] + j
		r.mt[i] &= 0xffffffff
		i++
		j++
		if i >= n {
			r.mt[0] = r.mt[n-1]
			i = 1
		}
		if j >= keyLength {
			j = 0
		}
	}
	for k = n - 1; k != 0; k-- {
		r.mt[i] = (r.mt[i] ^ ((r.mt[i-1] ^ (r.mt[i-1] >> 30)) * 1566083941)) - i
		r.mt[i] &= 0xffffffff
		i++
		if i >= n {
			r.mt[0] = r.mt[n-1]
			i = 1
		}
	}

	r.mt[0] = 0x80000000
}

// UInt32 generates a random 32bit unsigned int number
func (r *MtRandom) UInt32() uint32 {
	var y uint32
	mag01 := [2]uint32{0x0, matrixA}

	if r.mti >= n {
		var kk int

		if r.mti == n+1 {
			r.Seed(5489)
		}
		for kk = 0; kk < n-m; kk++ {
			y = (r.mt[kk] & upperMask) | (r.mt[kk+1] & lowerMask)
			r.mt[kk] = r.mt[kk+m] ^ (y >> 1) ^ mag01[y&0x1]
		}
		for ; kk < n-1; kk++ {
			y = (r.mt[kk] & upperMask) | (r.mt[kk+1] & lowerMask)
			r.mt[kk] = r.mt[kk+(m-n)] ^ (y >> 1) ^ mag01[y&0x1]
		}
		y = (r.mt[n-1] & upperMask) | (r.mt[0] & lowerMask)
		r.mt[n-1] = r.mt[m-1] ^ (y >> 1) ^ mag01[y&0x1]

		r.mti = 0
	}

	y = r.mt[r.mti]
	r.mti++

	y ^= (y >> 11)
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= (y >> 18)

	return y
}

// GenrandInt31 generates a 31bit unsigned int random number
// note: return type is uint32
func (r *MtRandom) GenrandInt31() uint32 {
	return uint32(r.UInt32() >> 1)
}

// GenrandReal1 generates a 32bit [0, 1] real random number
// note: return type is float64, not float32
func (r *MtRandom) GenrandReal1() float64 {
	return float64(r.UInt32()) * (1.0 / 4294967295.0)
}

// GenrandReal2 generates a 32bit [0, 1) real random number
// note: return type is float64, not float32
func (r *MtRandom) GenrandReal2() float64 {
	return float64(r.UInt32()) * (1.0 / 4294967296.0)
}

// GenrandReal3 generates a 32bit (0, 1) real random  number
// note: return type is float64, not float32
func (r *MtRandom) GenrandReal3() float64 {
	return ((float64(r.UInt32())) + 0.5) * (1.0 / 4294967296.0)
}

// GenrandRes53 generates a [0, 1) random number with 53-bit resolution
func (r *MtRandom) GenrandRes53() float64 {
	a := r.UInt32() >> 5
	b := r.UInt32() >> 6
	return (float64(a)*67108864.0 + float64(b)) * (1.0 / 9007199254740992.0)
}

func (r *MtRandom) Serialize() []byte {
	var res = make([]byte, (n+1)*4)
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], r.mti)
	copy(res, buf[:])

	for i, v := range r.mt {
		binary.LittleEndian.PutUint32(buf[:], v)
		copy(res[(i+1)*4:], buf[:])
	}
	return res
}

func (r *MtRandom) Deserialize(data []byte) {
	r.mti = binary.LittleEndian.Uint32(data[:4])
	for i := 0; i < n; i++ {
		start := (i + 1) * 4
		r.mt[i] = binary.LittleEndian.Uint32(data[start:])
	}
}

var mt = New()

func InitGenrand(s uint32) {
	mt.Seed(s)
}

func InitByArray(initKey []uint32) {
	mt.InitByArray(initKey)
}

func UInt32() uint32 {
	return mt.UInt32()
}

func GenrandInt31() uint32 {
	return mt.GenrandInt31()
}

func GenrandReal1() float64 {
	return mt.GenrandReal1()
}

func GenrandReal2() float64 {
	return mt.GenrandReal2()
}

func GenrandReal3() float64 {
	return mt.GenrandReal3()
}

func GenrandRes53() float64 {
	return mt.GenrandRes53()
}
