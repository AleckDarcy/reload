#Prep

1. Install: Java 8, mvn

2. execute
```shell
mvn install -Dmaven.test.skip=true
```
and
```shell
mvn install -Dmaven.test.failure.ignore=true
# stop it when tests are running
```

#Demo test

Location: zookeeper-server/src/test/java/org/apache/zookeeper/server/quorum/QuorumPeerMainTest.Java QuorumPeerMainTest::test3MileBeach()

1) Leader election

Search "new LeaderZookeeperServer", there will be a few logs saying "[TMB_Store] [quorum-..." from the above.

```text
[3MileBeach] TMB_Store.java:128 [1] [TMB_Store] [quorum-1645607776] 1: 1, {"id":1,"req_event":0,"events":[{"service":"quorum-59903575","timestamp":1647116482502871416,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000003","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510118125,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000003","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510973416,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000014","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482511244541,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000014","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482511901500,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000016","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482511974791,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000017","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482512042333,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000018","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482512094583,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000017","processor":"FastLeaderElection"}],"tfis":[]}
[3MileBeach] TMB_Store.java:128 [2] [TMB_Store] [quorum-59903575] 1: 1, {"id":1,"req_event":0,"events":[{"service":"quorum-59903575","timestamp":1647116482502871416,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000003","processor":"FastLeaderElection"},{"service":"quorum-59903575","timestamp":1647116482506834000,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000006","processor":"FastLeaderElection"},{"service":"quorum-59903575","timestamp":1647116482507104083,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000009","processor":"FastLeaderElection"},{"service":"quorum-59903575","timestamp":1647116482507238750,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000009","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510118125,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000003","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482511026458,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000015","processor":"FastLeaderElection"},{"service":"quorum-59903575","timestamp":1647116482511420083,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000015","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510973416,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000014","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482511244541,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000014","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482512042333,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000018","processor":"FastLeaderElection"},{"service":"quorum-59903575","timestamp":1647116482512256916,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000018","processor":"FastLeaderElection"}],"tfis":[]}
[3MileBeach] TMB_Store.java:128 [0] [TMB_Store] [quorum-363400259] 1: 1, {"id":1,"req_event":0,"events":[{"service":"quorum-59903575","timestamp":1647116482502871416,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000003","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510118125,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000003","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510907416,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000013","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510973416,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000014","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482511026458,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000015","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482511042166,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000013","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482511244541,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000014","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482511901500,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000016","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482512183958,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000016","processor":"FastLeaderElection"}],"tfis":[]}
[3MileBeach] TMB_Store.java:128 [1] [TMB_Store] [quorum-1645607776] 3: 1, {"id":3,"req_event":0,"events":[{"service":"quorum-1645607776","timestamp":1647116482502845958,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000002","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482506791791,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000005","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482506991250,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000008","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482507080083,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000005","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482509784083,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000002","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510427958,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000011","processor":"FastLeaderElection"},{"service":"quorum-1645607776","timestamp":1647116482510591625,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000011","processor":"FastLeaderElection"}],"tfis":[]}
[3MileBeach] TMB_Store.java:128 [0] [TMB_Store] [quorum-363400259] 2: 1, {"id":2,"req_event":0,"events":[{"service":"quorum-363400259","timestamp":1647116482502845958,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000001","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482504289166,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000004","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482504790666,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000001","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482506877833,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000007","processor":"FastLeaderElection"}],"tfis":[]}
[3MileBeach] TMB_Store.java:128 [0] [TMB_Store] [quorum-363400259] 3: 1, {"id":3,"req_event":0,"events":[{"service":"quorum-1645607776","timestamp":1647116482502845958,"type":"SERVICE_SEND","message_name":"LeaderElection(New)","uuid":"0000000002","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482509784083,"type":"SERVICE_RECV","message_name":"LeaderElection(New)","uuid":"0000000002","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510361625,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000010","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510427958,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000011","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510472000,"type":"SERVICE_RECV","message_name":"LeaderElection(Vote)","uuid":"0000000010","processor":"FastLeaderElection"},{"service":"quorum-363400259","timestamp":1647116482510774458,"type":"SERVICE_SEND","message_name":"LeaderElection(Vote)","uuid":"0000000012","processor":"FastLeaderElection"}],"tfis":[]}
[3MileBeach] FollowerZooKeeperServer.java:64 [0] [quorum-363400259] new FollowerZookeeperServer
[3MileBeach] LeaderZooKeeperServer.java:59 [2] [quorum-59903575] new LeaderZookeeperServer
[3MileBeach] FollowerZooKeeperServer.java:64 [1] [quorum-1645607776] new FollowerZookeeperServer
```

