package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/fadhilaf/s-tech/internal/model"
	repositoryModel "github.com/fadhilaf/s-tech/internal/repository/postgres/sqlc"
	"github.com/fadhilaf/s-tech/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- 1. Mock Store Definition ---
type MockStore struct { // repository.Store yang ditambah utility mock
	repository.Store 
	mock.Mock // digabung dengan utility mock
}

// Implementasi GetProductByName untuk kebutuhan test
func (m *MockStore) GetProductByName(ctx context.Context, name string) (repositoryModel.Product, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(repositoryModel.Product), args.Error(1)
}

// Implementasi CreateProduct untuk kebutuhan test
func (m *MockStore) CreateProduct(ctx context.Context, arg repositoryModel.CreateProductParams) (sql.Result, error) {
	args := m.Called(ctx, arg)
	// SQLC return sql.Result. Kita bisa return nil untuk unit test ini.
	res, _ := args.Get(0).(sql.Result)
	return res, args.Error(1)
}

// --- 2. Unit Test Logic ---
func TestCreateProduct(t *testing.T) {
	testCases := []struct {
		name           string
		input          model.CreateProductRequest
		setupMock      func(m *MockStore)
		expectedStatus int
		expectedMsg    string
	}{
		{
			name: "Success - Berhasil Simpan",
			input: model.CreateProductRequest{
				NotFile: model.CreateProductNoFileRequest{
					Name: "Keyboard RGB",
				},
			},
			setupMock: func(m *MockStore) {
				m.On("GetProductByName", mock.Anything, "Keyboard RGB").
					Return(repositoryModel.Product{}, errors.New("sql: no rows in result set"))
				m.On("CreateProduct", mock.Anything, mock.Anything).
					Return((sql.Result)(nil), nil)
			},
			expectedStatus: http.StatusCreated,
			expectedMsg:    "Berhasil memasukkan produk ke database",
		},
		{
			name: "Conflict - Nama Sudah Ada",
			input: model.CreateProductRequest{
				NotFile: model.CreateProductNoFileRequest{
					Name: "Keyboard RGB",
				},
			},
			setupMock: func(m *MockStore) {
				// Return nil error berarti produk ditemukan
				m.On("GetProductByName", mock.Anything, "Keyboard RGB").
					Return(repositoryModel.Product{Name: "Keyboard RGB"}, nil)
			},
			expectedStatus: http.StatusConflict,
			expectedMsg:    "Produk dengan nama yang sama sudah ada",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := new(MockStore)
			tc.setupMock(mockStore)

			// Inisialisasi usecase dengan mockStore
			uc := NewProductUsecase(mockStore)

			resp := uc.CreateProduct(tc.input)

			assert.Equal(t, tc.expectedStatus, resp.Status)
			assert.Equal(t, tc.expectedMsg, resp.Message)
			mockStore.AssertExpectations(t)
		})
	}
}
