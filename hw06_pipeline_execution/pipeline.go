package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	stageWorker := func(done In, out Out) Out {
		take := make(Bi)
		go func() {
			defer close(take)
			for {
				select {
				case <-done:
					return
				case value, ok := <-out:
					if !ok {
						return
					}
					take <- value
				}
			}
		}()
		return take
	}

	in = stageWorker(done, in)

	for _, stage := range stages {
		in = stageWorker(done, stage(in))
	}

	return in
}
