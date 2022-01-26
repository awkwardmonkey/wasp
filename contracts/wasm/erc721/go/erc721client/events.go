// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc721client

import (
	wasmclient2 "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmclient"
)

var erc721Handlers = map[string]func(*Erc721Events, []string){
	"erc721.approval":       func(evt *Erc721Events, msg []string) { evt.onErc721ApprovalThunk(msg) },
	"erc721.approvalForAll": func(evt *Erc721Events, msg []string) { evt.onErc721ApprovalForAllThunk(msg) },
	"erc721.init":           func(evt *Erc721Events, msg []string) { evt.onErc721InitThunk(msg) },
	"erc721.mint":           func(evt *Erc721Events, msg []string) { evt.onErc721MintThunk(msg) },
	"erc721.transfer":       func(evt *Erc721Events, msg []string) { evt.onErc721TransferThunk(msg) },
}

type Erc721Events struct {
	approval       func(e *EventApproval)
	approvalForAll func(e *EventApprovalForAll)
	init           func(e *EventInit)
	mint           func(e *EventMint)
	transfer       func(e *EventTransfer)
}

func (h *Erc721Events) CallHandler(topic string, params []string) {
	handler := erc721Handlers[topic]
	if handler != nil {
		handler(h, params)
	}
}

func (h *Erc721Events) OnErc721Approval(handler func(e *EventApproval)) {
	h.approval = handler
}

func (h *Erc721Events) OnErc721ApprovalForAll(handler func(e *EventApprovalForAll)) {
	h.approvalForAll = handler
}

func (h *Erc721Events) OnErc721Init(handler func(e *EventInit)) {
	h.init = handler
}

func (h *Erc721Events) OnErc721Mint(handler func(e *EventMint)) {
	h.mint = handler
}

func (h *Erc721Events) OnErc721Transfer(handler func(e *EventTransfer)) {
	h.transfer = handler
}

type EventApproval struct {
	wasmclient2.Event
	Approved wasmclient2.AgentID
	Owner    wasmclient2.AgentID
	TokenID  wasmclient2.Hash
}

func (h *Erc721Events) onErc721ApprovalThunk(message []string) {
	if h.approval == nil {
		return
	}
	e := &EventApproval{}
	e.Init(message)
	e.Approved = e.NextAgentID()
	e.Owner = e.NextAgentID()
	e.TokenID = e.NextHash()
	h.approval(e)
}

type EventApprovalForAll struct {
	wasmclient2.Event
	Approval bool
	Operator wasmclient2.AgentID
	Owner    wasmclient2.AgentID
}

func (h *Erc721Events) onErc721ApprovalForAllThunk(message []string) {
	if h.approvalForAll == nil {
		return
	}
	e := &EventApprovalForAll{}
	e.Init(message)
	e.Approval = e.NextBool()
	e.Operator = e.NextAgentID()
	e.Owner = e.NextAgentID()
	h.approvalForAll(e)
}

type EventInit struct {
	wasmclient2.Event
	Name   string
	Symbol string
}

func (h *Erc721Events) onErc721InitThunk(message []string) {
	if h.init == nil {
		return
	}
	e := &EventInit{}
	e.Init(message)
	e.Name = e.NextString()
	e.Symbol = e.NextString()
	h.init(e)
}

type EventMint struct {
	wasmclient2.Event
	Balance uint64
	Owner   wasmclient2.AgentID
	TokenID wasmclient2.Hash
}

func (h *Erc721Events) onErc721MintThunk(message []string) {
	if h.mint == nil {
		return
	}
	e := &EventMint{}
	e.Init(message)
	e.Balance = e.NextUint64()
	e.Owner = e.NextAgentID()
	e.TokenID = e.NextHash()
	h.mint(e)
}

type EventTransfer struct {
	wasmclient2.Event
	From    wasmclient2.AgentID
	To      wasmclient2.AgentID
	TokenID wasmclient2.Hash
}

func (h *Erc721Events) onErc721TransferThunk(message []string) {
	if h.transfer == nil {
		return
	}
	e := &EventTransfer{}
	e.Init(message)
	e.From = e.NextAgentID()
	e.To = e.NextAgentID()
	e.TokenID = e.NextHash()
	h.transfer(e)
}
