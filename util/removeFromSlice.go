package util

func Remove(slice []string, s int) []string {
	a := append(slice[:s], slice[s+1:]...)
	return a
}
