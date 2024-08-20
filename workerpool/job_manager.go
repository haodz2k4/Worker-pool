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

func (jm *JobManager) SetStatus(job *Job, status JobStatus) {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	jm.jobStatus[job] = status
}

func (jm *JobManager) GetStatus(job *Job) JobStatus {
	jm.mu.Lock()
	defer jm.mu.Unlock()
	return jm.jobStatus[job]
}
