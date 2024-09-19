# IDX-Admin

- 适应企业应用开发  
- 支持OIDC  
- 集成了Casbin的权限控制.  
- 简化CRUD的操作  
  
- 支持Code First的DB开发  
- 允许各模块按需整合，允许每个模块指定自己的数据库  
    譬如:  
    - system/controller
    - system/models  

- 集中测试
- 使用Vipper作为配置管理工具（todo...）
- 支持Cache (todo...)

使用方法
1. 通过main.go 的syncdb构建数据库
2. 通过test的testseedData初始化必要的原始数据
3. 通过test的testcasbin_test/TestMenuPermission初始化权限数据，或者将policy/中的matcher更换成不需要权限以便开发测试。

注意:
1. 使用gorm后time的null处理会导致mysql报错。所以在构建mysql连接串是需要加上parseTime=true参数  
例子:  
` root:1yrw9oLgUb@tcp(localhost:3306)/OrgDb?parseTime=true&loc=local`
2. 初始化原始数据失败
` select @@GLOBAL.sql_mode`
` set @@GLOBAL.sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';`


BaseController支持基本的CRUD和Page。
  默认路由
    - PUT {modelName}/ 保存对象，通过http body发送json对象，
    - Get {modelName}/get?id=<id>获取对象json
    - DELETE {modelName}/del?id=<id> 删除对象
    - Get {modelName}/page 分页获取数据 
page的用法  
` GET      "/api/v1/system/user/page?page=1&size=10&filters=UserName+like+%3F+and+id%3D%3F&args=%25admin%25+7a45cb54-b0ff-4ecd-95b9-074d33aaac1e&cols=UserName,ID&ordertype=asc&ordername=username"`
- @page 当前页
- @size 每页的大小
- @filters 查询条件
- @ages 查询参数
- @cols 需要查询的字段
- @ordertype 排序方式asc ,desc
- @ordername 排序字段
其中fitlers和args是通过windows.encodeURI()方法编码后的字串  
`fitlers=windows.encodeURI("UserName like ?")`   
args需要将每个参数用空格隔开  
要注意使用数据库字段，而非json后的字段名称

# Test
docker run -d -p 5432:5432 --name postgresql -e POSTGRES_PASSWORD=pass@word1 postgres