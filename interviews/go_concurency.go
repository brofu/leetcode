package interviews

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Implement a worker pool that processes a list of integers by squaring them concurrently, then collects results in order.

Input:  []int{1, 2, 3, 4, 5}
Output: []int{1, 4, 9, 16, 25}

Input:  []int{}
Output: []int{}

Input:  []int{-1, -2, 3}
Output: []int{1, 4, 9}
*/

type Job interface {
	JobId() int
	GetInput() int
	SetResult(int)
	GetResult() int
}

type Worker interface {
	Work(context.Context, <-chan Job, chan<- Job)
}

type Dispatcher struct {
	workerNumber int
	workers      []Worker
	jobs         []Job
	results      []int
}

func NewDispatcher(jobs []Job, workerNumber int) *Dispatcher {

	workers := make([]Worker, workerNumber)
	d := &Dispatcher{
		workerNumber: workerNumber,
		workers:      workers,
		jobs:         jobs,
		results:      make([]int, len(jobs)),
	}
	return d
}

func (d *Dispatcher) Start(ctx context.Context) {

	input := make(chan Job, d.workerNumber)
	output := make(chan Job, d.workerNumber)

	for _, worker := range d.workers {
		//Log here
		go worker.Work(ctx, input, output)
	}

	// dispatch the job
	for _, job := range d.jobs {
		input <- job
	}
	close(input)

	// collect the results
	finished := 0
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-output:
			d.results[job.JobId()] = job.GetResult()
			finished++
		default:
		}
		if finished == len(d.jobs) {
			return
		}
	}
}

func (d *Dispatcher) Results(ctx context.Context) []int {
	return d.results
}

type DefalutJob struct {
	number int
	idx    int
	result int
}

func (j *DefalutJob) GetResult() int {
	return j.result
}

func (j *DefalutJob) SetResult(r int) {
	j.result = r
}

func (j *DefalutJob) GetInput() int {
	return j.number
}

func (j *DefalutJob) JobId() int {
	return j.idx
}

type DefaultWorker struct {
}

func (w *DefaultWorker) Work(ctx context.Context, input <-chan Job, output chan<- Job) { // Result abstraction later

	for {
		select {
		case <-ctx.Done():
			// Log later
			return
		case job, ok := <-input: // a separeate function for business logic
			if !ok { // closed
				//Log later
				return
			}
			num := job.GetInput()
			job.SetResult(num * num)
			output <- job
		}
	}
}

const (
	WorkerNumber = 2
)

func main() {

	input := []int{}

	ctx := context.Background()

	jobs := make([]Job, len(input))
	for idx := range jobs {
		jobs[idx] = &DefalutJob{
			number: input[idx],
			idx:    idx,
		}
	}

	dp := NewDispatcher(jobs, WorkerNumber)

	ctx, cancelFunc := context.WithCancel(ctx)
	go dp.Start(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	for {
		select {
		case <-sig:
			// Log
			cancelFunc()
		}
	}

	//TODO: use context to cancel
	results := dp.Results(ctx)

	fmt.Println("flag", results)
}
