package components

import (
	"fmt"
	"time"

	"github.com/delaneyj/gomponents-iconify/iconify/mdi"
	"github.com/henrriusdev/portfolio/pkg/model"
	"github.com/willoma/bulma-gomponents"
	. "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func TimelineItem(exp model.Experience) Node {
	endDate := "Current"
	if exp.EndDate != nil {
		endDate = exp.EndDate.Format(time.DateOnly)
	}

	timeText := fmt.Sprintf("%s - %s", exp.StartDate.Format(time.DateOnly), endDate)

	return bulma.Content(
		bulma.Medium,
		bulma.Padding(bulma.Spacing4),
		html.Div(
			html.Class("is-flex is-align-items-center mb-2"),
			mdi.Calendar(html.Class("mr-2 has-text-info")),
			html.P(
				html.Class("is-size-7 has-text-grey"),
				Text(timeText),
			),
		),
		html.H3(
			html.Class("title is-4 mb-5"),
			Text(exp.Role),
		),
		html.H4(
			html.Class("subtitle is-6 mb-3"),
			Text(exp.Company),
		),
		html.P(
			html.Class("has-text-grey-dark"),
			Text(exp.Description),
		),
	)
}

func Timeline(items ...Node) Node {
	return html.Div(items...) // Spread the slice when passing to html.Div
}
