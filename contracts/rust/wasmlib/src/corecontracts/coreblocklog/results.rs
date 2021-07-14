// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::corecontracts::coreblocklog::*;
use crate::host::*;

#[derive(Clone, Copy)]
pub struct ImmutableControlAddressesResults {
    pub(crate) id: i32,
}

impl ImmutableControlAddressesResults {
    pub fn block_index(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, RESULT_BLOCK_INDEX.get_key_id())
    }

    pub fn governing_address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.id, RESULT_GOVERNING_ADDRESS.get_key_id())
    }

    pub fn state_controller_address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.id, RESULT_STATE_CONTROLLER_ADDRESS.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableControlAddressesResults {
    pub(crate) id: i32,
}

impl MutableControlAddressesResults {
    pub fn block_index(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, RESULT_BLOCK_INDEX.get_key_id())
    }

    pub fn governing_address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.id, RESULT_GOVERNING_ADDRESS.get_key_id())
    }

    pub fn state_controller_address(&self) -> ScMutableAddress {
        ScMutableAddress::new(self.id, RESULT_STATE_CONTROLLER_ADDRESS.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetBlockInfoResults {
    pub(crate) id: i32,
}

impl ImmutableGetBlockInfoResults {
    pub fn block_info(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.id, RESULT_BLOCK_INFO.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetBlockInfoResults {
    pub(crate) id: i32,
}

impl MutableGetBlockInfoResults {
    pub fn block_info(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.id, RESULT_BLOCK_INFO.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetLatestBlockInfoResults {
    pub(crate) id: i32,
}

impl ImmutableGetLatestBlockInfoResults {
    pub fn block_info(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.id, RESULT_BLOCK_INFO.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetLatestBlockInfoResults {
    pub(crate) id: i32,
}

impl MutableGetLatestBlockInfoResults {
    pub fn block_info(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.id, RESULT_BLOCK_INFO.get_key_id())
    }
}

pub struct ArrayOfImmutableBytes {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableBytes {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bytes(&self, index: i32) -> ScImmutableBytes {
        ScImmutableBytes::new(self.obj_id, Key32(index))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetRequestIDsForBlockResults {
    pub(crate) id: i32,
}

impl ImmutableGetRequestIDsForBlockResults {
    pub fn request_id(&self) -> ArrayOfImmutableBytes {
        let arr_id = get_object_id(self.id, RESULT_REQUEST_ID.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBytes { obj_id: arr_id }
    }
}

pub struct ArrayOfMutableBytes {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableBytes {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bytes(&self, index: i32) -> ScMutableBytes {
        ScMutableBytes::new(self.obj_id, Key32(index))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetRequestIDsForBlockResults {
    pub(crate) id: i32,
}

impl MutableGetRequestIDsForBlockResults {
    pub fn request_id(&self) -> ArrayOfMutableBytes {
        let arr_id = get_object_id(self.id, RESULT_REQUEST_ID.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBytes { obj_id: arr_id }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetRequestLogRecordResults {
    pub(crate) id: i32,
}

impl ImmutableGetRequestLogRecordResults {
    pub fn block_index(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, RESULT_BLOCK_INDEX.get_key_id())
    }

    pub fn request_index(&self) -> ScImmutableInt16 {
        ScImmutableInt16::new(self.id, RESULT_REQUEST_INDEX.get_key_id())
    }

    pub fn request_record(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.id, RESULT_REQUEST_RECORD.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetRequestLogRecordResults {
    pub(crate) id: i32,
}

impl MutableGetRequestLogRecordResults {
    pub fn block_index(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, RESULT_BLOCK_INDEX.get_key_id())
    }

    pub fn request_index(&self) -> ScMutableInt16 {
        ScMutableInt16::new(self.id, RESULT_REQUEST_INDEX.get_key_id())
    }

    pub fn request_record(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.id, RESULT_REQUEST_RECORD.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetRequestLogRecordsForBlockResults {
    pub(crate) id: i32,
}

impl ImmutableGetRequestLogRecordsForBlockResults {
    pub fn request_record(&self) -> ArrayOfImmutableBytes {
        let arr_id = get_object_id(self.id, RESULT_REQUEST_RECORD.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBytes { obj_id: arr_id }
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetRequestLogRecordsForBlockResults {
    pub(crate) id: i32,
}

impl MutableGetRequestLogRecordsForBlockResults {
    pub fn request_record(&self) -> ArrayOfMutableBytes {
        let arr_id = get_object_id(self.id, RESULT_REQUEST_RECORD.get_key_id(), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBytes { obj_id: arr_id }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableIsRequestProcessedResults {
    pub(crate) id: i32,
}

impl ImmutableIsRequestProcessedResults {
    pub fn request_processed(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, RESULT_REQUEST_PROCESSED.get_key_id())
    }
}

#[derive(Clone, Copy)]
pub struct MutableIsRequestProcessedResults {
    pub(crate) id: i32,
}

impl MutableIsRequestProcessedResults {
    pub fn request_processed(&self) -> ScMutableString {
        ScMutableString::new(self.id, RESULT_REQUEST_PROCESSED.get_key_id())
    }
}