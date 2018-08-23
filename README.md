# DonDiscord

Peter wanted me to provide memes on discord.
This project was just an easy way for me to learn golang. Feedback, criticism and beer is always appreciated

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

Golang1.9.


### Installing

```
go get github.com/bwmarrin/discordgo
go get github.com/erraa/dondiscord

cp example.config.json $HOME/dondiscord.json
vim $HOME/dondiscord.json
go build
```

And then either just run the binary or create a servicefile /etc/systemd/system

```
dondiscord
```

### And coding style tests

Fmt my friends

```
import "fmt"
```

## Authors

The Don himself

## License

Buy me a beer

## Acknowledgments

* Hat tip to anyone who's code was used
* People who buys me beer

<img src="http://www.vipbacking.eu/midicovers/CR2169.jpg" width="50%" height="50%" >
