# Rate limitting

## Blogs 
- [Rate limiting your Go application](https://blog.logrocket.com/rate-limiting-go-application/#what-is-rate-limiting)
- [tollbooth](https://github.com/didip/tollbooth)

## Documentation

### What is rate limiting?
Essentially, rate limiting is the process of controlling how many requests your application users can make within a specified time frame. There are a few reasons you might want to do this, for example:
- Ensuring that your application will function regardless of how much incoming traffic it receives
- Placing usage limits on some users, like implementing a capped free tier or a Fair Use policy
- Protecting your application from DoS (Denial of Service) and other cyber attacks

#### Token bucket algorithm :
In the token bucket algorithm technique, tokens are added to a bucket, which is some sort of storage, at a fixed rate per time. Each request the application processes consumes a token from the bucket. The bucket has a fixed size, so tokens can't pile up infinitely. When the bucket runs out of tokens, new requests are rejected.

#### Leaky bucket algorithm :
In the leaky bucket algorithm technique, requests are added to a bucket as they arrive and removed from the bucket for processing at a fixed rate. If the bucket fills up, additional requests are either rejected or delayed.

#### Fixed window algorithm :
The fixed window algorithm technique tracks the number of requests made in a fixed time window, for example, every five minutes. If the number of requests in a window exceeds a predetermined limit, other requests are either rejected or delayed until the beginning of the next window.

#### Sliding window algorithm :
Like the fixed window algorithm, the sliding window algorithm technique tracks the number of requests over a sliding window of time. Although the size of the window is fixed, the start of the window is determined by the time of the user's first request instead of arbitrary time intervals. If the number of requests in a window exceeds a preset limit, subsequent requests are either dropped or delayed.

Each algorithm has its unique advantages and disadvantages, so choosing the appropriate technique depends on the needs of the application. With that said, the token and leaky bucket variants are the most popular forms of rate limiting.