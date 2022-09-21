# ob-example
Oceanbase基于giipod建立了快速在线体验平台, 点击下面按钮一键体验(建议使用chrome浏览器):
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/akaError/ob-example)

进入后会自动部署一个OceanBase本地实例，请等待Oceanbase boot Success, 可在右边终端中操作
![](https://cn-hangzhou.oss-cdn.aliyun-inc.com/git/force/uploads/comment/292665/34479482796045550/image.png)

仓库中提供了不同语言和工具连接Oceanbase的示例，按下面步骤进行操作
等待
```
//1. open 
cd xxxx
//2. prepare relative env
sh env.sh
//3. run to query
sh run.sh
```
这里以python3为例
```
cd python3-pymysql
sh env.sh
sh run.sh
```

关于更多Oceanbase的细节请参考 [OceanBase](https://open.oceanbase.com).
