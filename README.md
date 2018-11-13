# Genetic DJ

[Currently running here.](http://54.215.128.40:8080/)

## About

This is a project to generate music inspired by the short C programs that
[countercomplex originally demoed][countercomplex]. This takes it a step
further than handcrafted C functions being piped to an audio device.

1. Functions are randomly generated. 
1. Functions are represented with S-Expressions for easy parsing.
   
   If an expression is `"(* t 7)"` for example, you can imagine the LISP definition:
   ```lisp
   (defun f (t)
     (* t 7))
   ```
   And then we evaluate this function for `t` = 0, 1, 2... and treat the
   resulting sequence as 8 kHz u8 audio.

1. Easy parsing allows "breeding" parse trees, enabling genetic algorithms to
   create "better" functions.
1. I'm writing the audio and serving it to the web. It's hard to judge fitness
   of music without people's opinions!

## Links

- [The original countercomplex blog post][countercomplex]
- Countercomplex's [first][video1], [second][video2], and [third][video3] videos

[countercomplex]: http://countercomplex.blogspot.com/2011/10/algorithmic-symphonies-from-one-line-of.html
[video1]: https://www.youtube.com/watch?v=GtQdIYUtAHg
[video2]: https://www.youtube.com/watch?v=qlrs2Vorw2Y
[video3]: https://www.youtube.com/watch?v=tCRPUv8V22o
