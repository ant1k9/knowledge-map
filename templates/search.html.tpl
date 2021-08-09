{{ template "header" }}

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
