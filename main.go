package main

import (
    "html/template"
    "path/filepath"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "fmt"
    "io/ioutil"
    "encoding/json"
    "os"
    "google.golang.org/api/googleapi/transport"
    "google.golang.org/api/youtube/v3"
    bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)
type Person struct {
  FirstName string
  LastName string
}
type Me struct {
  MyDetails []interface{}
}

type MyLinkedFaceAllTogether struct {
  MyAllLinkedFace []interface{}
}
type GlobalAllMine struct {
    Cdn *[]string
    Me *Me
    Linked [][]string
}
type GlobalAll struct {
    Cdn *[]string
    Me *Me
    Linked *MyLinkedFaceAllTogether
}
func (g *GlobalAll) setGlobalsAll(cd []string,me *Me,linked *MyLinkedFaceAllTogether) {
    for _,v:=range cd{
        *g.Cdn=append(*g.Cdn,v)
    }
    *g.Me=*me
    *g.Linked=*linked
}
func (g *GlobalAll) getGlobalsAll() ([]string,Me,MyLinkedFaceAllTogether) {
    return *g.Cdn,*g.Me,*g.Linked
}
type Global struct {
    cdn *[]string
}
type Auth struct {
    Email string
    Password string
}
func (g *Global) setGlobals(m ... string) {
    for _,v:=range m{
        *g.cdn=append(*g.cdn,v)
    }
}
func (g *Global) getGlobals() []string {
    return *g.cdn
}
type Music struct {
  Content []interface{}
}

type MyLinkedFaces struct {
  All []interface{}
}

