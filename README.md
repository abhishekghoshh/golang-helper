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
- **go project file structure**
  - **repos**
    - [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
    - [evrone/go-clean-template](https://github.com/evrone/go-clean-template)
    - [go-kit/kit](https://github.com/go-kit/kit)
    - [facebookarchive/grace](https://github.com/facebookarchive/grace)
    - [astaxie/beego](https://github.com/astaxie/beego)
  - **blogs**
    - [Organizing a Go module](https://go.dev/doc/modules/layout)
    - [Golang project directory structure](https://stackoverflow.com/questions/46646559/golang-project-directory-structure)
    - [Go - Project Structure and Guidelines](https://dev.to/jinxankit/go-project-structure-and-guidelines-4ccm)
    - [Tips to Create a Proper Go Project Layout](https://www.developer.com/languages/go-project-layout/)
    - **medium**
      - [My Favourite Go Project Structure](https://martengartner.medium.com/my-favourite-go-project-setup-479563662834)
  - **youtube**
    - [This Is The BEST Way To Structure Your GO Projects](https://www.youtube.com/watch?v=dxPakeBsgl4)
    - [How I Structure New Projects In Golang](https://www.youtube.com/watch?v=dJIUxvfSg6A)
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
  - **GopherConAU**
    - [GopherConAU 2023](https://www.youtube.com/playlist?list=PLN_36A3Rw5hFsJqqs7olOAxxU-WJGlXS0)
    - [GopherConAU 2019](https://www.youtube.com/playlist?list=PLN_36A3Rw5hFJVoIf31_MeN67Pqj2NGrB)
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
  - [domain driven design](https://www.youtube.com/playlist?list=PLeoD63TPS-_ZofX56-vg8gJVXpDEpsIDW)
  - [Anthony GG](https://www.youtube.com/@anthonygg_/playlists)
  - [High Available Microservices With Apache Kafka In Golang](https://www.youtube.com/watch?v=-yVxChp7HoQ)
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


Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.</br>
Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.</br>
The language is often referred to as Golang because of its domain name, golang.org, but its proper name is Go.</br>
How do we run GO projects in CMD :</br>
- **go build** : Compiles a bunch of go source code files
- **go run** : Compiles and executes one or two files
- **go fmt** : Formats all the code in each file in the current directory
- **go install** : Compiles and "installs" a package/current porject in bin directory
- **go get** : Downloads the raw source code of someone else's package
- **go test** : Runs any tests associated with the current project
