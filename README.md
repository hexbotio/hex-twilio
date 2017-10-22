# Hex Plugin - Twilio

Hex Plugin which will send SMS messages via your Twilio account.

```
{
  "rule": "example twilio rule",
  "match": "page somebody",
  "actions": [
    {
      "type": "hex-twilio",
      "command": "Page from a Bot!",
      "config": {
        "account_sid": "${TWILIO_ACCOUNT_SID}",
        "auth_token": "${TWILIO_AUTH_TOKEN}",
        "send_to": "317-555-1234",
        "send_from": "317-555-5555"
      }
    }
  ]
}
```
