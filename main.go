package main

import ( "fmt"; "./xmlrpc"; "log"; "reflect" )

func main() {
   client, _ := xmlrpc.NewClient("http://10.0.1.89/bugzilla/xmlrpc.cgi", nil)
   result := xmlrpc.Struct{}

   // login
   err := client.Call("User.login",
                      xmlrpc.Struct{"login":"test@localhost.localdomain",
                                    "password":"password"},
                      &result)
   if err != nil {
      log.Fatal(err)
      return
   }
   //fmt.Printf("User.login returned: %v\n", result)

   bug_number := "3"
   // get attachment data
   err = client.Call(
                        "Bug.attachments",
                        xmlrpc.Struct{
                           "ids":[]string{bug_number},
                            "include_fields":[]string{
                               "file_name","id",
                               "last_change_time",
                               "is_obsolete"},
                        },
                        &result,
                    )
   if err != nil {
      log.Fatal(err)
      return
   }
   fmt.Printf("Bug.attachments returned: %v\n\n", result)

   a := reflect.ValueOf(result["bugs"].(xmlrpc.Struct)[bug_number])
   for i := 0; i < a.Len(); i++ {
      attachment := a.Index(i).Interface().(xmlrpc.Struct)
      //fmt.Printf("attachment [%d] : %v\n", i, attachment)

      fmt.Printf("attachment [%d] ::\n", i)
      fmt.Printf("\tfile_name : %v\n", attachment["file_name"])
      fmt.Printf("\tid : %v\n", attachment["id"])
      fmt.Printf("\tis_obsolete : %v\n", attachment["is_obsolete"])
      fmt.Printf("\tlast_change_time : %v\n", attachment["last_change_time"])
   }
}
