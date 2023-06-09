/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "coswap.swap";

export interface MsgMintCoins {
  admin: string;
  amount: Coin[];
}

export interface MsgMintCoinsResponse {
}

export interface MsgDistributeCoins {
  admin: string;
  toAddress: string;
  amount: Coin[];
}

export interface MsgDistributeCoinsResponse {
}

export interface MsgAddLiquidity {
  creator: string;
  tokenA: Coin | undefined;
  tokenB: Coin | undefined;
}

export interface MsgAddLiquidityResponse {
}

export interface MsgSwapCoin {
  creator: string;
  amountIn: Coin | undefined;
  amountOutMin: Coin | undefined;
}

export interface MsgSwapCoinResponse {
}

function createBaseMsgMintCoins(): MsgMintCoins {
  return { admin: "", amount: [] };
}

export const MsgMintCoins = {
  encode(message: MsgMintCoins, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.admin !== "") {
      writer.uint32(10).string(message.admin);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMintCoins {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMintCoins();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.admin = reader.string();
          break;
        case 2:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMintCoins {
    return {
      admin: isSet(object.admin) ? String(object.admin) : "",
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgMintCoins): unknown {
    const obj: any = {};
    message.admin !== undefined && (obj.admin = message.admin);
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMintCoins>, I>>(object: I): MsgMintCoins {
    const message = createBaseMsgMintCoins();
    message.admin = object.admin ?? "";
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgMintCoinsResponse(): MsgMintCoinsResponse {
  return {};
}

export const MsgMintCoinsResponse = {
  encode(_: MsgMintCoinsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgMintCoinsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgMintCoinsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgMintCoinsResponse {
    return {};
  },

  toJSON(_: MsgMintCoinsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgMintCoinsResponse>, I>>(_: I): MsgMintCoinsResponse {
    const message = createBaseMsgMintCoinsResponse();
    return message;
  },
};

function createBaseMsgDistributeCoins(): MsgDistributeCoins {
  return { admin: "", toAddress: "", amount: [] };
}

export const MsgDistributeCoins = {
  encode(message: MsgDistributeCoins, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.admin !== "") {
      writer.uint32(10).string(message.admin);
    }
    if (message.toAddress !== "") {
      writer.uint32(18).string(message.toAddress);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDistributeCoins {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDistributeCoins();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.admin = reader.string();
          break;
        case 2:
          message.toAddress = reader.string();
          break;
        case 3:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDistributeCoins {
    return {
      admin: isSet(object.admin) ? String(object.admin) : "",
      toAddress: isSet(object.toAddress) ? String(object.toAddress) : "",
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgDistributeCoins): unknown {
    const obj: any = {};
    message.admin !== undefined && (obj.admin = message.admin);
    message.toAddress !== undefined && (obj.toAddress = message.toAddress);
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDistributeCoins>, I>>(object: I): MsgDistributeCoins {
    const message = createBaseMsgDistributeCoins();
    message.admin = object.admin ?? "";
    message.toAddress = object.toAddress ?? "";
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgDistributeCoinsResponse(): MsgDistributeCoinsResponse {
  return {};
}

export const MsgDistributeCoinsResponse = {
  encode(_: MsgDistributeCoinsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDistributeCoinsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDistributeCoinsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDistributeCoinsResponse {
    return {};
  },

  toJSON(_: MsgDistributeCoinsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDistributeCoinsResponse>, I>>(_: I): MsgDistributeCoinsResponse {
    const message = createBaseMsgDistributeCoinsResponse();
    return message;
  },
};

function createBaseMsgAddLiquidity(): MsgAddLiquidity {
  return { creator: "", tokenA: undefined, tokenB: undefined };
}

export const MsgAddLiquidity = {
  encode(message: MsgAddLiquidity, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenA !== undefined) {
      Coin.encode(message.tokenA, writer.uint32(18).fork()).ldelim();
    }
    if (message.tokenB !== undefined) {
      Coin.encode(message.tokenB, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddLiquidity {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddLiquidity();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenA = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.tokenB = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddLiquidity {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      tokenA: isSet(object.tokenA) ? Coin.fromJSON(object.tokenA) : undefined,
      tokenB: isSet(object.tokenB) ? Coin.fromJSON(object.tokenB) : undefined,
    };
  },

  toJSON(message: MsgAddLiquidity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenA !== undefined && (obj.tokenA = message.tokenA ? Coin.toJSON(message.tokenA) : undefined);
    message.tokenB !== undefined && (obj.tokenB = message.tokenB ? Coin.toJSON(message.tokenB) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddLiquidity>, I>>(object: I): MsgAddLiquidity {
    const message = createBaseMsgAddLiquidity();
    message.creator = object.creator ?? "";
    message.tokenA = (object.tokenA !== undefined && object.tokenA !== null)
      ? Coin.fromPartial(object.tokenA)
      : undefined;
    message.tokenB = (object.tokenB !== undefined && object.tokenB !== null)
      ? Coin.fromPartial(object.tokenB)
      : undefined;
    return message;
  },
};

function createBaseMsgAddLiquidityResponse(): MsgAddLiquidityResponse {
  return {};
}

export const MsgAddLiquidityResponse = {
  encode(_: MsgAddLiquidityResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddLiquidityResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddLiquidityResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddLiquidityResponse {
    return {};
  },

  toJSON(_: MsgAddLiquidityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddLiquidityResponse>, I>>(_: I): MsgAddLiquidityResponse {
    const message = createBaseMsgAddLiquidityResponse();
    return message;
  },
};

function createBaseMsgSwapCoin(): MsgSwapCoin {
  return { creator: "", amountIn: undefined, amountOutMin: undefined };
}

export const MsgSwapCoin = {
  encode(message: MsgSwapCoin, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amountIn !== undefined) {
      Coin.encode(message.amountIn, writer.uint32(18).fork()).ldelim();
    }
    if (message.amountOutMin !== undefined) {
      Coin.encode(message.amountOutMin, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSwapCoin {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSwapCoin();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amountIn = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.amountOutMin = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwapCoin {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amountIn: isSet(object.amountIn) ? Coin.fromJSON(object.amountIn) : undefined,
      amountOutMin: isSet(object.amountOutMin) ? Coin.fromJSON(object.amountOutMin) : undefined,
    };
  },

  toJSON(message: MsgSwapCoin): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn ? Coin.toJSON(message.amountIn) : undefined);
    message.amountOutMin !== undefined
      && (obj.amountOutMin = message.amountOutMin ? Coin.toJSON(message.amountOutMin) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSwapCoin>, I>>(object: I): MsgSwapCoin {
    const message = createBaseMsgSwapCoin();
    message.creator = object.creator ?? "";
    message.amountIn = (object.amountIn !== undefined && object.amountIn !== null)
      ? Coin.fromPartial(object.amountIn)
      : undefined;
    message.amountOutMin = (object.amountOutMin !== undefined && object.amountOutMin !== null)
      ? Coin.fromPartial(object.amountOutMin)
      : undefined;
    return message;
  },
};

function createBaseMsgSwapCoinResponse(): MsgSwapCoinResponse {
  return {};
}

export const MsgSwapCoinResponse = {
  encode(_: MsgSwapCoinResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSwapCoinResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSwapCoinResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSwapCoinResponse {
    return {};
  },

  toJSON(_: MsgSwapCoinResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSwapCoinResponse>, I>>(_: I): MsgSwapCoinResponse {
    const message = createBaseMsgSwapCoinResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  MintCoins(request: MsgMintCoins): Promise<MsgMintCoinsResponse>;
  DistributeCoins(request: MsgDistributeCoins): Promise<MsgDistributeCoinsResponse>;
  AddLiquidity(request: MsgAddLiquidity): Promise<MsgAddLiquidityResponse>;
  SwapCoin(request: MsgSwapCoin): Promise<MsgSwapCoinResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.MintCoins = this.MintCoins.bind(this);
    this.DistributeCoins = this.DistributeCoins.bind(this);
    this.AddLiquidity = this.AddLiquidity.bind(this);
    this.SwapCoin = this.SwapCoin.bind(this);
  }
  MintCoins(request: MsgMintCoins): Promise<MsgMintCoinsResponse> {
    const data = MsgMintCoins.encode(request).finish();
    const promise = this.rpc.request("coswap.swap.Msg", "MintCoins", data);
    return promise.then((data) => MsgMintCoinsResponse.decode(new _m0.Reader(data)));
  }

  DistributeCoins(request: MsgDistributeCoins): Promise<MsgDistributeCoinsResponse> {
    const data = MsgDistributeCoins.encode(request).finish();
    const promise = this.rpc.request("coswap.swap.Msg", "DistributeCoins", data);
    return promise.then((data) => MsgDistributeCoinsResponse.decode(new _m0.Reader(data)));
  }

  AddLiquidity(request: MsgAddLiquidity): Promise<MsgAddLiquidityResponse> {
    const data = MsgAddLiquidity.encode(request).finish();
    const promise = this.rpc.request("coswap.swap.Msg", "AddLiquidity", data);
    return promise.then((data) => MsgAddLiquidityResponse.decode(new _m0.Reader(data)));
  }

  SwapCoin(request: MsgSwapCoin): Promise<MsgSwapCoinResponse> {
    const data = MsgSwapCoin.encode(request).finish();
    const promise = this.rpc.request("coswap.swap.Msg", "SwapCoin", data);
    return promise.then((data) => MsgSwapCoinResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
