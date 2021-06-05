/*
 * @lc app=leetcode id=355 lang=golang
 *
 * [355] Design Twitter
 */

// @lc code=start
import (
	"container/heap"
)

var time = 0

type Tweet struct {
	id   int
	time int
}

type Twitter struct {
	users  map[int]map[int]struct{}
	tweets map[int][]Tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users:  make(map[int]map[int]struct{}),
		tweets: make(map[int][]Tweet),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if v, ok := this.tweets[userId]; ok {
		this.tweets[userId] = append([]Tweet{Tweet{id: tweetId, time: time}}, v...)
	} else {
		this.tweets[userId] = []Tweet{Tweet{id: tweetId, time: time}}
	}
	time++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := TweetHeap{}
	for _, tweet := range this.tweets[userId] {
		h = append(h, tweet)
	}
	for followee, _ := range this.users[userId] {
		for _, tweet := range this.tweets[followee] {
			h = append(h, tweet)
		}
	}
	res := []int{}
	cnt := 10
	if len(h) < 10 {
		cnt = len(h)
	}
	heap.Init(&h)
	for i := 0; i < cnt; i++ {
		res = append(res, heap.Pop(&h).(Tweet).id)
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.users[followerId]; !ok {
		this.users[followerId] = make(map[int]struct{})
	}
	this.users[followerId][followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.users[followerId], followeeId)
}

type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

