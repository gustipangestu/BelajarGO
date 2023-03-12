package gross

import (
	b "multi/logic"
)

func GrossSalary() int {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}
