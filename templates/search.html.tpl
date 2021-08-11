{{ template "header" }}

<div class="container">
  {{ range $index, $element := . }}
  {{ if mod $index 0 3 }}<div class="card-group">{{ end }}
  <div class="card">
    <div class="card-body">
      <p class="card-title">
        <a href="{{ $element.Path }}">{{ $element.Name }}</a>
      </p>
      <div class="card-text">{{ $element.Description }}</div>
    </div>
  </div>
  {{ if mod $index 1 3 }}</div>{{ end }}
  {{ end }}
</div>
