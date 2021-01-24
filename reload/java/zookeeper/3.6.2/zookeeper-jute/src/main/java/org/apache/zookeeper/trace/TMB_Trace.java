// File generated by hadoop record compiler. Do not edit.
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

package org.apache.zookeeper.trace;

import org.apache.jute.*;
import org.apache.jute.Record;
import org.apache.yetus.audience.InterfaceAudience;

import java.util.ArrayList;

@InterfaceAudience.Public
public class TMB_Trace implements Record {
  private long id;
  private long req_event;
  private java.util.List<TMB_Event> events;
  private java.util.List<TMB_TFI> tfis;
  public TMB_Trace() {
    this.events=new java.util.ArrayList<>();
    this.tfis=new java.util.ArrayList<>();
  }
  public TMB_Trace(
        long id,
        long req_event,
        java.util.List<TMB_Event> events,
        java.util.List<TMB_TFI> tfis) {
    this.id=id;
    this.req_event=req_event;
    this.events=events;
    this.tfis=tfis;
  }

  // 3MileBeach
  public String toJSON() {
    StringBuffer buffer = new StringBuffer(String.format("{\"id\":%d,\"req_event\":%d,\"events\":[", id, req_event));
    int i = 0;
    for (TMB_Event event: events) {
      String type = event.getType() == 1? "SEND": "RECV";
      if (i != 0) {
        buffer.append(String.format(",\n{\"service\":\"%s\",\"type\":\"%s\",\"timestamp\":%d,\"message_name\":\"%s\",\"uuid\":\"%s\"}",
                event.getService(), type, event.getTimestamp(), event.getMessage_name(), event.getUuid()));
      } else {
        buffer.append(String.format("{\"service\":\"%s\",\"type\":\"%s\",\"timestamp\":%d,\"message_name\":\"%s\",\"uuid\":\"%s\"}",
                event.getService(), type, event.getTimestamp(), event.getMessage_name(), event.getUuid()));
      }

      i++;
    }

    buffer.append("],\"tfis\":[]}");

    return buffer.toString();
  }

  public void addEvent(TMB_Event e) {
    // TODO: remove
    if (events == null) {
      events = new ArrayList<>();
    }

    events.add(e);
  }

  public long getReqEvent() {
    return req_event;
  }
  public void setReqEvent(long m_) { req_event=m_; }
  public long getId() {
    return id;
  }
  public void setId(long m_) {
    id=m_;
  }
  public java.util.List<TMB_Event> getEvents() {
    return events;
  }
  public void setEvents(java.util.List<TMB_Event> m_) {
    events=m_;
  }
  public java.util.List<TMB_TFI> getTfis() {
    return tfis;
  }
  public void setTfis(java.util.List<TMB_TFI> m_) {
    tfis=m_;
  }
  public org.apache.zookeeper.trace.TMB_Trace getTrace() {
    return null;
  }
  public void setTrace(org.apache.zookeeper.trace.TMB_Trace m_) {}
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    a_.writeLong(id,"id");
    a_.writeLong(req_event,"req_event");
    {
      a_.startVector(events,"events");
      if (events!= null) {          int len1 = events.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            TMB_Event e1 = (TMB_Event) events.get(vidx1);
    a_.writeRecord(e1,"e1");
          }
      }
      a_.endVector(events,"events");
    }
    {
      a_.startVector(tfis,"tfis");
      if (tfis!= null) {          int len1 = tfis.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            TMB_TFI e1 = (TMB_TFI) tfis.get(vidx1);
    a_.writeRecord(e1,"e1");
          }
      }
      a_.endVector(tfis,"tfis");
    }
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    id=a_.readLong("id");
    req_event=a_.readLong("req_event");
    {
      Index vidx1 = a_.startVector("events");
      if (vidx1!= null) {          events=new java.util.ArrayList<TMB_Event>();
          for (; !vidx1.done(); vidx1.incr()) {
    TMB_Event e1;
    e1= new TMB_Event();
    a_.readRecord(e1,"e1");
            events.add(e1);
          }
      }
    a_.endVector("events");
    }
    {
      Index vidx1 = a_.startVector("tfis");
      if (vidx1!= null) {          tfis=new java.util.ArrayList<TMB_TFI>();
          for (; !vidx1.done(); vidx1.incr()) {
    TMB_TFI e1;
    e1= new TMB_TFI();
    a_.readRecord(e1,"e1");
            tfis.add(e1);
          }
      }
    a_.endVector("tfis");
    }
    a_.endRecord(tag);
}
  public String toString() {
    try {
      java.io.ByteArrayOutputStream s =
        new java.io.ByteArrayOutputStream();
      ToStringOutputArchive a_ =
        new ToStringOutputArchive(s);
      a_.startRecord(this,"");
    a_.writeLong(id,"id");
    a_.writeLong(req_event,"req_event");
    {
      a_.startVector(events,"events");
      if (events!= null) {          int len1 = events.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            TMB_Event e1 = (TMB_Event) events.get(vidx1);
    a_.writeRecord(e1,"e1");
          }
      }
      a_.endVector(events,"events");
    }
    {
      a_.startVector(tfis,"tfis");
      if (tfis!= null) {          int len1 = tfis.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            TMB_TFI e1 = (TMB_TFI) tfis.get(vidx1);
    a_.writeRecord(e1,"e1");
          }
      }
      a_.endVector(tfis,"tfis");
    }
      a_.endRecord(this,"");
      return new String(s.toByteArray(), "UTF-8");
    } catch (Throwable ex) {
      ex.printStackTrace();
    }
    return "ERROR";
  }
  public void write(java.io.DataOutput out) throws java.io.IOException {
    BinaryOutputArchive archive = new BinaryOutputArchive(out);
    serialize(archive, "");
  }
  public void readFields(java.io.DataInput in) throws java.io.IOException {
    BinaryInputArchive archive = new BinaryInputArchive(in);
    deserialize(archive, "");
  }
  public int compareTo (Object peer_) throws ClassCastException {
    throw new UnsupportedOperationException("comparing TMB_Trace is unimplemented");
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof TMB_Trace)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    TMB_Trace peer = (TMB_Trace) peer_;
    boolean ret = false;
    ret = (id==peer.id);
    if (!ret) return ret;
    ret = (req_event==peer.req_event);
    if (!ret) return ret;
    ret = events.equals(peer.events);
    if (!ret) return ret;
    ret = tfis.equals(peer.tfis);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = (int) (id^(id>>>32));
    result = 37*result + ret;
    ret = (int) (req_event^(req_event>>>32));
    result = 37*result + ret;
    ret = events.hashCode();
    result = 37*result + ret;
    ret = tfis.hashCode();
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LTMB_Trace(ll[LTMB_Event(ilsss)][LTMB_TFI(isl[LTMB_TFIMeta(sll)])])";
  }
}
