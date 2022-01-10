Objectives:\
You must follow the same instructions as in "ascii-art" but this time with colors.

The output should manipulate colors using the flag --color=<color>, in which --color is the flag and <color> is the color desired by the user. These colors can be achieved using different notations (color code systems, like RGB, hsl, ANSI...), it is up to you to choose which one you want to use.

You should be able to choose between coloring a single letter or a set of letters (use your imagination for this one).
If the letter is not specified, the whole string should be colored.
The flag must have exactly the same format as above, any other formats must return the following usage message:
Usage: go run . [STRING] [OPTION]

EX: go run . something --color=<color>
  
Instructions:\
Your project must be written in Go.\
The code must respect the good practices.\

  
  
Allowed packages:\
Only the standard Go packages are allowed
  
This project will help you learn about:\
The Go file system(fs) API\
Color converters\
Data manipulation\
Terminal display\





How to run the code:\
In order to color a specific letter please use third argument as follow: [string] [option] [firstelement] = "hello" --color=blue 3<br>
In order to color more than one letter (consecutives) please use third and fourth argument as follow:\
[string] [option] [firstelement] [secondelement] = "hello" --color=blue 0 2
