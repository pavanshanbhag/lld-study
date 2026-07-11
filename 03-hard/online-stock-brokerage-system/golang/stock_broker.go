package onlinestockbrokeragesystem

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type StockBroker struct {
    accounts        map[string]*Account
    stocks          map[string]*Stock
    orderQueue      chan Order
    accountIDCount  int64
    mu             sync.RWMutex
}

func NewStockBroker() *StockBroker {
    sb := &StockBroker{
        accounts:   make(map[string]*Account),
        stocks:     make(map[string]*Stock),
        orderQueue: make(chan Order, 100),
    }
    go sb.processOrders()
    return sb
}

func (sb *StockBroker) CreateAccount(user *User, initialBalance float64) {
    sb.mu.Lock()
    defer sb.mu.Unlock()

    accountID := sb.generateAccountID()
    account := NewAccount(accountID, user, initialBalance)
    sb.accounts[accountID] = account
}

func (sb *StockBroker) GetAccount(accountID string) *Account {
    sb.mu.RLock()
    defer sb.mu.RUnlock()
    return sb.accounts[accountID]
}

func (sb *StockBroker) AddStock(stock *Stock) {
    sb.mu.Lock()
    defer sb.mu.Unlock()
    sb.stocks[stock.Symbol] = stock
}

func (sb *StockBroker) GetStock(symbol string) *Stock {
    sb.mu.RLock()
    defer sb.mu.RUnlock()
    return sb.stocks[symbol]
}

func (sb *StockBroker) PlaceOrder(order Order) {
    sb.orderQueue <- order
}

func (sb *StockBroker) processOrders() {
    for order := range sb.orderQueue {
        if err := order.Execute(); err != nil {
            fmt.Printf("Order failed: %v\n", err)
        }
    }
}

func (sb *StockBroker) generateAccountID() string {
    id := atomic.AddInt64(&sb.accountIDCount, 1)
    return fmt.Sprintf("A%03d", id)
}