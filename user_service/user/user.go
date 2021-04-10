package user

import (
	"fmt"
	"regexp"

	proto "user_service/proto/user"
)

type validateUser struct {
	*proto.User
}

// Validate User fields
func (v validateUser) Validate() error {

	if v.Name == "" {
		return fmt.Errorf("invalid name")
	}

	// Validates user email using regex
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(v.Email) {
		return fmt.Errorf("invalid email %s", v.Email)
	}

	// Validates user phone using regex supports only 10 digits
	// validates with spaces, dashes, brackets etc.
	re = regexp.MustCompile(`^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`)

	if !re.MatchString(v.Phone) {
		return fmt.Errorf("invalid phone %s", v.Phone)
	}

	return nil
}
