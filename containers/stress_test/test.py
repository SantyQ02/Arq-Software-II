import requests
from time import sleep

print('Working...')
while(True):
    try:
        requests.get('http://localhost:83/api/search?city=9&check_in_date=2023-11-21&check_out_date=2023-11-28')
    except:
        sleep(2)
    