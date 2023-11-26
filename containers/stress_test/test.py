import requests
from time import sleep

print('Working...')
while(True):
    try:
        requests.get('http://localhost:81')
    except:
        sleep(2)
    