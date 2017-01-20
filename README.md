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

Make sure to remove the domain you want redirected from the app container and only place it on the redirect container.

##### Example Redirect `http://www.example.com` to `https://www.example.com`

```
docker run -e VIRTUAL_HOST=www.example.com -e VIRTUAL_PORT 80 lukehowell/redirect:ssl
docker run -e VIRTUAL_HOST=www.example.com -e VIRTUAL_PORT 443 nginx
```

##### Example Redirect `http://example.com` to `http://www.example.com`

```
docker run -e VIRTUAL_HOST=example.com lukehowell/redirect:www
docker run -e VIRTUAL_HOST=www.example.com nginx
```

##### Example Redirect `http://www.example.com` to `http://example.com`

```
docker run -e VIRTUAL_HOST=www.example.com lukehowell/redirect:non-www
docker run -e VIRTUAL_HOST=example.com nginx
```