# kumotest
this is a test for Kumojin.

## Description
The purpose of the test is to display the time for different time zones.

## How to test?
1: Run: docker-compose up -d
2: build CLI app: 
```
cd client
go build
```
3: exec binaire: "client" (windows: "client.exe"
4: The client display the server timezone and the local.

## API Developer

/timezone
Method: GET
Return content-type: json
```
{
	TimeZone: "Server_timezone"
}
```

### Credits
Thanks to Kumojin for the technical test. The moral of the story, there are several facets to an exercise.
