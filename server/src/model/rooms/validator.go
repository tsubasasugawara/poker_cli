package rooms

func ValidatePassword(password string) bool {
	ok := true

	if password == "" {
		ok = false
	}

	return ok
}
