package utils

import (
	"fmt"
	"math"
)

var units = []string{"","Ki","Mi","Gi","Ti","Pi","Ei","Zi"}

func SizeOf(size int64) string {
	s64 := float64(size)
	for _, unit := range units {
		if math.Abs(s64) < 1024 {
			return fmt.Sprintf("%3.1f%sB",s64, unit)
		}
		s64/=1024
	}
}


def sizeof_fmt(num, suffix="B"):
for unit in ["","Ki","Mi","Gi","Ti","Pi","Ei","Zi"]:
if abs(num) < 1024.0:
return "%3.1f%s%s" % (num, unit, suffix)
num /= 1024.0
return "%.1f%s%s" % (num, "Yi", suffix)
