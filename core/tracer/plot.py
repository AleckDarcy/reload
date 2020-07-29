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

    plt.plot(data["trace_off_throughputs"],data["trace_off_e2eLatencies"],'ro-',label='No Trace')
    plt.plot(data["trace_on_throughputs"],data["trace_on_e2eLatencies"],'g+-',label='3MileBeach')
    plt.plot(data["jaeger_on_throughputs"],data["jaeger_on_e2eLatencies"],'b^-',label='Jaeger')

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
    "trace_off_throughputs":[254,501,928,1331,1721,2106,2193,2563],
    "trace_off_e2eLatencies":[3,3,4,5,9,14,28,44],
    "trace_off_throughputs_error_bar":[3.292047,4.296512,11.492266,30.641450,49.686994,81.521936,78.662827,44.123972],
    "trace_off_e2eLatency_error_bar":[0.051042,0.033866,0.053447,0.143327,0.284052,0.586616,1.026366,0.680046],
    "trace_on_throughputs":[220,430,767,1117,1389,1636,1749,1801],
    "trace_on_e2eLatencies":[4,4,5,7,11,19,34,65],
    "trace_on_throughputs_error_bar":[2.040650,3.424259,8.811807,11.701441,44.472094,29.542293,25.108137,27.927665],
    "trace_on_e2eLatencies_error_bar":[0.040912,0.036917,0.059314,0.074871,0.387813,0.349088,0.671887,1.093445],
    "process_e2eLatencies":[4,4,5,7,12,22,43,78],
    "process_e2eLatencies_error_bar":[0.049379,0.040440,0.074217,0.206136,0.267050,0.211105,0.761292,1.859171],
    "process_feLatencies":[3,3,4,5,9,14,24,39],
    "process_feLatencies_error_bar":[0.040242,0.044431,0.067680,0.170299,0.300009,0.272257,0.625682,2.518285],
    "jaeger_on_throughputs":[194,396,753,1177,1472,1765,1447,1431],
    "jaeger_on_e2eLatencies":[5,5,5,6,11,17,42,82],
    "jaeger_on_throughputs_error_bar":[3.601900,5.433704,21.905864,41.866766,108.609070,71.050787,97.985333,46.959156],
    "jaeger_on_e2eLatencies_error_bar":[0.089715,0.066943,0.142326,0.232803,1.024864,0.661293,2.162755,3.447784],
    "trace_on_e2eLatencies_overhead":[14.796647,15.836994,20.512405,18.836493,23.462934,27.655561,23.642696,47.707126],
    "trace_on_throughputs_loss":[-13.407783,-14.135952,-17.332822,-16.128863,-19.269823,-22.298228,-20.225446,-29.713996],
    "jaeger_on_e2eLatencies_overhead":[30.826760,26.476884,23.705430,13.582274,20.388366,18.183713,49.736058,85.094918],
    "jaeger_on_throughputs_loss":[-23.454319,-20.970052,-18.890491,-11.568058,-14.498057,-16.186080,-33.973505,-44.151343]
}

# GKE 6 nodes (4 cores CPU, 15 GB Mem)
# data2={
#     "trace_off_throughputs":[85,170,284,386,480,543,579,574],
#     "trace_off_e2eLatencies":[11,11,14,20,33,58,106,208],
#     "trace_off_throughputs_error_bar":[1.285507,0.546539,1.384542,2.637761,1.145332,4.410817,4.940174,5.491252],
#     "trace_off_e2eLatency_error_bar":[0.189024,0.036097,0.069401,0.139820,0.082971,0.483050,1.072576,2.100476],
#     "trace_on_throughputs":[80,154,242,326,368,413,432,435],
#     "trace_on_e2eLatencies":[12,12,16,24,43,76,143,275],
#     "trace_on_throughputs_error_bar":[0.272481,0.306640,1.119529,1.384611,0.771429,1.813762,2.387579,2.413813],
#     "trace_on_e2eLatencies_error_bar":[0.041439,0.027938,0.075305,0.100775,0.081564,0.328951,0.869464,1.189746],
#     "process_e2eLatencies":[12,12,15,23,41,75,145,285],
#     "process_e2eLatencies_error_bar":[0.168876,0.027308,0.118206,0.153190,0.262906,0.638229,1.045647,2.074211],
#     "process_feLatencies":[10,10,13,21,39,71,140,276],
#     "process_feLatencies_error_bar":[0.153638,0.025707,0.119448,0.152503,0.227509,0.610930,1.165278,2.316837],
#     "jaeger_on_throughputs":[74,141,222,261,268,333,357,389],
#     "jaeger_on_e2eLatencies":[13,14,18,30,59,96,175,313],
#     "jaeger_on_throughputs_error_bar":[0.995998,0.957805,4.635125,7.371776,2.230132,10.943304,6.884534,7.270097],
#     "jaeger_on_e2eLatencies_error_bar":[0.179764,0.097339,0.347656,0.811587,0.480248,3.256228,3.522646,5.988417],
#     "trace_on_e2eLatencies_overhead":[4.516512,10.034411,17.016894,18.264788,30.342682,31.358779,34.273695,32.026263],
#     "trace_on_throughputs_loss":[-4.944488,-9.262465,-14.608704,-15.541638,-23.404625,-23.970023,-25.339855,-24.251672],
#     "jaeger_on_e2eLatencies_overhead":[14.517061,20.461734,28.723930,49.973338,79.670870,66.283562,64.257898,50.293130],
#     "jaeger_on_throughputs_loss":[-12.951954,-17.065928,-21.734721,-32.448727,-44.262004,-38.660014,-38.334506,-32.139368]
# }

