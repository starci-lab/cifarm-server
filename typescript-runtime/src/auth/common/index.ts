export enum Chain {
    Avalanche = "Avalanche",
}
export interface AuthRequestPayload {
    address: string;
    signature: string;
    message: string;
    chain: Chain
}