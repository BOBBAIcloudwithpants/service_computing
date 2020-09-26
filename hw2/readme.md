# 《服务计算》作业2: Go入门
## 目录结构
```
.
├── go.mod              // go module文件
├── readme.md
├── repeat.go           // 《learn go with tests》'迭代'的例子复现
├── repeat_test.go      // 《learn go with tests》'迭代'的例子测试
├── strings.go          // 《learn go with tests》'迭代'的练习题
├── strings_test.go     // 《learn go with tests》'迭代'的练习题测试
├── tdd.go              // tdd实践算法部分
├── tdd_test.go         // tdd实践测试部分
└── test.sh             // 脚本

```

## 使用方法
- 确保您的电脑操作系统为 linux/macOS
- 配置了golang 1.11或以上的环境，并开启go module（没有的话可以在终端执行命令: `export GO111MODULE=on`）

在 `$GOPATH/src/github.com/bobbaicloudwithpants/service_computing/hw2` 目录下进入终端，输入:
```shell script
bash test.sh
```
即可看到结果。    
    
    
祝好！     
白家栋   2020.9.26