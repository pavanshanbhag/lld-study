package socialnetworkingservice

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	socialNetwork := NewSocialNetwork()

	user1 := NewUser("1", "John Doe", "john@example.com", "password", "profile1.jpg", "I love coding!")
	user2 := NewUser("2", "Jane Smith", "jane@example.com", "password", "profile2.jpg", "Exploring the world!")

	if err := socialNetwork.RegisterUser(user1); err != nil {
		logger.Error("register failed", "err", err)
		return
	}
	if err := socialNetwork.RegisterUser(user2); err != nil {
		logger.Error("register failed", "err", err)
		return
	}

	loggedInUser, err := socialNetwork.LoginUser("john@example.com", "password")
	if err != nil {
		logger.Error("login failed", "err", err)
		return
	}
	logger.Info("user logged in", "name", loggedInUser.Name)

	if err := socialNetwork.SendFriendRequest(user1.ID, user2.ID); err != nil {
		logger.Error("friend request failed", "err", err)
		return
	}
	if err := socialNetwork.AcceptFriendRequest(user2.ID, user1.ID); err != nil {
		logger.Error("accept friend failed", "err", err)
		return
	}

	post1 := NewPost("post1", user1.ID, "My first post!", []string{}, []string{})
	post2 := NewPost("post2", user2.ID, "Having a great day!", []string{}, []string{})
	if err := socialNetwork.CreatePost(post1); err != nil {
		logger.Error("create post failed", "err", err)
		return
	}
	if err := socialNetwork.CreatePost(post2); err != nil {
		logger.Error("create post failed", "err", err)
		return
	}

	if err := socialNetwork.LikePost(user2.ID, post1.ID); err != nil {
		logger.Error("like failed", "err", err)
		return
	}
	if err := socialNetwork.CommentOnPost(NewComment("comment1", user2.ID, post1.ID, "Great post!")); err != nil {
		logger.Error("comment failed", "err", err)
		return
	}

	newsfeed, err := socialNetwork.GetNewsfeed(user1.ID)
	if err != nil {
		logger.Error("newsfeed failed", "err", err)
		return
	}
	for _, post := range newsfeed {
		logger.Info("newsfeed post", "content", post.Content, "likes", len(post.GetLikes()), "comments", len(post.GetComments()))
	}

	notifications, err := socialNetwork.GetNotifications(user1.ID)
	if err != nil {
		logger.Error("notifications failed", "err", err)
		return
	}
	for _, notification := range notifications {
		logger.Info("notification", "type", notification.Type, "content", notification.Content)
	}
}
