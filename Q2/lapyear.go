package q2

func CekLapyear(year int) bool {
	status := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				status = true
			} else {
				status = false
			}
		} else {
			status = true
		}
	} else {
		status = false
	}
	return status
}
