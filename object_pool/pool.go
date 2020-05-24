package main

import (
	"errors"
	"fmt"
	"sync"
)

type pool struct {
	idles 		[]iPoolObject
	active 		[]iPoolObject
	maxObject 	int
	lock 		*sync.Mutex
}

func initPool(objects []iPoolObject) (*pool,error){
	if len(objects) == 0 {
		return nil, errors.New("could not find any object in pool list")
	}

	pool := &pool{
		idles: objects,
		active: make([]iPoolObject, 0),
		maxObject: len(objects),
		lock: new(sync.Mutex),
	}

	return pool, nil
}

func (p *pool)Loan() (iPoolObject, error){
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.idles) == 0{
		return nil, errors.New("could not find an idle object")
	}

	obj := p.idles[0]
	p.idles = p.idles[1:]
	p.active = append(p.active, obj)
	fmt.Println("loan an object with ID: ", obj.GetId())
	return obj, nil
}

func (p *pool)receive(obj iPoolObject) error{
	p.lock.Lock()
	defer p.lock.Unlock()

	sLen := len(p.active)
	for i, object := range p.active{
		if object.GetId() == obj.GetId(){
			p.active[sLen - 1], p.active[i] = p.active[i], p.active[sLen - 1]
			p.active = p.active[:sLen - 1]
			fmt.Println("object with ID: ", obj.GetId(), " was received")
			p.idles = append(p.idles, obj)
			return nil
		}
	}

	return errors.New("could not find object in this pool")
}