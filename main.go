package main

import ( "fmt"; "./xmlrpc"; "log"; "reflect" )

func check(err error) {
   if err != nil {
      log.Fatal(err)
   }
}

func main() {
   client, _ := xmlrpc.NewClient("http://10.0.1.89/bugzilla/xmlrpc.cgi", nil)
   result := xmlrpc.Struct{}

   // login
   err := client.Call("User.login",
                      xmlrpc.Struct{"login":"test@localhost.localdomain",
                                    "password":"password"},
                      &result)
   check(err)
   //fmt.Printf("User.login returned: %v\n", result)

   // get attachment data
   bug_number := "3"

   err = client.Call(
                        "Bug.attachments",
                        xmlrpc.Struct{
                           "ids":[]string{bug_number},
                            "include_fields":[]string{
                               "file_name","id","description",
                               "last_change_time",
                               "is_obsolete"},
                        },
                        &result,
                    )
   check(err)
   fmt.Printf("Bug.attachments() returned: %v\n\n", result)

   a := reflect.ValueOf(result["bugs"].(xmlrpc.Struct)[bug_number])

   if a.Len() == 0 {
      fmt.Printf("Bug id %s has no attachments.\n", bug_number)
      return
   }

   id_list := []int64{}
   for i := 0; i < a.Len(); i++ {
      attachment := a.Index(i).Interface().(xmlrpc.Struct)
      //fmt.Printf("attachment [%d] : %v\n", i, attachment)

      id_list = append(id_list, attachment["id"].(int64))
      fmt.Printf("attachment [%d] ::\n", i)
      fmt.Printf("\tfile_name : %v\n", attachment["file_name"])
      fmt.Printf("\tdescription : %v\n", attachment["description"])
      fmt.Printf("\tid : %v\n", attachment["id"])
      fmt.Printf("\tis_obsolete : %v\n", attachment["is_obsolete"])
      fmt.Printf("\tlast_change_time : %v\n", attachment["last_change_time"])
   }

   fmt.Printf("\nid_list : %v\n", id_list)

   // fetch attachments
   err = client.Call(
                       "Bug.attachments",
                       xmlrpc.Struct{
                          "attachment_ids":id_list,
                          "include_fields":[]string{
                             "id","file_name",
                             "is_obsolete","data"},
                       },
                       &result,
                    )
   check(err)
   fmt.Printf("Bug.attachments() returned: %v\n\n", result)
}
