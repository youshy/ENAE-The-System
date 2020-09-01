# ENAE - The System

AKA - Nerdy Music Video.

Watch our version on [youtube](https://www.youtube.com/watch?v=-_-2EpUqb9g). Read more on [my blog](https://akondas.com/blog/How-To-Write-A-Music-Video-In-Go)

## Why?

I always liked music videos (or movies) that were different to anything else. Remember video to [I Can Talk by Two Door Cinema Club](https://www.youtube.com/watch?v=bJDCMth8poM)? I'm not a fan of the band, but this music video has influenced me so much, that till this very day I'm wondering how it's made.

I'm not a Guy Richie, I can't shoot super-cool slow-mo shoots (even if I'd like to!), but I can code. Thus, the idea for this music video!

I like to think about it being an experience. Best viewed on big screens and LOUD!

## How can I run it?

Whole thing is written in Go, so crucial - make sure you have [Go](https://golang.org/doc/install) installed.

Then it's easy peasy - clone this repo and `go run .`!

## Can I break it?

Please do! As it's an experience, the crucial bit is that you can make it your own! You can read more about it on my [blog](https://akondas.com) (post is still yet to come, but it'll be there!)

## How can I break it?

So, this is where the fun starts. I'll keep it brief and let you do the tour!

Main things are the two function, namely `play` and `printer`. One deals with playing the song, the other one with printing on the screen. We'll focus on the printer.

The song is in 143 BPM and for the ease of use, I've calculated milliseconds values of each notes. You can find them in `printer.go`, lines `13` to `22`. Each song part has it's own function, which means that if you want to change a part of the song, then you go to the function that deals with that part!

Happy breaking The System!
