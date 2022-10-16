import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdatePool } from "./types/exchange/tx";
import { MsgDeletePool } from "./types/exchange/tx";
import { MsgCreatePool } from "./types/exchange/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/exchange.exchange.MsgUpdatePool", MsgUpdatePool],
    ["/exchange.exchange.MsgDeletePool", MsgDeletePool],
    ["/exchange.exchange.MsgCreatePool", MsgCreatePool],
    
];

export { msgTypes }