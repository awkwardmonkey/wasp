// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ControlAddressesCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewControlAddresses);
	results: sc.ImmutableControlAddressesResults = new sc.ImmutableControlAddressesResults(wasmlib.ScView.nilProxy);
}

export class GetBlockInfoCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetBlockInfo);
	params: sc.MutableGetBlockInfoParams = new sc.MutableGetBlockInfoParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetBlockInfoResults = new sc.ImmutableGetBlockInfoResults(wasmlib.ScView.nilProxy);
}

export class GetEventsForBlockCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetEventsForBlock);
	params: sc.MutableGetEventsForBlockParams = new sc.MutableGetEventsForBlockParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetEventsForBlockResults = new sc.ImmutableGetEventsForBlockResults(wasmlib.ScView.nilProxy);
}

export class GetEventsForContractCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetEventsForContract);
	params: sc.MutableGetEventsForContractParams = new sc.MutableGetEventsForContractParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetEventsForContractResults = new sc.ImmutableGetEventsForContractResults(wasmlib.ScView.nilProxy);
}

export class GetEventsForRequestCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetEventsForRequest);
	params: sc.MutableGetEventsForRequestParams = new sc.MutableGetEventsForRequestParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetEventsForRequestResults = new sc.ImmutableGetEventsForRequestResults(wasmlib.ScView.nilProxy);
}

export class GetLatestBlockInfoCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetLatestBlockInfo);
	results: sc.ImmutableGetLatestBlockInfoResults = new sc.ImmutableGetLatestBlockInfoResults(wasmlib.ScView.nilProxy);
}

export class GetRequestIDsForBlockCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetRequestIDsForBlock);
	params: sc.MutableGetRequestIDsForBlockParams = new sc.MutableGetRequestIDsForBlockParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetRequestIDsForBlockResults = new sc.ImmutableGetRequestIDsForBlockResults(wasmlib.ScView.nilProxy);
}

export class GetRequestReceiptCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetRequestReceipt);
	params: sc.MutableGetRequestReceiptParams = new sc.MutableGetRequestReceiptParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetRequestReceiptResults = new sc.ImmutableGetRequestReceiptResults(wasmlib.ScView.nilProxy);
}

export class GetRequestReceiptsForBlockCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetRequestReceiptsForBlock);
	params: sc.MutableGetRequestReceiptsForBlockParams = new sc.MutableGetRequestReceiptsForBlockParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetRequestReceiptsForBlockResults = new sc.ImmutableGetRequestReceiptsForBlockResults(wasmlib.ScView.nilProxy);
}

export class IsRequestProcessedCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewIsRequestProcessed);
	params: sc.MutableIsRequestProcessedParams = new sc.MutableIsRequestProcessedParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableIsRequestProcessedResults = new sc.ImmutableIsRequestProcessedResults(wasmlib.ScView.nilProxy);
}

export class ScFuncs {
    static controlAddresses(_ctx: wasmlib.ScViewCallContext): ControlAddressesCall {
        const f = new ControlAddressesCall();
		f.results = new sc.ImmutableControlAddressesResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getBlockInfo(_ctx: wasmlib.ScViewCallContext): GetBlockInfoCall {
        const f = new GetBlockInfoCall();
		f.params = new sc.MutableGetBlockInfoParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetBlockInfoResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getEventsForBlock(_ctx: wasmlib.ScViewCallContext): GetEventsForBlockCall {
        const f = new GetEventsForBlockCall();
		f.params = new sc.MutableGetEventsForBlockParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetEventsForBlockResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getEventsForContract(_ctx: wasmlib.ScViewCallContext): GetEventsForContractCall {
        const f = new GetEventsForContractCall();
		f.params = new sc.MutableGetEventsForContractParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetEventsForContractResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getEventsForRequest(_ctx: wasmlib.ScViewCallContext): GetEventsForRequestCall {
        const f = new GetEventsForRequestCall();
		f.params = new sc.MutableGetEventsForRequestParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetEventsForRequestResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getLatestBlockInfo(_ctx: wasmlib.ScViewCallContext): GetLatestBlockInfoCall {
        const f = new GetLatestBlockInfoCall();
		f.results = new sc.ImmutableGetLatestBlockInfoResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getRequestIDsForBlock(_ctx: wasmlib.ScViewCallContext): GetRequestIDsForBlockCall {
        const f = new GetRequestIDsForBlockCall();
		f.params = new sc.MutableGetRequestIDsForBlockParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetRequestIDsForBlockResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getRequestReceipt(_ctx: wasmlib.ScViewCallContext): GetRequestReceiptCall {
        const f = new GetRequestReceiptCall();
		f.params = new sc.MutableGetRequestReceiptParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetRequestReceiptResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getRequestReceiptsForBlock(_ctx: wasmlib.ScViewCallContext): GetRequestReceiptsForBlockCall {
        const f = new GetRequestReceiptsForBlockCall();
		f.params = new sc.MutableGetRequestReceiptsForBlockParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetRequestReceiptsForBlockResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static isRequestProcessed(_ctx: wasmlib.ScViewCallContext): IsRequestProcessedCall {
        const f = new IsRequestProcessedCall();
		f.params = new sc.MutableIsRequestProcessedParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableIsRequestProcessedResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }
}