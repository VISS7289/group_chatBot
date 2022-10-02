<template>
	<view>
		<scroll-view class="scrollY" :style="{height:pageHeight+'px'}">
			<view class="content">
				<view class="top-bar">
					<view class="top-bar-left" @tap="backOne">
						<image src="../../static/general/return.png"></image>
					</view>
					<view class="top-bar-right" v-if="relation==0||relation==3">
						<view class="more">
							<image src="../../static/general/more.png"></image>
						</view>
					</view>
				</view>
			</view>
			<view class="background">
				<view class="bgWight" :animation="animationData4"></view>
				<image :src="user.img" class="bgImg" mode="aspectFill"></image>
			</view>
			<view class="main">
				<view class="userHeader">
					<image :src="user.img" mode="aspectFill" class="userImg" :animation="animationData2"></image>
					<image src="../../static/general/man.png" class="userSex man" v-if="userDetail.Gender==0"
						:animation="animationData3"></image>
					<image src="../../static/general/woman.png" class="userSex woman" v-if="userDetail.Gender==1"
						:animation="animationData3"></image>
					<image src="../../static/general/alien.png" class="userSex alien" v-if="userDetail.Gender==2"
						:animation="animationData3"></image>
				</view>
				<view class="titleUser">{{userNick}}</view>
				<view class="nickUser">昵称：{{user.name}}</view>
				<!-- <view class="intrUser">{{user.intr}}</view> -->
				<view class="intrUser">{{userDetail.Intr}}</view>
			</view>
			<view class="buttomBig">
				<!-- <view class="buttomText" @tap="addFriendAnmi">加为好友</view> -->
				<view class="buttomText" @tap="getUserDetail" v-if="relation==0">发送消息</view>
				<view class="buttomText" @tap="getUserDetail" v-if="relation==1">申请中</view>
				<view class="buttomText" @tap="getUserDetail" v-if="relation==2">加为好友</view>
			</view>
			<view class="addFriend" :animation="animationData">
				<view class="titleUser">{{user.name}}</view>
				<textarea :value="myname+'请求加为好友~'" maxlength="120" class="addText"></textarea>
			</view>
			<view class="addButton" :animation="animationData1">
				<view class="close" @tap="addFriendAnmi">取消</view>
				<view class="send">发送</view>
			</view>
		</scroll-view>
	</view>
</template>

