// # sos-bot

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


type classSchedule struct {
    weekDay string
    dayTime string
    month string
    year string
    className string
    classID string
    
}

type studentProfiles struct {
    studentName string
    studentID string
    
}

type paymentInformation struct {
    studentName string
    studentID string
    statusDate string
    paymentStatus string
    paymentHistory []string
    
}


type studentAttendence struct {
    
    
}

type classEvent struct {
    
    
}

type studentHomework struct {
    studentName string
    studentID string
    
    
}

type studentFeildtrips struct {
    
    
}

type classQuestions struct {
    
    
}

type classMovies struct {
    
    
}

type studentWorkshops struct {
    
    
}


var classSchedule[] {
    
    weekDays : ['Mon','Wed','Fri'],
    timeDay : ['10:am','2:pm']
    
    
}



func addClassToList(ctx context.Context, client *firestore.Client, addClass) error {
  _ClassListData := addClass

  _, err := client.Collection("ClassList").Doc("ClassList").Set(ctx, _ClassListData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}





func listClass() {
    
    
}

func loadClass() {
    
    
}

func joinClass() {
    
    
}

func leaveClass() {
    
    
}


func addDocWithID(ctx context.Context, client *firestore.Client) error {
  var data = make(map[string]interface{})

  _, err := client.Collection("class").Doc("new-class-id").Set(ctx, data)
  
  
  if err != nil {
     log.Printf("An error has occurred: %s", err)
  }
  return err
}

func addDocDataTypes(ctx context.Context, client *firestore.Client) error {
        doc := make(map[string]interface{})
        doc["stringExample"] = "Hello world!"
        doc["booleanExample"] = true
        doc["numberExample"] = 3.14159265
        doc["dateExample"] = time.Now()
        doc["arrayExample"] = []interface{}{5, true, "hello"}
        doc["nullExample"] = nil
        doc["objectExample"] = map[string]interface{}{
                "a": 5,
                "b": true,
        }

  _, err := client.Collection("data").Doc("one").Set(ctx, doc)
        
  if err != nil {
   // Handle any errors
    log.Printf("An error has occurred: %s", err)
 }
  return err
}

func addDocAsEntity(ctx context.Context, client *firestore.Client) error {
  studentData := studentData{
    Name:    "Phoenix",
    Status: "Active",
}

  _, err := client.Collection("studentData").Doc("studentData").Set(ctx, studentData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}



func addHomeworkToFile(ctx context.Context, client *firestore.Client, saveHomework) error {
  homeworkData := saveHomework

  _, err := client.Collection("HomeworkFile").Doc("studentData").Set(ctx, homeworkData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}


func addStudentAttendence(ctx context.Context, client *firestore.Client, attendenceData) error {
  studentData := attendenceData

  _, err := client.Collection("AttendenceData").Doc("studentAttendence").Set(ctx, studentData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}





func startClass() {
    var user := studentProfiles[] {}
    
    var class := classSchedule[] {}
    
    var here := studentAttendence[] {}
    
    var classTimer := ':'
    
    addStudentAttendece(attendenceData)
    
}

func meetTeacher() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}
    

}

func doEvent() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}
    
    var event := eventData[] {}
    
}


func homework() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}
    
// receives homework files from Discord User Interface
    
    var saveHomework := homeworkData[] {}
    
    action := getActionType() {
        
    }
    if action == turnIn {
        
        addHomeworkToFile(saveHomework)
        
        
    }
    
    
}

func askQuestions() {
    var user := studentProfiles[] {}
    var class := classSchedule[] {}
    
    var question := classQuestions[] {}
    
}

func meetStudents() {
    
    var class := classSchedule[] {}
       
    var here := studentAttendence[] {}
     
    
}

func studentFeildtrip() {
    
    var class := classSchedule[] {}
    
    var here := studentAttendence[] {}
    
    var trip := studentFeildtrips[] {}
    
}




func addMovieToList(ctx context.Context, client *firestore.Client, addMovie) error {
  _MovieListData := addMovie

  _, err := client.Collection("MovieList").Doc("MovieList").Set(ctx, _MovieListData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
}



func watchMovie() {
    var class := classSchedule[] {}
    var here := studentAttendence[] {}
    
}

func studentWorkshops() {
    
    
}


func endClass() {
    var user := studentProfiles[] {}
    
    var class := classSchedule[] {}
    
    var here := studentAttendence[] {}
    
}

// . indexHandler
func indexHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "SOS Bot App"
  pagePath := r.URL.Path
    
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
  
}  //  .  indexHandler



func main() {

  appName := "SOS Bot App"
  
  http.HandleFunc("/", indexHandler)
  
  http.HandleFunc("/action/about", aboutBot)
  
  http.HandleFunc("/action/listclass", listClass)
  
  http.HandleFunc("/action/loadclass", loadClass)
  http.HandleFunc("/action/joinclass", joinClass)
  http.HandleFunc("/action/leaveclass", leaveClass)
  

    
 http.HandleFunc("/action/startclass", startClass)
 http.HandleFunc("/action/endclass", endClass)
 
 http.HandleFunc("/action/veiwprojects", veiwProjects)
 

 http.HandleFunc("/action/doevent", doEvents)
 http.HandleFunc("/action/homework", homework)
 
 http.HandleFunc("/action/watchmovie", watchMovie)
 http.HandleFunc("/action/studentfeildtrip", studentFeildtrip)

 
  http.HandleFunc("/action/meetteacher", meetTeacher)
  http.HandleFunc("/action/meetstudents", meetStudents)
  http.HandleFunc("/action/askquestions", askQuestions)
 
 

 
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