/*
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

package org.apache.zookeeper.server.quorum;

import org.apache.zookeeper.server.Request;
import org.apache.zookeeper.server.RequestProcessor;
import org.apache.zookeeper.server.ServerMetrics;
import org.apache.zookeeper.server.TMB_Utils;
import org.apache.zookeeper.trace.TMB_Helper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;

/**
 * This is a very simple RequestProcessor that simply forwards a request from a
 * previous stage to the leader as an ACK.
 */
class AckRequestProcessor implements RequestProcessor {

    private static final Logger LOG = LoggerFactory.getLogger(AckRequestProcessor.class);
    Leader leader;

    // 3MileBeach starts
    int quorumId;
    String quorumName;

    AckRequestProcessor(Leader leader, QuorumPeer self) {
        this.leader = leader;
        this.quorumId = self.hashCode();
        this.quorumName = String.format("quorum-%d", this.quorumId);
    }
    // 3MileBeach ends

    AckRequestProcessor(Leader leader) {
        this.leader = leader;
        this.quorumName = "quorum-standalone"; // 3MileBeach
    }

    /**
     * Forward the request as an ACK to the leader
     */
    public void processRequest(Request request) {
        QuorumPeer self = leader.self;
        TMB_Utils.printRequestForProcessor("AckRequestProcessor", quorumName, self, request); // 3MileBeach
        if (self != null) {
            request.logLatency(ServerMetrics.getMetrics().PROPOSAL_ACK_CREATION_LATENCY);
            leader.processAck(self.getId(), request.zxid, null);
        } else {
            LOG.error("Null QuorumPeer");
        }
    }

    public void shutdown() {
        // TODO No need to do anything
    }

}
