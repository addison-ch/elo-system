package db

import (
	"context"
	"database/sql"
	"elo-system/main-api/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreatePicture(t *testing.T) {
	var desc sql.NullString
	desc.Valid = true
	desc.String = util.RandomDescription()

	input := CreatePictureParams{
		Description: desc,
		Url:         util.RandomUrl(),
		Matches:     util.RandomMatches(),
		Rating:      util.RandomRating(),
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

func createRandomPicture(t *testing.T) Picture {
	var desc sql.NullString
	desc.Valid = true
	desc.String = util.RandomDescription()

	input := CreatePictureParams{
		Description: desc,
		Url:         util.RandomUrl(),
		Matches:     util.RandomMatches(),
		Rating:      util.RandomRating(),
	}

	picture, err := testQueries.CreatePicture(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, picture)

	require.Equal(t, input.Url, picture.Url)
	require.Equal(t, input.Matches, picture.Matches)
	require.Equal(t, input.Rating, picture.Rating)

	require.NotZero(t, picture.ID)
	require.NotZero(t, picture.CreatedAt)

	return picture
}

func TestGetPicture(t *testing.T) {
	picture1 := createRandomPicture(t)
	picture2, err := testQueries.GetPicture(context.Background(), picture1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, picture2)

	require.Equal(t, picture1.ID, picture2.ID)
	require.Equal(t, picture1.Description, picture2.Description)
	require.Equal(t, picture1.Url, picture2.Url)
	require.Equal(t, picture1.Matches, picture2.Matches)
	require.Equal(t, picture1.Rating, picture2.Rating)
	require.WithinDuration(t, picture1.CreatedAt, picture2.CreatedAt, time.Second)
}

func TestUpdatePicture(t *testing.T) {
	picture1 := createRandomPicture(t)

	arg := UpdatePictureParams{
		ID:      picture1.ID,
		Rating: picture1.Rating - 15,
		Matches: picture1.Matches + 1,
	}

	picture2, err := testQueries.UpdatePicture(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, picture2)

	require.Equal(t, picture1.ID, picture2.ID)
	require.Equal(t, picture1.Description, picture2.Description)
	require.Equal(t, picture1.Url, picture2.Url)
	require.Equal(t, arg.Matches, picture2.Matches)
	require.Equal(t, arg.Rating, picture2.Rating)
	require.WithinDuration(t, picture1.CreatedAt, picture2.CreatedAt, time.Second)
}

func TestListPictures(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPicture(t)
	}

	arg := ListPicturesParams{
		Limit:  5,
		Offset: 5,
	}

	pictures, err := testQueries.ListPictures(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, pictures, 5)

	for _, picture := range pictures {
		require.NotEmpty(t, picture)
	}
}