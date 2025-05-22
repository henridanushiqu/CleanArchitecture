package service_test

import (
	"context"
	"testing"

	"myapp/pkg/entity"
	"myapp/pkg/mocks"
	"myapp/pkg/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCountries(t *testing.T) {
	// Step 1: Create mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Step 2: Create mock repository
	mockRepo := mocks.NewMockCountryRepository(ctrl)

	// Step 3: Define expected behavior
	mockRepo.EXPECT().
		GetCountries(gomock.Any(), gomock.Any()).
		Return([]entity.Country{{ID: 1, Country: "Canada"}}, nil)

	// Step 4: Inject mock into service
	svc := service.NewCountryService(context.Background(), mockRepo)

	// Step 5: Call the method
	result, err := svc.GetCountries(context.Background(), entity.Pagination{
		RowsNumber: "10",
		PageNumber: "1",
	})

	// Step 6: Assertions
	require.NoError(t, err)
	require.Len(t, result, 1)
	require.Equal(t, "Canada", result[0].Country)
}
