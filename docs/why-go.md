# Why Go

So why use Go?

Based on my experience using Go for the last 2 years I highly recommend you learn and use Go for new services that you
create. Why? I feel the simplicity of Go allows me to more easily solve problems and not waste time worrying 
about the intricacies of a language or it's run time. And it's fun!

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

![oo-dev](images/go-dev.jpeg)