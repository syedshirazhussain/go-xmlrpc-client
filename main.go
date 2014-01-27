package main

import ( "fmt"; "./xmlrpc"; "log" )

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
   fmt.Printf("User.login returned: %v\n", result)

   // get attachment data
   err = client.Call(
                        "Bug.attachments",
                        xmlrpc.Struct{
                           "ids":[]string{"3"},
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
   fmt.Printf("Bug.attachments returned: %v\n", result)

   // version
   /*
   client.Call("Bugzilla.version",nil,&result)
   fmt.Printf("Bugzilla.version returned: %v\n",result)
   */
}
