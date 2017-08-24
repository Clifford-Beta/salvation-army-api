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

+ Request  (application/json)
    
        {
                "name": "Alliance High School",
                "email": "alliance@school.ac.ke",
                "phone": "91289083",
                "postal_address": "912-00100",
                "category": 1,
                "logo": "alliance.jpg",
                "location": "Kikuyu Town",
                "description": "A top notch school"
          }
   
+   Response 200 (application/json)

    +  Body
        
            {
            "id": 18,
            "name": "Alliance High School",
            "email": "alliance@school.ac.ke",
            "phone": "91289083",
            "postal_address": "912-00100",
            "category": 1,
            "logo": "alliance.jpg",
            "location": "Kikuyu Town",
            "description": "A top notch school",
            "date_registered": "2017-08-24T10:02:09.815821178+03:00",
            "time_stamp": "0001-01-01T00:00:00Z",
            "status": 1
            }

## Categories Collection [/category]

These are used to classify school, student and teacher performance.

### List All Categories [GET]

+ Response 200 (application/json)

          {
          "data": [
            {
          "id": 1,
          "name": "Special",
          "description": "A school for those who are disabled",
          "time_stamp": "2017-08-17T16:10:34Z",
          "status": 1
          },
            {
          "id": 2,
          "name": "Normal",
          "description": "A school for those without disability",
          "time_stamp": "2017-08-17T16:10:57Z",
          "status": 1
          }
          ],
          }


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
            
            
## Tier Collection [/tier]

These are used to classify school

### List All Tiers [GET]

+ Response 200 (application/json)

          {
          "data": [
            {
          "id": 1,
          "name": "Special",
          "description": "A school for those who are disabled",
          "time_stamp": "2017-08-17T16:10:34Z",
          "status": 1
          },
            {
          "id": 2,
          "name": "Normal",
          "description": "A school for those without disability",
          "time_stamp": "2017-08-17T16:10:57Z",
          "status": 1
          }
          ],
          }


### Create a New Tier [POST]


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

          {
          "data": [
            {
          "id": 1,
          "name": "Clifford Beta ",
          "email": "betaclifford@gmail.com",
          "date_add": "2017-08-21T10:43:10Z",
          "password": "",
          "status": 1
          },
            {
          "id": 2,
          "name": " Beta Clifford",
          "email": "b.clifford@sendy.co.ke",
          "date_add": "2017-08-21T10:43:29Z",
          "password": "",
          "status": 1
          },
            {
          "id": 3,
          "name": "Willyss Beta",
          "email": "betawillys@gmail.com",
          "date_add": "2017-08-22T16:30:35Z",
          "password": "",
          "status": 1
          }
          ],
          }

