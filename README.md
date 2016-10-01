terimg
======

Display pictures in your Linux terminal with a retro 256 color design.


* Licence: MIT
* Author:  Julien CHAUMONT
* Contact: github[at]julienc.io
* Version: 1.0.1 - 29/01/2014

Introduction
------------

This program was made for practising GoLang and does not claim any
skill in this language. Just enjoy the result!

Examples
--------

Video demonstration: https://www.youtube.com/watch?v=89tYEL3nyLI

![Original image](https://github.com/julienc91/terimg/raw/master/img/autumn.jpg)
 
![Screenshot 1](https://github.com/julienc91/terimg/raw/master/img/autumn_120.png "Screenshot (120px width")
![Screenshot 2](https://github.com/julienc91/terimg/raw/master/img/autumn_250.png "Screenshot (250px width")


Requirements
------------

In order to use `terimg`, you will need:

- A 256 color compatible shell (`gnome-terminal` does the trick)
- A `go` compiler such as `go`


Usage
-----

Just run the compiled program:

    terimg <png_file> [terminal_width] [preserve_ratio]

Just set `[preserve_ratio]` to any value in order to try to preserve
the picture size ratio.

Currently, `terimg` is working with `png` and `jpg` files.


Configuration
-------------

`terimg` uses a default configuration which is set in `rgb.csv`. It
gives all the rgb values of the 256 colors the user's terminal can
display. You can use the Python script `display_colors.py` to print
all the available colors of your own terminal in the correct order.
You will then need to modify `rgb.csv` in order to have better
results.

