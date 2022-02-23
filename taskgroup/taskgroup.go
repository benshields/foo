// Implement TaskGroup, a type that represents a collection of goroutines, each of which can return an error. Users should be able to add goroutines to the TaskGroup with Run. Users should be able to use Wait to wait until one of the running goroutines returns a non-nil error, at which point the TaskGroup should stop all other routines and surface the error.



package taskgroup

// A TaskGroup is a collection of goroutines.
type TaskGroup struct {
  // num of tasks, default & allow override
  numTasks int
  
  // chan for executing tasks, unneeded?
  
  // chan for cancelling running tasks
  // chan for receiving task results
  results, cancel chan error
}

// NewTaskGroup returns a new TaskGroup.
func NewTaskGroup( /* params */ numTasks int ) *TaskGroup {
  // set num, make cancel/results with matching buffer size
  tg := TaskGroup{}
  tg.numTasks = numTasks
  tg.results = make(chan error, numTasks)
  tg.cancel = make(chan error)
  return &tg
}

// Run calls the given function f in a new goroutine.
func (tg *TaskGroup) Run(f func( /* params */ ) error) {
  // sending on results chan
  // checking cancel chan to exit early
  
  // launch goroutine to perform work with scope of a local results channel
  // select receive from cancel or local results
  
  res := make(chan error)
  go func(){
    res <-f(/*param*/)
  }()
  
  select {
    case <-tg.cancel:
      return
    case r:= <-res:
      tg.results <- r
      return
  }
}

// Wait blocks until all function calls from Run have returned, then returns the
// first non-nil error (if any) from them.
func (tg *TaskGroup) Wait() error {
  /* looping over receive from results chan
      if nil error continue
      else cancel running tasks via cancel chan
  */
  
  var err error
  for i := 0; i < tg.numTasks; i++ {
    err <-tg.results
    if err != nil {
      close(tg.cancel)
      return err
    }
  }
  
  return nil
}

func main() {
  tg := NewTaskGroup()
  endpoints := []string{...}
  for _, e := range endpoints {
    tg.Run(func(/*params*/) error {
      doThing(e)
    })
  }
  
  fmt.Println(tg.Wait())
}