### Get One User [GET] [/user/{id}]

 + Response 200
    
        {
        "id": 1,
        "name": "Clifford Beta ",
        "email": "betaclifford@gmail.com",
        "date_add": "2017-08-21T10:43:10Z",
        "password": "",
        "status": 1
        }
        
 

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

          {
          "data": [
            {
          "id": 2,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:06:07Z",
          "time_stamp": "2017-08-16T10:06:07Z",
          "status": 1
          },
            {
          "id": 3,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:08:38Z",
          "time_stamp": "2017-08-16T10:08:38Z",
          "status": 1
          },
            {
          "id": 4,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:09:46Z",
          "time_stamp": "2017-08-16T10:09:46Z",
          "status": 1
          },
            {
          "id": 5,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:10:28Z",
          "time_stamp": "2017-08-16T10:10:28Z",
          "status": 1
          },
            {
          "id": 6,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:12:12Z",
          "time_stamp": "2017-08-16T10:12:12Z",
          "status": 1
          },
            {
          "id": 7,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:13:31Z",
          "time_stamp": "2017-08-16T10:13:31Z",
          "status": 1
          },
            {
          "id": 8,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:13:55Z",
          "time_stamp": "2017-08-16T10:13:55Z",
          "status": 1
          },
            {
          "id": 9,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:15:33Z",
          "time_stamp": "2017-08-16T10:15:33Z",
          "status": 1
          },
            {
          "id": 10,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:16:07Z",
          "time_stamp": "2017-08-16T10:16:07Z",
          "status": 1
          },
            {
          "id": 11,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:16:44Z",
          "time_stamp": "2017-08-16T10:16:44Z",
          "status": 1
          },
            {
          "id": 12,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:17:31Z",
          "time_stamp": "2017-08-16T10:17:31Z",
          "status": 1
          },
            {
          "id": 13,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T10:31:46Z",
          "time_stamp": "2017-08-16T10:31:46Z",
          "status": 1
          },
            {
          "id": 14,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T17:59:23Z",
          "time_stamp": "2017-08-16T17:59:23Z",
          "status": 1
          },
            {
          "id": 15,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-16T19:30:14Z",
          "time_stamp": "2017-08-16T19:30:14Z",
          "status": 1
          },
            {
          "id": 16,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T07:36:51Z",
          "time_stamp": "2017-08-17T07:36:51Z",
          "status": 1
          },
            {
          "id": 17,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T07:42:28Z",
          "time_stamp": "2017-08-17T07:42:28Z",
          "status": 1
          },
            {
          "id": 18,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T08:01:21Z",
          "time_stamp": "2017-08-17T08:01:21Z",
          "status": 1
          },
            {
          "id": 19,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T08:02:16Z",
          "time_stamp": "2017-08-17T08:02:16Z",
          "status": 1
          },
            {
          "id": 20,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:31:23Z",
          "time_stamp": "2017-08-17T13:31:23Z",
          "status": 1
          },
            {
          "id": 21,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:32:58Z",
          "time_stamp": "2017-08-17T13:32:58Z",
          "status": 1
          },
            {
          "id": 22,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:37:35Z",
          "time_stamp": "2017-08-17T13:37:35Z",
          "status": 1
          },
            {
          "id": 23,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:38:15Z",
          "time_stamp": "2017-08-17T13:38:15Z",
          "status": 1
          },
            {
          "id": 24,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:39:43Z",
          "time_stamp": "2017-08-17T13:39:43Z",
          "status": 1
          },
            {
          "id": 25,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:40:44Z",
          "time_stamp": "2017-08-17T13:40:44Z",
          "status": 1
          },
            {
          "id": 26,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:41:38Z",
          "time_stamp": "2017-08-17T13:41:38Z",
          "status": 1
          },
            {
          "id": 27,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T13:42:27Z",
          "time_stamp": "2017-08-17T13:42:27Z",
          "status": 1
          },
            {
          "id": 28,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:01:10Z",
          "time_stamp": "2017-08-17T14:01:10Z",
          "status": 1
          },
            {
          "id": 29,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:02:17Z",
          "time_stamp": "2017-08-17T14:02:17Z",
          "status": 1
          },
            {
          "id": 30,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:05:59Z",
          "time_stamp": "2017-08-17T14:05:59Z",
          "status": 1
          },
            {
          "id": 31,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:06:59Z",
          "time_stamp": "2017-08-17T14:06:59Z",
          "status": 1
          },
            {
          "id": 32,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:07:30Z",
          "time_stamp": "2017-08-17T14:07:30Z",
          "status": 1
          },
            {
          "id": 33,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:09:41Z",
          "time_stamp": "2017-08-17T14:09:41Z",
          "status": 1
          },
            {
          "id": 34,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:12:25Z",
          "time_stamp": "2017-08-17T14:12:25Z",
          "status": 1
          },
            {
          "id": 35,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T14:15:08Z",
          "time_stamp": "2017-08-17T14:15:08Z",
          "status": 1
          },
            {
          "id": 36,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T15:25:48Z",
          "time_stamp": "2017-08-17T15:25:48Z",
          "status": 1
          },
            {
          "id": 37,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-17T15:33:37Z",
          "time_stamp": "2017-08-17T15:33:37Z",
          "status": 1
          },
            {
          "id": 38,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T09:34:00Z",
          "time_stamp": "2017-08-18T09:34:00Z",
          "status": 1
          },
            {
          "id": 39,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T11:45:12Z",
          "time_stamp": "2017-08-18T11:45:12Z",
          "status": 1
          },
            {
          "id": 40,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T12:14:30Z",
          "time_stamp": "2017-08-18T12:14:30Z",
          "status": 1
          },
            {
          "id": 41,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T13:00:09Z",
          "time_stamp": "2017-08-18T13:00:09Z",
          "status": 1
          },
            {
          "id": 42,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T13:10:07Z",
          "time_stamp": "2017-08-18T13:10:07Z",
          "status": 1
          },
            {
          "id": 43,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T13:11:59Z",
          "time_stamp": "2017-08-18T13:11:59Z",
          "status": 1
          },
            {
          "id": 44,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-18T13:13:21Z",
          "time_stamp": "2017-08-18T13:13:21Z",
          "status": 1
          },
            {
          "id": 45,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-22T16:30:35Z",
          "time_stamp": "2017-08-22T16:30:35Z",
          "status": 1
          },
            {
          "id": 46,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-22T16:50:40Z",
          "time_stamp": "2017-08-22T16:50:40Z",
          "status": 1
          },
            {
          "id": 47,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-22T16:56:30Z",
          "time_stamp": "2017-08-22T16:56:30Z",
          "status": 1
          },
            {
          "id": 48,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-22T16:57:18Z",
          "time_stamp": "2017-08-22T16:57:18Z",
          "status": 1
          },
            {
          "id": 49,
          "name": "Zeph Adar",
          "email": "adarzeph@gmail.com",
          "phone": "0712379144",
          "role": "Head Teacher",
          "photo": "snisnyuegbe.jpg",
          "school": "Juliter Beta",
          "title": "H/T",
          "date_created": "2017-08-23T06:54:12Z",
          "time_stamp": "2017-08-23T06:54:12Z",
          "status": 1
          }
          ],
          }



