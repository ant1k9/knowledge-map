{{ template "header" }}

<div class="container text-center">
  <div class="card">
    <div class="card-body">
      <div class="card-title"><a href="{{ .Link }}">{{ .Name }}</a></div>

      {{ if .Docs }}
      <p class="card-text"><a href="{{ .Docs }}">Documentation</a></p>
      {{ end }}
      
      <p class="card-text">{{ .Description }}</p>
      
      {{ if .Examples }}
      <pre class="card-text text-left">
        <code class="language-{{ .ExamplesLanguage }} hljs">{{ .Examples }}</code>
      </pre>
      {{ end }}
      
      <p class="card-footer">
      {{ range .Labels }}
        <a href="/search/{{ . }}">#{{ . }}</a>
      {{ end }}
      </p>
    </div>
  </div>
</div>
