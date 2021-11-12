package example1

import (
	"errors"
	"fmt"
	"sync"
)

type pool struct {
	idle     []iPoolobjecter
	active   []iPoolobjecter
	capacity int
	mux      *sync.Mutex
}

func initPool(poolObjects []iPoolobjecter) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, errors.New("can not create a pool of 0 len")
	}
	active := make([]iPoolobjecter, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mux:      new(sync.Mutex),
	}
	return pool, nil
}

func (p *pool) loan() (iPoolobjecter, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	if len(p.idle) == 0 {
		return nil, errors.New("no pool object free")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Println("loan pool object with ID:", obj.getID())
	return obj, nil
}

func (p *pool) receive(target iPoolobjecter) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Println("return pool object with ID:", target.getID())
	return nil
}

func (p *pool) remove(target iPoolobjecter) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("targe pool object doesn't belong to the pool")
}
