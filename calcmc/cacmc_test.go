package calcmc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceNew(t *testing.T) {
	cc, err := NewCalcMQ(WithAltura(10), WithBase(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
}

func TestOptionsParamsAltura(t *testing.T) {
	cc, err := NewCalcMQ(WithAltura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetAltura())
}

func TestOptionsParamsBase(t *testing.T) {
	cc, err := NewCalcMQ(WithBase(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetBase())
}

func TestCalculaMQ(t *testing.T) {
	cc, err := NewCalcMQ(WithAltura(10), WithBase(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 100.0, cc.CalculaMQ())
}

func TestCalculaMQWithError(t *testing.T) {
	cc, err := NewCalcMQ(WithAltura(0), WithBase(10))
	assert.Error(t, err)
	assert.Nil(t, cc)

	cc, err = NewCalcMQ(WithAltura(10), WithBase(0))
	assert.Error(t, err)
	assert.Nil(t, cc)
}

func TestSetAltura(t *testing.T) {
	cc, err := NewCalcMQ(WithAltura(10), WithBase(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetAltura(20)
	assert.Equal(t, 20.0, cc.GetAltura())
}

func TestSetBase(t *testing.T) {
	cc, err := NewCalcMQ(WithAltura(10), WithBase(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetBase(20)
	assert.Equal(t, 20.0, cc.GetBase())
}
