package pages

import (
	"github.com/henrriusdev/portfolio/src/layout"
	b "github.com/willoma/bulma-gomponents"
	. "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

// HomePage es la página de inicio
func HomePage() Node {
	return layout.Base("Inicio",
		b.Container(
			b.Section(
				b.Title(Text("Bienvenido a mi Portafolio")),
				b.Subtitle(Text("Desarrollador de software con experiencia en Go.")),
				b.Button(html.Class("is-primary"), Text("Ver Proyectos")),
			),
		),
	)
}
