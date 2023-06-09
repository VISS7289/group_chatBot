<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-right" @tap="toSignUp">
				<view class="text">注册</view>
			</view>
		</view>
		<view class="logo">
			<image src='../../static/general/logo.png'></image>
		</view>
		<view class="main">
			<view class="title">登录</view>
			<view class="slogan">您好，欢迎来到AI聊天室！</view>
			<!-- <view class="inputs">
				<input type="text" placeholder="用户名/邮箱" class="user" placeholder-style="color:#bbb;font-weight:400;" @blur="subUsername"/>
				<input type="password" placeholder="密码" class="paw" placeholder-style="color:#bbb;font-weight:400;" @blur="subPassword"/>
			</view> -->
			<view class="inputs">
				<input type="text" v-model="username" placeholder="用户名/邮箱" class="user"
					placeholder-style="color:#bbb;font-weight:400;" />
				<input type="password" v-model="password" placeholder="密码" class="paw"
					placeholder-style="color:#bbb;font-weight:400;" />
			</view>
			<view class="tips" v-if="wrong">{{this.errInfo}}</view>
		</view>
		<view class="submit" @tap="subInfo">登录</view>
		<!-- <view class="submit" @tap="Test">Test</view> -->
	</view>
</template>

<script>
	import config from '../../commons/js/config.js'
	import common from '../../commons/js/common.js'
	export default {
		data() {
			return {
				wrong: false,
				username: '',
				password: '',
				errInfo: '输入用户名或密码错误！'
			}
		},

		methods: {
			// 页面跳转
			toSignUp: function() {
				uni.navigateTo({
					url: '../signup/signup',
				})
			},
			// 更新用户名
			subUsername: function(e) {
				this.username = e.detail.value
			},
			// 更新密码
			subPassword: function(e) {
				this.password = e.detail.value
			},
			// 登录
			subInfo: function() {
				if (this.username && this.password) {
					// 发起登录请求
					uni.request({
						url: config.myurl + '/login', // 登录接口的URL
						method: 'POST', // 请求方法为POST
						data: {
							'username': this.username, // 提交用户名
							'password': this.password, // 提交密码
						},
						success: data => {
							// 请求成功回调函数
							// console.log(data.data)
							if (data.data.Code != 1000) {
								// 登录失败，显示错误信息
								this.wrong = true
								this.errInfo = data.data.Msg
							} else {
								// 登录成功
								this.wrong = false
								let res = data.data.Data
								try {
									// 将用户信息存储到本地缓存
									uni.setStorageSync('user', {
										'id': res.UserID,
										'img': res.Img,
										'email': res.Email,
										'name': res.Username
									})
									uni.setStorageSync('atoken', res.AToken)
									uni.setStorageSync('rtoken', res.RToken)
								} catch (e) {
									console.log('数据存储出错')
									console.log(e)
								}
								console.log(res)
								// 跳转到首页
								uni.navigateTo({
									url: '../index/index'
								})
							}
						}
					})
				} else {
					// 用户名或密码为空，显示错误
					this.wrong = true
				}
			},

			Test: function() {
				console.log(uni.getStorageSync('aToken')) // 打印本地缓存中的aToken
				console.log(uni.getStorageSync('rToken')) // 打印本地缓存中的rToken

				// 发起ping请求
				uni.request({
					url: config.myurl + '/ping', // ping接口的URL
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + uni.getStorageSync('aToken')
					}, // 在请求头中添加授权信息
					success: data => {
						console.log(data.data) // 打印ping请求的返回数据
						if (data.data.Code == 1009) {
							// 若返回Code为1009，则表示需要刷新令牌
							common.refersh(config.myurl) // 调用刷新令牌的方法
						}
					}
				})
			}

		}
	}
</script>

<style lang="scss">
	@import "../../commons/css/mycss.scss";

	.top-bar {
		border-bottom: 1px solid $uni-border-color;

		.top-bar-right {
			.text {
				font-size: 40rpx;
			}
		}
	}

	.logo {
		text-align: center;

		image {
			padding-top: 192rpx;
			width: 194rpx;
			height: 100rpx;
		}
	}

	.main {
		padding: 54rpx $uni-spacing-row-lg 120rpx;
		height: 458rpx;

		.title {
			font-size: 56rpx;
			font-weight: 500;
			color: $uni-text-color;
			line-height: 80rpx;
		}

		.slogan {
			font-size: 40rpx;
			color: $uni-text-color-grey;
			line-height: 56rpx;
		}

		.inputs {
			padding-top: 8rpx;

			input {
				padding-top: 40rpx;
				height: 88rpx;
				font-size: $uni-font-size-lg;
				font-weight: 500;
				color: $uni-text-color;
				line-height: 88rpx;
				border-bottom: 1px solid $uni-border-color;
			}
		}

		.tips {
			float: left;
			font-size: $uni-font-size-lg;
			color: $uni-color-warning;
			line-height: 56rpx;
		}
	}

	.submit {
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
</style>
