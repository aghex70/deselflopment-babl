package pkg

import "github.com/gofrs/uuid"

func GenerateUUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return u.String()
}
