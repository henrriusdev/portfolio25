package pages

import (
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/henrriusdev/portfolio/src/layout"
	b "github.com/willoma/bulma-gomponents"
	e "github.com/willoma/gomplements"
	. "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

// HomePage es la página de inicio
func HomePage(experiences ...model.Experience) Node {
	if len(experiences) == 0 {
		endDate := time.Date(2024, 7, 4, 0, 0, 0, 0, time.UTC)
		experiences = append(experiences, model.Experience{
			Company:     "LinkTIC SAS",
			StartDate:   time.Date(2022, 11, 5, 0, 0, 0, 0, time.UTC),
			EndDate:     &endDate,
			Description: "",
			Role:        "Developer",
		})
	}
	return layout.Base("Inicio",
		e.Div(
			b.Container(
				b.Hero(
					b.FullHeightWithNavbar,
					e.Div(
						b.Title(html.P, "Hello, I'm Henrry", b.TextPrimaryDark, b.FontSize(1), b.MarginBottom(b.Spacing6)),
						b.Subtitle(html.P, "Go Backend Developer with more about 2 years professional experience. Also with knowledge about TypeScript, Java and Python and frontend frameworks.", b.TextGreyLighter, b.FontSize(4)),
					),
					b.ImageImg("/static/header.png", b.Small),
				),
			),
			b.Section(
				b.Small,
				e.Class("bg-sky-950"),
				b.Title("Experience", e.ID("experience"), b.TextGreyLighter, b.FontSize(2), b.MarginBottom(b.Spacing3), b.TextCentered),
				b.Content(
					Map(experiences, func(exp model.Experience) Node {
						return b.Section(
							b.Title(exp.Role+". "+exp.StartDate.Format(time.DateOnly)+" - "+exp.EndDate.Format(time.DateOnly), b.TextWhite, b.TitleLevel3),
							b.Subtitle("At "+exp.Company, b.TitleLevel5, b.TextGreyLighter),
						)
					}),
				),
			),
		),
	)
}
