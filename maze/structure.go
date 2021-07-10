package maze

type Structure interface {
	IsBlocked() bool
	print()
	Cost() int
}
