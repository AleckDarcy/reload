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

import java.io.Flushable;
import java.io.IOException;

import org.apache.zookeeper.ZooDefs.OpCode;
import org.apache.zookeeper.server.*;
import org.apache.zookeeper.trace.FaultInjectedException;
import org.apache.zookeeper.trace.TMB_Helper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class SendAckRequestProcessor implements RequestProcessor, Flushable {

    private static final Logger LOG = LoggerFactory.getLogger(SendAckRequestProcessor.class);

    Learner learner;
    // 3MileBeach starts
    int quorumId;
    String quorumName;

    SendAckRequestProcessor(Learner peer, QuorumPeer self) {
        this.learner = peer;
        this.quorumId = self.hashCode();
        this.quorumName = String.format("quorum-%d", this.quorumId);
    }
    // 3MileBeach ends

    SendAckRequestProcessor(Learner peer) {
        this.learner = peer;
        this.quorumName = "quorum-standalone"; // 3MileBeach
    }

    public void processRequest(Request si) {
        if (si.type != OpCode.sync) {
            // 3MileBeach starts
            TMB_Utils.printRequestForProcessor("SendAckRequestProcessor", quorumName, learner, si);
            byte[] data;
            try {
                TMB_Helper.printf("[%s] let's ack! request %s\n", quorumName, si.getTxn());
                data = TMB_Utils.ackHelper(si, TMB_Utils.QUORUM_ACK, quorumName, quorumId);
            } catch (FaultInjectedException e) {
                TMB_Helper.printf("[%s] Fault injected, won't send ACK to the leader\n", quorumName);
                return;
            }
            QuorumPacket qp = new QuorumPacket(Leader.ACK, si.getHdr().getZxid(), data, null);
            // 3MileBeach ends

//            QuorumPacket qp = new QuorumPacket(Leader.ACK, si.getHdr().getZxid(), null, null);
            try {
                si.logLatency(ServerMetrics.getMetrics().PROPOSAL_ACK_CREATION_LATENCY);

                learner.writePacket(qp, false);
            } catch (IOException e) {
                LOG.warn("Closing connection to leader, exception during packet send", e);
                try {
                    if (!learner.sock.isClosed()) {
                        learner.sock.close();
                    }
                } catch (IOException e1) {
                    // Nothing to do, we are shutting things down, so an exception here is irrelevant
                    LOG.debug("Ignoring error closing the connection", e1);
                }
            }
        } else {
            TMB_Helper.printf("[%s] SendAckRequestProcessor\n", quorumName); // 3MileBeach
        }
    }

    public void flush() throws IOException {
        try {
            learner.writePacket(null, true);
        } catch (IOException e) {
            LOG.warn("Closing connection to leader, exception during packet send", e);
            try {
                if (!learner.sock.isClosed()) {
                    learner.sock.close();
                }
            } catch (IOException e1) {
                // Nothing to do, we are shutting things down, so an exception here is irrelevant
                LOG.debug("Ignoring error closing the connection", e1);
            }
        }
    }

    public void shutdown() {
        // Nothing needed
    }

}