### Get one Staff member [GET][staff/{id}]

+ Response 200 (application/json)


    {
    "id": 2,
    "name": "Zeph Adar",
    "email": "adarzeph@gmail.com",
    "phone": "0712379144",
    "role": "Head Teacher",
    "photo": "snisnyuegbe.jpg",
    "school": "Juliter Beta",
    "title": "H/T",
    "date_created": "2017-08-16T10:06:07Z",
    "time_stamp": "2017-08-16T10:06:07Z",
    "status": 1
    }

### Create a New Staff member [POST]


+ Request (application/json)

             {
                     "name": "Kennedy Muhavi",
                     "email": "kenmuhavi@gmail.com",
                     "phone": "0712399144",
                     "role": 2,
                     "photo": "kenmuhavi.jpg",
                     "school": 1,
                     "title": "H/T",
                     "password": "123456"
               }


+ Response 201 (application/json)

    + Body

            {
            "id": 50,
            "name": "Kennedy Muhavi",
            "email": "kenmuhavi@gmail.com",
            "phone": "0712399144",
            "role": 2,
            "photo": "kenmuhavi.jpg",
            "school": 1,
            "title": "H/T",
            "password": "123456",
            "date_created": "0001-01-01T00:00:00Z",
            "time_stamp": "0001-01-01T00:00:00Z",
            "status": 1
            }



## Staff Role Collection [/role]


### List All Roles [GET]

+ Response 200 (application/json)

          {
          "data": [
            {
          "id": 1,
          "name": "Head Teacher",
          "description": "The primary administrator",
          "time_stamp": "2017-08-16T10:04:18Z",
          "status": 1
          },
            {
          "id": 2,
          "name": "Head Teacher",
          "description": "The primary administrator",
          "time_stamp": "2017-08-16T10:06:07Z",
          "status": 1
          },
            {
          "id": 3,
          "name": "Head Teacher",
          "description": "The primary administrator",
          "time_stamp": "2017-08-16T10:08:38Z",
          "status": 1
          }
          ],
          }


### Create a New Role [POST]


+ Request (application/json)

           {
                  "name": "Deputy Head Teacher",
                  "description": "Second in command"
            }


+ Response 201 (application/json)

    + Body

            {
            "id": 50,
            "name": "Deputy Head Teacher",
            "description": "Second in command",
            "time_stamp": "0001-01-01T00:00:00Z",
            "status": 1
            }



## Message Collection [/message]


### List All Messages [GET]

