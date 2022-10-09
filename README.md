# ASCIIMATOR #

## A Silly Little ASCII Animator ##

## Building the Binary ##

```bash
go build ./...
```

## Running Asciimator ##

```bash
./asciimator <animation file>

```

## Animation File ##

   1. Animation files are just text files
   2. the first line contains two numbers separated by a space:
      1. number of lines per frame
      2. number of milliseconds to show each frame
   3. The rest of the file is ASCII art.
      1. Each frame is same number of lines
      2. Frames are played in order found in the file

```bash



                            HAVE FUN!




                                 /|
                                 ||
                                 ||
                              /====~/\
                             / 0 0  /=~~
                             /         =~~
                            |[. .]/      ~~~~====== ~
                                   \                ~~
                {0O0}             /   / /-----/    /~~~
                 (0)              ||  \ \     [   ]  ~
                \ | /             [=    =]     | /
---Vv---vv---VvvV\|/vvv---------------------------------------------------
```
