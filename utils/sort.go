package utils

import "sort"

func SortInt64Slice(ds []int64) {
	if len(ds) == 0 {
		return
	}

	sort.Slice(ds, func(i, j int) bool {
		return ds[i] < ds[j]
	})
}

func SortInt64SliceStable(ds []int64) {
	if len(ds) == 0 {
		return
	}

	sort.SliceStable(ds, func(i, j int) bool {
		return ds[i] < ds[j]
	})
}
