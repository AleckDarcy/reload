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

package org.apache.zookeeper.proto;

import org.apache.jute.*;
import org.apache.jute.Record; // JDK14 needs explicit import due to clash with java.lang.Record
import org.apache.yetus.audience.InterfaceAudience;
@InterfaceAudience.Public
public class SetWatches2 implements Record {
  private long relativeZxid;
  private java.util.List<String> dataWatches;
  private java.util.List<String> existWatches;
  private java.util.List<String> childWatches;
  private java.util.List<String> persistentWatches;
  private java.util.List<String> persistentRecursiveWatches;
  private org.apache.zookeeper.trace.TMB_Trace trace;
  public SetWatches2() {
    this.trace = new org.apache.zookeeper.trace.TMB_Trace();
  }
  public SetWatches2(
        long relativeZxid,
        java.util.List<String> dataWatches,
        java.util.List<String> existWatches,
        java.util.List<String> childWatches,
        java.util.List<String> persistentWatches,
        java.util.List<String> persistentRecursiveWatches) {
    this.relativeZxid=relativeZxid;
    this.dataWatches=dataWatches;
    this.existWatches=existWatches;
    this.childWatches=childWatches;
    this.persistentWatches=persistentWatches;
    this.persistentRecursiveWatches=persistentRecursiveWatches;
    this.trace = new org.apache.zookeeper.trace.TMB_Trace();
  }
  public long getRelativeZxid() {
    return relativeZxid;
  }
  public void setRelativeZxid(long m_) {
    relativeZxid=m_;
  }
  public java.util.List<String> getDataWatches() {
    return dataWatches;
  }
  public void setDataWatches(java.util.List<String> m_) {
    dataWatches=m_;
  }
  public java.util.List<String> getExistWatches() {
    return existWatches;
  }
  public void setExistWatches(java.util.List<String> m_) {
    existWatches=m_;
  }
  public java.util.List<String> getChildWatches() {
    return childWatches;
  }
  public void setChildWatches(java.util.List<String> m_) {
    childWatches=m_;
  }
  public java.util.List<String> getPersistentWatches() {
    return persistentWatches;
  }
  public void setPersistentWatches(java.util.List<String> m_) {
    persistentWatches=m_;
  }
  public java.util.List<String> getPersistentRecursiveWatches() {
    return persistentRecursiveWatches;
  }
  public void setPersistentRecursiveWatches(java.util.List<String> m_) {
    persistentRecursiveWatches=m_;
  }
  public org.apache.zookeeper.trace.TMB_Trace getTrace() {
    return trace;
  }
  public void setTrace(org.apache.zookeeper.trace.TMB_Trace m_) {
    trace=m_;
  }
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    a_.writeLong(relativeZxid,"relativeZxid");
    {
      a_.startVector(dataWatches,"dataWatches");
      if (dataWatches!= null) {          int len1 = dataWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) dataWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(dataWatches,"dataWatches");
    }
    {
      a_.startVector(existWatches,"existWatches");
      if (existWatches!= null) {          int len1 = existWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) existWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(existWatches,"existWatches");
    }
    {
      a_.startVector(childWatches,"childWatches");
      if (childWatches!= null) {          int len1 = childWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) childWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(childWatches,"childWatches");
    }
    {
      a_.startVector(persistentWatches,"persistentWatches");
      if (persistentWatches!= null) {          int len1 = persistentWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) persistentWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(persistentWatches,"persistentWatches");
    }
    {
      a_.startVector(persistentRecursiveWatches,"persistentRecursiveWatches");
      if (persistentRecursiveWatches!= null) {          int len1 = persistentRecursiveWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) persistentRecursiveWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(persistentRecursiveWatches,"persistentRecursiveWatches");
    }
    a_.writeRecord(trace,"trace");
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    relativeZxid=a_.readLong("relativeZxid");
    {
      Index vidx1 = a_.startVector("dataWatches");
      if (vidx1!= null) {          dataWatches=new java.util.ArrayList<String>();
          for (; !vidx1.done(); vidx1.incr()) {
    String e1;
        e1=a_.readString("e1");
            dataWatches.add(e1);
          }
      }
    a_.endVector("dataWatches");
    }
    {
      Index vidx1 = a_.startVector("existWatches");
      if (vidx1!= null) {          existWatches=new java.util.ArrayList<String>();
          for (; !vidx1.done(); vidx1.incr()) {
    String e1;
        e1=a_.readString("e1");
            existWatches.add(e1);
          }
      }
    a_.endVector("existWatches");
    }
    {
      Index vidx1 = a_.startVector("childWatches");
      if (vidx1!= null) {          childWatches=new java.util.ArrayList<String>();
          for (; !vidx1.done(); vidx1.incr()) {
    String e1;
        e1=a_.readString("e1");
            childWatches.add(e1);
          }
      }
    a_.endVector("childWatches");
    }
    {
      Index vidx1 = a_.startVector("persistentWatches");
      if (vidx1!= null) {          persistentWatches=new java.util.ArrayList<String>();
          for (; !vidx1.done(); vidx1.incr()) {
    String e1;
        e1=a_.readString("e1");
            persistentWatches.add(e1);
          }
      }
    a_.endVector("persistentWatches");
    }
    {
      Index vidx1 = a_.startVector("persistentRecursiveWatches");
      if (vidx1!= null) {          persistentRecursiveWatches=new java.util.ArrayList<String>();
          for (; !vidx1.done(); vidx1.incr()) {
    String e1;
        e1=a_.readString("e1");
            persistentRecursiveWatches.add(e1);
          }
      }
    a_.endVector("persistentRecursiveWatches");
    }
    trace= new org.apache.zookeeper.trace.TMB_Trace();
    a_.readRecord(trace,"trace");
    a_.endRecord(tag);
}
  public String toString() {
    try {
      java.io.ByteArrayOutputStream s =
        new java.io.ByteArrayOutputStream();
      ToStringOutputArchive a_ = 
        new ToStringOutputArchive(s);
      a_.startRecord(this,"");
    a_.writeLong(relativeZxid,"relativeZxid");
    {
      a_.startVector(dataWatches,"dataWatches");
      if (dataWatches!= null) {          int len1 = dataWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) dataWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(dataWatches,"dataWatches");
    }
    {
      a_.startVector(existWatches,"existWatches");
      if (existWatches!= null) {          int len1 = existWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) existWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(existWatches,"existWatches");
    }
    {
      a_.startVector(childWatches,"childWatches");
      if (childWatches!= null) {          int len1 = childWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) childWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(childWatches,"childWatches");
    }
    {
      a_.startVector(persistentWatches,"persistentWatches");
      if (persistentWatches!= null) {          int len1 = persistentWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) persistentWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(persistentWatches,"persistentWatches");
    }
    {
      a_.startVector(persistentRecursiveWatches,"persistentRecursiveWatches");
      if (persistentRecursiveWatches!= null) {          int len1 = persistentRecursiveWatches.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) persistentRecursiveWatches.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(persistentRecursiveWatches,"persistentRecursiveWatches");
    }
    a_.writeRecord(trace,"trace");
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
    throw new UnsupportedOperationException("comparing SetWatches2 is unimplemented");
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof SetWatches2)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    SetWatches2 peer = (SetWatches2) peer_;
    boolean ret = false;
    ret = (relativeZxid==peer.relativeZxid);
    if (!ret) return ret;
    ret = dataWatches.equals(peer.dataWatches);
    if (!ret) return ret;
    ret = existWatches.equals(peer.existWatches);
    if (!ret) return ret;
    ret = childWatches.equals(peer.childWatches);
    if (!ret) return ret;
    ret = persistentWatches.equals(peer.persistentWatches);
    if (!ret) return ret;
    ret = persistentRecursiveWatches.equals(peer.persistentRecursiveWatches);
    if (!ret) return ret;
    ret = trace.equals(peer.trace);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = (int) (relativeZxid^(relativeZxid>>>32));
    result = 37*result + ret;
    ret = dataWatches.hashCode();
    result = 37*result + ret;
    ret = existWatches.hashCode();
    result = 37*result + ret;
    ret = childWatches.hashCode();
    result = 37*result + ret;
    ret = persistentWatches.hashCode();
    result = 37*result + ret;
    ret = persistentRecursiveWatches.hashCode();
    result = 37*result + ret;
    ret = trace.hashCode();
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LSetWatches2(l[s][s][s][s][s]LTMB_Trace(l[LTMB_Event(ilsss)][LTMB_TFI(isl[LTMB_TFIMeta(sll)])]))";
  }
}