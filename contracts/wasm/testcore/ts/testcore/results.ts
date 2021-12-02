// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class ImmutableCallOnChainResults extends wasmlib.ScMapID {
    intValue(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultIntValue]);
	}
}

export class MutableCallOnChainResults extends wasmlib.ScMapID {
    intValue(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultIntValue]);
	}
}

export class ImmutableGetMintedSupplyResults extends wasmlib.ScMapID {
    mintedColor(): wasmlib.ScImmutableColor {
		return new wasmlib.ScImmutableColor(this.mapID, sc.idxMap[sc.IdxResultMintedColor]);
	}

    mintedSupply(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultMintedSupply]);
	}
}

export class MutableGetMintedSupplyResults extends wasmlib.ScMapID {
    mintedColor(): wasmlib.ScMutableColor {
		return new wasmlib.ScMutableColor(this.mapID, sc.idxMap[sc.IdxResultMintedColor]);
	}

    mintedSupply(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultMintedSupply]);
	}
}

export class ImmutableRunRecursionResults extends wasmlib.ScMapID {
    intValue(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultIntValue]);
	}
}

export class MutableRunRecursionResults extends wasmlib.ScMapID {
    intValue(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultIntValue]);
	}
}

export class ImmutableTestChainOwnerIDFullResults extends wasmlib.ScMapID {
    chainOwnerID(): wasmlib.ScImmutableAgentID {
		return new wasmlib.ScImmutableAgentID(this.mapID, sc.idxMap[sc.IdxResultChainOwnerID]);
	}
}

export class MutableTestChainOwnerIDFullResults extends wasmlib.ScMapID {
    chainOwnerID(): wasmlib.ScMutableAgentID {
		return new wasmlib.ScMutableAgentID(this.mapID, sc.idxMap[sc.IdxResultChainOwnerID]);
	}
}

export class ImmutableFibonacciResults extends wasmlib.ScMapID {
    intValue(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultIntValue]);
	}
}

export class MutableFibonacciResults extends wasmlib.ScMapID {
    intValue(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultIntValue]);
	}
}

export class ImmutableGetCounterResults extends wasmlib.ScMapID {
    counter(): wasmlib.ScImmutableInt64 {
		return new wasmlib.ScImmutableInt64(this.mapID, sc.idxMap[sc.IdxResultCounter]);
	}
}

export class MutableGetCounterResults extends wasmlib.ScMapID {
    counter(): wasmlib.ScMutableInt64 {
		return new wasmlib.ScMutableInt64(this.mapID, sc.idxMap[sc.IdxResultCounter]);
	}
}

export class MapStringToImmutableInt64 {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getInt64(key: string): wasmlib.ScImmutableInt64 {
        return new wasmlib.ScImmutableInt64(this.objID, wasmlib.Key32.fromString(key));
    }
}

export class ImmutableGetIntResults extends wasmlib.ScMapID {
    values(): sc.MapStringToImmutableInt64 {
		return new sc.MapStringToImmutableInt64(this.mapID);
	}
}

export class MapStringToMutableInt64 {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    getInt64(key: string): wasmlib.ScMutableInt64 {
        return new wasmlib.ScMutableInt64(this.objID, wasmlib.Key32.fromString(key));
    }
}

export class MutableGetIntResults extends wasmlib.ScMapID {
    values(): sc.MapStringToMutableInt64 {
		return new sc.MapStringToMutableInt64(this.mapID);
	}
}

export class MapStringToImmutableString {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getString(key: string): wasmlib.ScImmutableString {
        return new wasmlib.ScImmutableString(this.objID, wasmlib.Key32.fromString(key));
    }
}

export class ImmutableGetStringValueResults extends wasmlib.ScMapID {
    vars(): sc.MapStringToImmutableString {
		return new sc.MapStringToImmutableString(this.mapID);
	}
}

export class MapStringToMutableString {
	objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    getString(key: string): wasmlib.ScMutableString {
        return new wasmlib.ScMutableString(this.objID, wasmlib.Key32.fromString(key));
    }
}

export class MutableGetStringValueResults extends wasmlib.ScMapID {
    vars(): sc.MapStringToMutableString {
		return new sc.MapStringToMutableString(this.mapID);
	}
}

export class ImmutableTestChainOwnerIDViewResults extends wasmlib.ScMapID {
    chainOwnerID(): wasmlib.ScImmutableAgentID {
		return new wasmlib.ScImmutableAgentID(this.mapID, sc.idxMap[sc.IdxResultChainOwnerID]);
	}
}

export class MutableTestChainOwnerIDViewResults extends wasmlib.ScMapID {
    chainOwnerID(): wasmlib.ScMutableAgentID {
		return new wasmlib.ScMutableAgentID(this.mapID, sc.idxMap[sc.IdxResultChainOwnerID]);
	}
}

export class ImmutableTestSandboxCallResults extends wasmlib.ScMapID {
    sandboxCall(): wasmlib.ScImmutableString {
		return new wasmlib.ScImmutableString(this.mapID, sc.idxMap[sc.IdxResultSandboxCall]);
	}
}

export class MutableTestSandboxCallResults extends wasmlib.ScMapID {
    sandboxCall(): wasmlib.ScMutableString {
		return new wasmlib.ScMutableString(this.mapID, sc.idxMap[sc.IdxResultSandboxCall]);
	}
}
