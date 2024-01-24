import requests
import json

def executeProcess(function_url, payload, specifier):
    response = requests.post(function_url, data=json.dumps(payload))

    if response.status_code == 200:
        print(f"Post sent successfully | Master-thread: {specifier}")
        print(response.json())
    else:
        print("ERROR: Process was not executable")
        print(f"Reason: {response.status_code}")

    print("\n")


function_urlMasterA = 'https://pdx3axs6nfkww2hia2irl55bwq0oubzi.lambda-url.us-east-1.on.aws/'
function_urlMasterB = 'https://pdx3axs6nfkww2hia2irl55bwq0oubzi.lambda-url.us-east-1.on.aws/'
function_urlMasterC = 'https://pdx3axs6nfkww2hia2irl55bwq0oubzi.lambda-url.us-east-1.on.aws/'

payload = {
}

executeProcess(function_urlMasterA, payload, "A") # Execution of A Master Thread
executeProcess(function_urlMasterB, payload, "B") # Execution of B Master Thread
executeProcess(function_urlMasterC, payload, "C") # Execution of C Master Thread