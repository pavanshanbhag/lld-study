package musicstreamingservice

import "testing"

func TestMusicStreamingServiceConstructor(t *testing.T) {
	t.Parallel()

	service := NewMusicStreamingService()
	if service.MusicLibrary == nil || service.UserManager == nil || service.MusicRecommender == nil {
		t.Fatal("expected all components to be initialized")
	}
}

func TestMusicLibrarySearchSongs(t *testing.T) {
	t.Parallel()

	library := NewMusicLibrary()
	library.AddSong(&Song{ID: "S1", Title: "One More Time", Artist: "Daft Punk", Album: "Discovery"})

	results := library.SearchSongs("daft")
	if len(results) != 1 || results[0].ID != "S1" {
		t.Fatalf("SearchSongs() = %+v, want one match", results)
	}
}

func TestUserManagerRegisterAndLogin(t *testing.T) {
	t.Parallel()

	manager := NewUserManager()
	user := NewUser("U1", "alice", "secret")
	manager.RegisterUser(user)

	if loggedIn := manager.LoginUser("alice", "secret"); loggedIn == nil {
		t.Fatal("expected successful login")
	}
}
