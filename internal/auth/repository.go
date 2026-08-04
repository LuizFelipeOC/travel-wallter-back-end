package auth

import (
	"fmt"
	"sync"
)

type AuthRepository struct {
	users map[string]*User
	mu    sync.RWMutex
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		users: make(map[string]*User),
	}
}

func (r *AuthRepository) CreateAccount(req CreateUserRequest) (*User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[req.Email]; exists {
		return nil, fmt.Errorf("email já cadastrado")
	}

	user := &User{
		ID:    fmt.Sprintf("user_%d", len(r.users)+1),
		Email: req.Email,
		Name:  req.Name,
	}

	r.users[req.Email] = user
	return user, nil
}

func (r *AuthRepository) Login(req LoginRequest) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[req.Email]
	if !exists {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	return user, nil
}
