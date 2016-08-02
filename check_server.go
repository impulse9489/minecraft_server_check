package main

import (
    "github.com/ammario/mcping"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
    "os"
)


func main() {
    fmt.Println("Service is running on port 8080")

    router := mux.NewRouter()
    router.HandleFunc("/", getMineCraftStatus).Methods("GET")
    http.ListenAndServe(":8080", router)


}


func getMineCraftStatus(res http.ResponseWriter, req *http.Request) {
    // Initialize and declare
    r := "SUCCESS"
    online := ""
    server, found := os.LookupEnv("MINECRAFT_SERVER")
    if !found {
                r = "MINECRAFT_SERVER required"
        fmt.Fprint(res, r)
    }
    resp, err := mcping.Ping(server)
    if err != nil  {
        fmt.Println(err)
        online = "OFFLINE"
        r = "Server "  + server + " is: " + online  + "\nError:" + err.Error()
        fmt.Fprint(res, r)

    } else {
        online = "ONLINE"

        players := ""
        for _,element := range resp.Sample {
            fmt.Println(element)
            players = players + element.Name + " "
        }
        fmt.Println("Server ", server ,"is: ", online, "\nPlayers online: ", players, "\nLatency: ", resp.Latency, "ms")

        r = "Server "  + server + " is: " + online + "\nPlayers online: " + players + "\nLatency: " + strconv.Itoa(int(resp.Latency)) + " ms"
        fmt.Fprint(res, r)
    }
}