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
		text-align:left;
	}
	</style>
  </head>
  <body>
	<div class="site-wrapper">

      <div class="site-wrapper-inner">

        <div class="cover-container">
			<div style="text-align:left; padding-bottom:10px; padding-top:10px">
				<img width="50%" src="http://7xji7p.com1.z0.glb.clouddn.com/steamer_logo.jpg" />
				<div style="display:inline; width:48%; float:right; letter-spacing: 3px; padding-top:160px;">
				Steamer /ˈstiːmə(r)/ 可以译成“汽船”，它是用来帮助你下载 Youtube 视频的，你可以填一个视频的链接，选择相应画质，然后 Steamer 将在下载完成后，把视频上传到七牛或其它你的云存储，或视频网站上去
				</div>
			</div>
		
			<form class="form-inline" action="/addurl" method="POST">
				<div class="">
				  <input type="url" style="width:90%" class="form-control" id="fetchurl" name="fetchurl" placeholder="输入要下载的Youtube链接">
				  <a href="#" id="search" style="width:8%" class="btn btn-default">搜索</a>
				<br /> <br />
				<ul id="detail" class="list-unstyled">
								
				</ul>
				  <button type="submit" id="submit" style="display:none" class="btn btn-primary">提交任务</button>
				</div>
			</form>
			
			<div style="height:10px;"></div>
			
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

      </div>

    </div>
	
	<script src="http://cdn.bootcss.com/jquery/1.11.2/jquery.min.js"></script>
    <script src="http://cdn.bootcss.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
	
	<script>
	$('#search').click(function (e) {
		url = $("input").val();
		$.post("/search", {fetchurl: url}, function(data){
			var jsonData = eval("(" + data + ")");
			var detail = $("#detail"); 
			detail.html("");
			$.each(jsonData, function(index, objVal) { //遍历对象数组，index是数组的索引号，objVal是遍历的一个对象。  					
				$("<li>").html('<label for="f'+objVal["FormatId"]+'"> <input id="f'+objVal["FormatId"]+'" type="radio" value="'+objVal["FormatId"]+'" name="formatId"> 分辨率: ' + objVal["Resolution"] + ' 扩展: ' + objVal["Extension"] + " 备注: " + objVal["Note"]+'</label>').appendTo(detail);
            });  
			$("#submit").show();
		});
	})
	</script>
	
  </body>
</html>