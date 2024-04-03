# Golang Basics

## Table of contents

### My resources
- [my cheetsheet](/cheetsheet.md)
- [my code samples](/codes_samples.md)
- [basic](./basics/)
- [design pattern](./design-pattern/)
- [protocol buffers](./grpc/potocol-buffers/)
- [grpc bacis](./grpc/grpc-basics/)
- [projects](./projects/)
  - [mysql-crud-with-jwt](./projects/crud-jwt/)
  - [blogpost-with-grpc-mongodb](./projects/blog-grpc/)
  - [go-cli](./projects/go-cli/)


### online resources and tools
- [gobyexample](https://gobyexample.com/)
- [go dev tutorial](https://go.dev/doc/tutorial/)
- [Effective Go](https://go.dev/doc/effective_go)
- [100go.co](https://100go.co/)
  - [100-go-mistakes](https://github.com/teivah/100-go-mistakes)
- [go101](https://go101.org/)
- [live reloading](https://github.com/cosmtrek/air)
- [mock generator](https://github.com/uber-go/mock)
- **database-helper**
  - [sqlc](https://docs.sqlc.dev/en/stable/index.html)
  - [ORM library](https://gorm.io/)
- [Go linter](https://github.com/golangci/golangci-lint)
- [Dependency Injection](https://blog.matthiasbruns.com/golang-the-ultimate-guide-to-dependency-injection)
  - [wire](https://github.com/google/wire/)
  - [dig](https://github.com/uber-go/dig)
  - [inject](https://github.com/facebookarchive/inject)

### Youtube videos
- [Google I/O 2012 - Meet the Go Team](https://www.youtube.com/watch?v=sln-gJaURzk)
- [How GO Was Created - Less Is More Prime Reacts](https://www.youtube.com/watch?v=4EMcm9vzlnI)

- **Data Structures**
  - [Data Structures in Golang Series](https://www.youtube.com/playlist?list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6)

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
  - [New Go Billion Row Challenge w/ Great Optimizations Prime Reacts](https://www.youtube.com/watch?v=SZ1PDS7iRU8)
    - [One Billion Row Challenge in Golang - From 95s to 1.96s](https://r2p.dev/b/2024-03-18-1brc-go/)
  
- **context**
  - [Golang Context Explained - How To Use With Timeout](https://www.youtube.com/watch?v=fXzzF5y6UEU)
  - [Learn Go context from code and its original blog post](https://www.youtube.com/watch?v=cXVI0vZFPkA)
  - [Understanding Contexts in Go in 5(-ish?) Minutes](https://www.youtube.com/watch?v=h2RdcrMLQAo)
  - [How To Use The Context Package In Golang?](https://www.youtube.com/watch?v=kaZOXRqFPCw)
  - [golang context package explained: the package that changed concurrency forever](https://www.youtube.com/watch?v=8omcakb31xQ)

- **live-reloading**
  - [cosmtrek/air](https://github.com/cosmtrek/air)
    - [This Package Makes It SO Much Easier To Write Golang Code](https://www.youtube.com/watch?v=erdDM_LmChs)
    - [16 Golang - Live Reloading Go app](https://www.youtube.com/watch?v=txVO9DaETZA)
    - [Live Reloading in Go / Golang - with Air](https://www.youtube.com/watch?v=ErVx9QyoBQI)
    - [HOT Reloading The Browser With Templ, Tailwind, And Golang](https://www.youtube.com/watch?v=6Pj-Vlhp31Y)

- **testing**
  - [Golang Testing (full tutorial)](https://www.youtube.com/watch?v=FjkSJ1iXKpg)

- **others**
  - [The standard library now has all you need for advanced routing in Go.](https://www.youtube.com/watch?v=H7tbjKFSg58)
  - [3 Golang Tips For Beginners I Wish I Knew Sooner](https://www.youtube.com/watch?v=PUPqnDYoMgU)
  - [7 Deadly Mistakes Beginner Go Developers Make (and how to fix them)](https://www.youtube.com/watch?v=biGr232TBwc)
  - [GoLang: 10+ UNIQUE Concepts/Conventions that Beginners Should Know About!](https://www.youtube.com/watch?v=CK5rLpZk5A8)
  - [domain driven design](https://www.youtube.com/playlist?list=PLeoD63TPS-_ZofX56-vg8gJVXpDEpsIDW)
  - [High Available Microservices With Apache Kafka In Golang](https://www.youtube.com/watch?v=-yVxChp7HoQ)
  

- **Some youtube channels**
  - [Melkey](https://www.youtube.com/@MelkeyDev/playlists)
  - [Mario Carrion](https://www.youtube.com/@MarioCarrion/playlists)
  - [Anthony GG](https://www.youtube.com/@anthonygg_/playlists)
  - [TheVimeagen](https://www.youtube.com/@TheVimeagen/videos)



### go project file structure
- **repos**
  - [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
  - [evrone/go-clean-template](https://github.com/evrone/go-clean-template)
  - [go-kit/kit](https://github.com/go-kit/kit)
  - [go-kit/examples](https://github.com/go-kit/examples)
  - [perkeep/perkeep](https://github.com/perkeep/perkeep)
  - [facebookarchive/grace](https://github.com/facebookarchive/grace)
  - [astaxie/beego](https://github.com/astaxie/beego)
  - [gravityblast/fresh](https://github.com/gravityblast/fresh)
- **blogs**
  - [Organizing a Go module](https://go.dev/doc/modules/layout)
  - [Golang project directory structure](https://stackoverflow.com/questions/46646559/golang-project-directory-structure)
  - [Go - Project Structure and Guidelines](https://dev.to/jinxankit/go-project-structure-and-guidelines-4ccm)
  - [Go Project Structure Best Practices](https://tutorialedge.net/golang/go-project-structure-best-practices/)
  - [Tips to Create a Proper Go Project Layout](https://www.developer.com/languages/go-project-layout/)
  - **medium**
    - [Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)
      - [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
      - [Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c)
    - [My Favourite Go Project Structure](https://martengartner.medium.com/my-favourite-go-project-setup-479563662834)
      - [I'll take pkg over internal](https://travisjeffery.com/b/2019/11/i-ll-take-pkg-over-internal/)
    - [Getting Started with Go: Project Structure](https://medium.com/evendyne/getting-started-with-go-project-structure-ab8814ded9c3)
    - [Package Organization in Go](https://medium.com/@leodahal4/package-organization-in-go-34efb1cd99a6)
    - [Go REST API Boilerplate](https://medium.com/@bhadange.atharv/go-rest-api-boilerplate-c9d25f99acbe)
- **youtube**
  - [This Is The BEST Way To Structure Your GO Projects](https://www.youtube.com/watch?v=dxPakeBsgl4)
  - [How I Structure New Projects In Golang](https://www.youtube.com/watch?v=dJIUxvfSg6A)


### Style guide
- [Google style guide for GO](https://google.github.io/styleguide/go/)
  - [Why Google stores billions of lines of code in a single repository](https://dl.acm.org/doi/pdf/10.1145/2854146?trk=public_post_comment-text)
  - [Style Guides and Rules](https://abseil.io/resources/swe-book/html/ch08.html#style_guides_and_rules)
  - [Knowledge Sharing](https://abseil.io/resources/swe-book/html/ch03.html#readability_standardized_mentorship_thr)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Golang style guide](https://developers.mattermost.com/contribute/more-info/server/style-guide/)


## Udemy courses

- [REST based microservices API development in Golang](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang)
- [Learn the Why's and How's of concurrency in Go.](https://udemy.com/course/concurrency-in-go-golang)
- [Working with Concurrency in Go (Golang)](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)
- [Complete Guide to Protocol Buffers 3](https://www.udemy.com/course/protocol-buffers/)
- [gRPC Golang Master Class: Build Modern API & Microservices](https://www.udemy.com/course/grpc-golang/)
- [Backend Master Class Golang + Postgres + Kubernetes + gRPC](https://udemy.com/course/backend-master-class-golang-postgresql-kubernetes/)
- [Golang For DevOps And Cloud Engineers](https://udemy.com/course/golang-for-devops-and-cloud-engineers/)


## Common go terminologies

Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.

Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases.

The language is often referred to as Golang because of its domain name, golang.org, but its proper name is Go.

How do we run GO projects in CMD :
- **go build** : Compiles a bunch of go source code files
- **go run** : Compiles and executes one or two files
- **go fmt** : Formats all the code in each file in the current directory
- **go install** : Compiles and "installs" a package/current porject in bin directory
- **go get** : Downloads the raw source code of someone else's package
- **go test** : Runs any tests associated with the current project