# GKE 8 nodes (2 cores CPU, 7.5 GB Mem)
data3={
    "trace_off_throughputs":[345,637,1026,1539,2097,2194,2230,2279],
    "trace_off_e2eLatencies":[2,3,3,5,7,14,27,51],
    "trace_off_throughputs_error_bar":[6.113438,10.282098,22.637012,71.433653,43.683530,99.119998,131.269296,98.969117],
    "trace_off_e2eLatency_error_bar":[0.053929,0.049625,0.087620,0.275125,0.155694,0.671645,1.854459,2.260247],
    "trace_on_throughputs":[255,513,928,1260,1554,1920,1889,1799],
    "trace_on_e2eLatencies":[3,3,4,6,10,16,32,65],
    "trace_on_throughputs_error_bar":[14.063761,4.138063,18.634663,60.548929,87.750687,33.164986,87.069104,72.481657],
    "trace_on_e2eLatencies_error_bar":[0.246363,0.032064,0.087547,0.319129,0.646123,0.298002,1.716654,3.066713],
    "process_e2eLatencies":[3,3,5,6,13,24,51,103],
    "process_e2eLatencies_error_bar":[0.046190,0.082728,0.137688,0.196663,0.345981,0.254629,1.421758,3.028254],
    "process_feLatencies":[2,3,3,4,7,9,22,36],
    "process_feLatencies_error_bar":[0.043098,0.070368,0.107077,0.101121,0.283080,0.321214,2.083674,2.900708],
    "jaeger_on_throughputs":[241,474,758,973,1389,1490,1515,1756],
    "jaeger_on_e2eLatencies":[4,4,5,8,11,21,55,66],
    "jaeger_on_throughputs_error_bar":[9.883722,14.782260,23.853099,94.653495,27.121531,75.152297,250.450849,59.803926],
    "jaeger_on_e2eLatencies_error_bar":[0.193239,0.140068,0.166752,0.991319,0.233964,1.143997,18.217100,2.391525],
    "trace_on_e2eLatencies_overhead":[37.260219,23.658583,10.300693,21.714161,37.465460,14.510076,17.850449,26.353853],
    "trace_on_throughputs_loss":[-26.186860,-19.505822,-9.583320,-18.164860,-25.884380,-12.472398,-15.277263,-21.061447],
    "jaeger_on_e2eLatencies_overhead":[44.495747,34.882833,35.778529,65.394514,51.681231,49.683941,98.476377,28.724587],
    "jaeger_on_throughputs_loss":[-30.197367,-25.548634,-26.102451,-36.748717,-33.757772,-32.068271,-32.064077,-22.934882]
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

client_throughput(data1,0,"1 node","plots/client_throughput_1.pdf")
# client_throughput(data2,0,"6 nodes","plots/client_throughput_6.pdf")
client_throughput(data3,0,"8 nodes","plots/client_throughput_8.pdf")

client_latency(data1,0,"1 node","plots/client_latency_1.pdf")
# client_latency(data2,0,"6 nodes","plots/client_latency_6.pdf")
client_latency(data3,0,"8 nodes","plots/client_latency_8.pdf")

throughput_latency(data1,0,"1 node","plots/throughput_latency_1.pdf")
# throughput_latency(data2,0,"6 nodes","plots/throughput_latency_6.pdf")
throughput_latency(data3,0,"8 nodes","plots/throughput_latency_8.pdf")

client_throughput_loss(data1,0,"1 node","plots/client_throughput_loss_1.pdf")
# client_throughput_loss(data2,0,"6 nodes","plots/client_throughput_loss_6.pdf")
client_throughput_loss(data3,0,"8 nodes","plots/client_throughput_loss_8.pdf")

client_e2eLatency_overhead(data1,0,"1 node","plots/client_e2eLatency_overhead_1.pdf")
# client_e2eLatency_overhead(data2,0,"6 nodes","plots/client_e2eLatency_overhead_6.pdf")
client_e2eLatency_overhead(data3,0,"8 nodes","plots/client_e2eLatency_overhead_8.pdf")

client_process_latency(data1,0,"1 node","plots/client_process_latency_1.pdf")
# client_process_latency(data2,0,"6 nodes","plots/client_process_latency_6.pdf")
client_process_latency(data3,0,"8 nodes","plots/client_process_latency_8.pdf")
