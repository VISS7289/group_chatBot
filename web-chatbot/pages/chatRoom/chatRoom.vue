<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png" class="back-img"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">{{friend.name}}</view>
			</view>
			<view class="top-bar-right">
				<view class="group-img">
					<image :src="friend.img" mode="aspectFill"></image>
				</view>
			</view>
		</view>
		<scroll-view class="chat" scroll-y="true" scroll-with-animation="true" :scroll-into-view="scrollToView">
			<view class="chat-main" :style="{paddingBottom : h + 'px'}">
				<view class="chat-ls" v-for="(item,index) in msgs" :key="index" :id="'msg'+item.tip">
					<view class="chat-time" v-if="changeTime3(index)!=''">{{changeTime3(index)}}</view>
					<view class="msg msg-left" v-if="item.id != user.id">
						<image :src="item.img" class="user-img" mode="aspectFill"></image>
						<view class="msg-msg" v-if="item.type==0">
							<view class="msg-text">{{item.msg}}</view>
						</view>
						<view class="msg-msg" v-if="item.type==1">
							<image :src="item.msg" class="msg-img" mode="widthFix" @tap="previewImg(item.msg)">
							</image>
						</view>
					</view>
					<view class="msg msg-right" v-if="item.id == user.id">
						<image :src="user.img" class="user-img" mode="aspectFill"></image>
						<view class="msg-msg" v-if="item.type==0">
							<view class="msg-text">{{item.msg}}</view>
						</view>
						<view class="msg-msg" v-if="item.type==1">
							<image :src="item.msg" class="msg-img" mode="widthFix" @tap="previewImg(item.msg)">
							</image>
						</view>
					</view>
				</view>
			</view>
			<view class="padbt"></view>
		</scroll-view>
		<submit @inputChatMsg="inputChat" @heights="heights"></submit>
	</view>
</template>

