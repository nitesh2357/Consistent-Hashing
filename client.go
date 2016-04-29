package main  
  
import (  
    "fmt"  
    "hash/crc32"  
    "sort"     
    "net/http"
   // "encoding/json" 
   // "io/ioutil"
    "os"
    "strings"
    "strconv"
)     

type HCirc []uint32  

type Node struct {  
    Id       int  
    IP       string    
}  

type KeyValuePair struct{
    Key int `json:"key,omitempty"`
    Value string `json:"value,omitempty"`
}

type ConsistentHashing struct {  
    Nodes       map[uint32]Node  
    Present   map[int]bool  
    Circle      HCirc  
    
}
var server [5]string

func (hc *ConsistentHashing) ReturnIP(node *Node) string {  
    return node.IP 
}  
  
func (hc *ConsistentHashing) Get(key string) Node {  
    hash := hc.GetHash(key)  
    i := hc.SearchNode(hash)  
    return hc.Nodes[hc.Circle[i]]  
}

func CreateNewNode(id int, ip string) *Node {  
    return &Node{  
        Id:       id,  
        IP:       ip,  
    }  
}  

func NConsistentHashing() *ConsistentHashing {  
    return &ConsistentHashing{  
        Nodes:     make(map[uint32]Node),   
        Present: make(map[int]bool),  
        Circle:      HCirc{},  
    }  
}

func (hc *ConsistentHashing) SortCircle() {  
    hc.Circle = HCirc{}  
    for k := range hc.Nodes {  
        hc.Circle = append(hc.Circle, k)  
    }  
    sort.Sort(hc.Circle)  
}  

func (hc *ConsistentHashing) GetHash(key string) uint32 {  
    return crc32.ChecksumIEEE([]byte(key))  
}  

func (hc *ConsistentHashing) SearchNode(hash uint32) int {  
    i := sort.Search(len(hc.Circle), func(i int) bool {return hc.Circle[i] >= hash })  
    if i < len(hc.Circle) {  
        if i == len(hc.Circle)-1 {  
            return 0  
        } else {  
            return i  
        }  
    } else {  
        return len(hc.Circle) - 1  
    }  
}  

func (hc HCirc) Len() int {  
    return len(hc)  
}  
  
func (hc HCirc) Less(i, j int) bool {  
    return hc[i] < hc[j]  
}  
  
func (hc HCirc) Swap(i, j int) {  
    hc[i], hc[j] = hc[j], hc[i]  
}

func (hr *ConsistentHashing) AddNewNode(node *Node) bool {   
    if _, ok := hr.Present[node.Id]; ok {  
        return false  
    }  
    str := hr.ReturnIP(node)  
    hr.Nodes[hr.GetHash(str)] = *(node)
    hr.Present[node.Id] = true  
    hr.SortCircle()  
    return true  
}  

func PutKeyValue(circ *ConsistentHashing, str string, inp string){
        ipAdd := circ.Get(str)  
        add := "http://"+ipAdd.IP+"/keys/"+str+"/"+inp
		fmt.Println(add)
        req,err := http.NewRequest("PUT",add,nil)
        client := &http.Client{}
        resp, err := client.Do(req)
        if err!=nil{
            fmt.Println("Error:",err)
        }else{
            defer resp.Body.Close()
            fmt.Println("PUT Request Done")
        }  
}  





func main() {   
    circ := NConsistentHashing() 
    key := strings.Split(os.Args[1],"-")
    //fmt.Println(key)
      a := key[0]
     var i, j int
      i, _ = strconv.Atoi(a)
      b := key[1]
      j, _ = strconv.Atoi(b)
      //fmt.Println(i,j)
      w := 0
      for i<=j{
           server[w] = strconv.Itoa(i)
           i = i+1
           w = w+1
       }
   //fmt.Println(server)
   for a := 0; a < len(server); a++ {
      circ.AddNewNode(CreateNewNode(a, "127.0.0.1:"+server[a]))
      fmt.Println("127.0.0.1:"+server[a])
         }     
    
   
	if(len(os.Args)==3){
		key1 := strings.Split(os.Args[2],",")
        for b := 0; b < len(key1); b++ {
         kv1 := strings.Split(key1[b],"->")
        
         PutKeyValue(circ,kv1[0],kv1[1])
        }
    } 
} 


	
	 