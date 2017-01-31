# riemann-gw
A golang RESTful gateway for submitting riemann data

# building for linux on a mac
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v fake/riemann-gw/gw.go

# testing with python (since there's that annoying protobuf+python3 issue)
```python
#!/usr/bin/env python
import requests
import json

url = 'http://localhost:8080/riemann'
payload = {'state': 'critical', 'host': 'foohost', 'service': 'foosvc', 'metric': "10"}

r = requests.post(url, data=json.dumps(payload))
print r.content
print r.status_code
```
