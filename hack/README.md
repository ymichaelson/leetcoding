> 使用步骤

- 配置config.json文件到hack/config目录下
- cd hack/cmd
- 在hack/cmd目录下执行 $./generate

> config
```json
{
    "packageName": "",
    "funcName": "",
    "args": "",
    "result": "",
    "return": "",
    "subject": {
        "order": 0,
        "name": "",
        "topic": "",
        "examples": [
            "",
            ""
        ]
    }
}
```

> config详解
```
{
    "packageName": "search",    文件夹名字以及包名
    "funcName": "Search",       函数名
    "args": "s string",         函数传入参数
    "result": "string",         函数返回结果
    "return": "\"\"",           return的结果
    "subject": {
        "order": 1771,          题目序号
        "name": "二分查找",      题目标题
        "topic": "给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target,写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。",     题目内容
        "examples": [                                 示例...
            "输入: nums = [-1,0,3,5,9,12], target = 9  输出: 4   解释: 9 出现在 nums 中并且下标为 4",
            "输入: nums = [-1,0,3,5,9,12], target = 2  输出: -1  解释: 2 不存在 nums 中因此返回 -1",
            ...
        ]
    }
}
```