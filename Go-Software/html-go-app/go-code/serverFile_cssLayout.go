// . . CSS Layout - learn - Website App

package main

// .

import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
		
  "cloud.google.com/go/firestore"
  "google.golang.org/api/iterator"
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}



// . appHandler
func appHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "CSS Layout - learn - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
  userAccountFocus := ".."
  userAccountChange := ".."
  userAccountActivity := ".."
  userAccountMetrics := ".."
  
  userAccountData := ".."
  userAccountJSON := ".."
  
  client, err := storage.NewClient(ctx)
  
    
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  pageFilePath := template.Must(template.ParseFiles("main_layout.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler


// . indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return }
// ,

  pageTitle := "CSS Layout - learn - Website App"
  pagePath := r.URL.Path
  
  pageList := []pageNav {
     { pageTitle: "one", pageLink: "one"},
     { pageTitle: "two", pageLink: "two"},
     { pageTitle: "three", pageLink: "three"},
    }
  
  if pagePath == "join/discord" {
      pageTitle = "Join Discord"
      pageList = pageList
  }
  
  if pagePath == "join/classroom" {
      pageTitle = "Join Google Classroom"
      pageList = pageList
  }
  
  if pagePath == "join/notion" {
      pageTitle = "Make Notion Profile"
      pageList = pageList
  }
  
  if pagePath == "join/payments" {
      pageTitle = "Student Payments"
      pageList = pageList
  }
    
// ,-
    
    
  if pagePath == "project_level/zero" {
      pageTitle = "Zero :: Project Level"
      pageList = pageList
  }

  if pagePath == "project_level/one" {
      pageTitle = "One :: Project Level"
      pageList = pageList
  }
    
  if pagePath == "project_level/two" {
      pageTitle = "Two :: Project Level"
      pageList = pageList
  }
    
    
  if pagePath == "project_level/three" {
      pageTitle = "Three :: Project Level"
      pageList = pageList
  }
    
    
// ,-
    
  pageData := htmlPageData {
    pageTitle: pageTitle,
    pagePath: pagePath,
      
    pageList: pageList,
  	
  }  //. .  pageData
  
  
  pageFilePath := template.Must(template.ParseFiles("main_layout.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler


//  .  html url routes as well as terminal cli logs

func main() {

    appName := "CSS Layout - learn - Website App"
    
    
    http.HandleFunc("/", indexHandler)
    
    
// ,-
    
    http.HandleFunc("/join/classroom", indexHandler)
    http.HandleFunc("/join/discord", indexHandler)
    http.HandleFunc("/join/notion", indexHandler)
    http.HandleFunc("/join/payments", indexHandler)
    
    
// ,-
    
    http.HandleFunc("/project_level/zero", indexHandler)
    http.HandleFunc("/project_level/one", indexHandler)
    http.HandleFunc("/project_level/two", indexHandler)
    http.HandleFunc("/project_level/three", indexHandler)
    
    
    http.HandleFunc("/project/level_00", indexHandler)
    http.HandleFunc("/project/level_01", indexHandler)
    http.HandleFunc("/project/level_02", indexHandler)
    http.HandleFunc("/project/level_03", indexHandler)
    
// ,-
    
    
    http.HandleFunc("/app", appHandler)




// -- -
 port := os.Getenv("PORT")
 if port == "" {
  port = "8080"
  log.Printf("Loading _webapp with default port")
 }
 
 
  log.Printf("_webapp is active and Listening on port %s", port)
        
  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)
        
// -- - 
 if err := http.ListenAndServe(":"+port, nil); err != nil {
 log.Fatal("Error Starting the HTTP Server :", err)
  return
 }


} // ,- func main app run


