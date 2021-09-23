# kratos-tpl

当前模版 api 仅用于wire依赖管理
###生成模版流程-模版仅适用语 git.dev.enbrands.com/scrm/bed/scrm 内部项目
```
kratos new shop -r https://github.com/china-xs/kratos-tpl
rm go.mod
替换包完成名称 "shop/  "git.dev.enbrands.com/scrm/bed/scrm/app/shop/
make config
make wire
edit configs/config.yaml 服务端口 注册地址、数据库、缓存

```

