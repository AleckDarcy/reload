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
import org.apache.zookeeper.trace._3MB_Trace;
@InterfaceAudience.Public
public class GetChildrenResponse implements Record {
  private java.util.List<String> children;
  private org.apache.zookeeper.trace._3MB_Trace trace;
  public GetChildrenResponse() {
  }
  public GetChildrenResponse(
        java.util.List<String> children) {
    this.children=children;
  }
  public java.util.List<String> getChildren() {
    return children;
  }
  public void setChildren(java.util.List<String> m_) {
    children=m_;
  }
  public org.apache.zookeeper.trace._3MB_Trace getTrace() { return trace; }
  public void setTrace(org.apache.zookeeper.trace._3MB_Trace t_) { trace = t_; }
  public void serialize(OutputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(this,tag);
    {
      a_.startVector(children,"children");
      if (children!= null) {          int len1 = children.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) children.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(children,"children");
    }
    a_.endRecord(this,tag);
  }
  public void deserialize(InputArchive a_, String tag) throws java.io.IOException {
    a_.startRecord(tag);
    {
      Index vidx1 = a_.startVector("children");
      if (vidx1!= null) {          children=new java.util.ArrayList<String>();
          for (; !vidx1.done(); vidx1.incr()) {
    String e1;
        e1=a_.readString("e1");
            children.add(e1);
          }
      }
    a_.endVector("children");
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
    {
      a_.startVector(children,"children");
      if (children!= null) {          int len1 = children.size();
          for(int vidx1 = 0; vidx1<len1; vidx1++) {
            String e1 = (String) children.get(vidx1);
        a_.writeString(e1,"e1");
          }
      }
      a_.endVector(children,"children");
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
    throw new UnsupportedOperationException("comparing GetChildrenResponse is unimplemented");
  }
  public boolean equals(Object peer_) {
    if (!(peer_ instanceof GetChildrenResponse)) {
      return false;
    }
    if (peer_ == this) {
      return true;
    }
    GetChildrenResponse peer = (GetChildrenResponse) peer_;
    boolean ret = false;
    ret = children.equals(peer.children);
    if (!ret) return ret;
     return ret;
  }
  public int hashCode() {
    int result = 17;
    int ret;
    ret = children.hashCode();
    result = 37*result + ret;
    return result;
  }
  public static String signature() {
    return "LGetChildrenResponse([s])";
  }
}
