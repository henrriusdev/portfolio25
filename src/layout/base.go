package layout

import (
	"github.com/henrriusdev/portfolio/src/components"
	"maragu.dev/gomponents/html"
)
import g "maragu.dev/gomponents"

func Base(title string, content g.Node) g.Node {
	return html.HTML(
		html.Head(
			Head(title),
		),
		html.Body(
			html.Div(html.Class("min-h-screen flex flex-col bg-gray-50 dark:bg-gray-900"),
				components.Navbar(),
				html.Main(html.Class("flex-grow"), content),
				components.Footer(),
			),
		),
	)
}