+ Response 200 (application/json)

          {
          "data": [
            {
          "id": 1,
          "title": "Test",
          "content": "This is a test message",
          "attachment": "",
          "from": "admin@system.com",
          "to": "user@system.com",
          "status": 1,
          "time_stamp": "2017-08-16T19:30:14Z",
          "date_sent": "2017-08-16T19:30:14Z"
          },
            {
          "id": 2,
          "title": "Test",
          "content": "This is a test message",
          "attachment": "",
          "from": "admin@system.com",
          "to": "user@system.com",
          "status": 1,
          "time_stamp": "2017-08-17T07:36:51Z",
          "date_sent": "2017-08-17T07:36:51Z"
          },
            {
          "id": 3,
          "title": "Test",
          "content": "This is a test message",
          "attachment": "",
          "from": "admin@system.com",
          "to": "user@system.com",
          "status": 1,
          "time_stamp": "2017-08-17T07:42:28Z",
          "date_sent": "2017-08-17T07:42:28Z"
          },
            {
          "id": 38,
          "title": "Test",
          "content": "This is a test message",
          "attachment": "",
          "from": "admin@system.com",
          "to": "user@system.com",
          "status": 1,
          "time_stamp": "0001-01-01T00:00:00Z",
          "date_sent": "2017-08-24T09:32:13Z"
          }
          ],
          }

### Get One Message [GET][message/{id}]

+ Response 200 (application/json)
    
    {
    "id": 1,
    "title": "Test",
    "content": "This is a test message",
    "attachment": "",
    "from": "admin@system.com",
    "to": "user@system.com",
    "status": 1,
    "time_stamp": "2017-08-16T19:30:14Z",
    "date_sent": "2017-08-16T19:30:14Z"
    }
    

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
            "id": 38,
            "title": "Test",
            "content": "This is a test message",
            "attachment": "",
            "from": "admin@system.com",
            "to": "user@system.com",
            "status": 1,
            "time_stamp": "0001-01-01T00:00:00Z",
            "date_sent": "2017-08-24T12:32:13.042121342+03:00"
            }



## Infrastructure Collection [/infrastructure]


### List All Infrastructure [GET]

