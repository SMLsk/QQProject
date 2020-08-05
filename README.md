用Go语言写仿QQ的这个GUI项目大概花了半个月左右的时间，中间学习了的比较多的反而是css等前端知识，现在从头捋一下整个过程，温故而知新吧。

在最开始写这个项目的时候，首先的问题是在于如何用Golang写一个GUI程序，为此大概纠结了一天，在网上查了一些资料，大部分都说用Go写GUI是非常吃力不讨好的事情，Go本身的用途主要在于需要高并发的服务器端，本身也不带GUI库，写出来效果比较差，写GUI目前最好的是C#，我在看了很多回答之后当时准备学C#，下了几本C#的书，第二天看了一天，仔细想了想，如果又是从头学一门语言，时间耗费不说，似乎有点在舍本逐末，在考量了成本和主要目标之后，我仍选择了Golang来做这个项目，原因是我之所以做这个项目是为了熟悉Go的使用，并且也很想尝试GUI，而不是单纯的为了做这个项目而做。

选定了Golang之后，如何写GUI程序又成为了一个问题，Golang写GUI需要借助第三方库，网上比较推荐的时walk和go-sciter，walk我没有细看，重点关注的是go-sciter，这是基于sciter技术来完成的，sciter是一个嵌入式的HTML/CSS脚本引擎，可以为桌面应用创建一个UI框架层，给桌面应用带来了很多web方面的技术，总而言之就是用HTML、CSS、TIScript（类似JavaScript）来编写桌面应用的界面，用起来还是蛮方便的，我没有了解过其他语言写GUI是怎么进行的，有什么更快更好的方式，对于现在的我而言，我觉得是一种比较熟悉且便捷的方式，唯一的问题在于，有关go-sciter的资料非常少，能找到的绝大部分都是出自博客园一位博主的文章，而sciter官方给出的文档是英文，且语焉不详，并不适合很好的学习，当然，这也是我在之后的实践中才体会出来。

开始写之前，我得学学怎么用sciter，本身的html页面倒还好，主要是go-sciter库以及script和Go交互的问题，现在想来最有用的是Go和sciter中的函数定义并互相调用的操作，其他的印象倒不深，另外，sciter也并非像浏览器一样能够对css和js的全部属性或方法都支持，所以中间走了不少弯路。

正式进入项目的开发，第一个步骤是写登录界面，我学到的第一个点是关于css和div的，通过编写登录界面，我对于div盒模型、相对/绝对定位、边框、圆形裁剪有了基本的概念，以前虽然我大致知道css和div，但是真叫我去写一个界面我是没有头绪的，做完这个之后，我有了大概的一套模式：

1. 将页面各个区域划给不同div，通过边框观察确定各div位置及所占区域
2. 将具体内容填充进div中
3. 涉及动态特效的，查找资料用script+css完成

写完一个完整的界面之后，我感觉对于css的理解增强了很多，对页面的编写脑海中有了确定的框架，当然，在这方面肯定有一套成熟的设计模式，之后得在这方面加强学习。

登录具体的业务逻辑由Go完成，大致的思路是前端界面在按钮被点击时，在点击事件中调用在Go中定义的函数，将参数传给Go，Go中客户端程序接收到之后，调用Model层的userModel，封装登录数据为一个server和client都能理解的message结构体，序列化后通过socket的接口发送出去，server端在收到之后，进行通过对数据库进行查询，获得用户数据，用户的登录口令数据确定之后，查询成功后，将查询到的用户数据和当前conn一并存入当前的UserProcess实例中，并将该结构体存入UserManager的一个map中，用用户id作为key，以此维护与客户端的连接。登录成功后还需要完成以下几件事：

1. 获取好友信息
2. 对在线的好友发送该用户登录成功的通知
3. 组装登录成功的message，将该用户的用户信息、个人信息、分组信息封装到其中，发送给该用户
4. 从数据库中读取因离线而未收到的信息，发送给用户

要实现上面的步骤，首先要靠的是数据库的支持，我使用的是MySQL，我封装了数据库的基本操作为一个DAO层，，为每一个表建立相应的model，在model中调用dao，读取数据时需要使用如下代码：

```go
func (this *Dao) ExecuteDql(sql string) (res []map[string]string, err error) {

	rows, err := this.Db.Query(sql)
	if err != nil {
		return
	}
	columns, err := rows.Columns()
	if err != nil {
		return
	}
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var i int
	for i = 0; rows.Next(); i++ {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		// fmt.Println(record)
		res = append(res, record)
	}
	if i == 0 {
		return res, errors.New("Empty Set")
	}
	return
}
```



将LoginRes传递给client之后，确定登录成功，然后client先处理其中的用户信息，并用户信息中的头像、昵称、签名等信息传递给前端用于构建页面，接着问题在于，头像的数据我该如何传输，如果直接将base64编码存到数据库中，显然占空间太大，传来传去也很不方便，因此，我的解决方案是将图片存入图床之中，数据库中只存入图床的地址，客户端接收到之后，对图片的地址进行处理，查询本地是否存有该图片，如果没有，则下载该图片到本地地址，并将用户数据中的图片地址替换为本地的地址，这里涉及到文件读写和通过http获取图片，代码如下：

```go
func (this *PhotoModel) DownPhoto(address string) (localAddress string) {
	localAddress = "../data/image/" + path.Base(address)
	_, err := os.Stat(localAddress)
	if err != nil {
		res, _ := http.Get(address)
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		ioutil.WriteFile(localAddress, data, os.ModePerm)
		return
	} else {
		return
	}

}

```

为了实现这个效果，当时在网上找了一些资料，关于文件操作，有os、ioutil、bufio等几个包，具体的我也没有去了解每个包的差异，综合之后上面的这个方案是最好的，之前用过一个网上的方案，但是下载之后的图片失真很严重，就放弃了。

关于图床的建立，我使用的是gitee来做，参照： 

> https://www.cnblogs.com/AhuntSun-blog/p/12675620.html







