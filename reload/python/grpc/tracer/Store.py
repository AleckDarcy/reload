#! /usr/bin/python

import uuid
import threading


class Store:
    uuid = str(uuid.uuid4())
    lock = threading.RLock()

    # hashmap of int(long threadID) to traces
    traces = dict()

    def SetTrace(self, threadID: int, t: Message.Trace):
        with self.lock:
            self.traces[threadID] = t

    def GetTrace(self, threadID: int) -> Message.Trace:
        t = None

        with self.lock:
            if self.traces[threadID]:
                t = self.traces[threadID]

        return t

    def FetchTrace(self, threadID: int) -> Message.Trace:
        t = None

        with self.lock:
            if self.traces[threadID]:
                t = self.traces[threadID]
                del self.traces[threadID]

        return t

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
