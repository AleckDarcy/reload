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
    Leader::sendPacket()
        Org:
            Sends packets to followers
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
