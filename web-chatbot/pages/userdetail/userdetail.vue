<template>
	<view class="content">
		<view class="top-bar">
			<view class="top-bar-left" @tap="backOne">
				<image src="../../static/general/return.png"></image>
			</view>
			<view class="top-bar-center">
				<view class="title">详情</view>
			</view>
		</view>
		<view class="main">
			<view class="line">
				<view class="lineHead">头像</view>
				<view class="lineInfo" v-if="userid != optid">
					<image :src="userDetail.img" mode="aspectFill" class="lineImg"></image>
				</view>
				<view class="lineInfo" @click="select" v-if="userid == optid">
					<image :src="userDetail.img" mode="aspectFill" class="lineImg"></image>
				</view>
				<ksp-cropper mode="free" :width="200" :height="140" :maxWidth="1024" :maxHeight="1024" :url="url"
					@cancel="oncancel" @ok="onok" v-if="userid == optid"></ksp-cropper>
				<view class="lineBtn" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">签名</view>
				<view class="linetext">{{userDetail.intr}}</view>
				<view class="lineBtn" @tap="goModify(2)" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">注册</view>
				<view class="linetext">{{userDetail.createTime}}</view>
			</view>
			<view class="line">
				<view class="lineHead">昵称</view>
				<view class="linetext">{{userDetail.name}}</view>
				<view class="lineBtn" @tap="goModify(0)" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line" v-if="userid != optid">
				<view class="lineHead">备注</view>
				<view class="linetext">{{userDetail.nick}}</view>
				<view class="lineBtn" @tap="goModify(5)">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">性别</view>
				<view class="linetext">
					<view class="uni-list-cell-db">
						<picker @change="bindPickerChange" :value="userDetail.sex" :range="array"
							v-if="userid == optid">
							<view class="uni-input">{{array[userDetail.sex]}}</view>
						</picker>
						<view class="uni-input" v-if="userid != optid">{{array[userDetail.sex]}}</view>
					</view>
				</view>
				<view class="lineBtn" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">生日</view>
				<view class="linetext">
					<view class="uni-list-cell-db">
						<picker mode="date" :value="userDetail.birthday" :start="startDate" :end="endDate"
							@change="bindDateChange" v-if="userid == optid">
							<view class="uni-input">{{userDetail.birthday}}</view>
						</picker>
						<view class="uni-input" v-if="userid != optid">{{userDetail.birthday}}</view>
					</view>
				</view>
				<view class="lineBtn" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">电话</view>
				<view class="linetext">{{userDetail.phone}}</view>
				<view class="lineBtn" @tap="goModify(4)" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line">
				<view class="lineHead">邮箱</view>
				<view class="linetext">{{userDetail.email}}</view>
				<view class="lineBtn" @tap="goModify(3)" v-if="userid == optid">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="line" v-if="userid == optid">
				<view class="lineHead">密码</view>
				<view class="linetext">*********</view>
				<view class="lineBtn" @tap="goModify(1)">
					<image src="../../static/general/more2.png">
				</view>
			</view>
			<view class="btn" v-if="userid == optid" @tap="exit()">
				退出登录
			</view>
			<view class="btn" v-if="userid != optid" @tap="delFriend()">
				删除好友
			</view>
			<canvas style="width: 100px; height: 100px; position:fixed;left:100%;" canvas-id="cv" type="2d"></canvas>
		</view>
	</view>
</template>

