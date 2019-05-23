package bitpack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPack(t *testing.T) {
	res, n := Pack([]int{1, 9, 17})
	assert.Len(t, res, 1)
	assert.Equal(t, uint64(0x010101), res[0])
	assert.Equal(t, 1, n)

	res, n = Pack([]int{1, 129})
	assert.Len(t, res, 3)
	assert.Equal(t, uint64(0x01), res[0])
	assert.Equal(t, uint64(0x00), res[1])
	assert.Equal(t, uint64(0x01), res[2])
	assert.Equal(t, 3, n)
}

func TestPackBase(t *testing.T) {
	res, n := PackBase(10, []int{1, 11, 21})
	assert.Len(t, res, 3)
	assert.Equal(t, uint64(0x01), res[0])
	assert.Equal(t, uint64(0x01), res[1])
	assert.Equal(t, uint64(0x01), res[2])
	assert.Equal(t, 3, n)
}

func TestUnpack(t *testing.T) {
	res, n := Unpack([]uint64{0x0F})
	assert.Len(t, res, 4)
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 2, res[1])
	assert.Equal(t, 3, res[2])
	assert.Equal(t, 4, res[3])
	assert.Equal(t, 4, n)

	res, n = UnpackBase(10, []uint64{0x01, 0x02})
	assert.Len(t, res, 2)
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 12, res[1])
	assert.Equal(t, 2, n)
}
