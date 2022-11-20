### changelog

Changelog converts the git commit history into a simple json format.  

### usage

`./changelog [path to repo]`


### output

``` json
[
  {
    "Time": "2022-11-20T10:51:21+01:00",
    "ID": "6e93c8b1c68f4594232a863c024f7e52ebc40dc5",
    "Author": "gokay",
    "Message": "update blog",
    "ChangedFiles": [
      "content/blog/2022/tty.md"
    ]
  },
  {
    "Time": "2022-11-19T18:39:28+01:00",
    "ID": "b5aad5d64330fb4aeced63adf053b44b1affdae1",
    "Author": "gokay",
    "Message": "add new post",
    "ChangedFiles": [
      "content/blog/2022/tty.md"
    ]
  },
  {
    "Time": "2022-11-19T14:49:13+01:00",
    "ID": "ebe892b732d89839960653c18ab0d90ba5230804",
    "Author": "gokay",
    "Message": "adjust pre code for small devices",
    "ChangedFiles": [
      "themes/tty/assets/css/base.css"
    ]
  }
]
```  
