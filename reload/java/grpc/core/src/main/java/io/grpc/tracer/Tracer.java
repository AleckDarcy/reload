package io.grpc.tracer;

public interface Tracer {
    public java.lang.String GetFI_Name();

    public Message.Trace GetFI_Trace();
    public void SetFI_Trace(Message.Trace trace);

    public Message.MessageType GetFI_MessageType();
}
