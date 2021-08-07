<div>
  <a href="/search/">Home</a><span> </span>
  <a href="/tree/">Menu</a><span> </span>
</div><br/>

<div>
  <form method="get">
    <input name=q>
  </form>
</div>

{{ range . }}
<div>
  <div>
    <a href="{{ .Path }}">{{ .Name }}</a>
  </div>
  <div>
    <em>Description:</em> {{ .Description }}
  </div>
</div>
<br/>
{{ end }}
