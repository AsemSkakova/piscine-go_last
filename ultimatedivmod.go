package piscine

func UltimateDivMod(a *int, b *int) {
	var division int
	var modulo int
	division = *a / (*b)
	modulo = *a % (*b)
	*a = division
	*b = modulo
}