<script>
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	export default {
		data() {
			return {
				dis: false,
				myname: 'duan',
				animationData: {},
				animationData1: {},
				animationData2: {},
				animationData3: {},
				animationData4: {},
				isAdd: false,
				pageTop: 0,
				imgTop: 0,
				pageHeight: 0,
				user: {},
				userDetail: {},
				userid: 0,
				userNick: '',
				atoken: '',
				rtoken: '',
				relation: 0, //0好友 1申请中 2非好友 3自己
			}
		},
		onLoad: function(option) {
			this.user = JSON.parse(decodeURIComponent(option.user))
			try {
				const value = uni.getStorageSync('user')
				if (value) {
					this.userid = value.id
					this.atoken = uni.getStorageSync('atoken')
					this.rtoken = uni.getStorageSync('rtoken')
					this.getUserDetail()
					this.getUserNick()
				} else {
					uni.navigateTo({ url: '../signin/signin', })
				}
			} catch (e) {
				// error
			}
		},
		onReady: function() {
			this.getElementStyle()
		},
		methods: {
			backOne: function() {
				uni.navigateBack({ delta: 1 })
			},
			getElementStyle: function() {
				const query = uni.createSelectorQuery().in(this)
				query.select('.background').boundingClientRect(data => {
					this.pageHeight = data.height
				}).exec()
				query.select('.userImg').boundingClientRect(data => {
					this.imgTop = data.top
				}).exec()
				this.pageTop = uni.getSystemInfoSync().statusBarHeight
			},
			getUserNick: function() {
				if (this.userid == this.user.id) {
					this.userNick = this.user.name
					this.relation = 3
					return
				}
				uni.request({
					url: config.myurl + '/serch/isfriend',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.user.id,
						'friend_id': this.userid
					},
					success: data => {
						if (data.data.Code == 1009) {
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.getUserNick()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(data.data.Data)
							this.relation = data.data.Data.State
							if (data.data.Data.State == 0) {
								if(len(data.data.Data.Markname)!=0){
									this.userNick = data.data.Data.Markname
								}else{
									this.userNick = this.user.name
								}
								
							} else {
								this.userNick = this.user.name
							}
						} else {
							// err
						}
					}
				})
			},
			addFriendAnmi: function() {
				this.isAdd = !this.isAdd
				var animation = uni.createAnimation({
					duration: 300,
					timingFunction: 'ease',
				})
				var animation1 = uni.createAnimation({
					duration: 300,
					timingFunction: 'ease',
				})
				var animation2 = uni.createAnimation({
					duration: 300,
					timingFunction: 'ease',
				})
				var animation3 = uni.createAnimation({
					duration: 300,
					timingFunction: 'ease',
				})
				var animation4 = uni.createAnimation({
					duration: 300,
					timingFunction: 'ease',
				})
				if (this.isAdd) {
					animation.bottom(0).step()
					animation1.bottom(0).step()
					animation2.height('240rpx').width('240rpx').top(0.17 * this.pageHeight - 2 * this.pageTop + 'px')
						.left('-30%').step()
					animation3.opacity(0).top(0.27 * this.pageHeight - 2 * this.pageTop + 'px').left('10%').step()
					animation4.backgroundColor('rgba(255,248,49,0.3)').step()
				} else {
					animation.bottom('-75%').step()
					animation1.bottom('-10%').step()
					animation2.height('400rpx').width('400rpx').top(this.imgTop - 2 * this.pageTop + 'px').left(0)
						.step()
					animation3.opacity(1).top('465rpx').left('324rpx').step()
					animation4.backgroundColor('rgba(255,248,49,0)').step()
				}

				this.animationData = animation.export()
				this.animationData1 = animation1.export()
				this.animationData2 = animation2.export()
				this.animationData3 = animation3.export()
				this.animationData4 = animation4.export()

			},
			getUserDetail: function() {
				uni.request({
					url: config.myurl + '/user/detial',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: { 'id': this.user.id },
					success: data => {
						if (data.data.Code == 1009) {
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.getUserDetail()
							} else {
								// err
							}

						} else if (data.data.Code == 1000) {
							// console.log(data.data.Data)
							this.userDetail = data.data.Data
							console.log(this.userDetail)
						} else {
							// err
						}
					}
				})
			},
		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	.top-bar {
		.top-bar-right {
			padding-right: $uni-spacing-col-base;
		}
	}

	.scrollY {
		width: 100%;
	}

	.background {
		position: fixed;
		left: 0rpx;
		top: 0rpx;
		width: 100%;
		height: 100%;

		.bgWight {
			width: 100%;
			height: 100%;
			//background-color: #eee;
		}

		.bgImg {
			opacity: 0.8;
			position: absolute;
			z-index: -1;
			left: -20rpx;
			top: -20rpx;
			width: 120%;
			height: 120%;
			filter: blur(12px);
		}
	}

	.main {
		display: flex;
		flex-direction: column;
		align-items: center;
		padding-top: var(--status-bar-height);

		.userHeader {
			position: relative;
			padding-top: 548rpx;
			width: 400rpx;

			.userImg {
				position: absolute;
				width: 400rpx;
				height: 400rpx;
				top: 148rpx;
				border-radius: 48rpx;
				border: 6rpx solid rgba(255, 255, 255, 1);
				z-index: 10;
				box-shadow: 0rpx 36rpx 40rpx 0rpx rgba(39, 40, 50, 0.1);
			}

			.userSex {
				position: absolute;
				top: 465rpx;
				left: 324rpx;
				width: 64rpx;
				height: 64rpx;
				z-index: 2;
				border-radius: 36rpx;
				padding: 4rpx;
				z-index: 11;
			}

			.man {
				background-color: $uni-color-success;
			}

			.woman {
				background-color: $uni-color-warning;
			}

			.alien {
				background-color: $uni-color-primary;
			}
		}

		.titleUser {
			font-size: 52rpx;
			color: $uni-text-color;
			line-height: 72rpx;
			margin-top: 42rpx;
		}

		.nickUser {
			font-size: $uni-font-size-base;
			color: $uni-text-color;
			line-height: 40rpx;
		}

		.intrUser {
			margin-top: 20rpx;
			width: 552rpx;
			font-size: $uni-font-size-base;
			color: $uni-text-color;
		}
	}

	.addFriend {
		width: 100%;
		height: 75%;
		background-color: rgba(255, 255, 255, 1);
		border-radius: $uni-border-radius-lg $uni-border-radius-lg 0rpx 0rpx;
		z-index: 9;
		position: absolute;
		bottom: -75%;
		padding: 0rpx 56rpx;
		box-sizing: border-box;

		.titleUser {
			font-size: 52rpx;
			color: $uni-text-color;
			line-height: 72rpx;
			margin: 168rpx 0rpx 40rpx 0rpx;
		}

		.addText {
			width: 100%;
			height: 50%;
			border-radius: $uni-border-radius-base;
			background-color: $uni-bg-color-grey;
			padding: 18rpx 22rpx;
			box-sizing: border-box;
			font-size: $uni-font-size-lg;
			color: $uni-text-color;
		}
	}

	.addButton {
		position: absolute;
		bottom: -10%;
		//background-color: #eee;
		width: 100%;
		height: 10%;
		display: flex;
		flex-direction: column;
		align-items: center;
		padding: 12rpx $uni-spacing-col-base;
		box-sizing: border-box;
		z-index: 11;
		display: flex;
		flex-direction: row;
		justify-content: center;

		.send {
			width: 384rpx;
			height: 80rpx;
			border-radius: $uni-border-radius-sm;
			background-color: $uni-color-primary;
			text-align: center;
			line-height: 80rpx;
			font-size: $uni-font-size-lg;
			color: $uni-text-color;
			margin-left: $uni-spacing-col-base;
		}

		.close {
			width: 184rpx;
			height: 80rpx;
			border-radius: $uni-border-radius-sm;
			background-color: $uni-bg-color-hover;
			text-align: center;
			line-height: 80rpx;
			font-size: $uni-font-size-lg;
			color: $uni-text-color;
		}
	}
</style>
