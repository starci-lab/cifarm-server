var InitModule = function InitModule(ctx, logger, nk, initializer) {
  initializer.registerRpc("typescript_healthcheck", rpcHealthcheck);
  logger.info("Hello World!");
};
!InitModule && InitModule.bind(null);
