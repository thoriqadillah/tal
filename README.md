# Terminal Artificial Life
Terminal Artificial Life is a reimagining of artificial life simulation in terminal inspired by this video https://youtu.be/0Kx4Y9TVMGg

This project aims to teach me Go programming language in general, but it mostly just for fun and to see how it looked like if the simulation rendered in ASCII

It is to be expected not looked as good as if it was rendered in actual canvas, but who cares? As long as it looks cool, it doesn't matter, and I think it looks pretty cool

## Example
![example](/asset/video-example.gif)

## The Motivation
I was watching youtube, and suddenly youtube recommended me the video and the video is really cool. I said to myself that I want to try the code, and I did. I did try the code in html and javascript environment. But it led my think that if I was going to try it, then maybe code it in different way. At the time I want to learn Golang because Golang is pretty popular in backend development. Well, technically, I did teach myself to learn Golang, but it was so boring because it was "yet another REST API project". I was having too many REST API project for learning purposse, and then I decided to make a project that has something to do with terminal, and then you know the rest of the story

## How It Works
In general, rendering in terminal is just printing bunch of strings from top to bottom, but have you ever thought about loading animation in terminal? Not actual animation, but like rendering from 1 to 100% but the cursor stays at the same place of the terminal. The simplest method we could do that is with just regular print and clearing the terminal and update the state of the string we want to animate. That would work, but that is not the optimal solution because that would hit the performance pretty bad

Another method we can use is to utilize ANSII escape characters. In general, ANSII escape character is used to control the cursor of the terminal. In this particular case, ANSII is used to place the cursor of the terminal at the top left corner everytime the canvas update the particle position. This way, the terminal will refresh the canvas with every iteration. The same method is used to colors the particle