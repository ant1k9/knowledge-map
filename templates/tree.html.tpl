<div>
  <a href="/search/">Home</a><span> </span>
  <a href="/tree/">Menu</a><span> </span>
</div><br/>

{{ range $path, $file := . }}
<div>
  <div>
    <a href="{{ $file.Link }}">{{ $file.Name }}</a>
  </div>
  <div>
    {{ if $file.Description }}
    <em>Description:</em> {{ $file.Description }}
    <div>
    {{ range .Labels }}
      <a href="/search/{{ . }}">#{{ . }}</a>
    {{ end }}
    </div>
    {{ end }}
  </div>
</div>
<br/>
{{ end }}
