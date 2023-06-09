import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAddLiquidity } from "./types/coswap/swap/tx";
import { MsgSwapCoin } from "./types/coswap/swap/tx";
import { MsgDistributeCoins } from "./types/coswap/swap/tx";
import { MsgMintCoins } from "./types/coswap/swap/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/coswap.swap.MsgAddLiquidity", MsgAddLiquidity],
    ["/coswap.swap.MsgSwapCoin", MsgSwapCoin],
    ["/coswap.swap.MsgDistributeCoins", MsgDistributeCoins],
    ["/coswap.swap.MsgMintCoins", MsgMintCoins],
    
];

export { msgTypes }