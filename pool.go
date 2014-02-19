package beanstalk

import (
	"math/rand"
	"time"
)

type Pool struct {
	conns []*Conn
}

func NewPool(servers []string) *Pool {
	p := new(Pool)
	for _, server := range servers {
		c := NewConn(server)
		p.conns = append(p.conns, c)
	}

	for i, _ := range p.conns {
		j := rand.Intn(i + 1)
		p.conns[i], p.conns[j] = p.conns[j], p.conns[i]
	}
	return p
}

func (p *Pool) Use(name string) error {
	if p == nil {
		return PoolError{"Use", ErrNilPool}
	}
	if err := checkName(name); err != nil {
		return err
	}
	for _, conn := range p.conns {
		conn.Tube = Tube{conn, name}
	}
	return nil
}

func (p *Pool) Put(body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error) {
	if p == nil {
		return 0, PoolError{"Put", ErrNilPool}
	}
	for _, conn := range p.conns {
		id, err = conn.Put(body, pri, delay, ttr)
		if err != nil {
			p.conns = append(p.conns[1:], p.conns[:1]...)
		} else {
			return id, nil
		}
	}
	return 0, err
}
