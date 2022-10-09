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
				<view class="add">
					<image src="../../static/index/add.png"></image>
				</view>
			</view>
		</view>
		<view class="main">
			<view class="apply"></view>
			<view class="noone" v-if="noone">
				<image src="../../static/index/lonely.png" mode="aspectFill"></image>
				<view class="no-title">您还没有好友~</view>
				<view class="serch-bt" @tap="toSerch">搜索好友</view>
			</view>
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
		<view class="submit" @tap="exit">Test2</view>
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
			this.userInit()
			this.getFriends()
			this.getFriendsReq()
			this.getFriends2()
		},
		onShow() {
			this.userInit()
			this.getFriends()
			this.getFriendsReq()
			this.getFriends2()
			this.join()
		},
		onPullDownRefresh() {
			this.userInit()
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
			join: function() {
				console.log('hello')
				uni.connectSocket({
					url: config.socketurl + '?send_id=' + this.user.id + '&accept_id=app',
					success: data => {
						console.log(data.errMsg)
					}
				})
				uni.onSocketError(res => {
					console.log('WebSocket连接打开失败，请检查！' + res);
					this.join()
				});
				uni.onSocketMessage(res => {
					let data = JSON.parse(res.data);
					console.log('收到服务器内容：' + res.data);
					if (data.code == 1000) {
						for (let i = 0; i < this.friends2.length; i++) {
							if (this.friends2[i].id == data.send_id) {
								this.friends2[i].info = data.message
								this.friends2[i].tip += 1
							}
						}
					}
				});
			},
			exit: function() {
				console.log(666)
				uni.closeSocket({
					success: function(res) {
						console.log("WebSocket关闭成功！");
					},
					fail: function(res) {
						console.log("WebSocket关闭失败！");
					}
				})

			},
			goRequest: function() {
				uni.navigateTo({
					url: '../friendReq/friendReq'
				})
			},
			goChatRoom: function(item) {
				uni.navigateTo({
					url: '../chatRoom/chatRoom?friendInfo=' + encodeURIComponent(JSON.stringify(item)) +
						'&type=0'
				})
			},
			changeTime: function(date) {
				return calT.dateTime(date)
			},
			changeTime2: function(date) {
				return calT.dateTime2(date)
			},
			getFriends: function() {
				this.friends = datas.friends()
				for (let i = 0; i < this.friends.length; i++) {
					this.friends[i].imgurl = '../../static/index/' + this.friends[i].imgurl
				}
				//console.log(this.friends);
			},
			getFriends2: function() {
				uni.request({
					url: config.myurl + '/friend/myFriend',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': this.user.id,
						'state': 0
					},
					success: async data => {
						// console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getFriends2()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
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
								this.friends2 = calT.mySortByTime(this.friends2, 'time', 1)
								this.refershMsg()
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
			getFriendsReq: function() {
				console.log('send')
				uni.request({
					url: config.myurl + '/friend/myFriend',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
					data: {
						'user_id': this.user.id,
						'state': 1
					},
					success: async data => {
						// console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getFriendsReq()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							// console.log('Req')
							if (data.data.Data != null) {
								this.friendReq.tip = data.data.Data.length
								this.friendReq.time = data.data.Data[0].LastTime
								for (let i = 0; i < data.data.Data.length; i++) {
									//1第1个大 2第2个大 0 一样大
									if (calT.compareTime(data.data.Data[i].LastTime, this.friendReq
											.time) ==
										1) {
										this.friendReq.time = data.data.Data[i].LastTime
									}

								}
								// console.log(this.friendReq.tip)
							} else {
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
				console.log('get')
			},
			refershMsg: async function() {
				for (let i = 0; i < this.friends2.length; i++) {
					this.friends2[i].info = await this.getNewMsg(this.friends2[i].id)
					this.friends2[i].tip = await this.getUnReadMsg(this.friends2[i].id)
					// this.friends2[i].info=this.getNewMsg(this.friends2[i].id)
				}
			},
			getNewMsg: function(fid) {
				return new Promise((resolve, reject) => {
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
								switch (data.data.Data.Type) {
									case 0:
										resolve(data.data.Data.Msg)
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
								resolve('unknow')
							} else {
								resolve('unknow')
							}
						}
					})
				})

			},
			getUnReadMsg: function(fid) {
				return new Promise((resolve, reject) => {
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
			toSerch: function() {
				uni.navigateTo({
					url: '../serch/serch?user=' + encodeURIComponent(JSON.stringify(this.user)),
				})
			},
			userInit: function() {
				try {
					const value = uni.getStorageSync('user')
					if (value) {
						this.user.img = 'data:image/png;base64,' + value.img
						this.user.id = value.id
						this.user.email = value.email
						this.user.name = value.name
						this.user.nick = value.name
						this.atoken = uni.getStorageSync('atoken')
						this.rtoken = uni.getStorageSync('rtoken')
					} else {
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
