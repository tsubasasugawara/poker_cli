package participants

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
