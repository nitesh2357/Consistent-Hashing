package main

import  (
  "github.com/julienschmidt/httprouter"
  "fmt"
  "net/http"
  "strconv"
  "encoding/json"
  "strings"
  "sort"
  "os"
)

type KeyValuePair struct{
  Key int `json:"key,omitempty"`
  Value string  `json:"value,omitempty"`
} 
var server [5]string
var n1,n2,n3,n4,n5 [] KeyValuePair
var idx1,idx2,idx3,idx4,idx5 int
type KeyPair []KeyValuePair
func (a KeyPair) Len() int           { return len(a) }
func (a KeyPair) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KeyPair) Less(i, j int) bool { return a[i].Key < a[j].Key }

func GetAllValue(rw http.ResponseWriter, req *http.Request,p httprouter.Params){
  port := strings.Split(req.Host,":")
  if(port[1]==server[0]){
      val := KeyPair(n1)
      if val ==nil {
       rw.WriteHeader(204)
		return
    }
    sort.Sort(KeyPair(n1))
    res,_ := json.Marshal(n1)
    
    
    fmt.Fprintln(rw,string(res))
  }else if(port[1]==server[1]){
      val := KeyPair(n2)
      if val ==nil {
       rw.WriteHeader(204)
		return
    }
    sort.Sort(KeyPair(n2))
    res,_:= json.Marshal(n2)
     
    fmt.Fprintln(rw,string(res))
  }else if(port[1]==server[2]){
      val := KeyPair(n3)
      if val ==nil {
       rw.WriteHeader(204)
		return
    }
    sort.Sort(KeyPair(n3))
    res,_:= json.Marshal(n3)
    fmt.Fprintln(rw,string(res))
  }else if(port[1]==server[3]){
      val := KeyPair(n4)
      if val ==nil {
       rw.WriteHeader(204)
		return
    }
    sort.Sort(KeyPair(n4))
    res,_:= json.Marshal(n4)
    fmt.Fprintln(rw,string(res))
  }else{
      val := KeyPair(n5)
      if val ==nil {
       rw.WriteHeader(204)
		return
    }
    sort.Sort(KeyPair(n5))
    res,_:= json.Marshal(n5)
    fmt.Fprintln(rw,string(res))
  }
}

func PutValue(rw http.ResponseWriter, req *http.Request,p httprouter.Params){
  port := strings.Split(req.Host,":")
  key,_ := strconv.Atoi(p.ByName("key_id"))
  if(port[1]==server[0]){

    n1 = append(n1,KeyValuePair{key,p.ByName("value")})
    idx1++
      
  }else if(port[1]==server[1]){
    n2 = append(n2,KeyValuePair{key,p.ByName("value")})
    idx2++
  }else if(port[1]==server[2]){
    n3 = append(n3,KeyValuePair{key,p.ByName("value")})
    idx3++
  }else if(port[1]==server[3]){
    n4 = append(n4,KeyValuePair{key,p.ByName("value")})
    idx4++
    }else{
    n5 = append(n5,KeyValuePair{key,p.ByName("value")})
    idx5++
  } 
}

func GetValue(rw http.ResponseWriter, req *http.Request,p httprouter.Params){ 
  out := n1
  ind := idx1
  port := strings.Split(req.Host,":")
  if(port[1]==server[1]){
    out = n2 
    ind = idx2
  }else if(port[1]==server[2]){
    out = n3
    ind = idx3
  }else if(port[1]==server[3]){
    out = n4
    ind = idx4
  }else if(port[1]==server[4]){
    out = n5
    ind = idx5
  }
  key,_ := strconv.Atoi(p.ByName("key_id"))
  for i:=0 ; i< ind ;i++{
    if(out[i].Key==key){
      res,_:= json.Marshal(out[i])
      fmt.Fprintln(rw,string(res))
    }else{
        rw.WriteHeader(204)
    }
  }
}

func main(){
  idx1 = 0
  idx2 = 0
  idx3 = 0
  mux := httprouter.New()
    mux.GET("/",GetAllValue)
    mux.GET("/:key_id",GetValue)
    mux.PUT("/keys/:key_id/:value",PutValue)
    key := strings.Split(os.Args[1],"-")
    
         a := key[0]
     var i, j int
      i, _ = strconv.Atoi(a)
      b := key[1]
      j, _ = strconv.Atoi(b)
      fmt.Println(i,j)
      w := 0
      for i<=j{
           server[w] = strconv.Itoa(i)
           i = i+1
           w = w+1
       }
   
   for a := 0; a < len(server); a++ {
      go http.ListenAndServe(":"+server[a],mux)
      fmt.Println("Starting server on Port : "+server[a])
   }
   
    select {}
}

  