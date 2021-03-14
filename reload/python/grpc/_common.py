# Copyright 2016 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""Shared implementation."""
import sys
import time
import logging
import uuid

import six

import grpc
from grpc._cython import cygrpc

from Store import Store
import message_pb2

serviceUUID = str(uuid.uuid4())
s = Store()
TMB_CTXIDKEY = "traceID"

_LOGGER = logging.getLogger(__name__)

CYGRPC_CONNECTIVITY_STATE_TO_CHANNEL_CONNECTIVITY = {
    cygrpc.ConnectivityState.idle:
    grpc.ChannelConnectivity.IDLE,
    cygrpc.ConnectivityState.connecting:
    grpc.ChannelConnectivity.CONNECTING,
    cygrpc.ConnectivityState.ready:
    grpc.ChannelConnectivity.READY,
    cygrpc.ConnectivityState.transient_failure:
    grpc.ChannelConnectivity.TRANSIENT_FAILURE,
    cygrpc.ConnectivityState.shutdown:
    grpc.ChannelConnectivity.SHUTDOWN,
}

CYGRPC_STATUS_CODE_TO_STATUS_CODE = {
    cygrpc.StatusCode.ok: grpc.StatusCode.OK,
    cygrpc.StatusCode.cancelled: grpc.StatusCode.CANCELLED,
    cygrpc.StatusCode.unknown: grpc.StatusCode.UNKNOWN,
    cygrpc.StatusCode.invalid_argument: grpc.StatusCode.INVALID_ARGUMENT,
    cygrpc.StatusCode.deadline_exceeded: grpc.StatusCode.DEADLINE_EXCEEDED,
    cygrpc.StatusCode.not_found: grpc.StatusCode.NOT_FOUND,
    cygrpc.StatusCode.already_exists: grpc.StatusCode.ALREADY_EXISTS,
    cygrpc.StatusCode.permission_denied: grpc.StatusCode.PERMISSION_DENIED,
    cygrpc.StatusCode.unauthenticated: grpc.StatusCode.UNAUTHENTICATED,
    cygrpc.StatusCode.resource_exhausted: grpc.StatusCode.RESOURCE_EXHAUSTED,
    cygrpc.StatusCode.failed_precondition: grpc.StatusCode.FAILED_PRECONDITION,
    cygrpc.StatusCode.aborted: grpc.StatusCode.ABORTED,
    cygrpc.StatusCode.out_of_range: grpc.StatusCode.OUT_OF_RANGE,
    cygrpc.StatusCode.unimplemented: grpc.StatusCode.UNIMPLEMENTED,
    cygrpc.StatusCode.internal: grpc.StatusCode.INTERNAL,
    cygrpc.StatusCode.unavailable: grpc.StatusCode.UNAVAILABLE,
    cygrpc.StatusCode.data_loss: grpc.StatusCode.DATA_LOSS,
}
STATUS_CODE_TO_CYGRPC_STATUS_CODE = {
    grpc_code: cygrpc_code
    for cygrpc_code, grpc_code in six.iteritems(
        CYGRPC_STATUS_CODE_TO_STATUS_CODE)
}


def encode(s):
    if isinstance(s, bytes):
        return s
    else:
        return s.encode('ascii')


def decode(b):
    if isinstance(b, str):
        return b
    else:
        try:
            return b.decode('utf8')
        except UnicodeDecodeError:
            _LOGGER.exception('Invalid encoding on %s', b)
            return b.decode('latin1')


def _transform(message, transformer, exception_message):
    if transformer is None:
        return message
    else:
        try:
            return transformer(message)
        except Exception:  # pylint: disable=broad-except
            _LOGGER.exception(exception_message)
            return None


# if we recieve a message we call serialize
# check whether context has valid trace id
# context: id -> use this id to get trace
# check for duplicate records
def serialize(message, serializer, Context = None):
    # 3mb start
    if hasattr(message, 'FI_Trace'):
        # default function has key -> do we have field TraceID added by deserialize / context propagation
        if Context is not None and Context.has_key(TMB_CTXIDKEY):
            # meta should have uuid and traceID
            meta = Context[TMB_CTXIDKEY]
            if meta:
                if s.CheckByContextMeta(meta):
                    # tmp = s.GetByContextMeta(meta)
                    Serialize_Uuid = ""
                    name = ""

                    if message.FI_Type == 1:
                        name = "Message_Request"
                        Serialize_Uuid = str(uuid.uuid4())
                    elif message.FI_Type == 2:
                        name = "Message_Response"
                        Serialize_Uuid = meta.uuid

                    # TODO: (if needed) if two records are same dont extend that record
                    record = message_pb2.Record(type=1, message_name=name, timestamp=int(time.time() * 1e9),
                                                uuid=Serialize_Uuid, service=serviceUUID)

                    def updateFunction(trace):
                        trace.records.extend([record])

                    trace, ok = Store.UpdateFunctionByContextMeta(meta, updateFunction)
                    if ok:
                        # TODO: fault injection
                        if name == "Message_Request":
                            # fault injection
                            pass
                        elif name == "Message_Response":
                            s.DeleteByContextMeta(meta)

                        # Set FI trace
                        message.FI_Trace = trace

                        print("SERIALIZE Global Store:", s.GetStore())

    # 3mb end

    return _transform(message, serializer, 'Exception serializing message!')

# assigns context id to context
# check for duplicate records
def deserialize(serialized_message, deserializer, Context=None):
    m = _transform(serialized_message, deserializer, 'Exception deserializing message!')

    # 3mb start
    if hasattr(m, 'FI_Trace'):
        # default function has key -> do we have field TraceID added by deserialize / context propagation
        if Context is not None and Context.has_key(TMB_CTXIDKEY):
            # meta should have uuid and traceID
            meta = Context[TMB_CTXIDKEY]
            if meta:
                if s.CheckByContextMeta(meta):
                    # tmp = s.GetByContextMeta(meta)
                    name = ""

                    if m.FI_Type == 1:
                        name = "Message_Request"
                    elif m.FI_Type == 2:
                        name = "Message_Response"

                    trace = m.FI_Trace
                    if trace is not None:
                        # TODO: fault injection

                        if name == "Message_Request":
                            print("[RELOAD] Unmarshal, receive request: ", name)
                            if len(trace.records) != 1:
                                print("[RELOAD] Unmarshal, receive invalid trace:", trace)
                            elif trace.records[0].uuid == "":
                                print("[RELOAD] Unmarshal, receive invalid uuid")
                            else:
                                meta.TMB_CTXIDKEY = trace.id
                                meta.uuid = trace.records[0].uuid

                                record = message_pb2.Record(type=2, message_name=name, timestamp=int(time.time() * 1e9),
                                                            uuid=trace.records[0].uuid, service=serviceUUID)

                                s.SetByContextMeta(meta, trace)

                        elif name == "Message_Response":
                            if len(trace.records) == 0:
                                print("[RELOAD] Unmarshal, receive empty trace")
                            elif trace.records[0].uuid == "":
                                print("[RELOAD] Unmarshal, receive invalid uuid: ", uuid)
                            else:
                                #def oldFunction(oldTrace):

                                m.FI_Trace = None


                        print("DESERIALIZE Global Store:", s.GetStore())

    # 3mb end

    return m


def fully_qualified_method(group, method):
    return '/{}/{}'.format(group, method)
