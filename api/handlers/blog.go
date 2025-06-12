package handlers

import (
	"bytes"
	"net/http"
	"strconv"

	chromahtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/henrriusdev/portfolio/api/middleware"
	"github.com/henrriusdev/portfolio/pkg/common"
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/henrriusdev/portfolio/pkg/service"
	"github.com/henrriusdev/portfolio/src/pages"
	enclave "github.com/quailyquaily/goldmark-enclave"
	"github.com/quailyquaily/goldmark-enclave/core"
	img64 "github.com/tenkoh/goldmark-img64"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
	"go.abhg.dev/goldmark/toc"
)

type Blog struct {
	services service.Service
}

func NewBlog(services service.Service) *Blog {
	return &Blog{services}
}

func (b *Blog) RegisterRoutes(f *fuego.Server) {
	// Get default user
	defaultUser, err := b.services.User.GetByID(1)
	if err != nil {
		panic("Default user not found")
	}

	// Public routes
	fuego.Get(f, "", b.List,
		option.Description("Show blog posts"),
		option.Summary("Blog"),
		option.Tags("blog"),
	)
	fuego.Get(f, "/{id}", b.View,
		option.Description("View blog post"),
		option.Summary("View post"),
		option.Tags("blog"),
	)

	// Protected routes
	fuego.Use(f, func(next http.Handler) http.Handler {
		return middleware.AuthMiddleware(http.HandlerFunc(next.ServeHTTP))
	})

	fuego.Get(f, "/new", func(c fuego.ContextNoBody) (fuego.Templ, error) {
		return b.New(c, defaultUser)
	},
		option.Description("New blog post"),
		option.Summary("New post"),
		option.Tags("blog", "admin"),
	)
	fuego.Get(f, "/{id}/edit", b.Edit,
		option.Description("Edit blog post"),
		option.Summary("Edit post"),
		option.Tags("blog", "admin"),
	)
	fuego.Post(f, "/save", func(c fuego.ContextWithBody[model.BlogPost]) (model.BlogPost, error) {
		return b.Save(c, defaultUser)
	},
		option.Description("Save blog post"),
		option.Summary("Save post"),
		option.Tags("blog", "admin"),
	)
	fuego.Delete(f, "/{id}", b.Delete,
		option.Description("Delete blog post"),
		option.Summary("Delete post"),
		option.Tags("blog", "admin"),
	)
}

func (b *Blog) List(c fuego.ContextNoBody) (fuego.Templ, error) {
	posts, err := b.services.Blog.GetAllWithAuthor()
	if err != nil {
		return nil, err
	}
	return pages.Blog(posts), nil
}

func (b *Blog) View(c fuego.ContextNoBody) (fuego.Templ, error) {
	param := c.PathParam("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := b.services.Blog.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	// Convert markdown to HTML with custom rendering
	markdown := goldmark.New(
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithExtensions(
			extension.GFM,
			&toc.Extender{
				Title: "Table of Contents",
			},
			highlighting.NewHighlighting(
				highlighting.WithStyle("tokyonight-night"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&anchor.Extender{
				Texter: anchor.Text("#"),
			},
			img64.Img64,
			enclave.New(&core.Config{}),
		),
		goldmark.WithRendererOptions(
			img64.WithPathResolver(img64.ParentLocalPathResolver("/src/assets/images")),
		),
	)

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(post.Content), &buf); err != nil {
		return nil, err
	}

	return pages.BlogPost(*post, common.Unsafe(buf.String())), nil
}

func (b *Blog) New(c fuego.ContextNoBody, user *model.User) (fuego.Templ, error) {
	return pages.BlogEdit(model.BlogPost{
		AuthorID: user.ID,
		Author:   *user,
	}), nil
}

type EditParams struct {
	ID uint `param:"id"`
}

func (b *Blog) Edit(c fuego.ContextNoBody) (fuego.Templ, error) {
	param := c.PathParam("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return nil, err
	}

	post, err := b.services.Blog.GetByID(uint(id))
	if err != nil {
		return nil, err
	}
	return pages.BlogEdit(*post), nil
}

func (b *Blog) Save(c fuego.ContextWithBody[model.BlogPost], user *model.User) (model.BlogPost, error) {
	post, err := c.Body()
	if err != nil {
		return model.BlogPost{}, err
	}

	// Set author ID
	post.AuthorID = user.ID
	post.Author = *user

	if post.ID > 0 {
		// Update existing post
		err = b.services.Blog.Update(&post)
	} else {
		// Create new post
		err = b.services.Blog.Create(&post)
	}

	if err != nil {
		return model.BlogPost{}, err
	}

	c.Redirect(http.StatusSeeOther, "/blog")
	return post, nil
}

type DeleteParams struct {
	ID uint `param:"id"`
}

func (b *Blog) Delete(c fuego.ContextNoBody) (string, error) {
	param := c.PathParam("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return "", err
	}

	if err := b.services.Blog.Delete(uint(id)); err != nil {
		return "", err
	}

	c.Redirect(http.StatusSeeOther, "/blog")
	return "", nil
}
