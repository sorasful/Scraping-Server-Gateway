# Scraping Server Gateway

First project written in Go, not external dependency !   
The purpose of this server is to provide some kind of **queue** to get url to scrape.
You can enqueue new urls, dequeue and send results (html code) back for further treatments.


## Add a new url in the queue

`/add?url=https://yolo.lol` - **GET** 

e.g :  
```bash
curl http://localhost:8080/add?url=https://googleqsdqds.fr
```

## Get a url to scrape 

`/to-scrape` - **GET** 

e.g :  
````bash
curl http://localhost:8080/to-scrape
````

## Send back page's result  

`/treat` - **POST**   
Body should contain the HTML code of the scraped page.
Headers must contains a header named **scraped_page** containing the url of the scraped page. 

e.g :  
````bash
curl -X POST localhost:8080/treat -H "scraped_page: http://lol.lol" --data "<html>......</html>"
````

## Build and run :  

### From sources :  

-  build it :  `go build server.py`  

-  run it : `./server`

### Using Docker

- Build image : 
`docker build -t scraping_server_gateway .`

Run on port 8080 :  `docker run  -p 8080:8080  -d scraping_server_gateway`