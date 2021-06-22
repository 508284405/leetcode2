package main
func main() {
	//type a struct {
	//	name string
	//	age  int
	//}
	////fmt.Println("hello word")m
	//var mapL = make(map[string]*a)
	//a2 := a{age: 15, name: "zhangsan"}
	//mapL["b"] = &a2
	//b := mapL["b"]
	//c := *b
	//c.name = "list"
	////mapL["b"] = b
	//fmt.Println(mapL)

}

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=02003390_8_hao_pg&wd=Java%20%E5%8F%98%E9%87%8F%E5%AE%9A%E4%B9%89%E5%9C%A8%E5%BE%AA%E7%8E%AF%E5%86%85%E5%92%8C%E5%BE%AA%E7%8E%AF%E5%A4%96&oq=go%25E5%2592%258CJava&rsv_pq=a6fb185200039a2c&rsv_t=6ce38uqn9mbOPcacB08kuXWUAroG4qSEx5JW3B560Y2dIqQoaMBZECIYaLJmQXnUx956lG1PvE8&rqlang=cn&rsv_enter=1&rsv_dl=tb&rsv_btype=t&inputT=8875&rsv_sug3=100&rsv_sug1=101&rsv_sug7=100&rsv_sug2=0&rsv_sug4=10626")
}
