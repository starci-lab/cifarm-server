# CiFarm Server
Developed using Go and the Nakama framework.
<br/>
<img src="https://cdn.icon-icons.com/icons2/2699/PNG/512/golang_logo_icon_171073.png" alt="Example Image" height="100" style="margin-right:16px">
<img src="https://heroiclabs.com/images/pages/nakama/dark-logo.460a1d86928149a160b025152207fea9659bedd7cb818ee8971d6ee536776bfe.svg" alt="Example Image" height="100">
<br/>
### Introduction
![CiFarmBackground](https://scontent.xx.fbcdn.net/v/t1.15752-9/456037244_1297281141439021_1125970913282244783_n.png?_nc_cat=111&ccb=1-7&_nc_sid=0024fc&_nc_eui2=AeEHNQutC5nQX6XCBg2wFN3Pc-swk58WUx5z6zCTnxZTHi-HeUcH4zWtDYjgW2Id-JUs4ReGCtJHsJ2tG1GxEiHG&_nc_ohc=fG5d9Mt-5fEQ7kNvgERdcZ_&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.xx&oh=03_Q7cD1QGH_oYTt_MZFFSq99GpENcx1L1iREcnflEujxQE6wEN_A&oe=670176C6)

## Getting Started
CiFarm Server is now live and fully operational in the production environment.<br>
<u>API Endpoint:</u> You can access the API at [https://api.cifarm-server.starci.net/](https://api.cifarm-server.starci.net/)

## Authentication
#### 1.  Verifying Blockchain Credentials
Access the Cibase API at https://blockchain-auth-service.starci.net/api and request the message by sending a POST request to https://blockchain-auth-service.starci.net/api/v1/authenticator/request-message. Sign the retrieved message using your blockchain credentials, and then submit it to the CiFarm server for verification. You have a time limit of 1 minute to complete this process.
#### 2. Implementation of Authentication Code (Unity)

```csharp   
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
#### 1. Buy Seed
RPC Name: `buy_seed`
<br/>
Request Body:  
```json
{
    "key" : "carrot",
    "quantity" : 5
}
```
Response:  
```json
{
    "totalCost" : 2500
}
```

The buy_seed function allows you to purchase seeds. You specify the type of seed (e.g., "carrot") and the quantity you want (e.g., 5). The response will show you the total cost of the seeds, which will be deducted from your wallet.
<br/>
Throw errors if there are issues such as insufficient balance or an invalid plant key, ...
#### 2. Buy Animal
RPC Name: `buy_animal`
<br/>
Request Body:  
```json
{
    "key" : "chicken"
}
```
Response:  
```json
{
    "cost" : 1000
}
```

The buy_animal function lets you purchase an animal by specifying its type (e.g., "chicken"). The response will display the cost of the animal, which will be deducted from your wallet. 
<br/>
Throw errors if there are issues such as insufficient balance or an invalid animal key, ...
#### 3. Plant Seed
RPC Name: `plant_seed`
<br/>
Request Body:  
```json
{
    "inventorySeedKey" : "d6a6a181-41e5-4389-96d9-89a1f5a58b13",
    "placedItemTileKey" : "e6df5293-0338-4e33-b6b2-24ae2ecdd24d"
}
```
Response:  
```json
{
    "harvestIn" : 36000
}
```
The plant_seed function allows you to plant a seed. You need to provide the unique key for the seed from your inventory (e.g., "d6a6a181-41e5-4389-96d9-89a1f5a58b13") and the key for the tile where you want to plant it (e.g., "e6df5293-0338-4e33-b6b2-24ae2ecdd24d"). The response will indicate the time until the seed is ready for harvest (e.g., "harvestIn": 36000 seconds).
<br/>
Throw errors if the seed key does not exist or if the tile has already been planted, ...

## Server Authoritative
#### 1. Get the match ID
Step 1: Get the value of centralMatchInfo object
```csharp
var personalStorageId = new StorageObjectId
{
    Collection = "System",
    UserId = session.UserId,
    Key = "centralMatchInfo"
};

var personalStorageObjects = await client.ReadStorageObjectsAsync(session, new[]
{
    personalStorageId
} );
var value = personalStorageObjects.Objects.ToList()[0].Value;
```
Step 2: Define class & Deserialize
```csharp
class CentralMatchInfo
{
    [JsonProperty("matchId")]   
    public string MatchId { get; set; }
}
```
```csharp
var matchId = "";
var result = JsonConvert.DeserializeObject<CentralMatchInfo>(value);
if (result != null)
{
    matchId = result.MatchId;
}
```
Step 3: Get the socket & Join the match
```csharp
var socket = client.NewSocket(true);
await socket.ConnectAsync(session);
var match = await socket.JoinMatchAsync(matchId);
```



#### 2. Get Realtime Tiles State
You can retrieve the real-time state of your farm's tiles, with updates occurring once per second.

```csharp
 socket.ReceivedMatchState += newState =>
 {
    //trigger once per second
    var content = enc.GetString(newState.State);
    switch (newState.OpCode)
    {
        case 1:
            Debug.Log(content);
            break;
        default:
            break;
    }
 };
```