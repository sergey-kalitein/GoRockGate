# RockGate
> **The RocketChat to OneSignal Gateway**

## Configuring

The configuration file path:

`conf/rockgate.yml`

Configuration file example:
```yaml
---
OneSignalUserKey:                 "Your-One-Signal-User-Key-Here"
GCMKey:                           "your-gcm-key-data-here"
AndroidGCMSenderID: &GCMSenderID  "0123456789012"  # not used in the library code, should implement?
ChromeKey:                        ""  # not used for web push
ChromeWebKey:                     ""  # not used for web push
ChromeWebGCMSenderID:             *GCMSenderID
APNSEnv:                          "production"
APNSP12:                          "...
                                   ...
                                   The BASE64 encoded P12 certificate file for Apple Push Messaging 
                                   ...
                                   ...
                                   ="
APNSP12Password:                  "your-P12-secret-password"
SafariAPNSP12:                    "Safari Web Push BASE64-Encoded P12 File Data"
SafariAPNSP12Password:            "your-safari-P12-secret-password"
```


## Building


`docker build -t sergey-kalitein/gorockgate:0.1 .`

`docker run -d -p 8181:8181 --name go-rock-gate sergey-kalitein/gorockgate:0.1 && docker attach go-rock-gate`

## API Endpoints

Available Endpoints:
- [List Apps](#list-apps)
  
- [Find (Get) App by Domain](#find-app-by-domain)
  
- [Get App (or Create)](#get-app-or-create)
  
- [Push Message](#push-message)


### List Apps


- **HTTP Method:** `GET`

- Endpoint Pattern: `/apps/list`

**HTTP Status:** `200 Ok`

**HTTP Response Example:**
```json
{
    "beautiful.chat": {
        "id": "293ca9ba-2543-46c4-8874-7ec5cd73da0b",
        "name": "beautiful.chat",
        "players": 0,
        "messagable_players": 0,
        "updated_at": "2021-03-16T09:55:33.71Z",
        "created_at": "2021-03-15T13:05:13.691Z",
        "gcm_key": "*******************",
        "chrome_key": "",
        "chrome_web_origin": "https://beautiful.chat",
        "chrome_web_gcm_sender_id": "",
        "chrome_web_default_notification_icon": "",
        "chrome_web_sub_domain": "",
        "apns_env": "production",
        "apns_certificates": "*******************\n",
        "safari_apns_cetificate": "",
        "safari_site_origin": "https://beautiful.chat",
        "safari_push_id": "",
        "safari_icon_16_16": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/16x16.png",
        "safari_icon_32_32": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/16x16@2x.png",
        "safari_icon_64_64": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/32x32@2x.png",
        "safari_icon_128_128": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/128x128.png",
        "safari_icon_256_256": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/128x128@2x.png",
        "site_name": "The beautiful.chat Website",
        "basic_auth_key": "*******************"
    },
    "sg.my.chat": {
        "id": "5e37ab88-ac42-4839-975e-9f0d18df7b8b",
        "name": "sg.my.chat",
        "players": 0,
        "messagable_players": 0,
        "updated_at": "2021-03-16T12:18:24.561Z",
        "created_at": "2021-03-16T12:16:26.09Z",
        "gcm_key": "*******************",
        "chrome_key": "",
        "chrome_web_origin": "https://sg.my.chat",
        "chrome_web_gcm_sender_id": "",
        "chrome_web_default_notification_icon": "",
        "chrome_web_sub_domain": "",
        "apns_env": "production",
        "apns_certificates": "*******************",
        "safari_apns_cetificate": "",
        "safari_site_origin": "",
        "safari_push_id": "web.onesignal.auto.28671d66-3da8-4a50-bcc4-1b29e015670b",
        "safari_icon_16_16": "public/safari_packages/5e37ab88-ac42-4839-975e-9f0d18df7b8b/icons/16x16.png",
        "safari_icon_32_32": "public/safari_packages/5e37ab88-ac42-4839-975e-9f0d18df7b8b/icons/16x16@2x.png",
        "safari_icon_64_64": "public/safari_packages/5e37ab88-ac42-4839-975e-9f0d18df7b8b/icons/32x32@2x.png",
        "safari_icon_128_128": "public/safari_packages/5e37ab88-ac42-4839-975e-9f0d18df7b8b/icons/128x128.png",
        "safari_icon_256_256": "public/safari_packages/5e37ab88-ac42-4839-975e-9f0d18df7b8b/icons/128x128@2x.png",
        "site_name": "'sg.my.chat' website",
        "basic_auth_key": "*******************"
    }
}
```

**HTTP Status:** `400 Bad Request`

**HTTP Response Example:**

```json
{
    "error_text": "error message text here"
}
```


### Find App by Domain

- **HTTP Method:** `GET`

- Endpoint Pattern: `/apps/find/{domain:[^/]+}`

- Endpoint Example: `/apps/find/sg.workspee.chat`


**HTTP Status:** `200 Ok` 

**HTTP Response Example:**
```json
{
    "id": "293ca9ba-2543-46c4-8874-7ec5cd73da0b",
    "name": "beautiful.chat",
    "players": 0,
    "messagable_players": 0,
    "updated_at": "2021-03-16T09:55:33.71Z",
    "created_at": "2021-03-15T13:05:13.691Z",
    "gcm_key": "*******************",
    "chrome_key": "",
    "chrome_web_origin": "https://beautiful.chat",
    "chrome_web_gcm_sender_id": "",
    "chrome_web_default_notification_icon": "",
    "chrome_web_sub_domain": "",
    "apns_env": "production",
    "apns_certificates": "*******************\n",
    "safari_apns_cetificate": "",
    "safari_site_origin": "https://beautiful.chat",
    "safari_push_id": "",
    "safari_icon_16_16": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/16x16.png",
    "safari_icon_32_32": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/16x16@2x.png",
    "safari_icon_64_64": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/32x32@2x.png",
    "safari_icon_128_128": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/128x128.png",
    "safari_icon_256_256": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/128x128@2x.png",
    "site_name": "The beautiful.chat Website",
    "basic_auth_key": "*******************"
}
```


**HTTP Status:** `404 Not Found`

**HTTP Response Example:**

```json
{
  "error_text": "application not found"
}
```


### Get App (or Create)

- **HTTP Method:** `GET`

- Endpoint Pattern: `/apps/find-or-create/{domain:[^/]+}`

- Endpoint Example: `/apps/find/my.new.beautiful.chat`


**HTTP Status:** `200 Ok`

**HTTP Response Example:**
```json
{
    "id": "293ca9ba-2543-46c4-8874-7ec5cd73da0b",
    "name": "my.new.beautiful.chat",
    "players": 0,
    "messagable_players": 0,
    "updated_at": "2021-03-16T09:55:33.71Z",
    "created_at": "2021-03-15T13:05:13.691Z",
    "gcm_key": "*******************",
    "chrome_key": "",
    "chrome_web_origin": "https://my.new.beautiful.chat",
    "chrome_web_gcm_sender_id": "",
    "chrome_web_default_notification_icon": "",
    "chrome_web_sub_domain": "",
    "apns_env": "production",
    "apns_certificates": "*******************\n",
    "safari_apns_cetificate": "",
    "safari_site_origin": "https://my.new.beautiful.chat",
    "safari_push_id": "",
    "safari_icon_16_16": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/16x16.png",
    "safari_icon_32_32": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/16x16@2x.png",
    "safari_icon_64_64": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/32x32@2x.png",
    "safari_icon_128_128": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/128x128.png",
    "safari_icon_256_256": "public/safari_packages/293ca9ba-2543-46c4-8874-7ec5cd73da0b/icons/128x128@2x.png",
    "site_name": "The My New beautiful.chat App",
    "basic_auth_key": "*******************"
}
```


**HTTP Status:** `400 Bad Request`

**HTTP Response Example:**

```json
{
  "error_text": "error text will be here"
}
```


### Push Message

These Endpoints are requested internally by the Rocket Chat instance:

> `/push/apn/send`
> 
> `/push/fcm/send`
> 
> `/push/web/send`

- **HTTP Method** `POST` 
  
- HTTP POST Body Payload Example:

```json
{
  "token": "f547c177ca7f6421910c673e2e09165121bcdc62d3a92e249d6ad0c627373359",
  "options": {
    "createdAt": "2021-03-12T10:09:17.925Z",
    "createdBy": "<SERVER>",
    "sent": false,
    "sending": 0,
    "from": "push",
    "title": "@sg",
    "text": "This is a push test message",
    "userId": "gR6Hhq5aEDdGswSQY",
    "sound": "default",
    "apn": {
      "text": "@sg:\nThis is a push test message"
    },
    "site_url": "https://sg.yourdomain.chat",
    "topic": "com.app.your-domain.chat",
    "uniqueId": "no33sYn6N2fb8JNXm"
  }
}
```

**HTTP Status:** `200 Ok`

**HTTP Response Example:**
```json
{
    "id": "",
    "recipients": 0,
    "errors": [
        "All included players are not subscribed"
    ]
}
```
