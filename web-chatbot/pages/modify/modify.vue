<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">{{infoType[modifyInfo.type]}}</view>
			</view>
			<view class="top-bar-right" @tap="infoChange">
				<view class="textB">确认</view>
			</view>
		</view>
		<view class="main">
			<view class="modifyInp" v-if="modifyInfo.type==0">
				<view class="inputs-div">
					<input type="text" placeholder="请输入用户名" class="user" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isUsername" />
					<view class="wText" v-if="occupy">用户名已被占用</view>
					<view class="wText" v-if="short">用户名过短</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==1">
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
				<view class="vertify-div" v-if="!verifi">
					<input type="text" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi" />
					<view class="vertifySend" @tap="sendVertify">
						<view v-if="send">发送验证码</view>
						<view v-if="unSend" class="unSelect">{{this.time}}s</view>
						<view v-if="reSend">重新发送</view>
					</view>
				</view>
				<view class="tips" v-if="wrong">{{this.errInfo}}</view>
			</view>
			<view class="modifyText" v-if="modifyInfo.type==2">
				<textarea v-model="modifyInfo.userinfo.intr" :placeholder="userRequest" class="modifyContent"
					@blur="isUserIntr"></textarea>
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==3">
				<view class="vertify-div" v-if="!verifi">
					<input type="text" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi" />
					<view class="vertifySend" @tap="sendVertify">
						<view v-if="send">发送验证码</view>
						<view v-if="unSend" class="unSelect">{{this.time}}s</view>
						<view v-if="reSend">重新发送</view>
					</view>
				</view>
				<view class="submit" @tap="subInfo" v-if="!verifi">确认验证码</view>
				<view class="inputs-div" v-if="verifi">
					<input type="text" placeholder="请输入邮箱" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isEmail" />
					<view class="wText" v-if="occupyEmail">邮箱已被占用</view>
					<view class="wText" v-if="invilid">邮箱无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
				<view class="vertify-div" v-if="verifi">
					<input type="text" placeholder="请重新输入验证码" class="paw"
						placeholder-style="color:#bbb;font-weight:400;" @blur="inputVerifi2" />
					<view class="vertifySend" @tap="sendVertify2">
						<view v-if="send2">发送验证码</view>
						<view v-if="unSend2" class="unSelect">{{this.time2}}s</view>
						<view v-if="reSend2">重新发送</view>
					</view>
				</view>
				<view class="tips" v-if="wrong">{{this.errInfo}}</view>
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==4">
				<view class="inputs-div">
					<input type="text" placeholder="请输入电话号码" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isPhone" />
					<view class="wText" v-if="invilid">电话号码无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
			</view>
			<view class="modifyText" v-if="modifyInfo.type==5">
				<textarea v-model="modifyInfo.userinfo.nick" :placeholder="userRequest" class="modifyContent"
					@blur="isUserNick"></textarea>
			</view>
			<view class="modifyText" v-if="modifyInfo.type==6">
				<textarea v-model="modifyInfo.userRequest" class="modifyContent" :placeholder="userRequest"
					@blur="isAddFriend"></textarea>
			</view>
		</view>
	</view>
</template>

