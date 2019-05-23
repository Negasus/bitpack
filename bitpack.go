package bitpack

import (
	"math/bits"
	"sort"
)

const (
	defaultBase int = 64
)

func Pack(values []int) ([]uint64, int) {
	return PackBase(defaultBase, values)
}

func PackBase(base int, values []int) ([]uint64, int) {
	return PackBaseBuf(base, values, make([]uint64, 0))
}

func PackBaseBuf(base int, values []int, buf []uint64) ([]uint64, int) {
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	return PackBaseBufSorted(base, values, buf)
}

func PackBaseBufSorted(base int, values []int, buf []uint64) ([]uint64, int) {
	if len(values) == 0 {
		return buf, 0
	}

	groupCount := (values[0] / base) + 1
	if values[0]%base == 0 {
		groupCount--
	}

	groups := make([][]int, groupCount)
	for _, v := range values {
		if v == 0 {
			continue
		}
		groupIndex := v / base
		if v%base == 0 {
			groupIndex--
		}
		groups[groupIndex] = append(groups[groupIndex], v-base*groupIndex)
	}

	masks := make([]uint64, len(groups))

	for groupIndex, group := range groups {
		if len(group) == 0 {
			masks[groupIndex] = 0
			continue
		}
		for i := 1; i < len(group); i++ {
			masks[groupIndex] = (masks[groupIndex] | 1) << uint(group[i-1]-group[i])
		}
		masks[groupIndex] = (masks[groupIndex] | 1) << uint(group[len(group)-1]-1)
	}

	return masks, len(groups)
}

func Unpack(values []uint64) ([]int, int) {
	return UnpackBase(defaultBase, values)
}

func UnpackBase(base int, values []uint64) ([]int, int) {
	return UnpackBaseBuf(base, values, make([]int, 0))
}

func UnpackBaseBuf(base int, values []uint64, buf []int) ([]int, int) {

	var count int

	if len(values) == 0 {
		return buf, count
	}

	for index, mask := range values {
		if mask == 0 {
			continue
		}
		var s uint
		for {
			m := uint(bits.TrailingZeros64(uint64(mask)))
			mask = mask >> m >> 1
			s += m + 1
			buf = append(buf, int(s)+(index*base))
			count++
			if mask == 0 {
				break
			}
		}
	}

	return buf, count
}
