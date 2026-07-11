package onlineauctionsystem

import (
	"testing"
	"time"
)

func TestAuctionSystemPlaceBid(t *testing.T) {
	t.Parallel()

	system := NewAuctionSystem()
	seller := NewUser("1", "Seller", "seller@example.com")
	bidder := NewUser("2", "Bidder", "bidder@example.com")
	system.RegisterUser(seller)
	system.RegisterUser(bidder)

	listing := NewAuctionListing("1", "Watch", "Vintage watch", 100.0, time.Hour, seller)
	system.CreateAuctionListing(listing)

	bid := NewBid("b1", bidder, 150.0)
	if !system.PlaceBid(listing.ID, bid) {
		t.Fatal("expected bid to succeed")
	}
	if listing.CurrentHighestBid != 150.0 {
		t.Fatalf("highest bid = %v, want 150", listing.CurrentHighestBid)
	}
}