<script>
	import refersh from '../../commons/js/refershToken.js'
	import config from '../../commons/js/config.js'
	export default {
		data() {
			return {
				infoType: ['昵称', '密码', '签名', '邮箱', '电话', '备注', '添加好友'],
				modifyInfo: {
					type: 0, //0text 1pwd 2emain
					userinfo: {}
				},
				type: 'password',
				occupyUsername: false,
				okUsername: false,
				occupyEmail: false,
				invilid: false,
				okEmail: false,
				ok: false,
				look: false,
				short: false,
				occupy: false,
				okRe: false,
				okPa: false,
				userValue: '',
				emailValue: '',
				phoneValue: '',
				userPassword: '',
				userPasswordRe: '',
				userRequest: '',
				atoken: '',
				rtoken: '',
				typeRes: '',
				reqRes: '',
				userid: '',
				verifi: false,
				shortPassword: false,
				shortRe: false,
				verifiValue: '',
				verifiValue2: '',
				time: 0,
				send: true,
				reSend: false,
				unSend: false,
				send2: true,
				reSend2: false,
				unSend2: false,
				okEmail: false,
				errInfo: '',
				wrong: false
			}
		},
		onLoad: function(option) {
			this.modifyInfo = JSON.parse(decodeURIComponent(option.modifyInfo))
			switch (this.modifyInfo.type) {
				case 6:
					this.userRequest = this.modifyInfo.userRequest
					this.ok = true
					break
				case 2:
					this.userRequest = this.modifyInfo.userinfo.intr
					this.ok = true
					break
				case 0:
					this.userRequest = this.modifyInfo.userinfo.name
					break
				case 4:
					this.userRequest = this.modifyInfo.userinfo.phone
					break
				case 5:
					this.userRequest = this.modifyInfo.userinfo.nick
					this.ok = true
					break
				case 3:
					this.userRequest = this.modifyInfo.userinfo.email
					this.ok = false
					break
				case 1:
					this.ok = false
					break
			}
			try {
				this.atoken = uni.getStorageSync('atoken')
				this.rtoken = uni.getStorageSync('rtoken')
				const value = uni.getStorageSync('user')
				this.userid = value.id
			} catch (e) {
				// error
			}
		},
		methods: {
			backOne: function() {
				uni.navigateBack({ delta: 1 })
			},
			inputVerifi: function(e) {
				this.verifiValue = e.detail.value
			},
			inputVerifi2: function(e) {
				this.verifiValue2 = e.detail.value
			},
			infoChange: function() {
				var fn = () => {
					console.log(2020)
					try {
						let value = uni.getStorageSync('user')
						if (value) {
							value[this.typeRes] = this.reqRes
							// console.log(value)
							this.modifyInfo[this.typeRes] = this.reqRes
							uni.setStorageSync('user', value)

							let pages = getCurrentPages() // 当前页面
							let beforePage = pages[pages.length - 2]
							beforePage.$vm.refersh({
								'type': this.typeRes,
								'req': this.reqRes
							})
							uni.navigateBack({ delta: 1 })
						} else {
							uni.navigateTo({ url: '../signin/signin', })
						}
					} catch (e) {
						console.log(e)
					}
				}
				console.log(this.ok)
				if (this.ok) {
					switch (this.modifyInfo.type) {
						case 6:
							this.addFriend()
							break
						case 2:
							this.typeRes = 'intr'
							this.reqRes = this.modifyInfo.userinfo.intr
							this.update('intr', this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.intr, '',
								fn)
							break
						case 0:
							this.typeRes = 'name'
							this.reqRes = this.modifyInfo.userinfo.name
							this.update('username', this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.name, '',
								fn)
							break
						case 4:
							this.typeRes = 'phone'
							this.reqRes = this.modifyInfo.userinfo.phone
							this.update('phone', this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.phone, '',
								fn)
							break
						case 5:
							this.typeRes = 'nick'
							this.reqRes = this.modifyInfo.userinfo.nick
							this.updateNick(this.modifyInfo.userinfo.id, this.modifyInfo.userinfo.nick, fn)
							break
						case 3:
							this.subInfo2()
							break
						case 1:
							this.subInfoPsw()
							break
					}

				}
			},
			toSignIn: function() {
				uni.navigateTo({ url: '../signin/signin', })
			},
			lookPassword: function() {
				if (this.look) {
					this.type = 'password'
				} else {
					this.type = 'text'
				}
				this.look = !this.look
			},
			isUsername: function(e) {
				if (e.detail.cursor > 0) {
					this.modifyInfo.userinfo.name = e.detail.value
					if (e.detail.cursor >= 6) {
						this.short = false
						uni.request({
							url: config.myurl + '/existName?username=' + this.modifyInfo.userinfo.name,
							method: 'GET',
							success: data => {
								console.log(data.data)
								if (data.data.Code == 1002) {
									this.occupy = true
								} else {
									this.occupy = false
								}
								this.ok = !(this.occupy || this.short)
							},
						})
					} else {
						this.short = true
						this.occupy = false
					}
				}
			},
			isEmail: function(e) {
				let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
				if (e.detail.cursor > 0) {
					if (reg.test(e.detail.value)) {
						this.okEmail = true
						this.modifyInfo.userinfo.email = e.detail.value
						this.invilid = false
						this.ok = true
					} else {
						this.okEmail = false
						this.invilid = true
						this.modifyInfo.userinfo.email = this.userRequest
						this.ok = false
					}
				}
			},
			isPhone: function(e) {
				let reg = /^([0-9])/
				if (e.detail.cursor > 0) {
					if (reg.test(e.detail.value) && e.detail.cursor > 6) {
						this.ok = true
						this.modifyInfo.userinfo.phone = e.detail.value
						this.invilid = false
					} else {
						this.modifyInfo.userinfo.phone = this.userRequest
						this.ok = false
						this.invilid = true
					}
				}
			},
			isUserPassword: function(e) {
				if (e.detail.cursor > 0) {
					this.userPassword = e.detail.value
					if (e.detail.cursor >= 6) {
						this.okPa = true
						this.shortPassword = false
						if (this.userPasswordRe && this.userPasswordRe != this.userPassword) {
							this.okRe = false
							this.shortRe = true
						}
					} else {
						this.okPa = false
						this.shortPassword = true
					}
				}
			},
			isUserPasswordRe: function(e) {
				if (e.detail.cursor > 0) {
					this.userPasswordRe = e.detail.value
					if (e.detail.cursor >= 6 && this.userPasswordRe == this.userPassword) {
						this.okRe = true
						this.shortRe = false
						this.ok = true
					} else {
						this.okRe = false
						this.shortRe = true
						this.ok = false
					}
				}
			},
			isUserNick: function(e) {
				if (e.detail.cursor > 0) {
					this.modifyInfo.userinfo.nick = e.detail.value
				} else {
					this.modifyInfo.userinfo.nick = this.userRequest
				}
			},
			isAddFriend: function(e) {
				this.ok = true
				if (e.detail.cursor > 0) {
					this.modifyInfo.userRequest = e.detail.value
				} else {
					this.modifyInfo.userRequest = this.userRequest
				}

			},
			isUserIntr: function(e) {
				this.ok = true
				if (e.detail.cursor > 0) {
					this.modifyInfo.userinfo.intr = e.detail.value
				} else {
					this.modifyInfo.userinfo.intr = this.userRequest
				}
			},
			update(type, optid, data, psw, fn) {
				console.log(fn)
				uni.request({
					url: config.myurl + '/change/update',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': optid,
						'data': data,
						'type': type,
						'psw': psw
					},
					success: data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.update(type, optid, data, psw)
							} else {
								//err
							}
						} else if (data.data.Code == 1000) {
							console.log(1010)
							fn.apply()
						} else {
							//err
						}
					}
				})
			},
			updateNick(optid, msg, fn) {
				console.log(fn)
				uni.request({
					url: config.myurl + '/change/nick',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.userid,
						'friend_id': optid,
						'msg': msg
					},
					success: data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.updateNick(optid, msg, fn)
							} else {
								//err
							}
						} else if (data.data.Code == 1000) {
							console.log(1010)
							fn.apply()
						} else {
							//err
						}
					}
				})
			},
			addFriend: function() {
				uni.request({
					url: config.myurl + '/friend/request',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.modifyInfo.userid,
						'friend_id': this.modifyInfo.friendid,
						'msg': this.modifyInfo.userRequest,
					},
					success: data => {
						if (data.data.Code == 1009) {
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.addFriend()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							uni.showToast({
								title: '好友请求发送成功',
								icon: 'none',
								duration: 2000
							})

							let pages = getCurrentPages() // 当前页面
							let beforePage = pages[pages.length - 2]
							uni.navigateBack({
								delta: 1,
								success: function() {
									beforePage.AddSuccess() // 执行上一页的onLoad方法
								}
							})
						} else {
							uni.showToast({
								title: '好友请求发送失败',
								icon: 'none',
								duration: 2000
							})
						}
					}
				})
			},
			subInfo: function() {
				console.log(this.modifyInfo.userinfo.email)
				console.log(this.verifiValue)
				uni.request({
					url: config.myurl + '/verifiExam',
					method: 'POST',
					data: {
						'email': this.modifyInfo.userinfo.email,
						'verifiCode': this.verifiValue,
					},
					success: data => {
						console.log(data.data)
						if (data.data.Code != 1000) {
							this.wrong = true
							this.errInfo = data.data.Msg
						} else {
							//SUCCESS
							this.wrong = false
							this.verifi = true
						}
					}
				})
			},
			subInfo2: function(){
				console.log(this.verifiValue2)
				uni.request({
					url: config.myurl + '/change/email',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.modifyInfo.userinfo.id,
						'n_email': this.modifyInfo.userinfo.email,
						'verifiCode': this.verifiValue2
					},
					success: data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							console.log(this.atoken)
							let newCode = refersh.refersh(config.myurl, this.atoken, this.rtoken)
							console.log(this.atoken)
							if (newCode == 1000) {
								this.subInfo2()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							let value = uni.getStorageSync('user')
							if (value) {
								value.email = this.modifyInfo.userinfo.email
								this.modifyInfo.email = this.modifyInfo.userinfo.email
								uni.setStorageSync('user', value)
							
								let pages = getCurrentPages() // 当前页面
								let beforePage = pages[pages.length - 2]
								beforePage.$vm.refersh({
									'type': 'email',
									'req': this.modifyInfo.userinfo.email
								})
								uni.navigateBack({ delta: 1 })
							} else {
								uni.navigateTo({ url: '../signin/signin', })
							}
						}
					}
				})
			},
			subInfoPsw: function() {
				uni.request({
					url: config.myurl + '/change/psw',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.modifyInfo.userinfo.id,
						'email': this.modifyInfo.userinfo.email,
						'password': this.userPassword,
						're_password': this.userPasswordRe,
						'verifiCode': this.verifiValue
					},
					success: data => {
						console.log(data.data)
						if (data.data.Code != 1000) {
							this.wrong = true
							this.errInfo = data.data.Msg
						} else {
							//SUCCESS
							this.wrong = false
							this.verifi = true
							uni.navigateBack({ delta: 1 })
						}
					}
				})
			},
			sendVertify: function() {
				if (!this.unSend) {
					this.send = false
					this.reSend = false
					this.unSend = true
					this.time = 60
					var obj = setInterval(() => {
						if (this.time <= 0) {
							this.time = 0
							this.reSend = true
							this.unSend = false
							clearInterval(obj)
						} else {
							this.time = this.time - 1
						}
					}, 1000)
					uni.request({
						url: config.myurl + '/verificationCode',
						method: 'POST',
						data: {
							'to': this.modifyInfo.userinfo.email,
							'username': this.modifyInfo.userinfo.name
						},
						success: data => {
							console.log(data.data)
							
						}
					})
				} else {
					this.invilid = !this.okEmail
				}
			},
			sendVertify2: function() {
				if (!this.unSend) {
					this.send2 = false
					this.reSend2 = false
					this.unSend2 = true
					this.time2 = 60
					var obj = setInterval(() => {
						if (this.time2 <= 0) {
							this.time2 = 0
							this.reSend2 = true
							this.unSend2 = false
							clearInterval(obj)
						} else {
							this.time2 = this.time2 - 1
						}
					}, 1000)
					uni.request({
						url: config.myurl + '/verificationCode',
						method: 'POST',
						data: {
							'to': this.modifyInfo.userinfo.email,
							'username': this.modifyInfo.userinfo.name
						},
						success: data => {
							console.log(data.data)
						}
					})
				} else {
					this.invilid = !this.okEmail
				}
			},
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
	.tips {
		float: left;
		font-size: $uni-font-size-lg;
		color: $uni-color-warning;
		line-height: 56rpx;
	}
	.main {
		display: flex;
		flex-direction: column;
		align-items: center;
		margin-top: 88rpx;
		padding-top: var(--status-bar-height);
		width: 100%;
		padding: $uni-spacing-col-base $uni-spacing-col-base 0 $uni-spacing-col-base;
		box-sizing: border-box;

		.modifyText {
			width: 100%;
		}

		.modifyContent {
			height: 386rpx;
			width: 100%;
			background: $uni-bg-color-grey;
			border-radius: $uni-border-radius-base;
			font-size: $uni-font-size-lg;
			color: $uni-text-color;
			line-height: 44rpx;
			padding: 16rpx 22rpx;
			box-sizing: border-box;
		}

		.modifyInp {
			width: 100%;
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

		.vertify-div {
			display: flex;
			flex-direction: cow;
			align-items: center;
			justify-content: center;
			margin-bottom: 120rpx;

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
	}
</style>
