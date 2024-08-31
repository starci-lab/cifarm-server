import { BeforeAvalancheAuthenticate } from "./auth";
import { rpcHealthcheck } from "./healthcheck";

let InitModule: nkruntime.InitModule =
        function(ctx: nkruntime.Context, logger: nkruntime.Logger, nk: nkruntime.Nakama, initializer: nkruntime.Initializer) {
        initializer.registerBeforeAuthenticateCustom(BeforeAvalancheAuthenticate);
        initializer.registerRpc("typescript_healthcheck", rpcHealthcheck);
    logger.info("Hello World!");
}

!InitModule && InitModule.bind(null);