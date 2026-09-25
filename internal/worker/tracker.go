package worker

import "sync"

// Tracker collects error messages from concurrent jobs.
type Tracker struct {
	errs []string
	wg   sync.WaitGroup
}

// Run executes each job in its own goroutine and records failures.
func (t *Tracker) Run(jobs []func() error) []string {
	for _, job := range jobs {
		t.wg.Add(1)
		go func(j func() error) {
			defer t.wg.Done()
			if err := j(); err != nil {
				t.errs = append(t.errs, err.Error())
			}
		}(job)
	}
	t.wg.Wait()
	return t.errs
}
