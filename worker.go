package pooler

type Work struct {
	ID  int
	Job string
	JS  JobSpec
}

type Worker interface {
	Start()
	Stop()
}

type worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool

	JS JobSpec
}

func (w *worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				// do work
				DoWork(job.Job, w.ID)
			case <-w.End:
				return
			}
		}
	}()
}

func (w *worker) Stop() {
	dbLogf("worker [%d] is stopping", w.ID)
	w.End <- true
}
