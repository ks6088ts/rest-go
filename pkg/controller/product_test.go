package controller

import (
	"testing"

	"github.com/ks6088ts/rest-go/pkg/repository"
	"github.com/ks6088ts/rest-go/pkg/service"
)

func TestProduct(t *testing.T) {
	session, err := repository.NewMockSession()
	if err != nil {
		t.Error(err)
	}

	service, err := service.NewService(session)
	if err != nil {
		t.Error(err)
	}

	_, err = NewController(service)
	if err != nil {
		t.Error(err)
	}
}
