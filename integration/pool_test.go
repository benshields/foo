package main

import (
	"testing"

	"github.com/benshields/foo/pool"
	"github.com/stretchr/testify/assert"
)

func TestPool(t *testing.T) {
	TestNew(t)
	TestCheckOut(t)
	TestCheckIn(t)
	TestInterlaceCheckOutCheckIn(t)
}

func TestNew(t *testing.T) {
	tests := []struct {
		options  []pool.Option
		expected pool.Pool
	}{
		{
			options: nil,
			expected: pool.Pool{
				Address:  pool.DefaultAddress,
				MinConns: pool.DefaultMinConns,
				MaxConns: pool.DefaultMaxConns,
			},
		},
		{
			options: []pool.Option{
				pool.WithAddress("atlas.mongodb.com"),
				pool.WithMaxConns(32),
			},
			expected: pool.Pool{
				Address:  "atlas.mongodb.com",
				MinConns: pool.DefaultMinConns,
				MaxConns: 32,
			},
		},
		{
			options: []pool.Option{
				pool.WithMinConns(3),
			},
			expected: pool.Pool{
				Address:  pool.DefaultAddress,
				MinConns: 3,
				MaxConns: pool.DefaultMaxConns,
			},
		},
		{
			options: []pool.Option{
				pool.WithMinConns(pool.MinConns - 1),
				pool.WithMaxConns(pool.MaxConns + 1),
			},
			expected: pool.Pool{
				Address:  pool.DefaultAddress,
				MinConns: pool.MinConns,
				MaxConns: pool.MaxConns,
			},
		},
		{
			options: []pool.Option{
				pool.WithMinConns(pool.MinConns - 1),
				pool.WithMaxConns(pool.MinConns - 2),
			},
			expected: pool.Pool{
				Address:  pool.DefaultAddress,
				MinConns: pool.MinConns,
				MaxConns: pool.MinConns,
			},
		},
		{
			options: []pool.Option{
				pool.WithMinConns(pool.MaxConns + 1),
				pool.WithMaxConns(pool.MinConns - 1),
			},
			expected: pool.Pool{
				Address:  pool.DefaultAddress,
				MinConns: pool.MaxConns,
				MaxConns: pool.MaxConns,
			},
		},
	}

	for _, test := range tests {
		p := pool.New(test.options...)
		assert.Equal(t, test.expected.Address, p.Address)
		assert.Equal(t, test.expected.MinConns, p.MinConns)
		assert.Equal(t, test.expected.MaxConns, p.MaxConns)
	}
}

func TestCheckOut(t *testing.T) {
	p := pool.New()

	for i := 0; i < pool.DefaultMaxConns; i++ {
		con, err := p.CheckOut()
		assert.NoError(t, err)
		assert.NotNil(t, con)
	}

	con, err := p.CheckOut()
	assert.ErrorIs(t, err, pool.ErrZeroConnectionsAvailable)
	assert.Nil(t, con)
}

func TestCheckIn(t *testing.T) {
	p := pool.New()

	con, err := p.CheckOut()
	assert.NoError(t, err)
	assert.NotNil(t, con)

	p.CheckIn(con)
}

func TestInterlaceCheckOutCheckIn(t *testing.T) {
	p := pool.New()

	var con *pool.DBConn
	var err error
	for i := 0; i < pool.DefaultMaxConns; i++ {
		con, err = p.CheckOut()
		assert.NoError(t, err)
		assert.NotNil(t, con)
	}
	con1 := con

	con, err = p.CheckOut()
	assert.ErrorIs(t, err, pool.ErrZeroConnectionsAvailable)
	assert.Nil(t, con)

	p.CheckIn(con1)

	for i := 0; i < 100; i++ {
		con, err = p.CheckOut()
		assert.NoError(t, err)
		assert.NotNil(t, con)

		p.CheckIn(con)
	}
}