+ Response 200 (application/json)

          {
          "data": [
            {
          "id": 1,
          "school": "Juliter Beta",
          "name": "Computer",
          "type": "Electronics",
          "quantity": 100,
          "description": "Desktop computers",
          "date_created": "2017-08-17T07:36:51Z"
          },
            {
          "id": 2,
          "school": "Juliter Beta",
          "name": "Computer",
          "type": "Electronics",
          "quantity": 100,
          "description": "Desktop computers",
          "date_created": "2017-08-17T07:42:28Z"
          }
          ],
          }


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

          {
          "data": [
            {
          "id": 1,
          "name": "Electronics",
          "description": "Electical appliances",
          "time_stamp": "2017-08-17T07:42:28Z",
          "status": 1
          },
            {
          "id": 34,
          "name": "Computer",
          "description": "Desktop computers",
          "time_stamp": "0001-01-01T00:00:00Z",
          "status": 1
          }
          ],
          }


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

         {
         "data": [
           {
         "id": 1,
         "name": "Memo",
         "description": "These are memos",
         "store": "Memoir",
         "status": 1,
         "time_stamp": "2017-08-17T08:01:21Z"
         },
           {
         "id": 2,
         "name": "Memo",
         "description": "These are memos",
         "store": "Memoir",
         "status": 1,
         "time_stamp": "2017-08-17T08:02:16Z"
         }
         ],
         }


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

          {
          "data": [
            {
          "id": 1,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T08:01:21Z",
          "time_stamp": "2017-08-17T08:01:21Z"
          },
            {
          "id": 2,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T08:02:16Z",
          "time_stamp": "2017-08-17T08:02:16Z"
          },
            {
          "id": 3,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:31:23Z",
          "time_stamp": "2017-08-17T13:31:23Z"
          },
            {
          "id": 4,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:32:58Z",
          "time_stamp": "2017-08-17T13:32:58Z"
          },
            {
          "id": 5,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:37:35Z",
          "time_stamp": "2017-08-17T13:37:35Z"
          },
            {
          "id": 6,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:38:15Z",
          "time_stamp": "2017-08-17T13:38:15Z"
          },
            {
          "id": 7,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:39:43Z",
          "time_stamp": "2017-08-17T13:39:43Z"
          },
            {
          "id": 8,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:40:44Z",
          "time_stamp": "2017-08-17T13:40:44Z"
          },
            {
          "id": 9,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:41:38Z",
          "time_stamp": "2017-08-17T13:41:38Z"
          },
            {
          "id": 10,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T13:42:27Z",
          "time_stamp": "2017-08-17T13:42:27Z"
          },
            {
          "id": 11,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:01:10Z",
          "time_stamp": "2017-08-17T14:01:10Z"
          },
            {
          "id": 12,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:02:17Z",
          "time_stamp": "2017-08-17T14:02:17Z"
          },
            {
          "id": 13,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:05:59Z",
          "time_stamp": "2017-08-17T14:05:59Z"
          },
            {
          "id": 14,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:06:59Z",
          "time_stamp": "2017-08-17T14:06:59Z"
          },
            {
          "id": 15,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:07:30Z",
          "time_stamp": "2017-08-17T14:07:30Z"
          },
            {
          "id": 16,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:09:41Z",
          "time_stamp": "2017-08-17T14:09:41Z"
          },
            {
          "id": 17,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:12:25Z",
          "time_stamp": "2017-08-17T14:12:25Z"
          },
            {
          "id": 18,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T14:15:08Z",
          "time_stamp": "2017-08-17T14:15:08Z"
          },
            {
          "id": 19,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T15:25:48Z",
          "time_stamp": "2017-08-17T15:25:48Z"
          },
            {
          "id": 20,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-17T15:33:37Z",
          "time_stamp": "2017-08-17T15:33:37Z"
          },
            {
          "id": 21,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-18T09:34:00Z",
          "time_stamp": "2017-08-18T09:34:00Z"
          },
            {
          "id": 22,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-18T11:45:12Z",
          "time_stamp": "2017-08-18T11:45:12Z"
          },
            {
          "id": 23,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-18T12:14:30Z",
          "time_stamp": "2017-08-18T12:14:30Z"
          },
            {
          "id": 24,
          "type": 1,
          "name": "Computer",
          "description": "Desktop computers",
          "status": 1,
          "date_created": "2017-08-18T13:00:09Z",
          "time_stamp": "2017-08-18T13:00:09Z"
          }
          ],
          }


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



## Extra Curricular  Collection [/activity]


### List All Extra Curricular Activities  [GET]

+ Response 200 (application/json)

        {
        "data": [
          {
        "id": 1,
        "name": "Drama",
        "description": "Acting and drama festivals",
        "time_stamp": "2017-08-16T08:59:40Z",
        "status": 1
        }
        ],
        }


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

          
    {
    "data": [
              {
            "id": 1,
            "name": "Drama",
            "school_name": "Juliter Beta",
            "description": "Acting and drama festivals",
            "level": "Nationals",
            "narrative": "Nationwide competitions",
            "performance": "First Runners Up",
            "date": "2017-08-16T09:12:03Z"
            },
              {
            "id": 1,
            "name": "Drama",
            "school_name": "Juliter Beta",
            "description": "Acting and drama festivals",
            "level": "Nationals",
            "narrative": "Nationwide competitions",
            "performance": "First Runners Up",
            "date": "2017-08-16T09:18:29Z"
            }
            ],
    }


### Create a New Extra Curricular Activity Perfromance [POST]


+ Request (application/json)

                     {
                        "school": 2,
                        "level": 1,
                        "activity": 1,
                        "performance": "Second Runners Up"
                    }

+ Response 201 (application/json)

    + Body

            {
            "id": 53,
            "school": 2,
            "level": 1,
            "activity": 1,
            "performance": "Second Runners Up",
            "date": "0001-01-01T00:00:00Z",
            "time_stamp": "2017-08-24T13:48:37.859754811+03:00",
            "status": 1
            }



## Extra Curricular  Activity Level [/extra_curricular_level]


### List All Extra Curricular Activities Level [GET]

+ Response 200 (application/json)

        {
        "data": [
          {
        "id": 1,
        "name": "Nationals",
        "description": "Nationwide competitions",
        "time_stamp": "2017-08-16T09:12:03Z",
        "status": 1
        },
          {
        "id": 2,
        "name": "Nationals",
        "description": "Nationwide competitions",
        "time_stamp": "2017-08-16T09:18:29Z",
        "status": 1
        }
        ],
        }

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


