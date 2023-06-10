<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">好友请求</view>
			</view>
		</view>
		<view class="main">
			<!-- <view class="requestLine"  v-for="(item,index) in requester" :key="index">
				<view class="requestTop">
					<view class="btn reject">拒绝</view>
					<view class="requestImg">
						<image :src="item.imgurl" mode="aspectFill"></image>
					</view>
					<view class="btn accept">接受</view>
				</view>
				<view class="requestCenter">
					<view class="name">{{item.name}}</view>
					<view class="time">{{changeTime(item.time)}}</view>
				</view>
				<view class="requestBottom">
					<view class="message">{{item.info}}</view>
				</view>
			</view> -->
			<view class="requestLine" v-for="(item,index) in requester2" :key="index">
				<view class="requestTop">
					<view class="btn reject" @tap="rejectFriend(item.id,index)">拒绝</view>
					<view class="requestImg">
						<image :src="item.img" mode="aspectFill"></image>
					</view>
					<view class="btn accept" @tap="acceptFriend(item.id,index)">接受</view>
				</view>
				<view class="requestCenter">
					<view class="name">{{item.name}}</view>
					<view class="time">{{changeTime2(item.time)}}</view>
				</view>
				<view class="requestBottom">
					<view class="message">{{item.info}}</view>
				</view>
			</view>
		</view>
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
				requester: [

				],
				requester2: [],
				user: {},
				atoken: '',
				rtoken: '',
			}
		},
		onLoad() {
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
			this.getRequester()
			this.getFriendsReq()
		},
		methods: {
			backOne: function() {
				let pages = getCurrentPages(); // 获取当前页面栈
				let beforePage = pages[pages.length - 2]; // 获取当前页面栈中的上一个页面
				uni.navigateBack({
					delta: 1, // 返回的页面数，这里是返回上一个页面
					success: function() {
						beforePage.getFriendsReq(); // 调用上一个页面的getFriendsReq函数
						beforePage.getFriends2(); // 调用上一个页面的getFriends2函数
					}
				});
			},

			changeTime: function(date) {
				return calT.dateTime(date); // 调用calT.dateTime函数并返回结果
			},

			changeTime2: function(date) {
				return calT.dateTime2(date); // 调用calT.dateTime2函数并返回结果
			},

			getRequester: function() {
				this.requester = datas.friends(); // 将datas.friends()的结果赋值给this.requester
				for (let i = 0; i < this.requester.length; i++) {
					this.requester[i].imgurl = '../../static/index/' + this.requester[i]
						.imgurl; // 修改this.requester中每个对象的imgurl属性
				}
				//console.log(this.friends);
			},

			getFriendsReq: function() {
				console.log('send'); // 打印日志：send
				uni.request({
					url: config.myurl + '/friend/myFriend', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken
					}, // 请求头中添加Authorization字段
					data: {
						'user_id': this.user.id, // 请求的用户ID
						'state': 1 // 请求的状态值为1
					},
					success: async data => { // 请求成功的回调函数
						console.log(typeof(data.data)); // 打印data.data的类型
						console.log(data.data.Code); // 打印data.data.Code的值
						if (data.data.Code == 1009) { // 如果返回的Code为1009
							let newCode = await refersh.refersh(config.myurl, this.atoken, this
								.rtoken); // 调用refersh.refersh函数获取新的Code
							if (newCode == 1000) { // 如果新的Code为1000
								this.atoken = uni.getStorageSync('atoken'); // 更新atoken值
								this.rtoken = uni.getStorageSync('rtoken'); // 更新rtoken值
								this.getFriendsReq(); // 重新调用getFriendsReq函数
							} else {
								// err
							}
						} else if (data.data.Code == 1000) { // 如果返回的Code为1000
							console.log('Req'); // 打印日志：Req
							this.requester2 = []; // 清空requester2数组
							for (let i = 0; i < data.data.Data.length; i++) {
								this.requester2.push({ // 将数据添加到requester2数组中
									id: data.data.Data[i].Friendid,
									info: data.data.Data[i].FriendNick,
									name: data.data.Data[i].FriendName,
									img: 'data:image/png;base64,' + data.data.Data[i].FriendImg,
									time: data.data.Data[i].LastTime,
									tip: 10
								});
							}
						} else {
							// err
							console.log(data.data.Code);
						}
					}
				});
				console.log('get'); // 打印日志：get
			},

			rejectFriend: function(optid, index) {
				uni.request({
					url: config.myurl + '/friend/reject', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken
					}, // 请求头中添加Authorization字段
					data: {
						'user_id': optid, // 请求的用户ID
						'friend_id': this.user.id // 请求的朋友ID
					},
					success: async data => { // 请求成功的回调函数
						console.log(data.data); // 打印data.data的值
						if (data.data.Code == 1009) { // 如果返回的Code为1009
							let newCode = await refersh.refersh(config.myurl, this.atoken, this
								.rtoken); // 调用refersh.refersh函数获取新的Code
							if (newCode == 1000) { // 如果新的Code为1000
								this.atoken = uni.getStorageSync('atoken'); // 更新atoken值
								this.rtoken = uni.getStorageSync('rtoken'); // 更新rtoken值
								this.rejectFriend(); // 重新调用rejectFriend函数
							} else {
								// err
							}
						} else if (data.data.Code == 1000) { // 如果返回的Code为1000
							console.log(this.requester2[index]); // 打印requester2数组中索引为index的元素
							console.log(index); // 打印index的值
							this.requester2.splice(index, 1); // 从requester2数组中移除索引为index的元素
							console.log(this.requester2); // 打印更新后的requester2数组
						} else {
							// err
							console.log(data.data.Code);
						}
					}
				});
			},

			acceptFriend: function(optid, index) {
				uni.request({
					url: config.myurl + '/friend/accept', // 请求的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken
					}, // 请求头中添加Authorization字段
					data: {
						'user_id': optid, // 请求的用户ID
						'friend_id': this.user.id // 请求的朋友ID
					},
					success: async data => { // 请求成功的回调函数
						console.log(data.data); // 打印data.data的值
						if (data.data.Code == 1009) { // 如果返回的Code为1009
							let newCode = await refersh.refersh(config.myurl, this.atoken, this
								.rtoken); // 调用refersh.refersh函数获取新的Code
							if (newCode == 1000) { // 如果新的Code为1000
								this.atoken = uni.getStorageSync('atoken'); // 更新atoken值
								this.rtoken = uni.getStorageSync('rtoken'); // 更新rtoken值
								this.rejectFriend(); // 重新调用rejectFriend函数
							} else {
								// err
							}
						} else if (data.data.Code == 1000) { // 如果返回的Code为1000
							console.log(this.requester2[index]); // 打印requester2数组中索引为index的元素
							console.log(index); // 打印index的值
							this.requester2.splice(index, 1); // 从requester2数组中移除索引为index的元素
							console.log(this.requester2); // 打印更新后的requester2数组
						} else {
							// err
							console.log(data.data.Code);
						}
					}
				});
			}

		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	.top-bar {
		background: rgba(255, 255, 255, 0.96);
		border-bottom: 1px solid $uni-border-color;
		margin-bottom: 30rpx;
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 188rpx;
		width: 100%;
		padding: var(--status-bar-height) $uni-spacing-col-base 0 $uni-spacing-col-base;
		box-sizing: border-box;

		.requestLine {
			background: rgba(255, 255, 255, 1);
			box-shadow: 0rpx 24rpx 64rpx -8rpx rgba(0, 0, 0, 0.1);
			border-radius: $uni-border-radius-base;
			margin-bottom: 112rpx;
			width: 100%;
			padding: $uni-spacing-row-sm $uni-spacing-col-base;
			box-sizing: border-box;
		}

		.requestTop {
			display: flex;
			flex-direction: cow;
			align-items: flex-end;
			justify-content: space-around;
			width: 100%;
			height: 144rpx;
			margin-top: -78rpx;

			.btn {
				width: 160rpx;
				height: 64rpx;
				border-radius: $uni-border-radius-lg;
				text-align: center;
				line-height: 64rpx;
				font-size: $uni-font-size-lg;
			}

			.reject {
				color: $uni-color-warning;
				background: rgba(255, 93, 91, 0.1);
			}

			.accept {
				color: $uni-text-color;
				background: $uni-color-primary;
			}

			.requestImg {
				text-align: center;

				image {
					width: 144rpx;
					height: 144rpx;
					border-radius: 50%;
				}
			}
		}

		.requestCenter {
			text-align: center;
			padding-top: 18rpx;
			padding-bottom: 40rpx;

			.name {
				font-size: 46rpx;
				font-weight: 500;
				color: $uni-text-color;
			}

			.time {
				font-size: $uni-font-size-sm;
				color: $uni-text-color-disable;
			}
		}

		.requestBottom {
			.message {
				padding: $uni-spacing-row-sm $uni-spacing-col-base;
				box-sizing: border-box;
				border-radius: $uni-border-radius-base;
				background: $uni-bg-color-grey;
				font-size: $uni-font-size-base;
				color: $uni-text-color;
				line-height: 40rpx;
			}
		}
	}
</style>
