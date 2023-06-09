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
	import {
		pathToBase64,
		base64ToPath
	} from '../../commons/js/imgDel.js'
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
				date: this.getDate({
					format: true
				}),
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
					uni.navigateTo({
						url: '../signin/signin',
					})
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
				uni.navigateTo({
					url: '../signin/signin'
				})
			},
			// 删除好友
			delFriend: function() {
				// 显示一个模态框，询问用户是否确认删除好友
				uni.showModal({
					title: '提示',
					content: '确认要删除好友吗？',
					success: res => {
						console.log(this.atoken)
						if (res.confirm) {
							// 如果用户确认删除好友
							uni.request({
								url: config.myurl + '/friend/deleate', // 请求的URL地址
								method: 'POST', // 请求方法为POST
								header: {
									'Authorization': 'Bearer ' + this
										.atoken // 设置请求头中的Authorization字段，使用当前的用户访问令牌
								},
								data: {
									'user_id': this.userid, // 当前用户的ID
									'friend_id': this.optid // 要删除的好友的ID
								},
								success: async data => {
									if (data.data.Code == 1009) {
										// 如果返回的Code为1009，表示访问令牌过期
										let newCode = await refersh.refersh(config.myurl, this
											.atoken, this.rtoken)
										if (newCode == 1000) {
											// 刷新令牌成功，更新访问令牌和刷新令牌，并重新执行删除好友操作
											this.atoken = uni.getStorageSync('atoken')
											this.rtoken = uni.getStorageSync('rtoken')
											this.delFriend()
										} else {
											// 刷新令牌失败，处理错误
											// err
										}
									} else if (data.data.Code == 1000) {
										// 如果返回的Code为1000，表示删除好友成功
										let pages = getCurrentPages() // 获取当前页面栈
										let beforePage = pages[pages.length - 2] // 获取上一个页面
										uni.navigateBack({
											delta: 1, // 返回上一个页面
											success: function() {
												beforePage
													.delfriend() // 执行上一页的delfriend方法
											}
										})
									} else {
										// 返回的Code既不是1009也不是1000，处理错误
										// err
									}
								}
							})
						} else if (res.cancel) {
							// 如果用户取消删除好友操作
							console.log('用户点击取消')
						}
					}
				})
			},

			backOne: function() {
				let pages = getCurrentPages() // 获取当前页面栈
				let beforePage = pages[pages.length - 2] // 获取上一个页面

				beforePage.$vm.refersh(this.userDetail) // 调用上一个页面的$vm对象的refersh方法，将this.userDetail作为参数传递进去
				uni.navigateBack({
					delta: 1 // 返回上一个页面
				})
			},

			refersh: function(parm) {
				this.userDetail[parm.type] = parm.req // 更新this.userDetail对象的某个属性，使用parm.type作为属性名，parm.req作为属性值
			},

			// 获取用户详情
			getUserDetail: function() {
				console.log(this.optid)
				uni.request({
					url: config.myurl + '/user/detial', // 请求的URL地址
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken // 设置请求头中的Authorization字段，使用当前的用户访问令牌
					},
					data: {
						'id': this.optid // 请求参数，指定用户ID
					},
					success: async data => {
						console.log(data.data)
						if (data.data.Code == 1009) {
							// 如果返回的Code为1009，表示访问令牌过期
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								// 刷新令牌成功，更新访问令牌和刷新令牌，并重新执行获取用户详情操作
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getUserDetail()
							} else {
								// 刷新令牌失败，处理错误
								// err
							}
						} else if (data.data.Code == 1000) {
							// 如果返回的Code为1000，表示获取用户详情成功
							console.log(data.data.Data)
							this.userDetail = {
								'id': this.optid, // 用户ID
								'img': 'data:image/png;base64,' + data.data.Data
									.Img, // 用户头像，使用Base64编码的格式
								'intr': data.data.Data.Intr, // 用户简介
								'createTime': data.data.Data.CreateTime.substring(0,
									10), // 用户创建时间，只取日期部分
								'name': data.data.Data.Username, // 用户名
								'sex': data.data.Data.Gender, // 用户性别
								'birthday': data.data.Data.Birthday.substring(0, 10), // 用户生日，只取日期部分
								'phone': data.data.Data.Phone, // 用户电话号码
								'email': data.data.Data.Email, // 用户邮箱
								'nick': this.userNick // 用户昵称
							}

						} else {
							// 返回的Code既不是1009也不是1000，处理错误
							// err
						}
					}
				})
				console.log(777)
			},

			// 获取用户昵称
			getUserNick: function() {
				if (this.userid == this.optid) {
					// 如果当前用户ID和目标用户ID相同，直接返回，不进行后续操作
					return
				}
				uni.request({
					url: config.myurl + '/serch/isfriend', // 请求的URL地址
					method: 'POST', // 请求方法为POST
					header: {
						'Authorization': 'Bearer ' + this.atoken // 设置请求头中的Authorization字段，使用当前的用户访问令牌
					},
					data: {
						'user_id': this.userid, // 当前用户的ID
						'friend_id': this.optid // 目标用户的ID
					},
					success: async data => {
						if (data.data.Code == 1009) {
							// 如果返回的Code为1009，表示访问令牌过期
							let newCode = await refersh.refersh(config.myurl, this.atoken, this.rtoken)
							if (newCode == 1000) {
								// 刷新令牌成功，更新访问令牌和刷新令牌，并重新执行获取用户昵称操作
								this.atoken = uni.getStorageSync('atoken')
								this.rtoken = uni.getStorageSync('rtoken')
								this.getUserNick()
							} else {
								// 刷新令牌失败，处理错误
								// err
							}
						} else if (data.data.Code == 1000) {
							console.log(data.data.Data)
							if (data.data.Data.State == 0) {
								// 如果State为0，表示是好友关系
								if (data.data.Data.Markname.length != 0) {
									// 如果Markname不为空，使用Markname作为用户昵称
									this.userNick = data.data.Data.Markname
								} else {
									// 如果Markname为空，使用用户名作为用户昵称
									this.userNick = this.user.name
								}
								this.userDetail.nick = this.userNick // 更新this.userDetail对象的nick属性
							} else {
								// 如果State不为0，表示非好友关系，使用用户名作为用户昵称
								this.userNick = this.user.name
							}
						} else {
							// 返回的Code既不是1009也不是1000，处理错误
							// err
						}
					}
				})
			},

			// 性别更改
			bindPickerChange: function(e) {
				if (this.userDetail.sex == e.detail.value) {
					// 如果选择的性别和当前性别相同，直接返回，不进行后续操作
					return
				}
				var fn = () => {
					this.userDetail['sex'] = e.detail.value // 更新this.userDetail对象的sex属性
				}
				this.update('gender', this.optid, e.detail.value.toString(), '', fn)
				// 调用update方法，更新用户性别，传入参数为更新类型、目标用户ID、要更新的值、空字符串和回调函数fn

				// this.userDetail.sex = e.detail.value
			},

			// 生日更改
			bindDateChange: function(e) {
				if (this.userDetail.birthday == e.detail.value) {
					// 如果选择的生日和当前生日相同，直接返回，不进行后续操作
					return
				}
				var fn = () => {
					this.userDetail['birthday'] = e.detail.value // 更新this.userDetail对象的birthday属性
				}
				this.update('birthday', this.optid, e.detail.value.toString(), '', fn)
				// 调用update方法，更新用户生日，传入参数为更新类型、目标用户ID、要更新的值、空字符串和回调函数fn
			},

			// 获取日期
			getDate(type) {
				const date = new Date() // 创建一个Date对象，表示当前时间
				let year = date.getFullYear() // 获取当前年份
				let month = date.getMonth() + 1 // 获取当前月份（注意月份从0开始，需要加1）
				let day = date.getDate() // 获取当前日期

				if (type === 'start') {
					// 如果传入的type为'start'，表示需要获取起始日期
					year = year - 60 // 将年份减去60，得到起始年份
				} else if (type === 'end') {
					// 如果传入的type为'end'，表示需要获取结束日期
					year = year + 2 // 将年份加上2，得到结束年份
				}
				month = month > 9 ? month : '0' + month // 如果月份小于10，补0使其两位数表示
				day = day > 9 ? day : '0' + day // 如果日期小于10，补0使其两位数表示
				return `${year}-${month}-${day}` // 返回格式为'YYYY-MM-DD'的日期字符串
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
				let res = await this.compressImg(ev.path)
				console.log(res)
				var fn = () => {
					console.log(2020)
					try {
						let value = uni.getStorageSync('user')
						if (value) {
							value.img = res.substring(23)
							this.userDetail['img'] = 'data:image/png;base64,' + res.substring(23)
							uni.setStorageSync('user', value)
						} else {
							uni.navigateTo({
								url: '../signin/signin',
							})
						}
						console.log(value)
					} catch (e) {
						console.log(e)
					}
				}
				this.update('imgBase64', this.optid, res.substring(23), '', fn)
			},
			// getImgInfo(url) {
			// 	return new Promise((resolve, reject) => {
			// 		let image = new Image()
			// 		image.src = url
			// 		let timer = setInterval(() => {
			// 			if (image.width > 0 || image.height > 0) {
			// 				resolve([image.width, image.height]) // 图片宽*高
			// 				clearInterval(timer)
			// 			}
			// 		}, 50)
			// 	})
			// },
			// 获取图片
			getImgInfo(url) {
				return new Promise((resolve, reject) => {
					wx.getImageInfo({
						src: url, // 图片的URL地址
						success: function(res) {
							resolve([res.width, res.height]) // 将获取到的图片宽度和高度作为Promise的解析结果进行resolve
						},
						fail: function(res) {
							reject(res) // 若获取图片信息失败，则将错误信息作为Promise的拒绝结果进行reject
						}
					})
				})
			},

			// 压缩图片
			compressImg: function(src) {
				return new Promise(async (resolve, reject) => {
					console.log(src)
					var ctx = uni.createCanvasContext('cv', this) // 创建一个canvas绘图上下文对象
					let w, h
					await this.getImgInfo(src).then(res => {
						w = res[0] // 获取图片的宽度
						h = res[1] // 获取图片的高度
					})
					let d = this.Afill(w, h, 100, 100) // 调用Afill方法计算绘制图片的位置和尺寸
					ctx.drawImage(src, d[0], d[1], d[2], d[3], 0, 0, 100, 100) // 在canvas上绘制指定区域的图片
					ctx.draw(false, () => {
						uni.canvasToTempFilePath({
							canvasId: 'cv', // canvas的ID
							fileType: 'jpg', // 保存的文件类型
							quality: 0.8, // 图片质量，取值范围为0到1
							success: function(res) {
								// #ifdef H5
								// 在H5平台下，tempFilePath为base64格式
								console.log(res)
								resolve(res.tempFilePath) // 将压缩后的图片路径进行resolve
								// #endif
								// #ifndef H5
								pathToBase64(res.tempFilePath)
									.then(base64 => {
										console.log(base64)
										resolve(base64) // 将压缩后的图片base64数据进行resolve
									})
									.catch(error => {
										console.error(error)
										reject(
											error
										) // 如果转换为base64失败，将错误信息进行reject
									})
								// #endif
							},
							fail: function() {
								reject(
									'error'
								) // 如果canvasToTempFilePath失败，将错误信息进行reject
							}
						}, this)
					}) // 绘制canvas
				})
			},

			Afill: function(imageWidth, imageHeight, canvasWidth, canvasHeight) {
				const imageRate = imageWidth / imageHeight // 计算原始图片的宽高比
				const canvasRate = canvasWidth / canvasHeight // 计算目标画布的宽高比
				let [sx, sy, sw, sh] = [] // 定义变量用于记录绘制图片的位置和尺寸

				if (imageRate >= canvasRate) { // 如果原始图片的宽高比大于等于目标画布的宽高比
					sw = imageHeight * canvasRate // 计算绘制图片的宽度
					sh = imageHeight // 绘制图片的高度与原始图片的高度相同
					sx = (imageWidth - sw) / 2 // 计算绘制图片的起始位置横坐标
					sy = 0 // 绘制图片的起始位置纵坐标为0
				} else { // 如果原始图片的宽高比小于目标画布的宽高比
					sh = imageWidth / canvasRate // 计算绘制图片的高度
					sw = imageWidth // 绘制图片的宽度与原始图片的宽度相同
					sx = 0 // 绘制图片的起始位置横坐标为0
					sy = (imageHeight - sh) / 2 // 计算绘制图片的起始位置纵坐标
				}

				return [sx, sy, sw, sh] // 返回绘制图片的位置和尺寸数组
			},

			oncancel() {
				// url设置为空，隐藏控件
				this.url = ''
			},
			// 修改页面
			goModify: function(type) {
				this.modifyInfo.type = type
				this.modifyInfo.userinfo = this.userDetail
				uni.navigateTo({
					url: '../modify/modify?modifyInfo=' + encodeURIComponent(JSON.stringify(this.modifyInfo))
				})
			},
			// 信息更新
			update(type, optid, data, psw, fn) {
				console.log(fn)
				uni.request({
					url: config.myurl + '/change/update',
					method: 'POST',
					header: {
						'Authorization': 'Bearer ' + this.atoken
					},
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
