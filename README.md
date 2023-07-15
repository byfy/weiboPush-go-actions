# weiboPush-go-actions
微博热搜(热、爆，新)标签推送微信，依赖WXPusher项目


## 用途

定时读取微博热搜，将热搜第一时间推送到微信，前排吃瓜。


## 教程

1. fork项目

2. 创建应用，用来推送消息

```html
https://wxpusher.zjiecode.com/admin/main
```
得到APP_TOKEN,UID

3. 修改定时器

.github/workflows/push_bao.yml 修改corn表达式

```yml
  schedule:
    # 每60分钟执行一次
    - cron:  '*/60 * * * *'
```
cron表达测试 ：https://crontab.guru

4. 配置APP_TOKEN,UID(多个uid请用,拼接),TAG(填写 爆或者 热 或者 新),COOKIE
COOKIE 获取：使用浏览器打开一次热榜页面，导出cookie即可

Settings--> secrets
![image.png](https://i.loli.net/2021/04/03/TNM2a8OSGXp6Z1F.png)

![image.png](https://i.loli.net/2021/04/03/yEPU5kdWz8RMecY.png)

5. 查看收到的消息

![image.png](https://i.loli.net/2021/04/16/jYKTorZRBfmkghD.png)

# Changelog
1. 20230715 修复原项目不能通知的bug