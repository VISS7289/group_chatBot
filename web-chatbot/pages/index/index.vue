<template>
	<view class="content">
		<view class="top-bar">
			<!-- <view class="top-bar-left"> -->
			<navigator :url="'../userhome/userhome?user='+encodeURIComponent(JSON.stringify(user))" hover-class="none"
				class="top-bar-left">
				<image :src="user.img" mode="aspectFill"></image>
			</navigator>
			<!-- </view> -->
			<view class="top-bar-center">
				<image src="../../static/general/logo.png" class="logo"></image>
			</view>
			<view class="top-bar-right">
				<view class="serch" @tap="toSerch">
					<image src="../../static/index/search.png"></image>
				</view>
				<view class="add" @tap="toBuildgroup">
					<image src="../../static/index/add.png"></image>
				</view>
			</view>
		</view>
		<view class="main">
			<view class="apply"></view>
			<!-- 			<view class="noone" v-if="noone">
				<image src="../../static/index/lonely.png" mode="aspectFill"></image>
				<view class="no-title">您还没有好友~</view>
				<view class="serch-bt" @tap="toSerch">搜索好友</view>
			</view> -->
			<view class="friends">
				<view class="friends-list" v-if="friendReq.tip>0" @tap="goRequest">
					<view class="friends-list-l">
						<text class="tip" v-if="friendReq.tip>0">{{friendReq.tip}}</text>
						<image :src="friendReq.img" mode="aspectFill"></image>
					</view>
					<view class="friends-list-r">
						<view class="top">
							<view class="name">{{friendReq.name}}</view>
							<view class="time">{{changeTime2(friendReq.time)}}</view>
						</view>
						<view class="info">{{friendReq.info}}</view>
					</view>
				</view>
			</view>
			<view class="friends">
				<view class="friends-list" @tap="goChatRoom(robotone)">
					<view class="friends-list-l">
						<text class="tip" v-if="robotone.tip>0">{{robotone.tip}}</text>
						<image :src="robotone.img" mode="aspectFill"></image>
					</view>
					<view class="friends-list-r">
						<view class="top">
							<view class="name">{{robotone.name}}</view>
							<view class="time">{{changeTime2(robotone.time)}}</view>
						</view>
						<view class="info">{{robotone.info}}</view>
					</view>
				</view>
			</view>
			<view class="friends" v-for="(item,index) in friends2" :key="index">
				<view class="friends-list" @tap="goChatRoom(item)">
					<view class="friends-list-l">
						<text class="tip" v-if="item.tip>0">{{item.tip}}</text>
						<image :src="item.img" mode="aspectFill"></image>
					</view>
					<view class="friends-list-r">
						<view class="top">
							<view class="name">{{item.name}}</view>
							<view class="time">{{changeTime2(item.time)}}</view>
						</view>
						<view class="info">{{item.info}}</view>
					</view>
				</view>
			</view>
		</view>
		<!-- <view class="submit" @tap="exit">Test2</view> -->
	</view>
</template>

