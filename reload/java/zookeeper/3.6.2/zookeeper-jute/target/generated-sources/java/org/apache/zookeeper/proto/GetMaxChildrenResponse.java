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
import org.apache.zookeeper.trace.Trace;
@InterfaceAudience.Public
public class GetMaxChildrenResponse implements Record {
  private int max;
  private org.apache.zookeeper.trace.Trace trace;
  public GetMaxChildrenResponse() {
  }
  public GetMaxChildrenResponse(
        int max) {
    this.max=max;
  }
  public int getMax() {
    return max;
  }
  public void setMax(int m_) {
    max=m_;
  }
  public org.apache.zookeeper.trace.Trace getTrace() { return trace; }
  public void setTrace(org.apache.zookeeper.trace.Trace t_) { trace = t_; }
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    a_.writeInt(max,"max");
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    max=a_.readInt("max");
    a_.endRecord(tag);
}
  public String toString() {
    try {
      java.io.ByteArrayOutputStream s =
        new java.io.ByteArrayOutputStream();
      ToStringOutputArchive a_ = 
        new ToStringOutputArchive(s);
      a_.startRecord(this,"");
    a_.writeInt(max,"max");
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
    if (!(peer_ instanceof GetMaxChildrenResponse)) {
      throw new ClassCastException("Comparing different types of records.");
    }
    GetMaxChildrenResponse peer = (GetMaxChildrenResponse) peer_;
    int ret = 0;
    ret = (max == peer.max)? 0 :((max<peer.max)?-1:1);
    if (ret != 0) return ret;
     return ret;
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof GetMaxChildrenResponse)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    GetMaxChildrenResponse peer = (GetMaxChildrenResponse) peer_;
    boolean ret = false;
    ret = (max==peer.max);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = (int)max;
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LGetMaxChildrenResponse(i)";
  }
}
