package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
	Crear un servidor web, Generar job con un worker pool estara consumiendo
	Mejorar fuente con dependencias externas
*/

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatch struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d started\n", w.id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d finished with result %d \n", w.id, fib)

			case <-w.QuitChan:
				fmt.Printf("Worker with %d stopped\n", w.id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatch {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatch{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

func (d *Dispatch) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				WorkerJobQueue := <-d.WorkerPool
				WorkerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatch) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}

// Server
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))

	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest) // La data enviada no es la correcta
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))

	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest) // La data enviada no es la correcta
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest) // La data enviada no es la correcta
		return
	}

	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()
	//http:localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})

	// Bloquea el programa cuando encuentre un error
	log.Fatal(http.ListenAndServe(port, nil))
}
