package users

func ValidateName(name string) bool {
	ok := true

	if name == "" {
		ok = false
	}

	return ok
}

func ValidatePassword(password string) bool {
	ok := true

	if password == "" {
		ok = false
	}

	return ok
}
