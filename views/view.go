package views

import "html/template"

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/main.html.tmpl",
		"views/layouts/navbar.html.tmpl",
		"views/layouts/sidebar.html.tmpl",
		"views/layouts/footer.html.tmpl",
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
