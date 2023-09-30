---
title: RouterOS REST Schema (v7.12beta3.28240783-all) v7.12beta3.28240783-all
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="routeros-rest-schema-v7-12beta3-28240783-all-">RouterOS REST Schema (v7.12beta3.28240783-all) v7.12beta3.28240783-all</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="https://{host}:{port}/rest">https://{host}:{port}/rest</a>

    * **host** -  Default: host

    * **port** -  Default: port

* <a href="http://{host}:{port}/rest">http://{host}:{port}/rest</a>

    * **host** -  Default: host

    * **port** -  Default: port

# Authentication

- HTTP Authentication, scheme: basic Mikrotik REST API only supports Basic Authentication, secured by HTTPS

<h1 id="routeros-rest-schema-v7-12beta3-28240783-all--default">Default</h1>

## POST_beep

<a id="opIdPOST_beep"></a>

> Code samples

```shell
# You can also use wget
curl -X POST https://{host}:{port}/rest/beep \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST https://{host}:{port}/rest/beep HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  ".proplist": "string",
  ".query": [],
  "as-value": "string",
  "frequency": "string",
  "length": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://{host}:{port}/rest/beep',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'https://{host}:{port}/rest/beep',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('https://{host}:{port}/rest/beep', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','https://{host}:{port}/rest/beep', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("https://{host}:{port}/rest/beep");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://{host}:{port}/rest/beep", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /beep`

> Body parameter

```json
{
  ".proplist": "string",
  ".query": [],
  "as-value": "string",
  "frequency": "string",
  "length": "string"
}
```

<h3 id="post_beep-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» .proplist|body|string|false|none|
|» .query|body|array|false|none|
|» as-value|body|string|false|none|
|» frequency|body|string|false|none|
|» length|body|string|false|none|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="post_beep-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|string|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad command or error|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|

<h3 id="post_beep-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## POST_blink

<a id="opIdPOST_blink"></a>

> Code samples

```shell
# You can also use wget
curl -X POST https://{host}:{port}/rest/blink \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST https://{host}:{port}/rest/blink HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  ".proplist": "string",
  ".query": [],
  "duration": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://{host}:{port}/rest/blink',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'https://{host}:{port}/rest/blink',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('https://{host}:{port}/rest/blink', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','https://{host}:{port}/rest/blink', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("https://{host}:{port}/rest/blink");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://{host}:{port}/rest/blink", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /blink`

> Body parameter

```json
{
  ".proplist": "string",
  ".query": [],
  "duration": "string"
}
```

<h3 id="post_blink-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» .proplist|body|string|false|none|
|» .query|body|array|false|none|
|» duration|body|string|false|none|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="post_blink-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|string|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad command or error|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|

<h3 id="post_blink-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## GET_certificate

<a id="opIdGET_certificate"></a>

> Code samples

```shell
# You can also use wget
curl -X GET https://{host}:{port}/rest/certificate \
  -H 'Accept: application/json'

```

```http
GET https://{host}:{port}/rest/certificate HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('https://{host}:{port}/rest/certificate',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get 'https://{host}:{port}/rest/certificate',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('https://{host}:{port}/rest/certificate', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','https://{host}:{port}/rest/certificate', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("https://{host}:{port}/rest/certificate");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://{host}:{port}/rest/certificate", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /certificate`

<h3 id="get_certificate-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|as-string|query|string|false|none|
|as-string-value|query|string|false|none|
|number|query|string|false|none|
|value-name|query|string|false|none|

> Example responses

> 200 Response

```json
[]
```

<h3 id="get_certificate-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad command or error|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|

<h3 id="get_certificate-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## PUT_certificate

<a id="opIdPUT_certificate"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT https://{host}:{port}/rest/certificate \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
PUT https://{host}:{port}/rest/certificate HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "common-name": "string",
  "copy-from": "string",
  "country": "string",
  "days-valid": "string",
  "digest-algorithm": "string",
  "key-size": "string",
  "key-usage": "string",
  "locality": "string",
  "name": "string",
  "organization": "string",
  "state": "string",
  "subject-alt-name": "string",
  "trusted": "string",
  "unit": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://{host}:{port}/rest/certificate',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.put 'https://{host}:{port}/rest/certificate',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.put('https://{host}:{port}/rest/certificate', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','https://{host}:{port}/rest/certificate', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("https://{host}:{port}/rest/certificate");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "https://{host}:{port}/rest/certificate", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /certificate`

> Body parameter

```json
{
  "common-name": "string",
  "copy-from": "string",
  "country": "string",
  "days-valid": "string",
  "digest-algorithm": "string",
  "key-size": "string",
  "key-usage": "string",
  "locality": "string",
  "name": "string",
  "organization": "string",
  "state": "string",
  "subject-alt-name": "string",
  "trusted": "string",
  "unit": "string"
}
```

