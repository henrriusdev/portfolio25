// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/henrriusdev/portfolio/src/components"
import "github.com/henrriusdev/portfolio/src/layout"
import "github.com/henrriusdev/portfolio/pkg/model"
import "time"

var seraStart = time.Date(2024, 8, 5, 0, 0, 0, 0, time.UTC)
var linkticStart = time.Date(2022, 12, 5, 0, 0, 0, 0, time.UTC)

func HomePage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<main><!-- Hero Section --><section class=\"min-h-screen relative\"><div class=\"container mx-auto px-4\"><div class=\"flex flex-col lg:flex-row items-center justify-between min-h-screen gap-8 py-16\"><!-- Text Content --><div class=\"w-full lg:w-1/2 z-10\"><h1 class=\"text-4xl md:text-5xl lg:text-6xl font-bold mb-6 text-gray-900 dark:text-white\">Hello, I'm Henrry</h1><p class=\"text-lg md:text-xl mb-8 text-gray-700 dark:text-gray-300 max-w-lg\">Go Backend Developer with more about 2 years professional experience. Also with knowledge about TypeScript, Java and Python and frontend frameworks.</p><div class=\"flex gap-4\"><a href=\"#contact\" class=\"bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-lg inline-block transition-colors\">Get in touch</a> <a href=\"#portfolio\" class=\"border-2 border-blue-600 text-blue-600 hover:bg-blue-600 hover:text-white px-8 py-3 rounded-lg inline-block transition-colors\">View Portfolio</a></div></div><!-- Image Side --><div class=\"w-full lg:w-1/2 h-full\"><div class=\"relative aspect-square\"><img src=\"/static/header.png\" alt=\"Hero Image\" class=\"rounded-2xl shadow-2xl object-cover w-full h-full\"><!-- Optional decorative elements --><div class=\"absolute -z-10 top-4 right-4 w-full h-full bg-blue-600/10 rounded-2xl\"></div><div class=\"absolute -z-20 top-8 right-8 w-full h-full bg-blue-600/5 rounded-2xl\"></div></div></div></div></div></section><!-- Experience Section -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.ExperienceTimeline([]model.Experience{
				{
					Role:        "Senior Developer",
					Company:     "Tech Corp",
					Description: "Led development of key features...",
					StartDate:   linkticStart,
					EndDate:     &seraStart,
				},
				{
					Role:        "Full Stack Developer",
					Company:     "Start Up",
					Description: "Developed full stack applications...",
					StartDate:   seraStart,
					EndDate:     nil,
				},
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!-- Projects Section --><section class=\"py-12 bg-gray-50 dark:bg-gray-900\"><div class=\"container mx-auto px-4\"><h2 class=\"text-3xl font-bold mb-8 text-center\">Projects</h2><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.ProjectCard(model.Project{
				Title:       "Project 1",
				Description: "A brief description of the project...",
				ImageURL:    "/assets/project1.png",
				URL:         "https://project1.com",
				Repo:        "https://github.com/username/project1",
				Techs: []model.Technology{
					{Name: "React"},
					{Name: "Node.js"},
					{Name: "MongoDB"},
				},
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.ProjectCard(model.Project{
				Title:       "Project 2",
				Description: "Another amazing project...",
				ImageURL:    "/assets/project2.png",
				URL:         "https://project2.com",
				Repo:        "https://github.com/username/project2",
				Techs: []model.Technology{
					{Name: "Vue.js"},
					{Name: "Firebase"},
					{Name: "Tailwind"},
				},
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div></div></section><!-- Blog Section --><section class=\"py-12\"><div class=\"container mx-auto px-4\"><h2 class=\"text-3xl font-bold mb-8 text-center\">Latest from the Blog</h2><div class=\"max-w-2xl mx-auto\"></div></div></section></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.Footer().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layout.Base("Portfolio").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