## Best Student  Collection [/rank_student]


### List All Best students [GET]

+ Response 200 (application/json)


        {
        "data": [
                      {
                    "id": 3,
                    "school": "Juliter Beta",
                    "name": "Mine",
                    "mark": 0,
                    "class": "",
                    "category": "Mine",
                    "age": 0,
                    "gender": "",
                    "technique": "",
                    "photo": "",
                    "year": 0,
                    "time_stamp": "0001-01-01T00:00:00Z"
                    },
                      {
                    "id": 4,
                    "school": "Juliter Beta",
                    "name": "Mine",
                    "mark": 0,
                    "class": "",
                    "category": "Mine",
                    "age": 0,
                    "gender": "",
                    "technique": "",
                    "photo": "",
                    "year": 0,
                    "time_stamp": "0001-01-01T00:00:00Z"
                    },
                      {
                    "id": 8,
                    "school": "Juliter Beta",
                    "name": "Mine",
                    "mark": 0,
                    "class": "",
                    "category": "Mine",
                    "age": 0,
                    "gender": "",
                    "technique": "",
                    "photo": "",
                    "year": 0,
                    "time_stamp": "0001-01-01T00:00:00Z"
                    },
                      {
                    "id": 12,
                    "school": "Thika School",
                    "name": "Mine",
                    "mark": 0,
                    "class": "",
                    "category": "Mine",
                    "age": 0,
                    "gender": "",
                    "technique": "",
                    "photo": "",
                    "year": 0,
                    "time_stamp": "0001-01-01T00:00:00Z"
                    },
                      {
                    "id": 20,
                    "school": "Thika School",
                    "name": "Mine",
                    "mark": 0,
                    "class": "",
                    "category": "Mine",
                    "age": 0,
                    "gender": "",
                    "technique": "",
                    "photo": "",
                    "year": 0,
                    "time_stamp": "0001-01-01T00:00:00Z"
                    }
                    ],
        }



### Create a Best Student  [POST] [/student]


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


### Retrieve best student [POST] [/best_student]
 
 
 + Request (application/json)
 
            {"from":2015,
                "to":2017
            }
    
 + Response 200 (application/json)
        
      + body
        
                {
                        "id": 3,
                        "school": "Juliter Beta",
                        "name": "Mine",
                        "mark": 0,
                        "class": "",
                        "category": "Mine",
                        "age": 0,
                        "gender": "",
                        "technique": "",
                        "photo": "",
                        "year": 0,
                        "time_stamp": "0001-01-01T00:00:00Z"
                }
        
  
## Best Teacher  Collection [/best_teacher]


### List All Best teachers [GET] [/rank_teacher]

+ Response 200 (application/json)

          {
                  "data": [
                    {
                  "id": 5,
                  "school": "Juliter Beta",
                  "name": "Mine",
                  "mark": 99,
                  "class": "",
                  "category": "Mine",
                  "gender": "",
                  "technique": "",
                  "photo": "",
                  "year": 0,
                  "time_stamp": "0001-01-01T00:00:00Z"
                  },
                    {
                  "id": 1,
                  "school": "Juliter Beta",
                  "name": "Mine",
                  "mark": 0,
                  "class": "",
                  "category": "Mine",
                  "gender": "",
                  "technique": "",
                  "photo": "",
                  "year": 0,
                  "time_stamp": "0001-01-01T00:00:00Z"
                  },
                    {
                  "id": 3,
                  "school": "Juliter Beta",
                  "name": "Mine",
                  "mark": 0,
                  "class": "",
                  "category": "Mine",
                  "gender": "",
                  "technique": "",
                  "photo": "",
                  "year": 0,
                  "time_stamp": "0001-01-01T00:00:00Z"
                  },
                    {
                  "id": 4,
                  "school": "Juliter Beta",
                  "name": "Mine",
                  "mark": 0,
                  "class": "",
                  "category": "Mine",
                  "gender": "",
                  "technique": "",
                  "photo": "",
                  "year": 0,
                  "time_stamp": "0001-01-01T00:00:00Z"
                  },
                    {
                  "id": 8,
                  "school": "Juliter Beta",
                  "name": "Mine",
                  "mark": 0,
                  "class": "",
                  "category": "Mine",
                  "gender": "",
                  "technique": "",
                  "photo": "",
                  "year": 0,
                  "time_stamp": "0001-01-01T00:00:00Z"
                  },
                    {
                  "id": 13,
                  "school": "Thika School",
                  "name": "Mine",
                  "mark": 0,
                  "class": "",
                  "category": "Mine",
                  "gender": "",
                  "technique": "",
                  "photo": "",
                  "year": 0,
                  "time_stamp": "0001-01-01T00:00:00Z"
                  }
                  ],
          }



