package size

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	_kB uint64 = 1 << 10
	_MB uint64 = 1 << 20
	_GB uint64 = 1 << 30
	_TB uint64 = 1 << 40
	_PB uint64 = 1 << 50
	_EB uint64 = 1 << 60
)

// Parse a string that has a size suffix (one of k, M, G, T, P, E).
// The suffix denotes multiples of 1024.
// e.g., "32k", "2M"
func ParseSize(s string) (uint64, error) {
	var suff string

	i := strings.LastIndexAny(s, "kKMGTPE")
	if i > 0 {
		suff = s[i:]
		s = s[:i]
	}

	if s == "" {
		return 0, nil
	}

	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	var multmap map[string]uint64 = map[string]uint64{
		"":  1,
		"k": _kB,
		"K": _kB,
		"M": _MB,
		"G": _GB,
		"T": _TB,
		"P": _PB,
		"E": _EB,
	}

	if m, ok := multmap[suff]; ok {
		v *= m
	} else {
		return 0, fmt.Errorf("unknown size suffix %s", suff)
	}

	return v, nil
}
