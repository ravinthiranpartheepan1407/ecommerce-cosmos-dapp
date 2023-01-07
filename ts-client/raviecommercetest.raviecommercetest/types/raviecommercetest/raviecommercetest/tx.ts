/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "raviecommercetest.raviecommercetest";

export interface MsgProductImage {
  creator: string;
  productPrice: string;
}

export interface MsgProductImageResponse {
}

function createBaseMsgProductImage(): MsgProductImage {
  return { creator: "", productPrice: "" };
}

export const MsgProductImage = {
  encode(message: MsgProductImage, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.productPrice !== "") {
      writer.uint32(18).string(message.productPrice);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProductImage {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProductImage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.productPrice = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgProductImage {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      productPrice: isSet(object.productPrice) ? String(object.productPrice) : "",
    };
  },

  toJSON(message: MsgProductImage): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.productPrice !== undefined && (obj.productPrice = message.productPrice);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgProductImage>, I>>(object: I): MsgProductImage {
    const message = createBaseMsgProductImage();
    message.creator = object.creator ?? "";
    message.productPrice = object.productPrice ?? "";
    return message;
  },
};

function createBaseMsgProductImageResponse(): MsgProductImageResponse {
  return {};
}

export const MsgProductImageResponse = {
  encode(_: MsgProductImageResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProductImageResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProductImageResponse();
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

  fromJSON(_: any): MsgProductImageResponse {
    return {};
  },

  toJSON(_: MsgProductImageResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgProductImageResponse>, I>>(_: I): MsgProductImageResponse {
    const message = createBaseMsgProductImageResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ProductImage(request: MsgProductImage): Promise<MsgProductImageResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.ProductImage = this.ProductImage.bind(this);
  }
  ProductImage(request: MsgProductImage): Promise<MsgProductImageResponse> {
    const data = MsgProductImage.encode(request).finish();
    const promise = this.rpc.request("raviecommercetest.raviecommercetest.Msg", "ProductImage", data);
    return promise.then((data) => MsgProductImageResponse.decode(new _m0.Reader(data)));
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
