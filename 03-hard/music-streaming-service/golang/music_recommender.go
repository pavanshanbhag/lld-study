package musicstreamingservice

import (
	"sync"
)

type MusicRecommender struct {
	userRecommendations map[string][]*Song
	mu                  sync.RWMutex
}

func NewMusicRecommender() *MusicRecommender {
	return &MusicRecommender{
		userRecommendations: make(map[string][]*Song),
	}
}

func (mr *MusicRecommender) RecommendSongs(user *User) []*Song {
	mr.mu.RLock()
	defer mr.mu.RUnlock()

	if recommendations, exists := mr.userRecommendations[user.ID]; exists {
		return recommendations
	}
	return make([]*Song, 0)
}
