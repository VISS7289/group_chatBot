<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="toSignIn">
				<image src="../../static/general/return.png"></image>
			</view>
		</view>
		<view class="logo">
			<image src='../../static/general/logo.png'></image>
		</view>
		<view class="main">
			<view class="title">注册</view>
			<view class="slogan">您好，欢迎来到AI聊天室！</view>
			<view class="inputs">
				<view class="inputs-div">
					<input type="text" placeholder="请输入用户名" class="user" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isUsername" />
					<view class="wText" v-if="occupyUsername">用户名已被占用</view>
					<view class="wText" v-if="shortUsername">用户名过短</view>
					<image src="../../static/general/right.png" class="ok" v-if="okUsername"></image>
				</view>
				<view class="inputs-div">
					<input type="text" placeholder="请输入邮箱" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isEmail" />
					<view class="wText" v-if="occupyEmail">邮箱已被占用</view>
					<view class="wText" v-if="invilid">邮箱无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="okEmail"></image>
				</view>
				<view class="inputs-div">
					<input :type="type" placeholder="请输入密码" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isUserPassword" />
					<view class="wTextPassword" v-if="shortPassword">密码过短</view>
					<image src="../../static/general/right.png" class="okPs" v-if="okPa"></image>
					<image src="../../static/general/look.png" class="look" v-if="look" @tap="lookPassword"></image>
					<image src="../../static/general/unlook.png" class="look" v-if="!look" @tap="lookPassword"></image>
				</view>
				<view class="inputs-div">
					<input :type="type" placeholder="请重新输入密码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="isUserPasswordRe" />
					<view class="wTextPassword" v-if="shortRe">密码错误</view>
					<image src="../../static/general/right.png" class="okPs" v-if="okRe"></image>
					<image src="../../static/general/look.png" class="look" v-if="look" @tap="lookPassword"></image>
					<image src="../../static/general/unlook.png" class="look" v-if="!look" @tap="lookPassword"></image>
				</view>
				<view class="vertify-div">
					<input :type="type" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi" />
					<view class="vertifySend" @tap="sendVertify">
						<view v-if="send">发送验证码</view>
						<view v-if="unSend" class="unSelect">{{this.time}}s</view>
						<view v-if="reSend">重新发送</view>
					</view>
				</view>
			</view>
			<view class="tips" v-if="wrong">{{this.errInfo}}</view>
		</view>
		<view class="submit" @tap="subInfo">注册</view>
	</view>
</template>

