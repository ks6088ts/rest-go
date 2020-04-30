package service

import (
	"strconv"
	"testing"

	"github.com/ks6088ts/rest-go/pkg/entity"
	"github.com/ks6088ts/rest-go/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func assertEqual(t *testing.T, p, q *entity.Product) {
	assert.Equal(t, p.Description, q.Description)
	assert.Equal(t, p.ID, q.ID)
	assert.Equal(t, p.Name, q.Name)
}

func TestProduct(t *testing.T) {
	session, err := repository.NewMockSession()
	if err != nil {
		t.Error(err)
	}

	s, err := NewService(session)
	if err != nil {
		t.Error(err)
	}

	s.session.MigrateProduct()

	ps := []entity.Product{
		{
			ID:          0,
			Description: "dummy description0",
			Name:        "dummy name0",
		},
		{
			ID:          1,
			Description: "dummy description1",
			Name:        "dummy name1",
		},
	}

	for _, p := range ps {
		if err = s.session.CreateProduct(&p); err != nil {
			t.Error(err)
		}
	}

	for _, p := range ps {
		id := strconv.Itoa(int(p.ID))
		actual, err := s.session.ReadProduct(id)
		if err != nil {
			t.Error(err)
		}
		assertEqual(t, &p, actual)
	}

	products, err := s.session.ReadProducts()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, len(ps), len(products))
	for i, p := range products {
		assertEqual(t, &p, &ps[i])
	}

	s.session.Close()
}
