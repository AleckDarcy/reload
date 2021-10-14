# coding:utf-8

import numpy as np
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc('pdf', fonttype=42)

# matplotlib.rcParams['pdf.fonttype'] = 42
# matplotlib.rcParams['ps.fonttype'] = 42

def client_throughput(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    plt.figure()

    x1=[0.7,1.7,2.7,3.7,4.7,5.7,6.7,7.7]
    plt.bar(x1, data["trace_off_throughputs"],fc="r",width=width,label="No Trace",tick_label=name_list)
    plt.errorbar(x1,data["trace_off_throughputs"],fmt="none",ecolor="black",yerr=data["trace_off_throughputs_error_bar"])

    x2=[1,2,3,4,5,6,7,8]
    plt.bar(x2, data["trace_on_throughputs"],fc="g",width=width,label="3MileBeach",tick_label=name_list)
    plt.errorbar(x2,data["trace_on_throughputs"],fmt="none",ecolor="black",yerr=data["trace_on_throughputs_error_bar"])

    x3=[1.3,2.3,3.3,4.3,5.3,6.3,7.3,8.3]
    plt.bar(x3, data["jaeger_on_throughputs"],fc="b",width=width,label="Jaeger",tick_label=name_list)
    plt.errorbar(x3,data["jaeger_on_throughputs"],fmt="none",ecolor="black",yerr=data["jaeger_on_throughputs_error_bar"])

    plt.xlabel('$N_{C}$')
    plt.xticks(x2)
    plt.ylabel('Throughput(op/s)')
    plt.subplots_adjust(left=0.13,right=0.96,top=0.95)

    if file != "":
        plt.legend()
        plt.savefig(file)
    else:
        plt.title(title)

def client_latency(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    plt.figure()

    x1=[0.7,1.7,2.7,3.7,4.7,5.7,6.7,7.7]
    plt.bar(x1, data["trace_off_e2eLatencies"],fc="r",width=width,label="No Trace",tick_label=name_list)
    plt.errorbar(x1,data["trace_off_e2eLatencies"],fmt="none",ecolor="black",yerr=data["trace_off_e2eLatency_error_bar"])

    x2=[1,2,3,4,5,6,7,8]
    plt.bar(x2, data["trace_on_e2eLatencies"],fc="g",width=width,label="3MileBeach",tick_label=name_list)
    plt.errorbar(x2,data["trace_on_e2eLatencies"],fmt="none",ecolor="black",yerr=data["trace_on_e2eLatencies_error_bar"])

    x3=[1.3,2.3,3.3,4.3,5.3,6.3,7.3,8.3]
    plt.bar(x3, data["jaeger_on_e2eLatencies"],fc="b",width=width,label="Jaeger",tick_label=name_list)
    plt.errorbar(x3,data["jaeger_on_e2eLatencies"],fmt="none",ecolor="black",yerr=data["jaeger_on_e2eLatencies_error_bar"])

    plt.xlabel('$N_{C}$')
    plt.xticks(x2)
    plt.ylabel('Latency(ms)')
    plt.subplots_adjust(left=0.11,right=0.96,top=0.95)

    if file != "":
        plt.savefig(file)
    else:
        plt.title(title)

def throughput_latency_all(cluster1,cluster2,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    plt.figure()

    plt.errorbar(cluster1["trace_off_throughputs"],cluster1["trace_off_e2eLatencies"],
                 yerr=cluster1["trace_off_e2eLatency_error_bar"],xerr=cluster1["trace_off_throughputs_error_bar"],
                 fmt='r.-',ecolor='black',label='Baseline (Cluster1)')
    plt.errorbar(cluster1["trace_on_throughputs"],cluster1["trace_on_e2eLatencies"],
                 yerr=cluster1["trace_on_e2eLatencies_error_bar"],xerr=cluster1["trace_on_throughputs_error_bar"],
                 fmt='g.-',ecolor='black',label='3MileBeach (Cluster1)')
    plt.errorbar(cluster1["jaeger_on_throughputs"],cluster1["jaeger_on_e2eLatencies"],
                 yerr=cluster1["jaeger_on_e2eLatencies_error_bar"],xerr=cluster1["jaeger_on_throughputs_error_bar"],
                 fmt='b.-',ecolor='black',label='Jaeger (Cluster1)')

    plt.errorbar(cluster2["trace_off_throughputs"],cluster2["trace_off_e2eLatencies"],
                 yerr=cluster2["trace_off_e2eLatency_error_bar"],xerr=cluster2["trace_off_throughputs_error_bar"],
                 fmt='r.--',ecolor='black',label='Baseline (Cluster2)')
    plt.errorbar(cluster2["trace_on_throughputs"],cluster2["trace_on_e2eLatencies"],
                 yerr=cluster2["trace_on_e2eLatencies_error_bar"],xerr=cluster1["trace_on_throughputs_error_bar"],
                 fmt='g.--',ecolor='black',label='3MileBeach (Cluster2)')
    plt.errorbar(cluster2["jaeger_on_throughputs"],cluster2["jaeger_on_e2eLatencies"],
                 yerr=cluster2["jaeger_on_e2eLatencies_error_bar"],xerr=cluster2["jaeger_on_throughputs_error_bar"],
                 fmt='b.--',ecolor='black',label='Jaeger (Cluster2)')

    plt.xlabel('Throughput(op/s)')
    plt.ylabel('Latency(ms)')
    plt.subplots_adjust(left=0.11,right=0.96,top=0.95)

    if file != "":
        plt.legend()
        plt.savefig(file)
    else:
        plt.title(title)

def throughput_latency(data,subplot,title,legend,file):
    if subplot != 0:
        plt.subplot(subplot)

    plt.figure()
    # plt.figure(figsize=(3.6,2.7))

    baseline_tp,baseline_lat=data["trace_off_throughputs"],data["trace_off_e2eLatencies"]
    milebeach_tp,milebeach_lat=data["trace_on_throughputs"],data["trace_on_e2eLatencies"]
    jaeger_tp,jaeger_lat=data["jaeger_on_throughputs"],data["jaeger_on_e2eLatencies"]

    plt.errorbar(baseline_tp,baseline_lat,
                 yerr=data["trace_off_e2eLatency_error_bar"],xerr=data["trace_off_throughputs_error_bar"],
                 fmt='r.-',ecolor='black',label='TraceOff')
    plt.errorbar(milebeach_tp,milebeach_lat,
             yerr=data["trace_on_e2eLatencies_error_bar"],xerr=data["trace_on_throughputs_error_bar"],
             fmt='g.-',ecolor='black',label='3MileBeachOn')
    plt.errorbar(jaeger_tp,jaeger_lat,
             yerr=data["jaeger_on_e2eLatencies_error_bar"],xerr=data["jaeger_on_throughputs_error_bar"],
             fmt='b.-',ecolor='black',label='JaegerOn')

    for i in (6,7):
        milebeach_o=100*(milebeach_lat[i]/baseline_lat[i]-1)
        milebeach_l=100*(milebeach_tp[i]/baseline_tp[i]-1)

        jaeger_o=100*(jaeger_lat[i]/baseline_lat[i]-1)
        jaeger_l=100*(jaeger_tp[i]/baseline_tp[i]-1)

        plt.annotate("(%0.1f%%)"%milebeach_o, (milebeach_tp[i],milebeach_lat[i]))
        plt.annotate("(%0.1f%%)"%jaeger_o, (jaeger_tp[i],jaeger_lat[i]))

        # plt.annotate("<%0.1f%%,%0.1f%%>"%(milebeach_o,milebeach_l), (milebeach_tp[i],milebeach_lat[i]))
        # plt.annotate("<%0.1f%%,%0.1f%%>"%(jaeger_o,jaeger_l), (jaeger_tp[i],jaeger_lat[i]))

    plt.xlabel('Throughput(op/s)')
    plt.xlim(xmin=0,xmax=3500)
    plt.ylim(ymin=0,ymax=120)
    plt.ylabel('Latency(ms)')
    # plt.subplots_adjust(left=0.16,right=0.97,top=0.97,bottom=0.16)

    if legend:
        plt.legend()

    if file != "":
        plt.savefig(file)
        # plt.show()
    else:
        plt.title(title)

def client_throughput_loss(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    plt.figure()

    x2=[0.85,1.85,2.85,3.85,4.85,5.85,6.85,7.85]
    plt.bar(x2, data["trace_on_throughputs_loss"],fc="g",width=width,label="3MileBeach",tick_label=name_list)

    x3=[1.15,2.15,3.15,4.15,5.15,6.15,7.15,8.15]
    plt.bar(x3, data["jaeger_on_throughputs_loss"],fc="b",width=width,label="Jaeger",tick_label=name_list)

    plt.xlabel('$N_{C}$')
    plt.ylabel('Throughput Loss(%)')
    plt.subplots_adjust(left=0.09,right=0.96,top=0.95)

    if file != "":
        plt.legend()
        plt.savefig(file)
    else:
        plt.title(title)

def client_e2eLatency_overhead(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    plt.figure()

    x2=[0.85,1.85,2.85,3.85,4.85,5.85,6.85,7.85]
    plt.bar(x2, data["trace_on_e2eLatencies_overhead"],fc="g",width=width,label="3MileBeach",tick_label=name_list)

    x3=[1.15,2.15,3.15,4.15,5.15,6.15,7.15,8.15]
    plt.bar(x3, data["jaeger_on_e2eLatencies_overhead"],fc="b",width=width,label="Jaeger",tick_label=name_list)

    plt.xlabel('$N_{C}$')
    plt.ylabel('E2E Latency Overhead(%)')
    plt.subplots_adjust(left=0.12,right=0.96,top=0.95)

    if subplot==336:
        plt.legend(loc="upper right",labels=["3MileBeach","Jaeger"])

    if file != "":
        plt.legend()
        plt.savefig(file)
    else:
        plt.title(title)

def client_process_latency(data,subplot,title,file):
    if subplot != 0:
        plt.subplot(subplot)

    name_list=[1,2,4,8,16,32,64,128]
    width = 0.3

    plt.figure()

    x2=[0.85,1.85,2.85,3.85,4.85,5.85,6.85,7.85]
    plt.bar(x2, data["process_e2eLatencies"],fc="black",width=width,label="E2E Latency",tick_label=name_list)
    plt.errorbar(x2,data["process_e2eLatencies"],fmt="none",ecolor="black",yerr=data["process_e2eLatencies_error_bar"])

    x3=[1.15,2.15,3.15,4.15,5.15,6.15,7.15,8.15]
    plt.bar(x3, data["process_feLatencies"],fc="g",width=width,label="Process Latency",tick_label=name_list)
    plt.errorbar(x3,data["process_feLatencies"],fmt="none",ecolor="black",yerr=data["process_feLatencies_error_bar"])

    plt.xlabel('$N_{C}$')
    plt.ylabel('Latency(ms)')
    plt.subplots_adjust(left=0.09,right=0.96,top=0.95)

    if subplot==339:
        plt.legend()

    if file != "":
        plt.legend()
        plt.savefig(file)
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


opt={
    # 0719, 1 node
    "before":{
        "e2e":{
            "data":[1.00,1.12,1.53,2.40,4.18,7.75,14.58,26.82],
            "T#":6,
            "color":".-",
        },
        "rt":{
            "data":[1.00,1.14,1.44,1.81,2.33,3.38,4.58,9.59],
            "T#":6,
            "color":"*-",
            # "color":"g*-",
        },
        "c1":{
            "data":[1.00,1.13,1.68,2.85,5.50,10.28,20.43,37.00],
            "T#":3,
            "color":"v-",
            # "color":"bv-",
        },
        "c2":{
            "data":[1.00,1.22,1.81,3.26,6.36,12.57,24.77,44.94],
            "T#":3,
            "color":"^-",
            # "color":"b^-",
        },
        "p1":{
            "data":[1.00,1.10,1.32,1.59,1.96,2.96,4.39,9.76],
            "T#":6,
            "color":"+-",
            # "color":"g+-",
        },
        "a1":{
            "data":[1.00,0.74,0.82,0.99,1.27,1.97,2.98,6.59],
            "T#":6,
            "color":"o-",
            # "color":"go-",
        },
    },
    # 0725, 1 node
    "after":{
        "e2e":{
            "data":[1.00,1.17,1.43,1.71,2.67,4.86,9.57,17.11],
            "T#":-1,
            "color":".-"
        },
        "rt":{
            "data":[1.00,1.21,1.58,2.15,3.49,8.70,24.78,52.16],
            "T#":4,
            "color":"*-",
            # "color":"b*-",
        },
        "c1":{
            "data":[1.00,1.25,1.52,1.90,3.32,5.98,9.89,16.96],
            "T#":-1,
            "color":"v-",
        },
        "c2":{
            "data":[1.00,1.22,1.51,1.97,3.46,6.10,10.76,17.99],
            "T#":-1,
            "color":"^-",
        },
        "p1":{
            "data":[1.00,1.14,1.33,1.56,2.21,3.15,4.91,7.40],
            "T#":-1,
            "color":"+-",
        },
        "a1":{
            "data":[1.00,1.08,1.32,1.37,2.09,3.48,5.35,6.73],
            "T#":-1,
            "color":"o-",
        },
    },
    # 0814, 8 nodes
    "cluster1":{
        "e2e":{
            "data":[1.00,1.03,1.28,1.83,3.32,6.91,14.20,28.28],
            "T#":4,
            "color":"k.-",
            # "color":"g.-",
        },
        "rt":{
            "data":[1.00,1.15,1.69,3.28,8.01,19.47,42.08,98.07],
            "T#":3,
            "color":"k*-",
            # "color":"b*-",
        },
        "c1":{
            "data":[1.00,0.97,1.22,1.60,2.50,4.00,9.33,14.05],
            "T#":-1,
            "color":"kv-",
        },
        "c2":{
            "data":[1.00,1.04,1.34,1.57,2.44,4.24,8.85,11.82],
            "T#":-1,
            "color":"k^-",
        },
        "p1":{
            "data":[1.00,1.04,1.20,1.47,1.89,3.65,6.76,10.25],
            "T#":-1,
            "color":"k+-",
        },
        "a1":{
            "data":[1.00,0.95,1.10,1.57,2.33,4.26,7.24,10.79],
            "T#":-1,
            "color":"ko-",
        },
    },
    # 0814, 1 node
    "cluster2":{
        "e2e":{
            "data":[1.00,1.07,1.39,2.02,3.35,6.27,11.80,22.20],
            "T#":-1,
            "color":"k.-",
        },
        "rt":{
            "data":[1.00,1.11,1.62,2.42,4.32,11.28,28.95,65.85],
            "T#":4,
            "color":"k*-",
            # "color":"b*-",
        },
        "c1":{
            "data":[1.00,1.07,1.47,2.22,4.09,7.16,12.27,21.79],
            "T#":-1,
            "color":"kv-",
        },
        "c2":{
            "data":[1.00,1.13,1.51,2.39,4.33,7.66,13.68,23.29],
            "T#":-1,
            "color":"k^-",
        },
        "p1":{
            "data":[1.00,1.03,1.27,1.82,2.71,4.08,5.56,7.54],
            "T#":-1,
            "color":"k+-",
        },
        "a1":{
            "data":[1.00,1.01,1.20,1.67,2.61,4.59,6.23,9.04],
            "T#":-1,
            "color":"ko-",
        },
    },
}

def tuning(opt,name,legend,file):
    x=[1,2,3,4,5,6,7]

    data=opt[name]

    e2e, rt, c1, c2, p1, a1 = data["e2e"],data["rt"],data["c1"],data["c2"],data["p1"],data["a1"]

    plt.figure()
    # plt.figure(figsize=(4.8,3.6))

    plt.plot(x, np.log2(e2e["data"][1:]),e2e["color"],label="E2E")
    plt.plot(x, np.log2(rt["data"][1:]),rt["color"],label="RT")
    plt.plot(x, np.log2(c1["data"][1:]),c1["color"],label="Req$_{C1}$")
    plt.plot(x, np.log2(c2["data"][1:]),c2["color"],label="Req$_{C2}$")
    plt.plot(x, np.log2(p1["data"][1:]),p1["color"],label="Req$_{P1}$")
    plt.plot(x, np.log2(a1["data"][1:]),a1["color"],label="Req$_{A1}$")

    # if e2e["T#"] != -1:
    #     plt.plot([e2e["T#"]],np.log2(e2e["data"])[e2e["T#"]],"r.--")
    # if rt["T#"] != -1:
    #     plt.plot([rt["T#"]],np.log2(rt["data"])[rt["T#"]],"r*--")
    # if c1["T#"] != -1:
    #     plt.plot([c1["T#"]],np.log2(c1["data"])[c1["T#"]],"rv--")
    # if c2["T#"] != -1:
    #     plt.plot([c2["T#"]],np.log2(c2["data"])[c2["T#"]],"r^--")
    # if p1["T#"] != -1:
    #     plt.plot([p1["T#"]],np.log2(p1["data"])[p1["T#"]],"r+--")
    # if a1["T#"] != -1:
    #     plt.plot([a1["T#"]],np.log2(a1["data"])[a1["T#"]],"ro--")

    plt.ylim(ymin=-1,ymax=7)
    plt.xlabel('log$_{2}$(N$_{C}$)')
    plt.ylabel('log$_{2}$(R$_{N_{C}}$)')
    # plt.subplots_adjust(left=0.12,right=0.99,top=0.98,bottom=0.13)

    if legend:
        plt.legend()

    if file != "":
        plt.savefig(file)
        # plt.show()
    else:
        plt.show()

# overhead and loss
o_l={
    "cluster1":{
        "overhead":{
            "case2":{
                "data":[8.3,12.0,13.4,20.4,17.7,20.9,28.6,35.9],
            },
            "case3":{
                "data":[33.3,40.1,41.5,49.8,49.0,65.0,85.8,93.6],
            },
        },
        "loss":{
            "case2":{
                "data":[-8.0,-11.0,-12.1,-17.2,-15.2,-17.0,-24.1,-25.2],
            },
            "case3":{
                "data":[-25.1,-28.5,-29.5,-33.8,-32.8,-40.2,-46.1,-47.2],
            },
        },
    },

    "cluster2":{
        "overhead":{
            "case2":{
                "data":[-2.6,10.3,13.9,18.1,18.7,24.4,35.4,42.8],
            },
            "case3":{
                "data":[3.8,20.6,27.6,27.8,29.6,42.9,84.9,167.4],
            },
        },
        "loss":{
            "case2":{
                "data":[1.6,-9.7,-12.6,-15.7,-15.8,-19.9,-25.5,-31.8],
            },
            "case3":{
                "data":[-5.0,-17.4,-22.1,-21.9,-23.3,-30.7,-44.1,-61.8],
            },
        },
    },
}

def overhead_and_loss(o_l,legend,file):
    x=[1,2,4,8,16,32,64,128]

    o,l=o_l["overhead"],o_l["loss"]

    plt.figure(figsize=(3.6,2.7))

    plt.plot(x,o["case2"]["data"],"r.-",label="Overhead (Case2)")
    plt.plot(x,o["case3"]["data"],"b.-",label="Overhead (Case3)")
    plt.plot(x,l["case2"]["data"],"r.--",label="Loss (Case2)")
    plt.plot(x,l["case3"]["data"],"b.--",label="Loss (Case3)")

    plt.ylim(ymin=-100,ymax=200)
    plt.xticks([1,4,8,16,32,64,128])
    plt.xlabel('$N_{C}$')
    plt.ylabel('Overhead and Loss(%)')
    plt.subplots_adjust(left=0.20,right=0.97,top=0.97,bottom=0.16)

    if legend:
        plt.legend()

    if file != "":
        plt.savefig(file)
        # plt.show()
    else:
        plt.show()

# client_throughput(data1,0,"1 node","plots/client_throughput_1.pdf")
# # client_throughput(data2,0,"6 nodes","plots/client_throughput_6.pdf")
# client_throughput(data3,0,"8 nodes","plots/client_throughput_8.pdf")
#
# client_latency(data1,0,"1 node","plots/client_latency_1.pdf")
# # client_latency(data2,0,"6 nodes","plots/client_latency_6.pdf")
# client_latency(data3,0,"8 nodes","plots/client_latency_8.pdf")

throughput_latency(data1,0,"1 node",False,"plots/throughput_latency_1.pdf")
# # throughput_latency(data2,0,"6 nodes","plots/throughput_latency_6.pdf")
throughput_latency(data3,0,"8 nodes",True,"plots/throughput_latency_8.pdf")

# throughput_latency_all(data3,data1,0,"8 nodes","plots/throughput_latency.pdf")

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

tuning(opt,"before",True,"plots/opt_before.pdf")
tuning(opt,"after",False,"plots/opt_after.pdf")
# tuning(opt,"cluster1",False,"plots/opt_cluster1.pdf")
# tuning(opt,"cluster2",False,"plots/opt_cluster2.pdf")

# overhead_and_loss(o_l["cluster1"],True,"plots/overhead_and_loss_8.pdf")
# overhead_and_loss(o_l["cluster2"],False,"plots/overhead_and_loss_1.pdf")

t_c={
    "before":{
        "throughput":{
            "mean":[125.69,225.08,330.64,419.52,490.51,527.01,544.98,578.52],
            "err":[0.996381,0.666127,1.640500,2.650524,2.818208,4.639243,2.109686,4.821631],
        },
        "e2eLatency":{
            "mean":[7.95,8.86,12.06,18.98,32.38,59.88,113.58,209.20],
            "err":[0.065898,0.026320,0.059825,0.123416,0.191008,0.565041,0.491299,2.492250],
        }
    },
    "after":{
        "throughput":{
            "mean":[186.72,412.11,766.74,1045.93,1196.84,1204.90,1253.71,1279.58],
            "err":[2.315363,20.804326,6.246531,11.626024,9.464491,14.528413,14.651856,7.103380],
        },
        "e2eLatency":{
            "mean":[5.32,4.88,5.18,7.58,13.19,25.94,48.21,90.35],
            "err":[0.068172,0.258526,0.040829,0.081321,0.093115,0.218108,0.159742,0.840485],
        }
    },
    "eventually":{
        "throughput":{
            "mean":[326.15,630.53,1025.64,1395.35,1781.50,1939.69,2121.31,2115.00],
            "err":[6.917360,4.287661,18.282140,22.434754,24.278297,67.586712,48.091394,39.288704],
        },
        "e2eLatency":{
            "mean":[3.05,3.15,3.87,5.68,8.87,16.25,29.12,54.80],
            "err":[0.068488,0.021945,0.065998,0.094117,0.124015,0.627419,0.750622,1.336895],
        }
    }
}

def tuning_compare(t_c,legend,file):
    before,after,eventually=t_c["before"],t_c["after"],t_c["eventually"]

    plt.figure(figsize=(4.8,3.6))

    plt.errorbar(before["throughput"]["mean"],before["e2eLatency"]["mean"],
                 yerr=before["e2eLatency"]["err"],xerr=before["throughput"]["err"],
                 fmt='r.-',ecolor='black',label='Before')
    plt.errorbar(after["throughput"]["mean"],after["e2eLatency"]["mean"],
                 yerr=after["e2eLatency"]["err"],xerr=after["throughput"]["err"],
                 fmt='g.-',ecolor='black',label='After')
    plt.errorbar(eventually["throughput"]["mean"],eventually["e2eLatency"]["mean"],
                 yerr=eventually["e2eLatency"]["err"],xerr=eventually["throughput"]["err"],
                 fmt='b.-',ecolor='black',label='Eventually')

    plt.xlabel('Throughput(op/s)')
    # plt.xlim(xmin=0,xmax=2500)
    # plt.ylim(ymin=0,ymax=250)
    plt.ylabel('Latency(ms)')
    plt.subplots_adjust(left=0.12,right=0.96,top=0.98,bottom=0.13)

    if legend:
        plt.legend()

    if file != "":
        plt.savefig(file)
        # plt.show()
    else:
        plt.show()

# tuning_compare(t_c,True,"plots/tuning_compare.pdf")