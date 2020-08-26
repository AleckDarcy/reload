# coding:utf-8

import numpy as np
import matplotlib.pyplot as plt

def client_throughput(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    x1=[0.7,1.7,2.7,3.7,4.7,5.7,6.7,7.7]
    plt.bar(x1, data["trace_off_throughputs"],fc="r",width=width,label="No Trace",tick_label=name_list)
    plt.errorbar(x1,data["trace_off_throughputs"],fmt="none",ecolor="black",yerr=data["trace_off_throughputs_error_bar"])

    x2=[1,2,3,4,5,6,7,8]
    plt.bar(x2, data["trace_on_throughputs"],fc="g",width=width,label="3MileBeach",tick_label=name_list)
    plt.errorbar(x2,data["trace_on_throughputs"],fmt="none",ecolor="black",yerr=data["trace_on_throughputs_error_bar"])

    x3=[1.3,2.3,3.3,4.3,5.3,6.3,7.3,8.3]
    plt.bar(x3, data["jaeger_on_throughputs"],fc="b",width=width,label="Jaeger",tick_label=name_list)
    plt.errorbar(x3,data["jaeger_on_throughputs"],fmt="none",ecolor="black",yerr=data["jaeger_on_throughputs_error_bar"])

    plt.xlabel('# Client')
    plt.xticks(x2)
    plt.ylabel('Throughput(op/s)')

    if file != "":
        plt.legend()
        plt.savefig(file)
        plt.figure()
    else:
        plt.title(title)

def client_latency(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    x1=[0.7,1.7,2.7,3.7,4.7,5.7,6.7,7.7]
    plt.bar(x1, data["trace_off_e2eLatencies"],fc="r",width=width,label="No Trace",tick_label=name_list)
    plt.errorbar(x1,data["trace_off_e2eLatencies"],fmt="none",ecolor="black",yerr=data["trace_off_e2eLatency_error_bar"])

    x2=[1,2,3,4,5,6,7,8]
    plt.bar(x2, data["trace_on_e2eLatencies"],fc="g",width=width,label="3MileBeach",tick_label=name_list)
    plt.errorbar(x2,data["trace_on_e2eLatencies"],fmt="none",ecolor="black",yerr=data["trace_on_e2eLatencies_error_bar"])

    x3=[1.3,2.3,3.3,4.3,5.3,6.3,7.3,8.3]
    plt.bar(x3, data["jaeger_on_e2eLatencies"],fc="b",width=width,label="Jaeger",tick_label=name_list)
    plt.errorbar(x3,data["jaeger_on_e2eLatencies"],fmt="none",ecolor="black",yerr=data["jaeger_on_e2eLatencies_error_bar"])

    plt.xlabel('# Client')
    plt.xticks(x2)
    plt.ylabel('Latency(ms)')

    if file != "":
        plt.legend()
        plt.savefig(file)
        plt.figure()
    else:
        plt.title(title)

def throughput_latency(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    print(data["trace_off_throughputs"])
    print(data["trace_off_throughputs_error_bar"])

    plt.errorbar(data["trace_off_throughputs"],data["trace_off_e2eLatencies"],
                 yerr=data["trace_off_e2eLatency_error_bar"],xerr=data["trace_off_throughputs_error_bar"],
                 fmt='r.-',ecolor='black',label='No Trace')
    plt.errorbar(data["trace_on_throughputs"],data["trace_on_e2eLatencies"],
             yerr=data["trace_on_e2eLatencies_error_bar"],xerr=data["trace_on_throughputs_error_bar"],
             fmt='g.-',ecolor='black',label='3MileBeach')
    plt.errorbar(data["jaeger_on_throughputs"],data["jaeger_on_e2eLatencies"],
             yerr=data["jaeger_on_e2eLatencies_error_bar"],xerr=data["jaeger_on_throughputs_error_bar"],
             fmt='b.-',ecolor='black',label='Jaeger')

    plt.xlabel('Throughput(op/s)')
    plt.ylabel('Latency(ms)')

    if file != "":
        plt.legend()
        plt.savefig(file)
        plt.figure()
    else:
        plt.title(title)

def client_throughput_loss(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    x2=[0.85,1.85,2.85,3.85,4.85,5.85,6.85,7.85]
    plt.bar(x2, data["trace_on_throughputs_loss"],fc="g",width=width,label="3MileBeach",tick_label=name_list)

    x3=[1.15,2.15,3.15,4.15,5.15,6.15,7.15,8.15]
    plt.bar(x3, data["jaeger_on_throughputs_loss"],fc="b",width=width,label="Jaeger",tick_label=name_list)

    plt.xlabel('# Client')
    plt.ylabel('Throughput Loss(%)')

    if file != "":
        plt.legend()
        plt.savefig(file)
        plt.figure()
    else:
        plt.title(title)

def client_e2eLatency_overhead(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    x2=[0.85,1.85,2.85,3.85,4.85,5.85,6.85,7.85]
    plt.bar(x2, data["trace_on_e2eLatencies_overhead"],fc="g",width=width,label="3MileBeach",tick_label=name_list)

    x3=[1.15,2.15,3.15,4.15,5.15,6.15,7.15,8.15]
    plt.bar(x3, data["jaeger_on_e2eLatencies_overhead"],fc="b",width=width,label="Jaeger",tick_label=name_list)

    plt.xlabel('# Client')
    plt.ylabel('E2E Latency Overhead(%)')

    if subplot==336:
        plt.legend(loc="upper right",labels=["3MileBeach","Jaeger"])

    if file != "":
        plt.legend()
        plt.savefig(file)
        plt.figure()
    else:
        plt.title(title)

def client_process_latency(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    x2=[0.85,1.85,2.85,3.85,4.85,5.85,6.85,7.85]
    plt.bar(x2, data["process_e2eLatencies"],fc="black",width=width,label="E2E Latency",tick_label=name_list)
    plt.errorbar(x2,data["process_e2eLatencies"],fmt="none",ecolor="black",yerr=data["process_e2eLatencies_error_bar"])

    x3=[1.15,2.15,3.15,4.15,5.15,6.15,7.15,8.15]
    plt.bar(x3, data["process_feLatencies"],fc="g",width=width,label="Process Latency",tick_label=name_list)
    plt.errorbar(x3,data["process_feLatencies"],fmt="none",ecolor="black",yerr=data["process_feLatencies_error_bar"])

    plt.xlabel('# Client')
    plt.ylabel('Latency(ms)')

    if subplot==339:
        plt.legend()

    if file != "":
        plt.legend()
        plt.savefig(file)
        plt.figure()
    else:
        plt.title(title)

# GKE 1 node (16 cores CPU,  GB Mem)
data1={
    "trace_off_throughputs":[321,698,1173,1655,2115,2420,2846,3101],
    "trace_off_e2eLatencies":[3,2,3,4,7,13,21,38],
    "trace_off_throughputs_error_bar":[12.628073,4.969167,19.476387,31.457585,51.387306,104.041066,39.031584,117.866471],
    "trace_off_e2eLatency_error_bar":[0.134979,0.019230,0.058308,0.095435,0.205318,0.618829,0.283309,1.836281],
    "trace_on_throughputs":[326,630,1025,1395,1781,1939,2121,2115],
    "trace_on_e2eLatencies":[3,3,3,5,8,16,29,54],
    "trace_on_throughputs_error_bar":[6.917360,4.287661,18.282140,22.434754,24.278297,67.586712,48.091394,39.288704],
    "trace_on_e2eLatencies_error_bar":[0.068488,0.021945,0.065998,0.094117,0.124015,0.627419,0.750622,1.336895],
    "process_e2eLatencies":[3,3,4,6,10,18,35,67],
    "process_e2eLatencies_error_bar":[0.017357,0.020096,0.082949,0.086671,0.196489,0.643649,0.110924,1.911390],
    "process_feLatencies":[2,2,3,4,7,12,20,32],
    "process_feLatencies_error_bar":[0.014761,0.018336,0.075746,0.064436,0.219597,0.641232,0.509715,1.577885],
    "jaeger_on_throughputs":[305,577,914,1291,1623,1677,1591,1185],
    "jaeger_on_e2eLatencies":[3,3,4,6,9,18,39,102],
    "jaeger_on_throughputs_error_bar":[1.864478,8.093197,7.179459,39.964485,48.761030,60.003804,105.331600,39.768894],
    "jaeger_on_e2eLatencies_error_bar":[0.019700,0.047679,0.035079,0.183898,0.302387,0.663611,2.542033,3.151264],
    "trace_on_e2eLatencies_overhead":[-2.610024,10.308382,13.893482,18.134608,18.703595,24.379007,35.427562,42.786232],
    "trace_on_throughputs_loss":[1.600399,-9.713173,-12.616272,-15.689904,-15.785435,-19.875306,-25.480644,-31.799839],
    "jaeger_on_e2eLatencies_overhead":[3.841933,20.557307,27.600992,27.839105,29.590759,42.869565,84.856698,167.431656],
    "jaeger_on_throughputs_loss":[-4.950880,-17.352461,-22.093349,-21.948664,-23.275313,-30.698806,-44.095637,-61.781602]
}

# GKE 8 nodes (2 cores CPU, 7.5 GB Mem)
data3={
    "trace_off_throughputs":[304,584,1008,1481,1790,2102,2191,2269],
    "trace_off_e2eLatencies":[3,3,3,5,8,15,28,50],
    "trace_off_throughputs_error_bar":[4.421816,6.534744,19.810174,44.195762,38.647658,126.596914,153.265018,47.488218],
    "trace_off_e2eLatency_error_bar":[0.047493,0.038734,0.075596,0.165479,0.192320,1.002471,2.427690,1.199647],
    "trace_on_throughputs":[279,519,886,1226,1518,1745,1663,1697],
    "trace_on_e2eLatencies":[3,3,4,6,10,18,36,69],
    "trace_on_throughputs_error_bar":[2.734937,7.639877,18.757106,30.730520,35.489215,107.064184,69.283500,57.417488],
    "trace_on_e2eLatencies_error_bar":[0.034375,0.055962,0.098029,0.175860,0.243026,1.426483,1.930305,2.727924],
    "process_e2eLatencies":[3,4,5,7,13,27,56,112],
    "process_e2eLatencies_error_bar":[0.046162,0.062444,0.085462,0.114093,0.268295,0.786349,0.629906,3.007378],
    "process_feLatencies":[3,3,3,4,6,11,22,32],
    "process_feLatencies_error_bar":[0.036371,0.054412,0.087543,0.080213,0.236487,1.404157,1.890498,2.461155],
    "jaeger_on_throughputs":[227,417,711,981,1203,1257,1180,1198],
    "jaeger_on_e2eLatencies":[4,4,5,8,13,25,53,98],
    "jaeger_on_throughputs_error_bar":[4.250812,11.396852,13.459501,20.225223,31.281583,56.122188,77.476744,29.871164],
    "jaeger_on_e2eLatencies_error_bar":[0.084432,0.134525,0.110976,0.176217,0.348444,1.218913,3.779545,2.828122],
    "trace_on_e2eLatencies_overhead":[8.290237,12.017934,13.430487,20.384088,17.685121,20.851156,28.625948,35.935835],
    "trace_on_throughputs_loss":[-8.027370,-11.024997,-12.104481,-17.239342,-15.193405,-16.950784,-24.106321,-25.214558],
    "jaeger_on_e2eLatencies_overhead":[33.287295,40.136966,41.533450,49.805435,48.991984,64.995231,85.767603,93.554357],
    "jaeger_on_throughputs_loss":[-25.136989,-28.545775,-29.520875,-33.760633,-32.769180,-40.157209,-46.138163,-47.202809]
}

# plt.figure(figsize=(12,9))
# client_throughput(data1,331,"1 node")
# client_throughput(data2,332,"6 nodes")
# client_throughput(data3,333,"8 nodes")
#
# client_latency(data1,334,"1 node")
# client_latency(data2,335,"6 nodes")
# client_latency(data3,336,"8 nodes")
#
# throughput_latency(data1,337,"1 node")
# throughput_latency(data2,338,"6 nodes")
# throughput_latency(data3,339,"8 nodes")
#
# plt.subplots_adjust(left=0.07,bottom=0.08,right=0.97,top=0.94,wspace=0.22,hspace=0.35)
# plt.legend(loc="upper right",labels=["No Trace","3MileBeach","Jaeger"])
# # plt.show()
#
# plt.savefig("plots.pdf")

# plt.figure(figsize=(12,9))
# client_throughput_loss(data1,331,"1 node")
# client_throughput_loss(data2,332,"6 nodes")
# client_throughput_loss(data3,333,"8 nodes")
#
# client_e2eLatency_overhead(data1,334,"1 node")
# client_e2eLatency_overhead(data2,335,"6 nodes")
# client_e2eLatency_overhead(data3,336,"8 nodes")
#
# client_process_latency(data1, 337, "1 node")
# client_process_latency(data2, 338, "6 nodes")
# client_process_latency(data3, 339, "8 nodes")
#
# plt.subplots_adjust(left=0.07,bottom=0.08,right=0.97,top=0.94,wspace=0.22,hspace=0.35)
# # plt.show()
#
# plt.savefig("overhead.pdf")

# client_throughput(data1,0,"1 node","plots/client_throughput_1.pdf")
# # client_throughput(data2,0,"6 nodes","plots/client_throughput_6.pdf")
# client_throughput(data3,0,"8 nodes","plots/client_throughput_8.pdf")
#
# client_latency(data1,0,"1 node","plots/client_latency_1.pdf")
# # client_latency(data2,0,"6 nodes","plots/client_latency_6.pdf")
# client_latency(data3,0,"8 nodes","plots/client_latency_8.pdf")

throughput_latency(data1,0,"1 node","plots/throughput_latency_1.pdf")
# throughput_latency(data2,0,"6 nodes","plots/throughput_latency_6.pdf")
throughput_latency(data3,0,"8 nodes","plots/throughput_latency_8.pdf")

# client_throughput_loss(data1,0,"1 node","plots/client_throughput_loss_1.pdf")
# # client_throughput_loss(data2,0,"6 nodes","plots/client_throughput_loss_6.pdf")
# client_throughput_loss(data3,0,"8 nodes","plots/client_throughput_loss_8.pdf")
#
# client_e2eLatency_overhead(data1,0,"1 node","plots/client_e2eLatency_overhead_1.pdf")
# # client_e2eLatency_overhead(data2,0,"6 nodes","plots/client_e2eLatency_overhead_6.pdf")
# client_e2eLatency_overhead(data3,0,"8 nodes","plots/client_e2eLatency_overhead_8.pdf")
#
# client_process_latency(data1,0,"1 node","plots/client_process_latency_1.pdf")
# # client_process_latency(data2,0,"6 nodes","plots/client_process_latency_6.pdf")
# client_process_latency(data3,0,"8 nodes","plots/client_process_latency_8.pdf")
