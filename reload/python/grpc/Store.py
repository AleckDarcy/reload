#! /usr/bin/python

import threading    # for locks

def merge(dst, src):
    dstI, srcI, dstLen, srcLen = 0, 0, len(dst.records), len(src.records)
    while dstI < dstLen and srcI < srcLen:
        dstRecord, srcRecord = dst.records[dstI], src.records[srcI]
        if dstRecord.timestamp > srcRecord.timestamp:
            #TODO: do this by indexing pythonic way instead of two for loops
            for j in range(0, dstI):
                dst.records.extend([dst.records[j]])
            for i in range(dstI + 1, len(dst.records)):
                dst.records.extend([dst.records[i]])
            dst.records[dstI] = srcRecord

            srcI += 1
            dstLen += 1
        elif dstRecord.timestamp == srcRecord.timestamp:
            if dstRecord.message_name == srcRecord.message_name and dstRecord.type == srcRecord.type:
                srcI += 1
        else:
            dstI += 1

        dstI += 1


class Store:
    def __init__(self):
        pass

    lock = threading.RLock()

    # hashmap ->
    # traceID: traces
    # traces ->
    # UUID: Trace
    store = dict()

    def CheckByContextMeta(self, meta):
        ok = False

        with self.lock:
            # check if meta.traceID exists in store and get traces dict
            traces = self.store.get(meta.traceID)
            if traces:
                # check if meta.uuid yields us a trace -> set ok true
                tmp = traces.get(meta.uuid)
                if tmp:
                    ok = True

        return ok

    # GUARD against trace being None (check return value)
    def GetByContextMeta(self, meta):
        ok = False
        t = None

        with self.lock:
            # check if meta.traceID exists in store and get traces dict
            traces = self.store.get(meta.traceID)
            if traces:
                # check if meta.uuid yields us a trace -> set ok true
                t = traces.get(meta.uuid)
                if t:
                    ok = True
                    # shallow copy in python 2.7
                    t = t[:]

        return t, ok

    def SetByContextMeta(self, meta, trace):
        with self.lock:
            # check if meta.traceID exists in store and get traces dict
            traces = self.store.get(meta.traceID)
            if traces:
                # check if meta.uuid yields us a trace -> set ok true
                t = traces.get(meta.uuid)
                if t:
                    merge(t, trace)
                else:
                    traces[meta.uuid] = trace
            else:
                tmpTrace = {meta.uuid: trace}
                self.store[meta.traceID] = tmpTrace

    def UpdateFunctionByContextMeta(self, meta, function):
        ok = False
        t = None

        with self.lock:
            # check if meta.traceID exists in store and get traces dict
            traces = self.store.get(meta.traceID)
            if traces:
                # check if meta.uuid yields us a trace -> set ok true
                t = traces.get(meta.uuid)
                if t:
                    ok = True

                    function(t)
                    # shallow copy in python 2.7
                    t = t[:]

        return t, ok

    def DeleteByContextMeta(self, meta):
        ok = False

        with self.lock:
            # check if meta.traceID exists in store and get traces dict
            traces = self.store.get(meta.traceID)
            if traces:
                # check if meta.uuid yields us a trace -> set ok true
                t = traces.get(meta.uuid)
                if t:
                    ok = True
                    del traces[meta.uuid]

                    if len(traces) == 0:
                        del self.store[meta.traceID]

        return ok

    # internal use
    def GetStore(self):
        return self.store
