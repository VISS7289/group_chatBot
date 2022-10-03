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
			</view>
			<view class="modifyInp" v-if="modifyInfo.type==3">
				<view class="inputs-div">
					<input type="text" placeholder="请输入邮箱" class="paw" placeholder-style="color:#bbb;font-weight:400;"
						@blur="isEmail" />
					<view class="wText" v-if="occupyEmail">邮箱已被占用</view>
					<view class="wText" v-if="invilid">邮箱无效</view>
					<image src="../../static/general/right.png" class="ok" v-if="ok"></image>
				</view>
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
				<textarea v-model="modifyInfo.userinfo.nick" class="modifyContent" @blur="isUserNick"></textarea>
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
				infoType: ['昵称', '密码', '简介', '邮箱', '电话', '备注', '添加好友'],
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
			}
		},
		onLoad: function(option) {
			this.modifyInfo = JSON.parse(decodeURIComponent(option.modifyInfo))
			if (this.modifyInfo.type == 6) {
				this.userRequest = this.modifyInfo.userRequest
				this.ok = true
			}
			try {
				this.atoken = uni.getStorageSync('atoken')
				this.rtoken = uni.getStorageSync('rtoken')
			} catch (e) {
				// error
			}
		},
		methods: {
			backOne: function() {
				uni.navigateBack({ delta: 1 })
			},
			infoChange: function() {
				if (this.ok) {
					if (this.modifyInfo.type != 6) {
						uni.navigateTo({
							url: '../userdetail/userdetail?user=' + encodeURIComponent(JSON.stringify(this
								.modifyInfo.userinfo))
						})
					} else {
						this.addFriend()
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
					if (e.detail.cursor >= 6) {
						this.ok = true
						this.modifyInfo.userinfo.name = e.detail.value
						this.shor = false
					} else {
						this.ok = false
						this.short = true
					}
				}
			},
			isEmail: function(e) {
				let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
				if (e.detail.cursor > 0) {
					if (reg.test(e.detail.value)) {
						this.ok = true
						this.modifyInfo.userinfo.email = e.detail.value
						this.invilid = false
					} else {
						this.ok = false
						this.invilid = true
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
					} else {
						this.okRe = false
						this.shortRe = true
					}
				}
			},
			isUserNick: function(e) {
				if (e.detail.cursor > 0) {
					this.ok = true
					this.modifyInfo.userinfo.nick = e.detail.value
				} else {
					this.ok = false
				}
			},
			subInfo: function() {
				if (this.okUsername && this.okEmail && this.okPa && this.okRe) {
					console.log(this.userValue)
					console.log(this.emailValue)
					console.log(this.userPassword)
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
