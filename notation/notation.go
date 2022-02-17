package notation

import (
	"fmt"
	"math"
	"strconv"
)

const (
	bitSize = 64
	shift   = 2

	maxStrconvBase = 36
)

// ShiftEquality return true, if the elements meet the conditions
// 	1. [n/2:] elements on shift, have an increasing sequence
// 	2. [:n/2] elements on shift, have a descending sequence
// 	3. sum([n/2:]) equal sum([:n/2])
//
// 	Example:
// 	isEqual, err := notation.ShiftEquality(10, "1", "2", "3", "3", "2", "1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(isEqual) // output: true
func ShiftEquality(m int, n ...string) (bool, error) {
	nums, err := convertToInt64Slice(append([]string(nil), n...), m)
	if err != nil {
		return false, fmt.Errorf("after convert to int64 slice: %v", err)
	}

	position := len(nums) / shift
	left, right := nums[:position], nums[position:]

	if !isNonStrictIncreasingSequence(left) || !isNonStrictDescendingSequence(right) {
		return false, nil
	}

	return sliceSum(left) == sliceSum(right), nil
}

func convertToInt64Slice(vals []string, base int) ([]int64, error) {
	if base > maxStrconvBase {
		return parseToInt64SliceMore36Base(vals, base)
	}

	return parseToInt64Slice(vals, base)
}

func parseToInt64Slice(vals []string, base int) ([]int64, error) {
	nums := make([]int64, 0, len(vals))
	for i := range vals {
		val, err := strconv.ParseInt(vals[i], base, bitSize)
		if err != nil {
			return nil, err
		}
		nums = append(nums, val)
	}

	return nums, nil
}

func parseToInt64SliceMore36Base(vals []string, base int) ([]int64, error) {
	nums := make([]int64, 0, len(vals))
	for i := range vals {
		val, err := parseNumberSystemValue(vals[i], int64(base))
		if err != nil {
			return nil, err
		}
		nums = append(nums, val)
	}

	return nums, nil
}

func parseNumberSystemValue(val string, base int64) (int64, error) {
	sum := int64(0)
	for i, r := range val {
		v, err := strconv.ParseInt(string(r), maxStrconvBase, bitSize)
		if err != nil {
			return 0, err
		}
		if v == 0 {
			continue
		}

		y := len(val) - 1 - i

		pow := int64(math.Pow(
			float64(base),
			float64(y),
		))

		sum += v * pow
	}

	return sum, nil
}

func isNonStrictIncreasingSequence(nums []int64) bool {
	lastValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if lastValue > nums[i] {
			return false
		}
		lastValue = nums[i]
	}

	return true
}

func isNonStrictDescendingSequence(nums []int64) bool {
	lastValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if lastValue < nums[i] {
			return false
		}
		lastValue = nums[i]
	}

	return true
}

func sliceSum(nums []int64) int64 {
	res := int64(0)
	for i := range nums {
		res += nums[i]
	}

	return res
}
