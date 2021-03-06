package repo

import (
	models2 "auth/internal/models"
	"errors"
)

type mockRepo struct {
	users map[int]map[string]string
	count int
}

func NewmockRepo() mockRepo {
	return mockRepo{count: 0, users: map[int]map[string]string{}}
}

func (rep mockRepo) CreateUser(user models2.User) (string, error) {
	rep.users[rep.count] = map[string]string{user.Username: user.Password}
	rep.count++
	return "", nil
}

func (rep mockRepo) GetUser(user models2.User) (string, string, error) {
	for _, value := range rep.users {
		if m, ok := value[user.Username]; ok {
			return "", m, nil
		}
	}

	return "", "", errors.New("no user")
}
