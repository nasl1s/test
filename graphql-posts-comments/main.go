package main

import (
	"log"
	"net/http"
	"time"
)

// Post представляет структуру данных для поста.
type Post struct {
	ID            int
	Title         string
	Content       string
	AuthorID      int
	Comments      []*Comment
	AllowComments bool
}

// Comment представляет структуру данных для комментария.
type Comment struct {
	ID        int
	Content   string
	AuthorID  int
	ParentID  int
	PostID    int
	Children  []*Comment
	CreatedAt time.Time
}

type InMemoryDB struct {
	posts        map[int]*Post
	comments     map[int][]*Comment
	postCount    int
	commentCount int
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		posts:        make(map[int]*Post),
		comments:     make(map[int][]*Comment),
		postCount:    0,
		commentCount: 0,
	}
}

func (db *InMemoryDB) AddPost(post *Post) {
	db.posts[post.ID] = post
	db.postCount++
}

func (db *InMemoryDB) GetPost(postID int) (*Post, bool) {
	post, exists := db.posts[postID]
	return post, exists
}

func (db *InMemoryDB) AddComment(comment *Comment) {
	postComments := db.comments[comment.PostID]
	postComments = append(postComments, comment)
	db.comments[comment.PostID] = postComments
	db.commentCount++
}

func (db *InMemoryDB) GetComments(postID int) ([]*Comment, bool) {
	comments, exists := db.comments[postID]
	return comments, exists
}

func main() {
	db := NewInMemoryDB()

	// Пример добавления поста
	post := &Post{
		ID:            1,
		Title:         "Example Post",
		Content:       "This is an example post content.",
		AuthorID:      123,
		AllowComments: true,
	}
	db.AddPost(post)

	// Пример добавления комментария
	comment := &Comment{
		ID:        1,
		Content:   "This is a comment.",
		AuthorID:  456,
		PostID:    1,
		CreatedAt: time.Now(),
	}
	db.AddComment(comment)

	// Пример получения поста и комментариев
	fetchedPost, exists := db.GetPost(1)
	if exists {
		log.Printf("Fetched post: %+v", fetchedPost)
		comments, _ := db.GetComments(1)
		for _, c := range comments {
			log.Printf("Comment: %+v", c)
		}
	} else {
		log.Println("Post not found")
	}

	// Начало HTTP-сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Обработка HTTP-запросов здесь
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
