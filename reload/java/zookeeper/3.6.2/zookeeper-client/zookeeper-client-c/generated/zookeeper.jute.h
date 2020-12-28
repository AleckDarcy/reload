/**
* Licensed to the Apache Software Foundation (ASF) under one
* or more contributor license agreements.  See the NOTICE file
* distributed with this work for additional information
* regarding copyright ownership.  The ASF licenses this file
* to you under the Apache License, Version 2.0 (the
* "License"); you may not use this file except in compliance
* with the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

#ifndef __ZOOKEEPER_JUTE__
#define __ZOOKEEPER_JUTE__
#include "recordio.h"

#ifdef __cplusplus
extern "C" {
#endif

struct TMB_Event {
    int32_t type;
    int64_t timestamp;
    char * message_name;
    char * uuid;
    char * service;
};
int serialize_TMB_Event(struct oarchive *out, const char *tag, struct TMB_Event *v);
int deserialize_TMB_Event(struct iarchive *in, const char *tag, struct TMB_Event*v);
void deallocate_TMB_Event(struct TMB_Event*);
struct TMB_TFIMeta {
    char * name;
    int64_t times;
    int64_t already;
};
int serialize_TMB_TFIMeta(struct oarchive *out, const char *tag, struct TMB_TFIMeta *v);
int deserialize_TMB_TFIMeta(struct iarchive *in, const char *tag, struct TMB_TFIMeta*v);
void deallocate_TMB_TFIMeta(struct TMB_TFIMeta*);
struct TMB_TFIMeta_vector {
    int32_t count;
    struct TMB_TFIMeta *data;

};
int serialize_TMB_TFIMeta_vector(struct oarchive *out, const char *tag, struct TMB_TFIMeta_vector *v);
int deserialize_TMB_TFIMeta_vector(struct iarchive *in, const char *tag, struct TMB_TFIMeta_vector *v);
int allocate_TMB_TFIMeta_vector(struct TMB_TFIMeta_vector *v, int32_t len);
int deallocate_TMB_TFIMeta_vector(struct TMB_TFIMeta_vector *v);
struct TMB_TFI {
    int32_t type;
    char * name;
    int64_t delay;
    struct TMB_TFIMeta_vector after;
};
int serialize_TMB_TFI(struct oarchive *out, const char *tag, struct TMB_TFI *v);
int deserialize_TMB_TFI(struct iarchive *in, const char *tag, struct TMB_TFI*v);
void deallocate_TMB_TFI(struct TMB_TFI*);
struct TMB_Event_vector {
    int32_t count;
    struct TMB_Event *data;

};
int serialize_TMB_Event_vector(struct oarchive *out, const char *tag, struct TMB_Event_vector *v);
int deserialize_TMB_Event_vector(struct iarchive *in, const char *tag, struct TMB_Event_vector *v);
int allocate_TMB_Event_vector(struct TMB_Event_vector *v, int32_t len);
int deallocate_TMB_Event_vector(struct TMB_Event_vector *v);
struct TMB_TFI_vector {
    int32_t count;
    struct TMB_TFI *data;

};
int serialize_TMB_TFI_vector(struct oarchive *out, const char *tag, struct TMB_TFI_vector *v);
int deserialize_TMB_TFI_vector(struct iarchive *in, const char *tag, struct TMB_TFI_vector *v);
int allocate_TMB_TFI_vector(struct TMB_TFI_vector *v, int32_t len);
int deallocate_TMB_TFI_vector(struct TMB_TFI_vector *v);
struct TMB_Trace {
    int64_t id;
    struct TMB_Event_vector events;
    struct TMB_TFI_vector tfis;
};
int serialize_TMB_Trace(struct oarchive *out, const char *tag, struct TMB_Trace *v);
int deserialize_TMB_Trace(struct iarchive *in, const char *tag, struct TMB_Trace*v);
void deallocate_TMB_Trace(struct TMB_Trace*);
struct Id {
    char * scheme;
    char * id;
};
int serialize_Id(struct oarchive *out, const char *tag, struct Id *v);
int deserialize_Id(struct iarchive *in, const char *tag, struct Id*v);
void deallocate_Id(struct Id*);
struct ACL {
    int32_t perms;
    struct Id id;
};
int serialize_ACL(struct oarchive *out, const char *tag, struct ACL *v);
int deserialize_ACL(struct iarchive *in, const char *tag, struct ACL*v);
void deallocate_ACL(struct ACL*);
struct Stat {
    int64_t czxid;
    int64_t mzxid;
    int64_t ctime;
    int64_t mtime;
    int32_t version;
    int32_t cversion;
    int32_t aversion;
    int64_t ephemeralOwner;
    int32_t dataLength;
    int32_t numChildren;
    int64_t pzxid;
};
int serialize_Stat(struct oarchive *out, const char *tag, struct Stat *v);
int deserialize_Stat(struct iarchive *in, const char *tag, struct Stat*v);
void deallocate_Stat(struct Stat*);
struct StatPersisted {
    int64_t czxid;
    int64_t mzxid;
    int64_t ctime;
    int64_t mtime;
    int32_t version;
    int32_t cversion;
    int32_t aversion;
    int64_t ephemeralOwner;
    int64_t pzxid;
};
int serialize_StatPersisted(struct oarchive *out, const char *tag, struct StatPersisted *v);
int deserialize_StatPersisted(struct iarchive *in, const char *tag, struct StatPersisted*v);
void deallocate_StatPersisted(struct StatPersisted*);
struct NullPointerResponse {
    char * requestName;
    struct TMB_Trace trace;
};
int serialize_NullPointerResponse(struct oarchive *out, const char *tag, struct NullPointerResponse *v);
int deserialize_NullPointerResponse(struct iarchive *in, const char *tag, struct NullPointerResponse*v);
void deallocate_NullPointerResponse(struct NullPointerResponse*);
struct ConnectRequest {
    int32_t protocolVersion;
    int64_t lastZxidSeen;
    int32_t timeOut;
    int64_t sessionId;
    struct buffer passwd;
    struct TMB_Trace trace;
};
int serialize_ConnectRequest(struct oarchive *out, const char *tag, struct ConnectRequest *v);
int deserialize_ConnectRequest(struct iarchive *in, const char *tag, struct ConnectRequest*v);
void deallocate_ConnectRequest(struct ConnectRequest*);
struct ConnectResponse {
    int32_t protocolVersion;
    int32_t timeOut;
    int64_t sessionId;
    struct buffer passwd;
    struct TMB_Trace trace;
};
int serialize_ConnectResponse(struct oarchive *out, const char *tag, struct ConnectResponse *v);
int deserialize_ConnectResponse(struct iarchive *in, const char *tag, struct ConnectResponse*v);
void deallocate_ConnectResponse(struct ConnectResponse*);
struct String_vector {
    int32_t count;
    char * *data;

};
int serialize_String_vector(struct oarchive *out, const char *tag, struct String_vector *v);
int deserialize_String_vector(struct iarchive *in, const char *tag, struct String_vector *v);
int allocate_String_vector(struct String_vector *v, int32_t len);
int deallocate_String_vector(struct String_vector *v);
struct SetWatches {
    int64_t relativeZxid;
    struct String_vector dataWatches;
    struct String_vector existWatches;
    struct String_vector childWatches;
    struct TMB_Trace trace;
};
int serialize_SetWatches(struct oarchive *out, const char *tag, struct SetWatches *v);
int deserialize_SetWatches(struct iarchive *in, const char *tag, struct SetWatches*v);
void deallocate_SetWatches(struct SetWatches*);
struct SetWatches2 {
    int64_t relativeZxid;
    struct String_vector dataWatches;
    struct String_vector existWatches;
    struct String_vector childWatches;
    struct String_vector persistentWatches;
    struct String_vector persistentRecursiveWatches;
    struct TMB_Trace trace;
};
int serialize_SetWatches2(struct oarchive *out, const char *tag, struct SetWatches2 *v);
int deserialize_SetWatches2(struct iarchive *in, const char *tag, struct SetWatches2*v);
void deallocate_SetWatches2(struct SetWatches2*);
struct RequestHeader {
    int32_t xid;
    int32_t type;
    struct TMB_Trace trace;
};
int serialize_RequestHeader(struct oarchive *out, const char *tag, struct RequestHeader *v);
int deserialize_RequestHeader(struct iarchive *in, const char *tag, struct RequestHeader*v);
void deallocate_RequestHeader(struct RequestHeader*);
struct MultiHeader {
    int32_t type;
    int32_t done;
    int32_t err;
    struct TMB_Trace trace;
};
int serialize_MultiHeader(struct oarchive *out, const char *tag, struct MultiHeader *v);
int deserialize_MultiHeader(struct iarchive *in, const char *tag, struct MultiHeader*v);
void deallocate_MultiHeader(struct MultiHeader*);
struct AuthPacket {
    int32_t type;
    char * scheme;
    struct buffer auth;
    struct TMB_Trace trace;
};
int serialize_AuthPacket(struct oarchive *out, const char *tag, struct AuthPacket *v);
int deserialize_AuthPacket(struct iarchive *in, const char *tag, struct AuthPacket*v);
void deallocate_AuthPacket(struct AuthPacket*);
struct ReplyHeader {
    int32_t xid;
    int64_t zxid;
    int32_t err;
    struct TMB_Trace trace;
};
int serialize_ReplyHeader(struct oarchive *out, const char *tag, struct ReplyHeader *v);
int deserialize_ReplyHeader(struct iarchive *in, const char *tag, struct ReplyHeader*v);
void deallocate_ReplyHeader(struct ReplyHeader*);
struct GetDataRequest {
    char * path;
    int32_t watch;
    struct TMB_Trace trace;
};
int serialize_GetDataRequest(struct oarchive *out, const char *tag, struct GetDataRequest *v);
int deserialize_GetDataRequest(struct iarchive *in, const char *tag, struct GetDataRequest*v);
void deallocate_GetDataRequest(struct GetDataRequest*);
struct SetDataRequest {
    char * path;
    struct buffer data;
    int32_t version;
    struct TMB_Trace trace;
};
int serialize_SetDataRequest(struct oarchive *out, const char *tag, struct SetDataRequest *v);
int deserialize_SetDataRequest(struct iarchive *in, const char *tag, struct SetDataRequest*v);
void deallocate_SetDataRequest(struct SetDataRequest*);
struct ReconfigRequest {
    char * joiningServers;
    char * leavingServers;
    char * newMembers;
    int64_t curConfigId;
    struct TMB_Trace trace;
};
int serialize_ReconfigRequest(struct oarchive *out, const char *tag, struct ReconfigRequest *v);
int deserialize_ReconfigRequest(struct iarchive *in, const char *tag, struct ReconfigRequest*v);
void deallocate_ReconfigRequest(struct ReconfigRequest*);
struct SetDataResponse {
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_SetDataResponse(struct oarchive *out, const char *tag, struct SetDataResponse *v);
int deserialize_SetDataResponse(struct iarchive *in, const char *tag, struct SetDataResponse*v);
void deallocate_SetDataResponse(struct SetDataResponse*);
struct GetSASLRequest {
    struct buffer token;
    struct TMB_Trace trace;
};
int serialize_GetSASLRequest(struct oarchive *out, const char *tag, struct GetSASLRequest *v);
int deserialize_GetSASLRequest(struct iarchive *in, const char *tag, struct GetSASLRequest*v);
void deallocate_GetSASLRequest(struct GetSASLRequest*);
struct SetSASLRequest {
    struct buffer token;
    struct TMB_Trace trace;
};
int serialize_SetSASLRequest(struct oarchive *out, const char *tag, struct SetSASLRequest *v);
int deserialize_SetSASLRequest(struct iarchive *in, const char *tag, struct SetSASLRequest*v);
void deallocate_SetSASLRequest(struct SetSASLRequest*);
struct SetSASLResponse {
    struct buffer token;
    struct TMB_Trace trace;
};
int serialize_SetSASLResponse(struct oarchive *out, const char *tag, struct SetSASLResponse *v);
int deserialize_SetSASLResponse(struct iarchive *in, const char *tag, struct SetSASLResponse*v);
void deallocate_SetSASLResponse(struct SetSASLResponse*);
struct ACL_vector {
    int32_t count;
    struct ACL *data;

};
int serialize_ACL_vector(struct oarchive *out, const char *tag, struct ACL_vector *v);
int deserialize_ACL_vector(struct iarchive *in, const char *tag, struct ACL_vector *v);
int allocate_ACL_vector(struct ACL_vector *v, int32_t len);
int deallocate_ACL_vector(struct ACL_vector *v);
struct CreateRequest {
    char * path;
    struct buffer data;
    struct ACL_vector acl;
    int32_t flags;
    struct TMB_Trace trace;
};
int serialize_CreateRequest(struct oarchive *out, const char *tag, struct CreateRequest *v);
int deserialize_CreateRequest(struct iarchive *in, const char *tag, struct CreateRequest*v);
void deallocate_CreateRequest(struct CreateRequest*);
struct CreateTTLRequest {
    char * path;
    struct buffer data;
    struct ACL_vector acl;
    int32_t flags;
    int64_t ttl;
    struct TMB_Trace trace;
};
int serialize_CreateTTLRequest(struct oarchive *out, const char *tag, struct CreateTTLRequest *v);
int deserialize_CreateTTLRequest(struct iarchive *in, const char *tag, struct CreateTTLRequest*v);
void deallocate_CreateTTLRequest(struct CreateTTLRequest*);
struct DeleteRequest {
    char * path;
    int32_t version;
    struct TMB_Trace trace;
};
int serialize_DeleteRequest(struct oarchive *out, const char *tag, struct DeleteRequest *v);
int deserialize_DeleteRequest(struct iarchive *in, const char *tag, struct DeleteRequest*v);
void deallocate_DeleteRequest(struct DeleteRequest*);
struct GetChildrenRequest {
    char * path;
    int32_t watch;
    struct TMB_Trace trace;
};
int serialize_GetChildrenRequest(struct oarchive *out, const char *tag, struct GetChildrenRequest *v);
int deserialize_GetChildrenRequest(struct iarchive *in, const char *tag, struct GetChildrenRequest*v);
void deallocate_GetChildrenRequest(struct GetChildrenRequest*);
struct GetAllChildrenNumberRequest {
    char * path;
    struct TMB_Trace trace;
};
int serialize_GetAllChildrenNumberRequest(struct oarchive *out, const char *tag, struct GetAllChildrenNumberRequest *v);
int deserialize_GetAllChildrenNumberRequest(struct iarchive *in, const char *tag, struct GetAllChildrenNumberRequest*v);
void deallocate_GetAllChildrenNumberRequest(struct GetAllChildrenNumberRequest*);
struct GetChildren2Request {
    char * path;
    int32_t watch;
    struct TMB_Trace trace;
};
int serialize_GetChildren2Request(struct oarchive *out, const char *tag, struct GetChildren2Request *v);
int deserialize_GetChildren2Request(struct iarchive *in, const char *tag, struct GetChildren2Request*v);
void deallocate_GetChildren2Request(struct GetChildren2Request*);
struct CheckVersionRequest {
    char * path;
    int32_t version;
    struct TMB_Trace trace;
};
int serialize_CheckVersionRequest(struct oarchive *out, const char *tag, struct CheckVersionRequest *v);
int deserialize_CheckVersionRequest(struct iarchive *in, const char *tag, struct CheckVersionRequest*v);
void deallocate_CheckVersionRequest(struct CheckVersionRequest*);
struct GetMaxChildrenRequest {
    char * path;
    struct TMB_Trace trace;
};
int serialize_GetMaxChildrenRequest(struct oarchive *out, const char *tag, struct GetMaxChildrenRequest *v);
int deserialize_GetMaxChildrenRequest(struct iarchive *in, const char *tag, struct GetMaxChildrenRequest*v);
void deallocate_GetMaxChildrenRequest(struct GetMaxChildrenRequest*);
struct GetMaxChildrenResponse {
    int32_t max;
    struct TMB_Trace trace;
};
int serialize_GetMaxChildrenResponse(struct oarchive *out, const char *tag, struct GetMaxChildrenResponse *v);
int deserialize_GetMaxChildrenResponse(struct iarchive *in, const char *tag, struct GetMaxChildrenResponse*v);
void deallocate_GetMaxChildrenResponse(struct GetMaxChildrenResponse*);
struct SetMaxChildrenRequest {
    char * path;
    int32_t max;
    struct TMB_Trace trace;
};
int serialize_SetMaxChildrenRequest(struct oarchive *out, const char *tag, struct SetMaxChildrenRequest *v);
int deserialize_SetMaxChildrenRequest(struct iarchive *in, const char *tag, struct SetMaxChildrenRequest*v);
void deallocate_SetMaxChildrenRequest(struct SetMaxChildrenRequest*);
struct SyncRequest {
    char * path;
    struct TMB_Trace trace;
};
int serialize_SyncRequest(struct oarchive *out, const char *tag, struct SyncRequest *v);
int deserialize_SyncRequest(struct iarchive *in, const char *tag, struct SyncRequest*v);
void deallocate_SyncRequest(struct SyncRequest*);
struct SyncResponse {
    char * path;
    struct TMB_Trace trace;
};
int serialize_SyncResponse(struct oarchive *out, const char *tag, struct SyncResponse *v);
int deserialize_SyncResponse(struct iarchive *in, const char *tag, struct SyncResponse*v);
void deallocate_SyncResponse(struct SyncResponse*);
struct GetACLRequest {
    char * path;
    struct TMB_Trace trace;
};
int serialize_GetACLRequest(struct oarchive *out, const char *tag, struct GetACLRequest *v);
int deserialize_GetACLRequest(struct iarchive *in, const char *tag, struct GetACLRequest*v);
void deallocate_GetACLRequest(struct GetACLRequest*);
struct SetACLRequest {
    char * path;
    struct ACL_vector acl;
    int32_t version;
    struct TMB_Trace trace;
};
int serialize_SetACLRequest(struct oarchive *out, const char *tag, struct SetACLRequest *v);
int deserialize_SetACLRequest(struct iarchive *in, const char *tag, struct SetACLRequest*v);
void deallocate_SetACLRequest(struct SetACLRequest*);
struct SetACLResponse {
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_SetACLResponse(struct oarchive *out, const char *tag, struct SetACLResponse *v);
int deserialize_SetACLResponse(struct iarchive *in, const char *tag, struct SetACLResponse*v);
void deallocate_SetACLResponse(struct SetACLResponse*);
struct AddWatchRequest {
    char * path;
    int32_t mode;
    struct TMB_Trace trace;
};
int serialize_AddWatchRequest(struct oarchive *out, const char *tag, struct AddWatchRequest *v);
int deserialize_AddWatchRequest(struct iarchive *in, const char *tag, struct AddWatchRequest*v);
void deallocate_AddWatchRequest(struct AddWatchRequest*);
struct WatcherEvent {
    int32_t type;
    int32_t state;
    char * path;
    struct TMB_Trace trace;
};
int serialize_WatcherEvent(struct oarchive *out, const char *tag, struct WatcherEvent *v);
int deserialize_WatcherEvent(struct iarchive *in, const char *tag, struct WatcherEvent*v);
void deallocate_WatcherEvent(struct WatcherEvent*);
struct ErrorResponse {
    int32_t err;
    struct TMB_Trace trace;
};
int serialize_ErrorResponse(struct oarchive *out, const char *tag, struct ErrorResponse *v);
int deserialize_ErrorResponse(struct iarchive *in, const char *tag, struct ErrorResponse*v);
void deallocate_ErrorResponse(struct ErrorResponse*);
struct CreateResponse {
    char * path;
    struct TMB_Trace trace;
};
int serialize_CreateResponse(struct oarchive *out, const char *tag, struct CreateResponse *v);
int deserialize_CreateResponse(struct iarchive *in, const char *tag, struct CreateResponse*v);
void deallocate_CreateResponse(struct CreateResponse*);
struct Create2Response {
    char * path;
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_Create2Response(struct oarchive *out, const char *tag, struct Create2Response *v);
int deserialize_Create2Response(struct iarchive *in, const char *tag, struct Create2Response*v);
void deallocate_Create2Response(struct Create2Response*);
struct ExistsRequest {
    char * path;
    int32_t watch;
    struct TMB_Trace trace;
};
int serialize_ExistsRequest(struct oarchive *out, const char *tag, struct ExistsRequest *v);
int deserialize_ExistsRequest(struct iarchive *in, const char *tag, struct ExistsRequest*v);
void deallocate_ExistsRequest(struct ExistsRequest*);
struct ExistsResponse {
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_ExistsResponse(struct oarchive *out, const char *tag, struct ExistsResponse *v);
int deserialize_ExistsResponse(struct iarchive *in, const char *tag, struct ExistsResponse*v);
void deallocate_ExistsResponse(struct ExistsResponse*);
struct GetDataResponse {
    struct buffer data;
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_GetDataResponse(struct oarchive *out, const char *tag, struct GetDataResponse *v);
int deserialize_GetDataResponse(struct iarchive *in, const char *tag, struct GetDataResponse*v);
void deallocate_GetDataResponse(struct GetDataResponse*);
struct GetChildrenResponse {
    struct String_vector children;
    struct TMB_Trace trace;
};
int serialize_GetChildrenResponse(struct oarchive *out, const char *tag, struct GetChildrenResponse *v);
int deserialize_GetChildrenResponse(struct iarchive *in, const char *tag, struct GetChildrenResponse*v);
void deallocate_GetChildrenResponse(struct GetChildrenResponse*);
struct GetAllChildrenNumberResponse {
    int32_t totalNumber;
    struct TMB_Trace trace;
};
int serialize_GetAllChildrenNumberResponse(struct oarchive *out, const char *tag, struct GetAllChildrenNumberResponse *v);
int deserialize_GetAllChildrenNumberResponse(struct iarchive *in, const char *tag, struct GetAllChildrenNumberResponse*v);
void deallocate_GetAllChildrenNumberResponse(struct GetAllChildrenNumberResponse*);
struct GetChildren2Response {
    struct String_vector children;
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_GetChildren2Response(struct oarchive *out, const char *tag, struct GetChildren2Response *v);
int deserialize_GetChildren2Response(struct iarchive *in, const char *tag, struct GetChildren2Response*v);
void deallocate_GetChildren2Response(struct GetChildren2Response*);
struct GetACLResponse {
    struct ACL_vector acl;
    struct Stat stat;
    struct TMB_Trace trace;
};
int serialize_GetACLResponse(struct oarchive *out, const char *tag, struct GetACLResponse *v);
int deserialize_GetACLResponse(struct iarchive *in, const char *tag, struct GetACLResponse*v);
void deallocate_GetACLResponse(struct GetACLResponse*);
struct CheckWatchesRequest {
    char * path;
    int32_t type;
    struct TMB_Trace trace;
};
int serialize_CheckWatchesRequest(struct oarchive *out, const char *tag, struct CheckWatchesRequest *v);
int deserialize_CheckWatchesRequest(struct iarchive *in, const char *tag, struct CheckWatchesRequest*v);
void deallocate_CheckWatchesRequest(struct CheckWatchesRequest*);
struct RemoveWatchesRequest {
    char * path;
    int32_t type;
    struct TMB_Trace trace;
};
int serialize_RemoveWatchesRequest(struct oarchive *out, const char *tag, struct RemoveWatchesRequest *v);
int deserialize_RemoveWatchesRequest(struct iarchive *in, const char *tag, struct RemoveWatchesRequest*v);
void deallocate_RemoveWatchesRequest(struct RemoveWatchesRequest*);
struct GetEphemeralsRequest {
    char * prefixPath;
    struct TMB_Trace trace;
};
int serialize_GetEphemeralsRequest(struct oarchive *out, const char *tag, struct GetEphemeralsRequest *v);
int deserialize_GetEphemeralsRequest(struct iarchive *in, const char *tag, struct GetEphemeralsRequest*v);
void deallocate_GetEphemeralsRequest(struct GetEphemeralsRequest*);
struct GetEphemeralsResponse {
    struct String_vector ephemerals;
    struct TMB_Trace trace;
};
int serialize_GetEphemeralsResponse(struct oarchive *out, const char *tag, struct GetEphemeralsResponse *v);
int deserialize_GetEphemeralsResponse(struct iarchive *in, const char *tag, struct GetEphemeralsResponse*v);
void deallocate_GetEphemeralsResponse(struct GetEphemeralsResponse*);
struct LearnerInfo {
    int64_t serverid;
    int32_t protocolVersion;
    int64_t configVersion;
};
int serialize_LearnerInfo(struct oarchive *out, const char *tag, struct LearnerInfo *v);
int deserialize_LearnerInfo(struct iarchive *in, const char *tag, struct LearnerInfo*v);
void deallocate_LearnerInfo(struct LearnerInfo*);
struct Id_vector {
    int32_t count;
    struct Id *data;

};
int serialize_Id_vector(struct oarchive *out, const char *tag, struct Id_vector *v);
int deserialize_Id_vector(struct iarchive *in, const char *tag, struct Id_vector *v);
int allocate_Id_vector(struct Id_vector *v, int32_t len);
int deallocate_Id_vector(struct Id_vector *v);
struct QuorumPacket {
    int32_t type;
    int64_t zxid;
    struct buffer data;
    struct Id_vector authinfo;
};
int serialize_QuorumPacket(struct oarchive *out, const char *tag, struct QuorumPacket *v);
int deserialize_QuorumPacket(struct iarchive *in, const char *tag, struct QuorumPacket*v);
void deallocate_QuorumPacket(struct QuorumPacket*);
struct QuorumAuthPacket {
    int64_t magic;
    int32_t status;
    struct buffer token;
};
int serialize_QuorumAuthPacket(struct oarchive *out, const char *tag, struct QuorumAuthPacket *v);
int deserialize_QuorumAuthPacket(struct iarchive *in, const char *tag, struct QuorumAuthPacket*v);
void deallocate_QuorumAuthPacket(struct QuorumAuthPacket*);
struct FileHeader {
    int32_t magic;
    int32_t version;
    int64_t dbid;
};
int serialize_FileHeader(struct oarchive *out, const char *tag, struct FileHeader *v);
int deserialize_FileHeader(struct iarchive *in, const char *tag, struct FileHeader*v);
void deallocate_FileHeader(struct FileHeader*);
struct TxnDigest {
    int32_t version;
    int64_t treeDigest;
};
int serialize_TxnDigest(struct oarchive *out, const char *tag, struct TxnDigest *v);
int deserialize_TxnDigest(struct iarchive *in, const char *tag, struct TxnDigest*v);
void deallocate_TxnDigest(struct TxnDigest*);
struct TxnHeader {
    int64_t clientId;
    int32_t cxid;
    int64_t zxid;
    int64_t time;
    int32_t type;
};
int serialize_TxnHeader(struct oarchive *out, const char *tag, struct TxnHeader *v);
int deserialize_TxnHeader(struct iarchive *in, const char *tag, struct TxnHeader*v);
void deallocate_TxnHeader(struct TxnHeader*);
struct CreateTxnV0 {
    char * path;
    struct buffer data;
    struct ACL_vector acl;
    int32_t ephemeral;
};
int serialize_CreateTxnV0(struct oarchive *out, const char *tag, struct CreateTxnV0 *v);
int deserialize_CreateTxnV0(struct iarchive *in, const char *tag, struct CreateTxnV0*v);
void deallocate_CreateTxnV0(struct CreateTxnV0*);
struct CreateTxn {
    char * path;
    struct buffer data;
    struct ACL_vector acl;
    int32_t ephemeral;
    int32_t parentCVersion;
};
int serialize_CreateTxn(struct oarchive *out, const char *tag, struct CreateTxn *v);
int deserialize_CreateTxn(struct iarchive *in, const char *tag, struct CreateTxn*v);
void deallocate_CreateTxn(struct CreateTxn*);
struct CreateTTLTxn {
    char * path;
    struct buffer data;
    struct ACL_vector acl;
    int32_t parentCVersion;
    int64_t ttl;
};
int serialize_CreateTTLTxn(struct oarchive *out, const char *tag, struct CreateTTLTxn *v);
int deserialize_CreateTTLTxn(struct iarchive *in, const char *tag, struct CreateTTLTxn*v);
void deallocate_CreateTTLTxn(struct CreateTTLTxn*);
struct CreateContainerTxn {
    char * path;
    struct buffer data;
    struct ACL_vector acl;
    int32_t parentCVersion;
};
int serialize_CreateContainerTxn(struct oarchive *out, const char *tag, struct CreateContainerTxn *v);
int deserialize_CreateContainerTxn(struct iarchive *in, const char *tag, struct CreateContainerTxn*v);
void deallocate_CreateContainerTxn(struct CreateContainerTxn*);
struct DeleteTxn {
    char * path;
};
int serialize_DeleteTxn(struct oarchive *out, const char *tag, struct DeleteTxn *v);
int deserialize_DeleteTxn(struct iarchive *in, const char *tag, struct DeleteTxn*v);
void deallocate_DeleteTxn(struct DeleteTxn*);
struct SetDataTxn {
    char * path;
    struct buffer data;
    int32_t version;
};
int serialize_SetDataTxn(struct oarchive *out, const char *tag, struct SetDataTxn *v);
int deserialize_SetDataTxn(struct iarchive *in, const char *tag, struct SetDataTxn*v);
void deallocate_SetDataTxn(struct SetDataTxn*);
struct CheckVersionTxn {
    char * path;
    int32_t version;
};
int serialize_CheckVersionTxn(struct oarchive *out, const char *tag, struct CheckVersionTxn *v);
int deserialize_CheckVersionTxn(struct iarchive *in, const char *tag, struct CheckVersionTxn*v);
void deallocate_CheckVersionTxn(struct CheckVersionTxn*);
struct SetACLTxn {
    char * path;
    struct ACL_vector acl;
    int32_t version;
};
int serialize_SetACLTxn(struct oarchive *out, const char *tag, struct SetACLTxn *v);
int deserialize_SetACLTxn(struct iarchive *in, const char *tag, struct SetACLTxn*v);
void deallocate_SetACLTxn(struct SetACLTxn*);
struct SetMaxChildrenTxn {
    char * path;
    int32_t max;
};
int serialize_SetMaxChildrenTxn(struct oarchive *out, const char *tag, struct SetMaxChildrenTxn *v);
int deserialize_SetMaxChildrenTxn(struct iarchive *in, const char *tag, struct SetMaxChildrenTxn*v);
void deallocate_SetMaxChildrenTxn(struct SetMaxChildrenTxn*);
struct CreateSessionTxn {
    int32_t timeOut;
};
int serialize_CreateSessionTxn(struct oarchive *out, const char *tag, struct CreateSessionTxn *v);
int deserialize_CreateSessionTxn(struct iarchive *in, const char *tag, struct CreateSessionTxn*v);
void deallocate_CreateSessionTxn(struct CreateSessionTxn*);
struct CloseSessionTxn {
    struct String_vector paths2Delete;
};
int serialize_CloseSessionTxn(struct oarchive *out, const char *tag, struct CloseSessionTxn *v);
int deserialize_CloseSessionTxn(struct iarchive *in, const char *tag, struct CloseSessionTxn*v);
void deallocate_CloseSessionTxn(struct CloseSessionTxn*);
struct ErrorTxn {
    int32_t err;
};
int serialize_ErrorTxn(struct oarchive *out, const char *tag, struct ErrorTxn *v);
int deserialize_ErrorTxn(struct iarchive *in, const char *tag, struct ErrorTxn*v);
void deallocate_ErrorTxn(struct ErrorTxn*);
struct Txn {
    int32_t type;
    struct buffer data;
};
int serialize_Txn(struct oarchive *out, const char *tag, struct Txn *v);
int deserialize_Txn(struct iarchive *in, const char *tag, struct Txn*v);
void deallocate_Txn(struct Txn*);
struct Txn_vector {
    int32_t count;
    struct Txn *data;

};
int serialize_Txn_vector(struct oarchive *out, const char *tag, struct Txn_vector *v);
int deserialize_Txn_vector(struct iarchive *in, const char *tag, struct Txn_vector *v);
int allocate_Txn_vector(struct Txn_vector *v, int32_t len);
int deallocate_Txn_vector(struct Txn_vector *v);
struct MultiTxn {
    struct Txn_vector txns;
};
int serialize_MultiTxn(struct oarchive *out, const char *tag, struct MultiTxn *v);
int deserialize_MultiTxn(struct iarchive *in, const char *tag, struct MultiTxn*v);
void deallocate_MultiTxn(struct MultiTxn*);

#ifdef __cplusplus
}
#endif

#endif //ZOOKEEPER_JUTE__
