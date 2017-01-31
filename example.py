#!/usr/bin/env python
import requests
import json

url = 'http://localhost:18001/riemann'
payload = {'state': 'critical', 'host': 'foohost', 'service': 'foosvc', 'metric': "10", 'ttl': "0", 'tags': ['test','one','two','three']}

r = requests.post(url, data=json.dumps(payload))
print r.content
print r.status_code
