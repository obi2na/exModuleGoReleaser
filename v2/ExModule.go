package exModuleGoReleaser

import (
	"golang.org/x/exp/constraints"
)

func Add[T Number](a, b T) T {
	return a + b
}

type Number interface {
	constraints.Integer | constraints.Float
}
