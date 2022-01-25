// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";

export class CallRequest {
    contract : wasmlib.ScHname = new wasmlib.ScHname(0); 
    function : wasmlib.ScHname = new wasmlib.ScHname(0); 
    params   : u8[] = []; 
    transfer : u8[] = []; 

    static fromBytes(bytes: u8[]): CallRequest {
        let decode = new wasmlib.BytesDecoder(bytes);
        let data = new CallRequest();
        data.contract = decode.hname();
        data.function = decode.hname();
        data.params   = decode.bytes();
        data.transfer = decode.bytes();
        decode.close();
        return data;
    }

    bytes(): u8[] {
        return new wasmlib.BytesEncoder().
		    hname(this.contract).
		    hname(this.function).
		    bytes(this.params).
		    bytes(this.transfer).
            data();
    }
}

export class ImmutableCallRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    value(): CallRequest {
        return CallRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class MutableCallRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    delete(): void {
        wasmlib.delKey(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    setValue(value: CallRequest): void {
        wasmlib.setBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES, value.bytes());
    }

    value(): CallRequest {
        return CallRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class DeployRequest {
    description : string = ""; 
    name        : string = ""; 
    params      : u8[] = []; 
    progHash    : wasmlib.ScHash = new wasmlib.ScHash(); 

    static fromBytes(bytes: u8[]): DeployRequest {
        let decode = new wasmlib.BytesDecoder(bytes);
        let data = new DeployRequest();
        data.description = decode.string();
        data.name        = decode.string();
        data.params      = decode.bytes();
        data.progHash    = decode.hash();
        decode.close();
        return data;
    }

    bytes(): u8[] {
        return new wasmlib.BytesEncoder().
		    string(this.description).
		    string(this.name).
		    bytes(this.params).
		    hash(this.progHash).
            data();
    }
}

export class ImmutableDeployRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    value(): DeployRequest {
        return DeployRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class MutableDeployRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    delete(): void {
        wasmlib.delKey(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    setValue(value: DeployRequest): void {
        wasmlib.setBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES, value.bytes());
    }

    value(): DeployRequest {
        return DeployRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class PostRequest {
    chainID  : wasmlib.ScChainID = new wasmlib.ScChainID(); 
    contract : wasmlib.ScHname = new wasmlib.ScHname(0); 
    delay    : u32 = 0; 
    function : wasmlib.ScHname = new wasmlib.ScHname(0); 
    params   : u8[] = []; 
    transfer : u8[] = []; 

    static fromBytes(bytes: u8[]): PostRequest {
        let decode = new wasmlib.BytesDecoder(bytes);
        let data = new PostRequest();
        data.chainID  = decode.chainID();
        data.contract = decode.hname();
        data.delay    = decode.uint32();
        data.function = decode.hname();
        data.params   = decode.bytes();
        data.transfer = decode.bytes();
        decode.close();
        return data;
    }

    bytes(): u8[] {
        return new wasmlib.BytesEncoder().
		    chainID(this.chainID).
		    hname(this.contract).
		    uint32(this.delay).
		    hname(this.function).
		    bytes(this.params).
		    bytes(this.transfer).
            data();
    }
}

export class ImmutablePostRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    value(): PostRequest {
        return PostRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class MutablePostRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    delete(): void {
        wasmlib.delKey(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    setValue(value: PostRequest): void {
        wasmlib.setBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES, value.bytes());
    }

    value(): PostRequest {
        return PostRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class SendRequest {
    address  : wasmlib.ScAddress = new wasmlib.ScAddress(); 
    transfer : u8[] = []; 

    static fromBytes(bytes: u8[]): SendRequest {
        let decode = new wasmlib.BytesDecoder(bytes);
        let data = new SendRequest();
        data.address  = decode.address();
        data.transfer = decode.bytes();
        decode.close();
        return data;
    }

    bytes(): u8[] {
        return new wasmlib.BytesEncoder().
		    address(this.address).
		    bytes(this.transfer).
            data();
    }
}

export class ImmutableSendRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    value(): SendRequest {
        return SendRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}

export class MutableSendRequest {
    objID: i32;
    keyID: wasmlib.Key32;

    constructor(objID: i32, keyID: wasmlib.Key32) {
        this.objID = objID;
        this.keyID = keyID;
    }

    delete(): void {
        wasmlib.delKey(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    exists(): boolean {
        return wasmlib.exists(this.objID, this.keyID, wasmlib.TYPE_BYTES);
    }

    setValue(value: SendRequest): void {
        wasmlib.setBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES, value.bytes());
    }

    value(): SendRequest {
        return SendRequest.fromBytes(wasmlib.getBytes(this.objID, this.keyID, wasmlib.TYPE_BYTES));
    }
}