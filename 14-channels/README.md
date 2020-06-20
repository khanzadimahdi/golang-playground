1. create a channel with **make** command
    - make(chan int)
2. Send message into channel
    - ch <- val
3. Receive message from channel
    - val := <- ch
4. Can have multiple senders and recievers