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

# function_urlMasterA = '' Add your AWS Lambda function URL
# function_urlMasterB = '' Add your AWS Lambda function URL
# function_urlMasterC = '' Add your AWS Lambda function URL

payload = {
}

# executeProcess(function_urlMasterA, payload, "A") # Execution of A Master Thread
# print("\n")
# executeProcess(function_urlMasterB, payload, "B") # Execution of B Master Thread
# print("\n")
# executeProcess(function_urlMasterC, payload, "C") # Execution of C Master Thread