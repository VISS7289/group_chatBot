<template>
	<view>
		<scroll-view class="scrollY" :style="{height:pageHeight+'px'}">
			<view class="content">
				<view class="top-bar">
					<view class="top-bar-left" @tap="backOne">
						<image src="../../static/general/return.png"></image>
					</view>
					<view class="top-bar-right" v-if="relation==0||relation==3">
						<view class="more" @tap="Touserdetail">
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
				<view class="buttomText" @tap="getUserDetail" v-if="user.state==0">发送消息</view>
				<view class="buttomText" @tap="addFriendAnmi" v-if="user.state==1">申请中</view>
				<view class="buttomText" @tap="addFriendAnmi" v-if="user.state==2">加为好友</view>
			</view>
			<view class="addFriend" :animation="animationData">
				<view class="titleUser">{{user.name}}</view>
				<textarea :placeholder="myname+'请求加为好友~'" maxlength="120" class="addText" @blur="AddFriend"></textarea>
			</view>
			<view class="addButton" :animation="animationData1">
				<view class="close" @tap="addFriendAnmi">取消</view>
				<view class="send" @tap="addFriend">发送</view>
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
				dis: false, // 控制某个元素是否可见的标志位
				myname: 'duan', // 当前用户的名字
				animationData: {}, // 动画数据对象
				animationData1: {}, // 动画数据对象1
				animationData2: {}, // 动画数据对象2
				animationData3: {}, // 动画数据对象3
				animationData4: {}, // 动画数据对象4
				isAdd: false, // 是否添加的标志位
				isDelete: false, // 是否删除的标志位
				pageTop: 0, // 页面顶部的位置
				imgTop: 0, // 图片顶部的位置
				pageHeight: 0, // 页面的高度
				user: {}, // 用户对象
				userDetail: {}, // 用户详细信息对象
				userid: '', // 用户ID
				userNick: '', // 用户昵称
				atoken: '', // 访问令牌
				rtoken: '', // 刷新令牌
				relation: 0, // 用户关系状态，0表示好友，1表示申请中，2表示非好友，3表示自己
				userRequest: '', // 用户请求
			}
		},

		// 初始化用户
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
					uni.navigateTo({
						url: '../signin/signin',
					})
				}
			} catch (e) {
				// error
			}
			this.userRequest = this.myname + '请求加为好友~'
		},
		onReady: function() {
			this.getElementStyle()
		},
		methods: {
			refersh: function(parm) {
				this.user.img = parm.img // 更新用户头像
				this.user.name = parm.name // 更新用户姓名
				this.userDetail.Intr = parm.intr // 更新用户简介
				this.userDetail.Gender = parm.sex // 更新用户性别
				this.userNick = parm.nick // 更新用户昵称
			},

			backOne: function() {
				if (!this.isDelete) { // 如果没有删除操作
					uni.navigateBack({
						delta: 1
					}) // 返回上一页
				} else {
					let pages = getCurrentPages() // 获取当前页面栈
					let beforePage = pages[pages.length - 2] // 获取上一页
					uni.navigateBack({
						delta: 1,
						success: function() {
							beforePage.DeleteSuccess() // 执行上一页的DeleteSuccess方法
						}
					})
				}
			},
			// 删除好友
			delfriend: function() {
				this.user.state = 2 // 将用户状态设置为2，表示非好友
				this.isDelete = true // 将isDelete标志位设置为true，表示进行了删除操作
				this.userNick = this.user.name // 将用户昵称设置为用户姓名
			},
			// 讲用户头像设为背景
			getElementStyle: function() {
				const query = uni.createSelectorQuery().in(this) // 创建选择器查询实例
				query.select('.background').boundingClientRect(data => {
					this.pageHeight = data.height // 获取背景元素的高度，并将其赋值给pageHeight属性
				}).exec()
				query.select('.userImg').boundingClientRect(data => {
					this.imgTop = data.top // 获取用户头像元素的top值，并将其赋值给imgTop属性
				}).exec()
				this.pageTop = uni.getSystemInfoSync().statusBarHeight // 获取系统状态栏的高度，并将其赋值给pageTop属性
			},
			// 获取好友昵称
			getUserNick: function() {
				if (this.userid == this.user.id) {
					this.userNick = this.user.name // 如果userid等于当前用户id，将用户昵称设置为用户姓名
					this.relation = 3 // 将关系状态设置为3，表示自己
					return
				}
				uni.request({
					url: config.myurl + '/serch/isfriend', // 请求好友关系接口的URL
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					}, // 请求头部，包含访问令牌
					data: {
						'user_id': this.userid, // 请求参数，当前用户id
						'friend_id': this.user.id // 请求参数，好友id
					},
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getUserNick()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							this.relation = data.data.Data.State // 将关系状态设置为返回的好友关系状态
							if (data.data.Data.State == 0) { // 如果关系状态为0，表示好友
								if (data.data.Data.Markname.length != 0) { // 如果有备注名
									console.log(this.userNick)
									this.userNick = data.data.Data.Markname // 将备注名设置为用户昵称
								} else {
									this.userNick = this.user.name // 否则，将用户姓名设置为用户昵称
								}
							} else {
								this.userNick = this.user.name // 如果关系状态非0，将用户姓名设置为用户昵称
							}
						} else {
							// err
						}
					}
				})
			},

			// 为页面加载动效
			addFriendAnmi: function() {
				this.isAdd = !this.isAdd // 切换isAdd标志位的值，用于控制动画效果的显示与隐藏

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

				if (this.isAdd) { // 如果isAdd为true，执行展开动画
					animation.bottom(0).step() // 动画1：底部上移至0位置
					animation1.bottom(0).step() // 动画2：底部上移至0位置
					animation2.height('240rpx').width('240rpx').top(0.17 * this.pageHeight - 2 * this.pageTop + 'px')
						.left('-30%').step() // 动画3：设置高度、宽度、位置
					animation3.opacity(0).top(0.27 * this.pageHeight - 2 * this.pageTop + 'px').left('10%')
						.step() // 动画4：设置透明度、位置
					animation4.backgroundColor('rgba(255,248,49,0.3)').step() // 动画5：设置背景颜色
				} else { // 如果isAdd为false，执行收起动画
					animation.bottom('-75%').step() // 动画1：底部下移至-75%位置
					animation1.bottom('-10%').step() // 动画2：底部下移至-10%位置
					animation2.height('400rpx').width('400rpx').top(this.imgTop - 2 * this.pageTop + 'px').left(0)
						.step() // 动画3：设置高度、宽度、位置
					animation3.opacity(1).top('465rpx').left('324rpx').step() // 动画4：设置透明度、位置
					animation4.backgroundColor('rgba(255,248,49,0)').step() // 动画5：设置背景颜色
				}

				// 将动画效果导出并赋值给相应的数据属性
				this.animationData = animation.export()
				this.animationData1 = animation1.export()
				this.animationData2 = animation2.export()
				this.animationData3 = animation3.export()
				this.animationData4 = animation4.export()
			},

			// 获取用户详情
			getUserDetail: function() {
				console.log(666) // 打印日志，用于调试
				uni.request({
					url: config.myurl + '/user/detial', // 请求用户详情接口的URL
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					}, // 请求头部，包含访问令牌
					data: {
						'id': this.user.id
					}, // 请求参数，用户id
					success: async data => {
						console.log(data.data) // 打印返回的数据，用于调试
						if (data.data.Code == 1009) { // 如果返回的Code等于1009，表示访问令牌过期
							let newCode = await refersh.refersh(config.myurl, this.atoken, this
								.rtoken) // 刷新访问令牌
							if (newCode == 1000) { // 如果刷新成功
								this.atoken = uni.getStorageSync('atoken') // 更新访问令牌
								this.rtoken = uni.getStorageSync('rtoken') // 更新刷新令牌
								this.getUserDetail() // 重新调用getUserDetail方法
							} else {
								// err
							}
						} else if (data.data.Code == 1000) { // 如果返回的Code等于1000，表示请求成功
							this.userDetail = data.data.Data // 将返回的用户详情数据赋值给userDetail
						} else {
							// err
						}
					}
				})
			},

			// 添加好友请求
			addFriend: function() {
				uni.request({
					url: config.myurl + '/friend/request', // 请求添加好友接口的URL
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken // 请求头部，包含访问令牌
					},
					data: {
						'user_id': this.userid, // 当前用户的id
						'friend_id': this.user.id, // 好友的id
						'msg': this.userRequest, // 好友请求的消息
					},
					success: async data => {
						if (data.data.Code == 1009) { // 如果返回的Code等于1009，表示访问令牌过期
							let newCode = await refersh.refersh(config.myurl, this.atoken, this
								.rtoken) // 刷新访问令牌
							if (newCode == 1000) { // 如果刷新成功
								this.atoken = uni.getStorageSync('atoken') // 更新访问令牌
								this.rtoken = uni.getStorageSync('rtoken') // 更新刷新令牌
								this.addFriend() // 重新调用addFriend方法
							} else {
								// err
							}
						} else if (data.data.Code == 1000) { // 如果返回的Code等于1000，表示请求成功
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
									beforePage.AddSuccess() // 执行上一页的AddSuccess方法
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

			// 请求添加好友
			AddFriend: function(e) {
				this.ok = true
				if (e.detail.cursor > 0) {
					this.userRequest = e.detail.value
				} else {
					this.userRequest = this.myname + '请求加为好友~'
				}

			},
			//跳转用户详情页面
			Touserdetail: function() {
				console.log(this.user.id)
				uni.navigateTo({
					url: '../userdetail/userdetail?id=' + this.user.id
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
