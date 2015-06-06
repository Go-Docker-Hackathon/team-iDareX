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
	<style>
	li {
		list-style:none;
	}
	</style>
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
				  <input type="url" style="width:300px" class="form-control" id="fetchurl" name="fetchurl" placeholder="输入要下载的Youtube链接">
				  <a href="#" id="search" class="btn btn-default">搜索</a>
				<br />  
				<ul id="detail">
								
				</ul>
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
							<a href="{{.Downloadurl}}">下载链接</a>	
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
	
	<script src="http://cdn.bootcss.com/jquery/1.11.2/jquery.min.js"></script>
    <script src="http://cdn.bootcss.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
	
	<script>
	$('#search').click(function (e) {
		url = $("input").val();
		$.post("/search", {url: url}, function(data){
			var jsonData = eval("(" + data + ")");
			var detail = $("#detail"); 
			$.each(jsonData, function(index, objVal) { //遍历对象数组，index是数组的索引号，objVal是遍历的一个对象。  					
				$("<li>").html('<label for="f'+objVal["format"]+'"> <input id="f'+objVal["format"]+'" type="radio" value="'+objVal["format"]+'" name="format"> ' + objVal["extension"] + objVal["note"]+'</label>').appendTo(detail);
            });  
		});
	})
	</script>
	
  </body>
</html>