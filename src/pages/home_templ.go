// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.898
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

func HomePage(projects []model.Project, experiences []model.Experience, technologies []model.Technology, links []model.Contact) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<main><!-- Hero Section --><header class=\"min-h-screen relative bg-gradient-to-b from-gray-50 to-white dark:from-gray-900 dark:to-gray-800 overflow-hidden transition-colors duration-300\"><!-- Decorative background elements --><div class=\"absolute inset-0 overflow-hidden\"><div class=\"absolute -top-40 -right-40 w-80 h-80 bg-blue-100 dark:bg-blue-900/20 rounded-full blur-3xl opacity-70\"></div><div class=\"absolute -bottom-40 -left-40 w-80 h-80 bg-indigo-100 dark:bg-indigo-900/20 rounded-full blur-3xl opacity-70\"></div></div><div class=\"container mx-auto px-4 relative z-10\"><div class=\"flex flex-col lg:flex-row items-center justify-between min-h-screen gap-12 py-16\"><!-- Text Content --><div class=\"w-full lg:w-1/2 z-10 animate-fadeIn\"><div class=\"space-y-6\"><span class=\"px-4 py-2 bg-blue-100 dark:bg-blue-900/30 text-blue-800 dark:text-blue-200 rounded-full text-sm font-medium inline-block\">Go Backend Developer</span><h1 class=\"text-5xl md:text-6xl lg:text-7xl font-bold text-gray-900 dark:text-white leading-tight\">Hello, I am <span class=\"text-blue-600 dark:text-blue-400\">Henrry</span></h1><p class=\"text-xl md:text-2xl text-gray-600 dark:text-gray-300 max-w-xl leading-relaxed\">Go Backend Developer with 2+ years professional experience. Skilled in TypeScript, Java, Python and frontend frameworks.</p><div class=\"flex flex-wrap gap-4 pt-4\"><a href=\"#contact\" class=\"bg-blue-600 hover:bg-blue-700 text-white px-8 py-4 rounded-xl font-medium inline-block transition-all duration-300 transform hover:scale-105 hover:shadow-lg\">Get in touch</a> <a href=\"#projects\" class=\"border-2 border-blue-600 text-blue-600 hover:bg-blue-600 hover:text-white px-8 py-4 rounded-xl font-medium inline-block transition-all duration-300 transform hover:scale-105 hover:shadow-lg\">View Portfolio</a></div></div></div><!-- Image Side --><div class=\"w-full lg:w-1/2 h-full\"><div class=\"relative aspect-square transform hover:scale-[1.02] transition-transform duration-500\"><div class=\"absolute inset-0 bg-gradient-to-tr from-blue-500/10 to-indigo-500/10 dark:from-blue-500/20 dark:to-indigo-500/20 rounded-3xl transform rotate-6\"></div><img src=\"/static/header.png\" alt=\"Hero Image\" class=\"rounded-3xl shadow-2xl object-cover w-full h-full relative z-10\"><!-- Decorative elements --><div class=\"absolute -z-10 -bottom-4 -right-4 w-full h-full bg-blue-600/10 rounded-3xl\"></div><div class=\"absolute -z-20 -bottom-8 -right-8 w-full h-full bg-indigo-600/5 rounded-3xl\"></div></div></div></div></div></header><!-- Experience Section --><section class=\"py-20 relative\" id=\"about\"><!-- Decorative background elements --><div class=\"absolute inset-0 overflow-hidden\"><div class=\"absolute top-20 left-20 w-64 h-64 bg-blue-50 dark:bg-blue-900/10 rounded-full blur-3xl opacity-70\"></div></div><div class=\"max-w-[900px] w-full mx-auto px-4\"><div class=\"text-center mb-16\"><span class=\"px-4 py-2 bg-indigo-100 dark:bg-indigo-900/30 text-indigo-800 dark:text-indigo-200 rounded-full text-sm font-medium inline-block mb-4\">My Journey</span><h2 class=\"text-5xl font-bold text-gray-900 dark:text-white mb-4\">Work Experience</h2><p class=\"text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto\">A timeline of my professional journey and the skills I've developed along the way.</p></div><div class=\"transform hover:scale-[1.01] transition-all duration-300\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.ExperienceTimeline(experiences, false).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div></section><!-- Projects Section --><section class=\"py-24 bg-gradient-to-b from-white to-gray-50 dark:from-gray-800 dark:to-gray-900 relative overflow-hidden\" id=\"projects\"><!-- Decorative elements --><div class=\"absolute inset-0 overflow-hidden\"><div class=\"absolute -top-40 left-1/3 w-72 h-72 bg-blue-50 dark:bg-blue-900/10 rounded-full blur-3xl opacity-70\"></div><div class=\"absolute bottom-0 right-0 w-full h-1/2 bg-gradient-to-t from-gray-50/80 to-transparent dark:from-gray-900/80\"></div></div><div class=\"container mx-auto px-4 relative z-10\"><div class=\"text-center mb-16\"><span class=\"px-4 py-2 bg-blue-100 dark:bg-blue-900/30 text-blue-800 dark:text-blue-200 rounded-full text-sm font-medium inline-block mb-4\">My Work</span><h2 class=\"text-5xl font-bold text-gray-900 dark:text-white mb-4\">Featured Projects</h2><p class=\"text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto\">A showcase of my recent work and the technologies I've used.</p></div><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 lg:gap-10 stagger-children\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, project := range projects {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"transform hover:scale-[1.03] hover:-rotate-1 transition-all duration-300 animate-fadeIn\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = components.ProjectCard(project, false).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div></div></section><!-- Blog Section --><!-- <sectionon class=\"py-12\">\n\t\t\t\t<div class=\"container mx-auto px-4\">\n\t\t\t\t\t<h2 class=\"text-3xl font-bold mb-8 text-center\">Latest from the Blog</h2>\n\t\t\t\t\t<div class=\"max-w-2xl mx-auto\">\n\t\t\t\t\t\t@BlogPreview(BlogPost{\n                        Title: \"Understanding Modern Web Development\",\n                        Preview: \"In this post, we explore the latest trends...\",\n                        Date: \"Mar 15, 2024\",\n                        ReadTime: \"5 min\",\n                        ImageURL: \"/assets/blog1.png\",\n                        Author: \"Your Name\",\n                        AuthorImg: \"/assets/author.png\",\n                    })\n\t\t\t\t\t</div>\n\t\t\t\t</div>\n\t\t\t</sectionon> --><!-- Technologies Section --><section class=\"py-24 bg-gray-50 dark:bg-gray-800 relative overflow-hidden\"><!-- Decorative elements --><div class=\"absolute inset-0 overflow-hidden\"><div class=\"absolute top-1/4 right-1/4 w-64 h-64 bg-indigo-50 dark:bg-indigo-900/10 rounded-full blur-3xl opacity-60\"></div><div class=\"absolute bottom-0 left-0 w-full h-1/3 bg-gradient-to-t from-gray-50 to-transparent dark:from-gray-800 dark:to-transparent\"></div></div><div class=\"container mx-auto px-4 relative z-10\"><div class=\"text-center mb-16\"><span class=\"px-4 py-2 bg-purple-100 dark:bg-purple-900/30 text-purple-800 dark:text-purple-200 rounded-full text-sm font-medium inline-block mb-4\">My Toolkit</span><h2 class=\"text-5xl font-bold text-gray-900 dark:text-white mb-4\">Technologies</h2><p class=\"text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto\">The tools and technologies I work with to bring ideas to life.</p></div><div class=\"grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-6 gap-8 lg:gap-10 stagger-children\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, tech := range technologies {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"flex flex-col items-center bg-white dark:bg-gray-700 rounded-xl p-6 shadow-sm hover:shadow-xl transition-all duration-300 transform hover:scale-105 hover:-translate-y-1 animate-fadeIn\"><div class=\"w-16 h-16 flex items-center justify-center mb-4 animate-float\"><img src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(tech.Icon)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `src/pages/home.templ`, Line: 146, Col: 25}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" alt=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(tech.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `src/pages/home.templ`, Line: 147, Col: 25}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" class=\"w-12 h-12 object-contain\"></div><p class=\"text-base font-medium text-center text-gray-800 dark:text-gray-200\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(tech.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `src/pages/home.templ`, Line: 151, Col: 97}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</p></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div></div></section></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.Footer(links...).Render(ctx, templ_7745c5c3_Buffer)
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
