[klipse](https://github.com/viebel/klipse)

*Description*: Klipse is a JavaScript plugin for embedding interactive code snippets in tech blogs.

*Labels*: #JavaScript #HTML

*Examples*:

```html
<link rel="stylesheet" type="text/css" href="https://storage.googleapis.com/app.klipse.tech/css/codemirror.css">

<script>
window.klipse_settings = {
    selector_golang: '.language-klipse-go', // css selector for the html elements you want to klipsify
};
</script>
<script src="https://storage.googleapis.com/app.klipse.tech/plugin_prod/js/klipse_plugin.min.js"></script>

<p>And now, we are going to <strong>klipsify</strong> this code snippet:</p>

<pre><code class="language-klipse-go">
import "fmt"

func main() {
      fmt.Println("Hello World!")
}
</code></pre>
```
