package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePicture(t *testing.T) {
	input := CreatePictureParams {
		Url: "https://images.unsplash.com/photo-1529778873920-4da4926a72c2",
		Matches: 0,
		Rating: 1500,
	}

	picture, err := testQueries.CreatePicture(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, picture)

	require.Equal(t, input.Url, picture.Url)
	require.Equal(t, input.Matches, picture.Matches)
	require.Equal(t, input.Rating, picture.Rating)

	require.NotZero(t, picture.ID)
	require.NotZero(t, picture.CreatedAt)
}