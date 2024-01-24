import requests
import json

function_url = 'https://pdx3axs6nfkww2hia2irl55bwq0oubzi.lambda-url.us-east-1.on.aws/'

payload = {
    'input': 'CAM-A' # Classify and Modify - A
}

response = requests.post(function_url, data=json.dumps(payload))

if response.status_code == 200:
  print("Post sent successfully")
  print(response.json())
else:
  print("ERROR: Process was not executable")
  print(f"Reason: {response.status_code}")