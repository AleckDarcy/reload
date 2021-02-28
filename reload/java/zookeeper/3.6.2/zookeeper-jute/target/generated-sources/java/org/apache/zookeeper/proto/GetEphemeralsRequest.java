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
public class GetEphemeralsRequest implements Record {
  private String prefixPath;
  private org.apache.zookeeper.trace.TMB_Trace trace;
  public GetEphemeralsRequest() {
    this.trace = new org.apache.zookeeper.trace.TMB_Trace();
  }
  public GetEphemeralsRequest(
        String prefixPath) {
    this.prefixPath=prefixPath;
    this.trace = new org.apache.zookeeper.trace.TMB_Trace();
  }
  public String getPrefixPath() {
    return prefixPath;
  }
  public void setPrefixPath(String m_) {
    prefixPath=m_;
  }
  public org.apache.zookeeper.trace.TMB_Trace getTrace() {
    return trace;
  }
  public void setTrace(org.apache.zookeeper.trace.TMB_Trace m_) {
    trace=m_;
  }
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    a_.writeString(prefixPath,"prefixPath");
    a_.writeRecord(trace,"trace");
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    prefixPath=a_.readString("prefixPath");
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
    a_.writeString(prefixPath,"prefixPath");
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
    if (!(peer_ instanceof GetEphemeralsRequest)) {
      throw new ClassCastException("Comparing different types of records.");
    }
    GetEphemeralsRequest peer = (GetEphemeralsRequest) peer_;
    int ret = 0;
    ret = prefixPath.compareTo(peer.prefixPath);
    if (ret != 0) return ret;
     return ret;
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof GetEphemeralsRequest)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    GetEphemeralsRequest peer = (GetEphemeralsRequest) peer_;
    boolean ret = false;
    ret = prefixPath.equals(peer.prefixPath);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = prefixPath.hashCode();
    result = 37*result + ret;
    ret = trace.hashCode();
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LGetEphemeralsRequest(sLTMB_Trace(l[LTMB_Event(ilsss)][LTMB_TFI(isl[LTMB_TFIMeta(sll)])]))";
  }
}
