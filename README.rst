Question
========

Implement a Task queue. Have a Task struct and create Task objects and push them to a queue. Have a go-routine
which periodically checks the tasks queue and inspect if the task is completed or not. If the task is
completed then remove it from the queue, if not completed push back at the end of the queue. If the task is
not completed after a certain amount of time then it should be removed from the queue and marked as a timeout.
Example of Task Struct:

``` code:go
type Task struct {
    Id string
    IsCompleted boolean // have a random function to mark the IsCompleted after a random period
    Status string //completed, failed, timeout
    Time string // when was the task created
    TaskData string // random string containing data about the task (can be a struct containing more information)
}
```

Implement the above mentioned logic with proper error handling.

Example

=======
Real-world scenario - 3 separate workers on the queue
1. Adding the task (add 9 emails in the queue (task))
2. Working on the task - picks up the task and processes them (the processing can fail or succeed) - read the
   email to send one by one and attempt to send it (sending can succeed or fail)
3. Cleans up the queue depending on fail or succeed or timeout status - cleans up the queue of emails (if
   email keeps failing then remove it from the queue and log the problem) !!!!

Upload to GitHub and send back a link to email address. Don’t send zip files in email, just the link.
