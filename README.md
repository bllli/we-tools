# we-tools 我们，工具人
golang + mysql 练手，顺便写点小工具

思路
1. 尽量少用框架，只用了sqlx CRUD；gin管理http路由和参数
2. 满足依赖倒置原则，依赖于抽象接口，不依赖于具体实现. 在main函数套娃
3. 尽量做到代码即注释

## 工程目录
```
- apps
  - {app_name}
    - models.go 模型
    - value_objects.go 值对象
    - repo.go repo层 封装对model的CRUD 使用 sqlx
    - usecase.go 业务层
    - queries.go 也是业务层，但只查询
    - api.go  接口层 解析输入参数 调用usecase/queries 输出结果 使用 gin
 
- common
  - storage 存储文件 目前用本地文件系统实现了一版，以后弄个七牛云
  - db 数据库
  - response.go 封装了gin的响应，统一返回格式
```


## apps

### memes 梗图 / 表情包
管理图片、gif图. 表情包大战大杀器

#### 一阶段 基本功能
- [ ] 存储
- [ ] 打标签
- [ ] 检索
- [ ] 收藏

#### 二阶段 社交属性
- [ ] 用户上传
- [ ] 用户打标签

持续挖坑不填中....