Quorum-59903575 is the leader, we can find the trace of leader election printed by "quorum-59903575" (2nd line).

2) Basic user operation

To enable fault injection, uncomment:
```java
client.TMB_ClientInitialize(new TMB_Trace(TMB_Helper.newTraceId(), 0, new ArrayList<>(), new ArrayList<TMB_TFI>(){{
//  add(tfi); // uncomment this line
}
}));
```
and uncomment one of the *tfi*s from the above.

If the test passed, search "capture trace" for traces.

```text
[3MileBeach] QuorumPeerMainTest.java:264 [1] client-762218386 capture trace: {"id":11,"req_event":0,"events":[{"service":"client-1504109396","timestamp":1647116491532586500,"type":"SERVICE_SEND","message_name":"CreateRequest","uuid":"0000000034","processor":"TMB_ClientPlugin"},{"service":"quorum-792674816","timestamp":1647116491535940750,"type":"SERVICE_RECV","message_name":"CreateRequest","uuid":"0000000034","processor":"FollowerRequestProcessor"},{"service":"quorum-792674816","timestamp":1647116491535966208,"type":"SERVICE_FRWD","message_name":"CreateRequest","uuid":"0000000035","processor":"FollowerRequestProcessor"},{"service":"quorum-370529458","timestamp":1647116491539319041,"type":"SERVICE_RECV","message_name":"CreateRequest","uuid":"0000000035","processor":"PrepRequestProcessor"},{"service":"quorum-370529458","timestamp":1647116491540387291,"type":"LOGICAL_PRPS_READY","message_name":"LeaderProposeReady","uuid":"0000000036","processor":"Leader"},{"service":"quorum-370529458","timestamp":1647116491540694041,"type":"SERVICE_PRPS","message_name":"CreateRequest","uuid":"0000000036-0000","processor":"Leader"},{"service":"quorum-370529458","timestamp":1647116491540867250,"type":"SERVICE_PRPS","message_name":"CreateRequest","uuid":"0000000036-0001","processor":"Leader"},{"service":"quorum-370529458","timestamp":1647116491541816458,"type":"SERVICE_RECV","message_name":"QuorumAck","uuid":"0000000036-FFFF","processor":"AckRequestProcessor"},{"service":"quorum-370529458","timestamp":1647116491541914916,"type":"PROCESSOR_RECV","message_name":"QuorumAck","uuid":"0000000036-FFFF","processor":"Leader"},{"service":"quorum-1354099948","timestamp":1647116491541609666,"type":"SERVICE_RECV","message_name":"CreateRequest","uuid":"0000000036-0000","processor":"Follower"},{"service":"quorum-1354099948","timestamp":1647116491542522291,"type":"SERVICE_SEND","message_name":"QuorumAck","uuid":"0000000036-0000","processor":"SendAckRequestProcessor"},{"service":"quorum-370529458","timestamp":1647116491542889500,"type":"SERVICE_RECV","message_name":"QuorumAck","uuid":"0000000036-0000","processor":"LearnerHandler"},{"service":"quorum-792674816","timestamp":1647116491541740500,"type":"SERVICE_RECV","message_name":"CreateRequest","uuid":"0000000036-0001","processor":"Follower"},{"service":"quorum-792674816","timestamp":1647116491542705208,"type":"SERVICE_SEND","message_name":"QuorumAck","uuid":"0000000036-0001","processor":"SendAckRequestProcessor"},{"service":"quorum-370529458","timestamp":1647116491543942291,"type":"SERVICE_RECV","message_name":"QuorumAck","uuid":"0000000036-0001","processor":"LearnerHandler"},{"service":"quorum-370529458","timestamp":1647116491543026041,"type":"PROCESSOR_RECV","message_name":"QuorumAck","uuid":"0000000036-0000","processor":"Leader"},{"service":"quorum-370529458","timestamp":1647116491543332250,"type":"LOGICAL_CMMT_READY","message_name":"LeaderCommitReady","uuid":"0000000037","processor":"Leader"},{"service":"quorum-370529458","timestamp":1647116491543411541,"type":"SERVICE_SEND","message_name":"LeaderCommit","uuid":"0000000037-0000","processor":"Leader"},{"service":"quorum-370529458","timestamp":1647116491543472875,"type":"SERVICE_SEND","message_name":"LeaderCommit","uuid":"0000000037-0001","processor":"Leader"},{"service":"quorum-792674816","timestamp":1647116491543760583,"type":"SERVICE_RECV","message_name":"LeaderCommit","uuid":"0000000037-0001","processor":"Follower"},{"service":"quorum-792674816","timestamp":1647116491543906000,"type":"PROCESSOR_RECV","message_name":"CreateTxn","uuid":"0000000034","processor":"FinalRequestProcessor"},{"service":"quorum-792674816","timestamp":1647116491545788458,"type":"SERVICE_SEND","message_name":"CreateResponse","uuid":"0000000034","processor":"FinalRequestProcessor"},{"service":"client-1504109396","timestamp":1647116491546669958,"type":"SERVICE_RECV","message_name":"CreateResponse","uuid":"0000000034","processor":"TMB_ClientPlugin"}],"tfis":[]}
```