<script>
	import config from '../../commons/js/config.js'
	export default {
		data() {
			return {
				type: 'password', // 密码输入框的类型
				occupyUsername: false, // 是否占用用户名（用于验证用户名是否已被占用）
				okUsername: false, // 用户名是否合法
				occupyEmail: false, // 是否占用电子邮件地址（用于验证电子邮件地址是否已被占用）
				invilid: false, // 输入是否无效（用于验证输入是否有效）
				okEmail: false, // 电子邮件地址是否合法
				shortUsername: false, // 用户名是否过短
				shortPassword: false, // 密码是否过短
				look: false, // 是否显示密码（用于切换密码输入框的可见性）
				shortRe: false, // 重新输入的密码是否过短
				okRe: false, // 重新输入的密码是否匹配原密码
				okPa: false, // 密码是否合法
				userValue: '', // 用户名输入框的值
				emailValue: '', // 电子邮件地址输入框的值
				userPassword: '', // 密码输入框的值
				userPasswordRe: '', // 重新输入的密码输入框的值
				verifiValue: '', // 验证码输入框的值
				wrong: false, // 输入是否错误
				send: true, // 是否发送验证码
				reSend: false, // 是否重新发送验证码
				unSend: false, // 是否无法发送验证码
				time: 60, // 验证码倒计时的时间（单位：秒）
				errInfo: '' // 错误信息
			}
		},

		methods: {
			toSignIn: function() {
				uni.navigateTo({
					url: '../signin/signin'
				}) // 导航到登录页面
			},

			lookPassword: function() {
				if (this.look) {
					this.type = 'password' // 如果密码输入框当前是可见的，将其类型设置为密码类型（隐藏密码）
				} else {
					this.type = 'text' // 如果密码输入框当前是隐藏的，将其类型设置为文本类型（显示密码）
				}
				this.look = !this.look // 切换密码输入框的可见性状态
			},
			// 用户名是否合法
			isUsername: function(e) {
				if (e.detail.cursor > 0) { // 检查光标位置，确保输入不为空
					this.userValue = e.detail.value // 获取用户名输入框的值
					if (e.detail.cursor >= 6) { // 检查用户名长度是否大于等于6
						this.shortUsername = false // 用户名长度合法
						uni.request({
							url: config.myurl + '/existName?username=' + this
								.userValue, // 向服务器发起验证用户名是否已被占用的请求
							method: 'GET',
							success: data => {
								console.log(data.data)
								if (data.data.Code == 1002) { // 判断服务器返回的结果码，1002表示用户名已被占用
									this.occupyUsername = true // 用户名已被占用
								} else {
									this.occupyUsername = false // 用户名未被占用
								}
								this.okUsername = !(this.occupyUsername || this
									.shortUsername) // 检查用户名的合法性，不能被占用且长度合法
							},
						})
					} else {
						this.shortUsername = true // 用户名长度过短
						this.occupyUsername = false // 用户名未被占用
					}
				}
			},
			// 邮箱是否合法
			isEmail: function(e) {
				let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
				if (e.detail.cursor > 0) { // 检查光标位置，确保输入不为空
					this.emailValue = e.detail.value // 获取邮箱输入框的值
					if (reg.test(this.emailValue)) { // 使用正则表达式验证邮箱格式是否合法
						this.okEmail = true // 邮箱合法
						this.invilid = false // 输入无效（即非法字符）
					} else {
						this.okEmail = false // 邮箱不合法
						this.invilid = true // 输入无效（即非法字符）
					}
				}
			},
			// 密码是否合法
			isUserPassword: function(e) {
				if (e.detail.cursor > 0) { // 检查光标位置，确保输入不为空
					this.userPassword = e.detail.value // 获取密码输入框的值
					if (e.detail.cursor >= 6) { // 检查密码长度是否大于等于6
						this.okPa = true // 密码合法
						this.shortPassword = false // 密码长度合法
						if (this.userPasswordRe && this.userPasswordRe != this.userPassword) { // 检查重新输入的密码是否存在且与原密码不匹配
							this.okRe = false // 重新输入的密码不匹配
							this.shortRe = true // 重新输入的密码长度合法
						}
					} else {
						this.okPa = false // 密码不合法
						this.shortPassword = true // 密码长度过短
					}
				}
			},
			// 验证密码是否合法
			isUserPasswordRe: function(e) {
				if (e.detail.cursor > 0) { // 检查光标位置，确保输入不为空
					this.userPasswordRe = e.detail.value // 获取重新输入的密码输入框的值
					if (e.detail.cursor >= 6 && this.userPasswordRe == this
						.userPassword) { // 检查重新输入的密码长度是否大于等于6且与原密码匹配
						this.okRe = true // 重新输入的密码合法且与原密码匹配
						this.shortRe = false // 重新输入的密码长度合法
					} else {
						this.okRe = false // 重新输入的密码不合法或与原密码不匹配
						this.shortRe = true // 重新输入的密码长度合法
					}
				}
			},

			// 输入验证码
			inputVerifi: function(e) {
				this.verifiValue = e.detail.value
			},
			// 注册
			subInfo: function() {
				this.wrong = !this.wrong // 切换错误状态，将错误状态设置为相反的值

				if (this.okUsername && this.okEmail && this.okPa && this.okRe) { // 检查用户名、邮箱、密码和重新输入的密码是否都合法
					uni.request({
						url: config.myurl + '/register', // 向服务器发送注册请求的URL
						method: 'POST',
						data: {
							'username': this.userValue, // 提交的用户名
							'password': this.userPassword, // 提交的密码
							're_password': this.userPasswordRe, // 提交的重新输入的密码
							'email': this.emailValue, // 提交的邮箱
							'verifiCode': this.verifiValue // 提交的验证码
						},
						success: data => {
							console.log(data.data)
							if (data.data.Code != 1000) { // 判断服务器返回的结果码，非1000表示注册失败
								this.wrong = true // 设置错误状态为true
								this.errInfo = data.data.Msg // 设置错误信息为服务器返回的错误信息
							} else {
								this.wrong = false // 设置错误状态为false
								let res = data.data.Data // 从服务器返回的数据中获取注册成功的用户信息
								try {
									uni.setStorageSync('user', { // 将用户信息存储到本地缓存中
										'id': res.UserID,
										'img': res.Img,
										'email': res.Email,
										'name': res.Username
									})
									uni.setStorageSync('atoken', res.AToken) // 存储访问令牌到本地缓存中
									uni.setStorageSync('rtoken', res.RToken) // 存储刷新令牌到本地缓存中
								} catch (e) {
									console.log('数据存储出错')
									console.log(e)
								}
								console.log(res)
								uni.navigateTo({ // 导航到首页
									url: '../index/index',
								})
							}
						}
					})
				} else {
					// 如果用户名、邮箱、密码或重新输入的密码有不合法的情况，则设置相应的状态为true，以提示用户输入错误
					this.shortUsername = !this.okUsername // 用户名长度过短
					this.invilid = !this.okEmail // 邮箱不合法
					this.shortPassword = !this.okPa // 密码长度过短
					this.shortRe = !this.okRe // 重新输入的密码长度过短
				}
			},

			// 发送验证码
			sendVertify: function() {
				if (this.okEmail && !this.unSend) { // 检查邮箱是否合法且未发送验证码
					this.send = false // 设置发送状态为false
					this.reSend = false // 设置重新发送状态为false
					this.unSend = true // 设置不发送状态为true
					this.time = 60 // 设置倒计时时间为60秒

					var obj = setInterval(() => { // 创建计时器
						if (this.time <= 0) { // 如果倒计时时间到达0
							this.time = 0 // 设置倒计时时间为0
							this.reSend = true // 设置重新发送状态为true
							this.unSend = false // 设置不发送状态为false
							clearInterval(obj) // 清除计时器
						} else {
							this.time = this.time - 1 // 将倒计时时间减1
						}
					}, 1000) // 每隔1秒执行一次计时器函数

					uni.request({
						url: config.myurl + '/verificationCode', // 向服务器发送获取验证码的请求的URL
						method: 'POST',
						data: {
							'to': this.emailValue, // 需要发送验证码的邮箱
							'username': this.userValue // 用户名
						},
						success: data => {
							console.log(data.data)
							// 登录成功
							if (data.data.Code == 1000) {
								uni.setStorageSync('aToken', data.data.Data[0]) // 将访问令牌存储到本地缓存中
								uni.setStorageSync('rToken', data.data.Data[1]) // 将刷新令牌存储到本地缓存中
							}
						}
					})
				} else {
					this.invilid = !this.okEmail // 如果邮箱不合法，则设置输入无效的状态为true，以提示用户输入错误
				}
			},

			timeReduce: function(t) {
				return // 这个方法没有实际操作，直接返回
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
				font-weight: 500;
				color: $uni-text-color;
				line-height: 80rpx;
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
		height: 850rpx;

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

	.vertify-div {
		display: flex;
		flex-direction: cow;
		align-items: center;

		.vertifySend {
			padding-top: 48rpx;
			padding-left: $uni-spacing-row-base;
			font-weight: 500;
			color: $uni-text-color;
			width: 200rpx;
			text-align: center;
		}

		.unSelect {
			color: $uni-text-color-disable;
		}
	}

	.inputs-div {
		position: relative;

		.wText,
		.wTextPassword {
			position: absolute;
			right: 0;
			top: 70rpx;
			font-size: $uni-font-size-base;
			font-weight: 500;
			color: $uni-color-warning;
			float: right;
		}

		image {
			position: absolute;
			right: 0;
			top: 70rpx;
			float: right;
			width: 40rpx;
			height: 40rpx;
		}

		.okPs {
			right: 40rpx;
		}

		.wTextPassword {
			right: 45rpx;
		}


	}
</style>
