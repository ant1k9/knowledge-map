{{ template "header" }}

{{ range $path, $file := . }}
<div>
  <div>
    <a href="{{ $path }}">{{ $file.Name }}</a>
  </div>
  <div>
    {{ if $file.Description }}
    <em>Description:</em> {{ $file.Description }}
    {{ end }}
  </div>
</div>
<br/>
{{ end }}
