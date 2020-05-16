# GOAL 1: Printing the Digits

The goal here is to create led digits and display them to the terminal. As per the course outline, we have the below pointers on generating the digits and the separator which will be used later.

## Creating the digits and separators.

Here the digits from `zero` through `nine` will be created using the below format specifications.

```
      Digit character       : █
      Separator character   : ░
```

All the digits may be put into an array

## As per the course outline, the below pointers can be used to print the digits side-by-side.

+ Loop for all the lines in a digit

+ Print each digit line by line

+ Don't forget printing a space after each digit

+ Don't forget printing a newline after each line

     EXAMPLE: Let's say you want to print 10.

```
      ██   ███ <--- Print a new line after printing a single line from
       █   █ █     all the digits.
       █   █ █
       █   █ █
      ███  ███
         ^^
         ||
         ++----> Add space between the digits
```

## Test run.

Here is the output obtained after the digit generation from the code

```bash
⇒  go run main.go
GOAL 1: Printing the Digits...
Place holder type for digits: [10][5]string

Priting only first digit...
███
█ █
█ █
█ █
███

Printing digits vertically...
███
█ █
█ █
█ █
███

██
 █
 █
 █
███

███
  █
███
█
███

███
  █
███
  █
███

█ █
█ █
███
  █
  █

███
█
███
  █
███

███
█
███
█ █
███

███
  █
  █
  █
  █

███
█ █
███
█ █
███

███
█ █
███
  █
███


Using field length...
   ███   ██    ███   ███   █ █   ███   ███   ███   ███   ███
   █ █    █      █     █   █ █   █     █       █   █ █   █ █
   █ █    █    ███   ███   ███   ███   ███     █   ███   ███
   █ █    █    █       █     █     █   █ █     █   █ █     █
   ███   ███   ███   ███     █   ███   ███     █   ███   ███

Using a print format...
███  ██   ███  ███  █ █  ███  ███  ███  ███  ███
█ █   █     █    █  █ █  █    █      █  █ █  █ █
█ █   █   ███  ███  ███  ███  ███    █  ███  ███
█ █   █   █      █    █    █  █ █    █  █ █    █
███  ███  ███  ███    █  ███  ███    █  ███  ███

```
