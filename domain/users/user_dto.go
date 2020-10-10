package users

// data transfer object
import (
	"strings"

	"github.com/emadghaffari/bookstore_uesrs-api/utils/errors"
)

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Validate Email for user
func (user *User) Validate() *errors.ResError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	if user.Email = strings.TrimSpace(strings.ToLower(user.Email)); user.Email == "" {
		return errors.HandlerBagRequest("Invalid Email Address")
	}

	return nil
}
