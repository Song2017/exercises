## 单机容器网络的实现原理和 docker0 网桥的作用
网卡（Network Interface）将数据转化为物理层01点信号、 
回环设备（Loopback Device）, 虚拟 的网络接口，你的计算机使用与自己沟通、 
路由表（Routing Table）网络IP映射 、 
iptables 规则 linux防火墙配置、 
这里的关键在于，容器要想跟外界进行通信，它发出的 IP 包就必须从它的 Network Namespace 里出来，来到宿主机上。
而解决这个问题的方法就是：为容器创建一个一端在容器里充当默认网卡、另一端在宿主机上的 Veth Pair 设备。
## 容器跨主机网络
### Flannel UDP 和 VXLAN 模式的工作原理
#### UDP
比于两台宿主机之间的直接通信，基于 Flannel UDP 模式的容器通信多了一个额外的步骤，即 flanneld 的处理过程。而这个过程，由于使用到了 flannel0 这个 TUN 设备，仅在发出 IP 包的过程中，就需要经过三次用户态与内核态之间的数据拷贝
第一次，用户态的容器进程发出的 IP 包经过 docker0 网桥进入内核态；
第二次，IP 包根据路由表进入 TUN（flannel0）设备，从而回到用户态的 flanneld 进程；
第三次，flanneld 进行 UDP 封包之后重新进入内核态，将 UDP 包通过宿主机的 eth0 发出去。
此外，我们还可以看到，Flannel 进行 UDP 封装（Encapsulation）和解封装（Decapsulation）的过程，也都是在用户态完成的。
在 Linux 操作系统中，上述这些上下文切换和用户态操作的代价其实是比较高的，这也正是造成 Flannel UDP 模式性能不好的主要原因。
所以说，我们在进行系统级编程的时候，有一个非常重要的优化原则，就是要减少用户态到内核态的切换次数，并且把核心的处理逻辑都放在内核态进行。
这也是为什么，Flannel 后来支持的VXLAN 模式，逐渐成为了主流的容器网络方案的原因。
#### VXLAN
VXLAN，即 Virtual Extensible LAN（虚拟可扩展局域网），是 Linux 内核本身就支持的一种网络虚似化技术
VXLAN 的覆盖网络的设计思想是：在现有的三层网络之上，“覆盖”一层虚拟的、由内核 VXLAN 模块负责维护的二层网络，使得连接在这个 VXLAN 二层网络上的“主机”（虚拟机或者容器都可以）之间，可以像在同一个局域网（LAN）里那样自由通信
为了能够在二层网络上打通“隧道”，VXLAN 会在宿主机上设置一个特殊的网络设备作为“隧道”的两端。这个设备就叫作 VTEP，即：VXLAN Tunnel End Point（虚拟隧道端点）。而 VTEP 设备的作用，其实跟前面的 flanneld 进程非常相似。只不过，它进行封装和解封装的对象，是二层数据帧（Ethernet frame）；而且这个工作的执行流程，全部是在内核里完成的（因为 VXLAN 本身就是 Linux 内核中的一个模块）
每台宿主机上名叫 flannel.1 的设备，就是 VXLAN 所需的 VTEP 设备，它既有 IP 地址，也有 MAC 地址
flannel.1 设备实际上要扮演一个“网桥”的角色，在二层网络进行 UDP 包的转发。而在 Linux 内核里面，“网桥”设备进行转发的依据，来自于一个叫作 FDB（Forwarding Database）的转发数据库
```
# 在Node 1上，使用“目的VTEP设备”的MAC地址进行查询
$ bridge fdb show flannel.1 | grep 5e:f8:4f:00:e3:37
5e:f8:4f:00:e3:37 dev flannel.1 dst 10.168.0.3 self permanent
```
### Kubernetes 网络模型与 CNI 网络插件
容器跨主机网络的两种实现方法：UDP 和 VXLAN, 有一个共性，那就是用户的容器都连接在 docker0 网桥上。
而网络插件则在宿主机上创建了一个特殊的设备（UDP 模式创建的是 TUN 设备，VXLAN 模式创建的则是 VTEP 设备），docker0 与这个设备之间，通过 IP 转发（路由表）进行协作。然后，网络插件真正要做的事情，则是通过某种方法，把不同宿主机上的特殊设备连通，从而达到容器跨主机通信的目的。

实际上，上面这个流程，也正是 Kubernetes 对容器网络的主要处理方法。只不过，Kubernetes 是通过一个叫作 CNI 的接口，维护了一个单独的网桥来代替 docker0。这个网桥的名字就叫作：CNI 网桥，它在宿主机上的设备名称默认是：cni0。