type Fauzi struct {
  Linkd *[][]string
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
    fp := filepath.Join("templat", "/login/loginMain.rita")
    cd := filepath.Join("templat", "/login/cdn.rita")
    st := filepath.Join("templat", "/login/loginStyle.rita")
    lh := filepath.Join("templat", "/login/loginHead.rita")
    lb := filepath.Join("templat", "/login/loginBody.rita")
    sn:= filepath.Join("templat", "/login/sanitize.rita")
    pl:= filepath.Join("templat", "/login/placeholderLog.rita")
    t, err:= template.ParseFiles(fp,cd,st,lh,lb,sn,pl)
    if err != nil {
        panic(err)
    }
    l:=make([]string,0)
    bn:=Global{cdn:&l}
    bn.setGlobals("https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react.js","https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react-dom.js","https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js")
    mnb:=bn.getGlobals();
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("charset", "utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    t.ExecuteTemplate(w, "login", mnb)
}
func serveSignupTemplate(w http.ResponseWriter,r *http.Request)  {
  fp := filepath.Join("templat", "/signup/signupMain.rita")
  cd := filepath.Join("templat", "/signup/cdn.rita")
  st := filepath.Join("templat", "/signup/signupStyle.rita")
  sh := filepath.Join("templat", "/signup/signupHead.rita")
  sb := filepath.Join("templat", "/signup/signupBody.rita")
  sn:= filepath.Join("templat", "/signup/verify.rita")
  pl:= filepath.Join("templat", "/signup/placeholder.rita")
  t, err:= template.ParseFiles(fp,cd,st,sh,sb,sn,pl)
  if err != nil {
      panic(err)
  }
  l:=make([]string,0)
  bn:=Global{cdn:&l}
  bn.setGlobals("https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react.js","https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react-dom.js","https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js")
  mnb:=bn.getGlobals();
  w.Header().Set("Content-Type", "text/html")
  w.Header().Set("charset", "utf-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  t.ExecuteTemplate(w, "signup", mnb)
}



func serveMainTemplate(w http.ResponseWriter,r *http.Request)  {
  /*fp := filepath.Join("templat", "/home/homeMain.rita")
  cd := filepath.Join("templat", "/home/cdn.rita")
  st := filepath.Join("templat", "/home/homeStyle.rita")
  sh := filepath.Join("templat", "/home/homeHead.rita")
  sb := filepath.Join("templat", "/home/homeBody.rita")
  sbOne := filepath.Join("templat", "/home/homeBodyPartOne.rita")
  sbThree := filepath.Join("templat", "/home/homeBodyPartThree.rita")
  pl:= filepath.Join("templat", "/home/placeholderHome.rita")
  chst := filepath.Join("templat", "/home/chatContainerStyle.rita")
  t, err:= template.ParseFiles(fp,cd,st,sh,pl,sb,sbOne,sbThree,chst)
  if err != nil {
      panic(err)
  }
  l:=make([]string,0)
  bn:=Global{cdn:&l}
  bn.setGlobals("https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react.js","https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react-dom.js","https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js","https://cdnjs.cloudflare.com/ajax/libs/rxjs/5.4.0/Rx.min.js")
  mnb:=bn.getGlobals()
  w.Header().Set("Content-Type", "text/html")
  w.Header().Set("charset", "utf-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  t.ExecuteTemplate(w, "home", mnb)*/



  //new

  c1 := make(chan []interface{})
  c2 := make(chan []interface{})

  go findMe(&c1,"9831296420")
  go findMyLinkedFaces(&c2,"9831296420")

  var ajay *Me
  var vijay *MyLinkedFaceAllTogether

  for i := 0; i < 2; i++ {
          select {
            case msg1 := <-c1:
                  ajay=&Me{MyDetails:msg1}
                  fmt.Println(ajay)
            case msg2 := <-c2:
                  vijay=&MyLinkedFaceAllTogether{MyAllLinkedFace:msg2}
                  fmt.Println(vijay)
            }
    }

    AllInternal:=make([][]string,0)
    //BllInternal:=make([][]string,0)
    for _,r:=range vijay.MyAllLinkedFace{
      Internal:=make([]string,0)
      //TempHolder:=make([]string,0)
      for _,rl:=range r.([]interface{}){
        //mbcv:=strings.Join(r.([]interface),",")
        Internal=append(Internal,rl.(string))
      }
      //mbcv:=strings.Join(Internal,",")
      //TempHolder=append(TempHolder,mbcv)
      //mndir:="["+mbcv+"]"
      AllInternal=append(AllInternal,Internal)
      //fmt.Println(mbcv)
      //fmt.Println(AllInternal)
    }
    fmt.Println(AllInternal)
    //gst:=strings.Join(AllInternal,",")
    //gstr:=`[`+gst+`]`
    //fmt.Println(gst)
    //lk:=&Fauzi{Linkd:&AllInternal}
    //b, errm := json.Marshal(lk)
    /*b,errm:=JSONMarshal(lk)
	   if errm != nil {
		     panic(errm)
	      }*/
        funcMap := template.FuncMap{
            "decreased": func(i int) int {
                return i - 1
            },
        }
    l:=make([]string,0)
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react.js")
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react-dom.js")
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js")
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/rxjs/5.4.0/Rx.min.js")
    bn:=GlobalAllMine{Cdn:&l,Me:ajay,Linked:AllInternal}
    fp := filepath.Join("templat", "/home/homeMain.rita")
    cd := filepath.Join("templat", "/home/cdn.rita")
    st := filepath.Join("templat", "/home/homeStyle.rita")
    sh := filepath.Join("templat", "/home/homeHead.rita")
    sb := filepath.Join("templat", "/home/homeBody.rita")
    sbOne := filepath.Join("templat", "/home/homeBodyPartOne.rita")
    sbThree := filepath.Join("templat", "/home/homeBodyPartThree.rita")
    pl:= filepath.Join("templat", "/home/placeholderHome.rita")
    chst := filepath.Join("templat", "/home/chatContainerStyle.rita")
    t:= template.Must(template.New("home").Funcs(funcMap).ParseFiles(fp,cd,st,sh,pl,sb,sbOne,sbThree,chst))
    /*if err != nil {
        panic(err)
    }*/
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("charset", "utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    t.ExecuteTemplate(w, "home", bn)
}

func serveAuth(w http.ResponseWriter, r *http.Request) {
    var o Auth
    b, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(b))
    json.Unmarshal(b, &o)
    fmt.Printf("%+v\n", o)
}

func serveAnyTemplate(w http.ResponseWriter, r *http.Request) {

vars := mux.Vars(r)
mkiloj:=vars["query"]
fmt.Println(mkiloj)

  yt := make(chan []byte)
    //seb:="crime patrol"
    i := 20
var maxResults int64
            query:= mkiloj
            maxResults=int64(i)

  go youTubeVideo(&yt,&query,&maxResults);

  msg:=<-yt
  fmt.Println("byte")
  fmt.Println(string(msg))
  w.Write(msg)



    /*driver := bolt.NewDriver()
    	conn, err := driver.OpenNeo("bolt://rita:b.PuhuqVThYfCn.fvurl1e25g7fzyCI@hobby-panhpmpgjildgbkepcdcklol.dbs.graphenedb.com:24786?tls=true")
    	defer conn.Close()
    	if err != nil {
    		panic(err)
    	}
    	data, rowsMetadata, _, _ := conn.QueryNeoAll("MATCH (n:Video) RETURN collect(distinct n.name) as name", nil)
    	fmt.Println("hooo")
    	fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
    	fmt.Printf("FIELDS: %s\n", data[0][0].([]interface{})) // FIELDS: 1 2.2
      m:=Music{
        Content:data[0][0].([]interface{})}
        resp, _ := json.Marshal(m)
      w.Header().Set("Content-Type", "application/json")
      w.Header().Set("charset", "utf-8")
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Write([]byte(resp))*/
}

func serveMyLinedFaceTemplate(w http.ResponseWriter, r *http.Request) {
    s1:="Match (ee:Rita)-[:LINKED]-(ff:Rita) where ee.email ='"
    b:="9831296420"
    s2:="' with ff optional Match (ff)-[:HAS_MANDATORY_DP]->(gg:ProfilePic) return collect([ff.name,ff.email,gg.title]) as all"
    s12 := fmt.Sprint(s1,b,s2)
    fmt.Println(s12)
    driver := bolt.NewDriver()
    	conn, err := driver.OpenNeo("bolt://rita:b.PuhuqVThYfCn.fvurl1e25g7fzyCI@hobby-panhpmpgjildgbkepcdcklol.dbs.graphenedb.com:24786?tls=true")
    	defer conn.Close()
    	if err != nil {
    		panic(err)
    	}
    	data, rowsMetadata, _, _ := conn.QueryNeoAll(s12, nil)
    	fmt.Println("hooo")
    	fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
    	fmt.Printf("FIELDS: %s\n", data[0][0].([]interface{})) // FIELDS: 1 2.2
      m:=MyLinkedFaces{
        All:data[0][0].([]interface{})}
        resp, _ := json.Marshal(m)
        fmt.Println(string(resp))
      w.Header().Set("Content-Type", "application/json")
      w.Header().Set("charset", "utf-8")
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Write([]byte(resp))
}
func findMe(a *chan []interface{},id string){
  s1:="Match (ee:Rita) where ee.email ='"
  b:=id
  s2:="' with ee optional Match (ee)-[:HAS_MANDATORY_DP]->(gg:ProfilePic) return collect([ee.name,ee.email,gg.title]) as all"
  s12 := fmt.Sprint(s1,b,s2)
  fmt.Println(s12)
  driver := bolt.NewDriver()
    conn, err := driver.OpenNeo("bolt://rita:b.PuhuqVThYfCn.fvurl1e25g7fzyCI@hobby-panhpmpgjildgbkepcdcklol.dbs.graphenedb.com:24786?tls=true")
    defer conn.Close()
    if err != nil {
      panic(err)
    }
    data, _, _, _ := conn.QueryNeoAll(s12, nil)
    //fmt.Println("hooo")
    //fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
    //fmt.Printf("FIELDS: %s\n", data[0][0].([]interface{})) // FIELDS: 1 2.2
        *a <- data[0][0].([]interface{})
}
func findMyLinkedFaces(a *chan []interface{},id string){

  s1:="Match (ee:Rita)-[:LINKED]-(ff:Rita) where ee.email ='"
  b:="9831296420"
  s2:="' with ff optional Match (ff)-[:HAS_MANDATORY_DP]->(gg:ProfilePic) return collect([ff.name,ff.email,gg.title]) as all"
  s12 := fmt.Sprint(s1,b,s2)
  fmt.Println(s12)
  driver := bolt.NewDriver()
    conn, err := driver.OpenNeo("bolt://rita:b.PuhuqVThYfCn.fvurl1e25g7fzyCI@hobby-panhpmpgjildgbkepcdcklol.dbs.graphenedb.com:24786?tls=true")
    defer conn.Close()
    if err != nil {
      panic(err)
    }
    data, _, _, _ := conn.QueryNeoAll(s12, nil)
    //fmt.Println("hooo")
    //fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
    //fmt.Printf("FIELDS: %s\n", data[0][0].([]interface{})) // FIELDS: 1 2.2
        *a <- data[0][0].([]interface{})
}
func serveDataTemplate(w http.ResponseWriter, r *http.Request) {
  c1 := make(chan []interface{})
  c2 := make(chan []interface{})

  go findMe(&c1,"9831296420")
  go findMyLinkedFaces(&c2,"9831296420")

  var ajay *Me
  var vijay *MyLinkedFaceAllTogether

  for i := 0; i < 2; i++ {
        	select {
        		case msg1 := <-c1:
                  ajay=&Me{MyDetails:msg1}
                  fmt.Println(ajay)
        		case msg2 := <-c2:
            			vijay=&MyLinkedFaceAllTogether{MyAllLinkedFace:msg2}
                  fmt.Println(vijay)
        		}
    }

    l:=make([]string,0)
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react.js")
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react-dom.js")
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js")
    l=append(l,"https://cdnjs.cloudflare.com/ajax/libs/rxjs/5.4.0/Rx.min.js")
    bn:=GlobalAll{Cdn:&l,Me:ajay,Linked:vijay}
    fmt.Println("hello writing my")
    fmt.Println(bn.Cdn)
    fmt.Println("hello writing my")
    //bn.setGlobals("https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react.js","https://cdnjs.cloudflare.com/ajax/libs/react/15.0.2/react-dom.js","https://cdnjs.cloudflare.com/ajax/libs/babel-core/5.8.23/browser.min.js","https://cdnjs.cloudflare.com/ajax/libs/rxjs/5.4.0/Rx.min.js")
    a,b,c:=bn.getGlobalsAll()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}


//youtube search










const developerKey = "AIzaSyCDsAnh9YH2s0qKjIhsyC5K0Xou3lfavpM"


type YouTube struct {
  Id string `json:"id"`
  Title string `json:"title"`
}
type AllYoutube struct {
  All []YouTube `json:"all"`
}

func youTubeVideo(a *chan []byte,c *string,e *int64){


          client := &http.Client{
                  Transport: &transport.APIKey{Key: developerKey},
          }

          service, err := youtube.New(client)
          if err != nil {
                  panic(err)
          }

          // Make the API call to YouTube.
          call := service.Search.List("id,snippet").
                  Q(*c).
                  MaxResults(*e)
          response, err := call.Do()
          if err != nil {
                  panic(err)
          }

          // Group video, channel, and playlist results in separate lists.
          videos := make(map[string]string)
          channels := make(map[string]string)
          playlists := make(map[string]string)
          sArr:=make([]YouTube,0)
          // Iterate through each item and add it to the correct list.
          for _, item := range response.Items {
                  switch item.Id.Kind {
                  case "youtube#video":
                            c:=YouTube{Id:item.Id.VideoId,Title:item.Snippet.Title}
                            sArr=append(sArr,c)
                            //jsonString, _ := json.Marshal(datas)
                            //fmt.Println(jsonString)
                          //videos[item.Id.VideoId] = item.Snippet.Title
                  case "youtube#channel":
                          channels[item.Id.ChannelId] = item.Snippet.Title
                  case "youtube#playlist":
                          playlists[item.Id.PlaylistId] = item.Snippet.Title
                  }
          }
          //*a<-videos
          datas:=AllYoutube{All:sArr}
          jsonString, _ := json.Marshal(datas)


          *a <- jsonString
          //fmt.Println(videos)
          printIDs("Videos", videos)

}

//end youtube search

func main() {



    /*yt := make(chan interface{})
    go youTubeVideo(&yt,"crime patrol");

    msg:=<-yt
    fmt.Println(msg)*/








    port := os.Getenv("PORT")
    if port == "" {
      log.Fatal("$PORT must be set")
    }
    r := mux.NewRouter()
    r.HandleFunc("/", serveMainTemplate)
    r.HandleFunc("/login", serveTemplate)
    r.HandleFunc("/signup", serveSignupTemplate)
    r.HandleFunc("/any/{query}", serveAnyTemplate)
    r.HandleFunc("/myLinkedFaces", serveMyLinedFaceTemplate)
    r.HandleFunc("/templateData", serveDataTemplate)
    r.HandleFunc("/linkAuth", serveAuth).Methods("POST")
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
    http.Handle("/", r)
    log.Println("Listening...to all")
    http.ListenAndServe(":"+port, r)





}


func printIDs(sectionName string, matches map[string]string) {
        fmt.Println("%v:\n", sectionName)
        for id, _ := range matches {
                fmt.Println("[%v]\n", id)
        }
        fmt.Println("\n\n")
}
