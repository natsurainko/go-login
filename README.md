# go-login

参照原作业 [api](https://apifox.com/apidoc/shared-e228cfbb-d012-45df-943f-268ea9ad8f60?pwd=hellojh_2024) 使用 `go` 语言编写的后端服务接口  

## 部分实现方式

### 检查

每个接口都做了参数检查，不符合条件的 Body Json 或是不完整的 Query 参数都会返回

``` json
{
    "code":400,
    "data":null,
    "msg":"missing parameters（或其他信息如 EOF）"
}
```

### 删除

删除的实现方式采用的是软删除，即在数据表中增加了 `IsDeleted` 字段用于标注删除的帖子，同时保证 `post_id` 唯一  
在 `获取所有发布的帖子 /api/student/post` 做了相应的处理，不会返回删除的帖子  
在其他 删除、修改、举报 的接口中做了处理，使其对象不能是被删除的帖子  

## 对于原接口中的几个 api 进行了以下调整

### 获取所有未审批的被举报帖子 GET /api/admin/report

该接口中的请求参数不变，请求方式不变  
唯一的改动是在返回内容中的 `report_list` 键下的子项中增加了 `repoet_id`   

返回示例如下
``` json
{
    "code": 200,
    "data": {
        "report_list": [
            {
                "post_id": 4,
                "reason": "string",
                "content": "string1",
                "username": "124",
                "report_id": 2
            },
            {
                "post_id": 4,
                "reason": "string",
                "content": "string1",
                "username": "124",
                "report_id": 3
            },
            {
                "post_id": 4,
                "reason": "string",
                "content": "string1",
                "username": "124",
                "report_id": 4
            }
        ]
    },
    "msg": "success"
}
```

`repoet_id` 键的数据类型为 `long` ，用于区分每一条举报记录，使后台管理员接口能清晰的分别处理每一条举报记录

### 审核被举报的帖子 POST /api/admin/report

修改了该接口的请求参数 Body
将请求 Json 中的 `post_id` 键改为了 `report_id`
返回结果不变

请求示例如下
``` json
{
  "user_id": 1,
  "report_id": 1,
  "approval": 1
}
```
当管理员同意举报时，将直接删帖
