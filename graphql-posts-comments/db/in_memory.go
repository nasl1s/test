// db/in_memory.go

package db

import (
	"sync"
	"time"
)

// Определяем структуру для поста
type Post struct {
	ID            int
	Title         string
	Content       string
	AuthorID      int
	Comments      []*Comment
	AllowComments bool
}

// Определяем структуру для комментария
type Comment struct {
	ID        int
	Content   string
	AuthorID  int
	ParentID  int
	PostID    int
	Children  []*Comment
	CreatedAt time.Time
}

// Определяем in-memory базу данных
type InMemoryDB struct {
	posts        map[int]*Post
	comments     map[int][]*Comment
	postCount    int
	commentCount int
	mu           sync.RWMutex // Для обеспечения безопасности доступа к данным
}

// Создаем новую in-memory базу данных
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		posts:        make(map[int]*Post),
		comments:     make(map[int][]*Comment),
		postCount:    0,
		commentCount: 0,
	}
}

// Добавляем пост в базу данных
func (db *InMemoryDB) AddPost(post *Post) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.posts[post.ID] = post
	db.postCount++
}

// Получаем пост из базы данных по его ID
func (db *InMemoryDB) GetPost(postID int) (*Post, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	post, exists := db.posts[postID]
	return post, exists
}

// Добавляем комментарий к посту в базу данных
func (db *InMemoryDB) AddComment(comment *Comment) {
	db.mu.Lock()
	defer db.mu.Unlock()
	postComments := db.comments[comment.PostID]
	postComments = append(postComments, comment)
	db.comments[comment.PostID] = postComments
	db.commentCount++
}

// Получаем комментарии для поста из базы данных
func (db *InMemoryDB) GetComments(postID int) ([]*Comment, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	comments, exists := db.comments[postID]
	return comments, exists
}
