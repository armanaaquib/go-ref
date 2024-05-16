package blogposts_test

import (
	"errors"
	"github.com/armanaaquib/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFs struct{}

func (StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("should return posts", func(t *testing.T) {
		firstBody := `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World!`
		secondBody := `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		tfs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(tfs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(tfs) {
			t.Errorf("got %d posts, want %d", len(posts), len(tfs))
		}

		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World!`,
		}
		assertPost(t, posts[0], want)
	})

	t.Run("should return error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFs{})

		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
