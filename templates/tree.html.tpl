{{ template "header" }}

<div class="container text-center">
  {{ range $index, $file := . }}
  <div class="card">
    <div class="card-title"><a href="{{ $file.Path }}">{{ $file.Name }}</a></div>
    {{ if $file.Description }}
    <div class="card-body"><em>Description:</em> {{ $file.Description }}</div>
    {{ end }}
  </div>
  {{ end }}
</div>
