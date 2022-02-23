package pool

import "errors"

type DBConn struct {
	address string
}

func NewDBConn(addr string) *DBConn {
	return &DBConn{addr}
}

////////////////////////

const (
	DefaultAddress  = "localhost"
	DefaultMinConns = 2
	DefaultMaxConns = 4

	MinConns = 1
	MaxConns = 1028
)

var ErrZeroConnectionsAvailable = errors.New("zero connections available")

type Pool struct {
	Address            string
	MinConns, MaxConns int
	availableConns     chan *DBConn
	bufferConns        chan struct{}
}

type Option func(*Pool)

func WithAddress(addr string) Option {
	return func(p *Pool) {
		p.Address = addr
	}
}

func WithMinConns(min int) Option {
	if min < MinConns {
		min = MinConns
	} else if min > MaxConns {
		min = MaxConns
	}
	return func(p *Pool) {
		p.MinConns = min
	}
}

func WithMaxConns(max int) Option {
	if max < MinConns {
		max = MinConns
	} else if max > MaxConns {
		max = MaxConns
	}
	return func(p *Pool) {
		p.MaxConns = max
	}
}

func New(options ...Option) Pool {
	p := Pool{
		Address:  DefaultAddress,
		MinConns: DefaultMinConns,
		MaxConns: DefaultMaxConns,
	}
	for _, opt := range options {
		opt(&p)
	}
	if p.MaxConns < p.MinConns {
		p.MaxConns = p.MinConns
	}

	p.availableConns = make(chan *DBConn, p.MinConns)
	for i := 0; i < p.MinConns; i++ {
		p.availableConns <- NewDBConn(p.Address)
	}

	buffer := p.MaxConns - p.MinConns
	p.bufferConns = make(chan struct{}, buffer)
	for i := 0; i < buffer; i++ {
		p.bufferConns <- struct{}{}
	}

	return p
}

func (p Pool) CheckIn(conn *DBConn) {
	p.availableConns <- conn
}

func (p Pool) CheckOut() (*DBConn, error) {
	select {
	case c := <-p.availableConns:
		return c, nil
	case _ = <-p.bufferConns:
		return NewDBConn(p.Address), nil
	default:
		return nil, ErrZeroConnectionsAvailable
	}
}
