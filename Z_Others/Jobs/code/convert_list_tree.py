def convert_format(ori_list: list) -> dict:
    """
    设输入列表长度为n
    时间复杂度 O(n) 遍历两次list,
        1. 构建辅助字典, 实现通过子节点查找父节点时间复杂度为O(1); 初始化子节点
        2. 不断将子节点追加到父节点

    空间复杂度: 需要构建辅助节点dict, 和返回结果dict, 空间复杂度是 O(n)
    """
    tree = {"root": {}}
    revert_dict = dict()
    for n in ori_list:
        # 构造反向键值字典
        revert_dict[n["name"]] = n.get("parent")
        # 初始化子节点
        tree[n["name"]] = {}
    for ele in ori_list:
        name = ele["name"]
        parent = revert_dict.get(name) or "root"
        # 子节点追加到父节点
        tree[parent].update({name: tree.get(name)})
    return tree["root"]


def convert_format2(ori_list: list) -> dict:
    """
    设输入列表长度为n
    时间复杂度 每个元素都要遍历一边list寻找其子节点, 所以时间复杂度是 O(n^2)
        考虑到list节点多的情况下, 可以去掉已经排序的过节点, 逐步缩小待排序空间, 得最差时间复杂度是 O(n^2)
    空间复杂度: 需要构建辅助子节点列表, 索引列表 和返回结果dict, 空间复杂度是 O(n)
    """
    tree = dict()

    def build_tree(tree_node, parent_val):
        children = list()
        # index_pop = list()
        # find children in parent_val
        for i, n in enumerate(input_list):
            if n.get("parent") == parent_val:
                children.append(n["name"])
        #     if n.get("name") == parent_val:
        #         index_pop.append(i)
        # # remove converted nodes
        # for i in index_pop:
        #     input_list.pop(i)
        # create new node and build new child tree
        for child in children:
            tree_node[child] = {}
            build_tree(tree_node[child], child)

    build_tree(tree, None)
    return tree


if __name__ == "__main__":
    input_list = [{
        "parent": "交易所",
        "name": "中国外汇交易中⼼"
    }, {
        "name": "交易所"
    }, {
        "parent": "交易所",
        "name": "聚合交易所"
    }, {
        "parent": "交易所",
        "name": "森浦Quebee"
    }, {
        "name": "交易模式"
    }, {
        "parent": "交易模式",
        "name": "报价驱动模式"
    }, {
        "parent": "报价驱动模式",
        "name": "可执⾏持续报价(ESP)"
    }]
    # python -m cProfile convert_list_tree.py
    for i in range(10000):
        output = convert_format(input_list)
    expect = {
        "交易所": {
            "中国外汇交易中⼼": {},
            "聚合交易所": {},
            "森浦Quebee": {}
        },
        "交易模式": {
            "报价驱动模式": {
                "可执⾏持续报价(ESP)": {}
            }
        }
    }
    print("o", output, "e", expect)
    assert (output == expect)
