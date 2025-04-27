# Golang

## Intro

Golang is a programming language that was designed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, and publicly announced in November of 2009.

If there’s one thing most people know about Go, is that it is designed for concurrency. No introduction to Go is complete without a demonstration of its goroutines and channels.

But when people hear the word concurrency they often think of parallelism, a related but quite distinct concept. In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

## Parallelism vs. concurrency

To put a simple example of concurrency: imagine you have the task to eat a whole cake and sing a whole song.
The rule is that you sing and eat simultaneously. You can eat the whole cake, then sing the whole song, or you can eat half a cake, then sing half a song, then do that again, etc.

Following the example, parallelism would be that a friend that knows how to sing joins you in this task, each of you gets a process (e.g: your friend is singing, and you are eating a cake) and you do that at the same time, in parallel.

Parallelism requires hardware with multiple processing units, essentially. In single-core CPU, you may get concurrency but NOT parallelism. Parallelism is a specific kind of concurrency where tasks are really executed simultaneously.

## What will I find here?

This book is not a dedicated golang learning book but rather an introduction into programming. I personally love golang and that's why I choose this language, as it also has a very big standard library meaning no need to download dependencies, garbage collection, widely adopted in the open source community, compiled language (generally more efficient than interpreted ones) and I could bore you by growing this list.

Take your time to understand things, feel free to stop and return (or desist) whenever you feel like.

## How do I install golang?

You can install golang [in it's official website](https://go.dev/doc/install)

Hope that you enjoy it!
