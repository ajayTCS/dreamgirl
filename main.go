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
    bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)
type Person struct {
  FirstName string
  LastName string
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
  fp := filepath.Join("templat", "/home/homeMain.rita")
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
  t.ExecuteTemplate(w, "home", mnb)
}

func serveAuth(w http.ResponseWriter, r *http.Request) {
    var o Auth
    b, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(b))
    json.Unmarshal(b, &o)
    fmt.Printf("%+v\n", o)
}

func serveAnyTemplate(w http.ResponseWriter, r *http.Request) {
    driver := bolt.NewDriver()
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
      w.Write([]byte(resp))
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

func main() {
    port := os.Getenv("PORT")
    if port == "" {
      log.Fatal("$PORT must be set")
    }
    r := mux.NewRouter()
    r.HandleFunc("/", serveMainTemplate)
    r.HandleFunc("/login", serveTemplate)
    r.HandleFunc("/signup", serveSignupTemplate)
    r.HandleFunc("/any", serveAnyTemplate)
    r.HandleFunc("/myLinkedFaces", serveMyLinedFaceTemplate)
    r.HandleFunc("/linkAuth", serveAuth).Methods("POST")
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
    http.Handle("/", r)
    log.Println("Listening...to all")
    http.ListenAndServe(":"+port, r)
}
