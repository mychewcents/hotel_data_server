# Hotel Data Server

## Installation steps
1. Go v1.19 and above should be present on your system to run the application. You can download go binary from [here](https://go.dev/doc/install).
2. Clone this repository in the GOPATH: `gopath/src/github.com/mychewcents/hotel_data_server`.
3. Inside the repository, run the command: `go install` -> This would install all the required package. Mostly this required to run the tests. Otherwise, the codebase doesn't use anything outside the base-Go binary libraries.

## Execution steps
1. While in the root of the directory, run: `go run main.go`
2. If everything is successful, you should see a message: `starting server on 8080...`

## Request/Response structure

1. Followed very simple data structure model. Stuck to the provided data model, for the sake of speed.
2. Request Format:

URL: `POST localhost:8080/hotels`

Request Body:
```JSON
{
  "hotel_ids": ["f8c9"],
  "destination_ids": [1122]
}
```

Response body:
```JSON
[
  {
    "id": "f8c9",
    "destination_id": 1122,
    "name": "Hilton Tokyo",
    "location": {
      "lat": 35.6926,
      "lng": 139.690965,
      "address": "160-0023, SHINJUKU-KU, 6-6-2 NISHI-SHINJUKU, JAPAN",
      "city": "Tokyo",
      "country": "Japan"
    },
    "description": "This sleek high-rise property is 10 minutes' walk from Shinjuku train station, 6 minutes' walk from the Tokyo Metropolitan Government Building and 3 km from Yoyogi Park. The polished rooms offer Wi-Fi and flat-screen TVs, plus minibars, sitting areas, and tea and coffeemaking facilities. Suites add living rooms, and access to a club lounge serving breakfast and cocktails. A free shuttle to Shinjuku station is offered. There's a chic Chinese restaurant, a sushi bar, and a grill restaurant with an open kitchen, as well as an English pub and a hip cocktail lounge. Other amenities include a gym, rooftop tennis courts, and a spa with an indoor pool.",
    "amenities": {
      "general": [
        "indoor pool",
        "business center",
        "wifi",
        "pool",
        "businesscenter",
        "drycleaning",
        "breakfast",
        "bar",
        "bathtub"
      ],
      "room": [
        "tv",
        "aircon",
        "minibar",
        "bathtub",
        "hair dryer"
      ]
    },
    "images": {
      "rooms": [
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i1_m.jpg",
          "description": "Suite"
        },
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i15_m.jpg",
          "description": "Double room"
        },
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i10_m.jpg",
          "description": "Suite"
        },
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i11_m.jpg",
          "description": "Suite - Living room"
        }
      ],
      "site": [
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i55_m.jpg",
          "description": "Bar"
        }
      ],
      "amenities": [
        {
          "link": "https://d2ey9sqrvkqdfs.cloudfront.net/YwAr/i57_m.jpg",
          "description": "Bar"
        }
      ]
    },
    "booking_conditions": [
      "All children are welcome. One child under 6 years stays free of charge when using existing beds. There is no capacity for extra beds in the room.",
      "Pets are not allowed.",
      "Wired internet is available in the hotel rooms and charges are applicable. WiFi is available in the hotel rooms and charges are applicable.",
      "Private parking is possible on site (reservation is not needed) and costs JPY 1500 per day.",
      "When booking more than 9 rooms, different policies and additional supplements may apply.",
      "The hotel's free shuttle is offered from Bus Stop #21 in front of Keio Department Store at Shinjuku Station. It is available every 20-minutes from 08:20-21:40. The hotel's free shuttle is offered from the hotel to Shinjuku Train Station. It is available every 20-minutes from 08:12-21:52. For more details, please contact the hotel directly. At the Executive Lounge a smart casual dress code is strongly recommended. Attires mentioned below are strongly discouraged and may not permitted: - Night attire (slippers, Yukata robe, etc.) - Gym clothes/sportswear (Tank tops, shorts, etc.) - Beachwear (flip-flops, sandals, etc.) and visible tattoos. Please note that due to renovation works, the Executive Lounge will be closed from 03 January 2019 until late April 2019. During this period, guests may experience some noise or minor disturbances. Smoking preference is subject to availability and cannot be guaranteed."
    ]
  },
  ...
]

```

## Implementation Steps
1. `internal -> datasources` -> 
   1. Because each source has their own format, we need to create multiple data sources module to help retrieve data and parse it into a common format.
   2. Because each data source needs to perform the same actions, we create an interface to maintain simplicity and consistency.

2. `intenal -> models` ->
   1. Defines the models used to translate API request/response and generic Hotel entity structure.

3. `internal -> controller` -> Acts as the business logic executioner for the query and helps in serving requests to the modules.

## Changes to the core Implementation

1. Have commented the code for using Lat/Lng from ACME provider because it's sending int and string, both type of values.

## Potential Improvements

1. Caching/Storage of the responses from the server to reduce multiple calls and latency increase.
2. Call the different providers at set intervals to refresh our storage for any new changes.
3. Can call all the sources in parallel, instead of sequential implementation, to save on latencies.
4. Code structure can be more abstracted with more controllers and functional endpoints.
5. Constants, like provider URLs, can move into a config management system instead of sitting in the code.