CNI 的设计思想，就是：Kubernetes 在启动 Infra 容器之后，就可以直接调用 CNI 网络插件，为这个 Infra 容器的 Network Namespace，配置符合预期的网络栈
CNI 的基础可执行文件，按照功能可以分为三类：
第一类，叫作 Main 插件，它是用来创建具体网络设备的二进制文件。
比如，bridge（网桥设备）、ipvlan、loopback（lo 设备）、macvlan、ptp（Veth Pair 设备），以及 vlan。
第二类，叫作 IPAM（IP Address Management）插件，它是负责分配 IP 地址的二进制文件。
比如，dhcp，这个文件会向 DHCP 服务器发起请求；host-local，则会使用预先配置的 IP 地址段来进行分配。
第三类，是由 CNI 社区维护的内置 CNI 插件。比如：flannel，就是专门为 Flannel 项目提供的 CNI 插件；
tuning，是一个通过 sysctl 调整网络设备参数的二进制文件；
portmap，是一个通过 iptables 配置端口映射的二进制文件；
bandwidth，是一个使用 Token Bucket Filter (TBF) 来进行限流的二进制文件。
Kubernetes 网络模型
所有容器都可以直接使用 IP 地址与其他容器通信，而无需使用 NAT。
所有宿主机都可以直接使用 IP 地址与所有容器通信，而无需使用 NAT。反之亦然。
容器自己“看到”的自己的 IP 地址，和别人（宿主机或者容器）看到的地址是完全一样的。
### Kubernetes三层网络方案
Host-gw模式通过在宿主机上添加一个路由规则：

    <目的容器IP地址段> via <网关的IP地址> dev eth0

IP包在封装成帧发出去的时候，会使用路由表里的“下一跳”来设置目的MAC地址。这样，它就会通过二层网络到达目的宿主机。
这个三层网络方案得以正常工作的核心，是为每个容器的IP地址，找到它所对应的，“下一跳”的网关。所以说，Flannel host-gw模式必须要求集群宿主机之间是二层连通的，如果宿主机分布在了不同的VLAN里（三层连通），由于需要经过的中间的路由器不一定有相关的路由配置（出于安全考虑，公有云环境下，宿主机之间的网关，肯定不会允许用户进行干预和设置），部分节点就无法找到容器IP的“下一条”网关了，host-gw就无法工作了。

Calico项目提供的网络解决方案，与Flannel的host-gw模式几乎一样，也会在宿主机上添加一个路由规则：

    <目的容器IP地址段> via <网关的IP地址> dev eth0

其中，网关的IP地址，正是目的容器所在宿主机的IP地址，而正如前面所述，这个三层网络方案得以正常工作的核心，是为每个容器的IP地址，找到它所对应的，“下一跳”的网关。区别是如何维护路由信息：
Host-gw : Flannel通过Etcd和宿主机上的flanneld来维护路由信息
Calico: 通过BGP（边界网关协议）来实现路由自治，所谓BGP，就是在大规模网络中实现节点路由信息共享的一种协议。

