package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100

	bonus := ((purchase / 100) / 100) * percent
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
