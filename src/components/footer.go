package components

import . "maragu.dev/gomponents"
import "maragu.dev/gomponents/html"

func Footer() Node {
	return html.Footer(html.Class("bg-gray-200 dark:bg-gray-800 p-4 text-center"),
		Text("© 2025 Mi Portafolio | Hecho con Go"),
	)
}
