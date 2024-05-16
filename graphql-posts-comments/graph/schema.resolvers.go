// graph/schema.resolvers.go

package graph

import (
	"context"
	"fmt"
	"time"

	"github.com/nasl1s/test/graphql-posts-comments/db"
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
	post, exists := r.db.GetPost(id)
	if !exists {
		return nil, fmt.Errorf("post with ID %d not found", id)
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
	// Определяем переменную для хранения родительского ID комментария
	var parentIDValue int
	if parentID != nil {
		parentIDValue = *parentID
	}

	// Добавить новый комментарий в базу данных
	newComment := &db.Comment{
		Content:   content,
		AuthorID:  authorID,
		PostID:    postID,
		ParentID:  parentIDValue, // Используем значение parentIDValue вместо parentID
		CreatedAt: time.Now(),
	}
	r.db.AddComment(newComment)
	return newComment, nil
}

func (r *Resolver) Comments(ctx context.Context, obj *db.Post) ([]*db.Comment, error) {
	// Получаем комментарии для указанного поста из базы данных
	comments, exists := r.db.GetComments(obj.ID)
	if !exists {
		// Если комментарии не существуют, возвращаем пустой срез комментариев и ошибку
		return []*db.Comment{}, fmt.Errorf("comments not found for post with ID %d", obj.ID)
	}
	// Возвращаем комментарии
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
