# AdCampaign


Commit 1:

created the service without integration of db

Commit 2: 

Added Postgresql 

Commit 3:

Added Prometheus

localhost:8080/metrics - this application metrics available

Sample Requests: 

Request 1 : GET http://localhost:8080/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android

Response Format:
[
  {
    "cid": "spotify",
    "cta": "Download",
    "img": "https://somelink"
  },
  {
    "cid": "subwaysurfer",
    "cta": "Play",
    "img": "https://somelink3"
  }
]

Request 2 : http://localhost:8080/v1/delivery?app=com.abc.xyz&country=germany&os=android

Response: 
[
  {
    "cid": "duolingo",
    "cta": "Install",
    "img": "https://somelink2"
  }
]

Request 3 : http://localhost:8080/v1/delivery?country=germany&os=android 

Response:
{
  "error": "missing app param"
}
