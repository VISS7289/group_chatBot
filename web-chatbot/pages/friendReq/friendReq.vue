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
			<view class="requestLine"  v-for="(item,index) in requester2" :key="index">
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
				requester:[
					
				],
				requester2:[],
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
					uni.navigateTo({ url: '../signin/signin', })
				}
				console.log(value)
			} catch (e) {
				// error
			}
			this.getRequester()
			this.getFriendsReq()
		},
		methods: {
			backOne: function(){
				let pages = getCurrentPages() // 当前页面
				let beforePage = pages[pages.length - 2]
				uni.navigateBack({
					delta: 1,
					success: function() {
						beforePage.getFriendsReq() 
						beforePage.getFriends2() 
					}
				})
			},
			changeTime:function(date){
				return calT.dateTime(date)
			},
			changeTime2: function(date) {
				return calT.dateTime2(date)
			},
			getRequester: function(){
				this.requester=datas.friends()
				for(let i=0;i<this.requester.length;i++){
					this.requester[i].imgurl='../../static/index/'+this.requester[i].imgurl
				}
				//console.log(this.friends);
			},
			getFriendsReq: function() {
				console.log('send')
				uni.request({
					url: config.myurl + '/friend/myFriend',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.user.id,
						'state': 1
					},
					success: async data => {
						console.log(typeof(data.data))
						console.log(data.data.Code)
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
							console.log('Req')
							this.requester2 = []
							for (let i = 0; i < data.data.Data.length; i++) {
								this.requester2.push({
									id: data.data.Data[i].Friendid,
									info: data.data.Data[i].FriendNick,
									name: data.data.Data[i].FriendName,
									img: 'data:image/png;base64,' + data.data.Data[i].FriendImg,
									time: data.data.Data[i].LastTime,
									tip: 10
								})
							}
						} else {
							//err
							console.log(data.data.Code)
						}
					}
				})
				console.log('get')
			},
			rejectFriend: function(optid,index) {
				uni.request({
					url: config.myurl + '/friend/reject',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': optid,
						'friend_id': this.user.id
					},
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.rejectFriend()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(this.requester2[index])
							console.log(index)
							this.requester2.splice(index, 1)
							console.log(this.requester2)
						} else {
							//err
							console.log(data.data.Code)
						}
					}
				})
			},
			acceptFriend: function(optid,index) {
				uni.request({
					url: config.myurl + '/friend/accept',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': optid,
						'friend_id': this.user.id
					},
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.rejectFriend()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(this.requester2[index])
							console.log(index)
							this.requester2.splice(index, 1)
							console.log(this.requester2)
						} else {
							//err
							console.log(data.data.Code)
						}
					}
				})
			}
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";
	.top-bar{
		background: rgba(255,255,255,0.96);
		border-bottom: 1px solid $uni-border-color;
		margin-bottom: 30rpx;
	}
	.main{
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 188rpx;
		width: 100%;
		padding: var(--status-bar-height) $uni-spacing-col-base 0 $uni-spacing-col-base;
		box-sizing: border-box;
		.requestLine{
			background: rgba(255,255,255,1);
			box-shadow: 0rpx 24rpx 64rpx -8rpx rgba(0,0,0,0.1);
			border-radius: $uni-border-radius-base;
			margin-bottom: 112rpx;
			width: 100%;
			padding: $uni-spacing-row-sm $uni-spacing-col-base;
			box-sizing: border-box;
		}
		.requestTop{
			display: flex;
			flex-direction: cow;
			align-items: flex-end;
			justify-content: space-around;
			width: 100%;
			height: 144rpx;
			margin-top: -78rpx;
			.btn{
				width: 160rpx;
				height: 64rpx;
				border-radius: $uni-border-radius-lg;
				text-align: center;
				line-height: 64rpx;
				font-size: $uni-font-size-lg;
			}
			.reject{
				color: $uni-color-warning;
				background: rgba(255,93,91,0.1);
			}
			.accept{
				color: $uni-text-color;
				background: $uni-color-primary;
			}
			.requestImg{
				text-align: center;
				image{
					width: 144rpx;
					height: 144rpx;
					border-radius: 50%;
				}
			}
		}
		.requestCenter{
			text-align: center;
			padding-top: 18rpx;
			padding-bottom: 40rpx;
			.name{
				font-size: 46rpx;
				font-weight: 500;
				color: $uni-text-color;
			}
			.time{
				font-size: $uni-font-size-sm;
				color: $uni-text-color-disable;
			}
		}
		.requestBottom{
			.message{
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
