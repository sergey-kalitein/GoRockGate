# RockGate
## The RocketChat to OneSignal Gateway

### Configuring

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


### Building


`docker build -t sergey-kalitein/gorockgate:0.1 .`

`docker run -d -p 8181:8181 --name go-rock-gate sergey-kalitein/gorockgate:0.1 && docker attach go-rock-gate`

### Endpoints

As a Rocket Chat Gateway:

> `https://your.server.gateway:8181`

Endpoints (requested internally by the Rocket Chat instance):

> `/push/apn/send`
> 
> `/push/fcm/send`
> 
> `/push/web/send`

HTTP POST Body Example of a Push Notification:

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
