fix length:
    发送方与接收方都使用固定大小的缓冲区，长度不够用空字符弥补，简单但不够灵活
    如：
        发送：AAAAAAAAAABBBBBBBBCCCCCCCC
        接收：AAAAAAAAAA
             BBBBBBBBCC
             CCCCCC
delimiter based:
    以特殊字符作为每条消息的结尾来区分，如"\n",碰到该字符，代表一条完整数据 数据长度可自定义，但容易被人破解
    如：
        发送：AAAAAAAAAAAA\nbbbbbbbbbbbbbbb\nccccccc\n
        接收：AAAAAAAAAAAA
             bbbbbbbbbbbbbbb
             ccccccc
length field based frame decoder:
    数据包装成数据头+数据正文，数据头可以分多帧，一般用固定长度形式存储元数据，可以根据数据头的元数据获取具体的数据信息，如goim的编码格式 
    自定义程度高，复杂度高
    数据头：packageLength 4bytes
           HeaderLength  2bytes
           ProtocolVersion 2bytes
           Operation 4bytes
           Sequence 4bytes
    数据正文：Body packageLength-HeaderLength