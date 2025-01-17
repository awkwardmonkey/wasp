// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ImmutableIncCounterState extends wasmlib.ScMapID {
    counter(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.StateCounter));
	}

    numRepeats(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, wasmlib.Key32.fromString(sc.StateNumRepeats));
	}
}

export class MutableIncCounterState extends wasmlib.ScMapID {
    asImmutable(): sc.ImmutableIncCounterState {
		const imm = new sc.ImmutableIncCounterState();
		imm.mapID = this.mapID;
		return imm;
	}

    counter(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.StateCounter));
	}

    numRepeats(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, wasmlib.Key32.fromString(sc.StateNumRepeats));
	}
}
