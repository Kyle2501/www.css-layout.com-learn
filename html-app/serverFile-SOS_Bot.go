// # sos-bot

package main



import (
    "os"
    "log"
		
	"text/template"
	"net/http"
	
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



func startClass() {
    var user := studentProfiles[] {}
    
    var class := classSchedule[] {}
    
    var here := studentAttendence[] {}
    
    var classTimer := ':'
    
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
    
    var homework := homeworkData[] {}
    
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
    
    
}