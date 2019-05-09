# Why Go

My favorite go proverbs from https://go-proverbs.github.io

> A little copying is better than a little dependency.

> Clear is better than clever.

> Reflection is never clear.

Based on my experience using Go for the last 2 years I highly recommend you learn and use Go for new services that you
create. 

**Why?** I feel the simplicity and speed of Go allows me to more easily solve problems and not let the language get in my way.
And everyone likes gophers!


## What is great about Go
- Simple but opinionated
- Statically typed (strong), modern and **fast** compiler
- Produces **fast**, runnable, static **binaries**. Easily generate binary for any platform without any dependencies
- Designed for **scalability** and **concurrency**
- Quickly and easily build CLI/APIs. You don’t need a web development framework! The standard library provides everything 
you need to build a web server, do routing, html templates, file serving, res/req handling, databases, etc..
- Strong standard library
- Strong standard tooling
- Strong open source community with abundance of packages
- Tests first class citizen
- While not a functional programming language, go has many functional programming features
    - functions as values (first class objects)
    - closures
    - function types
    - anonymous functions
    

## What Go doesn't have (intentionally)
- Inheritance and Classes. It uses structs, interfaces, and struct embedding to achieve powerful but simple OOP.
- Generics. Generics implementations add complexity and slow compilation or execution. Using interfaces and type aliases 
can suffice for generics. And after a while you won't miss `List<Dictionary<string<IEnumerable<HttpRequest>>>>`
- Annotations
- Exceptions

Some of the points above you may find shocking and think that go lacks "modern" language features... but trust me, the
small amount of extra code you may write is worth the time you will save.  

![oo-dev](images/go-dev.jpeg)