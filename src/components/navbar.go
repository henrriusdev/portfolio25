package components

import (
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
	. "maragu.dev/gomponents"
)

// Navbar genera la barra de navegación con Bulma
func Navbar() Node {
	return b.Navbar(
		e.AriaNavigation,
		b.FixedTop,
		b.Info,
		e.AriaLabel("main navigation"),
		b.NavbarBrand(
			b.NavbarAHref(
				"/",
				"Henrry Bourgeot",
				b.FontSize(3),
				b.TextBlack,
				b.WeightBold,
				b.PaddingHorizontal(b.Spacing5),
				b.Monospace,
			),
		),
		b.NavbarStart(
			b.NavbarAHref("#", "Home", b.TextBlack),
			b.NavbarAHref("#", "Documentation", b.TextBlack),
			b.NavbarDropdown(
				b.OnLink(
					e.Span,
					b.TextBlack,
					"More",
				),
				b.Hoverable,
				b.NavbarAHref("#", "About"),
				b.NavbarAHref("#", "Jobs"),
				b.NavbarAHref("#", "Contact"),
				b.NavbarDivider(),
				b.NavbarAHref("#", "Report an issue"),
			),
		),
		b.NavbarEnd(
			b.NavbarItem(
				b.Buttons(
					b.ButtonA(
						b.Primary,
						e.Strong("Sign up"),
					),
					b.ButtonA(
						b.Light,
						"Log in",
					),
				),
			),
		),
	)
}
