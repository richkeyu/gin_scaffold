package client

import (
	"github.com/stretchr/testify/assert"
	"pay/ptest"
	"testing"
)

func TestBaseCli_GetProjectById(t *testing.T) {
	m, err := NewBaseCli().GetProjectById(ptest.GetCtx(), 1)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestBaseCli_GetMerchant(t *testing.T) {
	m, err := NewBaseCli().GetMerchant(ptest.GetCtx(), 54)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}
func TestBaseCli_GetProjectByIdWithCache(t *testing.T) {
	m, err := NewBaseCli().GetProjectByIdWithCache(ptest.GetCtx(), 1)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestBaseCli_GetProjectsWithCache(t *testing.T) {
	m, err := NewBaseCli().GetProjectsWithCache(ptest.GetCtx())
	assert.Nil(t, err)
	assert.NotNil(t, m)
}
func TestBaseCli_GetMerchantWithCache(t *testing.T) {
	m, err := NewBaseCli().GetMerchantWithCache(ptest.GetCtx(), 54)
	assert.Nil(t, err)
	assert.NotNil(t, m)
	m, err = NewBaseCli().GetMerchantWithCache(ptest.GetCtx(), 54)
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestBaseCli_GetMerchantAllWithCache(t *testing.T) {
	m, err := NewBaseCli().GetMerchantAllWithCache(ptest.GetCtx())
	assert.Nil(t, err)
	assert.NotNil(t, m)
	m, err = NewBaseCli().GetMerchantAllWithCache(ptest.GetCtx())
	assert.Nil(t, err)
	assert.NotNil(t, m)
}

func TestBaseCli_GetCountriesAllWithCache(t *testing.T) {
	m, err := NewBaseCli().GetCountriesAllWithCache(ptest.GetCtx())
	assert.Nil(t, err)
	assert.NotNil(t, m)
	m, err = NewBaseCli().GetCountriesAllWithCache(ptest.GetCtx())
	assert.Nil(t, err)
	assert.NotNil(t, m)
}
