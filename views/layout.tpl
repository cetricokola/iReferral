<!doctype html>
<html class="no-js" lang="en">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">

    <title>{{.Title}}</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="shortcut icon" href="/static/img/logo.png" />
    <!-- Font Awesome -->
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.7.0/css/all.css">
    
    {{range .HeadStyles}}
    <link rel="stylesheet" href="{{.}}">
    {{end}}


</head>

<body>
    {{.LayoutContent}}
   
{{.Footer}}
    {{range .HeadScripts}}
    <script src="{{.}}"></script>
    {{end}}
</body>