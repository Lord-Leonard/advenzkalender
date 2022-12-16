package day_three

func DeterminePriority(inhalt byte) byte {
	priority := inhalt

	if priority > 95 {
		priority -= 96
	} else {
		priority -= 38
	}
	return priority
}
