package stackoverflow

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	system := NewStackOverflow()

	alice := system.CreateUser("Alice", "alice@example.com")
	bob := system.CreateUser("Bob", "bob@example.com")
	charlie := system.CreateUser("Charlie", "charlie@example.com")
	logger.Info("users created", "alice", alice.Username, "bob", bob.Username, "charlie", charlie.Username)

	javaQuestion, err := system.AskQuestion(alice,
		"What is polymorphism in Java?",
		"Can someone explain polymorphism in Java with an example?",
		[]string{"java", "oop"})
	if err != nil {
		logger.Error("ask question failed", "err", err)
		return
	}

	bobAnswer, err := system.AnswerQuestion(bob, javaQuestion,
		"Polymorphism in Java is the ability of an object to take on many forms...")
	if err != nil {
		logger.Error("answer failed", "err", err)
		return
	}

	if _, err = system.AddComment(charlie, javaQuestion, "Great question!"); err != nil {
		logger.Error("comment failed", "err", err)
		return
	}
	if _, err = system.AddComment(alice, bobAnswer, "Thanks for the explanation!"); err != nil {
		logger.Error("comment failed", "err", err)
		return
	}
	if err := system.VoteQuestion(charlie, javaQuestion, 1); err != nil {
		logger.Error("vote failed", "err", err)
		return
	}
	if err := system.VoteAnswer(charlie, bobAnswer, 1); err != nil {
		logger.Error("vote failed", "err", err)
		return
	}
	if err := system.AcceptAnswer(bobAnswer); err != nil {
		logger.Error("accept failed", "err", err)
		return
	}

	pythonQuestion, err := system.AskQuestion(bob,
		"How to use list comprehensions in Python?",
		"Can someone explain list comprehensions?",
		[]string{"python", "list-comprehension"})
	if err != nil {
		logger.Error("ask question failed", "err", err)
		return
	}
	aliceAnswer, err := system.AnswerQuestion(alice, pythonQuestion,
		"List comprehensions provide a concise way to create lists...")
	if err != nil {
		logger.Error("answer failed", "err", err)
		return
	}
	_ = system.VoteQuestion(charlie, pythonQuestion, 1)
	_ = system.VoteAnswer(charlie, aliceAnswer, 1)

	logger.Info("java question",
		"title", javaQuestion.Title,
		"votes", javaQuestion.VoteCount(),
		"comments", len(javaQuestion.Comments()),
		"answer_votes", bobAnswer.VoteCount(),
		"accepted", bobAnswer.IsAccepted(),
	)
	logger.Info("reputations",
		"alice", alice.Reputation(),
		"bob", bob.Reputation(),
		"charlie", charlie.Reputation(),
	)
	for _, q := range system.SearchQuestions("java") {
		logger.Info("search hit", "query", "java", "title", q.Title)
	}
	for _, q := range system.QuestionsByUser(bob) {
		logger.Info("user question", "user", bob.Username, "title", q.Title)
	}
}
