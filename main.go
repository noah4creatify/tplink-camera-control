package main

const (
	shortToken = "RDpbLfCPsJZ7fiv"
	longToken  = "yLwVl0zKqws7LgKPRQ84Mdt708T1qQ3Ha7xv3H7NyU84p21BriUWBU43odz3iP4rBL3cD02KZciXTysVXiV8ngg6vL48rPJyAUw0HurW20xqxv9aYb4M9wK1Ae0wlro510qXeU07kV57fQMc8L6aLgMLwygtc0F10a0Dg70TOoouyFhdysuRMO51yY5ZlOZZLEal1h0t9YQW0Ko7oBwmCAHoic4HYbUyVeU3sfQ1xtXcPcf1aT303wAQhv66qzW"
)

func securityEncode(password string) (result string) {
	if "" != password {
		var limitLength int

		passLenth := len(password)
		shortLength := len(shortToken)
		longLength := len(longToken)

		if passLenth > shortLength {
			limitLength = passLenth
		} else {
			limitLength = shortLength
		}

		for i := 0; limitLength > i; i++ {
			n1 := 187
			n2 := 187

			if passLenth <= i {
				n1 = int(shortToken[i])
			} else if shortLength <= i {
				n2 = int(password[i])
			} else {
				n1 = int(shortToken[i])
				n2 = int(password[i])
			}

			result += string(longToken[(n1^n2)%longLength])
		}
	}

	return
}

func main() {}
