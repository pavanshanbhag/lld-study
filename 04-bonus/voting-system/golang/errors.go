package votingsystem

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExists    = errors.New("already exists")
	ErrInvalidState     = errors.New("invalid election state")
	ErrAlreadyVoted     = errors.New("voter already cast a ballot")
	ErrUnknownCandidate = errors.New("unknown candidate")
	ErrUnknownVoter     = errors.New("unknown voter")
)
