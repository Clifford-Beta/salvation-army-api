FORMAT: 1A
HOST: http://mine.salvation_army_api.org/v1/

# Salvation Army

This is an API for managing Salvation Army Schools.

## Schools Collection [/school]

### List All Schools [GET]

You could filter for an individual school using a school id in the url,
/school/{id}
+ Response 200 (application/json)

        [
              {
                  "id": 1,
                  "name": "Thika School",
                  "email": "thika@slvarmy.com",
                  "phone": "0912687904",
                  "postal_address": "1812912-00100",
                  "category": 2,
                  "logo": "absiujmwiuhndowd.jpg",
                  "location": "Thika Town",
                  "description": "A school for the blind",
                  "date_registered": "2017-07-26T13:21:19.883499515+03:00",
                  "time_stamp": "0001-01-01T00:00:00Z",
                  "status": 1
            }
        ]

### Create a New School [POST]

You may create a new school using this action. It takes a JSON
object containing the school details.

+ Request (application/json)

        {
              "name": "Thika School",
              "email": "thika@slvarmy.com",
              "phone": "0912687904",
              "postal_address": "1812912-00100",
              "category": 2,
              "logo": "absiujmwiuhndowd.jpg",
              "location": "Thika Town",
              "description": "A school for the blind"
        }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Thika School",
                  "email": "thika@slvarmy.com",
                  "phone": "0912687904",
                  "postal_address": "1812912-00100",
                  "category": 2,
                  "logo": "absiujmwiuhndowd.jpg",
                  "location": "Thika Town",
                  "description": "A school for the blind",
                  "date_registered": "2017-07-26T13:21:19.883499515+03:00",
                  "time_stamp": "0001-01-01T00:00:00Z",
                  "status": 1
            }



## Categories Collection [/category]

These are used to classify school, student and teacher performance.

### List All Categories [GET]

+ Response 200 (application/json)

          [
            {
                  "id": 1,
                  "name": "Overall",
                  "description": "Cumulative Performance",
                  "time_stamp": "2017-07-26T13:26:40.367497929+03:00",
                  "status": 1
            },
            {
                  "id": 2,
                  "name": "English",
                  "description": "Performance in English",
                  "time_stamp": "2017-07-26T13:26:40.367497929+03:00",
                  "status": 1
            }
        ]


### Create a New Category [POST]


+ Request (application/json)

            {
                  "name": "English",
                  "description": "Performance in English"
            }


+ Response 201 (application/json)

    + Body

            {
                  "id": 2,
                  "name": "English",
                  "description": "Performance in English",
                  "time_stamp": "2017-07-26T13:26:40.367497929+03:00",
                  "status": 1
            }


## User Collection [/user]


### List All Users [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "name": "Clifford Beta",
                      "email": "betaclifford@gmail.com",
                      "date_add": "2017-07-26T12:32:31.820413256+03:00",
                      "password": "",
                      "status": 1
                }

        ]


### Create a New User [POST]


+ Request (application/json)

           {
                  "name": "Clifford Beta",
                  "email": "betaclifford@gmail.com",
                  "password": "1234567"
            }


+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Clifford Beta",
                  "email": "betaclifford@gmail.com",
                  "date_add": "2017-07-26T12:32:31.820413256+03:00",
                  "password": "",
                  "status": 1
            }


## Staff Collection [/staff]


### List All Staff members [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "name": "Zeph Adar",
                      "email": "adarzeph@gmail.com",
                      "phone": "0712379144",
                      "role": "Head Teacher",
                      "photo": "snisnyuegbe.jpg",
                      "school": "Thika Schol for the Blind,
                      "title": "H/T",
                      "password": "",
                      "date_created": "2017-07-26T12:45:57.574926023+03:00",
                      "time_stamp": "2017-07-26T12:45:57.574926055+03:00",
                      "status": 1
                }

        ]


### Create a New Staff member [POST]


