# Crypto API tester

Provides an API with some endpoints to create MD5, SHA-1 or SHA-256 hashes. For testing FIPS-140 compliance

# Usage
```bash
$ curl -X POST http://127.0.0.1:9000/md5 -d "Hello"
$ curl -X POST http://127.0.0.1:9000/sha1 -d "Hello"
$ curl -X POST http://127.0.0.1:9000/sha256 -d "Hello"
```
