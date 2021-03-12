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
import time
import logging
import uuid

import six

import grpc
from grpc._cython import cygrpc

from Store import Store
import message_pb2

s = Store()

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


def serialize(message, serializer):
    # 3mb start
    if hasattr(message, 'FI_Trace'):
        if message.FI_Trace:
            print("serialize records before extend message FI_trace: ", message.FI_Trace.records)

            name = ""
            if message.FI_Type == 1:
                name = "Message_Request"
            elif message.FI_Type == 2:
                name = "Message_Response"

            record = message_pb2.Record(type=1, service=str(uuid.uuid4()), message_name=name,
                                        timestamp=int(time.time() * 1e9),
                                        uuid=message.FI_Trace.records[0].uuid)

            message.FI_Trace.records.extend([record])

            print("serialize records after extend: ", message.FI_Trace.records)

            # add to store
            s.SetTrace(str(message.FI_Trace.id), message.FI_Trace)

            #if name == "Message_Response":
                # delete trace
                #s.FetchTrace(message.FI_Trace.id)
            # else:
                # SimulateFault(t)

            # get from store and update current message FI_Trace
            tmp = s.GetTrace(str(message.FI_Trace.id))
            if tmp.records:
                for i in range(len(tmp.records)):
                    record = tmp.records[i]
                    message.FI_Trace.records.extend([record])
            if tmp.rlfis:
                for i in range(len(tmp.rlfis)):
                    rlfi = tmp.rlfis[i]
                    message.FI_Trace.rlfis.extend([rlfi])
            if tmp.tfis:
                for i in range(len(tmp.tfis)):
                    tfi = tmp.tfis[i]
                    message.FI_Trace.tfis.extend([tfi])

            print("SERIALIZE Global Store:", s.PrintStore())
    # 3mb end

    return _transform(message, serializer, 'Exception serializing message!')


def deserialize(serialized_message, deserializer):
    m = _transform(serialized_message, deserializer, 'Exception deserializing message!')

    # 3mb start
    if hasattr(m, 'FI_Trace'):
        if m.FI_Trace:
            print("deserialize records before extend message FI_trace: ", m.FI_Trace.records)

            name = ""
            if m.FI_Type == 1:
                name = "Message_Request"
            elif m.FI_Type == 2:
                name = "Message_Response"

            record = message_pb2.Record(type=2, service=str(uuid.uuid4()), message_name=name,
                                        timestamp=int(time.time() * 1e9),
                                        uuid=m.FI_Trace.records[len(m.FI_Trace.records) - 1].uuid)

            m.FI_Trace.records.extend([record])

            print("deserialize records after extend: ", m.FI_Trace.records)

            # add to store
            s.SetTrace(str(m.FI_Trace.id), m.FI_Trace)

            # get from store and update current message FI_Trace
            tmp = s.GetTrace(str(m.FI_Trace.id))
            if tmp.records:
                for i in range(len(tmp.records)):
                    record = tmp.records[i]
                    m.FI_Trace.records.extend([record])
            if tmp.rlfis:
                for i in range(len(tmp.rlfis)):
                    rlfi = tmp.rlfis[i]
                    m.FI_Trace.rlfis.extend([rlfi])
            if tmp.tfis:
                for i in range(len(tmp.tfis)):
                    tfi = tmp.tfis[i]
                    m.FI_Trace.tfis.extend([tfi])

           print("DESERIALIZE Global Store:", s.PrintStore())
    # 3mb end

    return m


def fully_qualified_method(group, method):
    return '/{}/{}'.format(group, method)