+ Request (application/json)

           {
                  "name": "Zeph Adar",
                  "email": "adarzeph@gmail.com",
                  "phone": "0712379144",
                  "role": 2,
                  "photo": "snisnyuegbe.jpg",
                  "school": 1,
                  "title": "H/T",
                  "password": "123456"
            }


+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Zeph Adar",
                  "email": "adarzeph@gmail.com",
                  "phone": "0712379144",
                  "role": 2,
                  "photo": "snisnyuegbe.jpg",
                  "school": 1,
                  "title": "H/T",
                  "password": "",
                  "date_created": "2017-07-26T12:45:57.574926023+03:00",
                  "time_stamp": "2017-07-26T12:45:57.574926055+03:00",
                  "status": 1
            }



## Staff Role Collection [/role]


### List All Roles [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 2,
                      "name": "Head Teacher",
                      "description": "The primary administrator",
                      "time_stamp": "2017-07-26T13:21:19.883413539+03:00",
                      "status": 1
                }

        ]


### Create a New Role [POST]


+ Request (application/json)

           {
                  "name": "Head Teacher",
                  "description": "The primary administrator"
            }


+ Response 201 (application/json)

    + Body

            {
                  "id": 2,
                  "name": "Head Teacher",
                  "description": "The primary administrator",
                  "time_stamp": "2017-07-26T13:21:19.883413539+03:00",
                  "status": 1
            }



## Message Collection [/message]


### List All Messages [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "title": "Test",
                      "content": "This is a test message",
                      "attachment": "",
                      "from": "admin@system.com",
                      "to": "user@system.com",
                      "status": 1,
                      "time_stamp": "2017-07-26T13:26:40.366941659+03:00",
                      "date_sent": "2017-07-26T13:26:40.366941687+03:00"
            }

        ]


### Create a New Message [POST]


+ Request (application/json)

           {
                  "title": "Test",
                  "content": "This is a test message",
                  "attachment": "",
                  "from": "admin@system.com",
                  "to": "user@system.com"
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 2,
                  "name": "Head Teacher",
                  "description": "The primary administrator",
                  "time_stamp": "2017-07-26T13:21:19.883413539+03:00",
                  "status": 1
            }



## Infrastrcture Collection [/infrastructure]


### List All Infrastructure [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "school": "Thika School For the Blind,
                      "name": "Computer",
                      "type": "Electronics",
                      "quantity": 100,
                      "description": "Desktop computers",
                      "date_created": "2017-07-26T13:26:40.367157831+03:00",
                      "time_stamp": "2017-07-26T13:26:40.367157866+03:00",
                      "status": 1
                }

        ]


### Register a New Infrastructure [POST]


+ Request (application/json)

           {
                  "school": 1,
                  "name": "Computer",
                  "type": 1,
                  "quantity": 100,
                  "description": "Desktop computers"
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "school": 1,
                  "name": "Computer",
                  "type": 1,
                  "quantity": 100,
                  "description": "Desktop computers",
                  "date_created": "2017-07-26T13:26:40.367157831+03:00",
                  "time_stamp": "2017-07-26T13:26:40.367157866+03:00",
                  "status": 1
            }


## Infrastrcture Type Collection [/infrastructure_type]


### List All Infrastructure Types [GET]

+ Response 200 (application/json)

          [
                {
                  "id": 1,
                  "name": "Electronics",
                  "description": "Electricl Appliances",
                  "time_stamp": "2017-07-26T13:26:40.367232947+03:00",
                  "status": 1
                }

        ]


### Register a New Infrastructure Type [POST]


+ Request (application/json)

           {
                  "name": "Electronics",
                  "description": "Electrical Appliances"

            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Electronics",
                  "description": "Electrical Appliances",
                  "time_stamp": "2017-07-26T13:26:40.367232947+03:00",
                  "status": 1
            }



## File Type Collection [/file_type]