<script>
	import datas from '../../commons/js/datas.js'
	import calT from '../../commons/js/calTime.js'
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	export default {
		data() {
			return {
				friends: [

				],
				friends2: [],
				user: {
					imgurl: '../../static/test/duan.png',
					id: 123,
					email: '123@12.com',
					img: '',
					name: 'duan',
					nick: 'duan',
					intr: '这个用户很懒，没有任何简介。这个用户很懒，没有任何简介。这个用户很懒，没有任何简介。'
				},
				robotone: {},
				friendReq: {
					img: '../../static/index/addUser.png',
					tip: 0,
					name: '好友申请',
					info: '茫茫人海，相见既是缘分。',
					time: '19:46:35'
				},
				atoken: '',
				rtoken: '',
				imgurl: '../../static/test/duan.png',
				noone: false,
			}
		},
		onLoad() {
			// 用户初始化
			this.userInit()
			// 获取聊天机器人
			this.getRobote()

			this.getFriends()
			// 获取好友请求
			this.getFriendsReq()
			// 获取好友列表
			this.getFriends2()
		},
		onShow() {
			this.userInit()
			this.getRobote()
			this.getFriends()
			this.getFriendsReq()
			this.getFriends2()
			this.join()
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.userInit()
			this.getRobote()
			this.getFriends()
			this.getFriendsReq()
			this.getFriends2()
			setTimeout(function() {
				uni.stopPullDownRefresh()
			}, 1000)
		},
		onHide() {
			this.exit();
		},
		onUnload() {
			this.exit();
		},
		methods: {
			// 获取聊天机器人
			getRobote: function() {
				// 发起网络请求
				uni.request({
					url: config.myurl + '/user/detial', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken // 设置请求头中的Authorization字段，包含访问令牌
					},
					data: {
						'id': '00000000000000000' // 提供的数据，包含一个id字段
					},
					success: async data => {
						// 请求成功的回调函数
						// console.log(data.data)
						if (data.data.Code == 1009) {
							// 判断返回的Code是否为1009
							// 通过调用其他函数重新获取访问令牌，并将新的令牌保存在atoken和rtoken变量中
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								// 如果刷新令牌成功（返回的Code为1000），更新atoken和rtoken的值
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								// 递归调用getRobote函数，重新获取聊天机器人信息
								this.getRobote()
							} else {
								// 刷新令牌失败，处理错误
								// err
							}
						} else if (data.data.Code == 1000) {
							// 返回的Code为1000，表示获取聊天机器人信息成功
							// console.log('friends')
							if (data.data.Data != null) {
								// 判断返回的数据Data是否为null
								this.noone = false
								console.log(data.data.Data)
								console.log(data.data.Data.CreateTime)
								// 将返回的数据解析并存储在robotone对象中
								this.robotone = {
									id: data.data.Data.UserID,
									name: data.data.Data.Username,
									img: 'data:image/png;base64,' + data.data.Data.Img,
									info: 'unknow',
									time: data.data.Data.CreateTime,
									tip: 0
								}
								console.log(this.robotone)
								// 调用getNewMsg函数获取与聊天机器人的最新消息
								let msg = await this.getNewMsg(this.robotone.id)
								console.log(msg)
								if (msg != null) {
									console.log(msg)
									// 如果有最新消息，更新robotone的info和time字段
									this.robotone.info = msg.Msg
									this.robotone.time = msg.Time
								}
								// 调用getUnReadMsg函数获取与聊天机器人的未读消息数量
								this.robotone.tip = await this.getUnReadMsg(this.robotone.id)
							}
						}
					}
				})
			},

			// 页面跳转
			toBuildgroup: function() {
				uni.navigateTo({
					url: '../buildgroup/buildgroup'
				})
			},
			// 与后端建立 WebSocket 连接
			join: function() {
				console.log('hello');
				uni.connectSocket({
					url: config.socketurl + '?send_id=' + this.user.id +
						'&accept_id=app', // 连接的 WebSocket 地址，包含发送者和接收者的ID
					success: data => {
						console.log(data.errMsg); // 输出连接成功的消息
					}
				});
				uni.onSocketError(res => {
					console.log('WebSocket连接打开失败，请检查！' + res); // WebSocket 连接失败的回调函数
					this.join(); // 重新尝试建立连接
				});
				uni.onSocketMessage(res => {
					let data = JSON.parse(res.data); // 解析接收到的消息数据为 JSON 对象
					console.log('收到服务器内容：' + res.data); // 输出接收到的服务器内容
					if (data.code == 1000) {
						// 如果接收到的消息 code 为 1000，表示有新的消息到达
						for (let i = 0; i < this.friends2.length; i++) {
							if (this.friends2[i].id == data.send_id) {
								// 找到与发送者ID匹配的朋友
								this.friends2[i].info = data.message; // 更新朋友的消息内容
								this.friends2[i].tip += 1; // 未读消息数量加一
							}
						}
					}
				});
			},

			// 关闭 WebSocket 连接
			exit: function() {
				console.log(666);
				uni.closeSocket({
					success: function(res) {
						console.log("WebSocket关闭成功！"); // WebSocket 关闭成功的回调函数
					},
					fail: function(res) {
						console.log("WebSocket关闭失败！"); // WebSocket 关闭失败的回调函数
					}
				});
			},

			// 跳转到好友请求页面
			goRequest: function() {
				uni.navigateTo({
					url: '../friendReq/friendReq' // 跳转到 friendReq 页面
				});
			},

			// 跳转到聊天室页面
			goChatRoom: function(item) {
				uni.navigateTo({
					url: '../chatRoom/chatRoom?friendInfo=' + encodeURIComponent(JSON.stringify(item)) +
						'&type=0'
					// 跳转到 chatRoom 页面，并传递参数 friendInfo 和 type
					// friendInfo 参数是 item 对象的 JSON 字符串形式，需要使用 encodeURIComponent 进行编码
					// type 参数为 0
				});
			},

			// 修改日期格式
			changeTime: function(date) {
				return calT.dateTime(date); // 调用 calT.dateTime() 函数，修改日期格式并返回结果
			},

			// 修改日期格式
			changeTime2: function(date) {
				if (date == undefined) {
					return; // 如果 date 为 undefined，直接返回
				}
				return calT.dateTime2(date); // 调用 calT.dateTime2() 函数，修改日期格式并返回结果
			},

			// 从缓存中获取好友列表
			getFriends: function() {
				this.friends = datas.friends(); // 调用 datas.friends() 函数，从缓存中获取好友列表，并赋值给 this.friends 变量
				for (let i = 0; i < this.friends.length; i++) {
					this.friends[i].imgurl = '../../static/index/' + this.friends[i]
						.imgurl; // 修改好友列表中每个好友对象的 imgurl 属性
				}
				//console.log(this.friends);
			},

			// 获取好友列表
			getFriends2: function() {
				uni.request({
					url: config.myurl + '/friend/myFriend', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken // 设置请求头中的Authorization字段，包含访问令牌
					},
					data: {
						'user_id': this.user.id, // 提供的数据，包含一个user_id字段和一个state字段
						'state': 0
					},
					success: async data => {
						// 请求成功的回调函数
						// console.log(data.data)
						if (data.data.Code == 1009) {
							// 判断返回的Code是否为1009
							// 通过调用其他函数重新获取访问令牌，并将新的令牌保存在atoken和rtoken变量中
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								// 如果刷新令牌成功（返回的Code为1000），更新atoken和rtoken的值
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								// 递归调用getFriends2函数，重新获取好友列表
								this.getFriends2()
							} else {
								// 刷新令牌失败，处理错误
								// err
							}
						} else if (data.data.Code == 1000) {
							// 返回的Code为1000，表示获取好友列表成功
							// console.log('friends')
							if (data.data.Data != null) {
								this.noone = false
								// console.log(data.data.Data)
								this.friends2 = []
								for (let i = 0; i < data.data.Data.length; i++) {
									let frinick = data.data.Data[i].FriendNick
									if (frinick == '') {
										frinick = data.data.Data[i].FriendName
									}
									// 将返回的数据解析并存储在friends2数组中
									this.friends2.push({
										id: data.data.Data[i].Friendid,
										name: frinick,
										truname: data.data.Data[i].FriendName,
										img: 'data:image/png;base64,' + data.data.Data[i]
											.FriendImg,
										info: 'unknow',
										time: data.data.Data[i].LastTime,
										tip: 0
									})
								}
								// 对friends2数组按照时间进行排序
								this.friends2 = calT.mySortByTime(this.friends2, 'time', 1)
								this.refershMsg() // 刷新消息
							} else {
								this.noone = true
								this.friends2 = []
							}
						} else {
							//err
						}
					}
				})
			},

			// 获取好友请求
			getFriendsReq: function() {
				console.log('send');
				uni.request({
					url: config.myurl + '/friend/myFriend', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken // 设置请求头中的Authorization字段，包含访问令牌
					},
					data: {
						'user_id': this.user.id, // 提供的数据，包含一个user_id字段和一个state字段
						'state': 1
					},
					success: async data => {
						// 请求成功的回调函数
						// console.log(data.data)
						if (data.data.Code == 1009) {
							// 判断返回的Code是否为1009
							// 通过调用其他函数重新获取访问令牌，并将新的令牌保存在atoken和rtoken变量中
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								// 如果刷新令牌成功（返回的Code为1000），更新atoken和rtoken的值
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								// 递归调用getFriendsReq函数，重新获取好友请求
								this.getFriendsReq()
							} else {
								// 刷新令牌失败，处理错误
								// err
							}
						} else if (data.data.Code == 1000) {
							// 返回的Code为1000，表示获取好友请求成功
							// console.log('Req')
							if (data.data.Data != null) {
								// 存在好友请求数据
								this.friendReq.tip = data.data.Data.length
								this.friendReq.time = data.data.Data[0].LastTime
								for (let i = 0; i < data.data.Data.length; i++) {
									// 比较时间，找到最新的请求时间
									//1第1个大 2第2个大 0 一样大
									if (calT.compareTime(data.data.Data[i].LastTime, this.friendReq
											.time) == 1) {
										this.friendReq.time = data.data.Data[i].LastTime
									}
								}
								// console.log(this.friendReq.tip)
							} else {
								// 不存在好友请求数据
								this.friendReq = {
									img: '../../static/index/addUser.png',
									tip: 0,
									name: '好友申请',
									info: '茫茫人海，相见既是缘分。',
									time: '19:46:35'
								}
							}
						} else {
							//err
						}
					}
				})
				console.log('get');
			},

			// 从缓存获取信息
			refershMsg: async function() {
				for (let i = 0; i < this.friends2.length; i++) {
					// 通过调用getNewMsg函数获取最新的消息
					let msg = await this.getNewMsg(this.friends2[i].id)
					if (msg != null) {
						// 如果消息不为空，则更新好友列表中的info字段
						this.friends2[i].info = msg.Msg
					}
					// 调用getUnReadMsg函数获取未读消息数量，并更新好友列表中的tip字段
					this.friends2[i].tip = await this.getUnReadMsg(this.friends2[i].id)
					// this.friends2[i].info=this.getNewMsg(this.friends2[i].id)
				}
			},

			// 获取最新消息
			getNewMsg: function(fid) {
				return new Promise((resolve, reject) => {
					// 发起请求获取最新消息
					uni.request({
						url: config.myurl + '/msg/newone',
						method: 'POST',
						header: {
							'Authorization': 'Bearer ' + this.atoken
						},
						data: {
							'send_id': this.user.id,
							'accept_id': fid
						},
						success: async data => {
							// console.log(data.data)
							if (data.data.Code == 1009) {
								let newCode = await refersh.refersh(config.myurl, this.atoken, this
									.rtoken)
								if (newCode == 1000) {
									this.atoken = uni.getStorageSync('atoken')
									this.rtoken = uni.getStorageSync('rtoken')
									this.getNewMsg()
								} else {
									// err
								}
							} else if (data.data.Code == 1000) {
								// 根据消息类型处理返回的数据
								switch (data.data.Data.Type) {
									case 0:
										resolve(data.data.Data)
										break
									case 1:
										resolve('[图片]')
										break
									case 2:
										resolve('[音频]')
										break
									case 3:
										resolve('[位置]')
										break
								}
								resolve(data.data.Data)
							} else {
								resolve(data.data.Data)
							}
						}
					})
				})
			},

			// 获取未读消息
			getUnReadMsg: function(fid) {
				return new Promise((resolve, reject) => {
					// 发起请求获取未读消息
					uni.request({
						url: config.myurl + '/msg/unread',
						method: 'POST',
						header: {
							'Authorization': 'Bearer ' + this.atoken
						},
						data: {
							'send_id': fid,
							'accept_id': this.user.id
						},
						success: async data => {
							// console.log(data.data)
							if (data.data.Code == 1009) {
								let newCode = await refersh.refersh(config.myurl, this.atoken, this
									.rtoken)
								if (newCode == 1000) {
									this.atoken = uni.getStorageSync('atoken')
									this.rtoken = uni.getStorageSync('rtoken')
									this.getUnReadMsg()
								} else {
									// err
								}
							} else if (data.data.Code == 1000) {
								resolve(data.data.Data)
							} else {
								resolve(0)
							}
						}
					})
				})
			},

			// 跳转到搜索页面
			toSerch: function() {
				uni.navigateTo({
					url: '../serch/serch?user=' + encodeURIComponent(JSON.stringify(this.user)),
				})
			},

			// 用户初始化
			userInit: function() {
				try {
					// 从缓存中获取用户信息
					const value = uni.getStorageSync('user')
					if (value) {
						// 如果存在用户信息，则将用户信息赋值给对应的属性
						this.user.img = 'data:image/png;base64,' + value.img
						this.user.id = value.id
						this.user.email = value.email
						this.user.name = value.name
						this.user.nick = value.name
						this.atoken = uni.getStorageSync('atoken')
						this.rtoken = uni.getStorageSync('rtoken')
					} else {
						// 如果不存在用户信息，则跳转到登录页面
						uni.navigateTo({
							url: '../signin/signin',
						})
					}
					console.log(value)
				} catch (e) {
					// error
				}
			},

			Test: function() {
				console.log("{\"type\": 1,\"message\": \"啦啦啦德玛西亚\"}")
				uni.sendSocketMessage({
					data: "{\"type\": 1,\"message\": \"啦啦啦德玛西亚\"}",
					// data: "{\"type\": 1,\"message\": \"啦啦啦德玛西亚}",
					success: data => {
						if (data.errMsg == 'sendSocketMessage:fail WebSocket is not connected') {
							this.join()
						}
						console.log(data)
					},
					fail: data => {
						console.log(data)
					}
				})
			}

		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	.top-bar {
		background-color: rgba(255, 255, 255, 0.98);
		border-bottom: 1px solid $uni-border-color;
	}

	.main {
		padding: 108rpx 0rpx $uni-spacing-col-base;
		//margin-top: var(--status-bar-height);
		width: 100%;
		//border: 1px solid red;
		box-sizing: border-box;
	}

	.noone {
		margin: 0 auto;
		text-align: center;
		padding-top: 200rpx;

		image {
			height: 500rpx;
			width: 316rpx;
		}

		.no-title {
			font-size: $uni-font-size-base;
			color: rgba(40, 30, 50, 0.4);
			line-height: 40rpx;
			padding-bottom: 32rpx;
			padding-top: 32rpx;
		}

		.serch-bt {
			margin: 0 auto;
			width: 520rpx;
			height: 96rpx;
			background: $uni-color-primary;
			box-shadow: 0rpx 50rpx 32rpx -36rpx rgba(255, 228, 49, 0.4);
			border-radius: 48rpx;
			font-size: $uni-font-size-lg;
			font-weight: 500;
			color: $uni-text-color;
			line-height: 96rpx;
			text-align: center;
		}
	}

	.friends-list {
		padding: 10rpx 30rpx 10rpx 30rpx;

		&:active {
			background-color: $uni-bg-color-grey;
		}

		.friends-list-l {
			position: relative;
			float: left;

			image {
				width: 96rpx;
				height: 96rpx;
				border-radius: $uni-border-radius-base;
				background-color: $uni-color-primary;
			}

			.tip {
				position: absolute;
				z-index: 100;
				top: -8rpx;
				left: 68rpx;
				min-width: 20rpx;
				height: 36rpx;
				padding: 0 8rpx;
				background-color: $uni-color-warning;
				border-radius: 18rpx;
				font-size: $uni-font-size-sm;
				color: $uni-text-color-inverse;
				line-height: 36rpx;
				text-align: center;
			}
		}

		.friends-list-r {
			padding-left: 128rpx;

			.top {
				height: 50rpx;

				.name {
					float: left;
					font-size: 36rpx;
					font-weight: 400;
					color: $uni-text-color;
					line-height: 50rpx;
				}

				.time {
					//display: none;
					float: right;
					font-size: $uni-font-size-sm;
					color: $uni-text-color-disable;
					line-height: 50rpx;
				}
			}

			.info {
				font-size: $uni-font-size-base;
				color: $uni-text-color-grey;
				line-height: 55rpx;
				overflow: hidden;
				text-overflow: ellipsis;
				white-space: nowrap;
			}
		}

	}
</style>
