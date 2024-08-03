package test

import (
	"media-server/pkg/model"

	"github.com/stretchr/testify/mock"
)

// MockMediaRepository é o mock do repositório de mídia
type MockMediaRepository struct {
	mock.Mock
}

// CreateMedia mock method
func (m *MockMediaRepository) CreateMedia(media model.Media) (model.Media, error) {
	args := m.Called(media)
	return args.Get(0).(model.Media), args.Error(1)
}

// ListMedias mock method
func (m *MockMediaRepository) ListMedias() ([]model.Media, error) {
	args := m.Called()
	return args.Get(0).([]model.Media), args.Error(1)
}
