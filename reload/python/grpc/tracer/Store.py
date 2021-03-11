#! /usr/bin/python

import uuid
import threading

class Store:
    def __init__(self):
        pass

    uuid = str(uuid.uuid4())
    lock = threading.RLock()

    # hashmap of int(long traceID) to traces
    traces = dict()

    def SetTrace(self, traceID, t):
        with self.lock:
            self.traces[traceID] = t

    def GetTrace(self, traceID):
        with self.lock:
            t = self.traces.get(traceID)

        return t

    def FetchTrace(self, traceID):
        t = None

        with self.lock:
            if self.traces[traceID]:
                t = self.traces.get(traceID)
                del self.traces[traceID]

        return t

    def PrintStore(self):
        return self.traces

    # class ContextMeta:
    #    traceID = int()
    #    uuid = str()
    #
    #     def ContextMeta(self, traceID: int, uuid: str):
    #         self.traceID = traceID
    #        self.uuid = uuid
    # class Traces:
    # traces
    # def CheckByContextMeta(self, meta: ContextMeta) -> bool:
    # ok = false
    # return ok
