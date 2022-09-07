package model

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

func ValidateUserId(id string) bool {
	ok := true

	if id == "" {
		ok = false
	}
	if len(id) != 36 {
		ok = false
	}

	return ok
}

func ValidateRoomId(id string) bool {
	ok := true

	if id == "" {
		ok = false
	}
	if len(id) != 36 {
		ok = false
	}

	return ok
}
