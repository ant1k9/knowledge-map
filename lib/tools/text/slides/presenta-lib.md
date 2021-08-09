[presenta-lib](https://github.com/presenta-software/presenta-lib)

*Description*: A javascript library to build expressive web presentations in seconds.

*Labels*: #Slides #Javascript

*Docs*: https://lib.presenta.cc/overview/

*Examples*:

```html
<script src="https://unpkg.com/@presenta/lib"></script>

<div id="app"></div>

<script>
new Presenta('#app', {
    scenes: [
        {
            blocks:[
                {
                    type: 'text',
                    text: '<h1>Welcome PRESENTA</h1>'

                },
            ]
        },
        {
            blocks:[
                {
                    type: 'text', 
                    text: '<h1>This <mark>Title</mark></h1>
   <p>This <high>paragraph</high> of <b>text</b>!</p>', 
                    scale:4
                },
            ],
        }
    ]
})
</script>
```