<script>
	import ksp from '../../commons/js/ksp.js'
	import config from '../../commons/js/config.js'
	import refersh from '../../commons/js/refershToken.js'
	//import {saveBase64Img,dataURItoBlob} from '../../commons/js/blobImg.js';
	export default {
		data() {
			return {
				user: {
					id: 4,
					imgurl: '../../static/index/1 (4).jpg',
					tip: 3,
					name: '文嘉',
					time: new Date(),
					info: '明日复明日，明日何其多！日日待明日，万事成蹉跎。世人皆被明日累，明日无穷老将至。晨昏滚滚水东流。今古悠悠日西坠。百年明日能几何？请君听我《明日歌》。',
					sex: 0,
					nick: '嘉嘉嘉',
					intr: '明日复明日，明日何其多！日日待明日，万事成蹉跎。世人皆被明日累，明日无穷老将至。晨昏滚滚水东流。今古悠悠日西坠。百年明日能几何？请君听我《明日歌》。',
					email: '123456@12.com',
					phone: '12345678',
				},
				userid: '',
				optid: '',
				userDetail: {},
				atoken: '',
				rtoken: '',
				userNick: '',
				modifyInfo: {
					type: 2, //infoType: ['昵称', '密码', '签名', '邮箱', '电话', '备注', '添加好友']
					userinfo: {}
				},
				array: ['男', '女', '未知'],
				date: this.getDate({ format: true }),
				url: '',
			}
		},
		onLoad: function(option) {
			this.optid = option.id
			try {
				const value = uni.getStorageSync('user')
				if (value) {
					this.userid = value.id
					this.userNick = value.username
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
		computed: {
			startDate() {
				return this.getDate('start')
			},
			endDate() {
				return this.getDate('end')
			}
		},
		methods: {
			exit: function() {
				uni.navigateTo({ url: '../signin/signin' })
			},
			delFriend: function() {
				uni.showModal({
					title: '提示',
					content: '确认要删除好友吗？',
					success: res => {
						console.log(this.atoken)
						if (res.confirm) {
							uni.request({
								url: config.myurl + '/friend/deleate',
								method: 'POST',
								header: { 'Authorization': 'Bearer ' + this.atoken },
								data: {
									'user_id': this.userid,
									'friend_id': this.optid
								},
								success: async data => {
									if (data.data.Code == 1009) {
										let newCode = await refersh.refersh(config.myurl, this
											.atoken, this.rtoken)
										if (newCode == 1000) {
											this.atoken = uni.getStorageSync('atoken')
											this.rtoken = uni.getStorageSync('rtoken')
											this.delFriend()
										} else {
											// err
										}
									} else if (data.data.Code == 1000) {
										//success
										let pages = getCurrentPages() // 当前页面
										let beforePage = pages[pages.length - 2]
										uni.navigateBack({
											delta: 1,
											success: function() {
												beforePage
												.delfriend() // 执行上一页的onLoad方法
											}
										})
									} else {
										//err
									}
								}
							})
						} else if (res.cancel) {
							console.log('用户点击取消')
						}
					}
				})

			},
			backOne: function() {
				let pages = getCurrentPages() // 当前页面
				let beforePage = pages[pages.length - 2]

				beforePage.$vm.refersh(this.userDetail)
				uni.navigateBack({ delta: 1 })
			},
			refersh: function(parm) {
				this.userDetail[parm.type] = parm.req
			},
			getUserDetail: function() {
				console.log(this.optid)
				uni.request({
					url: config.myurl + '/user/detial',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: { 'id': this.optid },
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getUserDetail()
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(data.data.Data)
							this.userDetail = {
								'id': this.optid,
								'img': 'data:image/png;base64,' + data.data.Data.Img,
								'intr': data.data.Data.Intr,
								'createTime': data.data.Data.CreateTime.substring(0, 10),
								'name': data.data.Data.Username,
								'sex': data.data.Data.Gender,
								'birthday': data.data.Data.Birthday.substring(0, 10),
								'phone': data.data.Data.Phone,
								'email': data.data.Data.Email,
								'nick': this.userNick
							}

						} else {
							// err
						}
					}
				})
				console.log(777)
			},
			getUserNick: function() {
				if (this.userid == this.optid) {
					return
				}
				uni.request({
					url: config.myurl + '/serch/isfriend',
					method: 'POST',
					header: { 'Authorization': 'Bearer ' + this.atoken },
					data: {
						'user_id': this.userid,
						'friend_id': this.optid
					},
					success: async data => {
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
							console.log(data.data.Data)
							if (data.data.Data.State == 0) {
								if (data.data.Data.Markname.length != 0) {
									this.userNick = data.data.Data.Markname
								} else {
									this.userNick = this.user.name
								}
								this.userDetail.nick = this.userNick
							} else {
								this.userNick = this.user.name
							}
						} else {
							// err
						}
					}
				})
			},
			bindPickerChange: function(e) {
				if (this.userDetail.sex == e.detail.value) {
					return
				}
				var fn = () => {
					this.userDetail['sex'] = e.detail.value
				}
				this.update('gender', this.optid, e.detail.value.toString(), '', fn)

				// this.userDetail.sex = e.detail.value
			},
			bindDateChange: function(e) {
				if (this.userDetail.birthday == e.detail.value) {
					return
				}
				var fn = () => {
					this.userDetail['birthday'] = e.detail.value
				}
				this.update('birthday', this.optid, e.detail.value.toString(), '', fn)
			},
			getDate(type) {
				const date = new Date()
				let year = date.getFullYear()
				let month = date.getMonth() + 1
				let day = date.getDate()

				if (type === 'start') {
					year = year - 60
				} else if (type === 'end') {
					year = year + 2
				}
				month = month > 9 ? month : '0' + month
				day = day > 9 ? day : '0' + day
				return `${year}-${month}-${day}`
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
			async onok(ev) {
				this.url = ''
				console.log(this.userDetail)
				var fn = () => {
					console.log(2020)
					try {
						let value = uni.getStorageSync('user')
						if (value) {
							value.img = ev.base64.substring(22)
							this.userDetail['img'] = ev.base64
							uni.setStorageSync('user', value)
						} else {
							uni.navigateTo({ url: '../signin/signin', })
						}
						console.log(value)
					} catch (e) {
						console.log(e)
					}
				}
				let res= await this.compressImg(ev.base64)
				this.update('imgBase64', this.optid, res.substring(22), '', fn)
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
			compressImg: function(src){
				return new Promise(async (resolve, reject)=> {
					var ctx = uni.createCanvasContext('cv',this)
					let w, h
					await this.getImgInfo(src).then(res => {
						w = res[0]
						h = res[1]
					})
					let d = this.Afill(w, h, 100, 100)
					ctx.drawImage(src, d[0], d[1], d[2], d[3], 0, 0, 100, 100)
					ctx.draw(false,()=>{
						uni.canvasToTempFilePath({
						  canvasId: 'cv',
						  fileType: 'jpg',
						  quality: 0.8,
						  success: function(res) {
						    // 在H5平台下，tempFilePath 为 base64
							console.log(res.tempFilePath)
							resolve(res.tempFilePath)
						  } ,
						  fail: function() {
							  reject('error')
						  }
						},this)
					}) //绘制
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
			oncancel() {
				// url设置为空，隐藏控件
				this.url = ''
			},
			goModify: function(type) {
				this.modifyInfo.type = type
				this.modifyInfo.userinfo = this.userDetail
				uni.navigateTo({ url: '../modify/modify?modifyInfo=' + encodeURIComponent(JSON.stringify(this.modifyInfo)) })
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
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.update(type, optid, data, psw, fn)
							} else {
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(1010)
							fn.apply()
						} else {
							//err
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

		.line {
			margin-top: 12rpx;
			width: 100%;
			height: 112rpx;
			display: flex;
			flex-direction: cow;
			align-items: center;

			.lineHead {
				flex: none;
				padding-left: $uni-spacing-col-base;
				font-size: $uni-font-size-lg;
				color: $uni-text-color;
				line-height: 112rpx;
			}

			.linetext {
				flex: auto;
				font-size: $uni-font-size-lg;
				color: $uni-text-color-grey;
				line-height: 112rpx;
				overflow: hidden;
				text-overflow: ellipsis;
				white-space: nowrap;
				padding: 0rpx $uni-spacing-col-base;
			}

			.lineInfo {
				flex: auto;
				display: flex;
				flex-direction: cow;
				align-items: center;
				padding: 0rpx $uni-spacing-col-base;

				.lineImg {
					width: $uni-img-size-lg;
					height: $uni-img-size-lg;
					border-radius: $uni-border-radius-base;
				}
			}

			.lineBtn {
				flex: none;
				padding-right: $uni-spacing-col-base;

				image {
					width: 50rpx;
					height: 50rpx;
				}
			}
		}

		.btn {
			margin-top: 148rpx;
			font-size: $uni-font-size-lg;
			color: $uni-color-warning;
		}
	}
</style>
