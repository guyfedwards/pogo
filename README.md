# PoGo
> GoGo PoGo.  

Quick link redirection for the browser.  

Search history is ok, but finicky across browsers and machines, some aliases I want consistent. Pogo is a simple http redirect server for aliases that you can point the browser omnibar at.

## Run
```sh
$ pogo 
```

## Config
Add aliases and their destinations to the config file in `$XDG_CONFIG_HOME/pogo/config.yaml`
### Example Config
```yaml
defaultSearch: "https://www.google.com/search?q=%s"
aliases:
  pandas: https://reddit.com/r/pandagifs
```

#### defaultSearch
Fallback url if no matching alias found. Replaces `%s` with the string passed from the browser.

#### aliases
Map of key/value pairs for name and values.

## Configuring the browser
### Firefox
https://superuser.com/questions/7327/how-to-add-a-custom-search-engine-to-firefox
