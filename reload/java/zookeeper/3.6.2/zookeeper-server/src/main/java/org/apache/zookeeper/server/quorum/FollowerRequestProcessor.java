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

import java.io.IOException;
import java.util.concurrent.LinkedBlockingQueue;

import org.apache.jute.Record;
import org.apache.zookeeper.KeeperException;
import org.apache.zookeeper.MultiOperationRecord;
import org.apache.zookeeper.ZooDefs.OpCode;
import org.apache.zookeeper.proto.*;
import org.apache.zookeeper.server.*;
import org.apache.zookeeper.trace.TMB_Event;
import org.apache.zookeeper.trace.TMB_Store;
import org.apache.zookeeper.txn.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * This RequestProcessor forwards any requests that modify the state of the
 * system to the Leader.
 */
public class FollowerRequestProcessor extends ZooKeeperCriticalThread implements RequestProcessor {

    private static final Logger LOG = LoggerFactory.getLogger(FollowerRequestProcessor.class);

    FollowerZooKeeperServer zks;

    RequestProcessor nextProcessor;

    LinkedBlockingQueue<Request> queuedRequests = new LinkedBlockingQueue<Request>();

    boolean finished = false;

    // 3MileBeach starts
    TMB_Store.QuorumMeta quorumMeta;

    public FollowerRequestProcessor(FollowerZooKeeperServer zks, RequestProcessor nextProcessor, QuorumPeer self) {
        super("FollowerRequestProcessor:" + zks.getServerId(), zks.getZooKeeperServerListener());
        this.zks = zks;
        this.nextProcessor = nextProcessor;
        this.quorumMeta = self.getQuorumMeta();
    }
    // 3MileBeach ends

    public FollowerRequestProcessor(FollowerZooKeeperServer zks, RequestProcessor nextProcessor) {
        super("FollowerRequestProcessor:" + zks.getServerId(), zks.getZooKeeperServerListener());
        this.zks = zks;
        this.nextProcessor = nextProcessor;
        this.quorumMeta = new TMB_Store.QuorumMeta(0, "quorum-standalone"); // 3MileBeach
    }

    @Override
    public void run() {
        try {
            while (!finished) {
                Request request = queuedRequests.take();
                if (LOG.isTraceEnabled()) {
                    ZooTrace.logRequest(LOG, ZooTrace.CLIENT_REQUEST_TRACE_MASK, 'F', request, "");
                }
                if (request == Request.requestOfDeath) {
                    break;
                }

                // Screen quorum requests against ACLs first
                if (!zks.authWriteRequest(request)) {
                    continue;
                }
                TMB_Utils.printRequestForProcessor("FollowerRequestProcessor starts", quorumMeta, nextProcessor, request); // 3MileBeach

                // We want to queue the request to be processed before we submit
                // the request to the leader so that we are ready to receive
                // the response
                nextProcessor.processRequest(request);

                Record empty = null; // 3MileBeach
                // We now ship the request to the leader. As with all
                // other quorum operations, sync also follows this code
                // path, but different from others, we need to keep track
                // of the sync operations this follower has pending, so we
                // add it to pendingSyncs.
                switch (request.type) {
                case OpCode.sync:
                    zks.pendingSyncs.add(request);
                    zks.getFollower().request(request);
                    break;
                case OpCode.create:
                case OpCode.create2:
                case OpCode.createTTL:
                case OpCode.createContainer:
                    empty = new CreateRequest(); // 3MileBeach
                    break; // 3MileBeach
                case OpCode.delete:
                case OpCode.deleteContainer:
                    empty = new DeleteRequest(); // 3MileBeach
                    break; // 3MileBeach
                case OpCode.setData:
                    empty = new SetDataRequest(); // 3MileBeach
                    break; // 3MileBeach
                case OpCode.reconfig:
                    empty = new ReconfigRequest(); // 3MileBeach
                    break; // 3MileBeach
                case OpCode.setACL:
                    empty = new SetACLRequest(); // 3MileBeach
                    break; // 3MileBeach
                case OpCode.multi:
                    empty = new MultiOperationRecord(); // 3MileBeach
                    break; // 3MileBeach
                case OpCode.check:
                    empty = new CheckVersionRequest(); // 3MileBeach
                    // zks.getFollower().request(request);
                    break;
                case OpCode.createSession:
                case OpCode.closeSession:
                    // Don't forward local sessions to the leader.
                    if (!request.isLocalSession()) {
                        zks.getFollower().request(request);
                    }
                    break;
                }

                // 3MileBeach starts
                if (empty != null) {
                    request.request = TMB_Utils.appendEvent(request.request, empty, TMB_Event.RECORD_FRWD, quorumMeta, this.getClass());
                    zks.getFollower().request(request);
                }

                TMB_Utils.printRequestForProcessor("FollowerRequestProcessor ends", quorumMeta, nextProcessor, request);
                // 3MileBeach ends
            }
        } catch (Exception e) {
            handleException(this.getName(), e);
        }
        LOG.info("FollowerRequestProcessor exited loop!");
    }

    public void processRequest(Request request) {
        processRequest(request, true);
    }

    void processRequest(Request request, boolean checkForUpgrade) {
        if (!finished) {
            if (checkForUpgrade) {
                // Before sending the request, check if the request requires a
                // global session and what we have is a local session. If so do
                // an upgrade.
                Request upgradeRequest = null;
                try {
                    upgradeRequest = zks.checkUpgradeSession(request);
                } catch (KeeperException ke) {
                    if (request.getHdr() != null) {
                        request.getHdr().setType(OpCode.error);
                        request.setTxn(new ErrorTxn(ke.code().intValue()));
                    }
                    request.setException(ke);
                    LOG.warn("Error creating upgrade request", ke);
                } catch (IOException ie) {
                    LOG.error("Unexpected error in upgrade", ie);
                }
                if (upgradeRequest != null) {
                    queuedRequests.add(upgradeRequest);
                }
            }

            queuedRequests.add(request);
        }
    }

    public void shutdown() {
        LOG.info("Shutting down");
        finished = true;
        queuedRequests.clear();
        queuedRequests.add(Request.requestOfDeath);
        nextProcessor.shutdown();
    }

}
