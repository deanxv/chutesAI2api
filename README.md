<p align="right">
   <strong>中文</strong> 
</p>
<div align="center">

# ChutesAI2API

_觉得有点意思的话 别忘了点个 ⭐_

<a href="https://t.me/+LGKwlC_xa-E5ZDk9">
    <img src="https://telegram.org/img/website_icon.svg" width="16" height="16" style="vertical-align: middle;">
    <span style="text-decoration: none; font-size: 12px; color: #0088cc; vertical-align: middle;">Telegram 交流群</span>
</a>

<sup><i>(原`coze-discord-proxy`交流群, 此项目仍可进此群**交流** / **反馈bug**)</i></sup>
<sup><i>(群内提供公益API、AI机器人)</i></sup>

</div>

## 功能

- [x] 支持对话接口(流式/非流式)(`/chat/completions`)
   - **deepseek-r1**
   - **deepseek-v3-0324**
   - **deepseek-v3**
   - **qwq-32b**
   - **qwen2.5-72b-instruct**
   - **qwen2.5-coder-32b-instruct**
   - **gemma-3-27b-it**
   - **olympiccoder-32b**
   - **reka-flash-3**
   - **ui-tars-72b-dpo**
- [x] 支持文生图接口(`/images/generations`)【目前仅支持返回`base64`编码】
   - **juggernautxl**
   - **realistic-vision-v51**
   - **dreamshaper-xl-v2-turbo**
   - **playground-v2.5**
   - **dreamshaper-xl-1-0**
   - **omnigen-v1**
   - **animepasteldream**
   - **psychedelictrees**
   - **orphic-lora**
   - **constshaper**
   - **flux.1-dev**
   - **flex.1-alpha**
   - **flux.1-schnell**
- [x] 支持自定义请求头校验值(Authorization)
- [x] 可配置代理请求(环境变量`PROXY_URL`)

### 接口文档:

略

### 示例:

略

## 如何使用

略

## 如何集成NextChat

略

## 如何集成one-api

略

## 部署

### 基于 Docker-Compose(All In One) 进行部署

```shell
docker-compose pull && docker-compose up -d
```

#### docker-compose.yml

```docker
version: '3.4'

services:
  chutesai2api:
    image: deanxv/chutesai2api:latest
    container_name: chutesai2api
    restart: always
    ports:
      - "7011:7011"
    volumes:
      - ./data:/app/chutesai2api/data
    environment:
      - API_SECRET=123456  # [可选]接口密钥-修改此行为请求头校验的值(多个请以,分隔)
      - TZ=Asia/Shanghai
```

### 基于 Docker 进行部署

```docker
docker run --name chutesai2api -d --restart always \
-p 7011:7011 \
-v $(pwd)/data:/app/chutesai2api/data \
-e API_SECRET="123456" \
-e TZ=Asia/Shanghai \
deanxv/chutesai2api
```

其中`API_SECRET`修改为自己的。

如果上面的镜像无法拉取,可以尝试使用 GitHub 的 Docker 镜像,将上面的`deanxv/chutesai2api`替换为`ghcr.io/deanxv/chutesai2api`即可。

### 部署到第三方平台

<details>
<summary><strong>部署到 Zeabur</strong></summary>
<div>

[![Deployed on Zeabur](https://zeabur.com/deployed-on-zeabur-dark.svg)](https://zeabur.com?referralCode=deanxv&utm_source=deanxv)

> Zeabur 的服务器在国外,自动解决了网络的问题,~~同时免费的额度也足够个人使用~~

1. 首先 **fork** 一份代码。
2. 进入 [Zeabur](https://zeabur.com?referralCode=deanxv),使用github登录,进入控制台。
3. 在 Service -> Add Service,选择 Git（第一次使用需要先授权）,选择你 fork 的仓库。
4. Deploy 会自动开始,先取消。
5. 添加环境变量

   `API_SECRET:123456` [可选]接口密钥-修改此行为请求头校验的值(多个请以,分隔)(与openai-API-KEY用法一致)

保存。

6. 选择 Redeploy。

</div>


</details>

<details>
<summary><strong>部署到 Render</strong></summary>
<div>

> Render 提供免费额度,绑卡后可以进一步提升额度

Render 可以直接部署 docker 镜像,不需要 fork 仓库：[Render](https://dashboard.render.com)

</div>
</details>

## 配置

### 环境变量

1. `PORT=7011`  [可选]端口,默认为7011
2. `DEBUG=true`  [可选]DEBUG模式,可打印更多信息[true:打开、false:关闭]
3. `API_SECRET=123456`  [可选]接口密钥-修改此行为请求头(Authorization)校验的值(同API-KEY)(多个请以,分隔)
4. `REQUEST_RATE_LIMIT=60`  [可选]每分钟下的单ip请求速率限制,默认:60次/min
5. `PROXY_URL=http://127.0.0.1:10801`  [可选]代理
6. `DEFAULT_WIDTH=1024` [可选]图片宽度,默认为1024
7. `DEFAULT_HEIGHT=1024` [可选]图片高度,默认为1024
8. `DEFAULT_GUIDANCE_SCALE=4` [可选]GUIDANCE_SCALE,默认为4
9. `DEFAULT_NUM_INFERENCE_STEPS=20` [可选]DEFAULT_NUM_INFERENCE_STEPS,默认为20

## 报错排查

> `403 Forbidden`
>

官方偶发的bug,请稍后再试。

## 其他

略
