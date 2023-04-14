### We have a business idea to create a web based service which could show train real time schedules within 12 hours and number of vacant seats in each class (sleeper, AC - for simplicity consider only these two classes) until 30 minutes before departure time. Just in case if the train is late, new departure time could be considered. 

<p> We would like to offer below services to our customers: <br>
    <ol>
      <li> check realtime train schedules and available tickets in each category </li>
      <li> train schedules and ticket availabilities shall be always shown in 12 hours window size </li>
         <ol>
             <li>  Order of display should be based on price(lower price means higher rank), number of available tickets(more available tickets means higher rank) and departure time(earlier departure time means higher rank) </li> 
            <li> In case if the train is late, new departure time should be considered </li> 
            <li> Price could be changed based on demand, available tickets and departure time, so always new price should be considered for the order of display </li>
         </ol>
      <li> last minute ticket booking facilities on some discounts i.e. we would like to offer Indian Railway last minutes ticket sell. [This requirement we would not do for now] </li>
    </ol>
</p>

Indian Railway has agreed to this business idea and given us 4 paid APIs (paid APIs means, Indian Railway would impose usage charges of these APIs on us everytime we make a request) --

```
1. POST: /register
   Response Body:
   {
      "companyName": "something"
      // other parameters are omitted for simplicity
   }

   Response Expected:
   Status Code: 200
   {
	    "companyName": "Affordmed",
	    "clientID": "59e73b59-f659-4b08-892d-c96342bcf225",
	    "clientSecret": "atwtPilfsfykSBGp"
   }

   This API is used to register a company who wish to use this paid service. Normally its a multi-steps process but 
   for simplicity consider success response means payment done and all necessary fields are supplied.
```

```
   2. GET: /trains?window-size=12
   This is a protected route and would require auth token in the header issued by Indian Railway. Auth token could be 
   acquired by 4th API
   
   Expected Response:
   Status Code: 200

   [
    {
     "trainName": "somthing exp",
     "trainNumber": "Fake321",
     "departureTime": "10:20 a.m.",
     "seatAvailable": {
        "sleeper": 5,
        "AC": 2
     },
     "price": {
        "sleeper": 432,
        "AC": 1232
     },
     "delayedBy": "10"
   }
   ]
```

```
  3. GET /trains/{trainNumber}
  This is a protected route and would require auth token in the header issued by Indian Railway. Auth token could be acquired by 4th API

   Expected Response:
   Status Code: 200
    {
     "trainName": "somthing exp",
     "trainNumber": "Fake321",
     "departureTime": "10:20 a.m.",
     "seatAvailable": {
        "sleeper": 5,
        "AC": 2
     },
     "price": {
        "sleeper": 432,
        "AC": 1232
     },
     "delayedBy": "10"
   }
```

```
   4. POST: /auth
    Request Body:
      {
      	"companyName": "Affordmed",
	"clientID": "0fc91eb2-5008-4eb5-b312-ba6dc8f65e5b",
	"clientSecret": "oJnNPGsiuzytMOJP"
      }
    
    Expected Response:
    Status Code: 200
    {
	    "token_type": "Bearer",
	    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE0NDk5MDksImNvbXBhbnlOYW1lIjoiQWZmb3JkbWVkIiwiY2xpZW50SUQiOiIwZmM5MWViMi01MDA4LTRlYjUtYjMxMi1iYTZkYzhmNjVlNWIifQ.5rrWy0NzpHhrOBzR1KHhAai0HxBOWe7gQ89IvIjoABA",
	    "expires_in": 1681449909
    }
```

> Acceptance Criteria:
> - There should be a REST interface developed for above requirements
> - Opearting cost should be minimized i.e. we should be optimizing on the charges by Indian Railway on us
> - Our web service should give its user the best experience in terms of performance
