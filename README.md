# FamPayProject
 Project to get Youtube video descriptions and search those videos as required

### Prerequisities
1) Postgres DB for persisting data
2) Youtube Data API Key for getting data from youtube

### Adding details
We have 2 layers of config
1) config.json file in FamPayProject>config
2) DB based configuration

**You'll have to add your DB creds and API related details in config.json**

#### config.json
This file has first level configs like your DB creds and your API key

#### DB config

You don't have to create any of these tables or add thing, some values will be inserted by default on the first run
which can be edited later as you wish

1) *video_search_config* table stores the config for which we will be fetching data from youtube
2) *api_configs* stores all the api keys, this will be helpful if you want to use multiple keys incase one key is exhausts

### APIs

#### Get all
URL - http://localhost:5000/list

Method - GET

query params - 
1) page - optional (default value is 1)

Response - 

```json
{
    "message": "OK",
    "body": [
        {
            "created_at": "2021-11-04T21:02:37.628055+05:30",
            "Name": "_I8GsD-51Bs",
            "Title": "✔️ LIVE MATCH : Bethel vs Kentridge - High School Football",
            "Description": "Watch Here : http://4ty.me/y4mic9 ™ STREAMING Today! High School Football, Bethel vs Kentridge Braves @ Chargers The Kentridge (Kent. WA) varsity ...",
            "URL": "https://i.ytimg.com/vi/_I8GsD-51Bs/default.jpg",
            "PublishedAt": "2021-11-04T15:32:18Z"
        },
        {
            "created_at": "2021-11-04T21:02:37.639637+05:30",
            "Name": "r2uCU1wSI3M",
            "Title": "FM22 | The Head Coach | Cefn Druids | EPISODE 15 - EUROPEAN DEBUT | Football Manager 2022",
            "Description": "Welcome along to episode 15 of my FM22 Head Coach story. We return for our annual journeyman story, combined with a director of football challenge. We start ...",
            "URL": "https://i.ytimg.com/vi/r2uCU1wSI3M/default.jpg",
            "PublishedAt": "2021-11-04T15:30:07Z"
        },
        {
            "created_at": "2021-11-04T20:57:37.574262+05:30",
            "Name": "7EIgsUamPkQ",
            "Title": "NCEUH (CO-OP) vs. Fertile-Beltrami - High School Football LIVE",
            "Description": "Watch Here : http://4ty.me/iw67nz ™ STREAMING Today! High School Football, NCEUH (CO-OP) vs Fertile-Beltrami Titans @ Falcons The Fertile-Beltrami ...",
            "URL": "https://i.ytimg.com/vi/7EIgsUamPkQ/default.jpg",
            "PublishedAt": "2021-11-04T15:26:53Z"
        },
        ....
    ],
    "page": 2
}
```

#### Smart Search
URL - http://localhost:5000/search

Method - GET

query params -
1) title
2) desc
3) page

Response - 

```json
{
    "message": "OK",
    "body": [
        {
            "created_at": "2021-11-04T21:02:37.628055+05:30",
            "Name": "_I8GsD-51Bs",
            "Title": "✔️ LIVE MATCH : Bethel vs Kentridge - High School Football",
            "Description": "Watch Here : http://4ty.me/y4mic9 ™ STREAMING Today! High School Football, Bethel vs Kentridge Braves @ Chargers The Kentridge (Kent. WA) varsity ...",
            "URL": "https://i.ytimg.com/vi/_I8GsD-51Bs/default.jpg",
            "PublishedAt": "2021-11-04T15:32:18Z"
        }
    ],
    "page": 2
}
```
