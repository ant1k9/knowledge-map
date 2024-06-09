{{ template "header" . }}

{{ with .ResponseFiles }}
<div class="container text-center">
  {{ range $index, $file := . }}
  {{ if mod $index 0 3 }}<div class="card-group">{{ end }}
  <div class="card">
    <h5 class="card-title pt-3"><a href="{{ $file.Path }}">{{ $file.Name }}</a></h5>
    {{ if $file.Description }}
    <div class="card-body"> {{ $file.Description }}</div>
    {{ end }}
  </div>
  {{ if mod $index 1 3 }}</div>{{ end }}
  {{ end }}
</div>
{{ end }}
