# CiFarm Server
Developed using Go and the Nakama framework.
<br/>
<img src="https://cdn.icon-icons.com/icons2/2699/PNG/512/golang_logo_icon_171073.png" alt="Example Image" height="50" style="margin-right:8px">
<img src="https://heroiclabs.com/images/pages/nakama/dark-logo.460a1d86928149a160b025152207fea9659bedd7cb818ee8971d6ee536776bfe.svg" alt="Example Image" height="50">
<br/>
### Introduction
![CiFarmBackground](https://scontent.xx.fbcdn.net/v/t1.15752-9/456037244_1297281141439021_1125970913282244783_n.png?_nc_cat=111&ccb=1-7&_nc_sid=0024fc&_nc_eui2=AeEHNQutC5nQX6XCBg2wFN3Pc-swk58WUx5z6zCTnxZTHi-HeUcH4zWtDYjgW2Id-JUs4ReGCtJHsJ2tG1GxEiHG&_nc_ohc=fG5d9Mt-5fEQ7kNvgERdcZ_&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.xx&oh=03_Q7cD1QGH_oYTt_MZFFSq99GpENcx1L1iREcnflEujxQE6wEN_A&oe=670176C6)

## Getting Started
CiFarm Server is now live and fully operational in the production environment.<br>
<u>API Endpoint:</u> You can access the API at [https://api.cifarm-server.starci.net/](https://api.cifarm-server.starci.net/)

### Authentication
#### 1  Verifying Blockchain Credentials
Access the Cibase API at https://blockchain-auth-service.starci.net/api and request the message by sending a POST request to https://blockchain-auth-service.starci.net/api/v1/authenticator/request-message. Sign the retrieved message using your blockchain credentials, and then submit it to the CiFarm server for verification. You have a time limit of 1 minute to complete this process.
#### 2 Implementation of Authentication Code (Unity)

```   
async void Start()
    {
        const string scheme = "https";
        const string host = "https://api.cifarm-server.starci.net/";
        const int port = 443;
        const string serverKey = "defaultkey";
        var client = new Client(scheme, host, port, serverKey);
        var session = await client.AuthenticateCustomAsync(
    "",
    "",
    false,
    new Dictionary<string, string>
    {
    { "message", "<your_message>" },
    { "publicKey", "<your_publicKey>" },
    { "signature", "<your_signature>" },
    { "chain", "<your_chain>" }
    }
);
        var healthcheck = await client.RpcAsync(session, "go_healthcheck");
        Debug.Log(JsonConvert.SerializeObject(healthcheck.Payload));
        }
```
## RPCs
1. Buy Seed