隧道技术（需要封装包和解包，因为需要伪装成宿主机的IP包，需要三层链通）：Flannel UDP / VXLAN / Calico IPIP
三层网络（不需要封包和解封包，需要二层链通）：Flannel host-gw / Calico 普通模式 host-gw 模式能够正常工作的核心，就在于 IP 包在封装成帧发送出去的时候，会使用路由表里的“下一跳”来设置目的 MAC 地址
### Kubernetes 里，网络隔离能力
NetworkPolicy 实际上只是宿主机上的一系列 iptables 规则。
这跟传统 IaaS 里面的安全组（Security Group）其实是非常类似的。
Kubernetes 的网络模型以及大多数容器网络实现，其实既不会保证容器之间二层网络的互通，也不会实现容器之间的二层网络隔离。
所以说，Kubernetes 从底层的设计和实现上，更倾向于假设你已经有了一套完整的物理基础设施。
然后，Kubernetes 负责在此基础上提供一种“弱多租户”（soft multi-tenancy）的能力。
并且，基于上述思路，Kubernetes 将来也不大可能把 Namespace 变成一个具有实质意义的隔离机制，或者把它映射成为“子网”或者“租户”。
毕竟你可以看到，NetworkPolicy 对象的描述能力，要比基于 Namespace 的划分丰富得多。
这也是为什么，到目前为止，Kubernetes 项目在云计算生态里的定位，其实是基础设施与 PaaS 之间的中间层。这是非常符合“容器”这个本质上就是进程的抽象粒度的。
当然，随着 Kubernetes 社区以及 CNCF 生态的不断发展，Kubernetes 项目也已经开始逐步下探，“吃”掉了基础设施领域的很多“蛋糕”。这也正是容器生态继续发展的一个必然方向。
### Service、DNS与服务发现
服务发现, Service 机制，以及 Kubernetes 里的 DNS 插件，都是在帮助你解决同样一个问题，即：如何找到我的某一个容器
Kubernetes 之所以需要 Service，一方面是因为 Pod 的 IP 不是固定的，另一方面则是因为一组 Pod 实例之间总会有负载均衡的需求
Service 提供的是 Round Robin 方式的负载均衡
1. ClusterIP 模式的 Service 
提供的，就是一个 Pod 的稳定的 IP 地址，即 VIP。并且，这里 Pod 和 Service 的关系是可以通过 Label 确定
iptables 规则 > DNAT 规则 (在 PREROUTING 检查点之前，也就是在路由之前，将流入 IP 包的目的地址和端口，改成–to-destination 所指定的新的目的地址和端口)
Service 是由 kube-proxy 组件，加上 iptables 来共同实现的
kube-proxy 设置–proxy-mode=ipvs
kube-proxy 通过 iptables 处理 Service 的过程，其实需要在宿主机上设置相当多的 iptables 规则。而且，kube-proxy 还需要在控制循环里不断地刷新这些规则来确保它们始终是正确的, 基于 iptables 的 Service 实现，都是制约 Kubernetes 项目承载更多量级的 Pod 的主要障碍
相比于 iptables，IPVS 在内核中的实现其实也是基于 Netfilter 的 NAT 模式，所以在转发这一层上，理论上 IPVS 并没有显著的性能提升。但是，IPVS 并不需要在宿主机上为每个 Pod 设置 iptables 规则，而是把对这些“规则”的处理放到了内核态，从而极大地降低了维护这些规则的代价
2. Headless Service 为你提供的，则是一个 Pod 的稳定的 DNS 名字，并且，这个名字是可以通过 Pod 名字和 Service 名字拼接出来的
kube-dns 生成的 DNS 记录
### 从外界连通 Service 与 Service 调试“三板斧”
外部访问 Service 的三种方式（NodePort、LoadBalancer 和 External Name）
公有云提供的 Kubernetes 服务里，都使用了一个叫作 CloudProvider 的转接层，来跟公有云本身的 API 进行对接。所以，在上述 LoadBalancer 类型的 Service 被提交后，Kubernetes 就会调用 CloudProvider 在公有云上为你创建一个负载均衡服务，并且把被代理的 Pod 的 IP 地址配置给负载均衡服务做后端
Service，其实就是 Kubernetes 为 Pod 分配的、固定的、基于 iptables（或者 IPVS）的访问入口。而这些访问入口代理的 Pod 信息，则来自于 Etcd，由 kube-proxy 通过控制循环来维护
当你的 Service 没办法通过 DNS 访问到的时候, 区分到底是 Service 本身的配置问题，还是集群的 DNS 出了问题
1. 检查 Kubernetes 自己的 Master 节点的 Service DNS 是否正常 nslookup kubernetes.default
2. Service 没办法通过 ClusterIP 访问, 检查的是这个 Service 是否有 Endpoints kubectl get endpoints hostnames
3. 确认 kube-proxy 是否在正确运行
4. 查看宿主机上的 iptables: KUBE-SERVICES 或者 KUBE-NODEPORTS 规则对应的 Service 的入口链，这个规则应该与 VIP 和 Service 端口一一对应；KUBE-SEP-(hash) 规则对应的 DNAT 链，这些规则应该与 Endpoints 一一对应；KUBE-SVC-(hash) 规则对应的负载均衡链，这些规则的数目应该与 Endpoints 数目一致；如果是 NodePort 模式的话，还有 POSTROUTING 处的 SNAT 链
5. Pod 没办法通过 Service 访问到自己: kubelet 的 hairpin-mode 没有被正确设置
for d in /sys/devices/virtual/net/cni0/brif/veth*/hairpin_mode; do echo "$d = $(cat $d)"; done
### Service 与 Ingress
这种全局的、为了代理不同后端 Service 而设置的负载均衡服务，就是 Kubernetes 里的 Ingress 服务
Ingress 对象，其实就是 Kubernetes 项目对“反向代理”的一种抽象
目前，Ingress 只能工作在七层，而 Service 只能工作在四层。
所以当你想要在 Kubernetes 里为应用进行 TLS 配置等 HTTP 相关的操作时，都必须通过 Ingress 来进行。
当然，正如同很多负载均衡项目可以同时提供七层和四层代理一样，将来 Ingress 的进化中，也会加入四层代理的能力。
这样，一个比较完善的“反向代理”机制就比较成熟了。
而 Kubernetes 提出 Ingress 概念的原因其实也非常容易理解，有了 Ingress 这个抽象，用户就可以根据自己的需求来自由选择 Ingress Controller