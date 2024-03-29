package client

var requestBody = []byte(`{
  "id": "0123456789ABCDEF0123456789ABCDEF",
  "at": 2,
  "tmax": 120,
  "imp": [
    {
      "id": "1",
      "pmp": {
        "private_auction":1,
        "deals": [{
            "id":"1452f.eadb4.7aaa",
            "bidfloor":5.3,
            "at":1,
            "wseats":[],
            "ext": {
                "priority":1,
                "wadvs":[]
              }
          }
        ]
      },
      "video": {
        "mimes": [
          "video/x-flv",
          "video/mp4",
          "application/x-shockwave-flash",
          "application/javascript"
        ],

        "api":             [1,2],
        "battr":           [13,14],
        "boxingallowed":   true,
        "delivery":        [2],
        "h":               480,
        "linearity":       1,
        "maxbitrate":      1500,
        "maxduration":     30,
        "minbitrate":      300,
        "minduration":     5,
        "playbackmethod":  [1],
        "pos":             1,
        "protocol":        [2,3],
        "sequence":        1,
        "startdelay":      0,
        "w":               640
      }
    },
    {
      "id": "2",
      "pmp": {
        "private_auction":1,
        "deals": [
          {
            "id":"1452f.eadb4.7aaa",
            "bidfloor":3.5,
            "at":1,
            "wseats":[],
            "ext": {
              "priority":1,
              "wadvs":[]
            }
          },
          {
            "id":"1452f.eadb4.f9bc",
            "bidfloor":2.5,
            "at":1,
            "wseats":["45","165","33"],
            "ext": {
              "priority":2,
              "wadvs":[]
            }
          }
        ]
      },
      "video": {
        "mimes": [
          "video/x-flv",
          "video/mp4",
          "application/x-shockwave-flash",
          "application/javascript"
        ],

        "api":             [1,2],
        "battr":           [13,14],
        "boxingallowed":   true,
        "delivery":        [2],
        "h":               480,
        "linearity":       1,
        "maxbitrate":      1500,
        "maxduration":     60,
        "minbitrate":      300,
        "minduration":     30,
        "playbackmethod":  [1],
        "pos":             1,
        "protocol":        [2,3],
        "sequence":        2,
        "startdelay":      300,
        "w":               640
      }
    },
    {
      "id": "3",
      "bidfloor": 2.00
      "video": {
        "mimes": [
          "video/x-flv",
          "video/mp4",
          "application/x-shockwave-flash",
          "application/javascript"
        ],

        "api":             [1,2],
        "battr":           [13,14],
        "boxingallowed":   true,
        "delivery":        [2],
        "h":               480,
        "linearity":       1,
        "maxbitrate":      1500,
        "maxduration":     60,
        "minbitrate":      300,
        "minduration":     30,
        "playbackmethod":  [1],
        "pos":             1,
        "protocol":        [2,3],
        "sequence":        3,
        "startdelay":      -2,
        "w":               640
      }
    }
  ],
  "site": {
    "id": "1345135123",   	
    "name": "Site ABCD", 	
    "domain": "siteabcd.com",
    "cat": [
      "IAB2-1",
      "IAB2-2"
    ],
    "page": "http://siteabcd.com/page.htm",
    "ref": "http://referringsite.com/referringpage.htm",
    "privacypolicy": true,
    "publisher": {
      "id": "pub12345",
      "name": "Publisher A"
    },
    "content": {

      "cat":      ["IAB2-2"],
      "episode":  23,
      "id":       "1234567",
      "keyword":  ["keyword a", "keyword b", "keyword c"],
      "season":   2,
      "series":   "All About Cars",
      "title":    "Car Show"

    }
  },
  "device": {
    "ip": "64.124.253.1",
    "ua": "Mozilla/5.0 (Mac; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20140420 Firefox/3.6.16",             	
    "os": "OS X",
    "flashversion": "10.1",
    "js": 1
  },
  "user":
  {
    "uid": "456789876567897654678987656789",                   	
    "buyeruid": "545678765467876567898765678987654",
    "data": [
      {
        "id": "6",
        "name": "Data Provider 1",
        "segment": [
          {
            "id": "12341318394918",
            "name": "auto intenders"
          },
          {
            "id": "1234131839491234",
            "name": "auto enthusiasts"		
          }
        ]
      }
    ]
  }
})`)
