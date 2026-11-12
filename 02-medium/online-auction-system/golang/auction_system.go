package onlineauctionsystem

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type AuctionSystem struct {
	users           map[string]*User
	auctionListings map[string]*AuctionListing
	mu              sync.RWMutex
}

func NewAuctionSystem() *AuctionSystem {
	return &AuctionSystem{
		users:           make(map[string]*User),
		auctionListings: make(map[string]*AuctionListing),
	}
}

func (as *AuctionSystem) RegisterUser(user *User) {
	as.mu.Lock()
	defer as.mu.Unlock()
	as.users[user.ID] = user
}

func (as *AuctionSystem) User(userID string) *User {
	as.mu.RLock()
	defer as.mu.RUnlock()
	return as.users[userID]
}

func (as *AuctionSystem) CreateAuctionListing(listing *AuctionListing) {
	as.mu.Lock()
	as.auctionListings[listing.ID] = listing
	as.mu.Unlock()

	// Start auction timer
	go func() {
		time.Sleep(listing.Duration)
		listing.CloseAuction()
	}()
}

func (as *AuctionSystem) SearchAuctionListings(keyword string) []*AuctionListing {
	as.mu.RLock()
	defer as.mu.RUnlock()

	var matchingListings []*AuctionListing
	keyword = strings.ToLower(keyword)

	for _, listing := range as.auctionListings {
		if strings.Contains(strings.ToLower(listing.ItemName), keyword) ||
			strings.Contains(strings.ToLower(listing.Description), keyword) {
			matchingListings = append(matchingListings, listing)
		}
	}
	return matchingListings
}

func (as *AuctionSystem) PlaceBid(auctionListingID string, bid *Bid) error {
	as.mu.RLock()
	listing, exists := as.auctionListings[auctionListingID]
	as.mu.RUnlock()

	if !exists {
		return fmt.Errorf("auction listing %s not found", auctionListingID)
	}
	return listing.PlaceBid(bid)
}
