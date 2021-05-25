package user

func (u *User) IsEqual(other *User) bool {
	if other == nil {
		return false
	}

	if u.Email == other.Email || u.UserID.String() == other.UserID.String() {
		return true
	}

	return false
}
