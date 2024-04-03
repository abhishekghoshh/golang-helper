# GRPC

#### Allways remember to add or export this use your local golang module inside project
```
export GO111MODULE=on
```

## Blogs

#### Introduction
- [Protocol Buffers Documentation](https://protobuf.dev/programming-guides/proto3/)
- [Introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)
- [Clément Jean personal blog](https://clement-jean.github.io/)
- some examples
  - [dreamsofcode-io/grpc](https://github.com/dreamsofcode-io/grpc)
  - [google spanner](https://github.com/googleapis/googleapis/blob/master/google/spanner/v1/spanner.proto)
  - [google pub/sub](https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto)


#### rest vs gRPC
- [REST vs gRPC](https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html)
- [gRPC vs REST: let the battle begin!](https://www.slideshare.net/borisovalex/grpc-vs-rest-let-the-battle-begin-81800634)
- [Compare gRPC services with HTTP APIs](https://learn.microsoft.com/en-us/aspnet/core/grpc/comparison?view=aspnetcore-3.0)
- [What's the Difference Between gRPC and REST?](https://aws.amazon.com/compare/the-difference-between-grpc-and-rest/)

#### tools for testing gRPC
- [grpcui](https://github.com/fullstorydev/grpcui)
- [grpcurl](https://github.com/fullstorydev/grpcurl)


##### reverse proxy for gRPC [envoyproxy](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/other_protocols/grpc)  </br>
##### A JavaScript implementation of gRPC for browser clients [grpc-web](https://github.com/grpc/grpc-web)


#### Medium blogs

- **introduction**
  - [Why Pick gRPC?](https://medium.com/@baaalaji.arumugam/why-choose-grpc-because-its-the-cool-kid-in-the-api-block-f13a5fd96198)
  - [gRPC demystified](https://medium.com/@positiveb16/grpc-demystified-a455b694af9c)
  - [Unlocking Efficiency: Why gRPC Might Be Your Next API Choice Over REST](https://medium.com/@coderviewer/unlocking-efficiency-why-grpc-might-be-your-next-api-choice-over-rest-4a3b2ad06390)
  - [What is gRPC?](https://blog.stackademic.com/wtf-is-grpc-part-1-authentication-and-authorization-in-flutter-and-golang-74edc9a62285#setup-basic-project-in-golang)
    - [Wtf-is-grpc](https://github.com/Djsmk123/wtf-is-grpc)

- **microservice**
  - [Step by step making a Golang gRPC Server in 5 Minutes](https://blog.stackademic.com/step-by-step-making-a-golang-grpc-server-in-5-minutes-c8a36287ca36)
  - [How To: Build a gRPC Server In Go](https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05)
  - [gRPC Explained: Part 1 Introduction](https://medium.com/@dwivedi.ankit21/grpc-explained-part-1-introduction-6582dc4c7977)
    - [gRPC Explained: Part 2 Protobuf](https://medium.com/@dwivedi.ankit21/grpc-explained-part-2-protobuf-19d3d7f26cfa)
  - [Why Use gRPC for Microservices Communication](https://medium.com/@ali.assar/why-use-grpc-for-microservices-communication-632c077e5ce9)
  - [gRPC Go Microservice](https://rancavil.medium.com/grpc-go-microservice-b10dba570ec2)
  - [Building Scalable Microservices: Creating a GRPC Service with Go and Consuming it in a React App via Envoy](https://medium.com/@digvijay17july/building-scalable-microservices-creating-a-grpc-service-with-go-and-consuming-it-in-a-react-app-1de3c4385c05)
  - [Building Microservices with Golang and ScyllaDB: A Scalable Approach](https://medium.com/@jitenderkmr/building-microservices-with-golang-and-scylladb-a-scalable-approach-0ccb74cb95b4)
    - [How to Build a Golang gRPC Microservice with ScyllaDB: A Step-by-Step Guide](https://medium.com/@jitenderkmr/how-to-build-a-golang-grpc-microservice-with-scylladb-a-step-by-step-guide-c146797e8b1c)

- **message streaming**
  - [gRPC and Protobuf: Socket-Based APIs](https://blog.stackademic.com/grpc-and-protobuf-socket-based-apis-45f8421ba467)
  - [gRPC vs. WebSockets — choosing the right bi-directional/server-push protocol](https://medium.com/@n.sobadjiev_2847/grpc-vs-websockets-choosing-the-right-bi-directional-server-push-protocol-ea8be3e733d5)
  - [Building a Real-Time Chat Application with gRPC and Go](https://medium.com/@bhadange.atharv/building-a-real-time-chat-application-with-grpc-and-go-aa226937ad3c)
    - [grpc-chat](https://github.com/atharv-bhadange/grpc-chat)
  - [gRPC Chat using client and bi-directional streaming](https://medium.com/@anton_tomchuk/grpc-chat-using-client-and-bi-directional-streaming-d3a7662482d4)
    - [Use gRPC server streaming within your chat](https://medium.com/@anton_tomchuk/use-grpc-server-streaming-within-your-chat-71b6cd53a6ad)
    - [GrpcFClient](https://github.com/5tommy5/GrpcFClient)

- **others**
  - [Using metadata in gRPC using Golang as an example: Understanding and transmitting contextual information](https://blog.stackademic.com/using-metadata-in-grpc-using-golang-as-an-example-understanding-and-transmitting-contextual-580ecdbc4b79)
  - [Adding interceptor into Golang gRPC server](https://medium.com/@mahes0/adding-interceptor-into-golang-grpc-server-0a5ea4d12f27)
    - [Golang gRPC Rate Limiter using Redis Counter](https://medium.com/@mahes0/golang-grpc-rate-limiter-using-redis-counter-9401da642a2e)
  - [Unlocking the Secrets of How gRPC Clients Work](https://medium.com/@satokenta940/unlocking-the-secrets-of-how-grpc-clients-work-436b93c14848)
  - [Integrating GRPC at Scale](https://medium.com/@pawansharma_28214/integrating-grpc-at-scale-70e32e05b112)
- **grpc web and reverse proxy**
  - [A story of gRPC on Web](https://blog.datamole.ai/a-story-of-grpc-on-web-f32d633b3aa3)
  - [What Are gRPC & gRPC Web](https://medium.com/@nxenon/what-are-grpc-grpc-web-ecc9c3094c82)
    - [Envoy and gRPC-Web: a fresh new alternative to REST](https://blog.envoyproxy.io/envoy-and-grpc-web-a-fresh-new-alternative-to-rest-6504ce7eb880)
    - [Hacking into gRPC-Web](https://infosecwriteups.com/hacking-into-grpc-web-a54053757a45)
  - [Scaling Microservices with gRPC and Envoy Proxy — Part I— with gRPC](https://levelup.gitconnected.com/scaling-microservices-with-grpc-and-envoy-72a64fc5bbb6)
    - [Scaling Microservices with gRPC and Envoy Proxy — Part II — with Envoy Proxy](https://medium.com/swlh/scaling-microservices-with-grpc-and-envoy-proxy-part-2-148f589b2a83)
  - [Unleashing gRPC Gateway in Golang: Crafting a Reverse Proxy for Effortless RESTful Integration](https://medium.com/@jitenderkmr/exploring-grpc-gateway-in-golang-building-a-reverse-proxy-for-seamless-restful-integration-d342fe5248c4)

- **on-cloud**
  - [Load Balancing in gRPC (K8s)](https://medium.com/@ujala2yz/load-balancing-in-grpc-k8s-1a1dcea726ae)
    - [gRPC_LoadBalancing_K8s](https://github.com/ujala-singh/gRPC_LoadBalancing_K8s#local-dev-setup)
  - [How three lines of configuration solved our gRPC scaling issues in Kubernetes](https://medium.com/jamf-engineering/how-three-lines-of-configuration-solved-our-grpc-scaling-issues-in-kubernetes-ca1ff13f7f06)
    - [Kubernetes network load balancing using OkHttp client](https://medium.com/jamf-engineering/kubernetes-network-load-balancing-using-okhttp-client-54a2db8fc668)
  - [GRPC Client side connection pooling](https://arpittech.medium.com/grpc-and-connection-pooling-49a4137095e7)
  - [Step by step deploying gRPC Server into GCP Cloud Run](https://medium.com/@mahes0/step-by-step-deploying-grpc-server-into-gcp-cloud-run-d80dd84e2a89)
  - [Exposing gRPC services through GKE Ingress: A step-by-step guide](https://chimbu.medium.com/exposing-grpc-services-through-gke-ingress-a-step-by-step-guide-2cdf09d74b2d)

## Youtube

#### gRPC introduction
- [When RESTful architecture isn't enough...](https://www.youtube.com/watch?v=_4TPM6clQjM)
- [Where should you use gRPC? And where NOT to use it!](https://www.youtube.com/watch?v=4SuFtQV8RCk)
- **comparisons**
	- [tRPC, gRPC, GraphQL or REST: when to use what?](https://www.youtube.com/watch?v=veAb1fSp1Lk)
	- [REST vs RPC vs GraphQL API - How do I pick the right API paradigm?](https://www.youtube.com/watch?v=hkXzsB8D_mo)
	- [gRPC vs REST - KEY differences and performance TEST](https://www.youtube.com/watch?v=JjsT-i-ZEBc)