<script>
	// 导入所需模块
	import datas from '../../commons/js/datas.js'
	import calT from '../../commons/js/calTime.js'
	import submit from '../../components/submit/submit.vue'
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	// 导出模块
	export default {
		data() {
			return {
				// 数据变量
				msgs: [], // 消息列表
				imgMsg: [], // 图片消息列表
				nowTime: calT.getNewTime(), // 当前时间
				scrollToView: '', // 滚动到视图
				h: '120', // 高度
				friend: {}, // 好友信息
				atoken: '', // 访问令牌
				rtoken: '', // 刷新令牌
				user: {}, // 用户信息
				nowpage: 0, // 当前页数
				maxpage: 10, // 最大页数
				loadOver: false, // 加载完毕标志
				botnum: 2, // 机器人消息数
				timer: null, // 定时器
			}
		},
		onLoad(option) {
			// 页面加载时触发的函数
			this.userInit() // 初始化用户信息
			this.friend = JSON.parse(decodeURIComponent(option.friendInfo)) // 解码好友信息
			console.log(this.friend)
			console.log(option.type)
			this.getMsg2() // 获取聊天消息
		},
		onShow() {
			// 页面显示时触发的函数
			this.join() // 加入聊天
			if (this.friend.id == '00000000000000000') {
				// 如果好友id为特定值
				this.timer = setInterval(() => {
					this.botnum -= 1
					console.log(this.botnum)
					if (this.botnum < 0) {
						this.botnum = 0
					} else {
						console.log(this.msgs[this.msgs.length - 1])
						// uni.sendSocketMessage({
						// 	data: "{\"type\": 0,\"message\": \"" + this.msgs[this.msgs.length - 1].msg +
						// 		"\"}",
						// 	success: data => {
						// 		if (data.errMsg ==
						// 			'sendSocketMessage:fail WebSocket is not connected') {
						// 			this.join()
						// 		}
						// 	},
						// 	fail: data => {
						// 		console.log(data.errMsg)
						// 	}
						// })
					}
				}, 4000 + Math.round(10000 * Math.random()));
			}
		},
		onHide() {
			// 页面隐藏时触发的函数
			this.exit(); // 退出聊天
		},
		onUnload() {
			// 页面卸载时触发的函数
			this.exit(); // 退出聊天
		},
		components: {
			submit,
		},
		methods: {
			join: function() {
				// 加入聊天
				console.log('hello'); // 输出"hello"到控制台

				// 连接WebSocket
				uni.connectSocket({
					url: config.socketurl + '?send_id=' + this.user.id + '&accept_id=' + this.friend.id,
					success: data => {
						console.log(data.errMsg); // 输出连接成功的消息到控制台
					}
				});

				// 监听WebSocket连接错误
				uni.onSocketError(res => {
					console.log('WebSocket连接打开失败，请检查！' + res); // 输出连接错误信息到控制台
					this.join(); // 重新调用join()函数，尝试重新连接
				});

				// 监听WebSocket消息
				uni.onSocketMessage(res => {
					let data = JSON.parse(res.data); // 解析收到的消息数据为JSON对象
					console.log('收到服务器内容：' + res.data); // 输出收到的服务器消息内容到控制台
					if (data.code == 1000) {
						// 如果收到的消息code为1000
						this.reciveMsg(this.friend.id, this.friend.img, data.message, 0);
						// 调用reciveMsg()函数，并传入好友ID、好友头像、消息内容和0作为参数
					}
				});
			},

			exit: function() {
				console.log(666); // 输出666到控制台

				// 关闭WebSocket连接
				uni.closeSocket({
					success: function(res) {
						console.log("WebSocket关闭成功！"); // 输出关闭成功的消息到控制台
					},
					fail: function(res) {
						console.log("WebSocket关闭失败！"); // 输出关闭失败的消息到控制台
					}
				});
			},

			socketSent: function(e) {
				return new Promise((resolve, reject) => {
					const items = ['0', '0', '1', '1', '2', '2', '3', '4'];
					this.botnum = items[Math.floor(Math.random() * items
						.length)]; // 从items数组中随机选择一个元素，并将结果赋值给this.botnum
					console.log(this.botnum); // 输出this.botnum到控制台

					console.log("{\"type\": 1,\"message\": \"" + e + "\"}"); // 输出发送的消息内容到控制台

					uni.sendSocketMessage({
						data: "{\"type\": 1,\"message\": \"" + e + "\"}", // 发送消息的数据格式
						success: data => {
							if (data.errMsg ==
								'sendSocketMessage:fail WebSocket is not connected') {
								this.join(); // WebSocket连接未打开，重新连接
								resolve(false); // 返回Promise并解析为false
							} else {
								resolve(true); // 返回Promise并解析为true
							}
							console.log(data.errMsg); // 输出发送消息的结果到控制台
						},
						fail: data => {
							resolve(false); // 返回Promise并解析为false
							console.log(data.errMsg); // 输出发送消息的失败原因到控制台
						}
					});
				});
			},

			userInit: function() {
				try {
					const value = uni.getStorageSync('user'); // 从本地缓存中获取'user'的值并赋给value变量

					if (value) {
						this.user.img = 'data:image/png;base64,' + value.img; // 从value中获取用户头像信息并赋值给this.user.img
						this.user.id = value.id; // 从value中获取用户ID并赋值给this.user.id
						this.user.email = value.email; // 从value中获取用户电子邮件并赋值给this.user.email
						this.user.name = value.name; // 从value中获取用户名并赋值给this.user.name
						this.user.nick = value.name; // 从value中获取用户名并赋值给this.user.nick
						this.atoken = uni.getStorageSync('atoken'); // 从本地缓存中获取'atoken'的值并赋给this.atoken
						this.rtoken = uni.getStorageSync('rtoken'); // 从本地缓存中获取'rtoken'的值并赋给this.rtoken
					} else {
						uni.navigateTo({
							url: '../signin/signin', // 跳转到signin页面
						});
					}

					console.log(value); // 输出value到控制台
				} catch (e) {
					// error
				}
			},

			backOne: function() {
				uni.navigateBack({
					delta: 1 // 返回上一个页面
				});
			},

			changeTime: function(date) {
				return calT.chatTime(date); // 使用calT对象的chatTime方法处理时间并返回结果
			},

			changeTime2: function(date) {
				return calT.dateTime2(date); // 使用calT对象的dateTime2方法处理时间并返回结果
			},

			changeTime3: function(index) {
				if (index == 0) {
					return calT.dateTime2(this.msgs[index].time); // 如果索引为0，使用calT对象的dateTime2方法处理时间并返回结果
				}

				let delta = calT.spaceTime2(this.msgs[index - 1].time, this.msgs[index]
					.time); // 使用calT对象的spaceTime2方法计算时间差
				let equal = (calT.dateTime2(this.msgs[index - 1].time) == calT.dateTime2(this.msgs[index]
					.time)); // 判断前一条消息和当前消息的日期是否相等

				if (delta > 5 && equal == false) {
					return calT.dateTime2(this.msgs[index].time); // 如果时间差大于5并且日期不相等，使用calT对象的dateTime2方法处理时间并返回结果
				} else {
					return ''; // 否则返回空字符串
				}
			},

			//获取聊天数据
			getMsg: function() {
				let msg = datas.message()
				for (var i = 0; i < msg.length; i++) {
					msg[i].imgurl = '../../static/index/' + msg[i].imgurl
					//时间间隔
					let t = calT.spaceTime(msg[i].time, this.nowTime)
					if (t) {
						this.nowTime = t
					}
					msg[i].time = t
					//信息为图片
					if (msg[i].types == 1) {
						msg[i].message = '../../static/index/' + msg[i].message
						this.imgMsg.unshift(msg[i].message)
					}
					this.msgs.unshift(msg[i])
				}
				//滑到当前聊天
				this.$nextTick(function() {
					this.scrollToView = 'msg' + this.msgs[i - 1].tip
				})
				//console.log(msg);
			},
			getMsg2: function() {
				if (this.nowpage == -1) {
					return;
				}

				uni.request({
					url: config.myurl + '/msg/getOldChatF', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken // 设置请求头的Authorization字段为Bearer加上atoken的值
					},
					data: {
						'user_id': this.user.id, // 请求的用户ID
						'friend_id': this.friend.id, // 请求的好友ID
						'now_page': this.nowpage, // 当前页码
						'max_page': this.maxpage, // 最大页码
					},
					success: async data => {
						console.log(data.data); // 输出返回的数据到控制台

						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken);
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken'); // 更新atoken
								this.rtoken = uni.getStorageSync('rtoken'); // 更新rtoken
								this.getMsg2(); // 重新调用getMsg2函数
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(data.data.Data);
							let l = this.msgs.length;
							if (data.data.Data == null) {
								return;
							}
							for (var i = data.data.Data.length - 1; i >= 0; i--) {
								this.msgs.push({
									id: data.data.Data[i].SendId,
									img: 'data:image/png;base64,' + data.data.Data[i].Img,
									msg: data.data.Data[i].Msg,
									time: data.data.Data[i].Time,
									type: data.data.Data[i].Type,
									tip: l + i
								});
							}
							// 滑到当前聊天
							this.$nextTick(function() {
								this.scrollToView = 'msg' + this.msgs[data.data.Data.length - 1]
									.tip;
							});
							if (data.data.Data.length == 10) {
								this.nowpage++;
							} else {
								// 数据提取完毕
								this.nowpage = -1;
							}
						} else {
							console.log('err');
						}
					}
				});
			},

			previewImg: function(urlImg) {
				uni.previewImage({
					current: urlImg, // 当前显示图片的链接
					urls: this.imgMsg, // 需要预览的图片链接列表
					longPressActions: {
						itemList: ['发送给朋友', '保存图片', '收藏'], // 长按图片时显示的操作列表
						success: function(data) {
							console.log('选中了第' + (data.tapIndex + 1) + '个按钮,第' + (data.index + 1) + '张图片');
						},
						fail: function(err) {
							console.log(err.errMsg);
						}
					}
				});
			},

			inputChat: async function(inf) {
				console.log(calT.getNewTime()); // 输出当前时间到控制台
				let res = await this.socketSent(inf); // 调用socketSent函数发送消息并等待返回结果

				if (res) {
					this.reciveMsg(this.user.id, this.user.img, inf, 0); // 调用reciveMsg函数，模拟接收到消息
				}
			},

			reciveMsg: function(id, img, msg, type) {
				let len = this.msgs.length;
				let newMsg = {
					id: id,
					img: img,
					msg: msg,
					type: type,
					time: calT.getNewTime(),
					tip: len,
				};
				console.log(this.changeTime2(newMsg.time)); // 输出处理后的时间到控制台
				this.msgs.push(newMsg); // 将新消息添加到消息列表
				this.goBottom(); // 滚动到底部
			},

			answerRandomUser: async function(msg) {
				let retMsg = await this.chat(msg); // 调用chat函数进行聊天，并等待返回结果
				let len = this.msgs.length;
				let theNum = Math.round(145 * Math.random()).toString(); // 生成随机数

				console.log(retMsg);
				let newMsg = {
					id: '1231',
					imgurl: '1 (' + theNum + ').jpg',
					message: retMsg,
					types: 0,
					time: new Date(),
					tip: len,
				};
				newMsg.imgurl = '../../static/index/' + newMsg.imgurl; // 设置图片路径
				let t = calT.spaceTime(newMsg.time, this.nowTime); // 计算时间间隔
				if (t) {
					this.nowTime = t; // 更新当前时间
				}
				newMsg.time = t;
				this.msgs.push(newMsg); // 将新消息添加到消息列表
				this.goBottom(); // 滚动到底部
			},

			heights: function(inf) {
				this.h = inf; // 更新高度信息
				this.goBottom(); // 滚动到底部
				console.log(inf); // 输出高度信息到控制台
			},

			//滚动到底部
			goBottom: function() {
				this.scrollToView = ''
				this.$nextTick(function() {
					let len = this.msgs.length - 1
					this.scrollToView = 'msg' + this.msgs[len].tip
				})
			},
			//发送聊天内容
			chat: function(message) {
				return new Promise((resolve, reject) => {
					let retMsg = '1'

					const xhr = new XMLHttpRequest()
					const url = 'http://localhost:8086/ping'

					xhr.open('GET', url)
					xhr.onreadystatechange = function() {
						console.log(xhr.responseText)
					}
					xhr.send()


					uni.request({
						url: 'http://localhost:8086/chat', //仅为示例，并非真实接口地址。
						method: 'POST',
						data: JSON.stringify({
							chat_content: message,
						}),
						header: {
							'custom-header': 'hello' //自定义请求头信息
						},
						success: res => {
							retMsg = res.data.Data
							resolve(retMsg)
						},
					})
					// uni.request({
					// 	//url: '/chat/chat',
					// 	url: '/chat/chat',
					// 	method: "POST",
					// 	data: JSON.stringify({
					// 		chat_content: message,
					// 	}),
					// 	sslVerify:  false,
					// 	success: res => {
					// 		retMsg = res.data.Data;
					// 		resolve(retMsg)
					// 	},
					// 	fail: res => {
					// 		console.log(0);
					// 		console.log(res);
					// 		//reject
					// 	},
					// 	complete: res => {
					// 		console.log(2);
					// 	},
					// })

				})

			},
			test: function() {
				//console.log(2);
			}

		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	page {
		height: 100%;
	}

	.content {
		height: 100%;
		background: rgba(244, 244, 244, 0.98);
	}

	.top-bar {
		background: rgba(244, 244, 244, 0.98);
		border-bottom: 1px solid $uni-border-color;
		position: fixed;
		top: 0;

		.top-bar-right {
			margin-right: 9rpx;

			.group-img {
				//float: right;
				// bottom: 10rpx;
				width: 68rpx;
				height: 68rpx;
				// margin-top: 10rpx;
				// margin-right: 10rpx;
				margin: 0;
				padding: 0rpx;

				image {
					margin: 9rpx 0;
					width: 68rpx;
					height: 68rpx;
					//display: inline-block;
					border-radius: $uni-border-radius-base;
				}
			}
		}

	}

	.chat {
		height: 100%;

		.padbt {
			padding-top: var(--status-bar-height);
		}

		.chat-main {
			padding-left: $uni-spacing-col-base;
			padding-right: $uni-spacing-col-base;
			padding-top: 100rpx;
			//padding-bottom: 120rpx;
			display: flex;
			flex-direction: column;
		}

		.chat-ls {
			.chat-time {
				font-size: $uni-font-size-sm;
				color: rgba(39, 40, 50, 0.3);
				line-height: 34rpx;
				padding: 20rpx 0;
				text-align: center;
			}

			.msg {
				display: flex;
				padding: 20rpx 0;

				.user-img {
					flex: none;
					width: $uni-img-size-sm;
					height: $uni-img-size-sm;
					border-radius: $uni-border-radius-base;
				}

				.msg-msg {
					flex: none;
					max-width: 480rpx;
				}

				.msg-text {
					//width: fit-content;
					font-size: $uni-font-size-lg;
					color: $uni-text-color;
					line-height: 44rpx;
					padding: 18rpx 24rpx;
				}

				.msg-img {
					max-width: 464rpx;
					border-radius: 20rpx;
				}
			}

			.msg-left {
				flex-direction: row;

				.msg-text {
					margin-left: 16rpx;
					background-color: #fff;
					border-radius: 0rpx 20rpx 20rpx 20rpx;
				}

				.msg-img {
					margin-left: 16rpx;
				}
			}

			.msg-right {
				flex-direction: row-reverse;

				.msg-text {
					margin-right: 16rpx;
					background-color: rgba(255, 228, 49, 0.8);
					border-radius: 20rpx 0rpx 20rpx 20rpx;
				}

				.msg-img {
					margin-right: 16rpx;
				}
			}
		}
	}
</style>
