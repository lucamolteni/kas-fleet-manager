import random, string

# random_string is used to create kafka_request names
def random_string(length):
   return ''.join(random.choice(string.ascii_lowercase) for i in range(length))

# random_status - get random status of kafka_cluster
def random_status():
  return random.choice(['accepted', 'preparing', 'provisioning', 'ready'])

# kafka_json is used to create json payload 
def kafka_json():
  return {
           'cloud_provider': 'aws',
           'region': 'us-east-1',
           'name': random_string(15),
           'multi_az': True
         }

# get_random_search query to get with search param
def get_random_search():
  return f'cloud_provider = aws and status = {random_status()} and name <> {random_string(15)}'

# get random id from list
def get_random_id(list):
  return random.choice(list)

# generate random service account id
def generate_random_svc_acc_id():
  return f'{random_alfa_numeric(7)}-{random_alfa_numeric(4)}-{random_alfa_numeric(4)}-{random_alfa_numeric(4)}-{random_alfa_numeric(11)}'

# generate random service account client ID
def generate_random_svc_acc_client_id():
  return f'srvc-acct-{random_alfa_numeric(8)}-{random_alfa_numeric(4)}-{random_alfa_numeric(4)}-{random_alfa_numeric(4)}-{random_alfa_numeric(12)}'

# generate random service account secret
def generate_random_svc_acc_secret():
  return f'{random_alfa_numeric(8)}-{random_alfa_numeric(4)}-{random_alfa_numeric(4)}-{random_alfa_numeric(4)}-{random_alfa_numeric(12)}'

# get random alpha-numeric string
def random_alfa_numeric(length):
  return ''.join(random.choices(string.ascii_lowercase + string.digits, k=length))

def safe_delete(list, item):
  if item in list:
    list.remove(item)

# generate and return service account json
def svc_acc_json(base_url):
  svc_acc_id = generate_random_svc_acc_id()
  return {
           'id': svc_acc_id,
           'href': f'{base_url}/serviceaccounts/{svc_acc_id}',
           'clientID': generate_random_svc_acc_client_id(),
           'name': random_string(10),
           'description': 'Created by the performance test suite of the API',
           'clientSecret': generate_random_svc_acc_secret()
         }
