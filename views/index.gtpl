<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <title>Steamer</title>

    <!-- Bootstrap -->
	<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.4/css/bootstrap.min.css">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="http://cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="http://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
	<div class="warpper">
    <h1>Steamer</h1>
	
	<form class="form-inline" action="/addurl" method="POST">
		<div class="">
		  <input type="url" class="form-control" id="fetchurl" name="fetchurl" placeholder="输入要下载的Youtube链接">
		  <button type="submit" class="btn btn-default">添加任务</button>
		</div>
	</form>
	
	{{range .}}
	<p class="bg-info">
	Youtube链接: <a href="{{.Fetchurl}}"> {{.Fetchurl}} </a>
	{{if eq 0 .Status}}
	队列中
	{{else if eq 1 .Status}}
	下载中
	{{else if eq 2 .Status}}
	下载完成
	{{else if eq 3 .Status}}
	上传中
	{{else if eq 4 .Status}}
	上传完成
	{{end}}
	

	{{if eq 4 .Status}}
		下载链接: <a href="{{.Downloadurl}}">{{.Downloadurl}}</a>	
	{{else}}
		{{.Downloadurl}}	
	{{end}}
	</p>
	{{end}}
	

	
	</div>

  </body>
</html>