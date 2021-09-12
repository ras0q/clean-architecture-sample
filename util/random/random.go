package random

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofrs/uuid"
)

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func AlphaNumeric(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}

	return string(b)
}

func Email() string {
	l := 5

	return fmt.Sprintf("%s@%s.com", AlphaNumeric(l), AlphaNumeric(l))
}

func UUID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
