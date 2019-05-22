# Create and serve a UI with go and go templates

We already saw that go has a built-in http server that allows you to create web applications. While you will likely use one of the many UI frameworks to do client side rendering, go has a template processing engine that works well for simple UIs.

## go templates

Go templates are powerful for customizing output. Some of the uses include

- html templates
- email templates
- kubernetes definitions

With go you easily bind the templates with data from an API, json/yaml, etc. and produce customized output. You can create custom template functions or use template libraries like <https://gohugo.io/templates/> that further extend templating in go.

For more info see <https://blog.gopheracademy.com/advent-2017/using-go-templates/>

## go-slalom ui

go-slalom provides a simple welcome page to illustrate go templates. The `templates` folder contains the files below:

- layout.html
- index.html
- welcome.html

The `layout.html` looks like below

```html
{{ define "layout" }}
<!DOCTYPE html>

<html>
   <head>
         <meta charset="UTF-8">
         <link rel="stylesheet" href="/static/stylesheets/main.css">
         <title>Welcome Slalom</title>
   </head>
   <body>
      {{ template "content" . }}
   </body>
</html>
{{ end }}
```

The `{{ define "layout" }}` names the `layout` template. The `{{ template "content" . }}` includes the *nested* `content` template. Nested templates are powerful and allow you to create complex layouts. Note the `.` in `{{ template "content" . }}`. This passes the current *pipeline* data to the nested template. This guide will not cover template pipelines, but you can just think of it as how data is passed around when processing templates.

Note that both the `index.html` and `welcome.html` define a `content` template. Which one is incldued? It depends on the route. See [pkg/api/index.go](../pkg/api/index.go) which handles the `/` route for a web browser. The handler function parses both the `layout.html` and `index.html`. So the `content` template from `index.html` is applied. Where as the handler function for `/welcome` uses `layout.html` and `welcome.html`

## Using API to customize output

The `welcome.html` appears below

```html
{{ define "content" }}
<div class="welcome center">Welcome {{.Name}}, it is {{.Time}}</div>
{{ end }}
```

The `{{.Name}}` and `{{.Time}}` values come from data that is applied in the `welcome.go` handler function. While Name is served from form data and Time is coming from a go library, you can see how `welcome.go` could call an API to provide data. For simple UIs using server side rendering this may be good enough.