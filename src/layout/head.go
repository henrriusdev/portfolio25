package layout

import "maragu.dev/gomponents/html"
import . "maragu.dev/gomponents"

// Head genera las metaetiquetas y estilos
func Head(title string) Node {
	return html.Head(
		html.Title(title),
		html.Meta(html.Charset("UTF-8")),
		html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
		html.Link(html.Rel("stylesheet"), html.Href("/static/app.css")),
		html.Link(html.Rel("stylesheet"), html.Href("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css")),
	)
}
