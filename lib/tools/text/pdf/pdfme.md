[pdfme](https://github.com/aFelipeSP/pdfme)

*Description*: This is a powerful library to create PDF documents easily.

*Labels*: #PDF #Python

*Docs*: https://pdfme.readthedocs.io/en/latest/tutorial.html

*Examples*:

```python
from pdfme import build_pdf

document = {}

document['running_sections'] = {
    "header": {
        "x": "left", "y": 20, "height": "top",
        "style": {"text_align": "r"},
        "content": [{".b": "This is a header"}]
    },
    "footer": {
        "x": "left", "y": 800, "height": "bottom",
        "style": {"text_align": "c"},
        "content": [{".": ["Page ", {"var": "$page"}]}]
    }
}

document['formats'] = {
    "url": {"c": "blue", "u": 1},
    "title": {"b": 1, "s": 13}
}

document['sections'] = []
section1 = {}
document['sections'].append(section1)

section1['style'] = {
    "page_numbering_style": 'roman'
}

section1['content'] = content1 = []

content1.append({
    ".": "A Title", "style": "title", "label": "title1",
    "outline": {"level": 1, "text": "A different title 1"}
})

content1.append(
    ["This is a paragraph with a ", {".b;c:green": "bold green part"}, ", a ",
    {".": "link", "style": "url", "uri": "https://some.url.com"},
    ", a footnote", {"footnote": "description of the footnote"},
    " and a reference to ",
    {".": "Title 2.", "style": "url", "ref": "title2"}]
)

with open('document.pdf', 'wb') as f:
    build_pdf(document, f)
```
