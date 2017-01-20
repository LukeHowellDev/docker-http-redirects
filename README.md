# Docker - Http Requests

## ssl

```
docker run lukehowell/redirect:ssl

```

## www

```
docker run lukehowell/redirect:www

```

## non www

```
docker run lukehowell/redirect:non-www

```

## Use with [jwilder/nginx-proxy](https://hub.docker.com/r/jwilder/nginx-proxy/)

To use with the nginx-proxy simply set the URL you want redirected in the `VIRTUAL_HOST` environment variable.
With the ssl redirect you will need to add the `VIRTUAL_PORT`

##### Example Redirect `http://www.example.com` to `https://www.example.com`

```
docker run -e VIRTUAL_HOST=www.example.com -e VIRTUAL_PORT 80 lukehowell/redirect:ssl

```

##### Example Redirect `http://example.com` to `http://www.example.com`

```
docker run -e VIRTUAL_HOST=example.com lukehowell/redirect:www

```

##### Example Redirect `http://www.example.com` to `http://example.com`

```
docker run -e VIRTUAL_HOST=www.example.com lukehowell/redirect:non-www

```