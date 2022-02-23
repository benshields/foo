// You are given a simplistic Go API for retrieving user events from a remote service.

type EventID int64 // Valid EventIDs are positive integers.
type UserID int64 // Valid UserIDs are positive integers.
const InvalidEventID EventID = 0

type Stream struct {}
// NewStream allocates a new stream for user ID uid. Does no I/O.
func NewStream (uid UserID) *Stream
// Err returns an error if an operation resulted in an error on the stream.
func (s *Stream) Err () error
// Open begins listening on the stream for events.
func (s *Stream) Open ()
// Iterate blocks until an event or error occurs. In the error case, this will return InvalidEventID and the error can be accessed via the Err() method. Only events that occurred after Open() was called will be returned.
func (s *Stream) Iterate () EventID
// Close closes the stream. This can be called concurrently with Iterate().
func (s *Stream) Close ()

/*
Using the Stream API, implement a MultiStream. A MultiStream must:
Accept a list of user IDs and internally create a Stream for each.
Merge the results of the underlying Streams, and return events one at a time.
Support a timeout on iteration.
*/

type MultiStream struct {
  mergedEvents chan EventID
  streams []Stream
}

// NewStream allocates a new stream for user ID uid. Does no I/O.
func NewStream (uid []UserID) *MultiStream {
  ms := MultiStream{}
  
  ms.mergedEvents = make(chan EventID)
  ms.streams = make([]Stream, len(uid))
  for i, id := range uid {
    s := NewStream(id)
    ms.streams[i] = s
  }
  
  return &ms
}

// Err returns an error if an operation resulted in an error on the stream.
func (ms *MultiStream) Err() error

// Open begins listening on the stream for events.
func (ms *MultiStream) Open() {
  for _, s := range ms.streams {
    s.Open()
  }
  go func(){
    for _, s := range ms.streams {
      go func(){
        while {
        ms.mergedEvents <-s.Iterate()
        }
      }()
    }
  }()
}

// Iterate blocks until an event or error occurs. In the error case, this will return InvalidEventID and the error can be accessed via the Err() method. Only events that occurred after Open() was called will be returned.
func (ms *MultiStream) Iterate(ctx Context.context) EventID {
  // loop calling s.Iterate() in goroutine, sending return value to mergedEvents
  done := make(chan struct{}, len())
  for _, s := range ms.streams {
    go func(){
      ms.mergedEvents <-s.Iterate()
    }()
  }
  
  var result EventID
  select {
    case result <-mergedEvents:
      if result != InvalidEventID {
        return result
      } else {
        // add support to return that specific stream's Err() value
      }
    case <-ctx.Done():
      return InvalidEventID
    default: // timeout
      // ???
  }
}

// Close closes the stream. This can be called concurrently with Iterate().
func (ms *MultiStream) Close ()