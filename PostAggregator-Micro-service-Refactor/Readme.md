# Social Media Feed Aggregator Microservice 

A Golang microservice implementing gRPC and GraphQL to manage users, their posts, and social feeds.

## Features

### Core Functionality
- **User Management**: Store users with follower/following relationships
- **Post Management**: Users can create posts 
- **Feed Generation**:
  - `getTimeline`: Fetch a user's own latest 20 posts
  - `getFeed`: Fetch latest 20 posts from users they follow

## Technical Approach


### Key Components
- **Models**:
  - `User`: Contains userId, userName, Following[] (users they follow), Posts (heap of posts)
  - `Post`: Contains postId, userId, content, timestamp

### Optimizations
- Used **min-heap** to efficiently maintain only the 20 most recent posts
- Implemented **concurrent fetching** of followed users' posts using go routines for each follower 
- **Dockerized** for easy deployment

## Challenges Faced

1. **Feed Aggregation Logic**
   - Implementing efficient heap-based storage instead of simple arrays.

   - Instead of building a full array of all followers' posts and sorting them to get the latest 20,
    we now use a more memory-efficient approach with a min-heap. The old method stored all posts (which could be large), while the new method keeps only the top 20 most recent posts in memory.

    - We use a min-heap of size 20 and remove the post with the oldest timestamp whenever the heap exceeds 20 entries. This ensures we're always keeping only the most recent posts, saving memory.

    - To improve performance, we also run goroutines for each followed user to fetch posts in parallel, while maintaining the top 20 results concurrently.

2. **Dockerization**
   - Passing data file (which have dummy data of users and posts),  through environment variables
   - Container networking between gRPC and GraphQL services

3. Variable Naming :`( 

   

## Local Setup

### Prerequisites
- Docker


### Installation
```bash
git clone git@github.com:Khemendra-Bhardwaj/PostAggregator-Micro-service.git
cd post-aggregator
make up  # or docker-compose up --build
```

### Sample Query

#### cURL Command

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"query": "{ getFeed(userId: \"1\") { postId content } }"}' \
http://localhost:8080/graphql
```

### Get user's timeline (their own posts)
``` bash
query {
  getTimeline(userId: "1") {
    postId
    content
    timestamp
  }
}
```

### Get user's feed (posts from followed users)
``` bash 
query {
  getFeed(userId: "1") {
    postId
    content
    timestamp
    userId
  }
}
```

### Sample outputs 

``` bash

Command:
 curl -X POST -H "Content-Type: application/json" \
-d '{"query": "{ getTimeline(userId: \"1\") { postId content } }"}' \
http://localhost:8080/graphql

Expected Response: 
{
	"data": {
		"getTimeline": [
			{
				"content": "Building my first microservice.",
				"postId": "2"
			},
			{
				"content": "Exploring gRPC with Golang!",
				"postId": "1"
			}
		]
	}
}

```

``` bash 

Command : 

curl -X POST -H "Content-Type: application/json" -d '{"query": "{ getFeed(userId: \"1\") { postId content } }"}' http://localhost:8080/graphql

Expected Response: 
{
  "data": {
    "getFeed": [
      {
        "postId": "3",
        "content": "Charlie's post",
        "timestamp": "2025-04-09T13:02:00Z"
      },
      {
        "postId": "1",
        "content": "Bob's post",
        "timestamp": "2025-04-09T12:45:00Z"
      }
    ]
  }
}
```


### Code Structure Overview
``` bash 

├── cmd
│   ├── grpc-server/       # gRPC server implementation
│   └── graphql-server/    # GraphQL server implementation
├── internal
│   ├── models/            # User and Post models
│   ├── grpcclient/        # gRPC client wrapper
│   └── schema/            # GraphQL schema
├── proto/                 # Protocol Buffer definitions
├── data.json              # Sample users data 
├── Dockerfile
└── docker-compose.yml

```

@ayush2025