<h3 id="put_certificate-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» common-name|body|string|false|none|
|» copy-from|body|string|false|none|
|» country|body|string|false|none|
|» days-valid|body|string|false|none|
|» digest-algorithm|body|string|false|none|
|» key-size|body|string|false|none|
|» key-usage|body|string|false|none|
|» locality|body|string|false|none|
|» name|body|string|false|none|
|» organization|body|string|false|none|
|» state|body|string|false|none|
|» subject-alt-name|body|string|false|none|
|» trusted|body|string|false|none|
|» unit|body|string|false|none|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="put_certificate-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|string|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad command or error|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|

<h3 id="put_certificate-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## POST_certificate-add

<a id="opIdPOST_certificate-add"></a>

> Code samples

```shell
# You can also use wget
curl -X POST https://{host}:{port}/rest/certificate/add \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST https://{host}:{port}/rest/certificate/add HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  ".proplist": "string",
  ".query": [],
  "common-name": "string",
  "copy-from": "string",
  "country": "string",
  "days-valid": "string",
  "digest-algorithm": "string",
  "key-size": "string",
  "key-usage": "string",
  "locality": "string",
  "name": "string",
  "organization": "string",
  "state": "string",
  "subject-alt-name": "string",
  "trusted": "string",
  "unit": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://{host}:{port}/rest/certificate/add',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'https://{host}:{port}/rest/certificate/add',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('https://{host}:{port}/rest/certificate/add', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','https://{host}:{port}/rest/certificate/add', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("https://{host}:{port}/rest/certificate/add");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://{host}:{port}/rest/certificate/add", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /certificate/add`

> Body parameter

```json
{
  ".proplist": "string",
  ".query": [],
  "common-name": "string",
  "copy-from": "string",
  "country": "string",
  "days-valid": "string",
  "digest-algorithm": "string",
  "key-size": "string",
  "key-usage": "string",
  "locality": "string",
  "name": "string",
  "organization": "string",
  "state": "string",
  "subject-alt-name": "string",
  "trusted": "string",
  "unit": "string"
}
```

<h3 id="post_certificate-add-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» .proplist|body|string|false|none|
|» .query|body|array|false|none|
|» common-name|body|string|false|none|
|» copy-from|body|string|false|none|
|» country|body|string|false|none|
|» days-valid|body|string|false|none|
|» digest-algorithm|body|string|false|none|
|» key-size|body|string|false|none|
|» key-usage|body|string|false|none|
|» locality|body|string|false|none|
|» name|body|string|false|none|
|» organization|body|string|false|none|
|» state|body|string|false|none|
|» subject-alt-name|body|string|false|none|
|» trusted|body|string|false|none|
|» unit|body|string|false|none|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="post_certificate-add-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|string|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad command or error|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|

<h3 id="post_certificate-add-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>

## POST_certificate-add-scep

<a id="opIdPOST_certificate-add-scep"></a>

> Code samples

```shell
# You can also use wget
curl -X POST https://{host}:{port}/rest/certificate/add-scep \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST https://{host}:{port}/rest/certificate/add-scep HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  ".proplist": "string",
  ".query": [],
  "as-value": "string",
  "ca-identity": "string",
  "challenge-password": "string",
  "name": "string",
  "on-smart-card": "string",
  "refresh": "string",
  "scep-url": "string",
  "template": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('https://{host}:{port}/rest/certificate/add-scep',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'https://{host}:{port}/rest/certificate/add-scep',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('https://{host}:{port}/rest/certificate/add-scep', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','https://{host}:{port}/rest/certificate/add-scep', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("https://{host}:{port}/rest/certificate/add-scep");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "https://{host}:{port}/rest/certificate/add-scep", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /certificate/add-scep`

> Body parameter

```json
{
  ".proplist": "string",
  ".query": [],
  "as-value": "string",
  "ca-identity": "string",
  "challenge-password": "string",
  "name": "string",
  "on-smart-card": "string",
  "refresh": "string",
  "scep-url": "string",
  "template": "string"
}
```

<h3 id="post_certificate-add-scep-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|object|true|none|
|» .proplist|body|string|false|none|
|» .query|body|array|false|none|
|» as-value|body|string|false|none|
|» ca-identity|body|string|false|none|
|» challenge-password|body|string|false|none|
|» name|body|string|false|none|
|» on-smart-card|body|string|false|none|
|» refresh|body|string|false|none|
|» scep-url|body|string|false|none|
|» template|body|string|false|none|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="post_certificate-add-scep-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|string|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad command or error|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|Inline|

<h3 id="post_certificate-add-scep-responseschema">Response Schema</h3>

<aside class="success">
This operation does not require authentication
</aside>
