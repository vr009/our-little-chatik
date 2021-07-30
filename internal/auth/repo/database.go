package repo

import (
	"errors"
	"our-little-chatik/internal/models"
)

type mockRepo struct {
	users map[int]map[string]string
	count int
}

func NewmockRepo() mockRepo {
	return mockRepo{count: 0, users: map[int]map[string]string{}}
}

func (rep mockRepo) CreateUser(user models.User) error {
	rep.users[rep.count] = map[string]string{user.UserName: user.Password}
	rep.count++
	return nil
}

func (rep mockRepo) GetUser(user models.User) (string, error) {
	for _, value := range rep.users {
		if m, ok := value[user.UserName]; ok {
			return m, nil
		}
	}

	return "", errors.New("no user")
}
