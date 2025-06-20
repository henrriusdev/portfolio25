// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.898
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func NavbarScript() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<script nonce=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `src/components/navbar.templ`, Line: 4, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\">\n\t\t// Función para manejar el menú móvil\n\t\tfunction setupMobileMenu() {\n\t\t\tconst mobileMenuButton = document.getElementById('mobileMenuButton');\n\t\t\tconst mobileMenu = document.getElementById('mobileMenu');\n\t\t\tlet isOpen = false;\n\t\t\t\n\t\t\t// Función para abrir/cerrar el menú\n\t\t\tfunction toggleMenu() {\n\t\t\t\tisOpen = !isOpen;\n\t\t\t\tif (isOpen) {\n\t\t\t\t\tmobileMenu.classList.remove('hidden');\n\t\t\t\t\t// Añadir animación de entrada\n\t\t\t\t\tmobileMenu.classList.add('animate-fade-in-scale');\n\t\t\t\t\tmobileMenuButton.querySelector('svg').classList.add('rotate-180');\n\t\t\t\t} else {\n\t\t\t\t\t// Añadir animación de salida y luego ocultar\n\t\t\t\t\tmobileMenu.classList.add('animate-fade-out-scale');\n\t\t\t\t\tmobileMenuButton.querySelector('svg').classList.remove('rotate-180');\n\t\t\t\t\tsetTimeout(() => {\n\t\t\t\t\t\tmobileMenu.classList.add('hidden');\n\t\t\t\t\t\tmobileMenu.classList.remove('animate-fade-out-scale');\n\t\t\t\t\t}, 150);\n\t\t\t\t}\n\t\t\t}\n\t\t\t\n\t\t\t// Evento para el botón del menú\n\t\t\tmobileMenuButton.addEventListener('click', toggleMenu);\n\t\t\t\n\t\t\t// Cerrar el menú cuando se hace clic fuera\n\t\t\tdocument.addEventListener('click', (event) => {\n\t\t\t\tif (isOpen && !mobileMenuButton.contains(event.target) && !mobileMenu.contains(event.target)) {\n\t\t\t\t\ttoggleMenu();\n\t\t\t\t}\n\t\t\t});\n\t\t}\n\t\t\n\t\t// Inicializar cuando el DOM esté listo\n\t\tdocument.addEventListener('DOMContentLoaded', setupMobileMenu);\n\t\t\n\t\t// Añadir estilos para las animaciones\n\t\tconst style = document.createElement('style');\n\t\tstyle.textContent = `\n\t\t\t@keyframes fadeInScale {\n\t\t\t\tfrom { opacity: 0; transform: scale(0.95); }\n\t\t\t\tto { opacity: 1; transform: scale(1); }\n\t\t\t}\n\t\t\t@keyframes fadeOutScale {\n\t\t\t\tfrom { opacity: 1; transform: scale(1); }\n\t\t\t\tto { opacity: 0; transform: scale(0.95); }\n\t\t\t}\n\t\t\t.animate-fade-in-scale {\n\t\t\t\tanimation: fadeInScale 0.2s ease-out forwards;\n\t\t\t}\n\t\t\t.animate-fade-out-scale {\n\t\t\t\tanimation: fadeOutScale 0.15s ease-in forwards;\n\t\t\t}\n\t\t`;\n\t\tdocument.head.appendChild(style);\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func Navbar() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = NavbarScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<nav class=\"border-b py-3 bg-white dark:bg-gray-800 top-0 sticky z-[1000] transition-colors duration-300\"><div class=\"container mx-auto px-4\"><div class=\"flex items-center justify-between\"><!-- Logo/Name --><div class=\"text-xl font-bold text-gray-800 dark:text-white\">Henrry Bourgeot</div><!-- Desktop Menu --><div class=\"hidden md:flex items-center space-x-4\"><a href=\"/home\" class=\"text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white\">Home</a> <a href=\"/home#about\" class=\"text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white\">About</a> <a href=\"/home#projects\" class=\"text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white\">Projects</a> <a href=\"/home#contact\" class=\"text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white\">Contact</a> <a href=\"/blog\" class=\"text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white\">Blog</a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ThemeSwitcher(ThemeSwitcherProps{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</div><!-- Mobile Menu Button (usando JavaScript vanilla) --><div class=\"md:hidden relative\"><button id=\"mobileMenuButton\" class=\"text-gray-600 hover:text-gray-900 dark:text-gray-300 dark:hover:text-white focus:outline-none\"><svg class=\"h-6 w-6 transition-transform duration-200\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16m-7 6h7\"></path></svg></button><!-- Mobile Menu Dropdown --><div id=\"mobileMenu\" class=\"hidden absolute right-0 mt-2 w-48 rounded-md shadow-lg bg-white dark:bg-gray-700 ring-1 ring-black ring-opacity-5\"><div class=\"py-1 rounded-md bg-white dark:bg-gray-700 shadow-xs\"><a href=\"/home\" class=\"block px-4 py-2 text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600\">Home</a> <a href=\"/home#about\" class=\"block px-4 py-2 text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600\">About</a> <a href=\"/home#projects\" class=\"block px-4 py-2 text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600\">Projects</a> <a href=\"/home#contact\" class=\"block px-4 py-2 text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600\">Contact</a> <a href=\"/blog\" class=\"block px-4 py-2 text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-gray-600\">Blog</a><div class=\"px-4 py-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ThemeSwitcher(ThemeSwitcherProps{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div></div></div></div></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
