import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/dex/tx";
import { MsgUpdatePool } from "./types/dex/tx";
import { MsgDeletePool } from "./types/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/exchange.dex.MsgCreatePool", MsgCreatePool],
    ["/exchange.dex.MsgUpdatePool", MsgUpdatePool],
    ["/exchange.dex.MsgDeletePool", MsgDeletePool],
    
];

export { msgTypes }