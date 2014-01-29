terimg
======

Display pictures in your Linux terminal with a retro 256 color design.


* Licence: MIT
* Author:  Julien CHAUMONT
* Contact: julienc91 [at] outlook.fr
* Version: 1.0.1 - 29/01/2014

The pictures which are used for testing purpose are under a Creative
Commons Licence:

- "Shipwreck" by [Abi Danial](http://500px.com/photo/51139820)


Introduction
------------

This program was made for practising GoLang and does not claim any
skill in this language. Just enjoy the result!

![Shipwreck by Abi Danial](https://github.com/julienc91/terimg/raw/master/img/shipwreck.png
 "Shipwreck by Abi Danial")
 
![Screenshot 1](https://github.com/julienc91/terimg/raw/master/img/shipwreck_120.png "Screenshot (120px width)
![Screenshot 2](https://github.com/julienc91/terimg/raw/master/img/shipwreck_250.png "Screenshot (250px width)


Requirements
------------

In order to use `terimg`, you will need:

- A 256 color compatible shell (`gnome-terminal` does the trick)
- A `go` compiler such as `go`


Usage
-----

Just run the compiled program:
    
    terimg <png_file> [terminal_width]
    
Currently, `terimg` is only working with `png` files.


Configuration
-------------

`terimg` uses a default configuration which is set in `rgb.csv`. It
gives all the rgb values of the 256 colors the user's terminal can
display. You can use the Python script `display_colors.py` to print
all the available colors of your own terminal in the correct order.
You will then need to modify `rgb.csv` in order to have better
results.


    
