# httpdump

`httpdump` is a basic HTTP server dumping received requests content, which can be handy to debug HTTP clients requests. It doesn't do anything besides dumping the headers and content of HTTP requests it receives, and always returns `200 OK`.

Note: this utility is currently in a very rudimentary form since it started as a quick'n dirty debugging tool, it may be improved in the future.

## Build

Building httpdump from sources requires the GNU Make and [gb](https://getgb.io/) utility. To build it, simply execute `make` at the root of the source tree, and if everything goes well the binary can be found in the `bin/` directory.

## Usage

```
Usage: httpdump [OPTIONS]

Options:
   -bind  Address to bind on. If this value has a colon, as in ":8000" or
		"127.0.0.1:9001", it will be treated as a TCP address. If it
		begins with a "/" or a ".", it will be treated as a path to a
		UNIX socket. If it begins with the string "fd@", as in "fd@3",
		it will be treated as a file descriptor (useful for use with
		systemd, for instance). If it begins with the string "einhorn@",
		as in "einhorn@0", the corresponding einhorn socket will be
		used. If an option is not explicitly passed, the implementation
		will automatically select among "einhorn@0" (Einhorn), "fd@3"
		(systemd), and ":8000" (fallback) based on its environment.
   -h  display this help and exit
   -version  display version and exit
```

Given the following HTTP request using `curl`:

```
$ curl -X PUT -d key=value localhost:8000
```

`httpdump` outputs (`[H]`-prefixed lines are HTTP request headers, `[B]`-prefixed line is the HTTP request body if any):

```
2016/01/24 22:47:43.374892 [oxide.local/5c0JlaPI6B-000002] Started PUT "/" from [::1]:50034
[H] Content-Length: 9
[H] Content-Type: application/x-www-form-urlencoded
[H] User-Agent: curl/7.43.0
[H] Accept: */*
[B] key=value
2016/01/24 22:47:43.375022 [oxide.local/5c0JlaPI6B-000002] Returning 200 in 94.127µs
```

If the request `Content-Type` is "application/json" and the body is valid JSON data, `httpdump` will pretty-print it:

```
$ curl -X POST -H 'Content-Type: application/json' -d '{"key1":"value","key2":42,"key3":["a","b","c"]}' localhost:8000
```

```
2016/01/24 22:48:03.387585 [oxide.local/5c0JlaPI6B-000003] Started POST "/" from [::1]:50038
[H] User-Agent: curl/7.43.0
[H] Accept: */*
[H] Content-Type: application/json
[H] Content-Length: 15
[B] {
  "key1": "value",
  "key2": 42,
  "key3": [
    "a",
    "b",
    "c"
  ]
}
2016/01/24 22:48:03.387870 [oxide.local/5c0JlaPI6B-000003] Returning 200 in 248.491µs
```

## License

`httpdump` is released under the MIT license, see the LICENSE.md file for details.