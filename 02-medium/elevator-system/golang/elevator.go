package elevatorsystem

import (
	"fmt"
	"sync"
	"time"
)

type Elevator struct {
	id               int
	capacity         int
	currentFloor     int
	currentDirection Direction
	requests         chan *Request
	stopChan         chan struct{}
	mu               sync.RWMutex
}

func NewElevator(id, capacity int) *Elevator {
	return &Elevator{
		id:               id,
		capacity:         capacity,
		currentFloor:     1,
		currentDirection: DirectionUp,
		requests:         make(chan *Request, capacity),
		stopChan:         make(chan struct{}),
	}
}

func (e *Elevator) AddRequest(request *Request) error {
	select {
	case e.requests <- request:
		return nil
	default:
		return fmt.Errorf("elevator %d at capacity", e.id)
	}
}

func (e *Elevator) CurrentFloor() int {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.currentFloor
}

func (e *Elevator) setCurrentFloor(floor int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.currentFloor = floor
}

func (e *Elevator) Run() {
	go func() {
		for {
			select {
			case request := <-e.requests:
				e.processRequest(request)
			case <-e.stopChan:
				return
			}
		}
	}()
}

func (e *Elevator) Stop() {
	close(e.stopChan)
}

func (e *Elevator) processRequest(request *Request) {
	startFloor := e.CurrentFloor()
	endFloor := request.DestinationFloor

	if startFloor < endFloor {
		e.currentDirection = DirectionUp
		for i := startFloor; i <= endFloor; i++ {
			e.setCurrentFloor(i)
			time.Sleep(10 * time.Millisecond)
		}
	} else if startFloor > endFloor {
		e.currentDirection = DirectionDown
		for i := startFloor; i >= endFloor; i-- {
			e.setCurrentFloor(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