### Create a Best Teacher  [POST] [/teacher]


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


### Retrieve  Best Teacher  [POST] [/best_teacher]



+ Response 200 (application/json)

    + Body

            {
                "id": 1,
                "school": "Juliter Beta",
                "name": "Mine",
                "mark": 99,
                "class": "",
                "category": "Mine",
                "gender": "",
                "technique": "",
                "photo": "",
                "year": 0,
                "time_stamp": "0001-01-01T00:00:00Z"
            }

## School Performance  Collection 

### Rank Schools Performance [POST]  [/ranking]

+ Response 200 (application/json)

          {
          "data": [
            {
          "id": 1,
          "mark": 97,
          "school": "Juliter Beta",
          "category": "Mine",
          "location": "",
          "tier": "Special",
          "description": "",
          "year": 2016,
          "date_registered": "2017-08-15T17:33:58Z"
          },
            {
          "id": 2,
          "mark": 85,
          "school": "Juliter Beta",
          "category": "Mine",
          "location": "",
          "tier": "Normal",
          "description": "",
          "year": 2016,
          "date_registered": "2017-08-15T17:35:08Z"
          }
          ],
          }


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


### Best School  [POST] [/best_school]


+ Request (application/json)

             {
                "from":2015,
                 "to":2017
             }

+ Response 200 (application/json)

          {
                  "id": 1,
                  "mark": 97,
                  "school": "Juliter Beta",
                  "category": "Mine",
                  "location": "",
                  "tier": "Special",
                  "description": "",
                  "year": 2016,
                  "date_registered": "2017-08-15T17:33:58Z"
          }


## Project   Collection [/project]

### List Projects [GET]

+ Response 200 (application/json)

          {
                  "data": [
                    {
                  "id": 1,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-17T15:25:48Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-17T15:25:48Z"
                  },
                    {
                  "id": 2,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-17T15:33:37Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-17T15:33:37Z"
                  },
                    {
                  "id": 3,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T09:34:00Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T09:34:00Z"
                  },
                    {
                  "id": 4,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T11:45:12Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T11:45:12Z"
                  },
                    {
                  "id": 5,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T12:14:30Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T12:14:30Z"
                  },
                    {
                  "id": 6,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T13:00:09Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T13:00:09Z"
                  },
                    {
                  "id": 7,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T13:10:07Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T13:10:07Z"
                  },
                    {
                  "id": 8,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T13:11:59Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T13:11:59Z"
                  },
                    {
                  "id": 9,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-18T13:13:21Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-18T13:13:21Z"
                  },
                    {
                  "id": 10,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-22T16:30:35Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-22T16:30:35Z"
                  },
                    {
                  "id": 11,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-22T16:50:40Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-22T16:50:40Z"
                  },
                    {
                  "id": 12,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-22T16:56:30Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-22T16:56:30Z"
                  },
                    {
                  "id": 13,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-22T16:57:18Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-22T16:57:18Z"
                  },
                    {
                  "id": 14,
                  "school": "Juliter Beta",
                  "name": "Road",
                  "description": "",
                  "start": "2017-08-23T06:54:12Z",
                  "duration": 100,
                  "progress": 0,
                  "time_stamp": "2017-08-23T06:54:12Z"
                  },
                    {
                  "id": 16,
                  "school": "Juliter Beta",
                  "name": "",
                  "description": "",
                  "start": "2017-08-24T11:27:00Z",
                  "duration": 0,
                  "progress": 0,
                  "time_stamp": "2017-08-24T11:27:00Z"
                  }
                  ],
          }


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

