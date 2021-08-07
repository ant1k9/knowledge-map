<div>
  <a href="/search/">Home</a><span> </span>
  <a href="/tree/">Menu</a><span> </span>
</div><br/>

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
