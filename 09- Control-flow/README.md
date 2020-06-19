# Control flow

1. Defer
    - Used to delay execution of a statement unfil function exists
    - Userful to group "open" and "close" function together
        - Be careful in loops
    - Run in LIFO (last-in, first-out) order
    - Arguments evaluated at time defer is executed, not at time of called function execution
2. Panic
    - Occur when program connot continue at all
        - Don't use when file can't be opened, unless it is citical
        - User for unrecoverable events-cannot optain TCP port for web server
    - Function will stop executing
        - Deferred functions will still fire
    - If nothing handles panic, program will exit 
3. Recover
    - Used to recover from panics
    - Only useful in deferred functions
    - Current function will not attempt to continue, but higher functions in call stack will
    