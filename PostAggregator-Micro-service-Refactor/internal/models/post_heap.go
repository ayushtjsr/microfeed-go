package models

type PostHeap []*Post

func (h PostHeap) Len() int           { return len(h) }
func (h PostHeap) Less(i, j int) bool { return h[i].TimeStamp.After(h[j].TimeStamp) }
func (h PostHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PostHeap) Push(x any) {
	*h = append(*h, x.(*Post))
}

func (h *PostHeap) Pop() any {
	old := *h
	n := len(old)
	post := old[n-1]
	*h = old[0 : n-1]
	return post
}
