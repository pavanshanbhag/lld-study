package linkedin

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	service := NewLinkedInService()

	user1 := NewUser("1", "John Doe", "john@example.com", "password")
	user2 := NewUser("2", "Jane Smith", "jane@example.com", "password")
	service.RegisterUser(user1)
	service.RegisterUser(user2)

	loggedInUser, err := service.LoginUser("john@example.com", "password")
	if err != nil {
		logger.Error("login failed", "err", err)
		return
	}
	logger.Info("user logged in", "name", loggedInUser.Name)

	loggedInUser.Profile.SetHeadline("Software Engineer")
	loggedInUser.Profile.SetSummary("Passionate about coding and problem-solving.")
	service.UpdateUserProfile(loggedInUser)

	service.SendConnectionRequest(user1, user2)
	service.AcceptConnectionRequest(user2, user1)

	service.PostJobListing(NewJobPosting("1", "Software Developer", "We are hiring!", []string{"Java", "Python"}, "San Francisco"))

	for _, user := range service.SearchUsers("John") {
		logger.Info("user search hit", "name", user.Name, "headline", user.Profile.Headline)
	}
	for _, posting := range service.SearchJobPostings("Software") {
		logger.Info("job search hit", "title", posting.Title, "description", posting.Description)
	}

	service.SendMessage(user1, user2, "Hi Jane, hope you're doing well!")

	for _, notification := range service.GetNotifications(user2.ID) {
		logger.Info("notification", "type", notification.Type, "content", notification.Content)
	}
}
