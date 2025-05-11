package models

import (
	"container/heap"
	"sync"
)

type User struct {
	UserId    int
	UserName  string
	Posts     map[int]*PostHeap
	Following []*User
	Followers []*User
}

func (u *User) AddPost(post *Post) {
	if u.Posts == nil {
		u.Posts = make(map[int]*PostHeap)
	}

	if _, ok := u.Posts[post.UserId]; !ok {
		h := &PostHeap{}
		heap.Init(h)
		u.Posts[post.UserId] = h
	}
	heap.Push(u.Posts[post.UserId], post)
}

func (u *User) GetLatestPostByUser(userId int) *Post {
	if h, ok := u.Posts[userId]; ok && h.Len() > 0 {
		return heap.Pop(h).(*Post)
	}
	return nil
}

// Returns 20 recents post by user
func (u *User) GetRecentPostsByUser(userId int) []*Post {
	count := 20
	if h, ok := u.Posts[userId]; ok && h.Len() > 0 {
		var recent []*Post
		var temp []*Post

		for i := 0; i < count && h.Len() > 0; i++ {
			post := heap.Pop(h).(*Post)
			recent = append(recent, post)
			temp = append(temp, post)
		}

		for _, post := range temp {
			heap.Push(h, post)
		}

		return recent
	}
	return nil
}

// Get User Feed From its Followings
func (u *User) GetUserFeed(userId int) []*Post {
	var wg sync.WaitGroup
	var mu sync.Mutex
	feedHeap := &PostHeap{}
	heap.Init(feedHeap)

	for _, followed := range u.Following {
		wg.Add(1)
		go func(f *User) {
			defer wg.Done()
			posts := f.GetRecentPostsByUser(f.UserId)

			mu.Lock()
			for _, post := range posts {
				if feedHeap.Len() < 20 {
					heap.Push(feedHeap, post)
				} else if post.TimeStamp.After((*feedHeap)[0].TimeStamp) {
					heap.Pop(feedHeap)
					heap.Push(feedHeap, post)
				}
			}
			mu.Unlock()
		}(followed)
	}

	wg.Wait()

	var result []*Post
	for feedHeap.Len() > 0 {
		result = append(result, heap.Pop(feedHeap).(*Post))
	}
	// for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
	// 	result[i], result[j] = result[j], result[i]
	// }

	return result
}
