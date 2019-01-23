## Docker stuff  
  
I am intending to put docker related projects here, but this is early on.

`dlaunch.go` - This is a project I am working on as a wrapper for Docker images running interpreters/compilers.  Ideally, I am hoping to have it detect file names and such in the command, copy that data into the container (or just mount as a volume, not sure yet).  I am also hoping to have it auto-detect images you have loaded, so you could do "dlaunch perl <file>" and have it check that you have a Perl image, send the data to that with -rm to clean up when done, and return the data.  We'll see how it goes.  Go is pretty new to me, so this is both a way to learn it a bit, and also I am hoping that it will be able to compile this cleanly on any platform.  

