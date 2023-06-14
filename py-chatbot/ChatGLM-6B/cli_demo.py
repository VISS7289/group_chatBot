import os
import platform
import signal

from transformers import AutoTokenizer, AutoModel

# 创建一个自动标记器（tokenizer）实例，使用预训练模型“..\\chatglm-6b-int4”，并启用远程代码的信任
tokenizer = AutoTokenizer.from_pretrained("..\\chatglm-6b-int4", trust_remote_code=True)

# 创建一个自动模型（model）实例，使用预训练模型“..\\chatglm-6b-int4”，并启用远程代码的信任
# 将模型转换为半精度浮点数格式（half）并在CUDA设备上运行（如果可用）
model = AutoModel.from_pretrained("..\\chatglm-6b-int4", trust_remote_code=True).half().cuda()

# 将模型设置为评估（inference）模式
model = model.eval()

# 获取操作系统的名称（Windows、Linux等）
os_name = platform.system()

# 根据操作系统的名称选择清屏命令
clear_command = 'cls' if os_name == 'Windows' else 'clear'

# 用于指示是否停止对话的变量
stop_stream = False


# 用于构建用户提示的函数，包括历史对话记录
def build_prompt(history):
    prompt = "欢迎使用 ChatGLM-6B 模型，输入内容即可进行对话，clear 清空对话历史，stop 终止程序"
    for query, response in history:
        prompt += f"\n\n用户：{query}"
        prompt += f"\n\nChatGLM-6B：{response}"
    return prompt


# 信号处理函数，用于捕获终止程序的信号
def signal_handler(signal, frame):
    global stop_stream
    stop_stream = True


# 主函数
def main():
    history = []  # 存储对话历史记录的列表
    global stop_stream
    print("欢迎使用 ChatGLM-6B 模型，输入内容即可进行对话，clear 清空对话历史，stop 终止程序")
    while True:
        query = input("\n用户：")  # 接收用户输入的查询内容
        if query.strip() == "stop":  # 如果用户输入"stop"，则退出循环，终止程序
            break
        if query.strip() == "clear":  # 如果用户输入"clear"，则清空对话历史记录，并清屏
            history = []
            os.system(clear_command)
            print("欢迎使用 ChatGLM-6B 模型，输入内容即可进行对话，clear 清空对话历史，stop 终止程序")
            continue
        count = 0
        # 使用模型的stream_chat方法进行对话生成，返回一个生成器
        # 在每次迭代中生成模型的响应，并将响应与查询添加到对话历史中
        for response, history in model.stream_chat(tokenizer, query, history=history):
            if stop_stream:  # 如果stop_stream标志为True，则停止对话生成
                stop_stream = False
                break
            else:
                count += 1
                if count % 8 == 0:  # 每8次迭代清屏并打印当前对话历史
                    os.system(clear_command)
                    print(build_prompt(history), flush=True)
                    signal.signal(signal.SIGINT, signal_handler)  # 捕获终止程序的信号（Ctrl+C）
        os.system(clear_command)
        print(build_prompt(history), flush=True)  # 打印最终的对话历史


# 程序入口
if __name__ == "__main__":
    main()
