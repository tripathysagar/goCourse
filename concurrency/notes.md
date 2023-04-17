- go routine is the independent threads, which can be initiated by any `go` key infront of any stmt. The code will create a thread and run the stmt there
- go channels are virtual tunnels. In which we can put data in one end and get the ooutput in another end. It is used for saheing objects between threads. Two types of the channel
   - buffered: buffer size of given during declaration
   - unbuffered: buffer size 1

- channel in go, will always going block if it is full