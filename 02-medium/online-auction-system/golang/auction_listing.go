package onlineauctionsystem

import (
	"fmt"
	"sync"
	"time"
)

type AuctionListing struct {
	ID                   string
	ItemName             string
	Description          string
	StartingPrice        float64
	Duration             time.Duration
	Seller               *User
	Status               AuctionStatus
	CurrentHighestBid    float64
	CurrentHighestBidder *User
	Bids                 []*Bid
	mu                   sync.Mutex
	observers            []chan struct{}
}

func NewAuctionListing(id, itemName, description string, startingPrice float64, duration time.Duration, seller *User) *AuctionListing {
	return &AuctionListing{
		ID:                id,
		ItemName:          itemName,
		Description:       description,
		StartingPrice:     startingPrice,
		Duration:          duration,
		Seller:            seller,
		Status:            StatusActive,
		CurrentHighestBid: startingPrice,
		Bids:              make([]*Bid, 0),
		observers:         make([]chan struct{}, 0),
	}
}

func (a *AuctionListing) PlaceBid(bid *Bid) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Status != StatusActive {
		return fmt.Errorf("auction %s is not active", a.ID)
	}
	if bid.Amount <= a.CurrentHighestBid {
		return fmt.Errorf("bid %.2f must exceed current highest %.2f", bid.Amount, a.CurrentHighestBid)
	}

	a.CurrentHighestBid = bid.Amount
	a.CurrentHighestBidder = bid.Bidder
	a.Bids = append(a.Bids, bid)
	a.notifyObservers()
	return nil
}

func (a *AuctionListing) CloseAuction() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Status == StatusActive {
		a.Status = StatusClosed
		a.notifyObservers()
	}
}

func (a *AuctionListing) notifyObservers() {
	for _, observer := range a.observers {
		select {
		case observer <- struct{}{}:
		default:
		}
	}
}
