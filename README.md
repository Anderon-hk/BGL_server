# BGL_server

Created by Tam Chi Fung

- To query the GET "/list"
    - `curl -X GET http://ip_address_of_the_server/list`

- To use the POST "/add"
    - `curl -X POST http://ip_address_of_the_server/add -H "Content-Type: application/json" -d "{\"Key\" : \"key_value\", \"Value\" : \"value\"}"`
