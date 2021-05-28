package memauth

import (
	"fmt"
	"gocoin/crypto/user"
	"time"

	"github.com/gofrs/uuid"
	"github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
)

func New() *MemAuth {
	auth := MemAuth{
		Cache: cache.New(25*time.Minute, 30*time.Minute),
	}

	return &auth
}

func (m *MemAuth) CreateUser(email, password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	if err != nil {
		return err
	}

	userUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	user := user.User{
		UserID:   userUUID,
		Email:    email,
		Password: string(pass),
	}

	if err := m.Cache.Add(email, &user, 25*time.Minute); err != nil {
		return err
	}

	return nil
}

func (m *MemAuth) AuthUser(email, password string) (*user.User, error) {
	if userIface, found := m.Cache.Get(email); found {
		if user, ok := userIface.(*user.User); ok {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
				return nil, err
			}

			return user, nil
		}
	}

	return nil, fmt.Errorf("unable to find user to auth")
}

func (m *MemAuth) GetUser(userID, email string) (*user.User, error) {
	if userIface, found := m.Cache.Get(email); found {
		if user, ok := userIface.(*user.User); ok {
			if user.UserID.String() == userID {
				return user, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find user")
}

func (m *MemAuth) GetUserByID(userID string) (*user.User, error) {
	for _, userObj := range m.Cache.Items() {
		if user, ok := userObj.Object.(*user.User); ok {
			if user.UserID.String() == userID {
				return user, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find user")
}
