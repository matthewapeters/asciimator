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

## Silly Ideas ##

* `asciimator` works with `lolcat`, but results vary.  May cause seisures.

  ```bash
   asciimator ./unicorn.txt | lolcat
   ```

* For better color control, add ANSI escaped color codes.  I find it easier to
  animate my frames first, and then add color instructions, as it will throw off
  you spacing,  You can find tables of color codes [on Wikipedia.](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors) To escape ANSI color codes, use the following pattern:  z

   `\x1b[`\<ANSI-CODE> `m`

   or

   `\033[`\<ANSI-CODE>`m`

   Remember to use the ANSI code `0` to turn off color.