### List All File Types [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "name": "Memo",
                      "description": "These are memos",
                      "store": "Memoir",
                      "status": 1,
                      "time_stamp": "2017-07-26T13:26:40.367017479+03:00"
                }

        ]


###   Create a New File Type [POST]


+ Request (application/json)

           {
                  "name": "Memo",
                  "description": "These are memos",
                  "store": "Memoir",
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Memo",
                  "description": "These are memos",
                  "store": "Memoir",
                  "status": 1,
                  "time_stamp": "2017-07-26T13:26:40.367017479+03:00"
            }



## File  Collection [/file]


### List All Files  [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "type": "Memo",
                      "name": "memo_1_2016.pdf",
                      "description": "First Memo of 2016",
                      "status": 1,
                      "date_created": "2017-07-26T13:26:40.367086424+03:00",
                      "time_stamp": "2017-07-26T13:26:40.367086454+03:00"
                }

        ]


### Upload a New File  [POST]


+ Request (application/json)

           {
                  "type": 1,
                  "name": "memo_1_2016.pdf",
                  "description": "First Memo of 2016"

            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "type": 1,
                  "name": "memo_1_2016.pdf",
                  "description": "First Memo of 2016",
                  "status": 0,
                  "date_created": "2017-07-26T13:26:40.367086424+03:00",
                  "time_stamp": "2017-07-26T13:26:40.367086454+03:00"
            }



## Extra Curricular  Collection [/extra_curricular]


### List All Extra Curricular Activities  [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "name": "Drama",
                      "description": "Acting and drama festivals",
                      "time_stamp": "2017-07-26T13:26:40.367365152+03:00",
                      "status": 1
                }

        ]


### Create a New Extra Curricular Activity  [POST]


+ Request (application/json)

          {
              "name": "Drama",
              "description": "Acting and drama festivals"
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Drama",
                  "description": "Acting and drama festivals",
                  "time_stamp": "2017-07-26T13:26:40.367365152+03:00",
                  "status": 1
            }


## Extra Curricular  Activity Collection [/extra_curricular_activity]


### List All Extra Curricular Activities Perfomance [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "school": "Thika School For the Blind,
                      "level": "Nationals",
                      "activity": "Drama",
                      "performance": "First Runners Up",
                      "date": "2017-07-26T13:26:40.367426741+03:00",
                      "time_stamp": "2017-07-26T13:26:40.367426772+03:00",
                      "status": 1
                }

        ]


### Create a New Extra Curricular Activity Perfromance [POST]


+ Request (application/json)

          {
              "school": 1,
              "level": 1,
              "activity": 1,
              "performance": "First Runners Up",
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "school": 1,
                  "level": 1,
                  "activity": 1,
                  "performance": "First Runners Up",
                  "date": "2017-07-26T13:26:40.367426741+03:00",
                  "time_stamp": "2017-07-26T13:26:40.367426772+03:00",
                  "status": 1
            }



## Extra Curricular  Activity Level [/extra_curricular_level]


### List All Extra Curricular Activities Level [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "name": "Nationals",
                      "description": "Nationwide competitions",
                      "time_stamp": "2017-07-26T13:26:40.367296638+03:00",
                      "status": 1
                }

        ]


### Create a New Extra Curricular Activity Level [POST]


+ Request (application/json)

          {
                  "name": "Nationals",
                  "description": "Nationwide competitions"
        }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "name": "Nationals",
                  "description": "Nationwide competitions",
                  "time_stamp": "2017-07-26T13:26:40.367296638+03:00",
                  "status": 1
            }


## Best Student  Collection [/best_student]


### List All Best students [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "school": "Thika School for the Blind",
                      "name": "Clifford Beta",
                      "mark": 98,
                      "class": "Form 4A",
                      "category": "Overall,
                      "age": 21,
                      "gender": "M",
                      "technique": "Nothing much",
                      "photo": "tiosjmidusyhwmeiwoe.jpg",
                      "status": 1
                }

        ]


