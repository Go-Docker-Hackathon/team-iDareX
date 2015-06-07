# Steamer —— Youtube 视频搬运汽船

`Steamer` /ˈstiːmə(r)/ 可以译成“汽船”，它是用来帮助你下载 `Youtube` 视频的，你可以填一个视频的链接，选择相应画质，然后 `Steamer` 将在下载完成后，把视频上传到七牛或其它你的云存储，或视频网站上去。

## 安装

### 依赖

- `Docker`
- `docker-compose`

### 下载

#### Git Conle 

```
$ git clone https://github.com/Go-Docker-Hackathon/team-iDareX.git
```

#### 修改七牛账号

```
# file: vendor/upload/qiniu/upload.go
ACCESS_KEY = ""
SECRET_KEY = ""
BUCKET_NAME := ""
```

### 构建镜像

```
$ cd team-iDareX ; docker-compose build && docker-compose up -d
```

## 访问

http://your-docker-host:18080