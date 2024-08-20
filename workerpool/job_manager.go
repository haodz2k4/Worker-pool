package workerpool

import "sync"

type JobStatus int

const (
	Pending JobStatus = iota
	Processing
	Completed
)

type JobManager struct {
	jobStatus map[*Job]JobStatus
	mu        sync.Mutex
}

func NewJobManager() *JobManager {
	return &JobManager{
		jobStatus: make(map[*Job]JobStatus),
	}
}
