package librarymanagementsystem

import "testing"

func TestLibraryManagerBorrowAndReturn(t *testing.T) {
	t.Parallel()

	library := NewLibraryManager()
	book := NewBook("ISBN1", "Clean Code", "Robert Martin", 2008)
	member := NewMember("M1", "Alice", "alice@example.com")

	library.AddBook(book)
	library.RegisterMember(member)

	if err := library.BorrowBook("M1", "ISBN1"); err != nil {
		t.Fatalf("BorrowBook: %v", err)
	}
	if book.IsAvailable() {
		t.Fatal("book should be unavailable after borrow")
	}

	if err := library.BorrowBook("M1", "ISBN1"); err == nil {
		t.Fatal("expected error borrowing unavailable book")
	}

	if err := library.ReturnBook("M1", "ISBN1"); err != nil {
		t.Fatalf("ReturnBook: %v", err)
	}
	if !book.IsAvailable() {
		t.Fatal("book should be available after return")
	}
}

func TestLibraryManagerSearchBooks(t *testing.T) {
	t.Parallel()

	library := NewLibraryManager()
	library.AddBook(NewBook("ISBN1", "The Hobbit", "Tolkien", 1937))
	library.AddBook(NewBook("ISBN2", "Dune", "Herbert", 1965))

	results := library.SearchBooks("hobbit")
	if len(results) != 1 || results[0].ISBN != "ISBN1" {
		t.Fatalf("SearchBooks() = %+v, want one hobbit match", results)
	}
}

func TestLibraryManagerMaxBooksPerMember(t *testing.T) {
	t.Parallel()

	library := NewLibraryManager()
	member := NewMember("M1", "Alice", "alice@example.com")
	library.RegisterMember(member)

	for i := range maxBooksPerMember {
		isbn := string(rune('A' + i))
		library.AddBook(NewBook(isbn, "Book", "Author", 2020))
		if err := library.BorrowBook("M1", isbn); err != nil {
			t.Fatalf("BorrowBook %s: %v", isbn, err)
		}
	}

	library.AddBook(NewBook("EXTRA", "Extra", "Author", 2020))
	if err := library.BorrowBook("M1", "EXTRA"); err == nil {
		t.Fatal("expected error when exceeding max books per member")
	}
}
