# E621Query CLI image downloader

## Requiriments:
- latest golang installed
- that's it! 

## Installing:
- download the tool with `go get github.com/R0X-withazero/E621Query`
- compile it with `go build -o e621 main.go`

## Usage:
### Get Command list help `./e621 -h`
  **all of the downloaded files are stored in the E621/ folder**
```
Usage of e621:
  -limit int
        Number of pics to display per page (default 1)
  -pages int
        Number of pages to query (default 1)
  -tags string
        Tags for query: tag+tag
```

### Common search:
```
 ./e621 -tags "furry+-bdsm" -pages 1 -limit 10
 Downloading https://static1.e621.net/data/69/9f/699fa96396418578aa3f54a1f635bf1c.jpg to 699fa96396418578aa3f54a1f635bf1c.jpg                                            829509 bytes downloaded.                                                                                                                                                Downloading https://static1.e621.net/data/22/ce/22ceea2ba867697eebe92564d42cda3a.jpg to 22ceea2ba867697eebe92564d42cda3a.jpg                                            535623 bytes downloaded.                                                                                                                                                Downloading https://static1.e621.net/data/89/ea/89eab03eafd84b9830917e20749a18d9.jpg to 89eab03eafd84b9830917e20749a18d9.jpg                                            344633 bytes downloaded.                                                                                                                                                Downloading https://static1.e621.net/data/a3/e1/a3e156064829efac31afc7e42f3ef958.jpg to a3e156064829efac31afc7e42f3ef958.jpg                                            258823 bytes downloaded.                                            ...                                                            
```