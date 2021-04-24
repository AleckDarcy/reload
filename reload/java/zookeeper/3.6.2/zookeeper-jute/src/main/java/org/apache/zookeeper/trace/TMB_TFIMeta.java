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

@InterfaceAudience.Public
public class TMB_TFIMeta implements Record {
  private String name;
  private int event_type;
  private long times;
  private long already;
  public TMB_TFIMeta() {
  }
  public TMB_TFIMeta(
        String name,
        int event_type,
        long times,
        long already) {
    this.name=name;
    this.event_type=event_type;
    this.times=times;
    this.already=already;
  }
  public TMB_TFIMeta(
          String name,
          int event_type,
          long times) {
    this.name=name;
    this.event_type=event_type;
    this.times=times;
  }
  public String getName() {
    return name;
  }
  public void setName(String m_) {
    name=m_;
  }
  public int getEvent_type() {
    return event_type;
  }
  public void setEvent_type(int m_) {
    event_type=m_;
  }
  public long getTimes() {
    return times;
  }
  public void setTimes(long m_) {
    times=m_;
  }
  public long getAlready() {
    return already;
  }
  public void setAlready(long m_) {
    already=m_;
  }
  public org.apache.zookeeper.trace.TMB_Trace getTrace() {
    return null;
  }
  public void setTrace(org.apache.zookeeper.trace.TMB_Trace m_) {}
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    a_.writeString(name,"name");
    a_.writeInt(event_type,"event_type");
    a_.writeLong(times,"times");
    a_.writeLong(already,"already");
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    name=a_.readString("name");
    event_type=a_.readInt("event_type");
    times=a_.readLong("times");
    already=a_.readLong("already");
    a_.endRecord(tag);
}
  public String toString() {
    try {
      java.io.ByteArrayOutputStream s =
        new java.io.ByteArrayOutputStream();
      ToStringOutputArchive a_ = 
        new ToStringOutputArchive(s);
      a_.startRecord(this,"");
    a_.writeString(name,"name");
    a_.writeInt(event_type,"event_type");
    a_.writeLong(times,"times");
    a_.writeLong(already,"already");
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
    if (!(peer_ instanceof TMB_TFIMeta)) {
      throw new ClassCastException("Comparing different types of records.");
    }
    TMB_TFIMeta peer = (TMB_TFIMeta) peer_;
    int ret = 0;
    ret = name.compareTo(peer.name);
    if (ret != 0) return ret;
    ret = (event_type==peer.event_type)? 0: ((event_type<peer.event_type)?-1:1);
    if (ret != 0) return ret;
    ret = (times == peer.times)? 0 :((times<peer.times)?-1:1);
    if (ret != 0) return ret;
    ret = (already == peer.already)? 0 :((already<peer.already)?-1:1);
    if (ret != 0) return ret;
     return ret;
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof TMB_TFIMeta)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    TMB_TFIMeta peer = (TMB_TFIMeta) peer_;
    boolean ret = false;
    ret = name.equals(peer.name);
    if (!ret) return ret;
    ret = (event_type==peer.event_type);
    if (!ret) return ret;
    ret = (times==peer.times);
    if (!ret) return ret;
    ret = (already==peer.already);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = name.hashCode();
    result = 37*result + ret;
    ret = (int)event_type;
    result = 37*result + ret;
    ret = (int) (times^(times>>>32));
    result = 37*result + ret;
    ret = (int) (already^(already>>>32));
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LTMB_TFIMeta(sill)";
  }
}
