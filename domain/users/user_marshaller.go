package users

// PrivateUser struct
type PrivateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

// PublicUser struct
type PublicUser struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
}

// Marshall meth for list of users
func (users Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, 0)
	for _, user := range users {
		result = append(result, user.Marshall(isPublic))
	}
	return result
}

// Marshall meth
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			Status:    user.Status,
		}
	}

	return PrivateUser{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Status:    user.Status,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
