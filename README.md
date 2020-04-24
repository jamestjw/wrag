# wrag - Wrapper Reddit API in Golang

Features:
* Return random post from a particular subreddit.

Example: 
```
package main

import "github.com/jamestjw/wrag"

func main() {
    wrag.Initialise("/path/to/config.yml)
    listing := wrag.Random("subreddit-name").Details()
    fmt.Println(listing.Title)
    fmt.Println(listing.MediaURL)
}
```

Refer to `config.yml.example` for a sample version of `config.yml`.
