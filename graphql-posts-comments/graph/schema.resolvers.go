// graph/schema.resolvers.go

package graph

import (
	"context"
	"time"

	"github.com/your_username/your_project_name/db"
)

type Resolver struct {
	db *db.InMemoryDB
}

func (r *Resolver) Posts(ctx context.Context) ([]*db.Post, error) {
	// Вернуть все посты из базы данных
	posts := r.db.GetAllPosts() // Предположим, что у вас есть метод GetAllPosts в вашей InMemoryDB
	return posts, nil
}

func (r *Resolver) Post(ctx context.Context, id int) (*db.Post, error) {
	// Вернуть пост с указанным ID из базы данных
	post, err := r.db.GetPost(id) // Предположим, что у вас есть метод GetPost в вашей InMemoryDB
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *Resolver) AddPost(ctx context.Context, title string, content string, authorID int, allowComments bool) (*db.Post, error) {
	// Добавить новый пост в базу данных
	newPost := &db.Post{
		Title:         title,
		Content:       content,
		AuthorID:      authorID,
		AllowComments: allowComments,
	}
	r.db.AddPost(newPost) // Предположим, что у вас есть метод AddPost в вашей InMemoryDB
	return newPost, nil
}

func (r *Resolver) AddComment(ctx context.Context, content string, authorID int, postID int, parentID *int) (*db.Comment, error) {
	// Добавить новый комментарий в базу данных
	newComment := &db.Comment{
		Content:   content,
		AuthorID:  authorID,
		PostID:    postID,
		ParentID:  parentID,
		CreatedAt: time.Now(),
	}
	r.db.AddComment(newComment) // Предположим, что у вас есть метод AddComment в вашей InMemoryDB
	return newComment, nil
}

func (r *Resolver) Comments(ctx context.Context, obj *db.Post) ([]*db.Comment, error) {
	// Вернуть комментарии для указанного поста из базы данных
	comments, err := r.db.GetComments(obj.ID) // Предположим, что у вас есть метод GetComments в вашей InMemoryDB
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Resolver) CommentAdded(ctx context.Context, postID int) (<-chan *db.Comment, error) {
	// Создать и вернуть канал для подписки на новые комментарии к посту с указанным ID
	commentChan := make(chan *db.Comment)
	// В этом месте вы можете добавить логику для создания канала подписки на новые комментарии
	// Например, вы можете добавить этот канал в список каналов для определенного поста и отправлять
	// новые комментарии в этот канал, когда они добавляются в базу данных
	return commentChan, nil
}
