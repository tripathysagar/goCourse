- go routine is the independent threads, which can be initiated by any `go` key infront of any stmt. The code will create a thread and run the stmt there
- go channels are virtual tunnels. In which we can put data in one end and get the ooutput in another end. It is used for saheing objects between threads. Two types of the channel
   - buffered: buffer size of given during declaration
   - unbuffered: buffer size 1

- channel in go, will always going block if it is full
- for buffered channel we have to close the channel at the end of the computation otherwise we will get the "deadlock error"
- `sync.Waitgroup{}` is used as a signal to check if execution is done.
`sync.Waitgroup{}.Add(1)` used for increasing the signal. `sync.Waitgroup{}.Done()` is used for releasing the signal. And `sync.Waitgroup{}.Wait()` will check the complete execution if all the signal is released.
- mutex for large data 
- atomic for small data