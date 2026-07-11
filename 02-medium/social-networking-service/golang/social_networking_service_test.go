package socialnetworkingservice

import "testing"

func TestSocialNetworkRegisterAndCreatePost(t *testing.T) {
	t.Parallel()

	sn := NewSocialNetwork()
	user := NewUser("1", "Alice", "alice@example.com", "pass", "", "")
	if err := sn.RegisterUser(user); err != nil {
		t.Fatalf("RegisterUser: %v", err)
	}

	post := NewPost("p1", user.ID, "Hello world", nil, nil)
	if err := sn.CreatePost(post); err != nil {
		t.Fatalf("CreatePost: %v", err)
	}

	newsfeed, err := sn.GetNewsfeed(user.ID)
	if err != nil {
		t.Fatalf("GetNewsfeed: %v", err)
	}
	if len(newsfeed) != 1 || newsfeed[0].Content != "Hello world" {
		t.Fatalf("GetNewsfeed() = %+v, want one post", newsfeed)
	}
}

func TestSocialNetworkFriendRequest(t *testing.T) {
	t.Parallel()

	sn := NewSocialNetwork()
	user1 := NewUser("1", "Alice", "alice@example.com", "pass", "", "")
	user2 := NewUser("2", "Bob", "bob@example.com", "pass", "", "")
	sn.RegisterUser(user1)
	sn.RegisterUser(user2)

	if err := sn.SendFriendRequest(user1.ID, user2.ID); err != nil {
		t.Fatalf("SendFriendRequest: %v", err)
	}
	if err := sn.AcceptFriendRequest(user2.ID, user1.ID); err != nil {
		t.Fatalf("AcceptFriendRequest: %v", err)
	}
}
