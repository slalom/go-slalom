# Why Go

> A little copying is better than a little dependency.

> Clear is better than clever.

> Reflection is never clear.

Based on my experience using Go for the last 2 years I highly recommend you learn and use Go for new services that you
create. 

**Why?** I feel the simplicity of Go allows me to more easily solve problems and not let the language get in my way.
And it's fun!


## What is great about Go
- Simple but opinionated
- Statically typed (strong), modern and **fast** compiler
- Produces **fast**, runnable **binaries**. Easily generate binary for any platform. Run anywhere with Docker
- Quickly and easily build CLI/APIs. You don’t need a web development framework! The standard library provides everything 
you need to build a web server, do routing, html templates, file serving, res/req handling, databases, etc..
- Strong standard library
- Strong standard tooling
- Strong open source community with abundance of packages
- Tests first class citizen
- Designed for **scalability** and **concurrency**

## What Go doesn't have (intentionally)
- Inheritance and Classes. It uses structs, interfaces, and struct embedding to achieve powerful but simple OOP.
- Generics. Generics implementations add complexity and slow compilation or execution. Using interfaces and type aliases 
can suffice for generics. And after a while you won't miss `List<Dictionary<string<IEnumerable<HttpRequest>>>>`
- Annotations
- Exceptions

Some of the points above you may find shocking and think that go lacks "modern" language features... but trust me, the
small amount of extra code you may write is worth the time you will save.

## 1-5 minute go 101

A quick intro to programming in go 

[go 101](go-101.md)
  
  
  
  

![oo-dev](images/go-dev.jpeg)