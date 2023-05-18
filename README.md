
# Go HTTP Directory File Server

Use the standard Go http file server to serve a directory.

## Install

```
go install github.com/rhomel/http-serve-directory
```

## Usage

The default configuration will serve the current directory:
```
http-serve-directory
# output:
# 2023/05/17 23:52:43 serving directory: . (/home/my/current/working/dir)
# 2023/05/17 23:52:43 listening address: http://0.0.0.0:3000
```

Use a specific directory and address:
```
http-serve-directory -directory=/my/path -address=127.0.0.1:9999
# output:
# 2023/05/17 23:54:05 serving directory: /my/path
# 2023/05/17 23:54:05 listening address: http://127.0.0.1:9999
```

Use `http-serve-directory -h` to see more options.

