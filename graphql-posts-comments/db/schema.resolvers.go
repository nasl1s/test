// graph/schema.resolvers.go

package graph

import (
	"context"

	"github.com/your_username/your_project_name/db"
)

type Resolver struct {
	db *db.InMemoryDB
}

func (r *Resolver) Posts(ctx context.Context) ([]*db.Post, error) {
	// Вернуть все посты из базы данных
}

func (r *Resolver) Post(ctx context.Context, id int) (*db.Post, error) {
	// Вернуть пост с указанным ID из базы данных
}

func (r *Resolver) AddPost(ctx context.Context, title string, content string, authorID int, allowComments bool) (*db.Post, error) {
	// Добавить новый пост в базу данных
}

func (r *Resolver) AddComment(ctx context.Context, content string, authorID int, postID int, parentID *int) (*db.Comment, error) {
	// Добавить новый комментарий в базу данных
}

func (r *Resolver) Comments(ctx context.Context, obj *db.Post) ([]*db.Comment, error) {
	// Вернуть комментарии для указанного поста из базы данных
}

func (r *Resolver) CommentAdded(ctx context.Context, postID int) (<-chan *db.Comment, error) {
	// Создать и вернуть канал для подписки на новые комментарии к посту с указанным ID
}
