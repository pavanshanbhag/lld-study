package fooddeliveryservice

import "sync"

type Restaurant struct {
	ID      string
	Name    string
	Address string
	menu    []*MenuItem
	mu      sync.RWMutex
}

func NewRestaurant(id, name, address string, menu []*MenuItem) *Restaurant {
	return &Restaurant{
		ID:      id,
		Name:    name,
		Address: address,
		menu:    menu,
	}
}

func (r *Restaurant) AddMenuItem(item *MenuItem) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.menu = append(r.menu, item)
}

func (r *Restaurant) RemoveMenuItem(item *MenuItem) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, menuItem := range r.menu {
		if menuItem == item {
			r.menu = append(r.menu[:i], r.menu[i+1:]...)
			break
		}
	}
}

func (r *Restaurant) Menu() []*MenuItem {
	r.mu.RLock()
	defer r.mu.RUnlock()
	menuCopy := make([]*MenuItem, len(r.menu))
	copy(menuCopy, r.menu)
	return menuCopy
}

