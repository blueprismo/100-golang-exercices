# 100 exercices to improve with GOLANG

## Intro

I hadn't find any good repository with 100 exercices for the GO Programming language, and as It's a beautiful language, I will create it myself :smile: feel free to contribute and use it!

The following exercises assume a fairly basic level of object oriented programming, but it is also OK if you don't have it.

## Overview

My main idea is that the exercices are structured like this:

- From 1 to 15   : Basic exercices
- From 16 to 19  : File management
- From 20 to 31  : Intermediate exercices
- From 32 to 44  : Concurrency
- From 45 to 48  : Time
- From 49 to 53  : JSON
- From 54 to 61  : Context
- From 62 to 68  : Interfaces
- From 69 to 82  : HTTP & APIs
- From 83 to 89  : Testing & misc
- From 90 to 100 : Database (non-relational(mongo)) + CRUD API with Fiber

Each folder will contain one exercise and one solution. There are multiple ways to acheive the same reault, so everything counts.

```bash
├── n-th_exercise       # <- exercise number
│    ├── README.md      # <- exercise description and info + tips (for the book!)
│    ├── main.go
│    ├── [.env]
│    └── solution       # <- solution folder
│       └── solution.go
├── go.mod
└── go.sum
```

## Prerequisites

### Git

Git is a distributed version control system that tracks versions of files.
It is often used to control source code by programmers who are developing software collaboratively, as an example more than 1000 persons could work in a single project simultaneously without interrupting themselves.

If you have never heard about git and you are interested in coding, it is a tool of vital importance for you to know.

This book text is stored in a git repository, and ideally you should be able to fork the repository so you can execute your code locally and play around.

[Download git](https://git-scm.com/downloads)

### Github / VCS

Github is an online platform to store your code. You can also store your code locally (without need of git) or in any other platform (Gitlab, bitbucket, azure, AWS codecommit, etc.)

While writing code, a version control system is highly advised, so you can have states stored where your code is working as expected, like checkpoints.

I personally love github, for it's wide community, the easy user experience and some of the most famous open source repositories are stored in there, like [linux](https://github.com/torvalds/linux), [bitcoin](https://github.com/bitcoin/bitcoin), [kubernetes](https://github.com/kubernetes/kubernetes) and even the [go](https://github.com/golang/go) programming language itself as a mirror.

### An IDE of your choice

An [IDE](https://en.wikipedia.org/wiki/Integrated_development_environment) is where you are going to write your code.
IDEs provide a great variety of features and plugins to help you code smoother, faster, with less errors and more efficiently.

Visual Studio Code is a great free IDE if you are new to programming. But feel free to use whatever you are comfortable with. (You can also write code in the simplest notepad or text editors)

[Download Visual Studio Code](https://code.visualstudio.com/)

### Golang

Golang is a programming language that was designed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, and publicly announced in November of 2009.

If there’s one thing most people know about Go, is that it is designed for concurrency. No introduction to Go is complete without a demonstration of its goroutines and channels.

But when people hear the word concurrency they often think of parallelism, a related but quite distinct concept. In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

#### Parallelism vs. concurrency

To put a simple example of concurrency: imagine you have the task to eat a whole cake and sing a whole song.
The rule is that you sing and eat simultaneously. You can eat the whole cake, then sing the whole song, or you can eat half a cake, then sing half a song, then do that again, etc.

Following the example, parallelism would be that a friend that knows how to sing joins you in this task, each of you gets a process (e.g: your friend is singing, and you are eating a cake) and you do that at the same time, in parallel.

Parallelism requires hardware with multiple processing units, essentially. In single-core CPU, you may get concurrency but NOT parallelism. Parallelism is a specific kind of concurrency where tasks are really executed simultaneously.

### How do I install golang?

You can install golang [in it's official website](https://go.dev/doc/install)

## How do I start if I want to do the execises?

After all the prerequisites above, fork the [100-golang-exercises](https://github.com/blueprismo/100-golang-exercices/fork) git repository or clone it if you don't have a github account and you'll be ready to go.

## Note

Please notice there are some exercises that are intended to have an execution failure, the important thing here is learning.
Feel free to contribute, comment, share your solutions and of course have lot of fun!

Thanks a lot
