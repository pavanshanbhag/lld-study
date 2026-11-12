package musicstreamingservice

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewMusicPlayer()
	_ = NewMusicStreamingService()
}
