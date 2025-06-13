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
	fuego.Get(f, "/search", b.Search,
		option.Description("Search blog posts"),
		option.Summary("Search posts"),
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
	fuego.Post(f, "/save", func(c fuego.ContextWithBody[BlogPostForm]) (model.BlogPost, error) {
		return b.Save(c, defaultUser)
	},
		option.Description("Save blog post"),
		option.Summary("Save post"),
		option.Tags("blog", "admin"),
	)
	fuego.Post(f, "/delete", b.Delete,
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
				Title:  "Table of Contents",
				ListID: "table-of-contents",
			},
			highlighting.NewHighlighting(
				highlighting.WithStyle("doom-one2"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
					chromahtml.LineNumbersInTable(false),
					chromahtml.TabWidth(2),
					chromahtml.Standalone(true),
					chromahtml.WrapLongLines(true),
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
	// Get all categories for the form
	categories, err := b.services.Category.GetAll()
	if err != nil {
		return nil, err
	}

	return pages.BlogEdit(model.BlogPost{
		AuthorID: user.ID,
		Author:   *user,
	}, categories), nil
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

	// Get post with categories
	post, err := b.services.Blog.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	// Get all categories for the form
	categories, err := b.services.Category.GetAll()
	if err != nil {
		return nil, err
	}

	return pages.BlogEdit(*post, categories), nil
}

// BlogPostForm representa el formulario para crear o editar un post de blog
type BlogPostForm struct {
	ID         uint     `form:"id"`
	Title      string   `form:"title"`
	Content    string   `form:"content"`
	Categories []string `form:"categories"`
}

func (b *Blog) Save(c fuego.ContextWithBody[BlogPostForm], user *model.User) (model.BlogPost, error) {
	// Get form data
	form, err := c.Body()
	if err != nil {
		return model.BlogPost{}, err
	}

	// Create or get post
	var post model.BlogPost
	if form.ID > 0 {
		// Get existing post for update
		existingPost, err := b.services.Blog.GetByID(form.ID)
		if err != nil {
			return model.BlogPost{}, err
		}
		post = *existingPost
	}

	// Update post fields
	post.ID = form.ID
	post.Title = form.Title
	post.Content = form.Content
	post.AuthorID = user.ID
	post.Author = *user
	post.Categories = []model.Category{} // Initialize empty slice

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

	// Clear existing categories for this post
	// This is a simple approach - in a production app you might want to
	// only update what changed to avoid unnecessary database operations
	b.services.Blog.ClearCategories(post.ID)

	// Add selected categories
	for _, categoryIDStr := range form.Categories {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			continue
		}
		b.services.Category.AddCategoryToPost(post.ID, uint(categoryID))
	}

	c.Redirect(http.StatusSeeOther, "/blog")
	return post, nil
}

type DeleteRequest struct {
	ID uint `form:"id"`
}

func (b *Blog) Delete(c fuego.ContextWithBody[DeleteRequest]) (string, error) {
	req, err := c.Body()
	if err != nil {
		return "", err
	}

	id := req.ID

	if err := b.services.Blog.Delete(uint(id)); err != nil {
		return "", err
	}

	c.Redirect(http.StatusSeeOther, "/blog")
	return "", nil
}

// Search busca posts por término de búsqueda
func (b *Blog) Search(c fuego.ContextNoBody) (fuego.Templ, error) {
	// Obtener el parámetro de búsqueda de la URL
	query := c.QueryParam("q")

	var posts []model.BlogPost
	var err error

	if query != "" {
		// Si hay un término de búsqueda, buscar posts que coincidan
		posts, err = b.services.Blog.SearchPosts(query)
	} else {
		// Si no hay término de búsqueda, mostrar todos los posts
		posts, err = b.services.Blog.GetAllWithAuthor()
	}

	if err != nil {
		return nil, err
	}

	// Renderizar la plantilla de búsqueda con los resultados
	return pages.BlogSearch(posts, query), nil
}
