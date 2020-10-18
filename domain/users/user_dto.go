package users

// data transfer object
import (
	"strings"

	"github.com/emadghaffari/rest_uesrs-api/utils/errors"
)

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

// Users list
type Users []User

// Validate Email for user
func (user *User) Validate() *errors.ResError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	if user.Email = strings.TrimSpace(strings.ToLower(user.Email)); user.Email == "" {
		return errors.HandlerBagRequest("Invalid Email Address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.HandlerBagRequest("Invalid Password")
	}

	return nil
}
