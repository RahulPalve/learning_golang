# Concurrency and Parallism

## Parallelism
- Utilising multiples cores of CPU, and processes executing in parallel
```
  p1 p2    p1 p2     p1 p2   
  |        |          |  |
  |            |      |  |
     |     |         
     |         |

General  Concurrent Parallel
```
-  Async task runs on a single thread, nothing related to multithreading. async task sends a task to an external process outside your application... i.e database, file reader... these are IO processes then a callback is added on the process to perform an action when the IO process is finished. 
- multithreading is concurrent
- a core of CPU can execute a thread at a time.
- thread is like execution flow control, 
- Process
Each process provides the resources needed to execute a program. A process has a virtual address space, executable code, open handles to system objects, a security context, a unique process identifier, environment variables, a priority class, minimum and maximum working set sizes, and at least one thread of execution. Each process is started with a single thread, often called the primary thread, but can create additional threads from any of its threads.

- Thread A thread is a single sequential flow of control within a program.
A thread is an entity within a process that can be scheduled for execution. All threads of a process share its virtual address space and system resources. In addition, each thread maintains exception handlers, a scheduling priority, thread local storage, a unique thread identifier, and a set of structures the system will use to save the thread context until it is scheduled. The thread context includes the thread's set of machine registers, the kernel stack, a thread environment block, and a user stack in the address space of the thread's process. Threads can also have their own security context, which can be used for impersonating clients.
- First unix only had process no thread, but this made context swiching done by os to schedule core to process slow cause it required, so unix created thread called light processes initially they share memory context, that way switchiing from one thread to another is faster
- In go os actually just process one thread main thread but go runtime schedular, utilise this main thread to share multiple goroutines think of them as a thread, so basically the scheduling task of os is done by go runtime schedular to main thread.
- go can do things in parallel, if required go create another logical process which handle diff goroutines and that process and other one can run on diff cores.


## Refs:     
- https://stackoverflow.com/questions/4844637/what-is-the-difference-between-concurrency-parallelism-and-asynchronous-methods
- https://stackoverflow.com/questions/20708707/why-was-the-comment-that-said-dont-format-a-floppy-at-the-same-time-funny-whe
- Coursera concurrency in go week 1 and week 2 are amzaing for understanding.