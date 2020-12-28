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
public class TMB_TFI implements Record {
  private int type;
  private String name;
  private long delay;
  private java.util.List<TMB_TFIMeta> after;
  public TMB_TFI() {
  }
  public TMB_TFI(
        int type,
        String name,
        long delay,
        java.util.List<TMB_TFIMeta> after) {
    this.type=type;
    this.name=name;
    this.delay=delay;
    this.after=after;
  }
  public int getType() {
    return type;
  }
  public void setType(int m_) {
    type=m_;
  }
  public String getName() {
    return name;
  }
  public void setName(String m_) {
    name=m_;
  }
  public long getDelay() {
    return delay;
  }
  public void setDelay(long m_) {
    delay=m_;
  }
  public java.util.List<TMB_TFIMeta> getAfter() {
    return after;
  }
  public void setAfter(java.util.List<TMB_TFIMeta> m_) {
    after=m_;
  }
  public org.apache.zookeeper.trace.TMB_Trace getTrace() {
    return null;
  }
  public void setTrace(org.apache.zookeeper.trace.TMB_Trace m_) {}
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    a_.writeInt(type,"type");
    a_.writeString(name,"name");
    a_.writeLong(delay,"delay");
    {
      a_.startVector(after,"after");
      if (after!= null) {          int len1 = after.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            TMB_TFIMeta e1 = (TMB_TFIMeta) after.get(vidx1);
    a_.writeRecord(e1,"e1");
          }
      }
      a_.endVector(after,"after");
    }
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    type=a_.readInt("type");
    name=a_.readString("name");
    delay=a_.readLong("delay");
    {
      Index vidx1 = a_.startVector("after");
      if (vidx1!= null) {          after=new java.util.ArrayList<TMB_TFIMeta>();
          for (; !vidx1.done(); vidx1.incr()) {
    TMB_TFIMeta e1;
    e1= new TMB_TFIMeta();
    a_.readRecord(e1,"e1");
            after.add(e1);
          }
      }
    a_.endVector("after");
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
    a_.writeInt(type,"type");
    a_.writeString(name,"name");
    a_.writeLong(delay,"delay");
    {
      a_.startVector(after,"after");
      if (after!= null) {          int len1 = after.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            TMB_TFIMeta e1 = (TMB_TFIMeta) after.get(vidx1);
    a_.writeRecord(e1,"e1");
          }
      }
      a_.endVector(after,"after");
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
    throw new UnsupportedOperationException("comparing TMB_TFI is unimplemented");
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof TMB_TFI)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    TMB_TFI peer = (TMB_TFI) peer_;
    boolean ret = false;
    ret = (type==peer.type);
    if (!ret) return ret;
    ret = name.equals(peer.name);
    if (!ret) return ret;
    ret = (delay==peer.delay);
    if (!ret) return ret;
    ret = after.equals(peer.after);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = (int)type;
    result = 37*result + ret;
    ret = name.hashCode();
    result = 37*result + ret;
    ret = (int) (delay^(delay>>>32));
    result = 37*result + ret;
    ret = after.hashCode();
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LTMB_TFI(isl[LTMB_TFIMeta(sll)])";
  }
}
