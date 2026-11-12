package librarymanagementsystem

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	libraryManager := NewLibraryManager()

	libraryManager.AddBook(NewBook("ISBN1", "Book 1", "Author 1", 2020))
	libraryManager.AddBook(NewBook("ISBN2", "Book 2", "Author 2", 2019))
	libraryManager.AddBook(NewBook("ISBN3", "Book 3", "Author 3", 2021))

	libraryManager.RegisterMember(NewMember("M1", "John Doe", "john@example.com"))
	libraryManager.RegisterMember(NewMember("M2", "Jane Smith", "jane@example.com"))

	if err := libraryManager.BorrowBook("M1", "ISBN1"); err != nil {
		logger.Error("borrow failed", "err", err)
	}
	if err := libraryManager.BorrowBook("M2", "ISBN2"); err != nil {
		logger.Error("borrow failed", "err", err)
	}
	if err := libraryManager.ReturnBook("M1", "ISBN1"); err != nil {
		logger.Error("return failed", "err", err)
	}

	for _, book := range libraryManager.SearchBooks("Book") {
		logger.Info("search hit", "title", book.Title, "author", book.Author)
	}
}
