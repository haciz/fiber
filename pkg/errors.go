package pkg

import "fmt"

var (
	ErrEntityAlreadyExists = fmt.Errorf("entity already exists")
	ErrEntityIDNotProvided = fmt.Errorf("entity ID not provided")
	ErrEntityNotFound      = fmt.Errorf("entity not found")
	ErrEntityOutOfRange    = fmt.Errorf("entity out of range")
)
