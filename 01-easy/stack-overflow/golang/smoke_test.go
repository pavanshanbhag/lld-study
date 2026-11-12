package stackoverflow

import "testing"

func TestNewStackOverflow(t *testing.T) {
	t.Parallel()
	so := NewStackOverflow()
	user := so.CreateUser("alice", "alice@example.com")
	if user == nil {
		t.Fatal("expected user")
	}
}

func TestAskVoteAndAccept(t *testing.T) {
	t.Parallel()
	so := NewStackOverflow()
	alice := so.CreateUser("alice", "alice@example.com")
	bob := so.CreateUser("bob", "bob@example.com")

	q, err := so.AskQuestion(alice, "What is Go?", "Explain goroutines", []string{"go"})
	if err != nil {
		t.Fatalf("ask: %v", err)
	}
	a, err := so.AnswerQuestion(bob, q, "Goroutines are lightweight threads")
	if err != nil {
		t.Fatalf("answer: %v", err)
	}
	if err := so.VoteAnswer(alice, a, 1); err != nil {
		t.Fatalf("vote: %v", err)
	}
	if err := so.AcceptAnswer(a); err != nil {
		t.Fatalf("accept: %v", err)
	}
	if a.VoteCount() != 1 {
		t.Fatalf("votes = %d, want 1", a.VoteCount())
	}
	if !a.IsAccepted() {
		t.Fatal("expected accepted answer")
	}
	if bob.Reputation() < AnswerReputation {
		t.Fatalf("bob reputation = %d, want >= %d", bob.Reputation(), AnswerReputation)
	}
}

func TestSearchByTag(t *testing.T) {
	t.Parallel()
	so := NewStackOverflow()
	u := so.CreateUser("u", "u@example.com")
	_, _ = so.AskQuestion(u, "Java Q", "content", []string{"java"})
	results := so.SearchQuestions("java")
	if len(results) != 1 {
		t.Fatalf("results = %d, want 1", len(results))
	}
}
