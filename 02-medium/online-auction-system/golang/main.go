package onlineauctionsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	auctionSystem := NewAuctionSystem()

	user1 := NewUser("1", "John Doe", "john@example.com")
	user2 := NewUser("2", "Jane Smith", "jane@example.com")
	auctionSystem.RegisterUser(user1)
	auctionSystem.RegisterUser(user2)

	listing1 := NewAuctionListing("1", "Item 1", "Description 1", 100.0, 2*time.Second, user1)
	listing2 := NewAuctionListing("2", "Item 2", "Description 2", 50.0, 3*time.Second, user2)
	auctionSystem.CreateAuctionListing(listing1)
	auctionSystem.CreateAuctionListing(listing2)

	for _, listing := range auctionSystem.SearchAuctionListings("Item") {
		logger.Info("listing found", "name", listing.ItemName)
	}

	if err := auctionSystem.PlaceBid(listing1.ID, NewBid("1", user2, 150.0)); err != nil {
		logger.Error("bid failed", "err", err)
	}
	if err := auctionSystem.PlaceBid(listing1.ID, NewBid("2", user1, 200.0)); err != nil {
		logger.Error("bid failed", "err", err)
	}

	time.Sleep(2500 * time.Millisecond)
	logger.Info("auction status", "id", listing1.ID, "status", listing1.Status)
}
