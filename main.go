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

type AuthedPlaylist struct {
  PlaylistOf []interface{}
}

type MyLinkedFaceAllTogether struct {
  MyAllLinkedFace []interface{}
}
type GlobalAllMine struct {
    Cdn *[]string
    Me []string
    Linked [][]string
    Playlist []string
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
    MeForMyself:=make([]string,0)
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
    for _,rmn:=range ajay.MyDetails{
      for _,rj:=range rmn.([]interface{}){
        MeForMyself=append(MeForMyself,rj.(string))
      }
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
    bn:=GlobalAllMine{Cdn:&l,Me:MeForMyself,Linked:AllInternal}
    fp := filepath.Join("templat", "/home/homeMain.rita")
    cd := filepath.Join("templat", "/home/cdn.rita")
    st := filepath.Join("templat", "/home/homeStyle.rita")
    sh := filepath.Join("templat", "/home/homeHead.rita")
    sb := filepath.Join("templat", "/home/homeBody.rita")
    sbOne := filepath.Join("templat", "/home/homeBodyPartOne.rita")
    sbThree := filepath.Join("templat", "/home/homeBodyPartThree.rita")
    pl:= filepath.Join("templat", "/home/placeholderHome.rita")
    chst := filepath.Join("templat", "/home/chatContainerStyle.rita")
    mona := filepath.Join("templat", "/home/popupListener.rita")
    t:= template.Must(template.New("home").Funcs(funcMap).ParseFiles(fp,cd,st,sh,pl,sb,sbOne,sbThree,chst,mona))
    /*if err != nil {
        panic(err)
    }*/
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("charset", "utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    t.ExecuteTemplate(w, "home", bn)
}

func AuthinticateMeYaar(a *chan []interface{},id string){
  s1:="Match (ee:Rita) where ee.email ='"
  b:=id
  s2:="' with ee optional Match (ee)-[:HAS_MANDATORY_DP]->(gg:ProfilePic) return collect([ee.password]) as all"
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

type PasswordMy struct {
  Password []interface{}
}




func serveAuth(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  e:=r.Form.Get("_em_o_phn")
  p:=r.Form.Get("_pwd")
  fmt.Println(e)
  fmt.Println(p)
  c6 := make(chan []interface{})
  go AuthinticateMeYaar(&c6,e)
  msg6 := <-c6
    c:=&PasswordMy{Password:msg6}

    for _,ght:=range c.Password{
      for _,dnt:=range ght.([]interface{}){
        if p==dnt.(string) {



          // main page after validation begining












          c1 := make(chan []interface{})
          c2 := make(chan []interface{})
          c3 := make(chan []interface{})
          go findMe(&c1,e)
          go findMyLinkedFaces(&c2,e)
          go findMyPlaylist(&c3,e)
          var ajay *Me
          var vijay *MyLinkedFaceAllTogether
          var sanjay *AuthedPlaylist
          for i := 0; i < 3; i++ {
                  select {
                    case msg1 := <-c1:
                          ajay=&Me{MyDetails:msg1}
                          fmt.Println(ajay)
                    case msg2 := <-c2:
                          vijay=&MyLinkedFaceAllTogether{MyAllLinkedFace:msg2}
                          fmt.Println(vijay)
                    case msg3 := <-c3:
                          sanjay=&AuthedPlaylist{PlaylistOf:msg3}
                          fmt.Println(sanjay)
                          fmt.Println("hi sanjay")
                          //fmt.Println(sanjay.PlaylistOf)
                    }
            }
            MeForMyself:=make([]string,0)
            MyGettingPlaylist:=make([]string,0)
            AllInternal:=make([][]string,0)
            for _,r:=range vijay.MyAllLinkedFace{
              Internal:=make([]string,0)
              for _,rl:=range r.([]interface{}){
                Internal=append(Internal,rl.(string))
              }
              AllInternal=append(AllInternal,Internal)
            }
            for _,rmn:=range ajay.MyDetails{
              for _,rj:=range rmn.([]interface{}){
                MeForMyself=append(MeForMyself,rj.(string))
              }
            }
            fmt.Println("hello")
            fmt.Println(sanjay)
            for _,gnm:=range sanjay.PlaylistOf{
                //fmt.Println("oooooooooooooooooooooooooooooooooooo")
                //fmt.Println(gnm)
                MyGettingPlaylist=append(MyGettingPlaylist,gnm.(string))
            }

            fmt.Println(AllInternal)
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
            bn:=GlobalAllMine{Cdn:&l,Me:MeForMyself,Linked:AllInternal,Playlist:MyGettingPlaylist}
            fp := filepath.Join("templat", "/home/homeMain.rita")
            cd := filepath.Join("templat", "/home/cdn.rita")
            st := filepath.Join("templat", "/home/homeStyle.rita")
            sh := filepath.Join("templat", "/home/homeHead.rita")
            sb := filepath.Join("templat", "/home/homeBody.rita")
            sbOne := filepath.Join("templat", "/home/homeBodyPartOne.rita")
            sbThree := filepath.Join("templat", "/home/homeBodyPartThree.rita")
            pl:= filepath.Join("templat", "/home/placeholderHome.rita")
            chst := filepath.Join("templat", "/home/chatContainerStyle.rita")
            mona := filepath.Join("templat", "/home/popupListener.rita")
            t:= template.Must(template.New("home").Funcs(funcMap).ParseFiles(fp,cd,st,sh,pl,sb,sbOne,sbThree,chst,mona))
            w.Header().Set("Content-Type", "text/html")
            w.Header().Set("charset", "utf-8")
            w.Header().Set("Access-Control-Allow-Origin", "*")
            t.ExecuteTemplate(w, "home", bn)















          // end of main page


        }else {
          fmt.Println("err")
        }
      }
    }

    /*var o Auth
    b, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(b))
    token := r.PostFormValue("_em_o_phn")
    fmt.Println("token")
    fmt.Println(token)
    json.Unmarshal(b, &o)
    fmt.Printf("%+v\n", o)*/
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
  s2:="' with ee optional Match (ee)-[:HAS_MANDATORY_DP]->(gg:ProfilePic) return collect([ee.name,ee.email,gg.title,ee.gender]) as all"
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
  b:=id
  s2:="' with ff optional Match (ff)-[:HAS_MANDATORY_DP]->(gg:ProfilePic) return collect([ff.name,ff.email,gg.title]) as all"
  s12 := fmt.Sprint(s1,b,s2)
  fmt.Println(s12)
  driver := bolt.NewDriver()
    conn, err := driver.OpenNeo("bolt://rita:b.PuhuqVThYfCn.fvurl1e25g7fzyCI@hobby-panhpmpgjildgbkepcdcklol.dbs.graphenedb.com:24786?tls=true")
    defer conn.Close()
    if err != nil {
      panic(err)
      fmt.Println("lollllllllll")
    }
    data, _, _, _ := conn.QueryNeoAll(s12, nil)
    //fmt.Println("hooo")
    //fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
    //fmt.Printf("FIELDS: %s\n", data[0][0].([]interface{})) // FIELDS: 1 2.2
        *a <- data[0][0].([]interface{})
}

func findMyPlaylist(a *chan []interface{},id string)  {
  s1:="MATCH (ee:Rita)-[:HAS_YOUTUBE_PLAYLIST]->(ll) where ee.email='"
  b:=id
  s2:="' return collect(ll.title) as all"
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


//youtube more videos Handler

type MoreYouTubeVideoHealper struct {
  Id string
  Title rune
}


type MoreYouTubeVideo struct {
  SkipCount int32
  Current string
}


func serveMoreVideosYoutube(w http.ResponseWriter, r *http.Request){
  var jsx MoreYouTubeVideo
  fmt.Println("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk")
  token := r.FormValue("data")
//u, _ := url.Parse(token)
  fmt.Println(token)
  json.Unmarshal([]byte(token),&jsx)

  yt := make(chan []byte)
    //seb:="crime patrol"

    i := 50
      mkiloj:=jsx.Current
    var maxResults int64
              query:= mkiloj
              maxResults=int64(i)
  fmt.Println(maxResults)
    go youTubeVideo(&yt,&query,&maxResults);

    msg:=<-yt
    fmt.Println("byte")
    fmt.Println(string(msg))
    w.Write(msg)



  fmt.Println(jsx)
}









// end youtube more videos handler




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

//youtube load more videos


func youTubeMoreVideo(a *chan []byte,c *string,e *int64){


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








//end youtube load more videos

//times of india

func timesOfIndia(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  mkiloj:=vars["newsType"]
  fmt.Println(mkiloj)
  resp, err := http.Get("https://newsapi.org/v1/articles?source="+mkiloj+"&sortBy=top&apiKey=9b9b05565f3e424fa70f46c06b6d10c8")
if err != nil {
	panic(err)
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	panic(err)
}
//fmt.Println(string(body))
w.Write(body)
}



//end times of india


func redirectHandler(w http.ResponseWriter, r *http.Request){
  newUrl:="http://localhost:5000/signup"
  http.Redirect(w, r, newUrl, 301)
  fmt.Println("biccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc")
}

type PlaylistMarsheler struct {
  Me string
  Playlist string
}

type PlaylistResponder struct {
  All []interface{} `json:all`
}

func newPlaylistCreationHandler(w http.ResponseWriter, r *http.Request)  {
  var playlistMe *PlaylistMarsheler
  token := r.FormValue("data")
  json.Unmarshal([]byte(token), &playlistMe)
  s1:="Match (ee:Rita) where ee.email ='"
  b:=playlistMe.Me
  s2:="' with ee MERGE (gg:YoutubePlaylist{title:'"
  fh:=playlistMe.Playlist
  s3:="'}) with ee,gg MERGE (ee)-[:HAS_YOUTUBE_PLAYLIST]->(gg) with ee MATCH (ee)-[:HAS_YOUTUBE_PLAYLIST]->(ll) return collect(ll.title) as all"
  s12 := fmt.Sprint(s1,b,s2,fh,s3)
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
    m:=PlaylistResponder{
      All:data[0][0].([]interface{})}
      resp, _ := json.Marshal(m)
      fmt.Println(string(resp))
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("charset", "utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write([]byte(resp))
}

func CreateSpaceForMe(a *chan []interface{},name string,email string,password string,gender string)  {
  s1:="CREATE (ee:Rita{name:'"
  cvIn:=name
  cv1:="',email:'"
  cvIn2:=email
  cv2:="',password:'"
  cvIn3:=password
  cv3:="',gender:'"
  cvIn4:=gender
  cv4:="',interest:'',dob:''}),(ff:ProfilePic{title:''}),(ee)-[:HAS_MANDATORY_DP]->(ff) return collect(ee.email)"
  s12 := fmt.Sprint(s1,cvIn,cv1,cvIn2,cv2,cvIn3,cv3,cvIn4,cv4)
  fmt.Println(s12)
  driver := bolt.NewDriver()
    conn, err := driver.OpenNeo("bolt://rita:b.PuhuqVThYfCn.fvurl1e25g7fzyCI@hobby-panhpmpgjildgbkepcdcklol.dbs.graphenedb.com:24786?tls=true")
    defer conn.Close()
    if err != nil {
      panic(err)
    }
    data, rowsMetadata, _, errx := conn.QueryNeoAll(s12, nil)
    if errx !=nil {
      var interfaceSlice []interface{} = make([]interface{}, 0)
      *a <- interfaceSlice
    }else {
      *a <- data[0][0].([]interface{})
    }
    fmt.Println("hooo")
    //fmt.Println(f)
    //fmt.Println(errx)
    fmt.Println(data)
    fmt.Println(rowsMetadata)
    //fmt.Printf("COLUMNS: %#v\n", rowsMetadata["fields"].([]interface{}))  // COLUMNS: n.foo,n.bar
    //fmt.Printf("FIELDS: %s\n", data) // FIELDS: 1 2.2
        //*a <- data[0][0].([]interface{})
}

func serveAuthAndSignUp(w http.ResponseWriter, r *http.Request)  {
    r.ParseForm()
    name:=r.Form.Get("name")
    email:=r.Form.Get("_em_o_phn")
    password:=r.Form.Get("_pwd")
    repassword:=r.Form.Get("_pwd_re")
    gender:=r.Form.Get("gender")
    fmt.Println(repassword)
    chnlSign := make(chan []interface{})
    go CreateSpaceForMe(&chnlSign,name,email,password,gender)
    msgSign := <-chnlSign
    if len(msgSign)==0{
      fmt.Println("User already exist")
    }else{
      fmt.Println("New User")
      fmt.Println(msgSign[0])
      e:=msgSign[0].(string)

      // main page after validation begining












      c1 := make(chan []interface{})
      c2 := make(chan []interface{})
      c3 := make(chan []interface{})
      go findMe(&c1,e)
      go findMyLinkedFaces(&c2,e)
      go findMyPlaylist(&c3,e)
      var ajay *Me
      var vijay *MyLinkedFaceAllTogether
      var sanjay *AuthedPlaylist
      for i := 0; i < 3; i++ {
              select {
                case msg1 := <-c1:
                      ajay=&Me{MyDetails:msg1}
                      fmt.Println(ajay)
                case msg2 := <-c2:
                      vijay=&MyLinkedFaceAllTogether{MyAllLinkedFace:msg2}
                      fmt.Println(vijay)
                case msg3 := <-c3:
                      sanjay=&AuthedPlaylist{PlaylistOf:msg3}
                      fmt.Println(sanjay)
                      fmt.Println("hi sanjay")
                      //fmt.Println(sanjay.PlaylistOf)
                }
        }
        MeForMyself:=make([]string,0)
        MyGettingPlaylist:=make([]string,0)
        AllInternal:=make([][]string,0)
        for _,r:=range vijay.MyAllLinkedFace{
          Internal:=make([]string,0)
          for _,rl:=range r.([]interface{}){
            Internal=append(Internal,rl.(string))
          }
          AllInternal=append(AllInternal,Internal)
        }
        for _,rmn:=range ajay.MyDetails{
          for _,rj:=range rmn.([]interface{}){
            MeForMyself=append(MeForMyself,rj.(string))
          }
        }
        fmt.Println("hello")
        fmt.Println(sanjay)
        for _,gnm:=range sanjay.PlaylistOf{
            //fmt.Println("oooooooooooooooooooooooooooooooooooo")
            //fmt.Println(gnm)
            MyGettingPlaylist=append(MyGettingPlaylist,gnm.(string))
        }

        fmt.Println(AllInternal)
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
        bn:=GlobalAllMine{Cdn:&l,Me:MeForMyself,Linked:AllInternal,Playlist:MyGettingPlaylist}
        fp := filepath.Join("templat", "/home/homeMain.rita")
        cd := filepath.Join("templat", "/home/cdn.rita")
        st := filepath.Join("templat", "/home/homeStyle.rita")
        sh := filepath.Join("templat", "/home/homeHead.rita")
        sb := filepath.Join("templat", "/home/homeBody.rita")
        sbOne := filepath.Join("templat", "/home/homeBodyPartOne.rita")
        sbThree := filepath.Join("templat", "/home/homeBodyPartThree.rita")
        pl:= filepath.Join("templat", "/home/placeholderHome.rita")
        chst := filepath.Join("templat", "/home/chatContainerStyle.rita")
        mona := filepath.Join("templat", "/home/popupListener.rita")
        t:= template.Must(template.New("home").Funcs(funcMap).ParseFiles(fp,cd,st,sh,pl,sb,sbOne,sbThree,chst,mona))
        w.Header().Set("Content-Type", "text/html")
        w.Header().Set("charset", "utf-8")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        t.ExecuteTemplate(w, "home", bn)















      // end of main page


    }
}

func uploadAndProcessMyNewDp(w http.ResponseWriter, r *http.Request)  {
  //var results []string
  if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
    fmt.Println(body)
    //ctx := context.Background()
    //ctx = NewContext(ctx, "cloudinary://864654217542164:fdQqxrCeKl_OJdwR84Bw9LhuUhM@hnruvsvqz")

	    //cloudinary.UploadStaticImage(ctx, "ajay", bytes.NewBuffer(body))
		//results = append(results, string(body))
    //fmt.Println(results)
		//fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}



func main() {


//timesOfIndia();
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
    r.HandleFunc("/anyMore",serveMoreVideosYoutube)
    r.HandleFunc("/redirect",redirectHandler)
    r.HandleFunc("/newPlaylistCreation",newPlaylistCreationHandler)
    r.HandleFunc("/mostPopularVideo/{newsType}",timesOfIndia)
    r.HandleFunc("/linkAuth", serveAuth).Methods("POST")
    r.HandleFunc("/signUpMePlease", serveAuthAndSignUp).Methods("POST")
    r.HandleFunc("/fileUploadItemIcon", uploadAndProcessMyNewDp).Methods("POST")
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
