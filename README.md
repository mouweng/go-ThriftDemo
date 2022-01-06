# go Thrift Demo
> This is a demo for go thrift

### 安装Thrift
Mac端安装方式
```
brew install thrift
```
安装完成后查看版本
```
$ thrift -version
```
### 编写thrift IDL
这里写了两个函数Echo、Add
```
namespace go echo

struct EchoReq {
    1: string msg;
}

struct EchoRes {
    1: string msg;
}

struct Num {
    1:required i32 id;
}

service Echo {
    EchoRes echo(1: EchoReq req);
    Num Add(1: Num num1, 2: Num num2);
}

```
### 代码生成
```
thrift -r --gen <language> <Thrift filename>

thrift -r --gen go echo.thrift 
```
### golang服务端
```go
type EchoServer struct {
}
func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
    fmt.Printf("message from client: %v\n", req.GetMsg())

    res := &echo.EchoRes{
        Msg: "success",
    }

    return res, nil
}

func (e *EchoServer) Add(ctx context.Context, num1 *echo.Num, num2 *echo.Num) (*echo.Num, error) {
    fmt.Printf("This is Add!")
    num := &echo.Num {
        ID : num1.ID + num2.ID,
    }
    return num, nil
}
```
重写接口，定义自己想要的功能。
### golang客户端
```go
// 测试Echo
req := &echo.EchoReq{Msg:"You are welcome."}
res, err := client.Echo(ctx, req)
if err != nil {
    log.Println("Echo failed:", err)
    return
}
log.Println("response:", res.Msg)
fmt.Println("well done")

// 测试Add
num1 := &echo.Num{ID:1}
num2 := &echo.Num{ID:2}
num, err := client.Add(ctx, num1, num2)
if err != nil {
    log.Println("Echo failed:", err)
    return
}
log.Println("result = ", num.ID)
fmt.Println("well done")
```
### 调用结果
server
![服务端](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220106230140.jpg)
client
![](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220106230213.jpg)
### 参考文献
- [Go thrift使用举例](https://blog.csdn.net/lanyang123456/article/details/80372977)
- [apache/thrift](https://github.com/apache/thrift)
- [从零开始基于go-thrift创建一个RPC服务](https://segmentfault.com/a/1190000019752111)
