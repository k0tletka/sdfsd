package handler

import "time"

type ServiceInfoResponse struct {
	VersionNumber string    `json:"version_number"`
	CommitHash    string    `json:"commit_hash"`
	BuildTime     time.Time `json:"build_time"`
}
