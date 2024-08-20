package workerpool

type Job struct {
	Priority int
	Process  func()
}

func newJob(priority int, process func()) *Job {
	return &Job{
		Priority: priority,
		Process:  process,
	}
}

func (i *Job) Execute() {
	i.Process()
}
