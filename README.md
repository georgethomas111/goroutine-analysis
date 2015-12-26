
README
======

GoRoutine Scaling
-----------------

This program gives an idea of the goroutines and how hungry they are to the cores. This gives an idea of the number of cpu usage statistics for the process that is running. 

DESIGN
======

Use "runtime" package to get the stats. 

Have ways to display it in a timely manner, i.e either by listening on a channel or by a heart beat function.

DONE
====

Basic Producer consumer model.
with channel usage statistic.

TODO
====

Add the run time support to do the analysis of the threads.

Use the run time information to schedule in the pool of servers available.
