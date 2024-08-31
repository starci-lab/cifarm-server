import { AuthRequestPayload, Chain } from "../common"
import { verifyMessage } from "ethers"

export const BeforeAvalancheAuthenticate: nkruntime.BeforeHookFunction<nkruntime.AuthenticateCustomRequest> = function (ctx: nkruntime.Context, logger: nkruntime.Logger, nk: nkruntime.Nakama, data: nkruntime.AuthenticateCustomRequest): nkruntime.AuthenticateCustomRequest | void {
    const { chain, message, address, signature } = JSON.parse(data.account.id) as AuthRequestPayload;
    if (chain !== Chain.Avalanche) {
        return data;
    }
    const verified = verifyMessage(message, signature) === address;
    if (!verified) {
        logger.error("Signature verification failed");
        return null;
    } else {
        return {
            ...data,
            account: {
                id: address,
            }
        }
    }
  };