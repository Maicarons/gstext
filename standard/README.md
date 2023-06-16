# 标准解析库

本库仅作为解析器的示例库,您可以根据需要定制您的解析器.

## 类型

| type  |  description  |    default     |
|:-----:|:-------------:|:--------------:|
| text  |      文本       |      text      |
|   r   |     文字换行      | &lt;empty&gt;  |
|   l   |   行末等待玩家点击    | &lt;empty&gt;  |
|   p   |    页末等待点击     | &lt;empty&gt;  |
| jump  | 跳转到剧本档中的某个标签  |       1        |


## 要求

- JSON文件

## 示例JSON

```json
{
  "scriptname": "test",
  "description": "test json",
  "author": "maicarons",
  "formatversion": "1.0.0",
  "text": [
    {
      "type": "text",
      "meta":"Maicarons",
      "remark": "test",
      "content": "Hello world!"
    },{
      "type": "p",
      "meta":"",
      "remark": "test",
      "content": ""
    },
    {
      "type": "text",
      "meta":"World",
      "remark": "test",
      "content": "Hello Maicarons!"
    }
  ]
}
```