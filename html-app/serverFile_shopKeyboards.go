// # shop keyboards

package main



import (
    "os"
    "log"
		
	"text/template"
	"net/http"
	
	"time"
	
	"cloud.google.com/go/firestore"
	
	firebase "firebase.google.com/go"
  "google.golang.org/api/option"
  
  
)


func open_browseShop() {
    
    
}


func openProduct_details() {
    
    
}


func addProduct_toCart() {
    
    
}

func create_newProduct() {
    
    
}


func editProduct_details() {



}


func open_shoppingCart() {
    
    
}




func main() {

  appName := "SOS Bot App"
  
  http.HandleFunc("/", indexHandler)
  
  http.HandleFunc("/action/about", indexHandler)
  
   
 
 http.HandleFunc("/homepage", indexHandler)
    
    
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

    
    
    
}