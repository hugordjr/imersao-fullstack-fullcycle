package model_test

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewUser(t *testing.T) {
	name := "My Name"
	email := "email@email.com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotNil(t, user)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))

	_, err = model.NewUser("", email)
	require.NotNil(t, err)

	_, err = model.NewUser(name, " ")
	require.NotNil(t, err)

}
