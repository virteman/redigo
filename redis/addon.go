package redis

import (
	"fmt"
	"strconv"
)

func Uint64s(reply interface{}, err error) ([]uint64, error) {
	var result []uint64
	err = sliceHelper(reply, err, "Uint64s", func(n int) { result = make([]uint64, n) }, func(i int, v interface{}) error {
		switch v := v.(type) {
		case uint64:
			result[i] = v
			return nil
		case []byte:
			n, err := strconv.ParseUint(string(v), 10, 64)
			result[i] = n
			return err
		default:
			return fmt.Errorf("redigo: unexpected element type for Int64s, got type %T", v)
		}
	})
	return result, err
}
