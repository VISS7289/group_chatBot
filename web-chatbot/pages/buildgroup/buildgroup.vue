<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">创建群聊</view>
			</view>
		</view>
		<view class="main">
			<!-- 群头像 -->
			<view class="group-img">
				<view @click="select">
					<image :src="img" mode="aspectFill"></image>
				</view>
				<ksp-cropper mode="free" :width="200" :height="140" :maxWidth="1024" :maxHeight="1024" :url="url"
					@cancel="oncancel" @ok="onok"></ksp-cropper>
			</view>
			<!-- 群名字 -->
			<view class="group-name">
				<input class="group-name-input" type="text" placeholder="请输入群名称"
					placeholder-style="color:#bbb;font-weight:400;" />
			</view>
			<!-- 选择用户 -->
			<view class="friends">
				<view class="title">用户</view>
				<scroll-view class="user-scr" scroll-y="true" scroll-with-animation="true"
					:scroll-into-view="scrollToView">
					<view class="friend" v-for="(item,index) in friends" :key="index">
						<view class="friendline" @tap="choose(index)">
							<view class="friendselect">
								<image src="../../static/general/right.png" class="selectimg" mode="aspectFill"
									v-if="item.select"></image>
								<!-- <image src="../../static/general/right.png" class="selectimg" mode="aspectFill"></image> -->
							</view>
							<image :src="item.img" class="friendhead" mode="aspectFill"></image>
							<view class="friendname">{{item.name}}</view>
						</view>
					</view>
				</scroll-view>
			</view>
		</view>
		<view class="buttomBig">
			<view class="buttomText" @tap="getUserDetail" :class="{unselect:select}">创建({{selectnum}})</view>
		</view>
		<canvas style="width: 100px; height: 100px; position:fixed;left:100%;" canvas-id="cv" type="2d"></canvas>
	</view>
</template>

<script>
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	export default {
		data() {
			return {
				url: '',
				img: '../../static/index/group.png',
				friends: [],
				atoken: '',
				rtoken: '',
				user: {},
				scrollToView: '',
				selectnum: 0,
			}
		},
		onLoad() {
			this.userInit()
			this.getFriends()
		},
		computed: {
			select: function() {
				return this.selectnum == 0;
			}
		},
		methods: {
			choose: function(index) {
				this.friends[index].select = !this.friends[index].select;
				this.selectnum += this.friends[index].select * 2 - 1;
				console.log(this.selectnum);
			},
			backOne: function() {
				uni.navigateBack({
					delta: 1
				})
			},
			getFriends: function() {
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
								this.getFriends()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							// console.log('friends')
							if (data.data.Data != null) {
								this.noone = false
								// console.log(data.data.Data)
								this.friends = []
								for (let i = 0; i < data.data.Data.length; i++) {
									let frinick = data.data.Data[i].FriendNick
									if (frinick == '') {
										frinick = data.data.Data[i].FriendName
									}
									this.friends.push({
										id: data.data.Data[i].Friendid,
										name: frinick,
										truname: data.data.Data[i].FriendName,
										img: 'data:image/png;base64,' + data.data.Data[i]
											.FriendImg,
										select: false,
									})
								}
								console.log(this.friends)
							} else {
								this.friends = []
							}

						} else {
							//err
						}
					}
				})
			},
			select() {
				uni.chooseImage({
					count: 1,
					success: rst => {
						// 设置url的值，显示控件
						this.url = rst.tempFilePaths[0]
					}
				})
			},
			oncancel() {
				// url设置为空，隐藏控件
				this.url = ''
			},
			compressImg: function(src) {
				return new Promise(async (resolve, reject) => {
					var ctx = uni.createCanvasContext('cv', this)
					let w, h
					await this.getImgInfo(src).then(res => {
						w = res[0]
						h = res[1]
					})
					let d = this.Afill(w, h, 100, 100)
					ctx.drawImage(src, d[0], d[1], d[2], d[3], 0, 0, 100, 100)
					ctx.draw(false, () => {
						uni.canvasToTempFilePath({
							canvasId: 'cv',
							fileType: 'jpg',
							quality: 0.8,
							success: function(res) {
								// 在H5平台下，tempFilePath 为 base64
								console.log(res.tempFilePath)
								resolve(res.tempFilePath)
							},
							fail: function() {
								reject('error')
							}
						}, this)
					}) //绘制
				})
			},
			getImgInfo(url) {
				return new Promise((resolve, reject) => {
					let image = new Image()
					image.src = url
					let timer = setInterval(() => {
						if (image.width > 0 || image.height > 0) {
							resolve([image.width, image.height]) // 图片宽*高
							clearInterval(timer)
						}
					}, 50)
				})
			},
			Afill: function(imageWidth, imageHeight, canvasWidth, canvasHeight) {
				const imageRate = imageWidth / imageHeight
				const canvasRate = canvasWidth / canvasHeight
				let [sx, sy, sw, sh] = []
				if (imageRate >= canvasRate) {
					sw = imageHeight * canvasRate
					sh = imageHeight
					sx = (imageWidth - sw) / 2
					sy = 0
				} else {
					sh = imageWidth / canvasRate
					sw = imageWidth
					sx = 0
					sy = (imageHeight - sh) / 2
				}
				return [sx, sy, sw, sh]
			},
			async onok(ev) {
				this.img = await this.compressImg(ev.base64)
				console.log(ev.base64)
				this.url = "";
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

		.group-img {
			width: 196rpx;
			height: 196rpx;
			background: $uni-color-primary;
			border-radius: $uni-border-radius-lg;
			box-shadow: 0rpx 36rpx 40rpx 0rpx rgba(39, 40, 50, 0.1);

			image {
				width: 196rpx;
				height: 196rpx;
				background: $uni-color-primary;
				border-radius: $uni-border-radius-lg;
			}
		}

		.group-name {
			padding: 62rpx $uni-spacing-col-base;

			.group-name-input {
				text-align: center;
				width: 682rpx;
				height: 92rpx;
				border-radius: $uni-border-radius-lg;
				background: $uni-bg-color-grey;
				font-size: 32rpx;
				color: $uni-text-color;
			}
		}

		.friends {
			padding: $uni-spacing-col-base 0;
			display: flex;
			flex-direction: column;
			width: 100%;

			.friendline {
				display: flex;
				flex-direction: row;
				padding: 0 $uni-spacing-row-base;
				height: 120rpx;
				align-items: center;

				.friendselect {
					flex: none;
					width: 40rpx;
					height: 40rpx;
					border-radius: 48rpx;
					background-color: $uni-color-primary;
					padding: 10rpx;

					.selectimg {
						width: 48rpx;
						height: 48rpx;
					}
				}

				.friendhead {
					width: $uni-img-size-base;
					height: $uni-img-size-base;
					border-radius: $uni-border-radius-base;
					margin: 0 30rpx;
				}

				.friendname {
					width: 400rpx;
					font-size: 40rpx;
					font-weight: 400;
					color: $uni-text-color;
					display: -webkit-box;
					-webkit-box-orient: vertical;
					-webkit-line-clamp: 1;
					overflow: hidden;
				}
			}

			.title {
				padding: 0 $uni-spacing-row-base;
				font-size: 48rpx;
				font-weight: 520;
				color: $uni-text-color;
			}

		}

		.user-scr {
			height: 800rpx;
		}
	}

	.buttomBig {
		background-color: $uni-bg-color-grey;

		.unselect {
			background-color: $uni-bg-color-hover;
		}

		box-shadow: 0rpx 1rpx 0rpx 0rpx $uni-bg-color-mask;
	}
</style>
