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
	
    <link href="http://v3.bootcss.com/examples/cover/cover.css" rel="stylesheet">

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="http://cdn.bootcss.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="http://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
	<div class="site-wrapper">

      <div class="site-wrapper-inner">

        <div class="cover-container">

          <div class="masthead clearfix">

          </div>

          <div class="inner cover">
		
			<form class="form-inline" action="/addurl" method="POST">
				<div class="">
				  <input type="url" style="width:500px" class="form-control" id="fetchurl" name="fetchurl" placeholder="输入要下载的Youtube链接">
				  <button type="submit" class="btn btn-default">提交任务</button>
				</div>
			</form>
			
			<div style="height:50px;"></div>
			
			<ul class="list-group">
				{{range .}}
					  <li class="list-group-item list-group-item-info">
						Youtube链接: <a href="{{.Fetchurl}}"> {{.Fetchurl}} </a>
							
						{{if eq 0 .Status}}
						<span class="label label-default">队列中</span>
						{{else if eq 1 .Status}}
						<span class="label label-warning">正在下载</span>
						{{else if eq 2 .Status}}
						<span class="label label-info">下载完成</span>
						{{else if eq 3 .Status}}
						<span class="label label-primary">正在上传</span>
						{{else if eq 4 .Status}}
						<span class="label label-success">上传完成</span>
						{{end}}
					
						{{if eq 4 .Status}}
							下载链接: <a href="{{.Downloadurl}}">{{.Downloadurl}}</a>	
						{{else}}
							{{.Downloadurl}}	
						{{end}}
					</li>
				{{end}}
			</ul>

          </div>

          <div class="mastfoot">
            <div class="inner">
            </div>
          </div>

        </div>

      </div>

    </div>
	
  </body>
</html>