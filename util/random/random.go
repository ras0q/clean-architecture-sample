package random

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	domain "github.com/ras0q/clean-architecture-sample/0_domain"
)

var (
	rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	defaultLen = 5
)

func AlphaNumeric(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}

	return string(b)
}

func Email() string {
	return fmt.Sprintf("%s@%s.com", AlphaNumeric(defaultLen), AlphaNumeric(defaultLen))
}

func UUID() uuid.UUID {
	return uuid.New()
}

func Error() error {
	return errors.New(AlphaNumeric(defaultLen))
}

func User() *domain.User {
	u := domain.NewUser(UUID(), AlphaNumeric(defaultLen), Email())
	return &u
}
