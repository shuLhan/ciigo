package ciigo

const templateSearch = `
<h3> Search result </h3>
{{range $result := .}}
<h4>
<a href="{{$result.Path}}">{{$result.Path}}</a>
</h4>
	{{range $result.Snippets}}
	<p>... {{.}} ...</p>
	{{end}}
{{end}}
`
