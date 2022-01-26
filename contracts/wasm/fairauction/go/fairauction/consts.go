// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "fairauction"
	ScDescription = "Decentralized auction to securely sell tokens to the highest bidder"
	HScName       = wasmtypes.ScHname(0x1b5c43b1)
)

const (
	ParamColor       = "color"
	ParamDescription = "description"
	ParamDuration    = "duration"
	ParamMinimumBid  = "minimumBid"
	ParamOwnerMargin = "ownerMargin"
)

const (
	ResultBidders       = "bidders"
	ResultColor         = "color"
	ResultCreator       = "creator"
	ResultDeposit       = "deposit"
	ResultDescription   = "description"
	ResultDuration      = "duration"
	ResultHighestBid    = "highestBid"
	ResultHighestBidder = "highestBidder"
	ResultMinimumBid    = "minimumBid"
	ResultNumTokens     = "numTokens"
	ResultOwnerMargin   = "ownerMargin"
	ResultWhenStarted   = "whenStarted"
)

const (
	StateAuctions    = "auctions"
	StateBidderList  = "bidderList"
	StateBids        = "bids"
	StateOwnerMargin = "ownerMargin"
)

const (
	FuncFinalizeAuction = "finalizeAuction"
	FuncPlaceBid        = "placeBid"
	FuncSetOwnerMargin  = "setOwnerMargin"
	FuncStartAuction    = "startAuction"
	ViewGetInfo         = "getInfo"
)

const (
	HFuncFinalizeAuction = wasmtypes.ScHname(0x8d534ddc)
	HFuncPlaceBid        = wasmtypes.ScHname(0x9bd72fa9)
	HFuncSetOwnerMargin  = wasmtypes.ScHname(0x1774461a)
	HFuncStartAuction    = wasmtypes.ScHname(0xd5b7bacb)
	HViewGetInfo         = wasmtypes.ScHname(0xcfedba5f)
)
