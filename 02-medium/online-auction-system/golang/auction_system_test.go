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
	if err := system.PlaceBid(listing.ID, bid); err != nil {
		t.Fatalf("PlaceBid: %v", err)
	}
	if listing.CurrentHighestBid != 150.0 {
		t.Fatalf("highest bid = %v, want 150", listing.CurrentHighestBid)
	}

	if err := system.PlaceBid(listing.ID, NewBid("b2", bidder, 120.0)); err == nil {
		t.Fatal("expected error for low bid")
	}
}
