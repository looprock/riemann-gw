#!/usr/bin/env python
import requests
import json
import sys
from optparse import OptionParser
import shlex

cmd_parser = OptionParser(version="%prog 0.1")
cmd_parser.add_option("-d", "--debug", action="store_true", dest="debug", help="debug mode")
cmd_parser.add_option("-m", "--metric", action="store", dest="metric", help="specify metric", default="10")
cmd_parser.add_option("-r", "--riemannserver", action="store", dest="riemannserver", help="specify riemann server", default="localhost")
cmd_parser.add_option("-v", "--service", action="store", dest="service", help="specify service name", default="fakesrv")
cmd_parser.add_option("-s", "--state", action="store", dest="state", help="state", default="ok")
cmd_parser.add_option("-x", "--hostname", action="store", dest="hostname", help="state", default="fakehost")
cmd_parser.add_option("-l", "--ttl", action="store", dest="ttl", help="ttl", default="30")
cmd_parser.add_option("-t", "--tags", action="store", dest="tags", help="commma separated tags", default="test,fake")
(cmd_options, cmd_args) = cmd_parser.parse_args()

tags = cmd_options.tags.split(",")

url = 'http://%s:18001/riemann' % cmd_options.riemannserver
payload = {'state': cmd_options.state, 'host': cmd_options.hostname, 'service': cmd_options.service, 'metric': cmd_options.metric, 'ttl': cmd_options.ttl, 'tags': tags}

r = requests.post(url, data=json.dumps(payload))
print "Status code: %d" % r.status_code
print "Response:"
print r.content
