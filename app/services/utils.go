package service

import "math"

func CalcPaginateTotalPages(total_rows, limit int) int {
	return int(math.Ceil(float64(total_rows) / float64(limit)))
}
