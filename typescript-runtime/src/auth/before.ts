// const BeforeAuthenticate: nkruntime.BeforeHookFunction<nkruntime.AuthenticateCustomRequest> = function (ctx: nkruntime.Context, logger: nkruntime.Logger, nk: nkruntime.Nakama, data: nkruntime.AuthenticateCustomRequest): nkruntime.AuthenticateCustomRequest | void {
//     const apiUrl = ctx.env['AUTHENTICATION_API_URL'];
//     if (!apiUrl) {
//       throw new Error('missing authentication api configuration');
//     }
  
//     // Construct a payload to send to the third-party API
//     const payload = JSON.stringify({
//       id: data.account.id
//     });
  
//     // Send the HTTP Post request to the API
//     /*
//     Expected API Response
//     HTTP 200
//     {
//       "userId": "<UserId>",
//       "username": "<Username>"  
//     }
//     */
//     const response = nk.httpRequest(apiUrl, 'post', { 'content-type': 'application/json' }, JSON.stringify(payload));
//     if (response.code > 299) {
//       logger.error(`API error: ${response.body}`);
//       return null
//     }
    
//     const userInfo = JSON.parse(response.body);
//     if (!userInfo.userId || !userInfo.username) {
//       logger.error(`invalid API response: ${response.body}`)
//       return null;
//     }
    
//     // Update the incoming authenticate request with the new user ID and username
//     data.account.id = userInfo.userId;
//     data.username = userInfo.username;
    
//     return data;
//   };