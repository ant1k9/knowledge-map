{{ template "header" }}

<div>
  <a href="{{ .Link }}">{{ .Name }}</a>
</div>
<div>
  <em>Description:</em> {{ .Description }}
</div>

{{ if .Docs }}
<div>
  <em>Docs:</em> {{ .Docs }}
</div>
{{ end }}

{{ if .Examples }}
<div><br/>
  <em>Examples:</em>
  <div>
    <pre><code>{{ .Examples }}</code></pre>
  </div>
</div>
{{ end }}

<div>
{{ range .Labels }}
  <a href="/search/{{ . }}">#{{ . }}</a>
{{ end }}
</div>
