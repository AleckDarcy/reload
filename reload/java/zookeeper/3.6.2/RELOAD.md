reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/ClientCnxn.java
    ClientCnxn::submitRequest()

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/FollowerRequestProcessor.java
    FollowerRequestProcessor::run()
        Org:
            Follower forwards request to leader without deserialization
        3MileBeach:
            Deserializes request, appends event to trace, serializes request
            Modified cases: OpCode.create

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/SendAckRequestProcessor.java
    SendAckRequestProcessor::processRequest()
        Org:
            Follower sends ACK messages to leader without data
        3MileBeach:
            Add data to ACK messages

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/LearnerHandler.java
    LearnerHandler::run()
        Org:
            Receives messages from peers
        3MileBeach:
            Modified cases: Leader.ACK

reload/java/zookeeper/3.6.2/zookeeper-server/src/main/java/org/apache/zookeeper/server/quorum/Leader.java
    Leader::sendPacket()
        Org:
            Sends packets to followers
        3MileBeach:
            Deserializes request, appends event to trace, serializes request
