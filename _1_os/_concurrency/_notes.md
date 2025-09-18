Concurrency in a single thread environment work like this.
Lets take an example of a person who is eating and singing at a same time, here his mouth is single core CPU. Which means either he can eat or sing. this is called concurrency via Context switching.  So to do both he eat and than sing and than eat again than sing. 
Just give an illusion parallelism but its not a true parallelism. Just Context switching until both done

Now lets understand Parallelism. Same person Cooking and Washing his clothes. which means the environment he is doing two of the tasks simuntaneasuly is multiple core CPU. So both task can perform without overwhelming each other regardless of which one started before.
Both can be done in parallel ways.

Now let understand ASYNCRONOUES EXECUTION OF TASKS ;
Same Guy cooking and washing clothes while he is talking on a cell using Bluetooth headphones. Talking is a main thread for this guy. 
Async make sure that none of the running tasks block the main thread while executing. So this guy doing both tasks in parallel while talking as well.
