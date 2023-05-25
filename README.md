# Command Line Tool to get Lyrics of malagasy protestants songs

- [ ] Songs 1-10
- [x] Songs from number 100 in ffpm
- [ ] Add <b>fihirana fanampiny</b> songs
- [ ] Add `--help` command to display commands
- [ ] Handle errors when network is down
- [ ] Handle errors when song doesn't exist
- [ ] Create SDK for Golang
- [ ] Create GRPC API
- [ ] Run GRPC API with command ``fihirana grpc``
- [ ] Put all songs lyrics in a json file
- [ ] Commands to choose `fihirana [ffpm|fanampiny|antema|fifohazana] -n [number]`
- [ ] Commands to seach by title
- [ ] Select searched song in an interactive mode

## Installation for developers

### Installation of needed packages
```
$ go mod download    
```

### Build the comand-line
```
$ go build main.go
```

### Search for a song
```
$ ./main -n [song_number]
```
