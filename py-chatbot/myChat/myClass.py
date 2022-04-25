class Voc:
    def __init__(self, name):
        self. PAD_token = 0  # 表示padding
        self.SOS_token = 1  # 句子的开始
        self.EOS_token = 2  # 句子的结束
        self.name = name
        self.trimmed = False
        self.word2index = {} # 词映射到位置
        self.word2count = {} # 词频
        self.index2word = {self.PAD_token: "PAD", self.SOS_token: "SOS", self.EOS_token: "EOS"} # 位置映射到词
        self.num_words = 3  # 初始化时便有SOS, EOS, PAD这3个token。

    def addSentence(self, sentence):
        for word in sentence.split(' '):
            self.addWord(word)

    def addWord(self, word):
        if word not in self.word2index:
            self.word2index[word] = self.num_words
            self.word2count[word] = 1
            self.index2word[self.num_words] = word
            self.num_words += 1
        else:
            self.word2count[word] += 1

    # 删除频次小于min_count的token
    def trim(self, min_count):
        if self.trimmed:
            return
        self.trimmed = True

        keep_words = []

        for k, v in self.word2count.items():
            if v >= min_count:
                keep_words.append(k)

        print('keep_words {} / {} = {:.4f}'.format(
            len(keep_words), len(self.word2index), len(keep_words) / len(self.word2index)
        ))

        # 重新构造词典
        self.word2index = {}
        self.word2count = {}
        self.index2word = {self.PAD_token: "PAD", self.SOS_token: "SOS", self.EOS_token: "EOS"}
        self.num_words = 3  # 初始化时便有SOS, EOS, PAD这3个token。

        # 重新构造后词频就没有意义了(都是1)
        for word in keep_words:
            self.addWord(word)