#Modifications

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/ClientCnxn.java
    ClientCnxn::submitRequest()

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/FollowerRequestProcessor.java
    FollowerRequestProcessor::run()
        Org:
            Follower forwards request to leader without deserialization
        3MileBeach:
            Deserializes request, appends event to trace, serializes request
            Modified cases:
                OpCode.create
                OpCode.create2
                OpCode.createTTL
                OpCode.createContainer
                OpCode.delete
                OpCode.deleteContainer
                OpCode.setData
                OpCode.reconfig
                OpCode.setACL
                OpCode.multi
                OpCode.check

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/SendAckRequestProcessor.java
    SendAckRequestProcessor::processRequest()
        Org:
            Follower sends ACK messages to leader without data
        3MileBeach:
            Add data to ACK messages
        TODO:
            Inject faults

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/LearnerHandler.java
    LearnerHandler::run()
        Org:
            Receives messages from peers
        3MileBeach:
            Modified cases: Leader.ACK
            
reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/Leader.java
    Leader::tryToCommit()
        Org:
            Leader sends COMMIT messages to followers without data
        3MileBeach:
            Add data to COMMIT messages
            Inject faults:
                Block commit messages to followers and observers
                TODO:
                    Block CommitProcessor and pendingSyncs
    Leader::sendPacket()
        Org:
            Sends proposal packets to followers
        3MileBeach:
            Deserializes request, appends event to trace, serializes request

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/PrepRequestProcessor.java
    PrepRequestProcessor::pRequest2Txn()
        Org:
            Converts byte[] data to Record
        3MileBeach:
            Sets trace to Record
            Modified cases:
                OpCode.create
                OpCode.create2
                OpCode.createTTL
                OpCode.createContainer
                OpCode.delete
                OpCode.deleteContainer
                OpCode.setData
                OpCode.reconfig
                OpCode.setACL
                OpCode.createSession
                OpCode.check
