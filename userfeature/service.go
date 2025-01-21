package userfeature

import (
	"api1/metrics"
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	users   map[uint]*User
	nextID  uint
	mu      sync.RWMutex
	logger  *logrus.Logger
	metrics *metrics.Metrics
}

func NewUserService(logger *logrus.Logger, metrics *metrics.Metrics) *UserService {
	return &UserService{
		users:   make(map[uint]*User),
		nextID:  1,
		logger:  logger,
		metrics: metrics,
	}
}

func (s *UserService) Create(req *CreateUserRequest) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check for duplicate email
	for _, u := range s.users {
		if u.Email == req.Email {
			s.metrics.UserOperationsTotal.WithLabelValues("create", "failure").Inc()
			s.logger.WithField("email", req.Email).Warn("Attempt to create user with duplicate email")
			return nil, errors.New("email already exists")
		}
	}

	user := &User{
		ID:        s.nextID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users[user.ID] = user
	s.nextID++

	// Update metrics
	s.metrics.UsersTotal.Inc()
	s.metrics.UserOperationsTotal.WithLabelValues("create", "success").Inc()

	s.logger.WithFields(logrus.Fields{
		"id":    user.ID,
		"email": user.Email,
	}).Info("Created new user")

	return user, nil
}

func (s *UserService) Get(id uint) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		s.logger.WithField("id", id).Warn("User not found")
		return nil, errors.New("user not found")
	}

	s.logger.WithField("id", id).Info("Retrieved user")
	return user, nil
}

func (s *UserService) List() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	s.logger.WithField("count", len(users)).Info("Listed all users")
	return users
}

func (s *UserService) Update(id uint, req *UpdateUserRequest) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		s.logger.WithField("id", id).Warn("Attempt to update non-existent user")
		return nil, errors.New("user not found")
	}

	if req.Email != "" {
		// Check for duplicate email
		for _, u := range s.users {
			if u.ID != id && u.Email == req.Email {
				s.logger.WithField("email", req.Email).Warn("Attempt to update user with duplicate email")
				return nil, errors.New("email already exists")
			}
		}
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	user.UpdatedAt = time.Now()

	s.logger.WithFields(logrus.Fields{
		"id":    user.ID,
		"email": user.Email,
	}).Info("Updated user")

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		s.metrics.UserOperationsTotal.WithLabelValues("delete", "failure").Inc()
		s.logger.WithField("id", id).Warn("Attempt to delete non-existent user")
		return errors.New("user not found")
	}

	delete(s.users, id)
	// Update metrics
	s.metrics.UsersTotal.Dec()
	s.metrics.UserOperationsTotal.WithLabelValues("delete", "success").Inc()

	s.logger.WithField("id", id).Info("Deleted user")
	return nil
}
