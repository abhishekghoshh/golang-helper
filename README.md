# Golang Basics

#### Table of contents
- **codes** 
  - [basic](./basics/)
  - [design pattern](./design-pattern/)
  - [protocol buffers](./grpc/potocol-buffers/)
  - [grpc bacis](./grpc/grpc-basics/)
  - [projects](./projects/)
    - [mysql-crud-with-jwt](./projects/crud-jwt/)
    - [blogpost-with-grpc-mongodb](./projects/blog-grpc/)
    - [go-cli](./projects/go-cli/)
- **Resources**
  - [my cheetsheet](/cheetsheet.md)
  - [my code samples](/codes_samples.md)
  - [gobyexample](https://gobyexample.com/)
  - [Effective Go](https://go.dev/doc/effective_go)
  - [100go.co](https://100go.co/)
    - [100-go-mistakes](https://github.com/teivah/100-go-mistakes)
  - [go101](https://go101.org/)
- **Some youtube videos**
  - [Google I/O 2012 - Meet the Go Team](https://www.youtube.com/watch?v=sln-gJaURzk)
  - [How GO Was Created - Less Is More | Prime Reacts](https://www.youtube.com/watch?v=4EMcm9vzlnI)
  - **concurrency**
    - [Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
      - [Concurrency is not parallelism](https://go.dev/blog/waza-talk)
      - [Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)
    - [Go Concurrency](https://www.youtube.com/playlist?list=PL7g1jYj15RUNqJStuwE9SCmeOKpgxC0HP)
    - [Concurrency in Go](https://www.youtube.com/watch?v=LvgVSSpwND8)
    - [Making It FAST - 1 Billion Row Challenge in Go](https://www.youtube.com/watch?v=cYng524S-MA)
      - [1brc](https://github.com/shraddhaag/1brc)
      - [One Billion Rows Challenge in Golang](https://www.bytesizego.com/blog/one-billion-row-challenge-go)
  - [GoLang: 10+ UNIQUE Concepts/Conventions that Beginners Should Know About!](https://www.youtube.com/watch?v=CK5rLpZk5A8)
  - [Data Structures in Golang Series](https://www.youtube.com/playlist?list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6)
- **Some youtube channels**
  - [Melkey](https://www.youtube.com/@MelkeyDev/playlists)
  - [Mario Carrion](https://www.youtube.com/@MarioCarrion/playlists)
  - [Anthony GG](https://www.youtube.com/@anthonygg_/playlists)
- **Style guide**
  - [Effective Go](https://go.dev/doc/effective_go)
  - [Google style guide for GO](https://google.github.io/styleguide/go/)
    - [Why Google stores billions of lines of code in a single repository](https://dl.acm.org/doi/pdf/10.1145/2854146?trk=public_post_comment-text)
    - [Style Guides and Rules](https://abseil.io/resources/swe-book/html/ch08.html#style_guides_and_rules)
    - [Knowledge Sharing](https://abseil.io/resources/swe-book/html/ch03.html#readability_standardized_mentorship_thr)
  - [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
  - [Golang style guide](https://developers.mattermost.com/contribute/more-info/server/style-guide/)
- **medium blogs**


Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.</br>
Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.</br>
The language is often referred to as Golang because of its domain name, golang.org, but its proper name is Go.</br>
How do we run GO projects in CMD :</br>
- **go build** : Compiles a bunch of go source code files
- **go run** : Compiles and executes one or two files
- **go fmt** : Formats all the code in each file in the current directory
- **go install** : Compiles and "installs" a package. 
- **go get** : Downloads the raw source code of someone else's package
- **go test** : Runs any tests associated with the current project
