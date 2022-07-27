# 建立一个简单的HTTP服务器

## http包的使用

### 第一次建立一个http服务器

```go
err := http.ListenAndServe(":8208", nil)

if err != nil {
    fmt.Println("err ->", err)
}
```

方法：
1. 使用`net/http`这个包建立对应的启动并监听一个http server
2. 判断`err`是否为`nil`来看启动服务器是否失效


### 接受请求和返回

使用`http.HandleFunc`这个方法向`http server`提供监听的方法

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "欢迎使用人力资源管理系统")
    w.Write([]byte("欢迎光临"))
})

err := http.ListenAndServe(":8208", nil)
```

**注意点**
1. 注册监听的路由需要在启动服务器之前
2. 返回http请求使用，`Fprintf`和`w.Write`都可以
3. 如果使用`w.Write`需要转换成`[]byte`格式


### 处理GET和POST请求

**注意点**

1. 在这个请求中`GET`和`POST`请求需要我们通过`r.Method`自己去区分做分支逻辑
2. 获取`query`使用`r.ParseForm`，然后通过r.FormValue来获取
3. 如果我们需要获取`POST`中的请求体`Body`的话，需要遵守以下几步
   1. `defer r.Body.Close()`在结束流程后要关闭Body获取的流
   2. 使用`ioutil.ReadAll`来读取Body，会将其转换成`[]byte`类型

```go
http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        err := r.ParseForm()

        if err != nil {
            fmt.Println("err ->", err)
        } else {
            name := r.FormValue("name")
            age, _ := strconv.Atoi(r.FormValue("age"))
            gender, _ := strconv.Atoi(r.FormValue("gender"))
            fmt.Println(name, age, gender)

            db = append(db, hr{
                Name:   name,
                Age:    age,
                Gender: gender,
            })

            fmt.Fprintf(w, "添加了"+name)
        }
    }

    if r.Method == "POST" {
        defer r.Body.Close()
        body, err := ioutil.ReadAll(r.Body)
        // ....后续处理逻辑

    }
})
```

## json格式化的方法

主要参考了官方的文档[Encoding JSON](https://pkg.go.dev/encoding/json)

刚才说到需要使用`Body`中的数据结构的话目前获得的是`[]byte`,如果我们现在需要`json`格式怎么办呢，很简单使用`encoding/json`这个包即可

### []byte 转 json

```go
type hr struct {
	Name   string `json:name`
	Age    int    `json:age`
	Gender int    `json:gender`
}

func main() {
    // 省略上述代码
	var people hr
    var db []hr

    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        fmt.Fprintln(w, err)
        panic(err)
    }

    // 这里如果转换成功会将对应的内容写到&people对应的地址内
    jsonErr := json.Unmarshal(body, &people)

    // 这里直接使用people即可
    db = append(db, hr{
        Name:   people.Name,
        Age:    people.Age,
        Gender: people.Gender,
    })
}

```

### json转[]byte

直接使用`json.Marshal`即可，这个没有什么需要注意的点

```go
http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        b, err := json.Marshal(request{
            Msg:  "ok",
            Code: 200,
            Data: db,
        })

        fmt.Println(b)

        if err != nil {
            fmt.Fprintln(w, err)
            panic(err)
        }

        w.Header().Set("Content-Type", "text/json")
        w.Write(b)
    }
})
```