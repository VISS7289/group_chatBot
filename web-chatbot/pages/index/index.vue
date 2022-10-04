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
			<view class="friends" v-for="(item,index) in friends" :key="index">
				<view class="friends-list">
					<view class="friends-list-l">
						<text class="tip" v-if="item.tip>0">{{item.tip}}</text>
						<image :src="item.imgurl"></image>
					</view>
					<view class="friends-list-r">
						<view class="top">
							<view class="name">{{item.name}}</view>
							<view class="time">{{changeTime(item.time)}}</view>
						</view>
						<view class="info">{{item.info}}</view>
					</view>
				</view>
			</view>
		</view>
		<view class="submit" @tap="Test">Test</view>
		<view class="submit" @tap="Test2">Test2</view>
		<image :src="imgData"></image>
		<!-- <image src="data:image/png;base64,{{imgData}}"></image> -->
	</view>
</template>

<script>
	import datas from '../../commons/js/datas.js'
	import calT from '../../commons/js/calTime.js'
	export default {
		data() {
			return {
				friends: [

				],
				user: {
					imgurl: '../../static/test/duan.png',
					id: 123,
					email: '123@12.com',
					img: '',
					name: 'duan',
					nick: 'duan',
					intr: '这个用户很懒，没有任何简介。这个用户很懒，没有任何简介。这个用户很懒，没有任何简介。'
				},
				atoken: '',
				rtoken: '',
				imgurl: '../../static/test/duan.png',
				}
		},
		onLoad() {
			this.getFriends()
		},
		onShow(){
			this.userInit()
		},
		methods: {
			changeTime: function(date) {
				return calT.dateTime(date)
			},
			getFriends: function() {
				this.friends = datas.friends()
				for (let i = 0; i < this.friends.length; i++) {
					this.friends[i].imgurl = '../../static/index/' + this.friends[i].imgurl
				}
				//console.log(this.friends);
			},
			toSerch: function() {
				uni.navigateTo({ url: '../serch/serch?user='+encodeURIComponent(JSON.stringify(this.user)), })
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
						this.atoken = value.atoken
						this.rtoken = value.rtoken
					}else{
						uni.navigateTo({ url: '../signin/signin', })
					}
					console.log(value)
				} catch (e) {
					// error
				}
			},
			Test: function() {
				var img = document.createElement('img')
				img.src = '../../static/test/1 (1).jpg'
				uni.request({
					url: img.src,
					method: 'GET',
					responseType: 'arraybuffer',
					success: ress => {
						let base64 = wx.arrayBufferToBase64(ress.data) //把arraybuffer转成base64 
						base64 = 'data:image/jpeg;base64,' + base64 //不加上这串字符，在页面无法显示的哦
						console.log(base64)
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
