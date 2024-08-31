let InitModule: nkruntime.InitModule =
        function(ctx: nkruntime.Context, logger: nkruntime.Logger, nk: nkruntime.Nakama, initializer: nkruntime.Initializer) {
        initializer.registerRpc("typescript_healthcheck", rpcHealthcheck);
    logger.info("Hello World!");
}

!InitModule && InitModule.bind(null);