### Create a Best Student  [POST]


+ Request (application/json)

          {
                  "school": 1,
                  "name": "Clifford Beta",
                  "mark": 98,
                  "class": "Form 4A",
                  "category": 1,
                  "age": 21,
                  "gender": "M",
                  "technique": "Nothing much",
                  "photo": "tiosjmidusyhwmeiwoe.jpg",
                  "status": 1
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "school": 1,
                  "name": "Clifford Beta",
                  "mark": 98,
                  "class": "Form 4A",
                  "category": 1,
                  "age": 21,
                  "gender": "M",
                  "technique": "Nothing much",
                  "photo": "tiosjmidusyhwmeiwoe.jpg",
                  "status": 1
            }


## Best Teacher  Collection [/best_teacher]


### List All Best teachers [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "school": "Thika School for the Blind",
                      "name": "Clifford Beta",
                      "mark": 98,
                      "class": "Form 4A",
                      "category": "Overall,
                      "year":2016,
                      "age": 21,
                      "gender": "M",
                      "technique": "Nothing much",
                      "photo": "tiosjmidusyhwmeiwoe.jpg",
                      "status": 1
                }

        ]


### Create a Best Teacher  [POST]


+ Request (application/json)

          {
                  "school": 1,
                  "name": "Clifford Beta",
                  "mark": 98,
                  "class": "Form 4A",
                  "category": 1,
                  "gender": "M",
                  "year":2016,
                  "technique": "Nothing much",
                  "photo": "tiosjmidusyhwmeiwoe.jpg",
                  "status": 1
            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "school": 1,
                  "name": "Clifford Beta",
                  "mark": 98,
                  "class": "Form 4A",
                  "category": 1,
                  "year":2016,
                  "age": 21,
                  "gender": "M",
                  "technique": "Nothing much",
                  "photo": "tiosjmidusyhwmeiwoe.jpg",
                  "status": 1
            }



## School Performance  Collection [/performance]

### List Schools' Performance [GET]

+ Response 200 (application/json)

          [
                {
                      "id": 1,
                      "mark": 87,
                      "school": "Thika School for the Blind",
                      "category": "English",
                      "year": 2016,
                      "time_stamp": "2017-07-26T13:21:19.883658494+03:00",
                      "status": 1
                }

        ]


### Register a new school performance  [POST]


+ Request (application/json)

          {
                  "mark": 87,
                  "school": 1,
                  "category": 2,
                  "year": 2016

            }

+ Response 201 (application/json)

    + Body

            {
                  "id": 1,
                  "mark": 87,
                  "school": 1,
                  "category": 2,
                  "year": 2016,
                  "time_stamp": "2017-07-26T13:21:19.883658494+03:00",
                  "status": 1
            }



## Project   Collection [/project]

### List Projects [GET]

+ Response 200 (application/json)

          [
               {
                      "id": 1,
                      "school": "Thika School for the Blind",
                      "name": "Road",
                      "description": "Roads to all hostels in the school",
                      "start": "2017-07-26T13:21:19.883578768+03:00",
                      "duration": 100,
                      "progress": 0,
                      "status": 1,
                      "time_stamp": "2017-07-26T13:21:19.883578791+03:00"
                }


        ]


### Register a new project  [POST]


+ Request (application/json)

          {
                  "school": 1,
                  "name": "Road",
                  "description": "Roads to all hostels in the school",
                  "start": "2017-07-26T13:21:19.883578768+03:00",
                  "duration": 100,
                  "progress": 0,
            }


+ Response 201 (application/json)

    + Body

             {
                      "id": 1,
                      "school": 1,
                      "name": "Road",
                      "description": "Roads to all hostels in the school",
                      "start": "2017-07-26T13:21:19.883578768+03:00",
                      "duration": 100,
                      "progress": 0,
                      "status": 1,
                      "time_stamp": "2017-07-26T13:21:19.883578791+03:00"
                }

